// random package comment here
// package dedicated to the reverse resolution of an ip address

package hostname

import (
	"net/netip"

	"akvorado/common/daemon"
	"akvorado/common/reporter"
	//"akvorado/common/schema"
)

// Component represents Hostname component
type Component struct {
	r      *reporter.Reporter
	config Configuration
	d      *Dependencies

	// Host represents the structure of the hostname component
	Host struct {
		// Defines if the component is enabled or not
		Enable bool
		// Hostname of the machine.
		IPAddr string
	}
}

// Dependencies define the dependencies of the Kafka exporter.
type Dependencies struct {
	Daemon daemon.Component
	//Schema *schema.Component
}

// New creates the hostname component.
func New(r *reporter.Reporter, configuration Configuration, dependencies Dependencies) (*Component, error) {
	c := Component{
		r:      r,
		config: configuration,
		d:      &dependencies,
	}
	return &c, nil
}

// ReverseResolver checks if the seach is enabled. If yes,
// calls ./lookuphostname.go
// If no, ends the func there
func (c *Component) ReverseResolver(IPAddr netip.Addr) string {
	if !c.config.Enable {
		c.r.Warn().Msg("skipping hostname component: component disabled")
		c.Host.IPAddr = "n/a"
		return "n/a"

	}
	c.r.Info().Msg("starting hostname resolution component")
	result := LookupHostname(IPAddr) // calls lookuphostname.go
	// The actual part where its supposed to be solving
	c.Host.IPAddr = result
	//fmt.Print("Test dans le root de hostname", result)
	return result
}
