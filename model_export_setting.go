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

// checks if the ExportSetting type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportSetting{}

// ExportSetting An export setting.
type ExportSetting struct {
	Suffix string `json:"suffix"`
	Format string `json:"format"`
	Constraint Constraint `json:"constraint"`
}

type _ExportSetting ExportSetting

// NewExportSetting instantiates a new ExportSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportSetting(suffix string, format string, constraint Constraint) *ExportSetting {
	this := ExportSetting{}
	this.Suffix = suffix
	this.Format = format
	this.Constraint = constraint
	return &this
}

// NewExportSettingWithDefaults instantiates a new ExportSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportSettingWithDefaults() *ExportSetting {
	this := ExportSetting{}
	return &this
}

// GetSuffix returns the Suffix field value
func (o *ExportSetting) GetSuffix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Suffix
}

// GetSuffixOk returns a tuple with the Suffix field value
// and a boolean to check if the value has been set.
func (o *ExportSetting) GetSuffixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Suffix, true
}

// SetSuffix sets field value
func (o *ExportSetting) SetSuffix(v string) {
	o.Suffix = v
}

// GetFormat returns the Format field value
func (o *ExportSetting) GetFormat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *ExportSetting) GetFormatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *ExportSetting) SetFormat(v string) {
	o.Format = v
}

// GetConstraint returns the Constraint field value
func (o *ExportSetting) GetConstraint() Constraint {
	if o == nil {
		var ret Constraint
		return ret
	}

	return o.Constraint
}

// GetConstraintOk returns a tuple with the Constraint field value
// and a boolean to check if the value has been set.
func (o *ExportSetting) GetConstraintOk() (*Constraint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Constraint, true
}

// SetConstraint sets field value
func (o *ExportSetting) SetConstraint(v Constraint) {
	o.Constraint = v
}

func (o ExportSetting) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportSetting) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["suffix"] = o.Suffix
	toSerialize["format"] = o.Format
	toSerialize["constraint"] = o.Constraint
	return toSerialize, nil
}

func (o *ExportSetting) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"suffix",
		"format",
		"constraint",
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

	varExportSetting := _ExportSetting{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExportSetting)

	if err != nil {
		return err
	}

	*o = ExportSetting(varExportSetting)

	return err
}

type NullableExportSetting struct {
	value *ExportSetting
	isSet bool
}

func (v NullableExportSetting) Get() *ExportSetting {
	return v.value
}

func (v *NullableExportSetting) Set(val *ExportSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableExportSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableExportSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportSetting(val *ExportSetting) *NullableExportSetting {
	return &NullableExportSetting{value: val, isSet: true}
}

func (v NullableExportSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


