package endpoints

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const (
	defaultProtocol = "https"
	defaultSigner   = "v4"
)

var (
	protocolPriority = []string{"https", "http"}
	signerPriority   = []string{"v4"}
)

// Options provide configuration needed to direct how endpoints are resolved.
type Options struct {
	// Disable usage of HTTPS (TLS / SSL)
	DisableHTTPS bool

	// Instruct the resolver to use a service endpoint that supports dual-stack.
	// If a service does not have a dual-stack endpoint an error will be returned by the resolver.
	UseDualStack bool
}

// Partitions is a slice of partition
type Partitions []Partition

// ResolveEndpoint resolves a service endpoint for the given region and options.
func (ps Partitions) ResolveEndpoint(region string, opts Options) (aws.Endpoint, error) {
	if len(ps) == 0 {
		return aws.Endpoint{}, fmt.Errorf("no partitions found")
	}

	for i := 0; i < len(ps); i++ {
		if !ps[i].canResolveEndpoint(region) {
			continue
		}

		return ps[i].ResolveEndpoint(region, opts)
	}

	// fallback to first partition format to use when resolving the endpoint.
	return ps[0].ResolveEndpoint(region, opts)
}

// Partition is an AWS partition description for a service and its' region endpoints.
type Partition struct {
	ID                 string
	RegionRegex        *regexp.Regexp
	PartitionEndpoint  string
	IsRegionalized     bool
	Defaults           Endpoint
	DualStackDefaults  Endpoint
	Endpoints          Endpoints
	DualStackEndpoints Endpoints
}

func (p Partition) canResolveEndpoint(region string) bool {
	_, ok := p.Endpoints[region]
	return ok || p.RegionRegex.MatchString(region)
}

// ResolveEndpoint resolves and service endpoint for the given region and options.
func (p Partition) ResolveEndpoint(region string, options Options) (resolved aws.Endpoint, err error) {
	if len(region) == 0 && len(p.PartitionEndpoint) != 0 {
		region = p.PartitionEndpoint
	}

	endpoints := p.Endpoints
	defaults := p.Defaults
	if options.UseDualStack {
		endpoints = p.DualStackEndpoints
		defaults = p.DualStackDefaults
	}

	return p.endpointForRegion(region, endpoints).resolve(p.ID, region, defaults, options)
}

func (p Partition) endpointForRegion(region string, endpoints Endpoints) Endpoint {
	if e, ok := endpoints[region]; ok {
		return e
	}

	if !p.IsRegionalized {
		return endpoints[p.PartitionEndpoint]
	}

	// Unable to find any matching endpoint, return
	// blank that will be used for generic endpoint creation.
	return Endpoint{}
}

// Endpoints is a map of service config regions to endpoints
type Endpoints map[string]Endpoint

// CredentialScope is the credential scope of a region and service
type CredentialScope struct {
	Region  string
	Service string
}

// Endpoint is a service endpoint description
type Endpoint struct {
	// True if the endpoint cannot be resolved for this partition/region/service
	Unresolveable aws.Ternary

	Hostname  string
	Protocols []string

	CredentialScope CredentialScope

	SignatureVersions []string `json:"signatureVersions"`
}

// IsZero returns whether the endpoint structure is an empty (zero) value.
func (e Endpoint) IsZero() bool {
	switch {
	case e.Unresolveable != aws.UnknownTernary:
		return false
	case len(e.Hostname) != 0:
		return false
	case len(e.Protocols) != 0:
		return false
	case e.CredentialScope != (CredentialScope{}):
		return false
	case len(e.SignatureVersions) != 0:
		return false
	}
	return true
}

func (e Endpoint) resolve(partition, region string, def Endpoint, options Options) (aws.Endpoint, error) {
	var merged Endpoint
	merged.mergeIn(def)
	merged.mergeIn(e)
	e = merged

	if e.IsZero() {
		return aws.Endpoint{}, fmt.Errorf("unable to resolve endpoint for region: %v", region)
	}

	var u string
	if e.Unresolveable != aws.TrueTernary {
		// Only attempt to resolve the endpoint if it can be resolved.
		hostname := strings.Replace(e.Hostname, "{region}", region, 1)

		scheme := getEndpointScheme(e.Protocols, options.DisableHTTPS)
		u = scheme + "://" + hostname
	}

	signingRegion := e.CredentialScope.Region
	if len(signingRegion) == 0 {
		signingRegion = region
	}
	signingName := e.CredentialScope.Service

	return aws.Endpoint{
		URL:           u,
		PartitionID:   partition,
		SigningRegion: signingRegion,
		SigningName:   signingName,
		SigningMethod: getByPriority(e.SignatureVersions, signerPriority, defaultSigner),
	}, nil
}

func (e *Endpoint) mergeIn(other Endpoint) {
	if other.Unresolveable != aws.UnknownTernary {
		e.Unresolveable = other.Unresolveable
	}
	if len(other.Hostname) > 0 {
		e.Hostname = other.Hostname
	}
	if len(other.Protocols) > 0 {
		e.Protocols = other.Protocols
	}
	if len(other.CredentialScope.Region) > 0 {
		e.CredentialScope.Region = other.CredentialScope.Region
	}
	if len(other.CredentialScope.Service) > 0 {
		e.CredentialScope.Service = other.CredentialScope.Service
	}
	if len(other.SignatureVersions) > 0 {
		e.SignatureVersions = other.SignatureVersions
	}
}

func getEndpointScheme(protocols []string, disableHTTPS bool) string {
	if disableHTTPS {
		return "http"
	}

	return getByPriority(protocols, protocolPriority, defaultProtocol)
}

func getByPriority(s []string, p []string, def string) string {
	if len(s) == 0 {
		return def
	}

	for i := 0; i < len(p); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] == p[i] {
				return s[j]
			}
		}
	}

	return s[0]
}
