// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/adaptive_concurrency/v2alpha/adaptive_concurrency.proto

package v2alpha

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
)

// Validate checks the field values on GradientControllerConfig with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GradientControllerConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GradientControllerConfig with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GradientControllerConfigMultiError, or nil if none found.
func (m *GradientControllerConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *GradientControllerConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetSampleAggregatePercentile()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GradientControllerConfigValidationError{
					field:  "SampleAggregatePercentile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GradientControllerConfigValidationError{
					field:  "SampleAggregatePercentile",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetSampleAggregatePercentile()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GradientControllerConfigValidationError{
				field:  "SampleAggregatePercentile",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetConcurrencyLimitParams() == nil {
		err := GradientControllerConfigValidationError{
			field:  "ConcurrencyLimitParams",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetConcurrencyLimitParams()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GradientControllerConfigValidationError{
					field:  "ConcurrencyLimitParams",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GradientControllerConfigValidationError{
					field:  "ConcurrencyLimitParams",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConcurrencyLimitParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GradientControllerConfigValidationError{
				field:  "ConcurrencyLimitParams",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetMinRttCalcParams() == nil {
		err := GradientControllerConfigValidationError{
			field:  "MinRttCalcParams",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if all {
		switch v := interface{}(m.GetMinRttCalcParams()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GradientControllerConfigValidationError{
					field:  "MinRttCalcParams",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GradientControllerConfigValidationError{
					field:  "MinRttCalcParams",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMinRttCalcParams()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GradientControllerConfigValidationError{
				field:  "MinRttCalcParams",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GradientControllerConfigMultiError(errors)
	}

	return nil
}

// GradientControllerConfigMultiError is an error wrapping multiple validation
// errors returned by GradientControllerConfig.ValidateAll() if the designated
// constraints aren't met.
type GradientControllerConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GradientControllerConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GradientControllerConfigMultiError) AllErrors() []error { return m }

// GradientControllerConfigValidationError is the validation error returned by
// GradientControllerConfig.Validate if the designated constraints aren't met.
type GradientControllerConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GradientControllerConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GradientControllerConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GradientControllerConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GradientControllerConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GradientControllerConfigValidationError) ErrorName() string {
	return "GradientControllerConfigValidationError"
}

// Error satisfies the builtin error interface
func (e GradientControllerConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGradientControllerConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GradientControllerConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GradientControllerConfigValidationError{}

// Validate checks the field values on AdaptiveConcurrency with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AdaptiveConcurrency) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AdaptiveConcurrency with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AdaptiveConcurrencyMultiError, or nil if none found.
func (m *AdaptiveConcurrency) ValidateAll() error {
	return m.validate(true)
}

func (m *AdaptiveConcurrency) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AdaptiveConcurrencyValidationError{
					field:  "Enabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AdaptiveConcurrencyValidationError{
					field:  "Enabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AdaptiveConcurrencyValidationError{
				field:  "Enabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	oneofConcurrencyControllerConfigPresent := false
	switch v := m.ConcurrencyControllerConfig.(type) {
	case *AdaptiveConcurrency_GradientControllerConfig:
		if v == nil {
			err := AdaptiveConcurrencyValidationError{
				field:  "ConcurrencyControllerConfig",
				reason: "oneof value cannot be a typed-nil",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}
		oneofConcurrencyControllerConfigPresent = true

		if m.GetGradientControllerConfig() == nil {
			err := AdaptiveConcurrencyValidationError{
				field:  "GradientControllerConfig",
				reason: "value is required",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

		if all {
			switch v := interface{}(m.GetGradientControllerConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, AdaptiveConcurrencyValidationError{
						field:  "GradientControllerConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, AdaptiveConcurrencyValidationError{
						field:  "GradientControllerConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetGradientControllerConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AdaptiveConcurrencyValidationError{
					field:  "GradientControllerConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		_ = v // ensures v is used
	}
	if !oneofConcurrencyControllerConfigPresent {
		err := AdaptiveConcurrencyValidationError{
			field:  "ConcurrencyControllerConfig",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return AdaptiveConcurrencyMultiError(errors)
	}

	return nil
}

// AdaptiveConcurrencyMultiError is an error wrapping multiple validation
// errors returned by AdaptiveConcurrency.ValidateAll() if the designated
// constraints aren't met.
type AdaptiveConcurrencyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AdaptiveConcurrencyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AdaptiveConcurrencyMultiError) AllErrors() []error { return m }

// AdaptiveConcurrencyValidationError is the validation error returned by
// AdaptiveConcurrency.Validate if the designated constraints aren't met.
type AdaptiveConcurrencyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdaptiveConcurrencyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdaptiveConcurrencyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdaptiveConcurrencyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdaptiveConcurrencyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdaptiveConcurrencyValidationError) ErrorName() string {
	return "AdaptiveConcurrencyValidationError"
}

// Error satisfies the builtin error interface
func (e AdaptiveConcurrencyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdaptiveConcurrency.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdaptiveConcurrencyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdaptiveConcurrencyValidationError{}

// Validate checks the field values on
// GradientControllerConfig_ConcurrencyLimitCalculationParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GradientControllerConfig_ConcurrencyLimitCalculationParams with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GradientControllerConfig_ConcurrencyLimitCalculationParamsMultiError, or
// nil if none found.
func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) ValidateAll() error {
	return m.validate(true)
}

func (m *GradientControllerConfig_ConcurrencyLimitCalculationParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if wrapper := m.GetMaxConcurrencyLimit(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			err := GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{
				field:  "MaxConcurrencyLimit",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if m.GetConcurrencyUpdateInterval() == nil {
		err := GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{
			field:  "ConcurrencyUpdateInterval",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetConcurrencyUpdateInterval(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{
				field:  "ConcurrencyUpdateInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{
					field:  "ConcurrencyUpdateInterval",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if len(errors) > 0 {
		return GradientControllerConfig_ConcurrencyLimitCalculationParamsMultiError(errors)
	}

	return nil
}

// GradientControllerConfig_ConcurrencyLimitCalculationParamsMultiError is an
// error wrapping multiple validation errors returned by
// GradientControllerConfig_ConcurrencyLimitCalculationParams.ValidateAll() if
// the designated constraints aren't met.
type GradientControllerConfig_ConcurrencyLimitCalculationParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GradientControllerConfig_ConcurrencyLimitCalculationParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GradientControllerConfig_ConcurrencyLimitCalculationParamsMultiError) AllErrors() []error {
	return m
}

// GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError is
// the validation error returned by
// GradientControllerConfig_ConcurrencyLimitCalculationParams.Validate if the
// designated constraints aren't met.
type GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) Key() bool {
	return e.key
}

// ErrorName returns error name.
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) ErrorName() string {
	return "GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError"
}

// Error satisfies the builtin error interface
func (e GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGradientControllerConfig_ConcurrencyLimitCalculationParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GradientControllerConfig_ConcurrencyLimitCalculationParamsValidationError{}

// Validate checks the field values on
// GradientControllerConfig_MinimumRTTCalculationParams with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GradientControllerConfig_MinimumRTTCalculationParams) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on
// GradientControllerConfig_MinimumRTTCalculationParams with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in
// GradientControllerConfig_MinimumRTTCalculationParamsMultiError, or nil if
// none found.
func (m *GradientControllerConfig_MinimumRTTCalculationParams) ValidateAll() error {
	return m.validate(true)
}

func (m *GradientControllerConfig_MinimumRTTCalculationParams) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if m.GetInterval() == nil {
		err := GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
			field:  "Interval",
			reason: "value is required",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if d := m.GetInterval(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
				field:  "Interval",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(0*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
					field:  "Interval",
					reason: "value must be greater than 0s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	if wrapper := m.GetRequestCount(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			err := GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
				field:  "RequestCount",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetJitter()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
					field:  "Jitter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
					field:  "Jitter",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetJitter()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
				field:  "Jitter",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if wrapper := m.GetMinConcurrency(); wrapper != nil {

		if wrapper.GetValue() <= 0 {
			err := GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
				field:  "MinConcurrency",
				reason: "value must be greater than 0",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetBuffer()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
					field:  "Buffer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
					field:  "Buffer",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBuffer()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GradientControllerConfig_MinimumRTTCalculationParamsValidationError{
				field:  "Buffer",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GradientControllerConfig_MinimumRTTCalculationParamsMultiError(errors)
	}

	return nil
}

// GradientControllerConfig_MinimumRTTCalculationParamsMultiError is an error
// wrapping multiple validation errors returned by
// GradientControllerConfig_MinimumRTTCalculationParams.ValidateAll() if the
// designated constraints aren't met.
type GradientControllerConfig_MinimumRTTCalculationParamsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GradientControllerConfig_MinimumRTTCalculationParamsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GradientControllerConfig_MinimumRTTCalculationParamsMultiError) AllErrors() []error { return m }

// GradientControllerConfig_MinimumRTTCalculationParamsValidationError is the
// validation error returned by
// GradientControllerConfig_MinimumRTTCalculationParams.Validate if the
// designated constraints aren't met.
type GradientControllerConfig_MinimumRTTCalculationParamsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) Field() string {
	return e.field
}

// Reason function returns reason value.
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) Reason() string {
	return e.reason
}

// Cause function returns cause value.
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) Cause() error {
	return e.cause
}

// Key function returns key value.
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) ErrorName() string {
	return "GradientControllerConfig_MinimumRTTCalculationParamsValidationError"
}

// Error satisfies the builtin error interface
func (e GradientControllerConfig_MinimumRTTCalculationParamsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGradientControllerConfig_MinimumRTTCalculationParams.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GradientControllerConfig_MinimumRTTCalculationParamsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GradientControllerConfig_MinimumRTTCalculationParamsValidationError{}
