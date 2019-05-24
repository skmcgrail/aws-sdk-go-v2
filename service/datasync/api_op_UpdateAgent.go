// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package datasync

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// UpdateAgentRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/UpdateAgentRequest
type UpdateAgentInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the agent to update.
	//
	// AgentArn is a required field
	AgentArn *string `type:"string" required:"true"`

	// The name that you want to use to configure the agent.
	Name *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateAgentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateAgentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateAgentInput"}

	if s.AgentArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("AgentArn"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/UpdateAgentResponse
type UpdateAgentOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateAgentOutput) String() string {
	return awsutil.Prettify(s)
}

const opUpdateAgent = "UpdateAgent"

// UpdateAgentRequest returns a request value for making API operation for
// AWS DataSync.
//
// Updates the name of an agent.
//
//    // Example sending a request using UpdateAgentRequest.
//    req := client.UpdateAgentRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/datasync-2018-11-09/UpdateAgent
func (c *Client) UpdateAgentRequest(input *UpdateAgentInput) UpdateAgentRequest {
	op := &aws.Operation{
		Name:       opUpdateAgent,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &UpdateAgentInput{}
	}

	req := c.newRequest(op, input, &UpdateAgentOutput{})
	return UpdateAgentRequest{Request: req, Input: input, Copy: c.UpdateAgentRequest}
}

// UpdateAgentRequest is the request type for the
// UpdateAgent API operation.
type UpdateAgentRequest struct {
	*aws.Request
	Input *UpdateAgentInput
	Copy  func(*UpdateAgentInput) UpdateAgentRequest
}

// Send marshals and sends the UpdateAgent API request.
func (r UpdateAgentRequest) Send(ctx context.Context) (*UpdateAgentResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAgentResponse{
		UpdateAgentOutput: r.Request.Data.(*UpdateAgentOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAgentResponse is the response type for the
// UpdateAgent API operation.
type UpdateAgentResponse struct {
	*UpdateAgentOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateAgent request.
func (r *UpdateAgentResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}