package endpoints

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/google/go-cmp/cmp"
)

func TestEndpointResolve(t *testing.T) {
	defs := Endpoint{
		Hostname:          "service.{region}.amazonaws.com",
		SignatureVersions: []string{"v4"},
	}

	e := Endpoint{
		Protocols:         []string{"http", "https"},
		SignatureVersions: []string{"v4"},
		CredentialScope: CredentialScope{
			Region:  "us-west-2",
			Service: "service",
		},
	}

	resolved, err := e.resolve("aws", "us-west-2", defs, Options{})
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}

	if e, a := "https://service.us-west-2.amazonaws.com", resolved.URL; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "aws", resolved.PartitionID; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "service", resolved.SigningName; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "us-west-2", resolved.SigningRegion; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "v4", resolved.SigningMethod; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestEndpointMergeIn(t *testing.T) {
	expected := Endpoint{
		Hostname:          "other hostname",
		Protocols:         []string{"http"},
		SignatureVersions: []string{"v4"},
		CredentialScope: CredentialScope{
			Region:  "region",
			Service: "service",
		},
	}

	actual := Endpoint{}
	actual.mergeIn(Endpoint{
		Hostname:          "other hostname",
		Protocols:         []string{"http"},
		SignatureVersions: []string{"v4"},
		CredentialScope: CredentialScope{
			Region:  "region",
			Service: "service",
		},
	})

	if e, a := expected, actual; !reflect.DeepEqual(e, a) {
		t.Errorf("expect %v, got %v", e, a)
	}
}

var testPartitions = Partitions{
	{
		ID: "part-id-1",
		RegionRegex: func() *regexp.Regexp {
			reg, _ := regexp.Compile("^(us)\\-\\w+\\-\\d+$")
			return reg
		}(),
		Defaults: Endpoint{
			Hostname:          "service.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
		},
		DualStackDefaults: Endpoint{
			Hostname:          "api.service.{region}.aws",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
			CredentialScope: CredentialScope{
				Service: "service",
			},
		},
		IsRegionalized: true,
		Endpoints: Endpoints{
			"us-west-1": {},
			"us-west-1-alt": {
				Hostname:          "service-alt.us-west-1.amazonaws.com",
				Protocols:         []string{"http"},
				SignatureVersions: []string{"vFoo"},
				CredentialScope: CredentialScope{
					Region:  "us-west-1",
					Service: "foo",
				},
			},
		},
		DualStackEndpoints: Endpoints{
			"us-west-2": Endpoint{},
		},
	},
	{
		ID: "part-id-2",
		RegionRegex: func() *regexp.Regexp {
			reg, _ := regexp.Compile("^(cn)\\-\\w+\\-\\d+$")
			return reg
		}(),
		Defaults: Endpoint{
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
			CredentialScope: CredentialScope{
				Service: "foo",
			},
		},
		IsRegionalized:    false,
		PartitionEndpoint: "partition",
		Endpoints: Endpoints{
			"partition": {
				Hostname: "some-global-thing.amazonaws.com.cn",
				CredentialScope: CredentialScope{
					Region: "cn-east-1",
				},
			},
			"fips-partition": {
				Hostname: "some-global-thing-fips.amazonaws.com.cn",
				CredentialScope: CredentialScope{
					Region: "cn-east-1",
				},
			},
		},
		DualStackEndpoints: Endpoints{
			"partition": {
				Hostname: "some-global-thing.global.aws",
				CredentialScope: CredentialScope{
					Region:  "cn-east-1",
					Service: "foo",
				},
			},
		},
	},
	{
		ID: "part-id-3",
		RegionRegex: func() *regexp.Regexp {
			reg, _ := regexp.Compile("^(eu)\\-\\w+\\-\\d+$")
			return reg
		}(),
		Defaults: Endpoint{
			Hostname:          "service.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
			CredentialScope: CredentialScope{
				Service: "foo",
			},
		},
		IsRegionalized: true,
	},
	{
		ID: "part-id-4",
		RegionRegex: func() *regexp.Regexp {
			reg, _ := regexp.Compile("^(ca)\\-\\w+\\-\\d+$")
			return reg
		}(),
		Defaults: Endpoint{
			Hostname:          "service.{region}.amazonaws.com",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
			CredentialScope: CredentialScope{
				Service: "foo",
			},
		},
		DualStackDefaults: Endpoint{
			Hostname:          "service.{region}.aws",
			Protocols:         []string{"https"},
			SignatureVersions: []string{"v4"},
			CredentialScope: CredentialScope{
				Service: "foo",
			},
		},
		IsRegionalized: true,
	},
}

