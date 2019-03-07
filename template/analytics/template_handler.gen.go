// Copyright 2017 Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// THIS FILE IS AUTOMATICALLY GENERATED.

package analytics

import (
	"context"
	"net"
	"time"

	"istio.io/istio/mixer/pkg/adapter"
)

// The `analytics` template represents a single request reported to Apigee's analytics processing system.
// Complete Apigee documentation on the concepts and usage of this adapter is also available on the
// [Apigee Adapter for Istio](https://docs.apigee.com/api-platform/istio-adapter/concepts) site.
// For more information and product support, please [contact Apigee support](https://apigee.com/about/support/portal).
//
// Example config:
//
// ```yaml
// apiVersion: config.istio.io/v1alpha2
// kind: analytics
// metadata:
//   name: apigee
//   namespace: istio-system
// spec:
//   api_key: request.api_key | request.headers["x-api-key"] | ""
//   api_proxy: api.service | destination.service | ""
//   response_status_code: response.code | 0
//   client_ip: source.ip | ip("0.0.0.0")
//   request_verb: request.method | ""
//   request_uri: request.path | ""
//   request_path: request.path | ""
//   useragent: request.useragent | ""
//   client_received_start_timestamp: request.time
//   client_received_end_timestamp: request.time
//   target_sent_start_timestamp: request.time
//   target_sent_end_timestamp: request.time
//   target_received_start_timestamp: response.time
//   target_received_end_timestamp: response.time
//   client_sent_start_timestamp: response.time
//   client_sent_end_timestamp: response.time
//   api_claims: # from jwt
//     json_claims: request.auth.raw_claims | ""
// ```

// Fully qualified name of the template
const TemplateName = "analytics"

// Instance is constructed by Mixer for the 'analytics' template.
//
// This Template provides Istio telemetry data to the Apigee Analytics engine.
// For additional information on this adapter or support please contact anchor-prega-support@google.com.
type Instance struct {
	// Name of the instance as specified in configuration.
	Name string

	// The name of the proxy (usually the Istio API or service name).
	ApiProxy string

	// HTTP response code
	ResponseStatusCode int64

	// Client IP address
	ClientIp net.IP

	// HTTP request verb
	RequestVerb string

	// HTTP request URI
	RequestUri string

	// HTTP request path
	RequestPath string

	// HTTP user agent header
	Useragent string

	// Timestamp of when the api_proxy started receiving the request.
	ClientReceivedStartTimestamp time.Time

	// Timestamp of when the api_proxy  finished receiving the request.
	ClientReceivedEndTimestamp time.Time

	// Timestamp of when the api_proxy started sending the request to the target.
	ClientSentStartTimestamp time.Time

	// Timestamp of when the api_proxy finished sending the request to the target.
	ClientSentEndTimestamp time.Time

	// Timestamp of when the api_proxy started request to target.
	TargetSentStartTimestamp time.Time

	// Timestamp of when the api_proxy finished sending request to target.
	TargetSentEndTimestamp time.Time

	// Timestamp of when the api_proxy started receiving response from target.
	TargetReceivedStartTimestamp time.Time

	// Timestamp of when the api_proxy finished receiving response from target.
	TargetReceivedEndTimestamp time.Time

	// The JWT claims that were used for authenticating the request (if any)
	// Use subkey "json_claims" for passing all claims in as a single JSON field.
	ApiClaims map[string]string

	// The API KEY that was used for authenticating the request (if any)
	ApiKey string
}

// HandlerBuilder must be implemented by adapters if they want to
// process data associated with the 'analytics' template.
//
// Mixer uses this interface to call into the adapter at configuration time to configure
// it with adapter-specific configuration as well as all template-specific type information.
type HandlerBuilder interface {
	adapter.HandlerBuilder

	// SetAnalyticsTypes is invoked by Mixer to pass the template-specific Type information for instances that an adapter
	// may receive at runtime. The type information describes the shape of the instance.
	SetAnalyticsTypes(map[string]*Type /*Instance name -> Type*/)
}

// Handler must be implemented by adapter code if it wants to
// process data associated with the 'analytics' template.
//
// Mixer uses this interface to call into the adapter at request time in order to dispatch
// created instances to the adapter. Adapters take the incoming instances and do what they
// need to achieve their primary function.
//
// The name of each instance can be used as a key into the Type map supplied to the adapter
// at configuration time via the method 'SetAnalyticsTypes'.
// These Type associated with an instance describes the shape of the instance
type Handler interface {
	adapter.Handler

	// HandleAnalytics is called by Mixer at request time to deliver instances to
	// to an adapter.
	HandleAnalytics(context.Context, []*Instance) error
}
