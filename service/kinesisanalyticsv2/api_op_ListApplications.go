// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisanalyticsv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisanalyticsv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of Amazon Kinesis Data Analytics applications in your account.
// For each application, the response includes the application name, Amazon
// Resource Name (ARN), and status. If you want detailed information about a
// specific application, use DescribeApplication ().
func (c *Client) ListApplications(ctx context.Context, params *ListApplicationsInput, optFns ...func(*Options)) (*ListApplicationsOutput, error) {
	stack := middleware.NewStack("ListApplications", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListApplicationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListApplications(options.Region), middleware.Before)
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
			OperationName: "ListApplications",
			Err:           err,
		}
	}
	out := result.(*ListApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationsInput struct {

	// The maximum number of applications to list.
	Limit *int32

	// If a previous command returned a pagination token, pass it into this value to
	// retrieve the next set of results. For more information about pagination, see
	// Using the AWS Command Line Interface's Pagination Options
	// (https://docs.aws.amazon.com/cli/latest/userguide/pagination.html).
	NextToken *string
}

type ListApplicationsOutput struct {

	// A list of ApplicationSummary objects.
	//
	// This member is required.
	ApplicationSummaries []*types.ApplicationSummary

	// The pagination token for the next set of results, or null if there are no
	// additional results. Pass this token into a subsequent command to retrieve the
	// next set of items For more information about pagination, see Using the AWS
	// Command Line Interface's Pagination Options
	// (https://docs.aws.amazon.com/cli/latest/userguide/pagination.html).
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListApplicationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListApplications{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListApplications{}, middleware.After)
}

func newServiceMetadataMiddleware_opListApplications(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisanalytics",
		OperationName: "ListApplications",
	}
}
