package durfmt_test

import (
	"strings"
	"testing"
	"time"

	"github.com/obaibula/durfmt"
)

func TestString(t *testing.T) {
	tests := []struct {
		name       string
		fmt        string
		dur        time.Duration
		want       string
		wantErrMsg string
	}{
		{
			name: "Valid d-h",
			fmt:  "d-h",
			dur:  4*durfmt.Day + 4*durfmt.Hour,
			want: "4d-4h",
		},
		{
			name:       "Invalid precedense h-d",
			fmt:        "h-d",
			wantErrMsg: "invalid precedense",
		},
		{
			name: "Valid M-d-h",
			fmt:  "m-d-h",
			dur:  5*durfmt.Month + 30*durfmt.Day + durfmt.Hour,
			want: "5m-30d-1h",
		},
		{
			name:       "Invalid precedense d-h-m",
			fmt:        "d-h-m",
			wantErrMsg: "invalid precedense",
		},
		{
			name:       "Invalid precedense h-d-m",
			fmt:        "h-d-m",
			wantErrMsg: "invalid precedense",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := durfmt.String(tt.fmt, tt.dur)
			if tt.wantErrMsg != "" {
				if err == nil {
					t.Fatalf("got nil, want error: %q", tt.wantErrMsg)
				}
				if !strings.Contains(err.Error(), tt.wantErrMsg) {
					t.Fatalf("got error: %q, want error: %q", err, tt.wantErrMsg)
				}
			} else {
				if err != nil {
					t.Fatalf("expected no errors, got: %q", err)
				}
			}

			if got != tt.want {
				t.Errorf("got: %q, want: %q", got, tt.want)
			}
		})
	}
}
