/*
Figma API

This is the OpenAPI specification for the [Figma REST API](https://www.figma.com/developers/api).  Note: we are releasing the OpenAPI specification as a beta given the large surface area and complexity of the REST API. If you notice any inaccuracies with the specification, please [file an issue](https://github.com/figma/rest-api-spec/issues).

API version: 0.20.0
Contact: support@figma.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ErrorResponsePayloadWithErrorBoolean type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponsePayloadWithErrorBoolean{}

// ErrorResponsePayloadWithErrorBoolean A response indicating an error occurred.
type ErrorResponsePayloadWithErrorBoolean struct {
	// For erroneous requests, this value is always `true`.
	Error bool `json:"error"`
	// Status code
	Status float32 `json:"status"`
	// A string describing the error
	Message string `json:"message"`
}

type _ErrorResponsePayloadWithErrorBoolean ErrorResponsePayloadWithErrorBoolean

// NewErrorResponsePayloadWithErrorBoolean instantiates a new ErrorResponsePayloadWithErrorBoolean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponsePayloadWithErrorBoolean(error_ bool, status float32, message string) *ErrorResponsePayloadWithErrorBoolean {
	this := ErrorResponsePayloadWithErrorBoolean{}
	this.Error = error_
	this.Status = status
	this.Message = message
	return &this
}

// NewErrorResponsePayloadWithErrorBooleanWithDefaults instantiates a new ErrorResponsePayloadWithErrorBoolean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponsePayloadWithErrorBooleanWithDefaults() *ErrorResponsePayloadWithErrorBoolean {
	this := ErrorResponsePayloadWithErrorBoolean{}
	return &this
}

// GetError returns the Error field value
func (o *ErrorResponsePayloadWithErrorBoolean) GetError() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *ErrorResponsePayloadWithErrorBoolean) GetErrorOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *ErrorResponsePayloadWithErrorBoolean) SetError(v bool) {
	o.Error = v
}

// GetStatus returns the Status field value
func (o *ErrorResponsePayloadWithErrorBoolean) GetStatus() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ErrorResponsePayloadWithErrorBoolean) GetStatusOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ErrorResponsePayloadWithErrorBoolean) SetStatus(v float32) {
	o.Status = v
}

// GetMessage returns the Message field value
func (o *ErrorResponsePayloadWithErrorBoolean) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorResponsePayloadWithErrorBoolean) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorResponsePayloadWithErrorBoolean) SetMessage(v string) {
	o.Message = v
}

func (o ErrorResponsePayloadWithErrorBoolean) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponsePayloadWithErrorBoolean) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error"] = o.Error
	toSerialize["status"] = o.Status
	toSerialize["message"] = o.Message
	return toSerialize, nil
}

func (o *ErrorResponsePayloadWithErrorBoolean) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"error",
		"status",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varErrorResponsePayloadWithErrorBoolean := _ErrorResponsePayloadWithErrorBoolean{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorResponsePayloadWithErrorBoolean)

	if err != nil {
		return err
	}

	*o = ErrorResponsePayloadWithErrorBoolean(varErrorResponsePayloadWithErrorBoolean)

	return err
}

type NullableErrorResponsePayloadWithErrorBoolean struct {
	value *ErrorResponsePayloadWithErrorBoolean
	isSet bool
}

func (v NullableErrorResponsePayloadWithErrorBoolean) Get() *ErrorResponsePayloadWithErrorBoolean {
	return v.value
}

func (v *NullableErrorResponsePayloadWithErrorBoolean) Set(val *ErrorResponsePayloadWithErrorBoolean) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponsePayloadWithErrorBoolean) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponsePayloadWithErrorBoolean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponsePayloadWithErrorBoolean(val *ErrorResponsePayloadWithErrorBoolean) *NullableErrorResponsePayloadWithErrorBoolean {
	return &NullableErrorResponsePayloadWithErrorBoolean{value: val, isSet: true}
}

func (v NullableErrorResponsePayloadWithErrorBoolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponsePayloadWithErrorBoolean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

