// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates read and write permissions on a dashboard.
func (c *Client) UpdateDashboardPermissions(ctx context.Context, params *UpdateDashboardPermissionsInput, optFns ...func(*Options)) (*UpdateDashboardPermissionsOutput, error) {
	stack := middleware.NewStack("UpdateDashboardPermissions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateDashboardPermissionsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateDashboardPermissionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDashboardPermissions(options.Region), middleware.Before)
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
			OperationName: "UpdateDashboardPermissions",
			Err:           err,
		}
	}
	out := result.(*UpdateDashboardPermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDashboardPermissionsInput struct {
	// The ID of the AWS account that contains the dashboard whose permissions you're
	// updating.
	AwsAccountId *string
	// The ID for the dashboard.
	DashboardId *string
	// The permissions that you want to grant on this resource.
	GrantPermissions []*types.ResourcePermission
	// The permissions that you want to revoke from this resource.
	RevokePermissions []*types.ResourcePermission
}

type UpdateDashboardPermissionsOutput struct {
	// The Amazon Resource Name (ARN) of the dashboard.
	DashboardArn *string
	// Information about the permissions on the dashboard.
	Permissions []*types.ResourcePermission
	// The ID for the dashboard.
	DashboardId *string
	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateDashboardPermissionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDashboardPermissions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDashboardPermissions{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateDashboardPermissions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "UpdateDashboardPermissions",
	}
}