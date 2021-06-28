// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the specified sending authorization policy for the given identity (an
// email address or a domain). This API returns successfully even if a policy with
// the specified name does not exist. This API is for the identity owner only. If
// you have not verified the identity, this API will return an error. Sending
// authorization is a feature that enables an identity owner to authorize other
// senders to use its identities. For information about using sending
// authorization, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization.html).
// You can execute this operation no more than once per second.
func (c *Client) UpdateEmailIdentityPolicy(ctx context.Context, params *UpdateEmailIdentityPolicyInput, optFns ...func(*Options)) (*UpdateEmailIdentityPolicyOutput, error) {
	if params == nil {
		params = &UpdateEmailIdentityPolicyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateEmailIdentityPolicy", params, optFns, c.addOperationUpdateEmailIdentityPolicyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateEmailIdentityPolicyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to update a sending authorization policy for an identity.
// Sending authorization is an Amazon SES feature that enables you to authorize
// other senders to use your identities. For information, see the Amazon SES
// Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-identity-owner-tasks-management.html).
type UpdateEmailIdentityPolicyInput struct {

	// The email identity for which you want to update policy.
	//
	// This member is required.
	EmailIdentity *string

	// The text of the policy in JSON format. The policy cannot exceed 4 KB. For
	// information about the syntax of sending authorization policies, see the Amazon
	// SES Developer Guide
	// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/sending-authorization-policies.html).
	//
	// This member is required.
	Policy *string

	// The name of the policy. The policy name cannot exceed 64 characters and can only
	// include alphanumeric characters, dashes, and underscores.
	//
	// This member is required.
	PolicyName *string

	noSmithyDocumentSerde
}

// An HTTP 200 response if the request succeeds, or an error message if the request
// fails.
type UpdateEmailIdentityPolicyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateEmailIdentityPolicyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateEmailIdentityPolicy{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateEmailIdentityPolicy{}, middleware.After)
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
	if err = addOpUpdateEmailIdentityPolicyValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateEmailIdentityPolicy(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateEmailIdentityPolicy(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "UpdateEmailIdentityPolicy",
	}
}
