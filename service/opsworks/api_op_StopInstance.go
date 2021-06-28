// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Stops a specified instance. When you stop a standard instance, the data
// disappears and must be reinstalled when you restart the instance. You can stop
// an Amazon EBS-backed instance without losing data. For more information, see
// Starting, Stopping, and Rebooting Instances
// (https://docs.aws.amazon.com/opsworks/latest/userguide/workinginstances-starting.html).
// Required Permissions: To use this action, an IAM user must have a Manage
// permissions level for the stack, or an attached policy that explicitly grants
// permissions. For more information on user permissions, see Managing User
// Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) StopInstance(ctx context.Context, params *StopInstanceInput, optFns ...func(*Options)) (*StopInstanceOutput, error) {
	if params == nil {
		params = &StopInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopInstance", params, optFns, c.addOperationStopInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopInstanceInput struct {

	// The instance ID.
	//
	// This member is required.
	InstanceId *string

	// Specifies whether to force an instance to stop. If the instance's root device
	// type is ebs, or EBS-backed, adding the Force parameter to the StopInstances API
	// call disassociates the AWS OpsWorks Stacks instance from EC2, and forces
	// deletion of only the OpsWorks Stacks instance. You must also delete the
	// formerly-associated instance in EC2 after troubleshooting and replacing the AWS
	// OpsWorks Stacks instance with a new one.
	Force *bool

	noSmithyDocumentSerde
}

type StopInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationStopInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopInstance{}, middleware.After)
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
	if err = addOpStopInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opStopInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opStopInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "StopInstance",
	}
}
