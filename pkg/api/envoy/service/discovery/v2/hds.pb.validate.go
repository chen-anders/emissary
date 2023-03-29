// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/service/discovery/v2/hds.proto

package discoveryv2

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	core "github.com/emissary-ingress/emissary/v3/pkg/api/envoy/api/v2/core"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort

	_ = core.HealthStatus(0)
)

// Validate checks the field values on Capability with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Capability) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Capability with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CapabilityMultiError, or
// nil if none found.
func (m *Capability) ValidateAll() error {
	return m.validate(true)
}

func (m *Capability) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return CapabilityMultiError(errors)
	}

	return nil
}

// CapabilityMultiError is an error wrapping multiple validation errors
// returned by Capability.ValidateAll() if the designated constraints aren't met.
type CapabilityMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CapabilityMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CapabilityMultiError) AllErrors() []error { return m }

// CapabilityValidationError is the validation error returned by
// Capability.Validate if the designated constraints aren't met.
type CapabilityValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CapabilityValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CapabilityValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CapabilityValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CapabilityValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CapabilityValidationError) ErrorName() string { return "CapabilityValidationError" }

// Error satisfies the builtin error interface
func (e CapabilityValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCapability.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CapabilityValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CapabilityValidationError{}

// Validate checks the field values on HealthCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckRequestMultiError, or nil if none found.
func (m *HealthCheckRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetNode()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HealthCheckRequestValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HealthCheckRequestValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetNode()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckRequestValidationError{
				field:  "Node",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCapability()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HealthCheckRequestValidationError{
					field:  "Capability",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HealthCheckRequestValidationError{
					field:  "Capability",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCapability()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckRequestValidationError{
				field:  "Capability",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return HealthCheckRequestMultiError(errors)
	}

	return nil
}

// HealthCheckRequestMultiError is an error wrapping multiple validation errors
// returned by HealthCheckRequest.ValidateAll() if the designated constraints
// aren't met.
type HealthCheckRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckRequestMultiError) AllErrors() []error { return m }

// HealthCheckRequestValidationError is the validation error returned by
// HealthCheckRequest.Validate if the designated constraints aren't met.
type HealthCheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckRequestValidationError) ErrorName() string {
	return "HealthCheckRequestValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckRequestValidationError{}

// Validate checks the field values on EndpointHealth with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EndpointHealth) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EndpointHealth with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EndpointHealthMultiError,
// or nil if none found.
func (m *EndpointHealth) ValidateAll() error {
	return m.validate(true)
}

func (m *EndpointHealth) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEndpoint()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EndpointHealthValidationError{
					field:  "Endpoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EndpointHealthValidationError{
					field:  "Endpoint",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEndpoint()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EndpointHealthValidationError{
				field:  "Endpoint",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for HealthStatus

	if len(errors) > 0 {
		return EndpointHealthMultiError(errors)
	}

	return nil
}

// EndpointHealthMultiError is an error wrapping multiple validation errors
// returned by EndpointHealth.ValidateAll() if the designated constraints
// aren't met.
type EndpointHealthMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EndpointHealthMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EndpointHealthMultiError) AllErrors() []error { return m }

// EndpointHealthValidationError is the validation error returned by
// EndpointHealth.Validate if the designated constraints aren't met.
type EndpointHealthValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointHealthValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointHealthValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointHealthValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointHealthValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointHealthValidationError) ErrorName() string { return "EndpointHealthValidationError" }

// Error satisfies the builtin error interface
func (e EndpointHealthValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpointHealth.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointHealthValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointHealthValidationError{}

// Validate checks the field values on EndpointHealthResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *EndpointHealthResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EndpointHealthResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// EndpointHealthResponseMultiError, or nil if none found.
func (m *EndpointHealthResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *EndpointHealthResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEndpointsHealth() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EndpointHealthResponseValidationError{
						field:  fmt.Sprintf("EndpointsHealth[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EndpointHealthResponseValidationError{
						field:  fmt.Sprintf("EndpointsHealth[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EndpointHealthResponseValidationError{
					field:  fmt.Sprintf("EndpointsHealth[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return EndpointHealthResponseMultiError(errors)
	}

	return nil
}

// EndpointHealthResponseMultiError is an error wrapping multiple validation
// errors returned by EndpointHealthResponse.ValidateAll() if the designated
// constraints aren't met.
type EndpointHealthResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EndpointHealthResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EndpointHealthResponseMultiError) AllErrors() []error { return m }

// EndpointHealthResponseValidationError is the validation error returned by
// EndpointHealthResponse.Validate if the designated constraints aren't met.
type EndpointHealthResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EndpointHealthResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EndpointHealthResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EndpointHealthResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EndpointHealthResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EndpointHealthResponseValidationError) ErrorName() string {
	return "EndpointHealthResponseValidationError"
}

// Error satisfies the builtin error interface
func (e EndpointHealthResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEndpointHealthResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EndpointHealthResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EndpointHealthResponseValidationError{}

// Validate checks the field values on
// HealthCheckRequestOrEndpointHealthResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *HealthCheckRequestOrEndpointHealthResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// HealthCheckRequestOrEndpointHealthResponse with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in
// HealthCheckRequestOrEndpointHealthResponseMultiError, or nil if none found.
func (m *HealthCheckRequestOrEndpointHealthResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckRequestOrEndpointHealthResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	switch v := m.RequestType.(type) {
	case *HealthCheckRequestOrEndpointHealthResponse_HealthCheckRequest:
		if v == nil {
			err := HealthCheckRequestOrEndpointHealthResponseValidationError{
				field:  "RequestType",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetHealthCheckRequest()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckRequestOrEndpointHealthResponseValidationError{
						field:  "HealthCheckRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckRequestOrEndpointHealthResponseValidationError{
						field:  "HealthCheckRequest",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetHealthCheckRequest()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckRequestOrEndpointHealthResponseValidationError{
					field:  "HealthCheckRequest",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *HealthCheckRequestOrEndpointHealthResponse_EndpointHealthResponse:
		if v == nil {
			err := HealthCheckRequestOrEndpointHealthResponseValidationError{
				field:  "RequestType",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetEndpointHealthResponse()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckRequestOrEndpointHealthResponseValidationError{
						field:  "EndpointHealthResponse",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckRequestOrEndpointHealthResponseValidationError{
						field:  "EndpointHealthResponse",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEndpointHealthResponse()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckRequestOrEndpointHealthResponseValidationError{
					field:  "EndpointHealthResponse",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}

	if len(errors) > 0 {
		return HealthCheckRequestOrEndpointHealthResponseMultiError(errors)
	}

	return nil
}

// HealthCheckRequestOrEndpointHealthResponseMultiError is an error wrapping
// multiple validation errors returned by
// HealthCheckRequestOrEndpointHealthResponse.ValidateAll() if the designated
// constraints aren't met.
type HealthCheckRequestOrEndpointHealthResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckRequestOrEndpointHealthResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckRequestOrEndpointHealthResponseMultiError) AllErrors() []error { return m }

// HealthCheckRequestOrEndpointHealthResponseValidationError is the validation
// error returned by HealthCheckRequestOrEndpointHealthResponse.Validate if
// the designated constraints aren't met.
type HealthCheckRequestOrEndpointHealthResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) ErrorName() string {
	return "HealthCheckRequestOrEndpointHealthResponseValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckRequestOrEndpointHealthResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckRequestOrEndpointHealthResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckRequestOrEndpointHealthResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckRequestOrEndpointHealthResponseValidationError{}

// Validate checks the field values on LocalityEndpoints with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LocalityEndpoints) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LocalityEndpoints with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LocalityEndpointsMultiError, or nil if none found.
func (m *LocalityEndpoints) ValidateAll() error {
	return m.validate(true)
}

func (m *LocalityEndpoints) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetLocality()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LocalityEndpointsValidationError{
					field:  "Locality",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LocalityEndpointsValidationError{
					field:  "Locality",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetLocality()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LocalityEndpointsValidationError{
				field:  "Locality",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetEndpoints() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, LocalityEndpointsValidationError{
						field:  fmt.Sprintf("Endpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, LocalityEndpointsValidationError{
						field:  fmt.Sprintf("Endpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return LocalityEndpointsValidationError{
					field:  fmt.Sprintf("Endpoints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return LocalityEndpointsMultiError(errors)
	}

	return nil
}

// LocalityEndpointsMultiError is an error wrapping multiple validation errors
// returned by LocalityEndpoints.ValidateAll() if the designated constraints
// aren't met.
type LocalityEndpointsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LocalityEndpointsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LocalityEndpointsMultiError) AllErrors() []error { return m }

// LocalityEndpointsValidationError is the validation error returned by
// LocalityEndpoints.Validate if the designated constraints aren't met.
type LocalityEndpointsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalityEndpointsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalityEndpointsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalityEndpointsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalityEndpointsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalityEndpointsValidationError) ErrorName() string {
	return "LocalityEndpointsValidationError"
}

// Error satisfies the builtin error interface
func (e LocalityEndpointsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocalityEndpoints.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalityEndpointsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalityEndpointsValidationError{}

// Validate checks the field values on ClusterHealthCheck with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ClusterHealthCheck) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClusterHealthCheck with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ClusterHealthCheckMultiError, or nil if none found.
func (m *ClusterHealthCheck) ValidateAll() error {
	return m.validate(true)
}

func (m *ClusterHealthCheck) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ClusterName

	for idx, item := range m.GetHealthChecks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClusterHealthCheckValidationError{
						field:  fmt.Sprintf("HealthChecks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClusterHealthCheckValidationError{
						field:  fmt.Sprintf("HealthChecks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterHealthCheckValidationError{
					field:  fmt.Sprintf("HealthChecks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetLocalityEndpoints() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ClusterHealthCheckValidationError{
						field:  fmt.Sprintf("LocalityEndpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ClusterHealthCheckValidationError{
						field:  fmt.Sprintf("LocalityEndpoints[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ClusterHealthCheckValidationError{
					field:  fmt.Sprintf("LocalityEndpoints[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ClusterHealthCheckMultiError(errors)
	}

	return nil
}

// ClusterHealthCheckMultiError is an error wrapping multiple validation errors
// returned by ClusterHealthCheck.ValidateAll() if the designated constraints
// aren't met.
type ClusterHealthCheckMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClusterHealthCheckMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClusterHealthCheckMultiError) AllErrors() []error { return m }

// ClusterHealthCheckValidationError is the validation error returned by
// ClusterHealthCheck.Validate if the designated constraints aren't met.
type ClusterHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterHealthCheckValidationError) ErrorName() string {
	return "ClusterHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterHealthCheckValidationError{}

// Validate checks the field values on HealthCheckSpecifier with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *HealthCheckSpecifier) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on HealthCheckSpecifier with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// HealthCheckSpecifierMultiError, or nil if none found.
func (m *HealthCheckSpecifier) ValidateAll() error {
	return m.validate(true)
}

func (m *HealthCheckSpecifier) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetClusterHealthChecks() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, HealthCheckSpecifierValidationError{
						field:  fmt.Sprintf("ClusterHealthChecks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, HealthCheckSpecifierValidationError{
						field:  fmt.Sprintf("ClusterHealthChecks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return HealthCheckSpecifierValidationError{
					field:  fmt.Sprintf("ClusterHealthChecks[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, HealthCheckSpecifierValidationError{
					field:  "Interval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, HealthCheckSpecifierValidationError{
					field:  "Interval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return HealthCheckSpecifierValidationError{
				field:  "Interval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return HealthCheckSpecifierMultiError(errors)
	}

	return nil
}

// HealthCheckSpecifierMultiError is an error wrapping multiple validation
// errors returned by HealthCheckSpecifier.ValidateAll() if the designated
// constraints aren't met.
type HealthCheckSpecifierMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m HealthCheckSpecifierMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m HealthCheckSpecifierMultiError) AllErrors() []error { return m }

// HealthCheckSpecifierValidationError is the validation error returned by
// HealthCheckSpecifier.Validate if the designated constraints aren't met.
type HealthCheckSpecifierValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckSpecifierValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckSpecifierValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckSpecifierValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckSpecifierValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckSpecifierValidationError) ErrorName() string {
	return "HealthCheckSpecifierValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckSpecifierValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckSpecifier.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckSpecifierValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckSpecifierValidationError{}
