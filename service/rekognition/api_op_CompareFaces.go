// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Compares a face in the source input image with each of the 100 largest faces
// detected in the target input image. If the source image contains multiple faces,
// the service detects the largest face and compares it with each face detected in
// the target image.  <p>You pass the input and target images either as
// base64-encoded image bytes or as references to images in an Amazon S3 bucket. If
// you use the AWS CLI to call Amazon Rekognition operations, passing image bytes
// isn't supported. The image must be formatted as a PNG or JPEG file. </p> <p>In
// response, the operation returns an array of face matches ordered by similarity
// score in descending order. For each face match, the response provides a bounding
// box of the face, facial landmarks, pose details (pitch, role, and yaw), quality
// (brightness and sharpness), and confidence value (indicating the level of
// confidence that the bounding box contains a face). The response also provides a
// similarity score, which indicates how closely the faces match. </p> <note> <p>By
// default, only faces with a similarity score of greater than or equal to 80% are
// returned in the response. You can change this value by specifying the
// <code>SimilarityThreshold</code> parameter.</p> </note> <p>
// <code>CompareFaces</code> also returns an array of faces that don't match the
// source image. For each face, it returns a bounding box, confidence value,
// landmarks, pose details, and quality. The response also returns information
// about the face in the source image, including the bounding box of the face and
// confidence value.</p> <p>The <code>QualityFilter</code> input parameter allows
// you to filter out detected faces that don’t meet a required quality bar. The
// quality bar is based on a variety of common use cases. Use
// <code>QualityFilter</code> to set the quality bar by specifying
// <code>LOW</code>, <code>MEDIUM</code>, or <code>HIGH</code>. If you do not want
// to filter detected faces, specify <code>NONE</code>. The default value is
// <code>NONE</code>. </p> <note> <p>To use quality filtering, you need a
// collection associated with version 3 of the face model or higher. To get the
// version of the face model associated with a collection, call
// <a>DescribeCollection</a>. </p> </note> <p>If the image doesn't contain Exif
// metadata, <code>CompareFaces</code> returns orientation information for the
// source and target images. Use these values to display the images with the
// correct image orientation.</p> <p>If no faces are detected in the source or
// target images, <code>CompareFaces</code> returns an
// <code>InvalidParameterException</code> error. </p> <note> <p> This is a
// stateless API operation. That is, data returned by this operation doesn't
// persist.</p> </note> <p>For an example, see Comparing Faces in Images in the
// Amazon Rekognition Developer Guide.</p> <p>This operation requires permissions
// to perform the <code>rekognition:CompareFaces</code> action.</p>
func (c *Client) CompareFaces(ctx context.Context, params *CompareFacesInput, optFns ...func(*Options)) (*CompareFacesOutput, error) {
	stack := middleware.NewStack("CompareFaces", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpCompareFacesMiddlewares(stack)
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
	addOpCompareFacesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCompareFaces(options.Region), middleware.Before)
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
			OperationName: "CompareFaces",
			Err:           err,
		}
	}
	out := result.(*CompareFacesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CompareFacesInput struct {

	// The input image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes is not
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	SourceImage *types.Image

	// The target image as base64-encoded bytes or an S3 object. If you use the AWS CLI
	// to call Amazon Rekognition operations, passing base64-encoded image bytes is not
	// supported. If you are using an AWS SDK to call Amazon Rekognition, you might not
	// need to base64-encode image bytes passed using the Bytes field. For more
	// information, see Images in the Amazon Rekognition developer guide.
	//
	// This member is required.
	TargetImage *types.Image

	// A filter that specifies a quality bar for how much filtering is done to identify
	// faces. Filtered faces aren't compared. If you specify AUTO, Amazon Rekognition
	// chooses the quality bar. If you specify LOW, MEDIUM, or HIGH, filtering removes
	// all faces that don’t meet the chosen quality bar.  The quality bar is based on a
	// variety of common use cases. Low-quality detections can occur for a number of
	// reasons. Some examples are an object that's misidentified as a face, a face
	// that's too blurry, or a face with a pose that's too extreme to use. If you
	// specify <code>NONE</code>, no filtering is performed. The default value is
	// <code>NONE</code>. </p> <p>To use quality filtering, the collection you are
	// using must be associated with version 3 of the face model or higher.</p>
	QualityFilter types.QualityFilter

	// The minimum level of confidence in the face matches that a match must meet to be
	// included in the FaceMatches array.
	SimilarityThreshold *float32
}

type CompareFacesOutput struct {

	// An array of faces in the target image that match the source image face. Each
	// CompareFacesMatch object provides the bounding box, the confidence level that
	// the bounding box contains a face, and the similarity score for the face in the
	// bounding box and the face in the source image.
	FaceMatches []*types.CompareFacesMatch

	// The face in the source image that was used for comparison.
	SourceImageFace *types.ComparedSourceImageFace

	// The value of SourceImageOrientationCorrection is always null. If the input image
	// is in .jpeg format, it might contain exchangeable image file format (Exif)
	// metadata that includes the image's orientation. Amazon Rekognition uses this
	// orientation information to perform image correction. The bounding box
	// coordinates are translated to represent object locations after the orientation
	// information in the Exif metadata is used to correct the image orientation.
	// Images in .png format don't contain Exif metadata. Amazon Rekognition doesn’t
	// perform image correction for images in .png format and .jpeg images without
	// orientation information in the image Exif metadata. The bounding box coordinates
	// aren't translated and represent the object locations before the image is
	// rotated.
	SourceImageOrientationCorrection types.OrientationCorrection

	// The value of TargetImageOrientationCorrection is always null. If the input image
	// is in .jpeg format, it might contain exchangeable image file format (Exif)
	// metadata that includes the image's orientation. Amazon Rekognition uses this
	// orientation information to perform image correction. The bounding box
	// coordinates are translated to represent object locations after the orientation
	// information in the Exif metadata is used to correct the image orientation.
	// Images in .png format don't contain Exif metadata. Amazon Rekognition doesn’t
	// perform image correction for images in .png format and .jpeg images without
	// orientation information in the image Exif metadata. The bounding box coordinates
	// aren't translated and represent the object locations before the image is
	// rotated.
	TargetImageOrientationCorrection types.OrientationCorrection

	// An array of faces in the target image that did not match the source image face.
	UnmatchedFaces []*types.ComparedFace

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpCompareFacesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpCompareFaces{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpCompareFaces{}, middleware.After)
}

func newServiceMetadataMiddleware_opCompareFaces(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "CompareFaces",
	}
}
