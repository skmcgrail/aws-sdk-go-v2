// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the user assigned to the specific file transfer protocol-enabled
// server, as identified by its ServerId property.  <p>The response from this call
// returns the properties of the user associated with the <code>ServerId</code>
// value that was specified.</p>
func (c *Client) DescribeUser(ctx context.Context, params *DescribeUserInput, optFns ...func(*Options)) (*DescribeUserOutput, error) {
	stack := middleware.NewStack("DescribeUser", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeUserMiddlewares(stack)
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
	addOpDescribeUserValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeUser(options.Region), middleware.Before)
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
			OperationName: "DescribeUser",
			Err:           err,
		}
	}
	out := result.(*DescribeUserOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeUserInput struct {

	// A system-assigned unique identifier for a file transfer protocol-enabled server
	// that has this user assigned.
	//
	// This member is required.
	ServerId *string

	// The name of the user assigned to one or more file transfer protocol-enabled
	// servers. User names are part of the sign-in credentials to use the AWS Transfer
	// Family service and perform file transfer tasks.
	//
	// This member is required.
	UserName *string
}

type DescribeUserOutput struct {

	// A system-assigned unique identifier for a file transfer protocol-enabled server
	// that has this user assigned.
	//
	// This member is required.
	ServerId *string

	// An array containing the properties of the user account for the ServerID value
	// that you specified.
	//
	// This member is required.
	User *types.DescribedUser

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeUserMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeUser{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeUser{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeUser(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "DescribeUser",
	}
}
