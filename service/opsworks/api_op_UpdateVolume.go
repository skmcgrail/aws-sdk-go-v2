// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an Amazon EBS volume's name or mount point. For more information, see
// Resource Management
// (https://docs.aws.amazon.com/opsworks/latest/userguide/resources.html). Required
// Permissions: To use this action, an IAM user must have a Manage permissions
// level for the stack, or an attached policy that explicitly grants permissions.
// For more information on user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) UpdateVolume(ctx context.Context, params *UpdateVolumeInput, optFns ...func(*Options)) (*UpdateVolumeOutput, error) {
	stack := middleware.NewStack("UpdateVolume", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateVolumeMiddlewares(stack)
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
	addOpUpdateVolumeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateVolume(options.Region), middleware.Before)
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
			OperationName: "UpdateVolume",
			Err:           err,
		}
	}
	out := result.(*UpdateVolumeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateVolumeInput struct {

	// The volume ID.
	//
	// This member is required.
	VolumeId *string

	// The new mount point.
	MountPoint *string

	// The new name.
	Name *string
}

type UpdateVolumeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateVolumeMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateVolume{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateVolume{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateVolume(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "UpdateVolume",
	}
}
