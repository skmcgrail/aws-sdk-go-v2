// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a version of a resource definition that has already been defined.
func (c *Client) CreateResourceDefinitionVersion(ctx context.Context, params *CreateResourceDefinitionVersionInput, optFns ...func(*Options)) (*CreateResourceDefinitionVersionOutput, error) {
	if params == nil {
		params = &CreateResourceDefinitionVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateResourceDefinitionVersion", params, optFns, c.addOperationCreateResourceDefinitionVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateResourceDefinitionVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourceDefinitionVersionInput struct {

	// The ID of the resource definition.
	//
	// This member is required.
	ResourceDefinitionId *string

	// A client token used to correlate requests and responses.
	AmznClientToken *string

	// A list of resources.
	Resources []types.Resource

	noSmithyDocumentSerde
}

type CreateResourceDefinitionVersionOutput struct {

	// The ARN of the version.
	Arn *string

	// The time, in milliseconds since the epoch, when the version was created.
	CreationTimestamp *string

	// The ID of the parent definition that the version is associated with.
	Id *string

	// The ID of the version.
	Version *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateResourceDefinitionVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateResourceDefinitionVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateResourceDefinitionVersion{}, middleware.After)
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
	if err = addOpCreateResourceDefinitionVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResourceDefinitionVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateResourceDefinitionVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateResourceDefinitionVersion",
	}
}
