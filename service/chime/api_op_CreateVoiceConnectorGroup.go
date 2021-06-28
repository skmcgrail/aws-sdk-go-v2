// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon Chime Voice Connector group under the administrator's AWS
// account. You can associate Amazon Chime Voice Connectors with the Amazon Chime
// Voice Connector group by including VoiceConnectorItems in the request. You can
// include Amazon Chime Voice Connectors from different AWS Regions in your group.
// This creates a fault tolerant mechanism for fallback in case of availability
// events.
func (c *Client) CreateVoiceConnectorGroup(ctx context.Context, params *CreateVoiceConnectorGroupInput, optFns ...func(*Options)) (*CreateVoiceConnectorGroupOutput, error) {
	if params == nil {
		params = &CreateVoiceConnectorGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVoiceConnectorGroup", params, optFns, c.addOperationCreateVoiceConnectorGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVoiceConnectorGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVoiceConnectorGroupInput struct {

	// The name of the Amazon Chime Voice Connector group.
	//
	// This member is required.
	Name *string

	// The Amazon Chime Voice Connectors to route inbound calls to.
	VoiceConnectorItems []types.VoiceConnectorItem

	noSmithyDocumentSerde
}

type CreateVoiceConnectorGroupOutput struct {

	// The Amazon Chime Voice Connector group details.
	VoiceConnectorGroup *types.VoiceConnectorGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateVoiceConnectorGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateVoiceConnectorGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateVoiceConnectorGroup{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateVoiceConnectorGroupValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVoiceConnectorGroup(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateVoiceConnectorGroup(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "CreateVoiceConnectorGroup",
	}
}
