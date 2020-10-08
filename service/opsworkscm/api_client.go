// Code generated by smithy-go-codegen DO NOT EDIT.

package opsworkscm

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

const ServiceID = "OpsWorksCM"
const ServiceAPIVersion = "2016-11-01"

// AWS OpsWorks CM AWS OpsWorks for configuration management (CM) is a service that
// runs and manages configuration management servers. You can use AWS OpsWorks CM
// to create and manage AWS OpsWorks for Chef Automate and AWS OpsWorks for Puppet
// Enterprise servers, and add or remove nodes for the servers to manage. Glossary
// of terms
//
//     * Server: A configuration management server that can be
// highly-available. The configuration management server runs on an Amazon Elastic
// Compute Cloud (EC2) instance, and may use various other AWS services, such as
// Amazon Relational Database Service (RDS) and Elastic Load Balancing. A server is
// a generic abstraction over the configuration manager that you want to use, much
// like Amazon RDS. In AWS OpsWorks CM, you do not start or stop servers. After you
// create servers, they continue to run until they are deleted.
//
//     * Engine: The
// engine is the specific configuration manager that you want to use. Valid values
// in this release include ChefAutomate and Puppet.
//
//     * Backup: This is an
// application-level backup of the data that the configuration manager stores. AWS
// OpsWorks CM creates an S3 bucket for backups when you launch the first server. A
// backup maintains a snapshot of a server's configuration-related attributes at
// the time the backup starts.
//
//     * Events: Events are always related to a
// server. Events are written during server creation, when health checks run, when
// backups are created, when system maintenance is performed, etc. When you delete
// a server, the server's events are also deleted.
//
//     * Account attributes: Every
// account has attributes that are assigned in the AWS OpsWorks CM database. These
// attributes store information about configuration limits (servers, backups, etc.)
// and your customer account.
//
// Endpoints AWS OpsWorks CM supports the following
// endpoints, all HTTPS. You must connect to one of the following endpoints. Your
// servers can only be accessed or managed within the endpoint in which they are
// created.
//
//     * opsworks-cm.us-east-1.amazonaws.com
//
//     *
// opsworks-cm.us-east-2.amazonaws.com
//
//     * opsworks-cm.us-west-1.amazonaws.com
//
//
// * opsworks-cm.us-west-2.amazonaws.com
//
//     *
// opsworks-cm.ap-northeast-1.amazonaws.com
//
//     *
// opsworks-cm.ap-southeast-1.amazonaws.com
//
//     *
// opsworks-cm.ap-southeast-2.amazonaws.com
//
//     *
// opsworks-cm.eu-central-1.amazonaws.com
//
//     *
// opsworks-cm.eu-west-1.amazonaws.com
//
// For more information, see AWS OpsWorks
// endpoints and quotas
// (https://docs.aws.amazon.com/general/latest/gr/opsworks-service.html) in the AWS
// General Reference. Throttling limits All API operations allow for five requests
// per second with a burst of 10 requests per second.
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
	awsmiddleware.AddUserAgentKey("opsworkscm")(stack)
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
