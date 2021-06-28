// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The specified container was not found for the specified account.
type ContainerNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ContainerNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ContainerNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ContainerNotFoundException) ErrorCode() string             { return "ContainerNotFoundException" }
func (e *ContainerNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service is temporarily unavailable.
type InternalServerError struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerError) ErrorCode() string             { return "InternalServerError" }
func (e *InternalServerError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// Could not perform an operation on an object that does not exist.
type ObjectNotFoundException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *ObjectNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ObjectNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ObjectNotFoundException) ErrorCode() string             { return "ObjectNotFoundException" }
func (e *ObjectNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The requested content range is not valid.
type RequestedRangeNotSatisfiableException struct {
	Message *string

	noSmithyDocumentSerde
}

func (e *RequestedRangeNotSatisfiableException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *RequestedRangeNotSatisfiableException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *RequestedRangeNotSatisfiableException) ErrorCode() string {
	return "RequestedRangeNotSatisfiableException"
}
func (e *RequestedRangeNotSatisfiableException) ErrorFault() smithy.ErrorFault {
	return smithy.FaultClient
}
