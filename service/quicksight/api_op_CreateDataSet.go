// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a dataset.
func (c *Client) CreateDataSet(ctx context.Context, params *CreateDataSetInput, optFns ...func(*Options)) (*CreateDataSetOutput, error) {
	stack := middleware.NewStack("CreateDataSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateDataSetMiddlewares(stack)
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
	addOpCreateDataSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDataSet(options.Region), middleware.Before)
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
			OperationName: "CreateDataSet",
			Err:           err,
		}
	}
	out := result.(*CreateDataSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDataSetInput struct {
	// The row-level security configuration for the data that you want to create.
	RowLevelPermissionDataSet *types.RowLevelPermissionDataSet
	// An ID for the dataset that you want to create. This ID is unique per AWS Region
	// for each AWS account.
	DataSetId *string
	// A list of resource permissions on the dataset.
	Permissions []*types.ResourcePermission
	// Configures the combination and transformation of the data from the physical
	// tables.
	LogicalTableMap map[string]*types.LogicalTable
	// Indicates whether you want to import the data into SPICE.
	ImportMode types.DataSetImportMode
	// Groupings of columns that work together in certain QuickSight features.
	// Currently, only geospatial hierarchy is supported.
	ColumnGroups []*types.ColumnGroup
	// Contains a map of the key-value pairs for the resource tag or tags assigned to
	// the dataset.
	Tags []*types.Tag
	// Declares the physical tables that are available in the underlying data sources.
	PhysicalTableMap map[string]*types.PhysicalTable
	// The display name for the dataset.
	Name *string
	// The AWS account ID.
	AwsAccountId *string
}

type CreateDataSetOutput struct {
	// The ID of the ingestion, which is triggered as a result of dataset creation if
	// the import mode is SPICE.
	IngestionId *string
	// The Amazon Resource Name (ARN) of the dataset.
	Arn *string
	// The AWS request ID for this operation.
	RequestId *string
	// The ARN for the ingestion, which is triggered as a result of dataset creation if
	// the import mode is SPICE.
	IngestionArn *string
	// The ID for the dataset that you want to create. This ID is unique per AWS Region
	// for each AWS account.
	DataSetId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateDataSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateDataSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDataSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateDataSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateDataSet",
	}
}