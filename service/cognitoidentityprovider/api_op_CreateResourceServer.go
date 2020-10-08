// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new OAuth2.0 resource server and defines custom scopes in it.
func (c *Client) CreateResourceServer(ctx context.Context, params *CreateResourceServerInput, optFns ...func(*Options)) (*CreateResourceServerOutput, error) {
	stack := middleware.NewStack("CreateResourceServer", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateResourceServerMiddlewares(stack)
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
	addOpCreateResourceServerValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateResourceServer(options.Region), middleware.Before)
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
			OperationName: "CreateResourceServer",
			Err:           err,
		}
	}
	out := result.(*CreateResourceServerOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateResourceServerInput struct {

	// A unique resource server identifier for the resource server. This could be an
	// HTTPS endpoint where the resource server is located. For example,
	// https://my-weather-api.example.com.
	//
	// This member is required.
	Identifier *string

	// A friendly name for the resource server.
	//
	// This member is required.
	Name *string

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// A list of scopes. Each scope is map, where the keys are name and description.
	Scopes []*types.ResourceServerScopeType
}

type CreateResourceServerOutput struct {

	// The newly created resource server.
	//
	// This member is required.
	ResourceServer *types.ResourceServerType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateResourceServerMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateResourceServer{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateResourceServer{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateResourceServer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "CreateResourceServer",
	}
}
