// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Detaches the specified principal from the specified thing. A principal can be
// X.509 certificates, IAM users, groups, and roles, Amazon Cognito identities or
// federated identities. This call is asynchronous. It might take several seconds
// for the detachment to propagate.
func (c *Client) DetachThingPrincipal(ctx context.Context, params *DetachThingPrincipalInput, optFns ...func(*Options)) (*DetachThingPrincipalOutput, error) {
	stack := middleware.NewStack("DetachThingPrincipal", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDetachThingPrincipalMiddlewares(stack)
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
	addOpDetachThingPrincipalValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDetachThingPrincipal(options.Region), middleware.Before)
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
			OperationName: "DetachThingPrincipal",
			Err:           err,
		}
	}
	out := result.(*DetachThingPrincipalOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DetachThingPrincipal operation.
type DetachThingPrincipalInput struct {
	// The name of the thing.
	ThingName *string
	// If the principal is a certificate, this value must be ARN of the certificate. If
	// the principal is an Amazon Cognito identity, this value must be the ID of the
	// Amazon Cognito identity.
	Principal *string
}

// The output from the DetachThingPrincipal operation.
type DetachThingPrincipalOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDetachThingPrincipalMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDetachThingPrincipal{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDetachThingPrincipal{}, middleware.After)
}

func newServiceMetadataMiddleware_opDetachThingPrincipal(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DetachThingPrincipal",
	}
}