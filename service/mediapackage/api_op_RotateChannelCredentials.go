// Code generated by smithy-go-codegen DO NOT EDIT.

package mediapackage

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediapackage/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Changes the Channel's first IngestEndpoint's username and password. WARNING -
// This API is deprecated. Please use RotateIngestEndpointCredentials instead
//
// Deprecated: This API is deprecated. Please use RotateIngestEndpointCredentials
// instead
func (c *Client) RotateChannelCredentials(ctx context.Context, params *RotateChannelCredentialsInput, optFns ...func(*Options)) (*RotateChannelCredentialsOutput, error) {
	if params == nil {
		params = &RotateChannelCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RotateChannelCredentials", params, optFns, c.addOperationRotateChannelCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RotateChannelCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RotateChannelCredentialsInput struct {

	// The ID of the channel to update.
	//
	// This member is required.
	Id *string

	noSmithyDocumentSerde
}

type RotateChannelCredentialsOutput struct {

	// The Amazon Resource Name (ARN) assigned to the Channel.
	Arn *string

	// A short text description of the Channel.
	Description *string

	// Configure egress access logging.
	EgressAccessLogs *types.EgressAccessLogs

	// An HTTP Live Streaming (HLS) ingest resource configuration.
	HlsIngest *types.HlsIngest

	// The ID of the Channel.
	Id *string

	// Configure ingress access logging.
	IngressAccessLogs *types.IngressAccessLogs

	// A collection of tags associated with a resource
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRotateChannelCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRotateChannelCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRotateChannelCredentials{}, middleware.After)
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
	if err = addOpRotateChannelCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRotateChannelCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRotateChannelCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediapackage",
		OperationName: "RotateChannelCredentials",
	}
}
