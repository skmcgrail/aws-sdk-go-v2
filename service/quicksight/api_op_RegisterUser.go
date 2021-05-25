// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an Amazon QuickSight user, whose identity is associated with the AWS
// Identity and Access Management (IAM) identity or role specified in the request.
func (c *Client) RegisterUser(ctx context.Context, params *RegisterUserInput, optFns ...func(*Options)) (*RegisterUserOutput, error) {
	if params == nil {
		params = &RegisterUserInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterUser", params, optFns, addOperationRegisterUserMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterUserInput struct {

	// The ID for the AWS account that the user is in. Currently, you use the ID for
	// the AWS account that contains your Amazon QuickSight account.
	//
	// This member is required.
	AwsAccountId *string

	// The email address of the user that you want to register.
	//
	// This member is required.
	Email *string

	// Amazon QuickSight supports several ways of managing the identity of users. This
	// parameter accepts two values:
	//
	// * IAM: A user whose identity maps to an existing
	// IAM user or role.
	//
	// * QUICKSIGHT: A user whose identity is owned and managed
	// internally by Amazon QuickSight.
	//
	// This member is required.
	IdentityType types.IdentityType

	// The namespace. Currently, you should set this to default.
	//
	// This member is required.
	Namespace *string

	// The Amazon QuickSight role for the user. The user role can be one of the
	// following:
	//
	// * READER: A user who has read-only access to dashboards.
	//
	// * AUTHOR:
	// A user who can create data sources, datasets, analyses, and dashboards.
	//
	// *
	// ADMIN: A user who is an author, who can also manage Amazon QuickSight
	// settings.
	//
	// * RESTRICTED_READER: This role isn't currently available for use.
	//
	// *
	// RESTRICTED_AUTHOR: This role isn't currently available for use.
	//
	// This member is required.
	UserRole types.UserRole

	// The URL of the custom OpenID Connect (OIDC) provider that provides identity to
	// let a user federate into QuickSight with an associated AWS Identity and Access
	// Management (IAM) role. This parameter should only be used when
	// ExternalLoginFederationProviderType parameter is set to CUSTOM_OIDC.
	CustomFederationProviderUrl *string

	// (Enterprise edition only) The name of the custom permissions profile that you
	// want to assign to this user. Customized permissions allows you to control a
	// user's access by restricting access the following operations:
	//
	// * Create and
	// update data sources
	//
	// * Create and update datasets
	//
	// * Create and update email
	// reports
	//
	// * Subscribe to email reports
	//
	// To add custom permissions to an existing
	// user, use UpdateUser instead. A set of custom permissions includes any
	// combination of these restrictions. Currently, you need to create the profile
	// names for custom permission sets by using the QuickSight console. Then, you use
	// the RegisterUser API operation to assign the named set of permissions to a
	// QuickSight user. QuickSight custom permissions are applied through IAM policies.
	// Therefore, they override the permissions typically granted by assigning
	// QuickSight users to one of the default security cohorts in QuickSight (admin,
	// author, reader). This feature is available only to QuickSight Enterprise edition
	// subscriptions that use SAML 2.0-Based Federation for Single Sign-On (SSO).
	CustomPermissionsName *string

	// The type of supported external login provider that provides identity to let a
	// user federate into Amazon QuickSight with an associated AWS Identity and Access
	// Management (IAM) role. The type of supported external login provider can be one
	// of the following.
	//
	// * COGNITO: Amazon Cognito. The provider URL is
	// cognito-identity.amazonaws.com. When choosing the COGNITO provider type, don’t
	// use the "CustomFederationProviderUrl" parameter which is only needed when the
	// external provider is custom.
	//
	// * CUSTOM_OIDC: Custom OpenID Connect (OIDC)
	// provider. When choosing CUSTOM_OIDC type, use the CustomFederationProviderUrl
	// parameter to provide the custom OIDC provider URL.
	ExternalLoginFederationProviderType *string

	// The identity ID for a user in the external login provider.
	ExternalLoginId *string

	// The ARN of the IAM user or role that you are registering with Amazon QuickSight.
	IamArn *string

	// You need to use this parameter only when you register one or more users using an
	// assumed IAM role. You don't need to provide the session name for other
	// scenarios, for example when you are registering an IAM user or an Amazon
	// QuickSight user. You can register multiple users using the same IAM role if each
	// user has a different session name. For more information on assuming IAM roles,
	// see assume-role
	// (https://docs.aws.amazon.com/cli/latest/reference/sts/assume-role.html) in the
	// AWS CLI Reference.
	SessionName *string

	// The Amazon QuickSight user name that you want to create for the user you are
	// registering.
	UserName *string
}

type RegisterUserOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The user's user name.
	User *types.User

	// The URL the user visits to complete registration and provide a password. This is
	// returned only for users with an identity type of QUICKSIGHT.
	UserInvitationUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterUserMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRegisterUser{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRegisterUser{}, middleware.After)
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
	if err = addOpRegisterUserValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterUser(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRegisterUser(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "RegisterUser",
	}
}
