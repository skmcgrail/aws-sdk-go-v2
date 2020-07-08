// Code generated by smithy-go-codegen DO NOT EDIT.
package jsonrpc

import (
	"bytes"
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/jsonrpc/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/httpbinding"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithytime "github.com/awslabs/smithy-go/time"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

type awsAwsjson11_serializeOpEmptyOperation struct {
}

func (*awsAwsjson11_serializeOpEmptyOperation) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpEmptyOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*EmptyOperationInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonProtocol.EmptyOperation")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentEmptyOperationInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpGreetingWithErrors struct {
}

func (*awsAwsjson11_serializeOpGreetingWithErrors) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpGreetingWithErrors) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*GreetingWithErrorsInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonProtocol.GreetingWithErrors")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentGreetingWithErrorsInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpJsonEnums struct {
}

func (*awsAwsjson11_serializeOpJsonEnums) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpJsonEnums) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*JsonEnumsInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonProtocol.JsonEnums")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentJsonEnumsInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpKitchenSinkOperation struct {
}

func (*awsAwsjson11_serializeOpKitchenSinkOperation) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpKitchenSinkOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*KitchenSinkOperationInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonProtocol.KitchenSinkOperation")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentKitchenSinkOperationInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpNullOperation struct {
}

func (*awsAwsjson11_serializeOpNullOperation) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpNullOperation) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*NullOperationInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonProtocol.NullOperation")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentNullOperationInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpOperationWithOptionalInputOutput struct {
}

func (*awsAwsjson11_serializeOpOperationWithOptionalInputOutput) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpOperationWithOptionalInputOutput) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*OperationWithOptionalInputOutputInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonProtocol.OperationWithOptionalInputOutput")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentOperationWithOptionalInputOutputInput(input, jsonEncoder.Value); err != nil {
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

type awsAwsjson11_serializeOpPutAndGetInlineDocuments struct {
}

func (*awsAwsjson11_serializeOpPutAndGetInlineDocuments) ID() string {
	return "OperationSerializer"
}

