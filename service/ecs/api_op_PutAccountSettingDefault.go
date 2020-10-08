// Code generated by smithy-go-codegen DO NOT EDIT.

package ecs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies an account setting for all IAM users on an account for whom no
// individual account setting has been specified. Account settings are set on a
// per-Region basis.
func (c *Client) PutAccountSettingDefault(ctx context.Context, params *PutAccountSettingDefaultInput, optFns ...func(*Options)) (*PutAccountSettingDefaultOutput, error) {
	stack := middleware.NewStack("PutAccountSettingDefault", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutAccountSettingDefaultMiddlewares(stack)
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
	addOpPutAccountSettingDefaultValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutAccountSettingDefault(options.Region), middleware.Before)
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
			OperationName: "PutAccountSettingDefault",
			Err:           err,
		}
	}
	out := result.(*PutAccountSettingDefaultOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutAccountSettingDefaultInput struct {

	// The resource name for which to modify the account setting. If
	// serviceLongArnFormat is specified, the ARN for your Amazon ECS services is
	// affected. If taskLongArnFormat is specified, the ARN and resource ID for your
	// Amazon ECS tasks is affected. If containerInstanceLongArnFormat is specified,
	// the ARN and resource ID for your Amazon ECS container instances is affected. If
	// awsvpcTrunking is specified, the ENI limit for your Amazon ECS container
	// instances is affected. If containerInsights is specified, the default setting
	// for CloudWatch Container Insights for your clusters is affected.
	//
	// This member is required.
	Name types.SettingName

	// The account setting value for the specified principal ARN. Accepted values are
	// enabled and disabled.
	//
	// This member is required.
	Value *string
}

type PutAccountSettingDefaultOutput struct {

	// The current account setting for a resource.
	Setting *types.Setting

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutAccountSettingDefaultMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutAccountSettingDefault{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutAccountSettingDefault{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutAccountSettingDefault(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecs",
		OperationName: "PutAccountSettingDefault",
	}
}
