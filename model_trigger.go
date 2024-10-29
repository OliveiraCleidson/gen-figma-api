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

// Trigger - The `\"ON_HOVER\"` and `\"ON_PRESS\"` trigger types revert the navigation when the trigger is finished (the result is temporary).  `\"MOUSE_ENTER\"`, `\"MOUSE_LEAVE\"`, `\"MOUSE_UP\"` and `\"MOUSE_DOWN\"` are permanent, one-way navigation. The `delay` parameter requires the trigger to be held for a certain duration of time before the action occurs. Both `timeout` and `delay` values are in milliseconds. The `\"ON_MEDIA_HIT\"` and `\"ON_MEDIA_END\"` trigger types can only trigger from a video.  They fire when a video reaches a certain time or ends. The `timestamp` value is in seconds.
type Trigger struct {
	AfterTimeoutTrigger *AfterTimeoutTrigger
	OnKeyDownTrigger *OnKeyDownTrigger
	OnMediaHitTrigger *OnMediaHitTrigger
	TriggerOneOf *TriggerOneOf
	TriggerOneOf1 *TriggerOneOf1
	TriggerOneOf2 *TriggerOneOf2
}

// AfterTimeoutTriggerAsTrigger is a convenience function that returns AfterTimeoutTrigger wrapped in Trigger
func AfterTimeoutTriggerAsTrigger(v *AfterTimeoutTrigger) Trigger {
	return Trigger{
		AfterTimeoutTrigger: v,
	}
}

// OnKeyDownTriggerAsTrigger is a convenience function that returns OnKeyDownTrigger wrapped in Trigger
func OnKeyDownTriggerAsTrigger(v *OnKeyDownTrigger) Trigger {
	return Trigger{
		OnKeyDownTrigger: v,
	}
}

// OnMediaHitTriggerAsTrigger is a convenience function that returns OnMediaHitTrigger wrapped in Trigger
func OnMediaHitTriggerAsTrigger(v *OnMediaHitTrigger) Trigger {
	return Trigger{
		OnMediaHitTrigger: v,
	}
}

// TriggerOneOfAsTrigger is a convenience function that returns TriggerOneOf wrapped in Trigger
func TriggerOneOfAsTrigger(v *TriggerOneOf) Trigger {
	return Trigger{
		TriggerOneOf: v,
	}
}

// TriggerOneOf1AsTrigger is a convenience function that returns TriggerOneOf1 wrapped in Trigger
func TriggerOneOf1AsTrigger(v *TriggerOneOf1) Trigger {
	return Trigger{
		TriggerOneOf1: v,
	}
}

