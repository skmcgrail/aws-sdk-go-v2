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

// Creates a new group in the specified user pool. Calling this action requires
// developer credentials.
func (c *Client) CreateGroup(ctx context.Context, params *CreateGroupInput, optFns ...func(*Options)) (*CreateGroupOutput, error) {
	stack := middleware.NewStack("CreateGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCreateGroupMiddlewares(stack)
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
	addOpCreateGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateGroup(options.Region), middleware.Before)
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
			OperationName: "CreateGroup",
			Err:           err,
		}
	}
	out := result.(*CreateGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateGroupInput struct {
	// A nonnegative integer value that specifies the precedence of this group relative
	// to the other groups that a user can belong to in the user pool. Zero is the
	// highest precedence value. Groups with lower Precedence values take precedence
	// over groups with higher or null Precedence values. If a user belongs to two or
	// more groups, it is the group with the lowest precedence value whose role ARN
	// will be used in the cognito:roles and cognito:preferred_role claims in the
	// user's tokens. Two groups can have the same Precedence value. If this happens,
	// neither group takes precedence over the other. If two groups with the same
	// Precedence have the same role ARN, that role is used in the
	// cognito:preferred_role claim in tokens for users in each group. If the two
	// groups have different role ARNs, the cognito:preferred_role claim is not set in
	// users' tokens. The default Precedence value is null.
	Precedence *int32
	// The name of the group. Must be unique.
	GroupName *string
	// The user pool ID for the user pool.
	UserPoolId *string
	// The role ARN for the group.
	RoleArn *string
	// A string containing the description of the group.
	Description *string
}

type CreateGroupOutput struct {
	// The group object for the group.
	Group *types.GroupType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCreateGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCreateGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "CreateGroup",
	}
}