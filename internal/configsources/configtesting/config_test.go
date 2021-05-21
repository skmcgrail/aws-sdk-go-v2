package configtesting

import (
	"github.com/aws/aws-sdk-go-v2/config"
	internalConfig "github.com/aws/aws-sdk-go-v2/internal/configsources"
)

// EnableEndpointDiscoveryProvider Assertions
var (
	_ internalConfig.EnableEndpointDiscoveryProvider = &config.EnvConfig{}
	_ internalConfig.EnableEndpointDiscoveryProvider = &config.SharedConfig{}
)

// EnableEndpointDiscoveryProvider Assertions
var (
	_ internalConfig.EnableEndpointDiscoveryProvider = &config.EnvConfig{}
	_ internalConfig.EnableEndpointDiscoveryProvider = &config.SharedConfig{}
)

// UseDualStackEndpointProvider Assertions
var (
	_ internalConfig.UseDualStackEndpointProvider = &config.EnvConfig{}
	_ internalConfig.UseDualStackEndpointProvider = &config.SharedConfig{}
	_ internalConfig.UseDualStackEndpointProvider = &config.LoadOptions{}
)
