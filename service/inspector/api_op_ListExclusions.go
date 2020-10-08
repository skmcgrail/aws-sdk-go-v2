// Code generated by smithy-go-codegen DO NOT EDIT.

package inspector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List exclusions that are generated by the assessment run.
func (c *Client) ListExclusions(ctx context.Context, params *ListExclusionsInput, optFns ...func(*Options)) (*ListExclusionsOutput, error) {
	stack := middleware.NewStack("ListExclusions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListExclusionsMiddlewares(stack)
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
	addOpListExclusionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListExclusions(options.Region), middleware.Before)
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
			OperationName: "ListExclusions",
			Err:           err,
		}
	}
	out := result.(*ListExclusionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListExclusionsInput struct {

	// The ARN of the assessment run that generated the exclusions that you want to
	// list.
	//
	// This member is required.
	AssessmentRunArn *string

	// You can use this parameter to indicate the maximum number of items you want in
	// the response. The default value is 100. The maximum value is 500.
	MaxResults *int32

	// You can use this parameter when paginating results. Set the value of this
	// parameter to null on your first call to the ListExclusionsRequest action.
	// Subsequent calls to the action fill nextToken in the request with the value of
	// nextToken from the previous response to continue listing data.
	NextToken *string
}

type ListExclusionsOutput struct {

	// A list of exclusions' ARNs returned by the action.
	//
	// This member is required.
	ExclusionArns []*string

	// When a response is generated, if there is more data to be listed, this
	// parameters is present in the response and contains the value to use for the
	// nextToken parameter in a subsequent pagination request. If there is no more data
	// to be listed, this parameter is set to null.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListExclusionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListExclusions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListExclusions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListExclusions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "inspector",
		OperationName: "ListExclusions",
	}
}
