// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/opsworks/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the permissions for a specified stack. Required Permissions: To use
// this action, an IAM user must have a Manage permissions level for the stack, or
// an attached policy that explicitly grants permissions. For more information on
// user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) DescribePermissions(ctx context.Context, params *DescribePermissionsInput, optFns ...func(*Options)) (*DescribePermissionsOutput, error) {
	if params == nil {
		params = &DescribePermissionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePermissions", params, optFns, c.addOperationDescribePermissionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePermissionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePermissionsInput struct {

	// The user's IAM ARN. This can also be a federated user's ARN. For more
	// information about IAM ARNs, see Using Identifiers
	// (https://docs.aws.amazon.com/IAM/latest/UserGuide/Using_Identifiers.html).
	IamUserArn *string

	// The stack ID.
	StackId *string

	noSmithyDocumentSerde
}

// Contains the response to a DescribePermissions request.
type DescribePermissionsOutput struct {

	// An array of Permission objects that describe the stack permissions.
	//
	// * If the
	// request object contains only a stack ID, the array contains a Permission object
	// with permissions for each of the stack IAM ARNs.
	//
	// * If the request object
	// contains only an IAM ARN, the array contains a Permission object with
	// permissions for each of the user's stack IDs.
	//
	// * If the request contains a stack
	// ID and an IAM ARN, the array contains a single Permission object with
	// permissions for the specified stack and IAM ARN.
	Permissions []types.Permission

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribePermissionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePermissions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePermissions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePermissions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribePermissions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "DescribePermissions",
	}
}
