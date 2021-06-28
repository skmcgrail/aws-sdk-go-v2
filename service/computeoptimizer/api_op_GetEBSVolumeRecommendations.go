// Code generated by smithy-go-codegen DO NOT EDIT.

package computeoptimizer

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/computeoptimizer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns Amazon Elastic Block Store (Amazon EBS) volume recommendations. AWS
// Compute Optimizer generates recommendations for Amazon EBS volumes that meet a
// specific set of requirements. For more information, see the Supported resources
// and requirements
// (https://docs.aws.amazon.com/compute-optimizer/latest/ug/requirements.html) in
// the AWS Compute Optimizer User Guide.
func (c *Client) GetEBSVolumeRecommendations(ctx context.Context, params *GetEBSVolumeRecommendationsInput, optFns ...func(*Options)) (*GetEBSVolumeRecommendationsOutput, error) {
	if params == nil {
		params = &GetEBSVolumeRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEBSVolumeRecommendations", params, optFns, c.addOperationGetEBSVolumeRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEBSVolumeRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEBSVolumeRecommendationsInput struct {

	// The ID of the AWS account for which to return volume recommendations. If your
	// account is the management account of an organization, use this parameter to
	// specify the member account for which you want to return volume recommendations.
	// Only one account ID can be specified per request.
	AccountIds []string

	// An array of objects that describe a filter that returns a more specific list of
	// volume recommendations.
	Filters []types.EBSFilter

	// The maximum number of volume recommendations to return with a single request. To
	// retrieve the remaining results, make another request with the returned NextToken
	// value.
	MaxResults *int32

	// The token to advance to the next page of volume recommendations.
	NextToken *string

	// The Amazon Resource Name (ARN) of the volumes for which to return
	// recommendations.
	VolumeArns []string

	noSmithyDocumentSerde
}

type GetEBSVolumeRecommendationsOutput struct {

	// An array of objects that describe errors of the request. For example, an error
	// is returned if you request recommendations for an unsupported volume.
	Errors []types.GetRecommendationError

	// The token to use to advance to the next page of volume recommendations. This
	// value is null when there are no more pages of volume recommendations to return.
	NextToken *string

	// An array of objects that describe volume recommendations.
	VolumeRecommendations []types.VolumeRecommendation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetEBSVolumeRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetEBSVolumeRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetEBSVolumeRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetEBSVolumeRecommendations(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetEBSVolumeRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "compute-optimizer",
		OperationName: "GetEBSVolumeRecommendations",
	}
}
