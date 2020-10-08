// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes one or more of your compute environments. If you are using an
// unmanaged compute environment, you can use the DescribeComputeEnvironment
// operation to determine the ecsClusterArn that you should launch your Amazon ECS
// container instances into.
func (c *Client) DescribeComputeEnvironments(ctx context.Context, params *DescribeComputeEnvironmentsInput, optFns ...func(*Options)) (*DescribeComputeEnvironmentsOutput, error) {
	stack := middleware.NewStack("DescribeComputeEnvironments", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeComputeEnvironmentsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeComputeEnvironments(options.Region), middleware.Before)
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
			OperationName: "DescribeComputeEnvironments",
			Err:           err,
		}
	}
	out := result.(*DescribeComputeEnvironmentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeComputeEnvironmentsInput struct {

	// A list of up to 100 compute environment names or full Amazon Resource Name (ARN)
	// entries.
	ComputeEnvironments []*string

	// The maximum number of cluster results returned by DescribeComputeEnvironments in
	// paginated output. When this parameter is used, DescribeComputeEnvironments only
	// returns maxResults results in a single page along with a nextToken response
	// element. The remaining results of the initial request can be seen by sending
	// another DescribeComputeEnvironments request with the returned nextToken value.
	// This value can be between 1 and 100. If this parameter is not used, then
	// DescribeComputeEnvironments returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// The nextToken value returned from a previous paginated
	// DescribeComputeEnvironments request where maxResults was used and the results
	// exceeded the value of that parameter. Pagination continues from the end of the
	// previous results that returned the nextToken value. This value is null when
	// there are no more results to return. This token should be treated as an opaque
	// identifier that is only used to retrieve the next items in a list and not for
	// other programmatic purposes.
	NextToken *string
}

type DescribeComputeEnvironmentsOutput struct {

	// The list of compute environments.
	ComputeEnvironments []*types.ComputeEnvironmentDetail

	// The nextToken value to include in a future DescribeComputeEnvironments request.
	// When the results of a DescribeJobDefinitions request exceed maxResults, this
	// value can be used to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeComputeEnvironmentsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeComputeEnvironments{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeComputeEnvironments{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeComputeEnvironments(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "DescribeComputeEnvironments",
	}
}
