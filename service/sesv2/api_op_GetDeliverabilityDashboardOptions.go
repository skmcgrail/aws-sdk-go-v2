// Code generated by smithy-go-codegen DO NOT EDIT.

package sesv2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sesv2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieve information about the status of the Deliverability dashboard for your
// account. When the Deliverability dashboard is enabled, you gain access to
// reputation, deliverability, and other metrics for the domains that you use to
// send email. You also gain the ability to perform predictive inbox placement
// tests.  <p>When you use the Deliverability dashboard, you pay a monthly
// subscription charge, in addition to any other fees that you accrue by using
// Amazon SES and other AWS services. For more information about the features and
// cost of a Deliverability dashboard subscription, see <a
// href="http://aws.amazon.com/ses/pricing/">Amazon SES Pricing</a>.</p>
func (c *Client) GetDeliverabilityDashboardOptions(ctx context.Context, params *GetDeliverabilityDashboardOptionsInput, optFns ...func(*Options)) (*GetDeliverabilityDashboardOptionsOutput, error) {
	stack := middleware.NewStack("GetDeliverabilityDashboardOptions", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetDeliverabilityDashboardOptionsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeliverabilityDashboardOptions(options.Region), middleware.Before)
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
			OperationName: "GetDeliverabilityDashboardOptions",
			Err:           err,
		}
	}
	out := result.(*GetDeliverabilityDashboardOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Retrieve information about the status of the Deliverability dashboard for your
// AWS account. When the Deliverability dashboard is enabled, you gain access to
// reputation, deliverability, and other metrics for your domains. You also gain
// the ability to perform predictive inbox placement tests.  <p>When you use the
// Deliverability dashboard, you pay a monthly subscription charge, in addition to
// any other fees that you accrue by using Amazon SES and other AWS services. For
// more information about the features and cost of a Deliverability dashboard
// subscription, see <a href="http://aws.amazon.com/pinpoint/pricing/">Amazon
// Pinpoint Pricing</a>.</p>
type GetDeliverabilityDashboardOptionsInput struct {
}

// An object that shows the status of the Deliverability dashboard.
type GetDeliverabilityDashboardOptionsOutput struct {

	// Specifies whether the Deliverability dashboard is enabled. If this value is
	// true, the dashboard is enabled.
	//
	// This member is required.
	DashboardEnabled *bool

	// The current status of your Deliverability dashboard subscription. If this value
	// is PENDING_EXPIRATION, your subscription is scheduled to expire at the end of
	// the current calendar month.
	AccountStatus types.DeliverabilityDashboardAccountStatus

	// An array of objects, one for each verified domain that you use to send email and
	// currently has an active Deliverability dashboard subscription that isn’t
	// scheduled to expire at the end of the current calendar month.
	ActiveSubscribedDomains []*types.DomainDeliverabilityTrackingOption

	// An array of objects, one for each verified domain that you use to send email and
	// currently has an active Deliverability dashboard subscription that's scheduled
	// to expire at the end of the current calendar month.
	PendingExpirationSubscribedDomains []*types.DomainDeliverabilityTrackingOption

	// The date, in Unix time format, when your current subscription to the
	// Deliverability dashboard is scheduled to expire, if your subscription is
	// scheduled to expire at the end of the current calendar month. This value is null
	// if you have an active subscription that isn’t due to expire at the end of the
	// month.
	SubscriptionExpiryDate *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetDeliverabilityDashboardOptionsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetDeliverabilityDashboardOptions{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetDeliverabilityDashboardOptions{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetDeliverabilityDashboardOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "GetDeliverabilityDashboardOptions",
	}
}
