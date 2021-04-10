package gomod

import (
	"github.com/aws/aws-sdk-go-v2/internal/repotools/git"
	"reflect"
	"testing"
)

func TestParseTags(t *testing.T) {
	tests := map[string]struct {
		args []string
		want map[string][]string
	}{
		"standard tags": {
			args: []string{
				"v0.2.1",
				"v1.0.0",
				"v1.3.0",
				"feature/ec2/imds/v0.1.0",
				"feature/ec2/imds/v1.0.1",
				"feature/ec2/imds/v1.0.6",
			},
			want: map[string][]string{
				"/":                {"v1.3.0", "v1.0.0", "v0.2.1"},
				"feature/ec2/imds": {"v1.0.6", "v1.0.1", "v0.1.0"},
			},
		},
		"invalid tags": {
			args: []string{
				"v0.2.1",
				"v1.0.0",
				"release-1-2021-04-09",
				"v1.3.0",
				"1.4.0",
				"feature/ec2/imds/v0.1.0",
				"feature/ec2/imds/v1.0.1",
				"feature/ec2/imds/v1.0.6",
				"feature/ec2/imds@v1.2.0",
			},
			want: map[string][]string{
				"/":                {"v1.3.0", "v1.0.0", "v0.2.1"},
				"feature/ec2/imds": {"v1.0.6", "v1.0.1", "v0.1.0"},
			},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := git.ParseTags(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTags() = %v, want %v", got, tt.want)
			}
		})
	}
}
