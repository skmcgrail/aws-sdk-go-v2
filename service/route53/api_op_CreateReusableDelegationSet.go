// Code generated by smithy-go-codegen DO NOT EDIT.

package route53

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a delegation set (a group of four name servers) that can be reused by
// multiple hosted zones that were created by the same AWS account. You can also
// create a reusable delegation set that uses the four name servers that are
// associated with an existing hosted zone. Specify the hosted zone ID in the
// CreateReusableDelegationSet request. You can't associate a reusable delegation
// set with a private hosted zone. For information about using a reusable
// delegation set to configure white label name servers, see Configuring White
// Label Name Servers
// (https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/white-label-name-servers.html).
// The process for migrating existing hosted zones to use a reusable delegation set
// is comparable to the process for configuring white label name servers. You need
// to perform the following steps:
//
// * Create a reusable delegation set.
//
// * Recreate
// hosted zones, and reduce the TTL to 60 seconds or less.
//
// * Recreate resource
// record sets in the new hosted zones.
//
// * Change the registrar's name servers to
// use the name servers for the new hosted zones.
//
// * Monitor traffic for the
// website or application.
//
// * Change TTLs back to their original values.
//
// If you
// want to migrate existing hosted zones to use a reusable delegation set, the
// existing hosted zones can't use any of the name servers that are assigned to the
// reusable delegation set. If one or more hosted zones do use one or more name
// servers that are assigned to the reusable delegation set, you can do one of the
// following:
//
// * For small numbers of hosted zones—up to a few hundred—it's
// relatively easy to create reusable delegation sets until you get one that has
// four name servers that don't overlap with any of the name servers in your hosted
// zones.
//
// * For larger numbers of hosted zones, the easiest solution is to use
// more than one reusable delegation set.
//
// * For larger numbers of hosted zones,
// you can also migrate hosted zones that have overlapping name servers to hosted
// zones that don't have overlapping name servers, then migrate the hosted zones
// again to use the reusable delegation set.
func (c *Client) CreateReusableDelegationSet(ctx context.Context, params *CreateReusableDelegationSetInput, optFns ...func(*Options)) (*CreateReusableDelegationSetOutput, error) {
	if params == nil {
		params = &CreateReusableDelegationSetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateReusableDelegationSet", params, optFns, c.addOperationCreateReusableDelegationSetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateReusableDelegationSetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateReusableDelegationSetInput struct {

	// A unique string that identifies the request, and that allows you to retry failed
	// CreateReusableDelegationSet requests without the risk of executing the operation
	// twice. You must use a unique CallerReference string every time you submit a
	// CreateReusableDelegationSet request. CallerReference can be any unique string,
	// for example a date/time stamp.
	//
	// This member is required.
	CallerReference *string

	// If you want to mark the delegation set for an existing hosted zone as reusable,
	// the ID for that hosted zone.
	HostedZoneId *string

	noSmithyDocumentSerde
}

type CreateReusableDelegationSetOutput struct {

	// A complex type that contains name server information.
	//
	// This member is required.
	DelegationSet *types.DelegationSet

	// The unique URL representing the new reusable delegation set.
	//
	// This member is required.
	Location *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateReusableDelegationSetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpCreateReusableDelegationSet{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpCreateReusableDelegationSet{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpCreateReusableDelegationSetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateReusableDelegationSet(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addSanitizeURLMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opCreateReusableDelegationSet(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53",
		OperationName: "CreateReusableDelegationSet",
	}
}
