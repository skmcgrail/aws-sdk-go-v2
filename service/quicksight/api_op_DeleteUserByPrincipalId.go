// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a user identified by its principal ID.
func (c *Client) DeleteUserByPrincipalId(ctx context.Context, params *DeleteUserByPrincipalIdInput, optFns ...func(*Options)) (*DeleteUserByPrincipalIdOutput, error) {
	stack := middleware.NewStack("DeleteUserByPrincipalId", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteUserByPrincipalIdMiddlewares(stack)
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
	addOpDeleteUserByPrincipalIdValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUserByPrincipalId(options.Region), middleware.Before)
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
			OperationName: "DeleteUserByPrincipalId",
			Err:           err,
		}
	}
	out := result.(*DeleteUserByPrincipalIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DeleteUserByPrincipalIdInput struct {

	// The ID for the AWS account that the user is in. Currently, you use the ID for
	// the AWS account that contains your Amazon QuickSight account.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace. Currently, you should set this to default.
	//
	// This member is required.
	Namespace *string

	// The principal ID of the user.
	//
	// This member is required.
	PrincipalId *string
}

type DeleteUserByPrincipalIdOutput struct {

	// The AWS request ID for this operation.
	RequestId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteUserByPrincipalIdMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteUserByPrincipalId{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteUserByPrincipalId{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteUserByPrincipalId(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteUserByPrincipalId",
	}
}
