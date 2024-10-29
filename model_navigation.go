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

// Navigation The method of navigation. The possible values are: - `\"NAVIGATE\"`: Replaces the current screen with the destination, also closing all overlays. - `\"OVERLAY\"`: Opens the destination as an overlay on the current screen. - `\"SWAP\"`: On an overlay, replaces the current (topmost) overlay with the destination. On a top-level frame,    behaves the same as `\"NAVIGATE\"` except that no entry is added to the navigation history. - `\"SCROLL_TO\"`: Scrolls to the destination on the current screen. - `\"CHANGE_TO\"`: Changes the closest ancestor instance of source node to the specified variant.
type Navigation string

// List of Navigation
const (
	NAVIGATE Navigation = "NAVIGATE"
	SWAP Navigation = "SWAP"
	SCROLL_TO Navigation = "SCROLL_TO"
	CHANGE_TO Navigation = "CHANGE_TO"
)

// All allowed values of Navigation enum
var AllowedNavigationEnumValues = []Navigation{
	"NAVIGATE",
	"SWAP",
	"OVERLAY",
	"SCROLL_TO",
	"CHANGE_TO",
}

func (v *Navigation) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Navigation(value)
	for _, existing := range AllowedNavigationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Navigation", value)
}

// NewNavigationFromValue returns a pointer to a valid Navigation
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNavigationFromValue(v string) (*Navigation, error) {
	ev := Navigation(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Navigation: valid values are %v", v, AllowedNavigationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Navigation) IsValid() bool {
	for _, existing := range AllowedNavigationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Navigation value
func (v Navigation) Ptr() *Navigation {
	return &v
}

type NullableNavigation struct {
	value *Navigation
	isSet bool
}

func (v NullableNavigation) Get() *Navigation {
	return v.value
}

func (v *NullableNavigation) Set(val *Navigation) {
	v.value = val
	v.isSet = true
}

func (v NullableNavigation) IsSet() bool {
	return v.isSet
}

func (v *NullableNavigation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNavigation(val *Navigation) *NullableNavigation {
	return &NullableNavigation{value: val, isSet: true}
}

func (v NullableNavigation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNavigation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

