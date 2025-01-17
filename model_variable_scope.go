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
	"fmt"
)

// VariableScope Scopes allow a variable to be shown or hidden in the variable picker for various fields. This declutters the Figma UI if you have a large number of variables. Variable scopes are currently supported on `FLOAT`, `STRING`, and `COLOR` variables.  `ALL_SCOPES` is a special scope that means that the variable will be shown in the variable picker for all variable fields. If `ALL_SCOPES` is set, no additional scopes can be set.  `ALL_FILLS` is a special scope that means that the variable will be shown in the variable picker for all fill fields. If `ALL_FILLS` is set, no additional fill scopes can be set.  Valid scopes for `FLOAT` variables: - `ALL_SCOPES` - `TEXT_CONTENT` - `WIDTH_HEIGHT` - `GAP` - `STROKE_FLOAT` - `EFFECT_FLOAT` - `OPACITY` - `FONT_WEIGHT` - `FONT_SIZE` - `LINE_HEIGHT` - `LETTER_SPACING` - `PARAGRAPH_SPACING` - `PARAGRAPH_INDENT`  Valid scopes for `STRING` variables: - `ALL_SCOPES` - `TEXT_CONTENT` - `FONT_FAMILY` - `FONT_STYLE`  Valid scopes for `COLOR` variables: - `ALL_SCOPES` - `ALL_FILLS` - `FRAME_FILL` - `SHAPE_FILL` - `TEXT_FILL` - `STROKE_COLOR` - `EFFECT_COLOR`
type VariableScope string

// List of VariableScope
const (
	ALL_SCOPES VariableScope = "ALL_SCOPES"
	TEXT_CONTENT VariableScope = "TEXT_CONTENT"
	CORNER_RADIUS VariableScope = "CORNER_RADIUS"
	WIDTH_HEIGHT VariableScope = "WIDTH_HEIGHT"
	GAP VariableScope = "GAP"
	ALL_FILLS VariableScope = "ALL_FILLS"
	FRAME_FILL VariableScope = "FRAME_FILL"
	SHAPE_FILL VariableScope = "SHAPE_FILL"
	TEXT_FILL VariableScope = "TEXT_FILL"
	STROKE_COLOR VariableScope = "STROKE_COLOR"
	STROKE_FLOAT VariableScope = "STROKE_FLOAT"
	EFFECT_FLOAT VariableScope = "EFFECT_FLOAT"
	EFFECT_COLOR VariableScope = "EFFECT_COLOR"
	OPACITY VariableScope = "OPACITY"
	FONT_FAMILY VariableScope = "FONT_FAMILY"
	FONT_STYLE VariableScope = "FONT_STYLE"
	FONT_WEIGHT VariableScope = "FONT_WEIGHT"
	FONT_SIZE VariableScope = "FONT_SIZE"
	LINE_HEIGHT VariableScope = "LINE_HEIGHT"
	LETTER_SPACING VariableScope = "LETTER_SPACING"
	PARAGRAPH_SPACING VariableScope = "PARAGRAPH_SPACING"
	PARAGRAPH_INDENT VariableScope = "PARAGRAPH_INDENT"
)

// All allowed values of VariableScope enum
var AllowedVariableScopeEnumValues = []VariableScope{
	"ALL_SCOPES",
	"TEXT_CONTENT",
	"CORNER_RADIUS",
	"WIDTH_HEIGHT",
	"GAP",
	"ALL_FILLS",
	"FRAME_FILL",
	"SHAPE_FILL",
	"TEXT_FILL",
	"STROKE_COLOR",
	"STROKE_FLOAT",
	"EFFECT_FLOAT",
	"EFFECT_COLOR",
	"OPACITY",
	"FONT_FAMILY",
	"FONT_STYLE",
	"FONT_WEIGHT",
	"FONT_SIZE",
	"LINE_HEIGHT",
	"LETTER_SPACING",
	"PARAGRAPH_SPACING",
	"PARAGRAPH_INDENT",
}

func (v *VariableScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VariableScope(value)
	for _, existing := range AllowedVariableScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VariableScope", value)
}

// NewVariableScopeFromValue returns a pointer to a valid VariableScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVariableScopeFromValue(v string) (*VariableScope, error) {
	ev := VariableScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VariableScope: valid values are %v", v, AllowedVariableScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VariableScope) IsValid() bool {
	for _, existing := range AllowedVariableScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VariableScope value
func (v VariableScope) Ptr() *VariableScope {
	return &v
}

type NullableVariableScope struct {
	value *VariableScope
	isSet bool
}

func (v NullableVariableScope) Get() *VariableScope {
	return v.value
}

func (v *NullableVariableScope) Set(val *VariableScope) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableScope) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableScope(val *VariableScope) *NullableVariableScope {
	return &NullableVariableScope{value: val, isSet: true}
}

func (v NullableVariableScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

