// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a specific database snapshot in Amazon Lightsail.
func (c *Client) GetRelationalDatabaseSnapshot(ctx context.Context, params *GetRelationalDatabaseSnapshotInput, optFns ...func(*Options)) (*GetRelationalDatabaseSnapshotOutput, error) {
	stack := middleware.NewStack("GetRelationalDatabaseSnapshot", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRelationalDatabaseSnapshotMiddlewares(stack)
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
	addOpGetRelationalDatabaseSnapshotValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRelationalDatabaseSnapshot(options.Region), middleware.Before)
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
			OperationName: "GetRelationalDatabaseSnapshot",
			Err:           err,
		}
	}
	out := result.(*GetRelationalDatabaseSnapshotOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRelationalDatabaseSnapshotInput struct {

	// The name of the database snapshot for which to get information.
	//
	// This member is required.
	RelationalDatabaseSnapshotName *string
}

type GetRelationalDatabaseSnapshotOutput struct {

	// An object describing the specified database snapshot.
	RelationalDatabaseSnapshot *types.RelationalDatabaseSnapshot

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRelationalDatabaseSnapshotMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRelationalDatabaseSnapshot{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRelationalDatabaseSnapshot{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRelationalDatabaseSnapshot(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRelationalDatabaseSnapshot",
	}
}
