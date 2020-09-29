// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of all currently running remote access sessions.
func (c *Client) ListRemoteAccessSessions(ctx context.Context, params *ListRemoteAccessSessionsInput, optFns ...func(*Options)) (*ListRemoteAccessSessionsOutput, error) {
	stack := middleware.NewStack("ListRemoteAccessSessions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListRemoteAccessSessionsMiddlewares(stack)
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
	addOpListRemoteAccessSessionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListRemoteAccessSessions(options.Region), middleware.Before)
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
			OperationName: "ListRemoteAccessSessions",
			Err:           err,
		}
	}
	out := result.(*ListRemoteAccessSessionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the request to return information about the remote access session.
type ListRemoteAccessSessionsInput struct {
	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string
	// The Amazon Resource Name (ARN) of the project about which you are requesting
	// information.
	Arn *string
}

// Represents the response from the server after AWS Device Farm makes a request to
// return information about the remote access session.
type ListRemoteAccessSessionsOutput struct {
	// A container that represents the metadata from the service about each remote
	// access session you are requesting.
	RemoteAccessSessions []*types.RemoteAccessSession
	// An identifier that was returned from the previous call to this operation, which
	// can be used to return the next set of items in the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListRemoteAccessSessionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListRemoteAccessSessions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListRemoteAccessSessions{}, middleware.After)
}

func newServiceMetadataMiddleware_opListRemoteAccessSessions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "ListRemoteAccessSessions",
	}
}