// Code generated by smithy-go-codegen DO NOT EDIT.

package eventbridge

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/eventbridge/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Running PutPermission permits the specified AWS account or AWS organization to
// put events to the specified event bus. Amazon EventBridge (CloudWatch Events)
// rules in your account are triggered by these events arriving to an event bus in
// your account. For another account to send events to your account, that external
// account must have an EventBridge rule with your account's event bus as a target.
// <p>To enable multiple AWS accounts to put events to your event bus, run
// <code>PutPermission</code> once for each of these accounts. Or, if all the
// accounts are members of the same AWS organization, you can run
// <code>PutPermission</code> once specifying <code>Principal</code> as "*" and
// specifying the AWS organization ID in <code>Condition</code>, to grant
// permissions to all accounts in that organization.</p> <p>If you grant
// permissions using an organization, then accounts in that organization must
// specify a <code>RoleArn</code> with proper permissions when they use
// <code>PutTarget</code> to add your account's event bus as a target. For more
// information, see <a
// href="https://docs.aws.amazon.com/eventbridge/latest/userguide/eventbridge-cross-account-event-delivery.html">Sending
// and Receiving Events Between AWS Accounts</a> in the <i>Amazon EventBridge User
// Guide</i>.</p> <p>The permission policy on the default event bus cannot exceed
// 10 KB in size.</p>
func (c *Client) PutPermission(ctx context.Context, params *PutPermissionInput, optFns ...func(*Options)) (*PutPermissionOutput, error) {
	stack := middleware.NewStack("PutPermission", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpPutPermissionMiddlewares(stack)
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
	addOpPutPermissionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutPermission(options.Region), middleware.Before)
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
			OperationName: "PutPermission",
			Err:           err,
		}
	}
	out := result.(*PutPermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutPermissionInput struct {

	// The action that you are enabling the other account to perform. Currently, this
	// must be events:PutEvents.
	//
	// This member is required.
	Action *string

	// The 12-digit AWS account ID that you are permitting to put events to your
	// default event bus. Specify "*" to permit any account to put events to your
	// default event bus.  <p>If you specify "*" without specifying
	// <code>Condition</code>, avoid creating rules that may match undesirable events.
	// To create more secure rules, make sure that the event pattern for each rule
	// contains an <code>account</code> field with a specific account ID from which to
	// receive events. Rules with an account field do not match any events sent from
	// other accounts.</p>
	//
	// This member is required.
	Principal *string

	// An identifier string for the external account that you are granting permissions
	// to. If you later want to revoke the permission for this external account,
	// specify this StatementId when you run RemovePermission ().
	//
	// This member is required.
	StatementId *string

	// This parameter enables you to limit the permission to accounts that fulfill a
	// certain condition, such as being a member of a certain AWS organization. For
	// more information about AWS Organizations, see What Is AWS Organizations
	// (https://docs.aws.amazon.com/organizations/latest/userguide/orgs_introduction.html)
	// in the AWS Organizations User Guide. If you specify Condition with an AWS
	// organization ID, and specify "*" as the value for Principal, you grant
	// permission to all the accounts in the named organization.  <p>The
	// <code>Condition</code> is a JSON string which must contain <code>Type</code>,
	// <code>Key</code>, and <code>Value</code> fields.</p>
	Condition *types.Condition

	// The event bus associated with the rule. If you omit this, the default event bus
	// is used.
	EventBusName *string
}

type PutPermissionOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpPutPermissionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpPutPermission{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutPermission{}, middleware.After)
}

func newServiceMetadataMiddleware_opPutPermission(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "events",
		OperationName: "PutPermission",
	}
}
