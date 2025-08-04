package hostname

import (
	"fmt"
	"net/netip"

	"testing"
)

func TestRetrieveHost(t *testing.T) {
	// cases here
	var tests = []struct {
		name  string
		input netip.Addr
		want  string
	}{
		// the table itself
		{"8.8.4.4 should be dns.google.", netip.MustParseAddr("8.8.4.4"), "dns.google."},                           // Ipv4 Google public DNS
		{"2001:4860:4860::8888 should be dns.google.", netip.MustParseAddr("2001:4860:4860::8888"), "dns.google."}, // Ipv6 Google public DNS.
		{"192.168.40.56 should be n/a", netip.MustParseAddr("192.168.40.56"), "n/a"},                               // Try to use an unused ip w/ no host resolution.
	}
	// execution loop
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := LookupHostname(tt.input)
			fmt.Println("ans", ans, "tt.want", tt.want)
			if ans != tt.want {
				t.Errorf("Hostname() returned %s, expected ", tt.want)
			}
		})
	}
}
