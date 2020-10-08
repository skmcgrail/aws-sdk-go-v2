// Code generated by smithy-go-codegen DO NOT EDIT.

package costexplorer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates recommendations that help you save cost by identifying idle and
// underutilized Amazon EC2 instances. Recommendations are generated to either
// downsize or terminate instances, along with providing savings detail and
// metrics. For details on calculation and function, see Optimizing Your Cost with
// Rightsizing Recommendations
// (https://docs.aws.amazon.com/awsaccountbilling/latest/aboutv2/ce-rightsizing.html)
// in the AWS Billing and Cost Management User Guide.
func (c *Client) GetRightsizingRecommendation(ctx context.Context, params *GetRightsizingRecommendationInput, optFns ...func(*Options)) (*GetRightsizingRecommendationOutput, error) {
	stack := middleware.NewStack("GetRightsizingRecommendation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetRightsizingRecommendationMiddlewares(stack)
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
	addOpGetRightsizingRecommendationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetRightsizingRecommendation(options.Region), middleware.Before)
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
			OperationName: "GetRightsizingRecommendation",
			Err:           err,
		}
	}
	out := result.(*GetRightsizingRecommendationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRightsizingRecommendationInput struct {

	// The specific service that you want recommendations for. The only valid value for
	// GetRightsizingRecommendation is "AmazonEC2".
	//
	// This member is required.
	Service *string

	// Enables you to customize recommendations across two attributes. You can choose
	// to view recommendations for instances within the same instance families or
	// across different instance families. You can also choose to view your estimated
	// savings associated with recommendations with consideration of existing Savings
	// Plans or RI benefits, or neither.
	Configuration *types.RightsizingRecommendationConfiguration

	// Use Expression to filter by cost or by usage. There are two patterns:
	//
	//     *
	// Simple dimension values - You can set the dimension name and values for the
	// filters that you plan to use. For example, you can filter for REGION==us-east-1
	// OR REGION==us-west-1. The Expression for that looks like this: { "Dimensions": {
	// "Key": "REGION", "Values": [ "us-east-1", “us-west-1” ] } } The list of
	// dimension values are OR'd together to retrieve cost or usage data. You can
	// create Expression and DimensionValues objects using either with* methods or set*
	// methods in multiple lines.
	//
	//     * Compound dimension values with logical
	// operations - You can use multiple Expression types and the logical operators
	// AND/OR/NOT to create a list of one or more Expression objects. This allows you
	// to filter on more advanced options. For example, you can filter on ((REGION ==
	// us-east-1 OR REGION == us-west-1) OR (TAG.Type == Type1)) AND (USAGE_TYPE !=
	// DataTransfer). The Expression for that looks like this: { "And": [ {"Or": [
	// {"Dimensions": { "Key": "REGION", "Values": [ "us-east-1", "us-west-1" ] }},
	// {"Tags": { "Key": "TagName", "Values": ["Value1"] } } ]}, {"Not": {"Dimensions":
	// { "Key": "USAGE_TYPE", "Values": ["DataTransfer"] }}} ] }  Because each
	// Expression can have only one operator, the service returns an error if more than
	// one is specified. The following example shows an Expression object that creates
	// an error.  { "And": [ ... ], "DimensionValues": { "Dimension": "USAGE_TYPE",
	// "Values": [ "DataTransfer" ] } }
	//
	// For GetRightsizingRecommendation action, a
	// combination of OR and NOT is not supported. OR is not supported between
	// different dimensions, or dimensions and tags. NOT operators aren't supported.
	// Dimensions are also limited to LINKED_ACCOUNT, REGION, or RIGHTSIZING_TYPE.
	Filter *types.Expression

	// The pagination token that indicates the next set of results that you want to
	// retrieve.
	NextPageToken *string

	// The number of recommendations that you want returned in a single response
	// object.
	PageSize *int32
}

type GetRightsizingRecommendationOutput struct {

	// Enables you to customize recommendations across two attributes. You can choose
	// to view recommendations for instances within the same instance families or
	// across different instance families. You can also choose to view your estimated
	// savings associated with recommendations with consideration of existing Savings
	// Plans or RI benefits, or neither.
	Configuration *types.RightsizingRecommendationConfiguration

	// Information regarding this specific recommendation set.
	Metadata *types.RightsizingRecommendationMetadata

	// The token to retrieve the next set of results.
	NextPageToken *string

	// Recommendations to rightsize resources.
	RightsizingRecommendations []*types.RightsizingRecommendation

	// Summary of this recommendation set.
	Summary *types.RightsizingRecommendationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetRightsizingRecommendationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetRightsizingRecommendation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRightsizingRecommendation{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetRightsizingRecommendation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ce",
		OperationName: "GetRightsizingRecommendation",
	}
}
