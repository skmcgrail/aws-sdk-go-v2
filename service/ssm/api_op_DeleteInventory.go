// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete a custom inventory type, or the data associated with a custom Inventory
// type. Deleting a custom inventory type is also referred to as deleting a custom
// inventory schema.
func (c *Client) DeleteInventory(ctx context.Context, params *DeleteInventoryInput, optFns ...func(*Options)) (*DeleteInventoryOutput, error) {
	stack := middleware.NewStack("DeleteInventory", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDeleteInventoryMiddlewares(stack)
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
	addIdempotencyToken_opDeleteInventoryMiddleware(stack, options)
	addOpDeleteInventoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteInventory(options.Region), middleware.Before)
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
			OperationName: "DeleteInventory",
			Err:           err,
		}
	}
	out := result.(*DeleteInventoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteInventoryInput struct {
	// The name of the custom inventory type for which you want to delete either all
	// previously collected data, or the inventory type itself.
	TypeName *string
	// Use the SchemaDeleteOption to delete a custom inventory type (schema). If you
	// don't choose this option, the system only deletes existing inventory data
	// associated with the custom inventory type. Choose one of the following options:
	// DisableSchema: If you choose this option, the system ignores all inventory data
	// for the specified version, and any earlier versions. To enable this schema
	// again, you must call the PutInventory action for a version greater than the
	// disabled version. DeleteSchema: This option deletes the specified custom type
	// from the Inventory service. You can recreate the schema later, if you want.
	SchemaDeleteOption types.InventorySchemaDeleteOption
	// Use this option to view a summary of the deletion request without deleting any
	// data or the data type. This option is useful when you only want to understand
	// what will be deleted. Once you validate that the data to be deleted is what you
	// intend to delete, you can run the same command without specifying the DryRun
	// option.
	DryRun *bool
	// User-provided idempotency token.
	ClientToken *string
}

type DeleteInventoryOutput struct {
	// Every DeleteInventory action is assigned a unique ID. This option returns a
	// unique ID. You can use this ID to query the status of a delete operation. This
	// option is useful for ensuring that a delete operation has completed before you
	// begin other actions.
	DeletionId *string
	// A summary of the delete operation. For more information about this summary, see
	// Deleting custom inventory
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-inventory-custom.html#sysman-inventory-delete-summary)
	// in the AWS Systems Manager User Guide.
	DeletionSummary *types.InventoryDeletionSummary
	// The name of the inventory data type specified in the request.
	TypeName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDeleteInventoryMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDeleteInventory{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeleteInventory{}, middleware.After)
}

type idempotencyToken_initializeOpDeleteInventory struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpDeleteInventory) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpDeleteInventory) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*DeleteInventoryInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *DeleteInventoryInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opDeleteInventoryMiddleware(stack *middleware.Stack, cfg Options) {
	stack.Initialize.Add(&idempotencyToken_initializeOpDeleteInventory{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opDeleteInventory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DeleteInventory",
	}
}