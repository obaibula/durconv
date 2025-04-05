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
			name: "Valid h-d",
			fmt:  "h-d",
			dur:  50*durfmt.Day + 4*durfmt.Hour,
			want: "4h-50d",
		},
		{
			name: "Valid M_d_h",
			fmt:  "M_d_h",
			dur:  5*durfmt.Month + 30*durfmt.Day + durfmt.Hour,
			want: "5M_30d_1h",
		},
		{
			name: "Valid h-d-M",
			fmt:  "h-d-M",
			dur:  500*durfmt.Month + durfmt.Day + 12*durfmt.Hour,
			want: "12h-1d-500M",
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
