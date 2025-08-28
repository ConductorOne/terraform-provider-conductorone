# StepUpOAuth2Settings

StepUpOAuth2Settings repersents an OAuth2 provider that supports RFC 9470 <https://www.rfc-editor.org/rfc/rfc9470>

 Common ACR values for OAuth2 providers include:
   - "urn:okta:loa:1fa:any" (okta)
   - "urn:okta:loa:1fa:pwd" (okta)
   - "urn:okta:loa:2fa:any" (okta)
   - "urn:okta:loa:2fa:any:ifpossible" (okta)
   - "phr" (okta)
   - "phrh" (okta)


## Fields

| Field                | Type                 | Required             | Description          |
| -------------------- | -------------------- | -------------------- | -------------------- |
| `AcrValues`          | []*string*           | :heavy_minus_sign:   | The acrValues field. |