func TestResolveEndpoint(t *testing.T) {
	var cases = map[string]struct {
		Region   string
		Options  Options
		Expected aws.Endpoint
		WantErr  bool
	}{
		"modeled region with no endpoint overrides": {
			Region: "us-west-1",
			Expected: aws.Endpoint{
				PartitionID:   "part-id-1",
				URL:           "https://service.us-west-1.amazonaws.com",
				SigningRegion: "us-west-1",
				SigningMethod: "v4",
			},
		},
		"modeled region with no endpoint overrides and https disabled": {
			Region:  "us-west-1",
			Options: Options{DisableHTTPS: true},
			Expected: aws.Endpoint{
				PartitionID:   "part-id-1",
				URL:           "http://service.us-west-1.amazonaws.com",
				SigningRegion: "us-west-1",
				SigningMethod: "v4",
			},
		},
		"modeled region with endpoint overrides": {
			Region: "us-west-1-alt",
			Expected: aws.Endpoint{
				PartitionID:   "part-id-1",
				URL:           "http://service-alt.us-west-1.amazonaws.com",
				SigningRegion: "us-west-1",
				SigningName:   "foo",
				SigningMethod: "vFoo",
			},
		},
		"partition endpoint": {
			Region: "cn-central-1",
			Expected: aws.Endpoint{
				PartitionID:   "part-id-2",
				URL:           "https://some-global-thing.amazonaws.com.cn",
				SigningRegion: "cn-east-1",
				SigningName:   "foo",
				SigningMethod: "v4",
			},
		},
		"specified partition endpoint": {
			Region: "partition",
			Expected: aws.Endpoint{
				PartitionID:   "part-id-2",
				URL:           "https://some-global-thing.amazonaws.com.cn",
				SigningRegion: "cn-east-1",
				SigningName:   "foo",
				SigningMethod: "v4",
			},
		},
		"fips partition endpoint": {
			Region: "fips-partition",
			Expected: aws.Endpoint{
				PartitionID:   "part-id-2",
				URL:           "https://some-global-thing-fips.amazonaws.com.cn",
				SigningRegion: "cn-east-1",
				SigningName:   "foo",
				SigningMethod: "v4",
			},
		},
		"region with unmodeled endpoints": {
			Region: "eu-west-1",
			Expected: aws.Endpoint{
				PartitionID:   "part-id-3",
				URL:           "https://service.eu-west-1.amazonaws.com",
				SigningRegion: "eu-west-1",
				SigningName:   "foo",
				SigningMethod: "v4",
			},
		},
		"us-west-2 dual-stack enabled, modeled region": {
			Region:  "us-west-2",
			Options: Options{UseDualStack: true},
			Expected: aws.Endpoint{
				PartitionID:   "part-id-1",
				URL:           "https://api.service.us-west-2.aws",
				SigningName:   "service",
				SigningRegion: "us-west-2",
				SigningMethod: "v4",
			},
		},
		"us-west-3 dual-stack enabled, not modeled fallback to defaults": {
			Region:  "us-west-3",
			Options: Options{UseDualStack: true},
			Expected: aws.Endpoint{
				PartitionID:   "part-id-1",
				URL:           "https://api.service.us-west-3.aws",
				SigningName:   "service",
				SigningRegion: "us-west-3",
				SigningMethod: "v4",
			},
		},
		"ca-west-1, dual-stack enabled, not modeled, and partition with no defaults": {
			Region:  "eu-west-1",
			Options: Options{UseDualStack: true},
			WantErr: true,
		},
		"ca-west-1, dual-stack enabled, no modeled endpoints, partition with defaults": {
			Region:  "ca-west-1",
			Options: Options{UseDualStack: true},
			Expected: aws.Endpoint{
				URL:           "https://service.ca-west-1.aws",
				PartitionID:   "part-id-4",
				SigningName:   "foo",
				SigningRegion: "ca-west-1",
				SigningMethod: "v4",
			},
		},
		"partition endpoint, dual-stack enabled": {
			Region:  "partition",
			Options: Options{UseDualStack: true},
			Expected: aws.Endpoint{
				PartitionID:   "part-id-2",
				URL:           "https://some-global-thing.global.aws",
				SigningRegion: "cn-east-1",
				SigningName:   "foo",
				SigningMethod: "v4",
			},
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			endpoint, err := testPartitions.ResolveEndpoint(tt.Region, tt.Options)
			if (err != nil) != (tt.WantErr) {
				t.Errorf("WantErr=%v, got error: %v", tt.WantErr, err)
			}
			if diff := cmp.Diff(tt.Expected, endpoint); len(diff) > 0 {
				t.Error(diff)
			}
		})
	}
}
