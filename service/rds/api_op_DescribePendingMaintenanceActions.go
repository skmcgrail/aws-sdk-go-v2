// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rds/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of resources (for example, DB instances) that have at least one
// pending maintenance action.
func (c *Client) DescribePendingMaintenanceActions(ctx context.Context, params *DescribePendingMaintenanceActionsInput, optFns ...func(*Options)) (*DescribePendingMaintenanceActionsOutput, error) {
	stack := middleware.NewStack("DescribePendingMaintenanceActions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpDescribePendingMaintenanceActionsMiddlewares(stack)
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
	addOpDescribePendingMaintenanceActionsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePendingMaintenanceActions(options.Region), middleware.Before)
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
			OperationName: "DescribePendingMaintenanceActions",
			Err:           err,
		}
	}
	out := result.(*DescribePendingMaintenanceActionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribePendingMaintenanceActionsInput struct {
	// The ARN of a resource to return pending maintenance actions for.
	ResourceIdentifier *string
	// An optional pagination token provided by a previous
	// DescribePendingMaintenanceActions request. If this parameter is specified, the
	// response includes only records beyond the marker, up to a number of records
	// specified by MaxRecords.
	Marker *string
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that you can retrieve the remaining results.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32
	// A filter that specifies one or more resources to return pending maintenance
	// actions for. Supported filters:
	//
	//     * db-cluster-id - Accepts DB cluster
	// identifiers and DB cluster Amazon Resource Names (ARNs). The results list will
	// only include pending maintenance actions for the DB clusters identified by these
	// ARNs.
	//
	//     * db-instance-id - Accepts DB instance identifiers and DB instance
	// ARNs. The results list will only include pending maintenance actions for the DB
	// instances identified by these ARNs.
	Filters []*types.Filter
}

// Data returned from the DescribePendingMaintenanceActions action.
type DescribePendingMaintenanceActionsOutput struct {
	// An optional pagination token provided by a previous
	// DescribePendingMaintenanceActions request. If this parameter is specified, the
	// response includes only records beyond the marker, up to a number of records
	// specified by MaxRecords.
	Marker *string
	// A list of the pending maintenance actions for the resource.
	PendingMaintenanceActions []*types.ResourcePendingMaintenanceActions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpDescribePendingMaintenanceActionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpDescribePendingMaintenanceActions{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribePendingMaintenanceActions{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePendingMaintenanceActions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribePendingMaintenanceActions",
	}
}