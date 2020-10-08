// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an endpoint configuration. The DeleteEndpointConfig API deletes only the
// specified configuration. It does not delete endpoints created using the
// configuration. You must not delete an EndpointConfig in use by an endpoint that
// is live or while the UpdateEndpoint or CreateEndpoint operations are being
// performed on the endpoint. If you delete the EndpointConfig of an endpoint that
// is active or being created or updated you may lose visibility into the instance
// type the endpoint is using. The endpoint must be deleted in order to stop
// incurring charges.
func (c *Client) DeleteEndpointConfig(ctx context.Context, params *DeleteEndpointConfigInput, optFns ...func(*Options)) (*DeleteEndpointConfigOutput, error) {
	stack := middleware.NewStack("DeleteEndpointConfig", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteEndpointConfigMiddlewares(stack)
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
	addOpDeleteEndpointConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEndpointConfig(options.Region), middleware.Before)
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
			OperationName: "DeleteEndpointConfig",
			Err:           err,
		}
	}
	out := result.(*DeleteEndpointConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEndpointConfigInput struct {

	// The name of the endpoint configuration that you want to delete.
	//
	// This member is required.
	EndpointConfigName *string
}

type DeleteEndpointConfigOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteEndpointConfigMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteEndpointConfig{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteEndpointConfig{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteEndpointConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "DeleteEndpointConfig",
	}
}
