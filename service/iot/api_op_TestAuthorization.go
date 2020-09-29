// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Tests if a specified principal is authorized to perform an AWS IoT action on a
// specified resource. Use this to test and debug the authorization behavior of
// devices that connect to the AWS IoT device gateway.
func (c *Client) TestAuthorization(ctx context.Context, params *TestAuthorizationInput, optFns ...func(*Options)) (*TestAuthorizationOutput, error) {
	stack := middleware.NewStack("TestAuthorization", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpTestAuthorizationMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpTestAuthorizationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opTestAuthorization(options.Region), middleware.Before)
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
			OperationName: "TestAuthorization",
			Err:           err,
		}
	}
	out := result.(*TestAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestAuthorizationInput struct {
	// When testing custom authorization, the policies specified here are treated as if
	// they are attached to the principal being authorized.
	PolicyNamesToAdd []*string
	// When testing custom authorization, the policies specified here are treated as if
	// they are not attached to the principal being authorized.
	PolicyNamesToSkip []*string
	// The principal.
	Principal *string
	// The Cognito identity pool ID.
	CognitoIdentityPoolId *string
	// The MQTT client ID.
	ClientId *string
	// A list of authorization info objects. Simulating authorization will create a
	// response for each authInfo object in the list.
	AuthInfos []*types.AuthInfo
}

type TestAuthorizationOutput struct {
	// The authentication results.
	AuthResults []*types.AuthResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpTestAuthorizationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpTestAuthorization{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpTestAuthorization{}, middleware.After)
}

func newServiceMetadataMiddleware_opTestAuthorization(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "TestAuthorization",
	}
}