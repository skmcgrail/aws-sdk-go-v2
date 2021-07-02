// Code generated by smithy-go-codegen DO NOT EDIT.

package endpoints

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/endpoints"
	"regexp"
	"strings"
)

// Options is the endpoint resolver configuration options
type Options struct {
	DisableHTTPS         bool
	UseDualStackEndpoint aws.DualStackEndpointState
}

func isDualStackEndpointEnabled(value aws.DualStackEndpointState) bool {
	return value == aws.DualStackEndpointStateEnabled
}

// Resolver S3 endpoint resolver
type Resolver struct {
	partitions endpoints.Partitions
}

// ResolveEndpoint resolves the service endpoint for the given region and options
func (r *Resolver) ResolveEndpoint(region string, options Options) (endpoint aws.Endpoint, err error) {
	if len(region) == 0 {
		return endpoint, &aws.MissingRegionError{}
	}

	opt := endpoints.Options{
		DisableHTTPS: options.DisableHTTPS,
		UseDualStack: isDualStackEndpointEnabled(options.UseDualStackEndpoint),
	}
	return r.partitions.ResolveEndpoint(region, opt)
}

// New returns a new Resolver
func New() *Resolver {
	return &Resolver{
		partitions: defaultPartitions,
	}
}

var partitionRegexp = struct {
	Aws      *regexp.Regexp
	AwsCn    *regexp.Regexp
	AwsIso   *regexp.Regexp
	AwsIsoB  *regexp.Regexp
	AwsUsGov *regexp.Regexp
}{

	Aws:      regexp.MustCompile("^(us|eu|ap|sa|ca|me|af)\\-\\w+\\-\\d+$"),
	AwsCn:    regexp.MustCompile("^cn\\-\\w+\\-\\d+$"),
	AwsIso:   regexp.MustCompile("^us\\-iso\\-\\w+\\-\\d+$"),
	AwsIsoB:  regexp.MustCompile("^us\\-isob\\-\\w+\\-\\d+$"),
	AwsUsGov: regexp.MustCompile("^us\\-gov\\-\\w+\\-\\d+$"),
}

