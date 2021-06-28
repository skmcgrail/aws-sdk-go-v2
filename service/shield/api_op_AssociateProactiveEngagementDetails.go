// Code generated by smithy-go-codegen DO NOT EDIT.

package shield

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Initializes proactive engagement and sets the list of contacts for the DDoS
// Response Team (DRT) to use. You must provide at least one phone number in the
// emergency contact list. After you have initialized proactive engagement using
// this call, to disable or enable proactive engagement, use the calls
// DisableProactiveEngagement and EnableProactiveEngagement. This call defines the
// list of email addresses and phone numbers that the DDoS Response Team (DRT) can
// use to contact you for escalations to the DRT and to initiate proactive customer
// support. The contacts that you provide in the request replace any contacts that
// were already defined. If you already have contacts defined and want to use them,
// retrieve the list using DescribeEmergencyContactSettings and then provide it to
// this call.
func (c *Client) AssociateProactiveEngagementDetails(ctx context.Context, params *AssociateProactiveEngagementDetailsInput, optFns ...func(*Options)) (*AssociateProactiveEngagementDetailsOutput, error) {
	if params == nil {
		params = &AssociateProactiveEngagementDetailsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateProactiveEngagementDetails", params, optFns, c.addOperationAssociateProactiveEngagementDetailsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateProactiveEngagementDetailsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateProactiveEngagementDetailsInput struct {

	// A list of email addresses and phone numbers that the DDoS Response Team (DRT)
	// can use to contact you for escalations to the DRT and to initiate proactive
	// customer support. To enable proactive engagement, the contact list must include
	// at least one phone number. The contacts that you provide here replace any
	// contacts that were already defined. If you already have contacts defined and
	// want to use them, retrieve the list using DescribeEmergencyContactSettings and
	// then provide it here.
	//
	// This member is required.
	EmergencyContactList []types.EmergencyContact

	noSmithyDocumentSerde
}

type AssociateProactiveEngagementDetailsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAssociateProactiveEngagementDetailsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpAssociateProactiveEngagementDetails{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpAssociateProactiveEngagementDetails{}, middleware.After)
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
	if err = addOpAssociateProactiveEngagementDetailsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateProactiveEngagementDetails(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateProactiveEngagementDetails(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "shield",
		OperationName: "AssociateProactiveEngagementDetails",
	}
}
