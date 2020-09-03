// Code generated by smithy-go-codegen DO NOT EDIT.

package marketplacemetering

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/marketplacemetering/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsAwsjson11_serializeOpBatchMeterUsage struct {
}

func (*awsAwsjson11_serializeOpBatchMeterUsage) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpBatchMeterUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*BatchMeterUsageInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSMPMeteringService.BatchMeterUsage")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentBatchMeterUsageInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpMeterUsage struct {
}

func (*awsAwsjson11_serializeOpMeterUsage) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpMeterUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*MeterUsageInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSMPMeteringService.MeterUsage")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentMeterUsageInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpRegisterUsage struct {
}

func (*awsAwsjson11_serializeOpRegisterUsage) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpRegisterUsage) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*RegisterUsageInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSMPMeteringService.RegisterUsage")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentRegisterUsageInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}

type awsAwsjson11_serializeOpResolveCustomer struct {
}

func (*awsAwsjson11_serializeOpResolveCustomer) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpResolveCustomer) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*ResolveCustomerInput)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown input parameters type %T", in.Parameters)}
	}

	request.Request.URL.Path = "/"
	request.Request.Method = "POST"
	httpBindingEncoder, err := httpbinding.NewEncoder(request.URL.Path, request.URL.RawQuery, request.Header)
	if err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	httpBindingEncoder.SetHeader("Content-Type").String("application/x-amz-json-1.1")
	httpBindingEncoder.SetHeader("X-Amz-Target").String("AWSMPMeteringService.ResolveCustomer")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentResolveCustomerInput(input, jsonEncoder.Value); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request, err = request.SetStream(bytes.NewReader(jsonEncoder.Bytes())); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}

	if request.Request, err = httpBindingEncoder.Encode(request.Request); err != nil {
		return out, metadata, &smithy.SerializationError{Err: err}
	}
	in.Request = request

	return next.HandleSerialize(ctx, in)
}
func awsAwsjson11_serializeDocumentUsageRecord(v *types.UsageRecord, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.CustomerIdentifier != nil {
		ok := object.Key("CustomerIdentifier")
		ok.String(*v.CustomerIdentifier)
	}

	if v.Dimension != nil {
		ok := object.Key("Dimension")
		ok.String(*v.Dimension)
	}

	if v.Quantity != nil {
		ok := object.Key("Quantity")
		ok.Integer(*v.Quantity)
	}

	if v.Timestamp != nil {
		ok := object.Key("Timestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.Timestamp))
	}

	return nil
}

func awsAwsjson11_serializeDocumentUsageRecordList(v []*types.UsageRecord, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		if err := awsAwsjson11_serializeDocumentUsageRecord(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentBatchMeterUsageInput(v *BatchMeterUsageInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.ProductCode != nil {
		ok := object.Key("ProductCode")
		ok.String(*v.ProductCode)
	}

	if v.UsageRecords != nil {
		ok := object.Key("UsageRecords")
		if err := awsAwsjson11_serializeDocumentUsageRecordList(v.UsageRecords, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentMeterUsageInput(v *MeterUsageInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.DryRun != nil {
		ok := object.Key("DryRun")
		ok.Boolean(*v.DryRun)
	}

	if v.ProductCode != nil {
		ok := object.Key("ProductCode")
		ok.String(*v.ProductCode)
	}

	if v.Timestamp != nil {
		ok := object.Key("Timestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.Timestamp))
	}

	if v.UsageDimension != nil {
		ok := object.Key("UsageDimension")
		ok.String(*v.UsageDimension)
	}

	if v.UsageQuantity != nil {
		ok := object.Key("UsageQuantity")
		ok.Integer(*v.UsageQuantity)
	}

	return nil
}

func awsAwsjson11_serializeDocumentRegisterUsageInput(v *RegisterUsageInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Nonce != nil {
		ok := object.Key("Nonce")
		ok.String(*v.Nonce)
	}

	if v.ProductCode != nil {
		ok := object.Key("ProductCode")
		ok.String(*v.ProductCode)
	}

	if v.PublicKeyVersion != nil {
		ok := object.Key("PublicKeyVersion")
		ok.Integer(*v.PublicKeyVersion)
	}

	return nil
}

func awsAwsjson11_serializeDocumentResolveCustomerInput(v *ResolveCustomerInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.RegistrationToken != nil {
		ok := object.Key("RegistrationToken")
		ok.String(*v.RegistrationToken)
	}

	return nil
}