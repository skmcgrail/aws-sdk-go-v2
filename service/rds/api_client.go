// Code generated by smithy-go-codegen DO NOT EDIT.

package rds

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

const ServiceID = "RDS"
const ServiceAPIVersion = "2014-10-31"

// Amazon Relational Database Service Amazon Relational Database Service (Amazon
// RDS) is a web service that makes it easier to set up, operate, and scale a
// relational database in the cloud. It provides cost-efficient, resizeable
// capacity for an industry-standard relational database and manages common
// database administration tasks, freeing up developers to focus on what makes
// their applications and businesses unique. Amazon RDS gives you access to the
// capabilities of a MySQL, MariaDB, PostgreSQL, Microsoft SQL Server, Oracle, or
// Amazon Aurora database server. These capabilities mean that the code,
// applications, and tools you already use today with your existing databases work
// with Amazon RDS without modification. Amazon RDS automatically backs up your
// database and maintains the database software that powers your DB instance.
// Amazon RDS is flexible: you can scale your DB instance's compute resources and
// storage capacity to meet your application's demand. As with all Amazon Web
// Services, there are no up-front investments, and you pay only for the resources
// you use. This interface reference for Amazon RDS contains documentation for a
// programming or command line interface you can use to manage Amazon RDS. Amazon
// RDS is asynchronous, which means that some interfaces might require techniques
// such as polling or callback functions to determine when a command has been
// applied. In this reference, the parameter descriptions indicate whether a
// command is applied immediately, on the next instance reboot, or during the
// maintenance window. The reference structure is as follows, and we list following
// some related topics from the user guide.  <p> <b>Amazon RDS API Reference</b>
// </p> <ul> <li> <p>For the alphabetical list of API actions, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Operations.html">API
// Actions</a>.</p> </li> <li> <p>For the alphabetical list of data types, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_Types.html">Data
// Types</a>.</p> </li> <li> <p>For a list of common query parameters, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/CommonParameters.html">Common
// Parameters</a>.</p> </li> <li> <p>For descriptions of the error codes, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/CommonErrors.html">Common
// Errors</a>.</p> </li> </ul> <p> <b>Amazon RDS User Guide</b> </p> <ul> <li>
// <p>For a summary of the Amazon RDS interfaces, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Welcome.html#Welcome.Interfaces">Available
// RDS Interfaces</a>.</p> </li> <li> <p>For more information about how to use the
// Query API, see <a
// href="https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/Using_the_Query_API.html">Using
// the Query API</a>.</p> </li> </ul>
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
	awsmiddleware.AddUserAgentKey("rds")(stack)
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
