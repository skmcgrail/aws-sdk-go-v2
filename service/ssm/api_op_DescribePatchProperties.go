// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the properties of available patches organized by product, product family,
// classification, severity, and other properties of available patches. You can use
// the reported properties in the filters you specify in requests for actions such
// as CreatePatchBaseline (), UpdatePatchBaseline (), DescribeAvailablePatches (),
// and DescribePatchBaselines (). The following section lists the properties that
// can be used in filters for each major operating system type: WINDOWS Valid
// properties: PRODUCT, PRODUCT_FAMILY, CLASSIFICATION, MSRC_SEVERITY AMAZON_LINUX
// Valid properties: PRODUCT, CLASSIFICATION, SEVERITY AMAZON_LINUX_2 Valid
// properties: PRODUCT, CLASSIFICATION, SEVERITY UBUNTU Valid properties: PRODUCT,
// PRIORITY REDHAT_ENTERPRISE_LINUX Valid properties: PRODUCT, CLASSIFICATION,
// SEVERITY SUSE Valid properties: PRODUCT, CLASSIFICATION, SEVERITY CENTOS Valid
// properties: PRODUCT, CLASSIFICATION, SEVERITY
func (c *Client) DescribePatchProperties(ctx context.Context, params *DescribePatchPropertiesInput, optFns ...func(*Options)) (*DescribePatchPropertiesOutput, error) {
	stack := middleware.NewStack("DescribePatchProperties", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribePatchPropertiesMiddlewares(stack)
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
	addOpDescribePatchPropertiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePatchProperties(options.Region), middleware.Before)
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
			OperationName: "DescribePatchProperties",
			Err:           err,
		}
	}
	out := result.(*DescribePatchPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribePatchPropertiesInput struct {
	// The patch property for which you want to view patch details.
	Property types.PatchProperty
	// The operating system type for which to list patches.
	OperatingSystem types.OperatingSystem
	// Indicates whether to list patches for the Windows operating system or for
	// Microsoft applications. Not applicable for Linux operating systems.
	PatchSet types.PatchSet
	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults *int32
}

type DescribePatchPropertiesOutput struct {
	// The token for the next set of items to return. (You use this token in the next
	// call.)
	NextToken *string
	// A list of the properties for patches matching the filter request parameters.
	Properties []map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribePatchPropertiesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePatchProperties{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePatchProperties{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribePatchProperties(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "DescribePatchProperties",
	}
}