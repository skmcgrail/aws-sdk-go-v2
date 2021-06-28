// Code generated by smithy-go-codegen DO NOT EDIT.

package organizations

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes the specified member AWS account as a delegated administrator for the
// specified AWS service. Deregistering a delegated administrator can have
// unintended impacts on the functionality of the enabled AWS service. See the
// documentation for the enabled service before you deregister a delegated
// administrator so that you understand any potential impacts. You can run this
// action only for AWS services that support this feature. For a current list of
// services that support it, see the column Supports Delegated Administrator in the
// table at AWS Services that you can use with AWS Organizations
// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_integrate_services_list.html)
// in the AWS Organizations User Guide. This operation can be called only from the
// organization's management account.
func (c *Client) DeregisterDelegatedAdministrator(ctx context.Context, params *DeregisterDelegatedAdministratorInput, optFns ...func(*Options)) (*DeregisterDelegatedAdministratorOutput, error) {
	if params == nil {
		params = &DeregisterDelegatedAdministratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterDelegatedAdministrator", params, optFns, c.addOperationDeregisterDelegatedAdministratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterDelegatedAdministratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterDelegatedAdministratorInput struct {

	// The account ID number of the member account in the organization that you want to
	// deregister as a delegated administrator.
	//
	// This member is required.
	AccountId *string

	// The service principal name of an AWS service for which the account is a
	// delegated administrator. Delegated administrator privileges are revoked for only
	// the specified AWS service from the member account. If the specified service is
	// the only service for which the member account is a delegated administrator, the
	// operation also revokes Organizations read action permissions.
	//
	// This member is required.
	ServicePrincipal *string

	noSmithyDocumentSerde
}

type DeregisterDelegatedAdministratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeregisterDelegatedAdministratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterDelegatedAdministrator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterDelegatedAdministrator{}, middleware.After)
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
	if err = addOpDeregisterDelegatedAdministratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterDelegatedAdministrator(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDeregisterDelegatedAdministrator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "organizations",
		OperationName: "DeregisterDelegatedAdministrator",
	}
}
