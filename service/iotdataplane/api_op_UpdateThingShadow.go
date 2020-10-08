// Code generated by smithy-go-codegen DO NOT EDIT.

package iotdataplane

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the shadow for the specified thing. For more information, see
// UpdateThingShadow
// (http://docs.aws.amazon.com/iot/latest/developerguide/API_UpdateThingShadow.html)
// in the AWS IoT Developer Guide.
func (c *Client) UpdateThingShadow(ctx context.Context, params *UpdateThingShadowInput, optFns ...func(*Options)) (*UpdateThingShadowOutput, error) {
	stack := middleware.NewStack("UpdateThingShadow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateThingShadowMiddlewares(stack)
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
	addOpUpdateThingShadowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateThingShadow(options.Region), middleware.Before)
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
			OperationName: "UpdateThingShadow",
			Err:           err,
		}
	}
	out := result.(*UpdateThingShadowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the UpdateThingShadow operation.
type UpdateThingShadowInput struct {

	// The state information, in JSON format.
	//
	// This member is required.
	Payload []byte

	// The name of the thing.
	//
	// This member is required.
	ThingName *string

	// The name of the shadow.
	ShadowName *string
}

// The output from the UpdateThingShadow operation.
type UpdateThingShadowOutput struct {

	// The state information, in JSON format.
	Payload []byte

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateThingShadowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateThingShadow{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateThingShadow{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateThingShadow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotdata",
		OperationName: "UpdateThingShadow",
	}
}
