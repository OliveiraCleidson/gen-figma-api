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

// VariableDataValue - struct for VariableDataValue
type VariableDataValue struct {
	Expression *Expression
	RGB *RGB
	RGBA *RGBA
	VariableAlias *VariableAlias
	Bool *bool
	Float32 *float32
	String *string
}

// ExpressionAsVariableDataValue is a convenience function that returns Expression wrapped in VariableDataValue
func ExpressionAsVariableDataValue(v *Expression) VariableDataValue {
	return VariableDataValue{
		Expression: v,
	}
}

// RGBAsVariableDataValue is a convenience function that returns RGB wrapped in VariableDataValue
func RGBAsVariableDataValue(v *RGB) VariableDataValue {
	return VariableDataValue{
		RGB: v,
	}
}

// RGBAAsVariableDataValue is a convenience function that returns RGBA wrapped in VariableDataValue
func RGBAAsVariableDataValue(v *RGBA) VariableDataValue {
	return VariableDataValue{
		RGBA: v,
	}
}

// VariableAliasAsVariableDataValue is a convenience function that returns VariableAlias wrapped in VariableDataValue
func VariableAliasAsVariableDataValue(v *VariableAlias) VariableDataValue {
	return VariableDataValue{
		VariableAlias: v,
	}
}

// boolAsVariableDataValue is a convenience function that returns bool wrapped in VariableDataValue
func BoolAsVariableDataValue(v *bool) VariableDataValue {
	return VariableDataValue{
		Bool: v,
	}
}

// float32AsVariableDataValue is a convenience function that returns float32 wrapped in VariableDataValue
func Float32AsVariableDataValue(v *float32) VariableDataValue {
	return VariableDataValue{
		Float32: v,
	}
}

// stringAsVariableDataValue is a convenience function that returns string wrapped in VariableDataValue
func StringAsVariableDataValue(v *string) VariableDataValue {
	return VariableDataValue{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *VariableDataValue) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Expression
	err = newStrictDecoder(data).Decode(&dst.Expression)
	if err == nil {
		jsonExpression, _ := json.Marshal(dst.Expression)
		if string(jsonExpression) == "{}" { // empty struct
			dst.Expression = nil
		} else {
			if err = validator.Validate(dst.Expression); err != nil {
				dst.Expression = nil
			} else {
				match++
			}
		}
	} else {
		dst.Expression = nil
	}

	// try to unmarshal data into RGB
	err = newStrictDecoder(data).Decode(&dst.RGB)
	if err == nil {
		jsonRGB, _ := json.Marshal(dst.RGB)
		if string(jsonRGB) == "{}" { // empty struct
			dst.RGB = nil
		} else {
			if err = validator.Validate(dst.RGB); err != nil {
				dst.RGB = nil
			} else {
				match++
			}
		}
	} else {
		dst.RGB = nil
	}

	// try to unmarshal data into RGBA
	err = newStrictDecoder(data).Decode(&dst.RGBA)
	if err == nil {
		jsonRGBA, _ := json.Marshal(dst.RGBA)
		if string(jsonRGBA) == "{}" { // empty struct
			dst.RGBA = nil
		} else {
			if err = validator.Validate(dst.RGBA); err != nil {
				dst.RGBA = nil
			} else {
				match++
			}
		}
	} else {
		dst.RGBA = nil
	}

	// try to unmarshal data into VariableAlias
	err = newStrictDecoder(data).Decode(&dst.VariableAlias)
	if err == nil {
		jsonVariableAlias, _ := json.Marshal(dst.VariableAlias)
		if string(jsonVariableAlias) == "{}" { // empty struct
			dst.VariableAlias = nil
		} else {
			if err = validator.Validate(dst.VariableAlias); err != nil {
				dst.VariableAlias = nil
			} else {
				match++
			}
		}
	} else {
		dst.VariableAlias = nil
	}

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

	// try to unmarshal data into Float32
	err = newStrictDecoder(data).Decode(&dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			if err = validator.Validate(dst.Float32); err != nil {
				dst.Float32 = nil
			} else {
				match++
			}
		}
	} else {
		dst.Float32 = nil
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
		dst.Expression = nil
		dst.RGB = nil
		dst.RGBA = nil
		dst.VariableAlias = nil
		dst.Bool = nil
		dst.Float32 = nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(VariableDataValue)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(VariableDataValue)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src VariableDataValue) MarshalJSON() ([]byte, error) {
	if src.Expression != nil {
		return json.Marshal(&src.Expression)
	}

	if src.RGB != nil {
		return json.Marshal(&src.RGB)
	}

	if src.RGBA != nil {
		return json.Marshal(&src.RGBA)
	}

	if src.VariableAlias != nil {
		return json.Marshal(&src.VariableAlias)
	}

	if src.Bool != nil {
		return json.Marshal(&src.Bool)
	}

	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *VariableDataValue) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Expression != nil {
		return obj.Expression
	}

	if obj.RGB != nil {
		return obj.RGB
	}

	if obj.RGBA != nil {
		return obj.RGBA
	}

	if obj.VariableAlias != nil {
		return obj.VariableAlias
	}

	if obj.Bool != nil {
		return obj.Bool
	}

	if obj.Float32 != nil {
		return obj.Float32
	}

	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableVariableDataValue struct {
	value *VariableDataValue
	isSet bool
}

func (v NullableVariableDataValue) Get() *VariableDataValue {
	return v.value
}

func (v *NullableVariableDataValue) Set(val *VariableDataValue) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableDataValue) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableDataValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableDataValue(val *VariableDataValue) *NullableVariableDataValue {
	return &NullableVariableDataValue{value: val, isSet: true}
}

func (v NullableVariableDataValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableDataValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


