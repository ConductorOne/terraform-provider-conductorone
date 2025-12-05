package sdk

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/base64"
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/go-jose/go-jose/v3"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/zap/ctxzap"
	"github.com/pquerna/xjwt"
	"golang.org/x/oauth2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/square/go-jose.v2/json"
	"gopkg.in/square/go-jose.v2/jwt"

	"github.com/conductorone/terraform-provider-conductorone/internal/sdk/uhttp"
)

const (
	assertionType = "urn:ietf:params:oauth:client-assertion-type:jwt-bearer"
)

var (
	ErrInvalidClientSecret = errors.New("invalid client secret")
	ErrInvalidClientID     = errors.New("invalid client id")

	v1SecretTokenIdentifier = []byte("v1")
)

type c1Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Expiry      int    `json:"expires_in"`
}

type c1TokenSource struct {
	clientID     string
	clientSecret *jose.JSONWebKey
	tokenHost    string
	httpClient   *http.Client
}

func parseSecret(input []byte) (*jose.JSONWebKey, error) {
	items := bytes.SplitN(input, []byte(":"), 4)
	if len(items) != 4 {
		return nil, ErrInvalidClientSecret
	}

	if !bytes.Equal(items[2], v1SecretTokenIdentifier) {
		return nil, ErrInvalidClientSecret
	}

	jwkData, err := base64.RawURLEncoding.DecodeString(string(items[3]))
	if err != nil {
		return nil, ErrInvalidClientSecret
	}

	npk := &jose.JSONWebKey{}
	err = npk.UnmarshalJSON(jwkData)
	if err != nil {
		return nil, ErrInvalidClientSecret
	}

	if npk.IsPublic() || !npk.Valid() {
		return nil, ErrInvalidClientSecret
	}

	_, ok := npk.Key.(ed25519.PrivateKey)
	if !ok {
		return nil, ErrInvalidClientSecret
	}

	return npk, nil
}

func (c *c1TokenSource) Token() (*oauth2.Token, error) {
	jsigner, err := jose.NewSigner(
		jose.SigningKey{
			Algorithm: jose.EdDSA,
			Key:       c.clientSecret,
		},
		&jose.SignerOptions{
			NonceSource: &xjwt.RandomNonce{Size: 16},
		})
	if err != nil {
		return nil, err
	}

	// Our token host may include a port, but the audience never expects a port
	aud := c.tokenHost
	host, _, err := net.SplitHostPort(aud)
	if err == nil {
		aud = host
	}
	now := time.Now()
	claims := &jwt.Claims{
		Issuer:    c.clientID,
		Subject:   c.clientID,
		Audience:  jwt.Audience{aud},
		Expiry:    jwt.NewNumericDate(now.Add(time.Minute * 2)),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now.Add(-time.Minute * 2)),
	}
	b, err := json.Marshal(claims)
	if err != nil {
		return nil, err
	}

	rv, err := jsigner.Sign(b)
	if err != nil {
		return nil, err
	}

	s, err := rv.CompactSerialize()
	if err != nil {
		return nil, err
	}

	body := url.Values{
		"client_id":             []string{c.clientID},
		"grant_type":            []string{"client_credentials"},
		"client_assertion_type": []string{assertionType},
		"client_assertion":      []string{s},
	}

	tokenUrl := url.URL{
		Scheme: "https",
		Host:   c.tokenHost,
		Path:   "auth/v1/token",
	}

	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, tokenUrl.String(), strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, status.Errorf(codes.Unauthenticated, "failed to get token: %s", resp.Status)
	}

	c1t := &c1Token{}
	err = json.NewDecoder(resp.Body).Decode(c1t)
	if err != nil {
		return nil, err
	}

	if c1t.AccessToken == "" {
		return nil, status.Errorf(codes.Unauthenticated, "failed to get token: empty access token")
	}

	return &oauth2.Token{
		AccessToken: c1t.AccessToken,
		TokenType:   c1t.TokenType,
		Expiry:      time.Now().Add(time.Duration(c1t.Expiry) * time.Second),
	}, nil
}

func NewTokenSource(ctx context.Context, clientID string, clientSecret string, tokenHost string) (oauth2.TokenSource, error) {
	secret, err := parseSecret([]byte(clientSecret))
	if err != nil {
		return nil, err
	}

	httpClient, err := uhttp.NewClient(ctx, uhttp.WithLogger(true, ctxzap.Extract(ctx)), uhttp.WithUserAgent("cone-c1-credential-provider"))
	if err != nil {
		return nil, err
	}
	return oauth2.ReuseTokenSource(nil, &c1TokenSource{
		clientID:     clientID,
		clientSecret: secret,
		tokenHost:    strings.TrimPrefix(tokenHost, "https://"),
		httpClient:   httpClient,
	}), nil
}

func NewStaticTokenSource(ctx context.Context, token string) oauth2.TokenSource {
	return oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: token,
	})
}
