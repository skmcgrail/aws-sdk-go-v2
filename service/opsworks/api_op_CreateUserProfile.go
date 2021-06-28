// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworks

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new user profile. Required Permissions: To use this action, an IAM
// user must have an attached policy that explicitly grants permissions. For more
// information about user permissions, see Managing User Permissions
// (https://docs.aws.amazon.com/opsworks/latest/userguide/opsworks-security-users.html).
func (c *Client) CreateUserProfile(ctx context.Context, params *CreateUserProfileInput, optFns ...func(*Options)) (*CreateUserProfileOutput, error) {
	if params == nil {
		params = &CreateUserProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateUserProfile", params, optFns, c.addOperationCreateUserProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateUserProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateUserProfileInput struct {

	// The user's IAM ARN; this can also be a federated user's ARN.
	//
	// This member is required.
	IamUserArn *string

	// Whether users can specify their own SSH public key through the My Settings page.
	// For more information, see Setting an IAM User's Public SSH Key
	// (https://docs.aws.amazon.com/opsworks/latest/userguide/security-settingsshkey.html).
	AllowSelfManagement *bool

	// The user's public SSH key.
	SshPublicKey *string

	// The user's SSH user name. The allowable characters are [a-z], [A-Z], [0-9], '-',
	// and '_'. If the specified name includes other punctuation marks, AWS OpsWorks
	// Stacks removes them. For example, my.name will be changed to myname. If you do
	// not specify an SSH user name, AWS OpsWorks Stacks generates one from the IAM
	// user name.
	SshUsername *string

	noSmithyDocumentSerde
}

// Contains the response to a CreateUserProfile request.
type CreateUserProfileOutput struct {

	// The user's IAM ARN.
	IamUserArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateUserProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateUserProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateUserProfile{}, middleware.After)
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
	if err = addOpCreateUserProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateUserProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateUserProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "opsworks",
		OperationName: "CreateUserProfile",
	}
}
