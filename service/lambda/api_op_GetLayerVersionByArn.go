// Code generated by smithy-go-codegen DO NOT EDIT.

package lambda

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a version of an AWS Lambda layer
// (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html), with a
// link to download the layer archive that's valid for 10 minutes.
func (c *Client) GetLayerVersionByArn(ctx context.Context, params *GetLayerVersionByArnInput, optFns ...func(*Options)) (*GetLayerVersionByArnOutput, error) {
	stack := middleware.NewStack("GetLayerVersionByArn", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetLayerVersionByArnMiddlewares(stack)
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
	addOpGetLayerVersionByArnValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLayerVersionByArn(options.Region), middleware.Before)
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
			OperationName: "GetLayerVersionByArn",
			Err:           err,
		}
	}
	out := result.(*GetLayerVersionByArnOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLayerVersionByArnInput struct {

	// The ARN of the layer version.
	//
	// This member is required.
	Arn *string
}

type GetLayerVersionByArnOutput struct {

	// The layer's compatible runtimes.
	CompatibleRuntimes []types.Runtime

	// Details about the layer version.
	Content *types.LayerVersionContentOutput

	// The date that the layer version was created, in ISO-8601 format
	// (https://www.w3.org/TR/NOTE-datetime) (YYYY-MM-DDThh:mm:ss.sTZD).
	CreatedDate *string

	// The description of the version.
	Description *string

	// The ARN of the layer.
	LayerArn *string

	// The ARN of the layer version.
	LayerVersionArn *string

	// The layer's software license.
	LicenseInfo *string

	// The version number.
	Version *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetLayerVersionByArnMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetLayerVersionByArn{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetLayerVersionByArn{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetLayerVersionByArn(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lambda",
		OperationName: "GetLayerVersionByArn",
	}
}
