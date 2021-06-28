// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Tests if a specified principal is authorized to perform an AWS IoT action on a
// specified resource. Use this to test and debug the authorization behavior of
// devices that connect to the AWS IoT device gateway.
func (c *Client) TestAuthorization(ctx context.Context, params *TestAuthorizationInput, optFns ...func(*Options)) (*TestAuthorizationOutput, error) {
	if params == nil {
		params = &TestAuthorizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "TestAuthorization", params, optFns, c.addOperationTestAuthorizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*TestAuthorizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type TestAuthorizationInput struct {

	// A list of authorization info objects. Simulating authorization will create a
	// response for each authInfo object in the list.
	//
	// This member is required.
	AuthInfos []types.AuthInfo

	// The MQTT client ID.
	ClientId *string

	// The Cognito identity pool ID.
	CognitoIdentityPoolId *string

	// When testing custom authorization, the policies specified here are treated as if
	// they are attached to the principal being authorized.
	PolicyNamesToAdd []string

	// When testing custom authorization, the policies specified here are treated as if
	// they are not attached to the principal being authorized.
	PolicyNamesToSkip []string

	// The principal. Valid principals are CertificateArn
	// (arn:aws:iot:region:accountId:cert/certificateId), thingGroupArn
	// (arn:aws:iot:region:accountId:thinggroup/groupName) and CognitoId (region:id).
	Principal *string

	noSmithyDocumentSerde
}

type TestAuthorizationOutput struct {

	// The authentication results.
	AuthResults []types.AuthResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationTestAuthorizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpTestAuthorization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpTestAuthorization{}, middleware.After)
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
	if err = addOpTestAuthorizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opTestAuthorization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opTestAuthorization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "TestAuthorization",
	}
}
