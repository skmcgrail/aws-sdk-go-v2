module github.com/aws/aws-sdk-go-v2/service/mwaa

go 1.15

require (
	github.com/aws/aws-sdk-go-v2 v1.7.0
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.0.0
	github.com/aws/smithy-go v1.5.0
)

replace github.com/aws/aws-sdk-go-v2 => ../../

replace github.com/aws/aws-sdk-go-v2/internal/configsources => ../../internal/configsources/
