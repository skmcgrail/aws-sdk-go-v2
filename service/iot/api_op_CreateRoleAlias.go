// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a role alias.
func (c *Client) CreateRoleAlias(ctx context.Context, params *CreateRoleAliasInput, optFns ...func(*Options)) (*CreateRoleAliasOutput, error) {
	stack := middleware.NewStack("CreateRoleAlias", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateRoleAliasMiddlewares(stack)
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
	addOpCreateRoleAliasValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateRoleAlias(options.Region), middleware.Before)
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
			OperationName: "CreateRoleAlias",
			Err:           err,
		}
	}
	out := result.(*CreateRoleAliasOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateRoleAliasInput struct {

	// The role alias that points to a role ARN. This allows you to change the role
	// without having to update the device.
	//
	// This member is required.
	RoleAlias *string

	// The role ARN.
	//
	// This member is required.
	RoleArn *string

	// How long (in seconds) the credentials will be valid.
	CredentialDurationSeconds *int32

	// Metadata which can be used to manage the role alias. For URI Request parameters
	// use format: ...key1=value1&key2=value2... For the CLI command-line parameter use
	// format: &&tags "key1=value1&key2=value2..." For the cli-input-json file use
	// format: "tags": "key1=value1&key2=value2..."
	Tags []*types.Tag
}

type CreateRoleAliasOutput struct {

	// The role alias.
	RoleAlias *string

	// The role alias ARN.
	RoleAliasArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateRoleAliasMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateRoleAlias{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateRoleAlias{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateRoleAlias(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateRoleAlias",
	}
}
