// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a logging level.
func (c *Client) DeleteV2LoggingLevel(ctx context.Context, params *DeleteV2LoggingLevelInput, optFns ...func(*Options)) (*DeleteV2LoggingLevelOutput, error) {
	stack := middleware.NewStack("DeleteV2LoggingLevel", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteV2LoggingLevelMiddlewares(stack)
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
	addOpDeleteV2LoggingLevelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteV2LoggingLevel(options.Region), middleware.Before)
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
			OperationName: "DeleteV2LoggingLevel",
			Err:           err,
		}
	}
	out := result.(*DeleteV2LoggingLevelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteV2LoggingLevelInput struct {

	// The name of the resource for which you are configuring logging.
	//
	// This member is required.
	TargetName *string

	// The type of resource for which you are configuring logging. Must be THING_Group.
	//
	// This member is required.
	TargetType types.LogTargetType
}

type DeleteV2LoggingLevelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteV2LoggingLevelMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteV2LoggingLevel{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteV2LoggingLevel{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteV2LoggingLevel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DeleteV2LoggingLevel",
	}
}
