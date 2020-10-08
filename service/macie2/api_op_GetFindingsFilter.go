// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves information about the criteria and other settings for a findings
// filter.
func (c *Client) GetFindingsFilter(ctx context.Context, params *GetFindingsFilterInput, optFns ...func(*Options)) (*GetFindingsFilterOutput, error) {
	stack := middleware.NewStack("GetFindingsFilter", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetFindingsFilterMiddlewares(stack)
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
	addOpGetFindingsFilterValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetFindingsFilter(options.Region), middleware.Before)
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
			OperationName: "GetFindingsFilter",
			Err:           err,
		}
	}
	out := result.(*GetFindingsFilterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFindingsFilterInput struct {

	// The unique identifier for the Amazon Macie resource or account that the request
	// applies to.
	//
	// This member is required.
	Id *string
}

type GetFindingsFilterOutput struct {

	// The action that's performed on findings that meet the filter criteria
	// (findingCriteria). Possible values are: ARCHIVE, suppress (automatically
	// archive) the findings; and, NOOP, don't perform any action on the findings.
	Action types.FindingsFilterAction

	// The Amazon Resource Name (ARN) of the filter.
	Arn *string

	// The custom description of the filter.
	Description *string

	// The criteria that's used to filter findings.
	FindingCriteria *types.FindingCriteria

	// The unique identifier for the filter.
	Id *string

	// The custom name of the filter.
	Name *string

	// The position of the filter in the list of saved filters on the Amazon Macie
	// console. This value also determines the order in which the filter is applied to
	// findings, relative to other filters that are also applied to the findings.
	Position *int32

	// A map of key-value pairs that identifies the tags (keys and values) that are
	// associated with the filter.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetFindingsFilterMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetFindingsFilter{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetFindingsFilter{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetFindingsFilter(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetFindingsFilter",
	}
}
