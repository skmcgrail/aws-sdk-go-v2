// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codepipeline/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the metadata, structure, stages, and actions of a pipeline. Can be used
// to return the entire structure of a pipeline in JSON format, which can then be
// modified and used to update the pipeline structure with UpdatePipeline.
func (c *Client) GetPipeline(ctx context.Context, params *GetPipelineInput, optFns ...func(*Options)) (*GetPipelineOutput, error) {
	if params == nil {
		params = &GetPipelineInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetPipeline", params, optFns, c.addOperationGetPipelineMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetPipelineOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetPipeline action.
type GetPipelineInput struct {

	// The name of the pipeline for which you want to get information. Pipeline names
	// must be unique under an AWS user account.
	//
	// This member is required.
	Name *string

	// The version number of the pipeline. If you do not specify a version, defaults to
	// the current version.
	Version *int32

	noSmithyDocumentSerde
}

// Represents the output of a GetPipeline action.
type GetPipelineOutput struct {

	// Represents the pipeline metadata information returned as part of the output of a
	// GetPipeline action.
	Metadata *types.PipelineMetadata

	// Represents the structure of actions and stages to be performed in the pipeline.
	Pipeline *types.PipelineDeclaration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetPipelineMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetPipeline{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetPipeline{}, middleware.After)
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
	if err = addOpGetPipelineValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetPipeline(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetPipeline(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "GetPipeline",
	}
}
