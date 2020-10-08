// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates a Type object.
func (c *Client) UpdateType(ctx context.Context, params *UpdateTypeInput, optFns ...func(*Options)) (*UpdateTypeOutput, error) {
	stack := middleware.NewStack("UpdateType", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateTypeMiddlewares(stack)
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
	addOpUpdateTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateType(options.Region), middleware.Before)
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
			OperationName: "UpdateType",
			Err:           err,
		}
	}
	out := result.(*UpdateTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateTypeInput struct {

	// The API ID.
	//
	// This member is required.
	ApiId *string

	// The new type format: SDL or JSON.
	//
	// This member is required.
	Format types.TypeDefinitionFormat

	// The new type name.
	//
	// This member is required.
	TypeName *string

	// The new definition.
	Definition *string
}

type UpdateTypeOutput struct {

	// The updated Type object.
	Type *types.Type

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateTypeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateType{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateType{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "UpdateType",
	}
}
