// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists the iSCSI stored volumes of a gateway. Results are sorted by volume ARN.
// The response includes only the volume ARNs. If you want additional volume
// information, use the DescribeStorediSCSIVolumes () or the
// DescribeCachediSCSIVolumes () API.  <p>The operation supports pagination. By
// default, the operation returns a maximum of up to 100 volumes. You can
// optionally specify the <code>Limit</code> field in the body to limit the number
// of volumes in the response. If the number of volumes returned in the response is
// truncated, the response includes a Marker field. You can use this Marker value
// in your subsequent request to retrieve the next set of volumes. This operation
// is only supported in the cached volume and stored volume gateway types.</p>
func (c *Client) ListVolumes(ctx context.Context, params *ListVolumesInput, optFns ...func(*Options)) (*ListVolumesOutput, error) {
	stack := middleware.NewStack("ListVolumes", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListVolumesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListVolumes(options.Region), middleware.Before)
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
			OperationName: "ListVolumes",
			Err:           err,
		}
	}
	out := result.(*ListVolumesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object that contains one or more of the following fields:  <ul> <li> <p>
// <a>ListVolumesInput$Limit</a> </p> </li> <li> <p> <a>ListVolumesInput$Marker</a>
// </p> </li> </ul>
type ListVolumesInput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
	// Specifies that the list of volumes returned be limited to the specified number
	// of items.
	Limit *int32
	// A string that indicates the position at which to begin the returned list of
	// volumes. Obtain the marker from the response of a previous List iSCSI Volumes
	// request.
	Marker *string
}

// A JSON object containing the following fields:  <ul> <li> <p>
// <a>ListVolumesOutput$Marker</a> </p> </li> <li> <p>
// <a>ListVolumesOutput$VolumeInfos</a> </p> </li> </ul>
type ListVolumesOutput struct {
	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways () operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string
	// An array of VolumeInfo () objects, where each object describes an iSCSI volume.
	// If no volumes are defined for the gateway, then VolumeInfos is an empty array
	// "[]".
	VolumeInfos []*types.VolumeInfo
	// Use the marker in your next request to continue pagination of iSCSI volumes. If
	// there are no more volumes to list, this field does not appear in the response
	// body.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListVolumesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListVolumes{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListVolumes{}, middleware.After)
}

func newServiceMetadataMiddleware_opListVolumes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ListVolumes",
	}
}