package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr error
	}{
		{
			name: "Valid API key",
			headers: http.Header{
				"Authorization": {"ApiKey abcdef12345"},
			},
			want:    "abcdef12345",
			wantErr: nil,
		},
		{
			name:    "No Authorization header",
			headers: http.Header{},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
		{
			name: "Malformed Authorization header",
			headers: http.Header{
				"Authorization": {"Bearer abcdef12345"},
			},
			want:    "",
			wantErr: errors.New("malformed authorization header"),
		},
		{
			name: "Empty Authorization header",
			headers: http.Header{
				"Authorization": {""},
			},
			want:    "",
			wantErr: ErrNoAuthHeaderIncluded,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if got != tt.want {
				t.Errorf("GetAPIKey() got = %v, want = %v", got, tt.want)
			}
			if (err != nil && tt.wantErr != nil && err.Error() != tt.wantErr.Error()) || (err != nil && tt.wantErr == nil) || (err == nil && tt.wantErr != nil) {
				t.Errorf("GetAPIKey() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
