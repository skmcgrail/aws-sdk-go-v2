// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentity

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Unlinks a DeveloperUserIdentifier from an existing identity. Unlinked developer
// users will be considered new identities next time they are seen. If, for a given
// Cognito identity, you remove all federated identities as well as the developer
// user identifier, the Cognito identity becomes inaccessible. You must use AWS
// Developer credentials to call this API.
func (c *Client) UnlinkDeveloperIdentity(ctx context.Context, params *UnlinkDeveloperIdentityInput, optFns ...func(*Options)) (*UnlinkDeveloperIdentityOutput, error) {
	stack := middleware.NewStack("UnlinkDeveloperIdentity", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUnlinkDeveloperIdentityMiddlewares(stack)
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
	addOpUnlinkDeveloperIdentityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUnlinkDeveloperIdentity(options.Region), middleware.Before)
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
			OperationName: "UnlinkDeveloperIdentity",
			Err:           err,
		}
	}
	out := result.(*UnlinkDeveloperIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input to the UnlinkDeveloperIdentity action.
type UnlinkDeveloperIdentityInput struct {

	// The "domain" by which Cognito will refer to your users.
	//
	// This member is required.
	DeveloperProviderName *string

	// A unique ID used by your backend authentication process to identify a user.
	//
	// This member is required.
	DeveloperUserIdentifier *string

	// A unique identifier in the format REGION:GUID.
	//
	// This member is required.
	IdentityId *string

	// An identity pool ID in the format REGION:GUID.
	//
	// This member is required.
	IdentityPoolId *string
}

type UnlinkDeveloperIdentityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUnlinkDeveloperIdentityMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUnlinkDeveloperIdentity{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUnlinkDeveloperIdentity{}, middleware.After)
}

func newServiceMetadataMiddleware_opUnlinkDeveloperIdentity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-identity",
		OperationName: "UnlinkDeveloperIdentity",
	}
}
