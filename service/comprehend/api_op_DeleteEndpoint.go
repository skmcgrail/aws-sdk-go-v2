// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a model-specific endpoint for a previously-trained custom model. All
// endpoints must be deleted in order for the model to be deleted.
func (c *Client) DeleteEndpoint(ctx context.Context, params *DeleteEndpointInput, optFns ...func(*Options)) (*DeleteEndpointOutput, error) {
	stack := middleware.NewStack("DeleteEndpoint", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteEndpointMiddlewares(stack)
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
	addOpDeleteEndpointValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteEndpoint(options.Region), middleware.Before)
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
			OperationName: "DeleteEndpoint",
			Err:           err,
		}
	}
	out := result.(*DeleteEndpointOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteEndpointInput struct {

	// The Amazon Resource Number (ARN) of the endpoint being deleted.
	//
	// This member is required.
	EndpointArn *string
}

type DeleteEndpointOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteEndpointMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteEndpoint{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteEndpoint{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteEndpoint(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "DeleteEndpoint",
	}
}
