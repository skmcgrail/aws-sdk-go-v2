// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the data for a thing.
func (c *Client) UpdateThing(ctx context.Context, params *UpdateThingInput, optFns ...func(*Options)) (*UpdateThingOutput, error) {
	stack := middleware.NewStack("UpdateThing", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpUpdateThingMiddlewares(stack)
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
	addOpUpdateThingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateThing(options.Region), middleware.Before)
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
			OperationName: "UpdateThing",
			Err:           err,
		}
	}
	out := result.(*UpdateThingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the UpdateThing operation.
type UpdateThingInput struct {
	// The name of the thing type.
	ThingTypeName *string
	// The expected version of the thing record in the registry. If the version of the
	// record in the registry does not match the expected version specified in the
	// request, the UpdateThing request is rejected with a VersionConflictException.
	ExpectedVersion *int64
	// The name of the thing to update. You can't change a thing's name. To change a
	// thing's name, you must create a new thing, give it the new name, and then delete
	// the old thing.
	ThingName *string
	// A list of thing attributes, a JSON string containing name-value pairs. For
	// example: {\"attributes\":{\"name1\":\"value2\"}} This data is used to add new
	// attributes or update existing attributes.
	AttributePayload *types.AttributePayload
	// Remove a thing type association. If true, the association is removed.
	RemoveThingType *bool
}

// The output from the UpdateThing operation.
type UpdateThingOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpUpdateThingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpUpdateThing{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateThing{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateThing(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "UpdateThing",
	}
}