// TriggerOneOf2AsTrigger is a convenience function that returns TriggerOneOf2 wrapped in Trigger
func TriggerOneOf2AsTrigger(v *TriggerOneOf2) Trigger {
	return Trigger{
		TriggerOneOf2: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *Trigger) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AfterTimeoutTrigger
	err = newStrictDecoder(data).Decode(&dst.AfterTimeoutTrigger)
	if err == nil {
		jsonAfterTimeoutTrigger, _ := json.Marshal(dst.AfterTimeoutTrigger)
		if string(jsonAfterTimeoutTrigger) == "{}" { // empty struct
			dst.AfterTimeoutTrigger = nil
		} else {
			if err = validator.Validate(dst.AfterTimeoutTrigger); err != nil {
				dst.AfterTimeoutTrigger = nil
			} else {
				match++
			}
		}
	} else {
		dst.AfterTimeoutTrigger = nil
	}

	// try to unmarshal data into OnKeyDownTrigger
	err = newStrictDecoder(data).Decode(&dst.OnKeyDownTrigger)
	if err == nil {
		jsonOnKeyDownTrigger, _ := json.Marshal(dst.OnKeyDownTrigger)
		if string(jsonOnKeyDownTrigger) == "{}" { // empty struct
			dst.OnKeyDownTrigger = nil
		} else {
			if err = validator.Validate(dst.OnKeyDownTrigger); err != nil {
				dst.OnKeyDownTrigger = nil
			} else {
				match++
			}
		}
	} else {
		dst.OnKeyDownTrigger = nil
	}

	// try to unmarshal data into OnMediaHitTrigger
	err = newStrictDecoder(data).Decode(&dst.OnMediaHitTrigger)
	if err == nil {
		jsonOnMediaHitTrigger, _ := json.Marshal(dst.OnMediaHitTrigger)
		if string(jsonOnMediaHitTrigger) == "{}" { // empty struct
			dst.OnMediaHitTrigger = nil
		} else {
			if err = validator.Validate(dst.OnMediaHitTrigger); err != nil {
				dst.OnMediaHitTrigger = nil
			} else {
				match++
			}
		}
	} else {
		dst.OnMediaHitTrigger = nil
	}

	// try to unmarshal data into TriggerOneOf
	err = newStrictDecoder(data).Decode(&dst.TriggerOneOf)
	if err == nil {
		jsonTriggerOneOf, _ := json.Marshal(dst.TriggerOneOf)
		if string(jsonTriggerOneOf) == "{}" { // empty struct
			dst.TriggerOneOf = nil
		} else {
			if err = validator.Validate(dst.TriggerOneOf); err != nil {
				dst.TriggerOneOf = nil
			} else {
				match++
			}
		}
	} else {
		dst.TriggerOneOf = nil
	}

	// try to unmarshal data into TriggerOneOf1
	err = newStrictDecoder(data).Decode(&dst.TriggerOneOf1)
	if err == nil {
		jsonTriggerOneOf1, _ := json.Marshal(dst.TriggerOneOf1)
		if string(jsonTriggerOneOf1) == "{}" { // empty struct
			dst.TriggerOneOf1 = nil
		} else {
			if err = validator.Validate(dst.TriggerOneOf1); err != nil {
				dst.TriggerOneOf1 = nil
			} else {
				match++
			}
		}
	} else {
		dst.TriggerOneOf1 = nil
	}

	// try to unmarshal data into TriggerOneOf2
	err = newStrictDecoder(data).Decode(&dst.TriggerOneOf2)
	if err == nil {
		jsonTriggerOneOf2, _ := json.Marshal(dst.TriggerOneOf2)
		if string(jsonTriggerOneOf2) == "{}" { // empty struct
			dst.TriggerOneOf2 = nil
		} else {
			if err = validator.Validate(dst.TriggerOneOf2); err != nil {
				dst.TriggerOneOf2 = nil
			} else {
				match++
			}
		}
	} else {
		dst.TriggerOneOf2 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AfterTimeoutTrigger = nil
		dst.OnKeyDownTrigger = nil
		dst.OnMediaHitTrigger = nil
		dst.TriggerOneOf = nil
		dst.TriggerOneOf1 = nil
		dst.TriggerOneOf2 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(Trigger)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Trigger)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Trigger) MarshalJSON() ([]byte, error) {
	if src.AfterTimeoutTrigger != nil {
		return json.Marshal(&src.AfterTimeoutTrigger)
	}

	if src.OnKeyDownTrigger != nil {
		return json.Marshal(&src.OnKeyDownTrigger)
	}

	if src.OnMediaHitTrigger != nil {
		return json.Marshal(&src.OnMediaHitTrigger)
	}

	if src.TriggerOneOf != nil {
		return json.Marshal(&src.TriggerOneOf)
	}

	if src.TriggerOneOf1 != nil {
		return json.Marshal(&src.TriggerOneOf1)
	}

	if src.TriggerOneOf2 != nil {
		return json.Marshal(&src.TriggerOneOf2)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Trigger) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AfterTimeoutTrigger != nil {
		return obj.AfterTimeoutTrigger
	}

	if obj.OnKeyDownTrigger != nil {
		return obj.OnKeyDownTrigger
	}

	if obj.OnMediaHitTrigger != nil {
		return obj.OnMediaHitTrigger
	}

	if obj.TriggerOneOf != nil {
		return obj.TriggerOneOf
	}

	if obj.TriggerOneOf1 != nil {
		return obj.TriggerOneOf1
	}

	if obj.TriggerOneOf2 != nil {
		return obj.TriggerOneOf2
	}

	// all schemas are nil
	return nil
}

type NullableTrigger struct {
	value *Trigger
	isSet bool
}

func (v NullableTrigger) Get() *Trigger {
	return v.value
}

func (v *NullableTrigger) Set(val *Trigger) {
	v.value = val
	v.isSet = true
}

func (v NullableTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrigger(val *Trigger) *NullableTrigger {
	return &NullableTrigger{value: val, isSet: true}
}

func (v NullableTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

