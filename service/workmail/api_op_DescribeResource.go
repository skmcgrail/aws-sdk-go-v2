// Code generated by smithy-go-codegen DO NOT EDIT.

package workmail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workmail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Returns the data available for the resource.
func (c *Client) DescribeResource(ctx context.Context, params *DescribeResourceInput, optFns ...func(*Options)) (*DescribeResourceOutput, error) {
	stack := middleware.NewStack("DescribeResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeResourceMiddlewares(stack)
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
	addOpDescribeResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeResource(options.Region), middleware.Before)
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
			OperationName: "DescribeResource",
			Err:           err,
		}
	}
	out := result.(*DescribeResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeResourceInput struct {

	// The identifier associated with the organization for which the resource is
	// described.
	//
	// This member is required.
	OrganizationId *string

	// The identifier of the resource to be described.
	//
	// This member is required.
	ResourceId *string
}

type DescribeResourceOutput struct {

	// The booking options for the described resource.
	BookingOptions *types.BookingOptions

	// The date and time when a resource was disabled from WorkMail, in UNIX epoch time
	// format.
	DisabledDate *time.Time

	// The email of the described resource.
	Email *string

	// The date and time when a resource was enabled for WorkMail, in UNIX epoch time
	// format.
	EnabledDate *time.Time

	// The name of the described resource.
	Name *string

	// The identifier of the described resource.
	ResourceId *string

	// The state of the resource: enabled (registered to Amazon WorkMail) or disabled
	// (deregistered or never registered to WorkMail).
	State types.EntityState

	// The type of the described resource.
	Type types.ResourceType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workmail",
		OperationName: "DescribeResource",
	}
}
