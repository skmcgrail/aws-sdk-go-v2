// Code generated by smithy-go-codegen DO NOT EDIT.

package codestar

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an IAM user to the team for an AWS CodeStar project.
func (c *Client) AssociateTeamMember(ctx context.Context, params *AssociateTeamMemberInput, optFns ...func(*Options)) (*AssociateTeamMemberOutput, error) {
	if params == nil {
		params = &AssociateTeamMemberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateTeamMember", params, optFns, c.addOperationAssociateTeamMemberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateTeamMemberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateTeamMemberInput struct {

	// The ID of the project to which you will add the IAM user.
	//
	// This member is required.
	ProjectId *string

	// The AWS CodeStar project role that will apply to this user. This role determines
	// what actions a user can take in an AWS CodeStar project.
	//
	// This member is required.
	ProjectRole *string

	// The Amazon Resource Name (ARN) for the IAM user you want to add to the AWS
	// CodeStar project.
	//
	// This member is required.
	UserArn *string

	// A user- or system-generated token that identifies the entity that requested the
	// team member association to the project. This token can be used to repeat the
	// request.
	ClientRequestToken *string

	// Whether the team member is allowed to use an SSH public/private key pair to
	// remotely access project resources, for example Amazon EC2 instances.
	RemoteAccessAllowed bool

	noSmithyDocumentSerde
}

type AssociateTeamMemberOutput struct {

	// The user- or system-generated token from the initial request that can be used to
	// repeat the request.
	ClientRequestToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateTeamMemberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateTeamMember{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateTeamMember{}, middleware.After)
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
	if err = addOpAssociateTeamMemberValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateTeamMember(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateTeamMember(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar",
		OperationName: "AssociateTeamMember",
	}
}
