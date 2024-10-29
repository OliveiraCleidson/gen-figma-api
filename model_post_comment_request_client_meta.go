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

// PostCommentRequestClientMeta - The position where to place the comment.
type PostCommentRequestClientMeta struct {
	FrameOffset *FrameOffset
	FrameOffsetRegion *FrameOffsetRegion
	Region *Region
	Vector *Vector
}

// FrameOffsetAsPostCommentRequestClientMeta is a convenience function that returns FrameOffset wrapped in PostCommentRequestClientMeta
func FrameOffsetAsPostCommentRequestClientMeta(v *FrameOffset) PostCommentRequestClientMeta {
	return PostCommentRequestClientMeta{
		FrameOffset: v,
	}
}

// FrameOffsetRegionAsPostCommentRequestClientMeta is a convenience function that returns FrameOffsetRegion wrapped in PostCommentRequestClientMeta
func FrameOffsetRegionAsPostCommentRequestClientMeta(v *FrameOffsetRegion) PostCommentRequestClientMeta {
	return PostCommentRequestClientMeta{
		FrameOffsetRegion: v,
	}
}

// RegionAsPostCommentRequestClientMeta is a convenience function that returns Region wrapped in PostCommentRequestClientMeta
func RegionAsPostCommentRequestClientMeta(v *Region) PostCommentRequestClientMeta {
	return PostCommentRequestClientMeta{
		Region: v,
	}
}

// VectorAsPostCommentRequestClientMeta is a convenience function that returns Vector wrapped in PostCommentRequestClientMeta
func VectorAsPostCommentRequestClientMeta(v *Vector) PostCommentRequestClientMeta {
	return PostCommentRequestClientMeta{
		Vector: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PostCommentRequestClientMeta) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into FrameOffset
	err = newStrictDecoder(data).Decode(&dst.FrameOffset)
	if err == nil {
		jsonFrameOffset, _ := json.Marshal(dst.FrameOffset)
		if string(jsonFrameOffset) == "{}" { // empty struct
			dst.FrameOffset = nil
		} else {
			if err = validator.Validate(dst.FrameOffset); err != nil {
				dst.FrameOffset = nil
			} else {
				match++
			}
		}
	} else {
		dst.FrameOffset = nil
	}

	// try to unmarshal data into FrameOffsetRegion
	err = newStrictDecoder(data).Decode(&dst.FrameOffsetRegion)
	if err == nil {
		jsonFrameOffsetRegion, _ := json.Marshal(dst.FrameOffsetRegion)
		if string(jsonFrameOffsetRegion) == "{}" { // empty struct
			dst.FrameOffsetRegion = nil
		} else {
			if err = validator.Validate(dst.FrameOffsetRegion); err != nil {
				dst.FrameOffsetRegion = nil
			} else {
				match++
			}
		}
	} else {
		dst.FrameOffsetRegion = nil
	}

	// try to unmarshal data into Region
	err = newStrictDecoder(data).Decode(&dst.Region)
	if err == nil {
		jsonRegion, _ := json.Marshal(dst.Region)
		if string(jsonRegion) == "{}" { // empty struct
			dst.Region = nil
		} else {
			if err = validator.Validate(dst.Region); err != nil {
				dst.Region = nil
			} else {
				match++
			}
		}
	} else {
		dst.Region = nil
	}

	// try to unmarshal data into Vector
	err = newStrictDecoder(data).Decode(&dst.Vector)
	if err == nil {
		jsonVector, _ := json.Marshal(dst.Vector)
		if string(jsonVector) == "{}" { // empty struct
			dst.Vector = nil
		} else {
			if err = validator.Validate(dst.Vector); err != nil {
				dst.Vector = nil
			} else {
				match++
			}
		}
	} else {
		dst.Vector = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.FrameOffset = nil
		dst.FrameOffsetRegion = nil
		dst.Region = nil
		dst.Vector = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PostCommentRequestClientMeta)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PostCommentRequestClientMeta)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PostCommentRequestClientMeta) MarshalJSON() ([]byte, error) {
	if src.FrameOffset != nil {
		return json.Marshal(&src.FrameOffset)
	}

	if src.FrameOffsetRegion != nil {
		return json.Marshal(&src.FrameOffsetRegion)
	}

	if src.Region != nil {
		return json.Marshal(&src.Region)
	}

	if src.Vector != nil {
		return json.Marshal(&src.Vector)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PostCommentRequestClientMeta) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.FrameOffset != nil {
		return obj.FrameOffset
	}

	if obj.FrameOffsetRegion != nil {
		return obj.FrameOffsetRegion
	}

	if obj.Region != nil {
		return obj.Region
	}

	if obj.Vector != nil {
		return obj.Vector
	}

	// all schemas are nil
	return nil
}

type NullablePostCommentRequestClientMeta struct {
	value *PostCommentRequestClientMeta
	isSet bool
}

func (v NullablePostCommentRequestClientMeta) Get() *PostCommentRequestClientMeta {
	return v.value
}

func (v *NullablePostCommentRequestClientMeta) Set(val *PostCommentRequestClientMeta) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCommentRequestClientMeta) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCommentRequestClientMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCommentRequestClientMeta(val *PostCommentRequestClientMeta) *NullablePostCommentRequestClientMeta {
	return &NullablePostCommentRequestClientMeta{value: val, isSet: true}
}

func (v NullablePostCommentRequestClientMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCommentRequestClientMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

