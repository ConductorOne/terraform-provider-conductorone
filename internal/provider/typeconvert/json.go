package typeconvert

import (
	"encoding/json"
)

// MarshalJSONStripNulls marshals the given value to JSON, then strips any null
// values from JSON objects (recursively). This is useful for normalizing JSON
// from API responses where the server includes explicit null values for
// optional fields, but the Terraform configuration omits those fields entirely.
// Without stripping, jsontypes.Normalized would see these as different values,
// causing "Provider produced inconsistent result after apply" errors.
func MarshalJSONStripNulls(v interface{}) ([]byte, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	var raw interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return nil, err
	}

	stripped := stripNulls(raw)
	return json.Marshal(stripped)
}

// stripNulls recursively removes null values from JSON objects.
// Arrays are preserved as-is (null elements remain), but object keys
// with null values are removed.
func stripNulls(v interface{}) interface{} {
	switch val := v.(type) {
	case map[string]interface{}:
		result := make(map[string]interface{}, len(val))
		for k, v := range val {
			if v == nil {
				continue
			}
			result[k] = stripNulls(v)
		}
		return result
	case []interface{}:
		result := make([]interface{}, len(val))
		for i, item := range val {
			result[i] = stripNulls(item)
		}
		return result
	default:
		return v
	}
}
