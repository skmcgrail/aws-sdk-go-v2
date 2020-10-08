// Code generated by smithy-go-codegen DO NOT EDIT.

package jsonrpc10

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc10/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// This operation uses unions for inputs and outputs.
func (c *Client) JsonUnions(ctx context.Context, params *JsonUnionsInput, optFns ...func(*Options)) (*JsonUnionsOutput, error) {
	stack := middleware.NewStack("JsonUnions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpJsonUnionsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	addRetryMiddlewares(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opJsonUnions(options.Region), middleware.Before)
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
			OperationName: "JsonUnions",
			Err:           err,
		}
	}
	out := result.(*JsonUnionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A shared structure that contains a single union member.
type JsonUnionsInput struct {

	// A union with a representative set of types for members.
	Contents types.MyUnion
}

// A shared structure that contains a single union member.
type JsonUnionsOutput struct {

	// A union with a representative set of types for members.
	Contents types.MyUnion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpJsonUnionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpJsonUnions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpJsonUnions{}, middleware.After)
}

func newServiceMetadataMiddleware_opJsonUnions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "JsonUnions",
	}
}
