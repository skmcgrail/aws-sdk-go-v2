// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of all valid regions for Amazon Lightsail. Use the include
// availability zones parameter to also return the Availability Zones in a region.
func (c *Client) GetRegions(ctx context.Context, params *GetRegionsInput, optFns ...func(*Options)) (*GetRegionsOutput, error) {
	if params == nil {
		params = &GetRegionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetRegions", params, optFns, c.addOperationGetRegionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetRegionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetRegionsInput struct {

	// A Boolean value indicating whether to also include Availability Zones in your
	// get regions request. Availability Zones are indicated with a letter: e.g.,
	// us-east-2a.
	IncludeAvailabilityZones *bool

	// A Boolean value indicating whether to also include Availability Zones for
	// databases in your get regions request. Availability Zones are indicated with a
	// letter (e.g., us-east-2a).
	IncludeRelationalDatabaseAvailabilityZones *bool

	noSmithyDocumentSerde
}

type GetRegionsOutput struct {

	// An array of key-value pairs containing information about your get regions
	// request.
	Regions []types.Region

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetRegionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetRegions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetRegions{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetRegions(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetRegions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetRegions",
	}
}
