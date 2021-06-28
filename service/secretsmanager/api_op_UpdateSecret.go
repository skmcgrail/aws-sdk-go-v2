// Code generated by smithy-go-codegen DO NOT EDIT.

package secretsmanager

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies many of the details of the specified secret. If you include a
// ClientRequestToken and either SecretString or SecretBinary then it also creates
// a new version attached to the secret. To modify the rotation configuration of a
// secret, use RotateSecret instead. The Secrets Manager console uses only the
// SecretString parameter and therefore limits you to encrypting and storing only a
// text string. To encrypt and store binary data as part of the version of a
// secret, you must use either the AWS CLI or one of the AWS SDKs.
//
// * If a version
// with a VersionId with the same value as the ClientRequestToken parameter already
// exists, the operation results in an error. You cannot modify an existing
// version, you can only create a new version.
//
// * If you include SecretString or
// SecretBinary to create a new secret version, Secrets Manager automatically
// attaches the staging label AWSCURRENT to the new version.
//
// * If you call an
// operation to encrypt or decrypt the SecretString or SecretBinary for a secret in
// the same account as the calling user and that secret doesn't specify a AWS KMS
// encryption key, Secrets Manager uses the account's default AWS managed customer
// master key (CMK) with the alias aws/secretsmanager. If this key doesn't already
// exist in your account then Secrets Manager creates it for you automatically. All
// users and roles in the same AWS account automatically have access to use the
// default CMK. Note that if an Secrets Manager API call results in AWS creating
// the account's AWS-managed CMK, it can result in a one-time significant delay in
// returning the result.
//
// * If the secret resides in a different AWS account from
// the credentials calling an API that requires encryption or decryption of the
// secret value then you must create and use a custom AWS KMS CMK because you can't
// access the default CMK for the account using credentials from a different AWS
// account. Store the ARN of the CMK in the secret when you create the secret or
// when you update it by including it in the KMSKeyId. If you call an API that must
// encrypt or decrypt SecretString or SecretBinary using credentials from a
// different account then the AWS KMS key policy must grant cross-account access to
// that other account's user or role for both the kms:GenerateDataKey and
// kms:Decrypt operations.
//
// Minimum permissions To run this command, you must have
// the following permissions:
//
// * secretsmanager:UpdateSecret
//
// * kms:GenerateDataKey
// - needed only if you use a custom AWS KMS key to encrypt the secret. You do not
// need this permission to use the account's AWS managed CMK for Secrets
// Manager.
//
// * kms:Decrypt - needed only if you use a custom AWS KMS key to encrypt
// the secret. You do not need this permission to use the account's AWS managed CMK
// for Secrets Manager.
//
// Related operations
//
// * To create a new secret, use
// CreateSecret.
//
// * To add only a new version to an existing secret, use
// PutSecretValue.
//
// * To get the details for a secret, use DescribeSecret.
//
// * To
// list the versions contained in a secret, use ListSecretVersionIds.
func (c *Client) UpdateSecret(ctx context.Context, params *UpdateSecretInput, optFns ...func(*Options)) (*UpdateSecretOutput, error) {
	if params == nil {
		params = &UpdateSecretInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSecret", params, optFns, c.addOperationUpdateSecretMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSecretOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSecretInput struct {

	// Specifies the secret that you want to modify or to which you want to add a new
	// version. You can specify either the Amazon Resource Name (ARN) or the friendly
	// name of the secret. If you specify an ARN, we generally recommend that you
	// specify a complete ARN. You can specify a partial ARN too—for example, if you
	// don’t include the final hyphen and six random characters that Secrets Manager
	// adds at the end of the ARN when you created the secret. A partial ARN match can
	// work as long as it uniquely matches only one secret. However, if your secret has
	// a name that ends in a hyphen followed by six characters (before Secrets Manager
	// adds the hyphen and six characters to the ARN) and you try to use that as a
	// partial ARN, then those characters cause Secrets Manager to assume that you’re
	// specifying a complete ARN. This confusion can cause unexpected results. To avoid
	// this situation, we recommend that you don’t create secret names ending with a
	// hyphen followed by six characters. If you specify an incomplete ARN without the
	// random suffix, and instead provide the 'friendly name', you must not include the
	// random suffix. If you do include the random suffix added by Secrets Manager, you
	// receive either a ResourceNotFoundException or an AccessDeniedException error,
	// depending on your permissions.
	//
	// This member is required.
	SecretId *string

	// (Optional) If you want to add a new version to the secret, this parameter
	// specifies a unique identifier for the new version that helps ensure idempotency.
	// If you use the AWS CLI or one of the AWS SDK to call this operation, then you
	// can leave this parameter empty. The CLI or SDK generates a random UUID for you
	// and includes that in the request. If you don't use the SDK and instead generate
	// a raw HTTP request to the Secrets Manager service endpoint, then you must
	// generate a ClientRequestToken yourself for new versions and include that value
	// in the request. You typically only need to interact with this value if you
	// implement your own retry logic and want to ensure that a given secret is not
	// created twice. We recommend that you generate a UUID-type
	// (https://wikipedia.org/wiki/Universally_unique_identifier) value to ensure
	// uniqueness within the specified secret. Secrets Manager uses this value to
	// prevent the accidental creation of duplicate versions if there are failures and
	// retries during the Lambda rotation function's processing.
	//
	// * If the
	// ClientRequestToken value isn't already associated with a version of the secret
	// then a new version of the secret is created.
	//
	// * If a version with this value
	// already exists and that version's SecretString and SecretBinary values are the
	// same as those in the request then the request is ignored (the operation is
	// idempotent).
	//
	// * If a version with this value already exists and that version's
	// SecretString and SecretBinary values are different from the request then an
	// error occurs because you cannot modify an existing secret value.
	//
	// This value
	// becomes the VersionId of the new version.
	ClientRequestToken *string

	// (Optional) Specifies an updated user-provided description of the secret.
	Description *string

	// (Optional) Specifies an updated ARN or alias of the AWS KMS customer master key
	// (CMK) to be used to encrypt the protected text in new versions of this secret.
	// You can only use the account's default CMK to encrypt and decrypt if you call
	// this operation using credentials from the same account that owns the secret. If
	// the secret is in a different account, then you must create a custom CMK and
	// provide the ARN of that CMK in this field. The user making the call must have
	// permissions to both the secret and the CMK in their respective accounts.
	KmsKeyId *string

	// (Optional) Specifies updated binary data that you want to encrypt and store in
	// the new version of the secret. To use this parameter in the command-line tools,
	// we recommend that you store your binary data in a file and then use the
	// appropriate technique for your tool to pass the contents of the file as a
	// parameter. Either SecretBinary or SecretString must have a value, but not both.
	// They cannot both be empty. This parameter is not accessible using the Secrets
	// Manager console.
	SecretBinary []byte

	// (Optional) Specifies updated text data that you want to encrypt and store in
	// this new version of the secret. Either SecretBinary or SecretString must have a
	// value, but not both. They cannot both be empty. If you create this secret by
	// using the Secrets Manager console then Secrets Manager puts the protected secret
	// text in only the SecretString parameter. The Secrets Manager console stores the
	// information as a JSON structure of key/value pairs that the default Lambda
	// rotation function knows how to parse. For storing multiple values, we recommend
	// that you use a JSON text string argument and specify key/value pairs. For
	// information on how to format a JSON parameter for the various command line tool
	// environments, see Using JSON for Parameters
	// (https://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#cli-using-param-json)
	// in the AWS CLI User Guide. For example:
	// [{"username":"bob"},{"password":"abc123xyz456"}] If your command-line tool or
	// SDK requires quotation marks around the parameter, you should use single quotes
	// to avoid confusion with the double quotes required in the JSON text. You can
	// also 'escape' the double quote character in the embedded JSON text by prefacing
	// each with a backslash. For example, the following string is surrounded by
	// double-quotes. All of the embedded double quotes are escaped:
	// "[{\"username\":\"bob\"},{\"password\":\"abc123xyz456\"}]"
	SecretString *string

	noSmithyDocumentSerde
}

type UpdateSecretOutput struct {

	// The ARN of the secret that was updated. Secrets Manager automatically adds
	// several random characters to the name at the end of the ARN when you initially
	// create a secret. This affects only the ARN and not the actual friendly name.
	// This ensures that if you create a new secret with the same name as an old secret
	// that you previously deleted, then users with access to the old secret don't
	// automatically get access to the new secret because the ARNs are different.
	ARN *string

	// The friendly name of the secret that was updated.
	Name *string

	// If a new version of the secret was created by this operation, then VersionId
	// contains the unique identifier of the new version.
	VersionId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSecretMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSecret{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSecret{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateSecretMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateSecretValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSecret(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateSecret struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateSecret) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateSecret) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateSecretInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateSecretInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateSecretMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateSecret{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateSecret(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "secretsmanager",
		OperationName: "UpdateSecret",
	}
}
