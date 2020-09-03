// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Searches for summary information about workflows.
func (c *Client) SearchFlowTemplates(ctx context.Context, params *SearchFlowTemplatesInput, optFns ...func(*Options)) (*SearchFlowTemplatesOutput, error) {
	stack := middleware.NewStack("SearchFlowTemplates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpSearchFlowTemplatesMiddlewares(stack)
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
	addOpSearchFlowTemplatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opSearchFlowTemplates(options.Region), middleware.Before)

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
			OperationName: "SearchFlowTemplates",
			Err:           err,
		}
	}
	out := result.(*SearchFlowTemplatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchFlowTemplatesInput struct {
	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string
	// The maximum number of results to return in the response.
	MaxResults *int32
	// An array of objects that limit the result set. The only valid filter is
	// DEVICE_MODEL_ID.
	Filters []*types.FlowTemplateFilter
}

type SearchFlowTemplatesOutput struct {
	// An array of objects that contain summary information about each workflow in the
	// result set.
	Summaries []*types.FlowTemplateSummary
	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpSearchFlowTemplatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpSearchFlowTemplates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpSearchFlowTemplates{}, middleware.After)
}

func newServiceMetadataMiddleware_opSearchFlowTemplates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "SearchFlowTemplates",
	}
}