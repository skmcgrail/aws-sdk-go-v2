// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Removes the suspension from up to 50 previously suspended users for the
// specified Amazon Chime EnterpriseLWA account. Only users on EnterpriseLWA
// accounts can be unsuspended using this action. For more information about
// different account types, see Managing Your Amazon Chime Accounts
// (https://docs.aws.amazon.com/chime/latest/ag/manage-chime-account.html) in the
// Amazon Chime Administration Guide. Previously suspended users who are
// unsuspended using this action are returned to Registered status. Users who are
// not previously suspended are ignored.
func (c *Client) BatchUnsuspendUser(ctx context.Context, params *BatchUnsuspendUserInput, optFns ...func(*Options)) (*BatchUnsuspendUserOutput, error) {
	stack := middleware.NewStack("BatchUnsuspendUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpBatchUnsuspendUserMiddlewares(stack)
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
	addOpBatchUnsuspendUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchUnsuspendUser(options.Region), middleware.Before)
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
			OperationName: "BatchUnsuspendUser",
			Err:           err,
		}
	}
	out := result.(*BatchUnsuspendUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchUnsuspendUserInput struct {
	// The request containing the user IDs to unsuspend.
	UserIdList []*string
	// The Amazon Chime account ID.
	AccountId *string
}

type BatchUnsuspendUserOutput struct {
	// If the BatchUnsuspendUser () action fails for one or more of the user IDs in the
	// request, a list of the user IDs is returned, along with error codes and error
	// messages.
	UserErrors []*types.UserError

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpBatchUnsuspendUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpBatchUnsuspendUser{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchUnsuspendUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opBatchUnsuspendUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "BatchUnsuspendUser",
	}
}