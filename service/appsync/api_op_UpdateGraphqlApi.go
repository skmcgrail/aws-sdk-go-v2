// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a GraphqlApi object.
func (c *Client) UpdateGraphqlApi(ctx context.Context, params *UpdateGraphqlApiInput, optFns ...func(*Options)) (*UpdateGraphqlApiOutput, error) {
	if params == nil {
		params = &UpdateGraphqlApiInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateGraphqlApi", params, optFns, c.addOperationUpdateGraphqlApiMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateGraphqlApiOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateGraphqlApiInput struct {

	// The API ID.
	//
	// This member is required.
	ApiId *string

	// The new name for the GraphqlApi object.
	//
	// This member is required.
	Name *string

	// A list of additional authentication providers for the GraphqlApi API.
	AdditionalAuthenticationProviders []types.AdditionalAuthenticationProvider

	// The new authentication type for the GraphqlApi object.
	AuthenticationType types.AuthenticationType

	// The Amazon CloudWatch Logs configuration for the GraphqlApi object.
	LogConfig *types.LogConfig

	// The OpenID Connect configuration for the GraphqlApi object.
	OpenIDConnectConfig *types.OpenIDConnectConfig

	// The new Amazon Cognito user pool configuration for the GraphqlApi object.
	UserPoolConfig *types.UserPoolConfig

	// A flag indicating whether to enable X-Ray tracing for the GraphqlApi.
	XrayEnabled bool

	noSmithyDocumentSerde
}

type UpdateGraphqlApiOutput struct {

	// The updated GraphqlApi object.
	GraphqlApi *types.GraphqlApi

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateGraphqlApiMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateGraphqlApi{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateGraphqlApi{}, middleware.After)
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
	if err = addOpUpdateGraphqlApiValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGraphqlApi(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateGraphqlApi(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "UpdateGraphqlApi",
	}
}
