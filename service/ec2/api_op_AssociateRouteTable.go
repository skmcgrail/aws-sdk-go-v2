// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateRouteTableRequest
type AssociateRouteTableInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The ID of the route table.
	//
	// RouteTableId is a required field
	RouteTableId *string `locationName:"routeTableId" type:"string" required:"true"`

	// The ID of the subnet.
	//
	// SubnetId is a required field
	SubnetId *string `locationName:"subnetId" type:"string" required:"true"`
}

// String returns the string representation
func (s AssociateRouteTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AssociateRouteTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AssociateRouteTableInput"}

	if s.RouteTableId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RouteTableId"))
	}

	if s.SubnetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SubnetId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateRouteTableResult
type AssociateRouteTableOutput struct {
	_ struct{} `type:"structure"`

	// The route table association ID. This ID is required for disassociating the
	// route table.
	AssociationId *string `locationName:"associationId" type:"string"`
}

// String returns the string representation
func (s AssociateRouteTableOutput) String() string {
	return awsutil.Prettify(s)
}

const opAssociateRouteTable = "AssociateRouteTable"

// AssociateRouteTableRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Associates a subnet with a route table. The subnet and route table must be
// in the same VPC. This association causes traffic originating from the subnet
// to be routed according to the routes in the route table. The action returns
// an association ID, which you need in order to disassociate the route table
// from the subnet later. A route table can be associated with multiple subnets.
//
// For more information, see Route Tables (https://docs.aws.amazon.com/AmazonVPC/latest/UserGuide/VPC_Route_Tables.html)
// in the Amazon Virtual Private Cloud User Guide.
//
//    // Example sending a request using AssociateRouteTableRequest.
//    req := client.AssociateRouteTableRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/AssociateRouteTable
func (c *Client) AssociateRouteTableRequest(input *AssociateRouteTableInput) AssociateRouteTableRequest {
	op := &aws.Operation{
		Name:       opAssociateRouteTable,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &AssociateRouteTableInput{}
	}

	req := c.newRequest(op, input, &AssociateRouteTableOutput{})
	return AssociateRouteTableRequest{Request: req, Input: input, Copy: c.AssociateRouteTableRequest}
}

// AssociateRouteTableRequest is the request type for the
// AssociateRouteTable API operation.
type AssociateRouteTableRequest struct {
	*aws.Request
	Input *AssociateRouteTableInput
	Copy  func(*AssociateRouteTableInput) AssociateRouteTableRequest
}

// Send marshals and sends the AssociateRouteTable API request.
func (r AssociateRouteTableRequest) Send(ctx context.Context) (*AssociateRouteTableResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssociateRouteTableResponse{
		AssociateRouteTableOutput: r.Request.Data.(*AssociateRouteTableOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssociateRouteTableResponse is the response type for the
// AssociateRouteTable API operation.
type AssociateRouteTableResponse struct {
	*AssociateRouteTableOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssociateRouteTable request.
func (r *AssociateRouteTableResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}