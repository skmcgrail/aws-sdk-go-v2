// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Provides a list of the configuration sets associated with your Amazon SES
// account in the current AWS Region. For information about using configuration
// sets, see Monitoring Your Amazon SES Sending Activity
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html)
// in the Amazon SES Developer Guide. You can execute this operation no more than
// once per second. This operation will return up to 1,000 configuration sets each
// time it is run. If your Amazon SES account has more than 1,000 configuration
// sets, this operation will also return a NextToken element. You can then execute
// the ListConfigurationSets operation again, passing the NextToken parameter and
// the value of the NextToken element to retrieve additional results.
func (c *Client) ListConfigurationSets(ctx context.Context, params *ListConfigurationSetsInput, optFns ...func(*Options)) (*ListConfigurationSetsOutput, error) {
	stack := middleware.NewStack("ListConfigurationSets", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpListConfigurationSetsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListConfigurationSets(options.Region), middleware.Before)
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
			OperationName: "ListConfigurationSets",
			Err:           err,
		}
	}
	out := result.(*ListConfigurationSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to list the configuration sets associated with your AWS
// account. Configuration sets enable you to publish email sending events. For
// information about using configuration sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type ListConfigurationSetsInput struct {

	// The number of configuration sets to return.
	MaxItems *int32

	// A token returned from a previous call to ListConfigurationSets to indicate the
	// position of the configuration set in the configuration set list.
	NextToken *string
}

// A list of configuration sets associated with your AWS account. Configuration
// sets enable you to publish email sending events. For information about using
// configuration sets, see the Amazon SES Developer Guide
// (https://docs.aws.amazon.com/ses/latest/DeveloperGuide/monitor-sending-activity.html).
type ListConfigurationSetsOutput struct {

	// A list of configuration sets.
	ConfigurationSets []*types.ConfigurationSet

	// A token indicating that there are additional configuration sets available to be
	// listed. Pass this token to successive calls of ListConfigurationSets.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpListConfigurationSetsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpListConfigurationSets{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpListConfigurationSets{}, middleware.After)
}

func newServiceMetadataMiddleware_opListConfigurationSets(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListConfigurationSets",
	}
}
