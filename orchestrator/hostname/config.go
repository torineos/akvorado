package hostname

import (
	"akvorado/common/helpers"
)

// Configuration describes the configuration for the Hostname component.
type Configuration struct {
	// HostnameEnabled tells if we want to collect Hostname data
	HostnameToggle bool
}

// DefaultConfiguration represents the default configuration for the Hostname component.
// If the reverse resolution isn't enabled, the component will return nothing.
func DefaultHostnameConfiguration() Configuration {
	return Configuration{
		// Disabled by default to preserve resources
		HostnameToggle: false,
	}
}

func init() {
	helpers.RegisterMapstructureUnmarshallerHook(
		helpers.DefaultValuesUnmarshallerHook(DefaultHostnameConfiguration()))
}
