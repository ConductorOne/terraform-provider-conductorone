package typeconvert

import (
	"encoding/json"
	"testing"
)

func TestMarshalJSONStripNulls(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "strips null values from objects",
			input:    `{"assignee":null,"instructions":"repro","userIds":["abc"]}`,
			expected: `{"instructions":"repro","userIds":["abc"]}`,
		},
		{
			name:     "strips nested null values",
			input:    `{"provisionSteps":[{"manual":{"assignee":null,"instructions":"repro"}}]}`,
			expected: `{"provisionSteps":[{"manual":{"instructions":"repro"}}]}`,
		},
		{
			name:     "preserves non-null values",
			input:    `{"appId":"123","implicit":false}`,
			expected: `{"appId":"123","implicit":false}`,
		},
		{
			name:     "handles empty objects",
			input:    `{}`,
			expected: `{}`,
		},
		{
			name:     "handles objects with all null values",
			input:    `{"a":null,"b":null}`,
			expected: `{}`,
		},
		{
			name:     "preserves null in arrays",
			input:    `{"items":[null,"a","b"]}`,
			expected: `{"items":[null,"a","b"]}`,
		},
		{
			name: "multi-step entitlement scenario",
			input: `{"provisionSteps":[{"delegated":{"appId":"abc","entitlementId":"def","implicit":false}},{"manual":{"assignee":null,"instructions":"repro","userIds":["ghi"]}}]}`,
			expected: `{"provisionSteps":[{"delegated":{"appId":"abc","entitlementId":"def","implicit":false}},{"manual":{"instructions":"repro","userIds":["ghi"]}}]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var input interface{}
			if err := json.Unmarshal([]byte(tt.input), &input); err != nil {
				t.Fatalf("failed to unmarshal input: %v", err)
			}

			result, err := MarshalJSONStripNulls(input)
			if err != nil {
				t.Fatalf("MarshalJSONStripNulls returned error: %v", err)
			}

			// Normalize expected for comparison (to handle key ordering)
			var expectedRaw, resultRaw interface{}
			if err := json.Unmarshal([]byte(tt.expected), &expectedRaw); err != nil {
				t.Fatalf("failed to unmarshal expected: %v", err)
			}
			if err := json.Unmarshal(result, &resultRaw); err != nil {
				t.Fatalf("failed to unmarshal result: %v", err)
			}

			expectedBytes, _ := json.Marshal(expectedRaw)
			resultBytes, _ := json.Marshal(resultRaw)

			if string(expectedBytes) != string(resultBytes) {
				t.Errorf("expected %s, got %s", string(expectedBytes), string(resultBytes))
			}
		})
	}
}
