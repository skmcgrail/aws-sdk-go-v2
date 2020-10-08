// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the properties of a DBProxyTargetGroup.
func (c *Client) ModifyDBProxyTargetGroup(ctx context.Context, params *ModifyDBProxyTargetGroupInput, optFns ...func(*Options)) (*ModifyDBProxyTargetGroupOutput, error) {
	stack := middleware.NewStack("ModifyDBProxyTargetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpModifyDBProxyTargetGroupMiddlewares(stack)
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
	addOpModifyDBProxyTargetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBProxyTargetGroup(options.Region), middleware.Before)
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
			OperationName: "ModifyDBProxyTargetGroup",
			Err:           err,
		}
	}
	out := result.(*ModifyDBProxyTargetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDBProxyTargetGroupInput struct {

	// The name of the new proxy to which to assign the target group.
	//
	// This member is required.
	DBProxyName *string

	// The name of the new target group to assign to the proxy.
	//
	// This member is required.
	TargetGroupName *string

	// The settings that determine the size and behavior of the connection pool for the
	// target group.
	ConnectionPoolConfig *types.ConnectionPoolConfiguration

	// The new name for the modified DBProxyTarget. An identifier must begin with a
	// letter and must contain only ASCII letters, digits, and hyphens; it can't end
	// with a hyphen or contain two consecutive hyphens.
	NewName *string
}

type ModifyDBProxyTargetGroupOutput struct {

	// The settings of the modified DBProxyTarget.
	DBProxyTargetGroup *types.DBProxyTargetGroup

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpModifyDBProxyTargetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBProxyTargetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBProxyTargetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opModifyDBProxyTargetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBProxyTargetGroup",
	}
}
