// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a user in AWS CodeStar and the user attributes across all projects.
func (c *Client) DescribeUserProfile(ctx context.Context, params *DescribeUserProfileInput, optFns ...func(*Options)) (*DescribeUserProfileOutput, error) {
	if params == nil {
		params = &DescribeUserProfileInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeUserProfile", params, optFns, c.addOperationDescribeUserProfileMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeUserProfileOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserProfileInput struct {

	// The Amazon Resource Name (ARN) of the user.
	//
	// This member is required.
	UserArn *string

	noSmithyDocumentSerde
}

type DescribeUserProfileOutput struct {

	// The date and time when the user profile was created in AWS CodeStar, in
	// timestamp format.
	//
	// This member is required.
	CreatedTimestamp *time.Time

	// The date and time when the user profile was last modified, in timestamp format.
	//
	// This member is required.
	LastModifiedTimestamp *time.Time

	// The Amazon Resource Name (ARN) of the user.
	//
	// This member is required.
	UserArn *string

	// The display name shown for the user in AWS CodeStar projects. For example, this
	// could be set to both first and last name ("Mary Major") or a single name
	// ("Mary"). The display name is also used to generate the initial icon associated
	// with the user in AWS CodeStar projects. If spaces are included in the display
	// name, the first character that appears after the space will be used as the
	// second character in the user initial icon. The initial icon displays a maximum
	// of two characters, so a display name with more than one space (for example "Mary
	// Jane Major") would generate an initial icon using the first character and the
	// first character after the space ("MJ", not "MM").
	DisplayName *string

	// The email address for the user. Optional.
	EmailAddress *string

	// The SSH public key associated with the user. This SSH public key is associated
	// with the user profile, and can be used in conjunction with the associated
	// private key for access to project resources, such as Amazon EC2 instances, if a
	// project owner grants remote access to those resources.
	SshPublicKey *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeUserProfileMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUserProfile{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUserProfile{}, middleware.After)
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
	if err = addOpDescribeUserProfileValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUserProfile(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeUserProfile(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar",
		OperationName: "DescribeUserProfile",
	}
}
