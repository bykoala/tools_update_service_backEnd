package settings

import (
	"testing"
)

func TestInit(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"success",
			false,
		},
		{
			"failed",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Init("config/config.ymal"); (err != nil) != tt.wantErr && tt.name == "success" {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := Init("config/config.ymal"); (err != nil) == tt.wantErr && tt.name == "failed" {
				t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
