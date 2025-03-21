package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "")
	testCases := []struct {
		authorization_header string
		expected             string
		expectedErr          error
	}{{"", "", ErrNoAuthHeaderIncluded}}

	for _, test := range testCases {
		t.Run("123", func(t *testing.T) {
			header.Set("Authorization", test.authorization_header)
			got, gotErr := GetAPIKey(header)
			if !reflect.DeepEqual(test.expected, got) {
				t.Fatalf("expected: %v, got: %v", test.expected, got)
			}
			if !reflect.DeepEqual(test.expectedErr, gotErr) {
				t.Fatalf("expected: %v, got: %v", test.expectedErr, gotErr)
			}
		})

	}
}
