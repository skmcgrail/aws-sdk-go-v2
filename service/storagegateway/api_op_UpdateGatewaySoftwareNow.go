// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the gateway virtual machine (VM) software. The request immediately
// triggers the software update.  <note> <p>When you make this request, you get a
// <code>200 OK</code> success response immediately. However, it might take some
// time for the update to complete. You can call <a>DescribeGatewayInformation</a>
// to verify the gateway is in the <code>STATE_RUNNING</code> state.</p> </note>
// <important> <p>A software update forces a system restart of your gateway. You
// can minimize the chance of any disruption to your applications by increasing
// your iSCSI Initiators' timeouts. For more information about increasing iSCSI
// Initiator timeouts for Windows and Linux, see <a
// href="https://docs.aws.amazon.com/storagegateway/latest/userguide/ConfiguringiSCSIClientInitiatorWindowsClient.html#CustomizeWindowsiSCSISettings">Customizing
// your Windows iSCSI settings</a> and <a
// href="https://docs.aws.amazon.com/storagegateway/latest/userguide/ConfiguringiSCSIClientInitiatorRedHatClient.html#CustomizeLinuxiSCSISettings">Customizing
// your Linux iSCSI settings</a>, respectively.</p> </important>
func (c *Client) UpdateGatewaySoftwareNow(ctx context.Context, params *UpdateGatewaySoftwareNowInput, optFns ...func(*Options)) (*UpdateGatewaySoftwareNowOutput, error) {
	stack := middleware.NewStack("UpdateGatewaySoftwareNow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateGatewaySoftwareNowMiddlewares(stack)
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
	addOpUpdateGatewaySoftwareNowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGatewaySoftwareNow(options.Region), middleware.Before)
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
			OperationName: "UpdateGatewaySoftwareNow",
			Err:           err,
		}
	}
	out := result.(*UpdateGatewaySoftwareNowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway to
// update.
type UpdateGatewaySoftwareNowInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string
}

// A JSON object containing the Amazon Resource Name (ARN) of the gateway that was
// updated.
type UpdateGatewaySoftwareNowOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateGatewaySoftwareNowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateGatewaySoftwareNow{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateGatewaySoftwareNow{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGatewaySoftwareNow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "UpdateGatewaySoftwareNow",
	}
}
