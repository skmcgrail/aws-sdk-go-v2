// Code generated by smithy-go-codegen DO NOT EDIT.

package dataexchange

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dataexchange/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// This operation returns information about a data set.
func (c *Client) GetDataSet(ctx context.Context, params *GetDataSetInput, optFns ...func(*Options)) (*GetDataSetOutput, error) {
	stack := middleware.NewStack("GetDataSet", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDataSetMiddlewares(stack)
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
	addOpGetDataSetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDataSet(options.Region), middleware.Before)
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
			OperationName: "GetDataSet",
			Err:           err,
		}
	}
	out := result.(*GetDataSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDataSetInput struct {

	// The unique identifier for a data set.
	//
	// This member is required.
	DataSetId *string
}

type GetDataSetOutput struct {

	// The ARN for the data set.
	Arn *string

	// The type of file your data is stored in. Currently, the supported asset type is
	// S3_SNAPSHOT.
	AssetType types.AssetType

	// The date and time that the data set was created, in ISO 8601 format.
	CreatedAt *time.Time

	// The description for the data set.
	Description *string

	// The unique identifier for the data set.
	Id *string

	// The name of the data set.
	Name *string

	// A property that defines the data set as OWNED by the account (for providers) or
	// ENTITLED to the account (for subscribers).
	Origin types.Origin

	// If the origin of this data set is ENTITLED, includes the details for the product
	// on AWS Marketplace.
	OriginDetails *types.OriginDetails

	// The data set ID of the owned data set corresponding to the entitled data set
	// being viewed. This parameter is returned when a data set owner is viewing the
	// entitled copy of its owned data set.
	SourceId *string

	// The tags for the data set.
	Tags map[string]*string

	// The date and time that the data set was last updated, in ISO 8601 format.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDataSetMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDataSet{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDataSet{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDataSet(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dataexchange",
		OperationName: "GetDataSet",
	}
}
