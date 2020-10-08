// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Edit or change an OpsItem. You must have permission in AWS Identity and Access
// Management (IAM) to update an OpsItem. For more information, see Getting started
// with OpsCenter
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-getting-started.html)
// in the AWS Systems Manager User Guide. Operations engineers and IT professionals
// use OpsCenter to view, investigate, and remediate operational issues impacting
// the performance and health of their AWS resources. For more information, see AWS
// Systems Manager OpsCenter
// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter.html) in
// the AWS Systems Manager User Guide.
func (c *Client) UpdateOpsItem(ctx context.Context, params *UpdateOpsItemInput, optFns ...func(*Options)) (*UpdateOpsItemOutput, error) {
	stack := middleware.NewStack("UpdateOpsItem", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateOpsItemMiddlewares(stack)
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
	addOpUpdateOpsItemValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateOpsItem(options.Region), middleware.Before)
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
			OperationName: "UpdateOpsItem",
			Err:           err,
		}
	}
	out := result.(*UpdateOpsItemOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateOpsItemInput struct {

	// The ID of the OpsItem.
	//
	// This member is required.
	OpsItemId *string

	// Specify a new category for an OpsItem.
	Category *string

	// Update the information about the OpsItem. Provide enough information so that
	// users reading this OpsItem for the first time understand the issue.
	Description *string

	// The Amazon Resource Name (ARN) of an SNS topic where notifications are sent when
	// this OpsItem is edited or changed.
	Notifications []*types.OpsItemNotification

	// Add new keys or edit existing key-value pairs of the OperationalData map in the
	// OpsItem object. Operational data is custom data that provides useful reference
	// details about the OpsItem. For example, you can specify log files, error
	// strings, license keys, troubleshooting tips, or other relevant data. You enter
	// operational data as key-value pairs. The key has a maximum length of 128
	// characters. The value has a maximum size of 20 KB.  <important> <p>Operational
	// data keys <i>can't</i> begin with the following: amazon, aws, amzn, ssm,
	// /amazon, /aws, /amzn, /ssm.</p> </important> <p>You can choose to make the data
	// searchable by other users in the account or you can restrict  search access.
	// Searchable data means that all users with access to the OpsItem Overview page
	// (as provided by the DescribeOpsItems () API action) can view and search on the
	// specified data. Operational data that is not searchable is only viewable by
	// users who have access to the OpsItem (as provided by the GetOpsItem () API
	// action).  <p>Use the <code>/aws/resources</code> key in OperationalData to
	// specify a related resource in  the request. Use the /aws/automations key in
	// OperationalData to associate an Automation runbook with the OpsItem. To view AWS
	// CLI example commands that use these keys, see Creating OpsItems manually
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-creating-OpsItems.html#OpsCenter-manually-create-OpsItems)
	// in the AWS Systems Manager User Guide.
	OperationalData map[string]*types.OpsItemDataValue

	// Keys that you want to remove from the OperationalData map.
	OperationalDataToDelete []*string

	// The importance of this OpsItem in relation to other OpsItems in the system.
	Priority *int32

	// One or more OpsItems that share something in common with the current OpsItems.
	// For example, related OpsItems can include OpsItems with similar error messages,
	// impacted resources, or statuses for the impacted resource.
	RelatedOpsItems []*types.RelatedOpsItem

	// Specify a new severity for an OpsItem.
	Severity *string

	// The OpsItem status. Status can be Open, In Progress, or Resolved. For more
	// information, see Editing OpsItem details
	// (https://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-working-with-OpsItems.html#OpsCenter-working-with-OpsItems-editing-details)
	// in the AWS Systems Manager User Guide.
	Status types.OpsItemStatus

	// A short heading that describes the nature of the OpsItem and the impacted
	// resource.
	Title *string
}

type UpdateOpsItemOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateOpsItemMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateOpsItem{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateOpsItem{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateOpsItem(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "UpdateOpsItem",
	}
}
