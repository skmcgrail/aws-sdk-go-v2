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

// Creates an authorizer.
func (c *Client) CreateAuthorizer(ctx context.Context, params *CreateAuthorizerInput, optFns ...func(*Options)) (*CreateAuthorizerOutput, error) {
	stack := middleware.NewStack("CreateAuthorizer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateAuthorizerMiddlewares(stack)
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
	addOpCreateAuthorizerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAuthorizer(options.Region), middleware.Before)
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
			OperationName: "CreateAuthorizer",
			Err:           err,
		}
	}
	out := result.(*CreateAuthorizerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAuthorizerInput struct {

	// The ARN of the authorizer's Lambda function.
	//
	// This member is required.
	AuthorizerFunctionArn *string

	// The authorizer name.
	//
	// This member is required.
	AuthorizerName *string

	// Specifies whether AWS IoT validates the token signature in an authorization
	// request.
	SigningDisabled *bool

	// The status of the create authorizer request.
	Status types.AuthorizerStatus

	// Metadata which can be used to manage the custom authorizer. For URI Request
	// parameters use format: ...key1=value1&key2=value2... For the CLI command-line
	// parameter use format: &&tags "key1=value1&key2=value2..." For the cli-input-json
	// file use format: "tags": "key1=value1&key2=value2..."
	Tags []*types.Tag

	// The name of the token key used to extract the token from the HTTP headers.
	TokenKeyName *string

	// The public keys used to verify the digital signature returned by your custom
	// authentication service.
	TokenSigningPublicKeys map[string]*string
}

type CreateAuthorizerOutput struct {

	// The authorizer ARN.
	AuthorizerArn *string

	// The authorizer's name.
	AuthorizerName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateAuthorizerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateAuthorizer{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAuthorizer{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateAuthorizer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "CreateAuthorizer",
	}
}
