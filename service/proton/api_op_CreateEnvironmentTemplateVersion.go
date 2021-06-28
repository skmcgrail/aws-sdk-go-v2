// Code generated by smithy-go-codegen DO NOT EDIT.

package proton

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/proton/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Create a new major or minor version of an environment template. A major version
// of an environment template is a version that isn't backwards compatible. A minor
// version of an environment template is a version that's backwards compatible
// within its major version.
func (c *Client) CreateEnvironmentTemplateVersion(ctx context.Context, params *CreateEnvironmentTemplateVersionInput, optFns ...func(*Options)) (*CreateEnvironmentTemplateVersionOutput, error) {
	if params == nil {
		params = &CreateEnvironmentTemplateVersionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEnvironmentTemplateVersion", params, optFns, c.addOperationCreateEnvironmentTemplateVersionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEnvironmentTemplateVersionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEnvironmentTemplateVersionInput struct {

	// An object that includes the template bundle S3 bucket path and name for the new
	// version of an template.
	//
	// This member is required.
	Source types.TemplateVersionSourceInput

	// The name of the environment template.
	//
	// This member is required.
	TemplateName *string

	// When included, if two identicial requests are made with the same client token,
	// AWS Proton returns the environment template version that the first request
	// created.
	ClientToken *string

	// A description of the new version of an environment template.
	Description *string

	// To create a new minor version of the environment template, include a
	// majorVersion. To create a new major and minor version of the environment
	// template, exclude majorVersion.
	MajorVersion *string

	// Create tags for a new version of an environment template.
	Tags []types.Tag

	noSmithyDocumentSerde
}

type CreateEnvironmentTemplateVersionOutput struct {

	// The environment template detail data that's returned by AWS Proton.
	//
	// This member is required.
	EnvironmentTemplateVersion *types.EnvironmentTemplateVersion

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEnvironmentTemplateVersionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateEnvironmentTemplateVersion{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateEnvironmentTemplateVersion{}, middleware.After)
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
	if err = addOpCreateEnvironmentTemplateVersionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEnvironmentTemplateVersion(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateEnvironmentTemplateVersion(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "proton",
		OperationName: "CreateEnvironmentTemplateVersion",
	}
}
