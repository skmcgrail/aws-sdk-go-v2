// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified SIP credentials used by your equipment to authenticate
// during call termination.
func (c *Client) DeleteVoiceConnectorTerminationCredentials(ctx context.Context, params *DeleteVoiceConnectorTerminationCredentialsInput, optFns ...func(*Options)) (*DeleteVoiceConnectorTerminationCredentialsOutput, error) {
	stack := middleware.NewStack("DeleteVoiceConnectorTerminationCredentials", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteVoiceConnectorTerminationCredentialsMiddlewares(stack)
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
	addOpDeleteVoiceConnectorTerminationCredentialsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteVoiceConnectorTerminationCredentials(options.Region), middleware.Before)
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
			OperationName: "DeleteVoiceConnectorTerminationCredentials",
			Err:           err,
		}
	}
	out := result.(*DeleteVoiceConnectorTerminationCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteVoiceConnectorTerminationCredentialsInput struct {
	// The RFC2617 compliant username associated with the SIP credentials, in US-ASCII
	// format.
	Usernames []*string
	// The Amazon Chime Voice Connector ID.
	VoiceConnectorId *string
}

type DeleteVoiceConnectorTerminationCredentialsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteVoiceConnectorTerminationCredentialsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteVoiceConnectorTerminationCredentials{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteVoiceConnectorTerminationCredentials{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteVoiceConnectorTerminationCredentials(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DeleteVoiceConnectorTerminationCredentials",
	}
}