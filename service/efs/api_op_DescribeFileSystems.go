// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the description of a specific Amazon EFS file system if either the file
// system CreationToken or the FileSystemId is provided. Otherwise, it returns
// descriptions of all file systems owned by the caller's AWS account in the AWS
// Region of the endpoint that you're calling.  <p>When retrieving all file system
// descriptions, you can optionally specify the <code>MaxItems</code> parameter to
// limit the number of descriptions in a response. Currently, this number is
// automatically set to 10. If more file system descriptions remain, Amazon EFS
// returns a <code>NextMarker</code>, an opaque token, in the response. In this
// case, you should send a subsequent request with the <code>Marker</code> request
// parameter set to the value of <code>NextMarker</code>. </p> <p>To retrieve a
// list of your file system descriptions, this operation is used in an iterative
// process, where <code>DescribeFileSystems</code> is called first without the
// <code>Marker</code> and then the operation continues to call it with the
// <code>Marker</code> parameter set to the value of the <code>NextMarker</code>
// from the previous response until the response has no <code>NextMarker</code>.
// </p> <p> The order of file systems returned in the response of one
// <code>DescribeFileSystems</code> call and the order of file systems returned
// across the responses of a multi-call iteration is unspecified. </p> <p> This
// operation requires permissions for the
// <code>elasticfilesystem:DescribeFileSystems</code> action. </p>
func (c *Client) DescribeFileSystems(ctx context.Context, params *DescribeFileSystemsInput, optFns ...func(*Options)) (*DescribeFileSystemsOutput, error) {
	stack := middleware.NewStack("DescribeFileSystems", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDescribeFileSystemsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeFileSystems(options.Region), middleware.Before)
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
			OperationName: "DescribeFileSystems",
			Err:           err,
		}
	}
	out := result.(*DescribeFileSystemsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeFileSystemsInput struct {

	// (Optional) Restricts the list to the file system with this creation token
	// (String). You specify a creation token when you create an Amazon EFS file
	// system.
	CreationToken *string

	// (Optional) ID of the file system whose description you want to retrieve
	// (String).
	FileSystemId *string

	// (Optional) Opaque pagination token returned from a previous DescribeFileSystems
	// operation (String). If present, specifies to continue the list from where the
	// returning call had left off.
	Marker *string

	// (Optional) Specifies the maximum number of file systems to return in the
	// response (integer). This number is automatically set to 100. The response is
	// paginated at 100 per page if you have more than 100 file systems.
	MaxItems *int32
}

type DescribeFileSystemsOutput struct {

	// An array of file system descriptions.
	FileSystems []*types.FileSystemDescription

	// Present if provided by caller in the request (String).
	Marker *string

	// Present if there are more file systems than returned in the response (String).
	// You can use the NextMarker in the subsequent request to fetch the descriptions.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDescribeFileSystemsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDescribeFileSystems{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeFileSystems{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeFileSystems(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DescribeFileSystems",
	}
}
