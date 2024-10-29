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

// VariableModeChange - struct for VariableModeChange
type VariableModeChange struct {
	VariableModeCreate *VariableModeCreate
	VariableModeDelete *VariableModeDelete
	VariableModeUpdate *VariableModeUpdate
}

// VariableModeCreateAsVariableModeChange is a convenience function that returns VariableModeCreate wrapped in VariableModeChange
func VariableModeCreateAsVariableModeChange(v *VariableModeCreate) VariableModeChange {
	return VariableModeChange{
		VariableModeCreate: v,
	}
}

// VariableModeDeleteAsVariableModeChange is a convenience function that returns VariableModeDelete wrapped in VariableModeChange
func VariableModeDeleteAsVariableModeChange(v *VariableModeDelete) VariableModeChange {
	return VariableModeChange{
		VariableModeDelete: v,
	}
}

// VariableModeUpdateAsVariableModeChange is a convenience function that returns VariableModeUpdate wrapped in VariableModeChange
func VariableModeUpdateAsVariableModeChange(v *VariableModeUpdate) VariableModeChange {
	return VariableModeChange{
		VariableModeUpdate: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *VariableModeChange) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into VariableModeCreate
	err = newStrictDecoder(data).Decode(&dst.VariableModeCreate)
	if err == nil {
		jsonVariableModeCreate, _ := json.Marshal(dst.VariableModeCreate)
		if string(jsonVariableModeCreate) == "{}" { // empty struct
			dst.VariableModeCreate = nil
		} else {
			if err = validator.Validate(dst.VariableModeCreate); err != nil {
				dst.VariableModeCreate = nil
			} else {
				match++
			}
		}
	} else {
		dst.VariableModeCreate = nil
	}

	// try to unmarshal data into VariableModeDelete
	err = newStrictDecoder(data).Decode(&dst.VariableModeDelete)
	if err == nil {
		jsonVariableModeDelete, _ := json.Marshal(dst.VariableModeDelete)
		if string(jsonVariableModeDelete) == "{}" { // empty struct
			dst.VariableModeDelete = nil
		} else {
			if err = validator.Validate(dst.VariableModeDelete); err != nil {
				dst.VariableModeDelete = nil
			} else {
				match++
			}
		}
	} else {
		dst.VariableModeDelete = nil
	}

	// try to unmarshal data into VariableModeUpdate
	err = newStrictDecoder(data).Decode(&dst.VariableModeUpdate)
	if err == nil {
		jsonVariableModeUpdate, _ := json.Marshal(dst.VariableModeUpdate)
		if string(jsonVariableModeUpdate) == "{}" { // empty struct
			dst.VariableModeUpdate = nil
		} else {
			if err = validator.Validate(dst.VariableModeUpdate); err != nil {
				dst.VariableModeUpdate = nil
			} else {
				match++
			}
		}
	} else {
		dst.VariableModeUpdate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.VariableModeCreate = nil
		dst.VariableModeDelete = nil
		dst.VariableModeUpdate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(VariableModeChange)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(VariableModeChange)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VariableModeChange) MarshalJSON() ([]byte, error) {
	if src.VariableModeCreate != nil {
		return json.Marshal(&src.VariableModeCreate)
	}

	if src.VariableModeDelete != nil {
		return json.Marshal(&src.VariableModeDelete)
	}

	if src.VariableModeUpdate != nil {
		return json.Marshal(&src.VariableModeUpdate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VariableModeChange) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.VariableModeCreate != nil {
		return obj.VariableModeCreate
	}

	if obj.VariableModeDelete != nil {
		return obj.VariableModeDelete
	}

	if obj.VariableModeUpdate != nil {
		return obj.VariableModeUpdate
	}

	// all schemas are nil
	return nil
}

type NullableVariableModeChange struct {
	value *VariableModeChange
	isSet bool
}

func (v NullableVariableModeChange) Get() *VariableModeChange {
	return v.value
}

func (v *NullableVariableModeChange) Set(val *VariableModeChange) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableModeChange) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableModeChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableModeChange(val *VariableModeChange) *NullableVariableModeChange {
	return &NullableVariableModeChange{value: val, isSet: true}
}

func (v NullableVariableModeChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableModeChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


