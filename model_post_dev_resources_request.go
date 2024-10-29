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

// checks if the PostDevResourcesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostDevResourcesRequest{}

// PostDevResourcesRequest struct for PostDevResourcesRequest
type PostDevResourcesRequest struct {
	// An array of dev resources.
	DevResources []PostDevResourcesRequestDevResourcesInner `json:"dev_resources"`
}

type _PostDevResourcesRequest PostDevResourcesRequest

// NewPostDevResourcesRequest instantiates a new PostDevResourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostDevResourcesRequest(devResources []PostDevResourcesRequestDevResourcesInner) *PostDevResourcesRequest {
	this := PostDevResourcesRequest{}
	this.DevResources = devResources
	return &this
}

// NewPostDevResourcesRequestWithDefaults instantiates a new PostDevResourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostDevResourcesRequestWithDefaults() *PostDevResourcesRequest {
	this := PostDevResourcesRequest{}
	return &this
}

// GetDevResources returns the DevResources field value
func (o *PostDevResourcesRequest) GetDevResources() []PostDevResourcesRequestDevResourcesInner {
	if o == nil {
		var ret []PostDevResourcesRequestDevResourcesInner
		return ret
	}

	return o.DevResources
}

// GetDevResourcesOk returns a tuple with the DevResources field value
// and a boolean to check if the value has been set.
func (o *PostDevResourcesRequest) GetDevResourcesOk() ([]PostDevResourcesRequestDevResourcesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.DevResources, true
}

// SetDevResources sets field value
func (o *PostDevResourcesRequest) SetDevResources(v []PostDevResourcesRequestDevResourcesInner) {
	o.DevResources = v
}

func (o PostDevResourcesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostDevResourcesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dev_resources"] = o.DevResources
	return toSerialize, nil
}

func (o *PostDevResourcesRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"dev_resources",
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

	varPostDevResourcesRequest := _PostDevResourcesRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostDevResourcesRequest)

	if err != nil {
		return err
	}

	*o = PostDevResourcesRequest(varPostDevResourcesRequest)

	return err
}

type NullablePostDevResourcesRequest struct {
	value *PostDevResourcesRequest
	isSet bool
}

func (v NullablePostDevResourcesRequest) Get() *PostDevResourcesRequest {
	return v.value
}

func (v *NullablePostDevResourcesRequest) Set(val *PostDevResourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostDevResourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostDevResourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostDevResourcesRequest(val *PostDevResourcesRequest) *NullablePostDevResourcesRequest {
	return &NullablePostDevResourcesRequest{value: val, isSet: true}
}

func (v NullablePostDevResourcesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostDevResourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

