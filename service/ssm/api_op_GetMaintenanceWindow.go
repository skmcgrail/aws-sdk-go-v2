// Code generated by smithy-go-codegen DO NOT EDIT.

package ssm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves a maintenance window.
func (c *Client) GetMaintenanceWindow(ctx context.Context, params *GetMaintenanceWindowInput, optFns ...func(*Options)) (*GetMaintenanceWindowOutput, error) {
	stack := middleware.NewStack("GetMaintenanceWindow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetMaintenanceWindowMiddlewares(stack)
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
	addOpGetMaintenanceWindowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMaintenanceWindow(options.Region), middleware.Before)
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
			OperationName: "GetMaintenanceWindow",
			Err:           err,
		}
	}
	out := result.(*GetMaintenanceWindowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMaintenanceWindowInput struct {

	// The ID of the maintenance window for which you want to retrieve information.
	//
	// This member is required.
	WindowId *string
}

type GetMaintenanceWindowOutput struct {

	// Whether targets must be registered with the maintenance window before tasks can
	// be defined for those targets.
	AllowUnassociatedTargets *bool

	// The date the maintenance window was created.
	CreatedDate *time.Time

	// The number of hours before the end of the maintenance window that Systems
	// Manager stops scheduling new tasks for execution.
	Cutoff *int32

	// The description of the maintenance window.
	Description *string

	// The duration of the maintenance window in hours.
	Duration *int32

	// Indicates whether the maintenance window is enabled.
	Enabled *bool

	// The date and time, in ISO-8601 Extended format, for when the maintenance window
	// is scheduled to become inactive. The maintenance window will not run after this
	// specified time.
	EndDate *string

	// The date the maintenance window was last modified.
	ModifiedDate *time.Time

	// The name of the maintenance window.
	Name *string

	// The next time the maintenance window will actually run, taking into account any
	// specified times for the maintenance window to become active or inactive.
	NextExecutionTime *string

	// The schedule of the maintenance window in the form of a cron or rate expression.
	Schedule *string

	// The number of days to wait to run a maintenance window after the scheduled CRON
	// expression date and time.
	ScheduleOffset *int32

	// The time zone that the scheduled maintenance window executions are based on, in
	// Internet Assigned Numbers Authority (IANA) format. For example:
	// "America/Los_Angeles", "etc/UTC", or "Asia/Seoul". For more information, see the
	// Time Zone Database (https://www.iana.org/time-zones) on the IANA website.
	ScheduleTimezone *string

	// The date and time, in ISO-8601 Extended format, for when the maintenance window
	// is scheduled to become active. The maintenance window will not run before this
	// specified time.
	StartDate *string

	// The ID of the created maintenance window.
	WindowId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetMaintenanceWindowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetMaintenanceWindow{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetMaintenanceWindow{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetMaintenanceWindow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ssm",
		OperationName: "GetMaintenanceWindow",
	}
}
