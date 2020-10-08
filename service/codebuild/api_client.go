// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

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

const ServiceID = "CodeBuild"
const ServiceAPIVersion = "2016-10-06"

// AWS CodeBuild AWS CodeBuild is a fully managed build service in the cloud. AWS
// CodeBuild compiles your source code, runs unit tests, and produces artifacts
// that are ready to deploy. AWS CodeBuild eliminates the need to provision,
// manage, and scale your own build servers. It provides prepackaged build
// environments for the most popular programming languages and build tools, such as
// Apache Maven, Gradle, and more. You can also fully customize build environments
// in AWS CodeBuild to use your own build tools. AWS CodeBuild scales automatically
// to meet peak build requests. You pay only for the build time you consume. For
// more information about AWS CodeBuild, see the AWS CodeBuild User Guide
// (https://docs.aws.amazon.com/codebuild/latest/userguide/welcome.html). AWS
// CodeBuild supports these operations:
//
//     * BatchDeleteBuilds: Deletes one or
// more builds.
//
//     * BatchGetBuilds: Gets information about one or more builds.
//
//
// * BatchGetProjects: Gets information about one or more build projects. A build
// project defines how AWS CodeBuild runs a build. This includes information such
// as where to get the source code to build, the build environment to use, the
// build commands to run, and where to store the build output. A build environment
// is a representation of operating system, programming language runtime, and tools
// that AWS CodeBuild uses to run a build. You can add tags to build projects to
// help manage your resources and costs.
//
//     * BatchGetReportGroups: Returns an
// array of report groups.
//
//     * BatchGetReports: Returns an array of reports.
//
//
// * CreateProject: Creates a build project.
//
//     * CreateReportGroup: Creates a
// report group. A report group contains a collection of reports.
//
//     *
// CreateWebhook: For an existing AWS CodeBuild build project that has its source
// code stored in a GitHub or Bitbucket repository, enables AWS CodeBuild to start
// rebuilding the source code every time a code change is pushed to the
// repository.
//
//     * DeleteProject: Deletes a build project.
//
//     * DeleteReport:
// Deletes a report.
//
//     * DeleteReportGroup: Deletes a report group.
//
//     *
// DeleteResourcePolicy: Deletes a resource policy that is identified by its
// resource ARN.
//
//     * DeleteSourceCredentials: Deletes a set of GitHub, GitHub
// Enterprise, or Bitbucket source credentials.
//
//     * DeleteWebhook: For an
// existing AWS CodeBuild build project that has its source code stored in a GitHub
// or Bitbucket repository, stops AWS CodeBuild from rebuilding the source code
// every time a code change is pushed to the repository.
//
//     * DescribeTestCases:
// Returns a list of details about test cases for a report.
//
//     *
// GetResourcePolicy: Gets a resource policy that is identified by its resource
// ARN.
//
//     * ImportSourceCredentials: Imports the source repository credentials
// for an AWS CodeBuild project that has its source code stored in a GitHub, GitHub
// Enterprise, or Bitbucket repository.
//
//     * InvalidateProjectCache: Resets the
// cache for a project.
//
//     * ListBuilds: Gets a list of build IDs, with each
// build ID representing a single build.
//
//     * ListBuildsForProject: Gets a list
// of build IDs for the specified build project, with each build ID representing a
// single build.
//
//     * ListCuratedEnvironmentImages: Gets information about Docker
// images that are managed by AWS CodeBuild.
//
//     * ListProjects: Gets a list of
// build project names, with each build project name representing a single build
// project.
//
//     * ListReportGroups: Gets a list ARNs for the report groups in the
// current AWS account.
//
//     * ListReports: Gets a list ARNs for the reports in the
// current AWS account.
//
//     * ListReportsForReportGroup: Returns a list of ARNs
// for the reports that belong to a ReportGroup.
//
//     * ListSharedProjects: Gets a
// list of ARNs associated with projects shared with the current AWS account or
// user.
//
//     * ListSharedReportGroups: Gets a list of ARNs associated with report
// groups shared with the current AWS account or user
//
//     * ListSourceCredentials:
// Returns a list of SourceCredentialsInfo objects. Each SourceCredentialsInfo
// object includes the authentication type, token ARN, and type of source provider
// for one set of credentials.
//
//     * PutResourcePolicy: Stores a resource policy
// for the ARN of a Project or ReportGroup object.
//
//     * StartBuild: Starts
// running a build.
//
//     * StopBuild: Attempts to stop running a build.
//
//     *
// UpdateProject: Changes the settings of an existing build project.
//
//     *
// UpdateReportGroup: Changes a report group.
//
//     * UpdateWebhook: Changes the
// settings of an existing webhook.
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
	awsmiddleware.AddUserAgentKey("codebuild")(stack)
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
