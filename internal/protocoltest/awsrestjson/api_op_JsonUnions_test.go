// Code generated by smithy-go-codegen DO NOT EDIT.

package awsrestjson

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/awsrestjson/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/aws/smithy-go/ptr"
	smithyrand "github.com/aws/smithy-go/rand"
	smithytesting "github.com/aws/smithy-go/testing"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_JsonUnions_awsRestjson1Serialize(t *testing.T) {
	cases := map[string]struct {
		Params        *JsonUnionsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes a string union value
		"RestJsonSerializeStringUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberStringValue{Value: "foo"},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "stringValue": "foo"
			    }
			}`))
			},
		},
		// Serializes a boolean union value
		"RestJsonSerializeBooleanUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberBooleanValue{Value: true},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "booleanValue": true
			    }
			}`))
			},
		},
		// Serializes a number union value
		"RestJsonSerializeNumberUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberNumberValue{Value: 1},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "numberValue": 1
			    }
			}`))
			},
		},
		// Serializes a blob union value
		"RestJsonSerializeBlobUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberBlobValue{Value: []byte("foo")},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "blobValue": "Zm9v"
			    }
			}`))
			},
		},
		// Serializes a timestamp union value
		"RestJsonSerializeTimestampUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberTimestampValue{Value: smithytime.ParseEpochSeconds(1398796238)},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "timestampValue": 1398796238
			    }
			}`))
			},
		},
		// Serializes an enum union value
		"RestJsonSerializeEnumUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberEnumValue{Value: types.FooEnum("Foo")},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "enumValue": "Foo"
			    }
			}`))
			},
		},
		// Serializes a list union value
		"RestJsonSerializeListUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberListValue{Value: []string{
					"foo",
					"bar",
				}},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "listValue": ["foo", "bar"]
			    }
			}`))
			},
		},
		// Serializes a map union value
		"RestJsonSerializeMapUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberMapValue{Value: map[string]string{
					"foo":  "bar",
					"spam": "eggs",
				}},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "mapValue": {
			            "foo": "bar",
			            "spam": "eggs"
			        }
			    }
			}`))
			},
		},
		// Serializes a structure union value
		"RestJsonSerializeStructureUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberStructureValue{Value: types.GreetingStruct{
					Hi: ptr.String("hello"),
				}},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "structureValue": {
			            "hi": "hello"
			        }
			    }
			}`))
			},
		},
		// Serializes a renamed structure union value
		"RestJsonSerializeRenamedStructureUnionValue": {
			Params: &JsonUnionsInput{
				Contents: &types.MyUnionMemberRenamedStructureValue{Value: types.RenamedGreeting{
					Salutation: ptr.String("hello!"),
				}},
			},
			ExpectMethod:  "PUT",
			ExpectURIPath: "/JsonUnions",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareJSONReaderBytes(actual, []byte(`{
			    "contents": {
			        "renamedStructureValue": {
			            "salutation": "hello!"
			        }
			    }
			}`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			url := server.URL
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               awshttp.NewBuildableClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.JsonUnions(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_JsonUnions_awsRestjson1Deserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *JsonUnionsOutput
	}{
		// Deserializes a string union value
		"RestJsonDeserializeStringUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "stringValue": "foo"
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberStringValue{Value: "foo"},
			},
		},
		// Deserializes a boolean union value
		"RestJsonDeserializeBooleanUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "booleanValue": true
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberBooleanValue{Value: true},
			},
		},
		// Deserializes a number union value
		"RestJsonDeserializeNumberUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "numberValue": 1
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberNumberValue{Value: 1},
			},
		},
		// Deserializes a blob union value
		"RestJsonDeserializeBlobUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "blobValue": "Zm9v"
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberBlobValue{Value: []byte("foo")},
			},
		},
		// Deserializes a timestamp union value
		"RestJsonDeserializeTimestampUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "timestampValue": 1398796238
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberTimestampValue{Value: smithytime.ParseEpochSeconds(1398796238)},
			},
		},
		// Deserializes an enum union value
		"RestJsonDeserializeEnumUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "enumValue": "Foo"
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberEnumValue{Value: types.FooEnum("Foo")},
			},
		},
		// Deserializes a list union value
		"RestJsonDeserializeListUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "listValue": ["foo", "bar"]
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberListValue{Value: []string{
					"foo",
					"bar",
				}},
			},
		},
		// Deserializes a map union value
		"RestJsonDeserializeMapUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "mapValue": {
			            "foo": "bar",
			            "spam": "eggs"
			        }
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberMapValue{Value: map[string]string{
					"foo":  "bar",
					"spam": "eggs",
				}},
			},
		},
		// Deserializes a structure union value
		"RestJsonDeserializeStructureUnionValue": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"application/json"},
			},
			BodyMediaType: "application/json",
			Body: []byte(`{
			    "contents": {
			        "structureValue": {
			            "hi": "hello"
			        }
			    }
			}`),
			ExpectResult: &JsonUnionsOutput{
				Contents: &types.MyUnionMemberStructureValue{Value: types.GreetingStruct{
					Hi: ptr.String("hello"),
				}},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			url := "http://localhost:8888/"
			client := New(Options{
				HTTPClient: smithyhttp.ClientDoFunc(func(r *http.Request) (*http.Response, error) {
					headers := http.Header{}
					for k, vs := range c.Header {
						for _, v := range vs {
							headers.Add(k, v)
						}
					}
					if len(c.BodyMediaType) != 0 && len(headers.Values("Content-Type")) == 0 {
						headers.Set("Content-Type", c.BodyMediaType)
					}
					response := &http.Response{
						StatusCode: c.StatusCode,
						Header:     headers,
						Request:    r,
					}
					if len(c.Body) != 0 {
						response.ContentLength = int64(len(c.Body))
						response.Body = ioutil.NopCloser(bytes.NewReader(c.Body))
					} else {

						response.Body = http.NoBody
					}
					return response, nil
				}),
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params JsonUnionsInput
			result, err := client.JsonUnions(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			opts := cmp.Options{
				cmpopts.IgnoreUnexported(
					middleware.Metadata{},
				),
				cmp.FilterValues(func(x, y float64) bool {
					return math.IsNaN(x) && math.IsNaN(y)
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
				cmp.FilterValues(func(x, y float32) bool {
					return math.IsNaN(float64(x)) && math.IsNaN(float64(y))
				}, cmp.Comparer(func(_, _ interface{}) bool { return true })),
			}
			if err := smithytesting.CompareValues(c.ExpectResult, result, opts...); err != nil {
				t.Errorf("expect c.ExpectResult value match:\n%v", err)
			}
		})
	}
}
