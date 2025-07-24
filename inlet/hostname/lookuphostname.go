package hostname

import (
	"net"
	"net/netip"
)

// LookupHostname makes a request to your DNS server
// (listed in your resolv.conf) and returns the hostname
// of the provided IP address
func LookupHostname(netip netip.Addr) string {
	names, err := net.LookupAddr(netip.String())
	if err == nil {
		//fmt.Println(names)
		// return only the first hostname
		return names[0]
	}
	// if no host is returned, return n/a
	return "n/a"
}
