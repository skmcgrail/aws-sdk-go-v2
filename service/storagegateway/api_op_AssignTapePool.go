// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Assigns a tape to a tape pool for archiving. The tape assigned to a pool is
// archived in the S3 storage class that is associated with the pool. When you use
// your backup application to eject the tape, the tape is archived directly into
// the S3 storage class (S3 Glacier or S3 Glacier Deep Archive) that corresponds to
// the pool.  <p>Valid Values: <code>GLACIER</code> | <code>DEEP_ARCHIVE</code>
// </p>
func (c *Client) AssignTapePool(ctx context.Context, params *AssignTapePoolInput, optFns ...func(*Options)) (*AssignTapePoolOutput, error) {
	stack := middleware.NewStack("AssignTapePool", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpAssignTapePoolMiddlewares(stack)
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
	addOpAssignTapePoolValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opAssignTapePool(options.Region), middleware.Before)
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
			OperationName: "AssignTapePool",
			Err:           err,
		}
	}
	out := result.(*AssignTapePoolOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssignTapePoolInput struct {
	// The ID of the pool that you want to add your tape to for archiving. The tape in
	// this pool is archived in the S3 storage class that is associated with the pool.
	// When you use your backup application to eject the tape, the tape is archived
	// directly into the storage class (S3 Glacier or S3 Glacier Deep Archive) that
	// corresponds to the pool.  <p>Valid Values: <code>GLACIER</code> |
	// <code>DEEP_ARCHIVE</code> </p>
	PoolId *string
	// The unique Amazon Resource Name (ARN) of the virtual tape that you want to add
	// to the tape pool.
	TapeARN *string
}

type AssignTapePoolOutput struct {
	// The unique Amazon Resource Names (ARN) of the virtual tape that was added to the
	// tape pool.
	TapeARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpAssignTapePoolMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpAssignTapePool{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssignTapePool{}, middleware.After)
}

func newServiceMetadataMiddleware_opAssignTapePool(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "AssignTapePool",
	}
}