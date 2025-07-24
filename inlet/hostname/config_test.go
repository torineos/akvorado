package hostname

import (
	"testing"

	"akvorado/common/helpers"

	"github.com/gin-gonic/gin"
)

func TestConfigurationUnmarshallerHook(t *testing.T) {
	helpers.TestConfigurationDecode(t, helpers.ConfigurationDecodeCases{
		{
			Description:    "nil",
			Initial:        func() interface{} { return Configuration{} },
			Configuration:  func() interface{} { return nil },
			Expected:       Configuration{},
			SkipValidation: true,
		},
		{
			Description:    "empty",
			Initial:        func() interface{} { return Configuration{} },
			Configuration:  func() interface{} { return gin.H{} },
			Expected:       Configuration{},
			SkipValidation: true,
		}, {
			Description: "enable = false",
			Initial:     func() interface{} { return Configuration{} },
			Configuration: func() interface{} {
				return gin.H{
					"enable": false,
				}
			},
			Expected: Configuration{
				Enable: false,
			},
		}, {
			Description: "enable = true",
			Initial:     func() interface{} { return Configuration{} },
			Configuration: func() interface{} {
				return gin.H{
					"enable": true,
				}
			},
			Expected: Configuration{
				Enable: true,
			},
		},
	})
}

func TestDefaultConfiguration(t *testing.T) {
	if err := helpers.Validate.Struct(DefaultConfiguration()); err != nil {
		t.Fatalf("validate.Struct() error:\n%+v", err)
	}
}
