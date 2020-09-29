// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Displays details about an import virtual machine or import snapshot tasks that
// are already created.
func (c *Client) DescribeImportImageTasks(ctx context.Context, params *DescribeImportImageTasksInput, optFns ...func(*Options)) (*DescribeImportImageTasksOutput, error) {
	stack := middleware.NewStack("DescribeImportImageTasks", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsEc2query_serdeOpDescribeImportImageTasksMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeImportImageTasks(options.Region), middleware.Before)
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
			OperationName: "DescribeImportImageTasks",
			Err:           err,
		}
	}
	out := result.(*DescribeImportImageTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeImportImageTasksInput struct {
	// A token that indicates the next page of results.
	NextToken *string
	// The IDs of the import image tasks.
	ImportTaskIds []*string
	// The maximum number of results to return in a single call.
	MaxResults *int32
	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
	// Filter tasks using the task-state filter and one of the following values:
	// active, completed, deleting, or deleted.
	Filters []*types.Filter
}

type DescribeImportImageTasksOutput struct {
	// A list of zero or more import image tasks that are currently active or were
	// completed or canceled in the previous 7 days.
	ImportImageTasks []*types.ImportImageTask
	// The token to use to get the next page of results. This value is null when there
	// are no more results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsEc2query_serdeOpDescribeImportImageTasksMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsEc2query_serializeOpDescribeImportImageTasks{}, middleware.After)
	stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeImportImageTasks{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeImportImageTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeImportImageTasks",
	}
}