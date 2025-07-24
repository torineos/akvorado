// Package hostname is dedicated to the reverse resolution of an ip address
package hostname

import (
	"reflect"

	"akvorado/common/helpers"

	"github.com/go-viper/mapstructure/v2"
)

// Configuration describes the configuration for the Hostname component.
type Configuration struct {
	// Enable specifies if the component is enabled or not.
	Enable bool `yaml:",omitempty"`
}

// DefaultConfiguration represents the default configuration for the Hostname component.
func DefaultConfiguration() Configuration {
	return Configuration{
		// Enable is disabled by default
		Enable: false,
	}
}

// ConfigurationUnmarshallerHook normalize core configuration because I have problems
func ConfigurationUnmarshallerHook() mapstructure.DecodeHookFunc {
	return func(from, to reflect.Value) (interface{}, error) {
		if from.Kind() != reflect.String || to.Type() != reflect.TypeOf(Configuration{}) {
			return from.Interface(), nil
		}
		return from.Interface(), nil
	}

}

// Default value is being unmarshalled (from yaml structure to go structure). Registers the default configuration of hostname I believe
func init() {
	helpers.RegisterMapstructureUnmarshallerHook(ConfigurationUnmarshallerHook())
}
