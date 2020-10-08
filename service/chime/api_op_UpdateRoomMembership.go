// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates room membership details, such as the member role, for a room in an
// Amazon Chime Enterprise account. The member role designates whether the member
// is a chat room administrator or a general chat room member. The member role can
// be updated only for user IDs.
func (c *Client) UpdateRoomMembership(ctx context.Context, params *UpdateRoomMembershipInput, optFns ...func(*Options)) (*UpdateRoomMembershipOutput, error) {
	stack := middleware.NewStack("UpdateRoomMembership", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateRoomMembershipMiddlewares(stack)
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
	addOpUpdateRoomMembershipValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRoomMembership(options.Region), middleware.Before)
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
			OperationName: "UpdateRoomMembership",
			Err:           err,
		}
	}
	out := result.(*UpdateRoomMembershipOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoomMembershipInput struct {

	// The Amazon Chime account ID.
	//
	// This member is required.
	AccountId *string

	// The member ID.
	//
	// This member is required.
	MemberId *string

	// The room ID.
	//
	// This member is required.
	RoomId *string

	// The role of the member.
	Role types.RoomMembershipRole
}

type UpdateRoomMembershipOutput struct {

	// The room membership details.
	RoomMembership *types.RoomMembership

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateRoomMembershipMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRoomMembership{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRoomMembership{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateRoomMembership(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "UpdateRoomMembership",
	}
}
