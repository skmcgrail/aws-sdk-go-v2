// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a link to a currently running remote access session.
func (c *Client) GetRemoteAccessSession(ctx context.Context, params *GetRemoteAccessSessionInput, optFns ...func(*Options)) (*GetRemoteAccessSessionOutput, error) {
	if params == nil {
		params = &GetRemoteAccessSessionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRemoteAccessSession", params, optFns, c.addOperationGetRemoteAccessSessionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRemoteAccessSessionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to get information about the specified remote access
// session.
type GetRemoteAccessSessionInput struct {

	// The Amazon Resource Name (ARN) of the remote access session about which you want
	// to get session information.
	//
	// This member is required.
	Arn *string

	noSmithyDocumentSerde
}

// Represents the response from the server that lists detailed information about
// the remote access session.
type GetRemoteAccessSessionOutput struct {

	// A container that lists detailed information about the remote access session.
	RemoteAccessSession *types.RemoteAccessSession

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRemoteAccessSessionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRemoteAccessSession{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRemoteAccessSession{}, middleware.After)
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
	if err = addOpGetRemoteAccessSessionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRemoteAccessSession(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRemoteAccessSession(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "GetRemoteAccessSession",
	}
}
