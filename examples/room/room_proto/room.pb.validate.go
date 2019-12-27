// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: examples/room/room_proto/room.proto

package room_proto

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/gogo/protobuf/types"
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
	_ = types.DynamicAny{}
)

// Validate checks the field values on UserModel with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *UserModel) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Nickname

	// no validation rules for Avatar

	return nil
}

// UserModelValidationError is the validation error returned by
// UserModel.Validate if the designated constraints aren't met.
type UserModelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserModelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserModelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserModelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserModelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserModelValidationError) ErrorName() string { return "UserModelValidationError" }

// Error satisfies the builtin error interface
func (e UserModelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserModel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserModelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserModelValidationError{}

// Validate checks the field values on RoomModel with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RoomModel) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for Name

	// no validation rules for Users

	return nil
}

// RoomModelValidationError is the validation error returned by
// RoomModel.Validate if the designated constraints aren't met.
type RoomModelValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RoomModelValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RoomModelValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RoomModelValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RoomModelValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RoomModelValidationError) ErrorName() string { return "RoomModelValidationError" }

// Error satisfies the builtin error interface
func (e RoomModelValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRoomModel.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RoomModelValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RoomModelValidationError{}

// Validate checks the field values on CreateRoomInput with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateRoomInput) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	return nil
}

// CreateRoomInputValidationError is the validation error returned by
// CreateRoomInput.Validate if the designated constraints aren't met.
type CreateRoomInputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRoomInputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRoomInputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRoomInputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRoomInputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRoomInputValidationError) ErrorName() string { return "CreateRoomInputValidationError" }

// Error satisfies the builtin error interface
func (e CreateRoomInputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRoomInput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRoomInputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRoomInputValidationError{}

// Validate checks the field values on ListRoomInput with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListRoomInput) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// ListRoomInputValidationError is the validation error returned by
// ListRoomInput.Validate if the designated constraints aren't met.
type ListRoomInputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoomInputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoomInputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoomInputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoomInputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoomInputValidationError) ErrorName() string { return "ListRoomInputValidationError" }

// Error satisfies the builtin error interface
func (e ListRoomInputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoomInput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoomInputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoomInputValidationError{}

// Validate checks the field values on ListRoomOutput with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ListRoomOutput) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Total

	for idx, item := range m.GetRooms() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return ListRoomOutputValidationError{
						field:  fmt.Sprintf("Rooms[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// ListRoomOutputValidationError is the validation error returned by
// ListRoomOutput.Validate if the designated constraints aren't met.
type ListRoomOutputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRoomOutputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRoomOutputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRoomOutputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRoomOutputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRoomOutputValidationError) ErrorName() string { return "ListRoomOutputValidationError" }

// Error satisfies the builtin error interface
func (e ListRoomOutputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRoomOutput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRoomOutputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRoomOutputValidationError{}

// Validate checks the field values on JoinRoomInput with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *JoinRoomInput) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// JoinRoomInputValidationError is the validation error returned by
// JoinRoomInput.Validate if the designated constraints aren't met.
type JoinRoomInputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e JoinRoomInputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e JoinRoomInputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e JoinRoomInputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e JoinRoomInputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e JoinRoomInputValidationError) ErrorName() string { return "JoinRoomInputValidationError" }

// Error satisfies the builtin error interface
func (e JoinRoomInputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sJoinRoomInput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = JoinRoomInputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = JoinRoomInputValidationError{}

// Validate checks the field values on ExitRoomInput with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ExitRoomInput) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	return nil
}

// ExitRoomInputValidationError is the validation error returned by
// ExitRoomInput.Validate if the designated constraints aren't met.
type ExitRoomInputValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ExitRoomInputValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ExitRoomInputValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ExitRoomInputValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ExitRoomInputValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ExitRoomInputValidationError) ErrorName() string { return "ExitRoomInputValidationError" }

// Error satisfies the builtin error interface
func (e ExitRoomInputValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sExitRoomInput.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ExitRoomInputValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ExitRoomInputValidationError{}

// Validate checks the field values on UserEnterEvent with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserEnterEvent) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RoomId

	{
		tmp := m.GetUser()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return UserEnterEventValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// UserEnterEventValidationError is the validation error returned by
// UserEnterEvent.Validate if the designated constraints aren't met.
type UserEnterEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserEnterEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserEnterEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserEnterEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserEnterEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserEnterEventValidationError) ErrorName() string { return "UserEnterEventValidationError" }

// Error satisfies the builtin error interface
func (e UserEnterEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserEnterEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserEnterEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserEnterEventValidationError{}

// Validate checks the field values on UserExitEvent with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *UserExitEvent) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RoomId

	// no validation rules for UserId

	return nil
}

// UserExitEventValidationError is the validation error returned by
// UserExitEvent.Validate if the designated constraints aren't met.
type UserExitEventValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserExitEventValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserExitEventValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserExitEventValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserExitEventValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserExitEventValidationError) ErrorName() string { return "UserExitEventValidationError" }

// Error satisfies the builtin error interface
func (e UserExitEventValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUserExitEvent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserExitEventValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserExitEventValidationError{}