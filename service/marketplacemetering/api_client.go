// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacemetering

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/awslabs/smithy-go/middleware"
	"net/http"
	"time"
)

const ServiceID = "Marketplace Metering"
const ServiceAPIVersion = "2016-01-14"

// AWS Marketplace Metering Service This reference provides descriptions of the
// low-level AWS Marketplace Metering Service API. AWS Marketplace sellers can use
// this API to submit usage data for custom usage dimensions. For information on
// the permissions you need to use this API, see AWS Marketing metering and
// entitlement API permissions
// (https://docs.aws.amazon.com/marketplace/latest/userguide/iam-user-policy-for-aws-marketplace-actions.html)
// in the AWS Marketplace Seller Guide. Submitting Metering Records
//
//     *
// MeterUsage- Submits the metering record for a Marketplace product. MeterUsage is
// called from an EC2 instance or a container running on EKS or ECS.
//
//     *
// BatchMeterUsage- Submits the metering record for a set of customers.
// BatchMeterUsage is called from a software-as-a-service (SaaS)
// application.
//
// Accepting New Customers
//
//     * ResolveCustomer- Called by a SaaS
// application during the registration process. When a buyer visits your website
// during the registration process, the buyer submits a Registration Token through
// the browser. The Registration Token is resolved through this API to obtain a
// CustomerIdentifier and Product Code.
//
// Entitlement and Metering for Paid
// Container Products
//
//     * Paid container software products sold through AWS
// Marketplace must integrate with the AWS Marketplace Metering Service and call
// the RegisterUsage operation for software entitlement and metering. Free and BYOL
// products for Amazon ECS or Amazon EKS aren't required to call RegisterUsage, but
// you can do so if you want to receive usage data in your seller reports. For more
// information on using the RegisterUsage operation, see Container-Based Products
// (https://docs.aws.amazon.com/marketplace/latest/userguide/container-based-products.html).
//
// BatchMeterUsage
// API calls are captured by AWS CloudTrail. You can use Cloudtrail to verify that
// the SaaS metering records that you sent are accurate by searching for records
// with the eventName of BatchMeterUsage. You can also use CloudTrail to audit
// records over time. For more information, see the AWS CloudTrail User Guide
// (http://docs.aws.amazon.com/awscloudtrail/latest/userguide/cloudtrail-concepts.html)
// .
type Client struct {
	options Options
}

// New returns an initialized Client based on the functional options. Provide
// additional functional options to further configure the behavior of the client,
// such as changing the client's endpoint or adding custom middleware behavior.
func New(options Options, optFns ...func(*Options)) *Client {
	options = options.Copy()

	resolveRetryer(&options)

	resolveHTTPClient(&options)

	resolveHTTPSignerV4(&options)

	resolveDefaultEndpointConfiguration(&options)

	for _, fn := range optFns {
		fn(&options)
	}

	client := &Client{
		options: options,
	}

	return client
}

type Options struct {
	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// The credentials object to use when signing requests.
	Credentials aws.CredentialsProvider

	// The endpoint options to be used when attempting to resolve an endpoint.
	EndpointOptions ResolverOptions

	// The service endpoint resolver.
	EndpointResolver EndpointResolver

	// Signature Version 4 (SigV4) Signer
	HTTPSignerV4 HTTPSignerV4

	// The region to send requests to. (Required)
	Region string

	// Retryer guides how HTTP requests should be retried in case of recoverable
	// failures. When nil the API client will use a default retryer.
	Retryer retry.Retryer

	// The HTTP client to invoke API calls with. Defaults to client's default HTTP
	// implementation if nil.
	HTTPClient HTTPClient
}

type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// Copy creates a clone where the APIOptions list is deep copied.
func (o Options) Copy() Options {
	to := o
	to.APIOptions = make([]func(*middleware.Stack) error, len(o.APIOptions))
	copy(to.APIOptions, o.APIOptions)
	return to
}

// NewFromConfig returns a new client from the provided config.
func NewFromConfig(cfg aws.Config, optFns ...func(*Options)) *Client {
	opts := Options{
		Region:      cfg.Region,
		Retryer:     cfg.Retryer,
		HTTPClient:  cfg.HTTPClient,
		Credentials: cfg.Credentials,
		APIOptions:  cfg.APIOptions,
	}
	resolveAWSEndpointResolver(cfg, &opts)
	return New(opts, optFns...)
}

func resolveHTTPClient(o *Options) {
	if o.HTTPClient != nil {
		return
	}
	o.HTTPClient = aws.NewBuildableHTTPClient()
}

func resolveRetryer(o *Options) {
	if o.Retryer != nil {
		return
	}
	o.Retryer = retry.NewStandard()
}

func resolveAWSEndpointResolver(cfg aws.Config, o *Options) {
	if cfg.EndpointResolver == nil {
		return
	}
	o.EndpointResolver = WithEndpointResolver(cfg.EndpointResolver, NewDefaultEndpointResolver())
}

func addClientUserAgent(stack *middleware.Stack) {
	awsmiddleware.AddUserAgentKey("marketplacemetering")(stack)
}

func addHTTPSignerV4Middleware(stack *middleware.Stack, o Options) {
	stack.Finalize.Add(v4.NewSignHTTPRequestMiddleware(o.Credentials, o.HTTPSignerV4), middleware.After)
}

type HTTPSignerV4 interface {
	SignHTTP(ctx context.Context, credentials aws.Credentials, r *http.Request, payloadHash string, service string, region string, signingTime time.Time) error
}

func resolveHTTPSignerV4(o *Options) {
	if o.HTTPSignerV4 != nil {
		return
	}
	o.HTTPSignerV4 = v4.NewSigner()
}

func addRetryMiddlewares(stack *middleware.Stack, o Options) error {
	mo := retry.AddRetryMiddlewaresOptions{
		Retryer: o.Retryer,
	}
	return retry.AddRetryMiddlewares(stack, mo)
}

func addRequestIDRetrieverMiddleware(stack *middleware.Stack) {
	awsmiddleware.AddRequestIDRetrieverMiddleware(stack)
}

func addResponseErrorMiddleware(stack *middleware.Stack) {
	awshttp.AddResponseErrorMiddleware(stack)
}
