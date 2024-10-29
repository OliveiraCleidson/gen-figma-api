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
	"gopkg.in/validator.v2"
	"fmt"
)

// ComponentPropertyDefinitionDefaultValue - Initial value of this property for instances.
type ComponentPropertyDefinitionDefaultValue struct {
	Bool *bool
	String *string
}

// boolAsComponentPropertyDefinitionDefaultValue is a convenience function that returns bool wrapped in ComponentPropertyDefinitionDefaultValue
func BoolAsComponentPropertyDefinitionDefaultValue(v *bool) ComponentPropertyDefinitionDefaultValue {
	return ComponentPropertyDefinitionDefaultValue{
		Bool: v,
	}
}

// stringAsComponentPropertyDefinitionDefaultValue is a convenience function that returns string wrapped in ComponentPropertyDefinitionDefaultValue
func StringAsComponentPropertyDefinitionDefaultValue(v *string) ComponentPropertyDefinitionDefaultValue {
	return ComponentPropertyDefinitionDefaultValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ComponentPropertyDefinitionDefaultValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Bool
	err = newStrictDecoder(data).Decode(&dst.Bool)
	if err == nil {
		jsonBool, _ := json.Marshal(dst.Bool)
		if string(jsonBool) == "{}" { // empty struct
			dst.Bool = nil
		} else {
			if err = validator.Validate(dst.Bool); err != nil {
				dst.Bool = nil
			} else {
				match++
			}
		}
	} else {
		dst.Bool = nil
	}

	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			if err = validator.Validate(dst.String); err != nil {
				dst.String = nil
			} else {
				match++
			}
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Bool = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ComponentPropertyDefinitionDefaultValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ComponentPropertyDefinitionDefaultValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ComponentPropertyDefinitionDefaultValue) MarshalJSON() ([]byte, error) {
	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ComponentPropertyDefinitionDefaultValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableComponentPropertyDefinitionDefaultValue struct {
	value *ComponentPropertyDefinitionDefaultValue
	isSet bool
}

func (v NullableComponentPropertyDefinitionDefaultValue) Get() *ComponentPropertyDefinitionDefaultValue {
	return v.value
}

func (v *NullableComponentPropertyDefinitionDefaultValue) Set(val *ComponentPropertyDefinitionDefaultValue) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentPropertyDefinitionDefaultValue) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentPropertyDefinitionDefaultValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentPropertyDefinitionDefaultValue(val *ComponentPropertyDefinitionDefaultValue) *NullableComponentPropertyDefinitionDefaultValue {
	return &NullableComponentPropertyDefinitionDefaultValue{value: val, isSet: true}
}

func (v NullableComponentPropertyDefinitionDefaultValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentPropertyDefinitionDefaultValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

