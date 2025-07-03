package helpers

import (
	"fmt"
	"net"
	"net/netip"

)

func Hostname(netip netip.Addr) string {
	// define context background
	names, err := net.LookupAddr(netip.String())
	if err == nil {
		fmt.Println(names)
		// return only the first hostname
		return names[0]
	}
	// if no host is returned, return n/a
	return "n/a"
}
