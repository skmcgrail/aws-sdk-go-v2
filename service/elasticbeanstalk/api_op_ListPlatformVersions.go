// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticbeanstalk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the platform versions available for your account in an AWS Region.
// Provides summary information about each platform version. Compare to
// DescribePlatformVersion (), which provides full details about a single platform
// version. For definitions of platform version and other platform-related terms,
// see AWS Elastic Beanstalk Platforms Glossary
// (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/platforms-glossary.html).
func (c *Client) ListPlatformVersions(ctx context.Context, params *ListPlatformVersionsInput, optFns ...func(*Options)) (*ListPlatformVersionsOutput, error) {
	stack := middleware.NewStack("ListPlatformVersions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListPlatformVersionsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPlatformVersions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListPlatformVersions",
			Err:           err,
		}
	}
	out := result.(*ListPlatformVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPlatformVersionsInput struct {

	// Criteria for restricting the resulting list of platform versions. The filter is
	// interpreted as a logical conjunction (AND) of the separate PlatformFilter terms.
	Filters []*types.PlatformFilter

	// The maximum number of platform version values returned in one call.
	MaxRecords *int32

	// For a paginated request. Specify a token from a previous response page to
	// retrieve the next response page. All other parameter values must be identical to
	// the ones specified in the initial request. If no NextToken is specified, the
	// first page is retrieved.
	NextToken *string
}

type ListPlatformVersionsOutput struct {

	// In a paginated request, if this value isn't null, it's the token that you can
	// pass in a subsequent request to get the next response page.
	NextToken *string

	// Summary information about the platform versions.
	PlatformSummaryList []*types.PlatformSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListPlatformVersionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListPlatformVersions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListPlatformVersions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPlatformVersions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticbeanstalk",
		OperationName: "ListPlatformVersions",
	}
}
