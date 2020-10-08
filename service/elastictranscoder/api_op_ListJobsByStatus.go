// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The ListJobsByStatus operation gets a list of jobs that have a specified status.
// The response body contains one element for each job that satisfies the search
// criteria.
func (c *Client) ListJobsByStatus(ctx context.Context, params *ListJobsByStatusInput, optFns ...func(*Options)) (*ListJobsByStatusOutput, error) {
	stack := middleware.NewStack("ListJobsByStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListJobsByStatusMiddlewares(stack)
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
	addOpListJobsByStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListJobsByStatus(options.Region), middleware.Before)
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
			OperationName: "ListJobsByStatus",
			Err:           err,
		}
	}
	out := result.(*ListJobsByStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListJobsByStatusRequest structure.
type ListJobsByStatusInput struct {

	// To get information about all of the jobs associated with the current AWS account
	// that have a given status, specify the following status: Submitted, Progressing,
	// Complete, Canceled, or Error.
	//
	// This member is required.
	Status *string

	// To list jobs in chronological order by the date and time that they were
	// submitted, enter true. To list jobs in reverse chronological order, enter false.
	Ascending *string

	// When Elastic Transcoder returns more than one page of results, use pageToken in
	// subsequent GET requests to get each successive page of results.
	PageToken *string
}

// The ListJobsByStatusResponse structure.
type ListJobsByStatusOutput struct {

	// An array of Job objects that have the specified status.
	Jobs []*types.Job

	// A value that you use to access the second and subsequent pages of results, if
	// any. When the jobs in the specified pipeline fit on one page or when you've
	// reached the last page of results, the value of NextPageToken is null.
	NextPageToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListJobsByStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListJobsByStatus{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListJobsByStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opListJobsByStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "ListJobsByStatus",
	}
}
