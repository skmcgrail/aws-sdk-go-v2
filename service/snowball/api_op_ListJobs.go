// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of JobListEntry objects of the specified length. Each
// JobListEntry object contains a job's state, a job's ID, and a value that
// indicates whether the job is a job part, in the case of export jobs. Calling
// this API action in one of the US regions will return jobs from the list of all
// jobs associated with this account in all US regions.
func (c *Client) ListJobs(ctx context.Context, params *ListJobsInput, optFns ...func(*Options)) (*ListJobsOutput, error) {
	stack := middleware.NewStack("ListJobs", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListJobsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListJobs(options.Region), middleware.Before)
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
			OperationName: "ListJobs",
			Err:           err,
		}
	}
	out := result.(*ListJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJobsInput struct {
	// HTTP requests are stateless. To identify what object comes "next" in the list of
	// JobListEntry objects, you have the option of specifying NextToken as the
	// starting point for your returned list.
	NextToken *string
	// The number of JobListEntry objects to return.
	MaxResults *int32
}

type ListJobsOutput struct {
	// Each JobListEntry object contains a job's state, a job's ID, and a value that
	// indicates whether the job is a job part, in the case of export jobs.
	JobListEntries []*types.JobListEntry
	// HTTP requests are stateless. If you use this automatically generated NextToken
	// value in your next ListJobs call, your returned JobListEntry objects will start
	// from this point in the array.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListJobsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListJobs{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListJobs{}, middleware.After)
}

func newServiceMetadataMiddleware_opListJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "ListJobs",
	}
}