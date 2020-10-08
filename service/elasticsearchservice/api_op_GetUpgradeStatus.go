// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the latest status of the last upgrade or upgrade eligibility check
// that was performed on the domain.
func (c *Client) GetUpgradeStatus(ctx context.Context, params *GetUpgradeStatusInput, optFns ...func(*Options)) (*GetUpgradeStatusOutput, error) {
	stack := middleware.NewStack("GetUpgradeStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetUpgradeStatusMiddlewares(stack)
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
	addOpGetUpgradeStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUpgradeStatus(options.Region), middleware.Before)
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
			OperationName: "GetUpgradeStatus",
			Err:           err,
		}
	}
	out := result.(*GetUpgradeStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for request parameters to GetUpgradeStatus () operation.
type GetUpgradeStatusInput struct {

	// The name of an Elasticsearch domain. Domain names are unique across the domains
	// owned by an account within an AWS region. Domain names start with a letter or
	// number and can contain the following characters: a-z (lowercase), 0-9, and -
	// (hyphen).
	//
	// This member is required.
	DomainName *string
}

// Container for response returned by GetUpgradeStatus () operation.
type GetUpgradeStatusOutput struct {

	// One of 4 statuses that a step can go through returned as part of the
	// GetUpgradeStatusResponse () object. The status can take one of the following
	// values:
	//
	//     * In Progress
	//
	//     * Succeeded
	//
	//     * Succeeded with Issues
	//
	//     *
	// Failed
	StepStatus types.UpgradeStatus

	// A string that describes the update briefly
	UpgradeName *string

	// Represents one of 3 steps that an Upgrade or Upgrade Eligibility Check does
	// through:
	//
	//     * PreUpgradeCheck
	//
	//     * Snapshot
	//
	//     * Upgrade
	UpgradeStep types.UpgradeStep

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetUpgradeStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetUpgradeStatus{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUpgradeStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUpgradeStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "GetUpgradeStatus",
	}
}
