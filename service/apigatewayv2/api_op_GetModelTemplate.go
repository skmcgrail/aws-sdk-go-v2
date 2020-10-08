// Code generated by smithy-go-codegen DO NOT EDIT.

package apigatewayv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a model template.
func (c *Client) GetModelTemplate(ctx context.Context, params *GetModelTemplateInput, optFns ...func(*Options)) (*GetModelTemplateOutput, error) {
	stack := middleware.NewStack("GetModelTemplate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetModelTemplateMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpGetModelTemplateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetModelTemplate(options.Region), middleware.Before)
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
			OperationName: "GetModelTemplate",
			Err:           err,
		}
	}
	out := result.(*GetModelTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetModelTemplateInput struct {

	// The API identifier.
	//
	// This member is required.
	ApiId *string

	// The model ID.
	//
	// This member is required.
	ModelId *string
}

type GetModelTemplateOutput struct {

	// The template value.
	Value *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetModelTemplateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetModelTemplate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetModelTemplate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetModelTemplate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetModelTemplate",
	}
}
