// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists a history of user activity and any risks detected as part of Amazon
// Cognito advanced security.
func (c *Client) AdminListUserAuthEvents(ctx context.Context, params *AdminListUserAuthEventsInput, optFns ...func(*Options)) (*AdminListUserAuthEventsOutput, error) {
	stack := middleware.NewStack("AdminListUserAuthEvents", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAdminListUserAuthEventsMiddlewares(stack)
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
	addOpAdminListUserAuthEventsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAdminListUserAuthEvents(options.Region), middleware.Before)
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
			OperationName: "AdminListUserAuthEvents",
			Err:           err,
		}
	}
	out := result.(*AdminListUserAuthEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AdminListUserAuthEventsInput struct {
	// The maximum number of authentication events to return.
	MaxResults *int32
	// A pagination token.
	NextToken *string
	// The user pool username or an alias.
	Username *string
	// The user pool ID.
	UserPoolId *string
}

type AdminListUserAuthEventsOutput struct {
	// The response object. It includes the EventID, EventType, CreationDate,
	// EventRisk, and EventResponse.
	AuthEvents []*types.AuthEventType
	// A pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAdminListUserAuthEventsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAdminListUserAuthEvents{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAdminListUserAuthEvents{}, middleware.After)
}

func newServiceMetadataMiddleware_opAdminListUserAuthEvents(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "AdminListUserAuthEvents",
	}
}