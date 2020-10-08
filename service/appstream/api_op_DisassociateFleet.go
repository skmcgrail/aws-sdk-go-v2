// Code generated by smithy-go-codegen DO NOT EDIT.

package appstream

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates the specified fleet from the specified stack.
func (c *Client) DisassociateFleet(ctx context.Context, params *DisassociateFleetInput, optFns ...func(*Options)) (*DisassociateFleetOutput, error) {
	stack := middleware.NewStack("DisassociateFleet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDisassociateFleetMiddlewares(stack)
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
	addOpDisassociateFleetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateFleet(options.Region), middleware.Before)
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
			OperationName: "DisassociateFleet",
			Err:           err,
		}
	}
	out := result.(*DisassociateFleetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateFleetInput struct {

	// The name of the fleet.
	//
	// This member is required.
	FleetName *string

	// The name of the stack.
	//
	// This member is required.
	StackName *string
}

type DisassociateFleetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDisassociateFleetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateFleet{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateFleet{}, middleware.After)
}

func newServiceMetadataMiddleware_opDisassociateFleet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appstream",
		OperationName: "DisassociateFleet",
	}
}
