// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new BasePathMapping resource.
func (c *Client) CreateBasePathMapping(ctx context.Context, params *CreateBasePathMappingInput, optFns ...func(*Options)) (*CreateBasePathMappingOutput, error) {
	if params == nil {
		params = &CreateBasePathMappingInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateBasePathMapping", params, optFns, c.addOperationCreateBasePathMappingMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateBasePathMappingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests API Gateway to create a new BasePathMapping resource.
type CreateBasePathMappingInput struct {

	// [Required] The domain name of the BasePathMapping resource to create.
	//
	// This member is required.
	DomainName *string

	// [Required] The string identifier of the associated RestApi.
	//
	// This member is required.
	RestApiId *string

	// The base path name that callers of the API must provide as part of the URL after
	// the domain name. This value must be unique for all of the mappings across a
	// single API. Specify '(none)' if you do not want callers to specify a base path
	// name after the domain name.
	BasePath *string

	// The name of the API's stage that you want to use for this mapping. Specify
	// '(none)' if you want callers to explicitly specify the stage name after any base
	// path name.
	Stage *string

	noSmithyDocumentSerde
}

// Represents the base path that callers of the API must provide as part of the URL
// after the domain name. A custom domain name plus a BasePathMapping specification
// identifies a deployed RestApi in a given stage of the owner Account. Use Custom
// Domain Names
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/how-to-custom-domains.html)
type CreateBasePathMappingOutput struct {

	// The base path name that callers of the API must provide as part of the URL after
	// the domain name.
	BasePath *string

	// The string identifier of the associated RestApi.
	RestApiId *string

	// The name of the associated stage.
	Stage *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateBasePathMappingMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateBasePathMapping{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateBasePathMapping{}, middleware.After)
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
	if err = addOpCreateBasePathMappingValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateBasePathMapping(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addAcceptHeader(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateBasePathMapping(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "CreateBasePathMapping",
	}
}
