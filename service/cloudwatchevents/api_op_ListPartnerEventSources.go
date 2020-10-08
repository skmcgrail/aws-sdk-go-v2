// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchevents

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchevents/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// An SaaS partner can use this operation to list all the partner event source
// names that they have created. This operation is not used by AWS customers.
func (c *Client) ListPartnerEventSources(ctx context.Context, params *ListPartnerEventSourcesInput, optFns ...func(*Options)) (*ListPartnerEventSourcesOutput, error) {
	stack := middleware.NewStack("ListPartnerEventSources", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListPartnerEventSourcesMiddlewares(stack)
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
	addOpListPartnerEventSourcesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListPartnerEventSources(options.Region), middleware.Before)
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
			OperationName: "ListPartnerEventSources",
			Err:           err,
		}
	}
	out := result.(*ListPartnerEventSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListPartnerEventSourcesInput struct {

	// If you specify this, the results are limited to only those partner event sources
	// that start with the string you specify.
	//
	// This member is required.
	NamePrefix *string

	// pecifying this limits the number of results returned by this operation. The
	// operation also returns a NextToken which you can use in a subsequent operation
	// to retrieve the next set of results.
	Limit *int32

	// The token returned by a previous call to this operation. Specifying this
	// retrieves the next set of results.
	NextToken *string
}

type ListPartnerEventSourcesOutput struct {

	// A token you can use in a subsequent operation to retrieve the next set of
	// results.
	NextToken *string

	// The list of partner event sources returned by the operation.
	PartnerEventSources []*types.PartnerEventSource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListPartnerEventSourcesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListPartnerEventSources{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPartnerEventSources{}, middleware.After)
}

func newServiceMetadataMiddleware_opListPartnerEventSources(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "ListPartnerEventSources",
	}
}