func (m *awsAwsjson11_serializeOpPutAndGetInlineDocuments) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	request, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, &smithy.SerializationError{Err: fmt.Errorf("unknown transport type %T", in.Request)}
	}

	input, ok := in.Parameters.(*PutAndGetInlineDocumentsInput)
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
	httpBindingEncoder.SetHeader("X-Amz-Target").String("JsonProtocol.PutAndGetInlineDocuments")

	jsonEncoder := smithyjson.NewEncoder()
	if err := awsAwsjson11_serializeDocumentPutAndGetInlineDocumentsInput(input, jsonEncoder.Value); err != nil {
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
func awsAwsjson11_serializeDocumentDocument(v smithy.Document, value smithyjson.Value) error {
	// TODO: implement document serialization.
	return nil
}

func awsAwsjson11_serializeDocumentEmptyStruct(v *types.EmptyStruct, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	return nil
}

func awsAwsjson11_serializeDocumentKitchenSink(v *types.KitchenSink, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Blob != nil {
		ok := object.Key("Blob")
		ok.Base64EncodeBytes(v.Blob)
	}

	if v.Boolean != nil {
		ok := object.Key("Boolean")
		ok.Boolean(*v.Boolean)
	}

	if v.Double != nil {
		ok := object.Key("Double")
		ok.Double(*v.Double)
	}

	if v.EmptyStruct != nil {
		ok := object.Key("EmptyStruct")
		if err := awsAwsjson11_serializeDocumentEmptyStruct(v.EmptyStruct, ok); err != nil {
			return err
		}
	}

	if v.Float != nil {
		ok := object.Key("Float")
		ok.Float(*v.Float)
	}

	if v.HttpdateTimestamp != nil {
		ok := object.Key("HttpdateTimestamp")
		ok.String(smithytime.FormatHTTPDate(*v.HttpdateTimestamp))
	}

	if v.Integer != nil {
		ok := object.Key("Integer")
		ok.Integer(*v.Integer)
	}

	if v.Iso8601Timestamp != nil {
		ok := object.Key("Iso8601Timestamp")
		ok.String(smithytime.FormatDateTime(*v.Iso8601Timestamp))
	}

	if v.JsonValue != nil {
		ok := object.Key("JsonValue")
		ok.String(*v.JsonValue)
	}

	if v.ListOfLists != nil {
		ok := object.Key("ListOfLists")
		if err := awsAwsjson11_serializeDocumentListOfListOfStrings(v.ListOfLists, ok); err != nil {
			return err
		}
	}

	if v.ListOfMapsOfStrings != nil {
		ok := object.Key("ListOfMapsOfStrings")
		if err := awsAwsjson11_serializeDocumentListOfMapsOfStrings(v.ListOfMapsOfStrings, ok); err != nil {
			return err
		}
	}

	if v.ListOfStrings != nil {
		ok := object.Key("ListOfStrings")
		if err := awsAwsjson11_serializeDocumentListOfStrings(v.ListOfStrings, ok); err != nil {
			return err
		}
	}

	if v.ListOfStructs != nil {
		ok := object.Key("ListOfStructs")
		if err := awsAwsjson11_serializeDocumentListOfStructs(v.ListOfStructs, ok); err != nil {
			return err
		}
	}

	if v.Long != nil {
		ok := object.Key("Long")
		ok.Long(*v.Long)
	}

	if v.MapOfListsOfStrings != nil {
		ok := object.Key("MapOfListsOfStrings")
		if err := awsAwsjson11_serializeDocumentMapOfListsOfStrings(v.MapOfListsOfStrings, ok); err != nil {
			return err
		}
	}

	if v.MapOfMaps != nil {
		ok := object.Key("MapOfMaps")
		if err := awsAwsjson11_serializeDocumentMapOfMapOfStrings(v.MapOfMaps, ok); err != nil {
			return err
		}
	}

	if v.MapOfStrings != nil {
		ok := object.Key("MapOfStrings")
		if err := awsAwsjson11_serializeDocumentMapOfStrings(v.MapOfStrings, ok); err != nil {
			return err
		}
	}

	if v.MapOfStructs != nil {
		ok := object.Key("MapOfStructs")
		if err := awsAwsjson11_serializeDocumentMapOfStructs(v.MapOfStructs, ok); err != nil {
			return err
		}
	}

	if v.RecursiveList != nil {
		ok := object.Key("RecursiveList")
		if err := awsAwsjson11_serializeDocumentListOfKitchenSinks(v.RecursiveList, ok); err != nil {
			return err
		}
	}

	if v.RecursiveMap != nil {
		ok := object.Key("RecursiveMap")
		if err := awsAwsjson11_serializeDocumentMapOfKitchenSinks(v.RecursiveMap, ok); err != nil {
			return err
		}
	}

	if v.RecursiveStruct != nil {
		ok := object.Key("RecursiveStruct")
		if err := awsAwsjson11_serializeDocumentKitchenSink(v.RecursiveStruct, ok); err != nil {
			return err
		}
	}

	if v.SimpleStruct != nil {
		ok := object.Key("SimpleStruct")
		if err := awsAwsjson11_serializeDocumentSimpleStruct(v.SimpleStruct, ok); err != nil {
			return err
		}
	}

	if v.String_ != nil {
		ok := object.Key("String")
		ok.String(*v.String_)
	}

	if v.StructWithLocationName != nil {
		ok := object.Key("StructWithLocationName")
		if err := awsAwsjson11_serializeDocumentStructWithLocationName(v.StructWithLocationName, ok); err != nil {
			return err
		}
	}

	if v.Timestamp != nil {
		ok := object.Key("Timestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.Timestamp))
	}

	if v.UnixTimestamp != nil {
		ok := object.Key("UnixTimestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.UnixTimestamp))
	}

	return nil
}

func awsAwsjson11_serializeDocumentListOfKitchenSinks(v []*types.KitchenSink, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		if err := awsAwsjson11_serializeDocumentKitchenSink(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentListOfListOfStrings(v [][]*string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		if err := awsAwsjson11_serializeDocumentListOfStrings(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentListOfMapsOfStrings(v []map[string]*string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		if err := awsAwsjson11_serializeDocumentMapOfStrings(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentListOfStrings(v []*string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		av.String(*v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentListOfStructs(v []*types.SimpleStruct, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		if err := awsAwsjson11_serializeDocumentSimpleStruct(v[i], av); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentMapOfKitchenSinks(v map[string]*types.KitchenSink, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			om.Null()
			continue
		}
		if err := awsAwsjson11_serializeDocumentKitchenSink(v[key], om); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentMapOfListsOfStrings(v map[string][]*string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			om.Null()
			continue
		}
		if err := awsAwsjson11_serializeDocumentListOfStrings(v[key], om); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentMapOfMapOfStrings(v map[string]map[string]*string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			om.Null()
			continue
		}
		if err := awsAwsjson11_serializeDocumentMapOfStrings(v[key], om); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentMapOfStrings(v map[string]*string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			om.Null()
			continue
		}
		om.String(*v[key])
	}
	return nil
}

func awsAwsjson11_serializeDocumentMapOfStructs(v map[string]*types.SimpleStruct, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			om.Null()
			continue
		}
		if err := awsAwsjson11_serializeDocumentSimpleStruct(v[key], om); err != nil {
			return err
		}
	}
	return nil
}

func awsAwsjson11_serializeDocumentSimpleStruct(v *types.SimpleStruct, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson11_serializeDocumentStructWithLocationName(v *types.StructWithLocationName, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Value != nil {
		ok := object.Key("RenamedMember")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson11_serializeDocumentFooEnumList(v []types.FooEnum, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsAwsjson11_serializeDocumentFooEnumMap(v map[string]types.FooEnum, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		om.String(string(v[key]))
	}
	return nil
}

func awsAwsjson11_serializeDocumentFooEnumSet(v []types.FooEnum, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		av.String(string(v[i]))
	}
	return nil
}

func awsAwsjson11_serializeDocumentStringList(v []*string, value smithyjson.Value) error {
	array := value.Array()
	defer array.Close()

	for i := range v {
		av := array.Value()
		if vv := v[i]; vv == nil {
			av.Null()
			continue
		}
		av.String(*v[i])
	}
	return nil
}

func awsAwsjson11_serializeDocumentStringMap(v map[string]*string, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	for key := range v {
		om := object.Key(key)
		if vv := v[key]; vv == nil {
			om.Null()
			continue
		}
		om.String(*v[key])
	}
	return nil
}

func awsAwsjson11_serializeDocumentEmptyOperationInput(v *EmptyOperationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	return nil
}

func awsAwsjson11_serializeDocumentGreetingWithErrorsInput(v *GreetingWithErrorsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	return nil
}

func awsAwsjson11_serializeDocumentJsonEnumsInput(v *JsonEnumsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if len(v.FooEnum1) > 0 {
		ok := object.Key("fooEnum1")
		ok.String(string(v.FooEnum1))
	}

	if len(v.FooEnum2) > 0 {
		ok := object.Key("fooEnum2")
		ok.String(string(v.FooEnum2))
	}

	if len(v.FooEnum3) > 0 {
		ok := object.Key("fooEnum3")
		ok.String(string(v.FooEnum3))
	}

	if v.FooEnumList != nil {
		ok := object.Key("fooEnumList")
		if err := awsAwsjson11_serializeDocumentFooEnumList(v.FooEnumList, ok); err != nil {
			return err
		}
	}

	if v.FooEnumMap != nil {
		ok := object.Key("fooEnumMap")
		if err := awsAwsjson11_serializeDocumentFooEnumMap(v.FooEnumMap, ok); err != nil {
			return err
		}
	}

	if v.FooEnumSet != nil {
		ok := object.Key("fooEnumSet")
		if err := awsAwsjson11_serializeDocumentFooEnumSet(v.FooEnumSet, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentKitchenSinkOperationInput(v *KitchenSinkOperationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Blob != nil {
		ok := object.Key("Blob")
		ok.Base64EncodeBytes(v.Blob)
	}

	if v.Boolean != nil {
		ok := object.Key("Boolean")
		ok.Boolean(*v.Boolean)
	}

	if v.Double != nil {
		ok := object.Key("Double")
		ok.Double(*v.Double)
	}

	if v.EmptyStruct != nil {
		ok := object.Key("EmptyStruct")
		if err := awsAwsjson11_serializeDocumentEmptyStruct(v.EmptyStruct, ok); err != nil {
			return err
		}
	}

	if v.Float != nil {
		ok := object.Key("Float")
		ok.Float(*v.Float)
	}

	if v.HttpdateTimestamp != nil {
		ok := object.Key("HttpdateTimestamp")
		ok.String(smithytime.FormatHTTPDate(*v.HttpdateTimestamp))
	}

	if v.Integer != nil {
		ok := object.Key("Integer")
		ok.Integer(*v.Integer)
	}

	if v.Iso8601Timestamp != nil {
		ok := object.Key("Iso8601Timestamp")
		ok.String(smithytime.FormatDateTime(*v.Iso8601Timestamp))
	}

	if v.JsonValue != nil {
		ok := object.Key("JsonValue")
		ok.String(*v.JsonValue)
	}

	if v.ListOfLists != nil {
		ok := object.Key("ListOfLists")
		if err := awsAwsjson11_serializeDocumentListOfListOfStrings(v.ListOfLists, ok); err != nil {
			return err
		}
	}

	if v.ListOfMapsOfStrings != nil {
		ok := object.Key("ListOfMapsOfStrings")
		if err := awsAwsjson11_serializeDocumentListOfMapsOfStrings(v.ListOfMapsOfStrings, ok); err != nil {
			return err
		}
	}

	if v.ListOfStrings != nil {
		ok := object.Key("ListOfStrings")
		if err := awsAwsjson11_serializeDocumentListOfStrings(v.ListOfStrings, ok); err != nil {
			return err
		}
	}

	if v.ListOfStructs != nil {
		ok := object.Key("ListOfStructs")
		if err := awsAwsjson11_serializeDocumentListOfStructs(v.ListOfStructs, ok); err != nil {
			return err
		}
	}

	if v.Long != nil {
		ok := object.Key("Long")
		ok.Long(*v.Long)
	}

	if v.MapOfListsOfStrings != nil {
		ok := object.Key("MapOfListsOfStrings")
		if err := awsAwsjson11_serializeDocumentMapOfListsOfStrings(v.MapOfListsOfStrings, ok); err != nil {
			return err
		}
	}

	if v.MapOfMaps != nil {
		ok := object.Key("MapOfMaps")
		if err := awsAwsjson11_serializeDocumentMapOfMapOfStrings(v.MapOfMaps, ok); err != nil {
			return err
		}
	}

	if v.MapOfStrings != nil {
		ok := object.Key("MapOfStrings")
		if err := awsAwsjson11_serializeDocumentMapOfStrings(v.MapOfStrings, ok); err != nil {
			return err
		}
	}

	if v.MapOfStructs != nil {
		ok := object.Key("MapOfStructs")
		if err := awsAwsjson11_serializeDocumentMapOfStructs(v.MapOfStructs, ok); err != nil {
			return err
		}
	}

	if v.RecursiveList != nil {
		ok := object.Key("RecursiveList")
		if err := awsAwsjson11_serializeDocumentListOfKitchenSinks(v.RecursiveList, ok); err != nil {
			return err
		}
	}

	if v.RecursiveMap != nil {
		ok := object.Key("RecursiveMap")
		if err := awsAwsjson11_serializeDocumentMapOfKitchenSinks(v.RecursiveMap, ok); err != nil {
			return err
		}
	}

	if v.RecursiveStruct != nil {
		ok := object.Key("RecursiveStruct")
		if err := awsAwsjson11_serializeDocumentKitchenSink(v.RecursiveStruct, ok); err != nil {
			return err
		}
	}

	if v.SimpleStruct != nil {
		ok := object.Key("SimpleStruct")
		if err := awsAwsjson11_serializeDocumentSimpleStruct(v.SimpleStruct, ok); err != nil {
			return err
		}
	}

	if v.String_ != nil {
		ok := object.Key("String")
		ok.String(*v.String_)
	}

	if v.StructWithLocationName != nil {
		ok := object.Key("StructWithLocationName")
		if err := awsAwsjson11_serializeDocumentStructWithLocationName(v.StructWithLocationName, ok); err != nil {
			return err
		}
	}

	if v.Timestamp != nil {
		ok := object.Key("Timestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.Timestamp))
	}

	if v.UnixTimestamp != nil {
		ok := object.Key("UnixTimestamp")
		ok.Double(smithytime.FormatEpochSeconds(*v.UnixTimestamp))
	}

	return nil
}

func awsAwsjson11_serializeDocumentNullOperationInput(v *NullOperationInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.String_ != nil {
		ok := object.Key("string")
		ok.String(*v.String_)
	}

	if v.StringList != nil {
		ok := object.Key("stringList")
		if err := awsAwsjson11_serializeDocumentStringList(v.StringList, ok); err != nil {
			return err
		}
	}

	if v.StringMap != nil {
		ok := object.Key("stringMap")
		if err := awsAwsjson11_serializeDocumentStringMap(v.StringMap, ok); err != nil {
			return err
		}
	}

	return nil
}

func awsAwsjson11_serializeDocumentOperationWithOptionalInputOutputInput(v *OperationWithOptionalInputOutputInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.Value != nil {
		ok := object.Key("Value")
		ok.String(*v.Value)
	}

	return nil
}

func awsAwsjson11_serializeDocumentPutAndGetInlineDocumentsInput(v *PutAndGetInlineDocumentsInput, value smithyjson.Value) error {
	object := value.Object()
	defer object.Close()

	if v.InlineDocument != nil {
		ok := object.Key("inlineDocument")
		if err := awsAwsjson11_serializeDocumentDocument(v.InlineDocument, ok); err != nil {
			return err
		}
	}

	return nil
}
