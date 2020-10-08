// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the topic rule destinations in your AWS account.
func (c *Client) ListTopicRuleDestinations(ctx context.Context, params *ListTopicRuleDestinationsInput, optFns ...func(*Options)) (*ListTopicRuleDestinationsOutput, error) {
	stack := middleware.NewStack("ListTopicRuleDestinations", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListTopicRuleDestinationsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTopicRuleDestinations(options.Region), middleware.Before)
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
			OperationName: "ListTopicRuleDestinations",
			Err:           err,
		}
	}
	out := result.(*ListTopicRuleDestinationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTopicRuleDestinationsInput struct {

	// The maximum number of results to return at one time.
	MaxResults *int32

	// The token to retrieve the next set of results.
	NextToken *string
}

type ListTopicRuleDestinationsOutput struct {

	// Information about a topic rule destination.
	DestinationSummaries []*types.TopicRuleDestinationSummary

	// The token to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListTopicRuleDestinationsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListTopicRuleDestinations{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListTopicRuleDestinations{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTopicRuleDestinations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "ListTopicRuleDestinations",
	}
}
