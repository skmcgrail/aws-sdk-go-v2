// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified virtual tape. This operation is only supported in the tape
// gateway type.
func (c *Client) DeleteTape(ctx context.Context, params *DeleteTapeInput, optFns ...func(*Options)) (*DeleteTapeOutput, error) {
	stack := middleware.NewStack("DeleteTape", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteTapeMiddlewares(stack)
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
	addOpDeleteTapeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTape(options.Region), middleware.Before)
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
			OperationName: "DeleteTape",
			Err:           err,
		}
	}
	out := result.(*DeleteTapeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DeleteTapeInput
type DeleteTapeInput struct {
	// The unique Amazon Resource Name (ARN) of the gateway that the virtual tape to
	// delete is associated with. Use the ListGateways () operation to return a list of
	// gateways for your account and AWS Region.
	GatewayARN *string
	// The Amazon Resource Name (ARN) of the virtual tape to delete.
	TapeARN *string
}

// DeleteTapeOutput
type DeleteTapeOutput struct {
	// The Amazon Resource Name (ARN) of the deleted virtual tape.
	TapeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteTapeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteTape{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteTape{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteTape(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DeleteTape",
	}
}