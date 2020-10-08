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

// Gets an OTA update.
func (c *Client) GetOTAUpdate(ctx context.Context, params *GetOTAUpdateInput, optFns ...func(*Options)) (*GetOTAUpdateOutput, error) {
	stack := middleware.NewStack("GetOTAUpdate", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetOTAUpdateMiddlewares(stack)
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
	addOpGetOTAUpdateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetOTAUpdate(options.Region), middleware.Before)
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
			OperationName: "GetOTAUpdate",
			Err:           err,
		}
	}
	out := result.(*GetOTAUpdateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetOTAUpdateInput struct {

	// The OTA update ID.
	//
	// This member is required.
	OtaUpdateId *string
}

type GetOTAUpdateOutput struct {

	// The OTA update info.
	OtaUpdateInfo *types.OTAUpdateInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetOTAUpdateMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetOTAUpdate{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetOTAUpdate{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetOTAUpdate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "GetOTAUpdate",
	}
}
