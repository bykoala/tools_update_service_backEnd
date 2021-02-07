package tools

import (
	"net"
	"reflect"
	"testing"
)

func TestGetOutboundIP(t *testing.T) {
	tests := []struct {
		name string
		want net.IP
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOutboundIP(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetOutboundIP() = %v, want %v", got, tt.want)
			}
		})
	}
}
