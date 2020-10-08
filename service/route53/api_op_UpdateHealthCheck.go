// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates an existing health check. Note that some values can't be updated. For
// more information about updating health checks, see Creating, Updating, and
// Deleting Health Checks
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/health-checks-creating-deleting.html)
// in the Amazon Route 53 Developer Guide.
func (c *Client) UpdateHealthCheck(ctx context.Context, params *UpdateHealthCheckInput, optFns ...func(*Options)) (*UpdateHealthCheckOutput, error) {
	stack := middleware.NewStack("UpdateHealthCheck", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestxml_serdeOpUpdateHealthCheckMiddlewares(stack)
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
	addOpUpdateHealthCheckValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateHealthCheck(options.Region), middleware.Before)
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
			OperationName: "UpdateHealthCheck",
			Err:           err,
		}
	}
	out := result.(*UpdateHealthCheckOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A complex type that contains information about a request to update a health
// check.
type UpdateHealthCheckInput struct {

	// The ID for the health check for which you want detailed information. When you
	// created the health check, CreateHealthCheck returned the ID in the response, in
	// the HealthCheckId element.
	//
	// This member is required.
	HealthCheckId *string

	// A complex type that identifies the CloudWatch alarm that you want Amazon Route
	// 53 health checkers to use to determine whether the specified health check is
	// healthy.
	AlarmIdentifier *types.AlarmIdentifier

	// A complex type that contains one ChildHealthCheck element for each health check
	// that you want to associate with a CALCULATED health check.
	ChildHealthChecks []*string

	// Stops Route 53 from performing health checks. When you disable a health check,
	// here's what happens:
	//
	//     * Health checks that check the health of endpoints:
	// Route 53 stops submitting requests to your application, server, or other
	// resource.
	//
	//     * Calculated health checks: Route 53 stops aggregating the status
	// of the referenced health checks.
	//
	//     * Health checks that monitor CloudWatch
	// alarms: Route 53 stops monitoring the corresponding CloudWatch metrics.
	//
	//
	// <p>After you disable a health check, Route 53 considers the status of the health
	// check to always be healthy. If you configured DNS failover, Route 53 continues
	// to route traffic to the corresponding resources. If you want to stop routing
	// traffic to a resource, change the value of <a
	// href="https://docs.aws.amazon.com/Route53/latest/APIReference/API_UpdateHealthCheck.html#Route53-UpdateHealthCheck-request-Inverted">Inverted</a>.
	// </p> <p>Charges for a health check still apply when the health check is
	// disabled. For more information, see <a
	// href="http://aws.amazon.com/route53/pricing/">Amazon Route 53 Pricing</a>.</p>
	Disabled *bool

	// Specify whether you want Amazon Route 53 to send the value of
	// FullyQualifiedDomainName to the endpoint in the client_hello message during TLS
	// negotiation. This allows the endpoint to respond to HTTPS health check requests
	// with the applicable SSL/TLS certificate. Some endpoints require that HTTPS
	// requests include the host name in the client_hello message. If you don't enable
	// SNI, the status of the health check will be SSL alert handshake_failure. A
	// health check can also have that status for other reasons. If SNI is enabled and
	// you're still getting the error, check the SSL/TLS configuration on your endpoint
	// and confirm that your certificate is valid. The SSL/TLS certificate on your
	// endpoint includes a domain name in the Common Name field and possibly several
	// more in the Subject Alternative Names field. One of the domain names in the
	// certificate should match the value that you specify for
	// FullyQualifiedDomainName. If the endpoint responds to the client_hello message
	// with a certificate that does not include the domain name that you specified in
	// FullyQualifiedDomainName, a health checker will retry the handshake. In the
	// second attempt, the health checker will omit FullyQualifiedDomainName from the
	// client_hello message.
	EnableSNI *bool

	// The number of consecutive health checks that an endpoint must pass or fail for
	// Amazon Route 53 to change the current status of the endpoint from unhealthy to
	// healthy or vice versa. For more information, see How Amazon Route 53 Determines
	// Whether an Endpoint Is Healthy
	// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/dns-failover-determining-health-of-endpoints.html)
	// in the Amazon Route 53 Developer Guide. If you don't specify a value for
	// FailureThreshold, the default value is three health checks.
	FailureThreshold *int32

	// Amazon Route 53 behavior depends on whether you specify a value for IPAddress.
	// <note> <p>If a health check already has a value for <code>IPAddress</code>, you
	// can change the value. However, you can't update an existing health check to add
	// or remove the value of <code>IPAddress</code>. </p> </note> <p> <b>If you
	// specify a value for</b> <code>IPAddress</code>:</p> <p>Route 53 sends health
	// check requests to the specified IPv4 or IPv6 address and passes the value of
	// <code>FullyQualifiedDomainName</code> in the <code>Host</code> header for all
	// health checks except TCP health checks. This is typically the fully qualified
	// DNS name of the endpoint on which you want Route 53 to perform health
	// checks.</p> <p>When Route 53 checks the health of an endpoint, here is how it
	// constructs the <code>Host</code> header:</p> <ul> <li> <p>If you specify a value
	// of <code>80</code> for <code>Port</code> and <code>HTTP</code> or
	// <code>HTTP_STR_MATCH</code> for <code>Type</code>, Route 53 passes the value of
	// <code>FullyQualifiedDomainName</code> to the endpoint in the <code>Host</code>
	// header.</p> </li> <li> <p>If you specify a value of <code>443</code> for
	// <code>Port</code> and <code>HTTPS</code> or <code>HTTPS_STR_MATCH</code> for
	// <code>Type</code>, Route 53 passes the value of
	// <code>FullyQualifiedDomainName</code> to the endpoint in the <code>Host</code>
	// header.</p> </li> <li> <p>If you specify another value for <code>Port</code> and
	// any value except <code>TCP</code> for <code>Type</code>, Route 53 passes <i>
	// <code>FullyQualifiedDomainName</code>:<code>Port</code> </i> to the endpoint in
	// the <code>Host</code> header.</p> </li> </ul> <p>If you don't specify a value
	// for <code>FullyQualifiedDomainName</code>, Route 53 substitutes the value of
	// <code>IPAddress</code> in the <code>Host</code> header in each of the above
	// cases.</p> <p> <b>If you don't specify a value for</b>
	// <code>IPAddress</code>:</p> <p>If you don't specify a value for
	// <code>IPAddress</code>, Route 53 sends a DNS request to the domain that you
	// specify in <code>FullyQualifiedDomainName</code> at the interval you specify in
	// <code>RequestInterval</code>. Using an IPv4 address that is returned by DNS,
	// Route 53 then checks the health of the endpoint.</p> <note> <p>If you don't
	// specify a value for <code>IPAddress</code>, Route 53 uses only IPv4 to send
	// health checks to the endpoint. If there's no resource record set with a type of
	// A for the name that you specify for <code>FullyQualifiedDomainName</code>, the
	// health check fails with a "DNS resolution failed" error.</p> </note> <p>If you
	// want to check the health of weighted, latency, or failover resource record sets
	// and you choose to specify the endpoint only by
	// <code>FullyQualifiedDomainName</code>, we recommend that you create a separate
	// health check for each endpoint. For example, create a health check for each HTTP
	// server that is serving content for www.example.com. For the value of
	// <code>FullyQualifiedDomainName</code>, specify the domain name of the server
	// (such as <code>us-east-2-www.example.com</code>), not the name of the resource
	// record sets (www.example.com).</p> <important> <p>In this configuration, if the
	// value of <code>FullyQualifiedDomainName</code> matches the name of the resource
	// record sets and you then associate the health check with those resource record
	// sets, health check results will be unpredictable.</p> </important> <p>In
	// addition, if the value of <code>Type</code> is <code>HTTP</code>,
	// <code>HTTPS</code>, <code>HTTP_STR_MATCH</code>, or
	// <code>HTTPS_STR_MATCH</code>, Route 53 passes the value of
	// <code>FullyQualifiedDomainName</code> in the <code>Host</code> header, as it
	// does when you specify a value for <code>IPAddress</code>. If the value of
	// <code>Type</code> is <code>TCP</code>, Route 53 doesn't pass a <code>Host</code>
	// header.</p>
	FullyQualifiedDomainName *string

	// A sequential counter that Amazon Route 53 sets to 1 when you create a health
	// check and increments by 1 each time you update settings for the health check. We
	// recommend that you use GetHealthCheck or ListHealthChecks to get the current
	// value of HealthCheckVersion for the health check that you want to update, and
	// that you include that value in your UpdateHealthCheck request. This prevents
	// Route 53 from overwriting an intervening update:
	//
	//     * If the value in the
	// UpdateHealthCheck request matches the value of HealthCheckVersion in the health
	// check, Route 53 updates the health check with the new settings.
	//
	//     * If the
	// value of HealthCheckVersion in the health check is greater, the health check was
	// changed after you got the version number. Route 53 does not update the health
	// check, and it returns a HealthCheckVersionMismatch error.
	HealthCheckVersion *int64

	// The number of child health checks that are associated with a CALCULATED health
	// that Amazon Route 53 must consider healthy for the CALCULATED health check to be
	// considered healthy. To specify the child health checks that you want to
	// associate with a CALCULATED health check, use the ChildHealthChecks and
	// ChildHealthCheck elements. Note the following:
	//
	//     * If you specify a number
	// greater than the number of child health checks, Route 53 always considers this
	// health check to be unhealthy.
	//
	//     * If you specify 0, Route 53 always considers
	// this health check to be healthy.
	HealthThreshold *int32

	// The IPv4 or IPv6 IP address for the endpoint that you want Amazon Route 53 to
	// perform health checks on. If you don't specify a value for IPAddress, Route 53
	// sends a DNS request to resolve the domain name that you specify in
	// FullyQualifiedDomainName at the interval that you specify in RequestInterval.
	// Using an IP address that is returned by DNS, Route 53 then checks the health of
	// the endpoint. Use one of the following formats for the value of IPAddress:
	//
	//
	// * IPv4 address: four values between 0 and 255, separated by periods (.), for
	// example, 192.0.2.44.
	//
	//     * IPv6 address: eight groups of four hexadecimal
	// values, separated by colons (:), for example,
	// 2001:0db8:85a3:0000:0000:abcd:0001:2345. You can also shorten IPv6 addresses as
	// described in RFC 5952, for example, 2001:db8:85a3::abcd:1:2345.
	//
	// If the endpoint
	// is an EC2 instance, we recommend that you create an Elastic IP address,
	// associate it with your EC2 instance, and specify the Elastic IP address for
	// IPAddress. This ensures that the IP address of your instance never changes. For
	// more information, see the applicable documentation:
	//
	//     * Linux: Elastic IP
	// Addresses (EIP)
	// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/elastic-ip-addresses-eip.html)
	// in the Amazon EC2 User Guide for Linux Instances
	//
	//     * Windows: Elastic IP
	// Addresses (EIP)
	// (https://docs.aws.amazon.com/AWSEC2/latest/WindowsGuide/elastic-ip-addresses-eip.html)
	// in the Amazon EC2 User Guide for Windows Instances
	//
	// If a health check already
	// has a value for IPAddress, you can change the value. However, you can't update
	// an existing health check to add or remove the value of IPAddress. For more
	// information, see FullyQualifiedDomainName
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_UpdateHealthCheck.html#Route53-UpdateHealthCheck-request-FullyQualifiedDomainName).
	// Constraints: Route 53 can't check the health of endpoints for which the IP
	// address is in local, private, non-routable, or multicast ranges. For more
	// information about IP addresses for which you can't create health checks, see the
	// following documents:
	//
	//     * RFC 5735, Special Use IPv4 Addresses
	// (https://tools.ietf.org/html/rfc5735)
	//
	//     * RFC 6598, IANA-Reserved IPv4 Prefix
	// for Shared Address Space (https://tools.ietf.org/html/rfc6598)
	//
	//     * RFC 5156,
	// Special-Use IPv6 Addresses (https://tools.ietf.org/html/rfc5156)
	IPAddress *string

	// When CloudWatch has insufficient data about the metric to determine the alarm
	// state, the status that you want Amazon Route 53 to assign to the health check:
	//
	//
	// * Healthy: Route 53 considers the health check to be healthy.
	//
	//     * Unhealthy:
	// Route 53 considers the health check to be unhealthy.
	//
	//     * LastKnownStatus:
	// Route 53 uses the status of the health check from the last time CloudWatch had
	// sufficient data to determine the alarm state. For new health checks that have no
	// last known status, the default status for the health check is healthy.
	InsufficientDataHealthStatus types.InsufficientDataHealthStatus

	// Specify whether you want Amazon Route 53 to invert the status of a health check,
	// for example, to consider a health check unhealthy when it otherwise would be
	// considered healthy.
	Inverted *bool

	// The port on the endpoint that you want Amazon Route 53 to perform health checks
	// on. Don't specify a value for Port when you specify a value for Type of
	// CLOUDWATCH_METRIC or CALCULATED.
	Port *int32

	// A complex type that contains one Region element for each region that you want
	// Amazon Route 53 health checkers to check the specified endpoint from.
	Regions []types.HealthCheckRegion

	// A complex type that contains one ResettableElementName element for each element
	// that you want to reset to the default value. Valid values for
	// ResettableElementName include the following:
	//
	//     * ChildHealthChecks: Amazon
	// Route 53 resets ChildHealthChecks
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheckConfig.html#Route53-Type-HealthCheckConfig-ChildHealthChecks)
	// to null.
	//
	//     * FullyQualifiedDomainName: Route 53 resets
	// FullyQualifiedDomainName
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_UpdateHealthCheck.html#Route53-UpdateHealthCheck-request-FullyQualifiedDomainName).
	// to null.
	//
	//     * Regions: Route 53 resets the Regions
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheckConfig.html#Route53-Type-HealthCheckConfig-Regions)
	// list to the default set of regions.
	//
	//     * ResourcePath: Route 53 resets
	// ResourcePath
	// (https://docs.aws.amazon.com/Route53/latest/APIReference/API_HealthCheckConfig.html#Route53-Type-HealthCheckConfig-ResourcePath)
	// to null.
	ResetElements []types.ResettableElementName

	// The path that you want Amazon Route 53 to request when performing health checks.
	// The path can be any value for which your endpoint will return an HTTP status
	// code of 2xx or 3xx when the endpoint is healthy, for example the file
	// /docs/route53-health-check.html. You can also include query string parameters,
	// for example, /welcome.html?language=jp&login=y. Specify this value only if you
	// want to change it.
	ResourcePath *string

	// If the value of Type is HTTP_STR_MATCH or HTTPS_STR_MATCH, the string that you
	// want Amazon Route 53 to search for in the response body from the specified
	// resource. If the string appears in the response body, Route 53 considers the
	// resource healthy. (You can't change the value of Type when you update a health
	// check.)
	SearchString *string
}

// A complex type that contains the response to the UpdateHealthCheck request.
type UpdateHealthCheckOutput struct {

	// A complex type that contains the response to an UpdateHealthCheck request.
	//
	// This member is required.
	HealthCheck *types.HealthCheck

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestxml_serdeOpUpdateHealthCheckMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestxml_serializeOpUpdateHealthCheck{}, middleware.After)
	stack.Deserialize.Add(&awsRestxml_deserializeOpUpdateHealthCheck{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateHealthCheck(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "UpdateHealthCheck",
	}
}