var defaultPartitions = endpoints.Partitions{
	{
		ID: "aws",
		Defaults: endpoints.Endpoint{
			Hostname:          "s3.{region}.amazonaws.com",
			Protocols:         []string{"http", "https"},
			SignatureVersions: []string{"s3v4"},
		},
		DualStackDefaults: endpoints.Endpoint{
			Hostname:          "s3.{region}.aws",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.Aws,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"accesspoint-af-south-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.af-south-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.ap-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-northeast-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.ap-northeast-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-northeast-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.ap-northeast-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-northeast-3": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.ap-northeast-3.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-south-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.ap-south-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-southeast-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.ap-southeast-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-southeast-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.ap-southeast-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ca-central-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.ca-central-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-central-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.eu-central-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-north-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.eu-north-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-south-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.eu-south-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-west-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.eu-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-west-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.eu-west-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-west-3": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.eu-west-3.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-me-south-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.me-south-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-sa-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.sa-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-us-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-us-east-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.us-east-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-us-west-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-us-west-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"af-south-1": endpoints.Endpoint{},
			"ap-east-1":  endpoints.Endpoint{},
			"ap-northeast-1": endpoints.Endpoint{
				Hostname:          "s3.ap-northeast-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-northeast-2": endpoints.Endpoint{},
			"ap-northeast-3": endpoints.Endpoint{},
			"ap-south-1":     endpoints.Endpoint{},
			"ap-southeast-1": endpoints.Endpoint{
				Hostname:          "s3.ap-southeast-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-southeast-2": endpoints.Endpoint{
				Hostname:          "s3.ap-southeast-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"aws-global": endpoints.Endpoint{
				Hostname:          "s3.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			"ca-central-1": endpoints.Endpoint{},
			"eu-central-1": endpoints.Endpoint{},
			"eu-north-1":   endpoints.Endpoint{},
			"eu-south-1":   endpoints.Endpoint{},
			"eu-west-1": endpoints.Endpoint{
				Hostname:          "s3.eu-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"eu-west-2": endpoints.Endpoint{},
			"eu-west-3": endpoints.Endpoint{},
			"fips-accesspoint-ca-central-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.ca-central-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-east-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.us-east-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-west-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-west-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"me-south-1": endpoints.Endpoint{},
			"s3-external-1": endpoints.Endpoint{
				Hostname:          "s3-external-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
				CredentialScope: endpoints.CredentialScope{
					Region: "us-east-1",
				},
			},
			"sa-east-1": endpoints.Endpoint{
				Hostname:          "s3.sa-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"us-east-1": endpoints.Endpoint{
				Hostname:          "s3.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"us-east-2": endpoints.Endpoint{},
			"us-west-1": endpoints.Endpoint{
				Hostname:          "s3.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"us-west-2": endpoints.Endpoint{
				Hostname:          "s3.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
		},
		DualStackEndpoints: endpoints.Endpoints{
			"accesspoint-af-south-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.af-south-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.ap-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-northeast-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.ap-northeast-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-northeast-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.ap-northeast-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-northeast-3": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.ap-northeast-3.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-south-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.ap-south-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-south-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.ap-south-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-southeast-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.ap-southeast-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-southeast-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.ap-southeast-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-southeast-3": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.ap-southeast-3.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ap-southeast-4": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.ap-southeast-4.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-ca-central-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.ca-central-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-central-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.eu-central-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-central-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.eu-central-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-north-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.eu-north-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-south-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.eu-south-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-south-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.eu-south-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-west-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.eu-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-west-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.eu-west-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-eu-west-3": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.eu-west-3.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-me-central-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.me-central-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-me-south-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.me-south-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-me-west-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.me-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-sa-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.sa-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-us-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-us-east-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.us-east-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-us-west-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-us-west-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"af-south-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.af-south-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-east-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.ap-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-northeast-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.ap-northeast-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-northeast-2": endpoints.Endpoint{
				Hostname:          "s3.dualstack.ap-northeast-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-northeast-3": endpoints.Endpoint{
				Hostname:          "s3.dualstack.ap-northeast-3.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-south-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.ap-south-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-south-2": endpoints.Endpoint{
				Hostname:          "s3.dualstack.ap-south-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-southeast-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.ap-southeast-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-southeast-2": endpoints.Endpoint{
				Hostname:          "s3.dualstack.ap-southeast-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-southeast-3": endpoints.Endpoint{
				Hostname:          "s3.dualstack.ap-southeast-3.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ap-southeast-4": endpoints.Endpoint{
				Hostname:          "s3.dualstack.ap-southeast-4.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"ca-central-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.ca-central-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"eu-central-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.eu-central-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"eu-central-2": endpoints.Endpoint{
				Hostname:          "s3.dualstack.eu-central-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"eu-north-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.eu-north-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"eu-south-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.eu-south-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"eu-south-2": endpoints.Endpoint{
				Hostname:          "s3.dualstack.eu-south-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"eu-west-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.eu-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"eu-west-2": endpoints.Endpoint{
				Hostname:          "s3.dualstack.eu-west-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"eu-west-3": endpoints.Endpoint{
				Hostname:          "s3.dualstack.eu-west-3.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"fips-accesspoint-ca-central-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.dualstack.ca-central-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.dualstack.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-east-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.dualstack.us-east-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-west-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.dualstack.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-west-2": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.dualstack.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"me-central-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.me-central-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"me-south-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.me-south-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"me-west-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.me-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"sa-east-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.sa-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"us-east-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.us-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"us-east-2": endpoints.Endpoint{
				Hostname:          "s3.dualstack.us-east-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"us-west-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.us-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"us-west-2": endpoints.Endpoint{
				Hostname:          "s3.dualstack.us-west-2.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
		},
	},
	{
		ID: "aws-cn",
		Defaults: endpoints.Endpoint{
			Hostname:          "s3.{region}.amazonaws.com.cn",
			Protocols:         []string{"http", "https"},
			SignatureVersions: []string{"s3v4"},
		},
		DualStackDefaults: endpoints.Endpoint{
			Hostname:          "s3.{region}.amazonwebservices.com.cn",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsCn,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"accesspoint-cn-north-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.cn-north-1.amazonaws.com.cn",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-cn-northwest-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.cn-northwest-1.amazonaws.com.cn",
				SignatureVersions: []string{"s3v4"},
			},
			"cn-north-1":     endpoints.Endpoint{},
			"cn-northwest-1": endpoints.Endpoint{},
		},
		DualStackEndpoints: endpoints.Endpoints{
			"accesspoint-cn-north-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.cn-north-1.amazonaws.com.cn",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-cn-northwest-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.cn-northwest-1.amazonaws.com.cn",
				SignatureVersions: []string{"s3v4"},
			},
			"cn-north-1": endpoints.Endpoint{
				Hostname: "s3.dualstack.cn-north-1.amazonaws.com.cn",
			},
			"cn-northwest-1": endpoints.Endpoint{
				Hostname: "s3.dualstack.cn-northwest-1.amazonaws.com.cn",
			},
		},
	},
	{
		ID: "aws-iso",
		Defaults: endpoints.Endpoint{
			Hostname:          "s3.{region}.c2s.ic.gov",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"s3v4"},
		},
		RegionRegex:    partitionRegexp.AwsIso,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"us-iso-east-1": endpoints.Endpoint{
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3v4"},
			},
		},
	},
	{
		ID: "aws-iso-b",
		Defaults: endpoints.Endpoint{
			Hostname:          "s3.{region}.sc2s.sgov.gov",
			Protocols:         []string{"http", "https"},
			SignatureVersions: []string{"s3v4"},
		},
		RegionRegex:    partitionRegexp.AwsIsoB,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"us-isob-east-1": endpoints.Endpoint{},
		},
	},
	{
		ID: "aws-us-gov",
		Defaults: endpoints.Endpoint{
			Hostname:          "s3.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"s3", "s3v4"},
		},
		DualStackDefaults: endpoints.Endpoint{
			Hostname:          "s3.{region}.aws",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		RegionRegex:    partitionRegexp.AwsUsGov,
		IsRegionalized: true,
		Endpoints: endpoints.Endpoints{
			"accesspoint-us-gov-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.us-gov-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-us-gov-west-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.us-gov-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-gov-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.us-gov-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-gov-west-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.us-gov-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-us-gov-west-1": endpoints.Endpoint{
				Hostname: "s3-fips.us-gov-west-1.amazonaws.com",
				CredentialScope: endpoints.CredentialScope{
					Region: "us-gov-west-1",
				},
			},
			"us-gov-east-1": endpoints.Endpoint{
				Hostname:  "s3.us-gov-east-1.amazonaws.com",
				Protocols: []string{"http", "https"},
			},
			"us-gov-west-1": endpoints.Endpoint{
				Hostname:  "s3.us-gov-west-1.amazonaws.com",
				Protocols: []string{"http", "https"},
			},
		},
		DualStackEndpoints: endpoints.Endpoints{
			"accesspoint-us-gov-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.us-gov-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"accesspoint-us-gov-west-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint.dualstack.us-gov-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-gov-east-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.dualstack.us-gov-east-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"fips-accesspoint-us-gov-west-1": endpoints.Endpoint{
				Hostname:          "s3-accesspoint-fips.dualstack.us-gov-west-1.amazonaws.com",
				SignatureVersions: []string{"s3v4"},
			},
			"us-gov-east-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.us-gov-east-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
			"us-gov-west-1": endpoints.Endpoint{
				Hostname:          "s3.dualstack.us-gov-west-1.amazonaws.com",
				SignatureVersions: []string{"s3", "s3v4"},
			},
		},
	},
}

// GetDNSSuffix returns the dnsSuffix URL component for the given partition id
func GetDNSSuffix(id string, options Options) (string, error) {
	switch {
	case strings.EqualFold(id, "aws"):
		if options.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled {
			return "aws", nil
		}
		return "amazonaws.com", nil

	case strings.EqualFold(id, "aws-cn"):
		if options.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled {
			return "amazonwebservices.com.cn", nil
		}
		return "amazonaws.com.cn", nil

	case strings.EqualFold(id, "aws-iso"):
		if options.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled {
			return "", fmt.Errorf("partition aws-iso does not have a dual-stack dns suffix")
		}
		return "c2s.ic.gov", nil

	case strings.EqualFold(id, "aws-iso-b"):
		if options.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled {
			return "", fmt.Errorf("partition aws-iso-b does not have a dual-stack dns suffix")
		}
		return "sc2s.sgov.gov", nil

	case strings.EqualFold(id, "aws-us-gov"):
		if options.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled {
			return "aws", nil
		}
		return "amazonaws.com", nil

	default:
		return "", fmt.Errorf("unknown partition")

	}
}

// GetDNSSuffixFromRegion returns the dnsSuffix URL component for the given
// partition id
func GetDNSSuffixFromRegion(region string, options Options) (string, error) {
	switch {
	case partitionRegexp.Aws.MatchString(region):
		if options.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled {
			return "aws", nil
		}
		return "amazonaws.com", nil

	case partitionRegexp.AwsCn.MatchString(region):
		if options.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled {
			return "amazonwebservices.com.cn", nil
		}
		return "amazonaws.com.cn", nil

	case partitionRegexp.AwsIso.MatchString(region):
		if options.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled {
			return "", fmt.Errorf("partition aws-iso does not have a dual-stack dns suffix")
		}
		return "c2s.ic.gov", nil

	case partitionRegexp.AwsIsoB.MatchString(region):
		if options.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled {
			return "", fmt.Errorf("partition aws-iso-b does not have a dual-stack dns suffix")
		}
		return "sc2s.sgov.gov", nil

	case partitionRegexp.AwsUsGov.MatchString(region):
		if options.UseDualStackEndpoint == aws.DualStackEndpointStateEnabled {
			return "aws", nil
		}
		return "amazonaws.com", nil

	default:
		return "", fmt.Errorf("unknown region partition")

	}
}
