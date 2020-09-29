// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates details for the specified Amazon Chime Voice Connector.
func (c *Client) UpdateVoiceConnector(ctx context.Context, params *UpdateVoiceConnectorInput, optFns ...func(*Options)) (*UpdateVoiceConnectorOutput, error) {
	stack := middleware.NewStack("UpdateVoiceConnector", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateVoiceConnectorMiddlewares(stack)
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
	addOpUpdateVoiceConnectorValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVoiceConnector(options.Region), middleware.Before)
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
			OperationName: "UpdateVoiceConnector",
			Err:           err,
		}
	}
	out := result.(*UpdateVoiceConnectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVoiceConnectorInput struct {
	// The name of the Amazon Chime Voice Connector.
	Name *string
	// When enabled, requires encryption for the Amazon Chime Voice Connector.
	RequireEncryption *bool
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId *string
}

type UpdateVoiceConnectorOutput struct {
	// The updated Amazon Chime Voice Connector details.
	VoiceConnector *types.VoiceConnector

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateVoiceConnectorMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateVoiceConnector{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateVoiceConnector{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateVoiceConnector(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "UpdateVoiceConnector",
	}
}