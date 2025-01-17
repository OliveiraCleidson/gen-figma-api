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

// checks if the VariableCollectionCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableCollectionCreate{}

// VariableCollectionCreate An object that contains details about creating a `VariableCollection`.
type VariableCollectionCreate struct {
	// The action to perform for the variable collection.
	Action string `json:"action"`
	// A temporary id for this variable collection.
	Id *string `json:"id,omitempty"`
	// The name of this variable collection.
	Name string `json:"name"`
	// The initial mode refers to the mode that is created by default. You can set a temporary id here, in order to reference this mode later in this request.
	InitialModeId *string `json:"initialModeId,omitempty"`
	// Whether this variable collection is hidden when publishing the current file as a library.
	HiddenFromPublishing *bool `json:"hiddenFromPublishing,omitempty"`
}

type _VariableCollectionCreate VariableCollectionCreate

// NewVariableCollectionCreate instantiates a new VariableCollectionCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableCollectionCreate(action string, name string) *VariableCollectionCreate {
	this := VariableCollectionCreate{}
	this.Action = action
	this.Name = name
	var hiddenFromPublishing bool = false
	this.HiddenFromPublishing = &hiddenFromPublishing
	return &this
}

// NewVariableCollectionCreateWithDefaults instantiates a new VariableCollectionCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableCollectionCreateWithDefaults() *VariableCollectionCreate {
	this := VariableCollectionCreate{}
	var hiddenFromPublishing bool = false
	this.HiddenFromPublishing = &hiddenFromPublishing
	return &this
}

// GetAction returns the Action field value
func (o *VariableCollectionCreate) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *VariableCollectionCreate) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *VariableCollectionCreate) SetAction(v string) {
	o.Action = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VariableCollectionCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableCollectionCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VariableCollectionCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VariableCollectionCreate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *VariableCollectionCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VariableCollectionCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VariableCollectionCreate) SetName(v string) {
	o.Name = v
}

// GetInitialModeId returns the InitialModeId field value if set, zero value otherwise.
func (o *VariableCollectionCreate) GetInitialModeId() string {
	if o == nil || IsNil(o.InitialModeId) {
		var ret string
		return ret
	}
	return *o.InitialModeId
}

// GetInitialModeIdOk returns a tuple with the InitialModeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableCollectionCreate) GetInitialModeIdOk() (*string, bool) {
	if o == nil || IsNil(o.InitialModeId) {
		return nil, false
	}
	return o.InitialModeId, true
}

// HasInitialModeId returns a boolean if a field has been set.
func (o *VariableCollectionCreate) HasInitialModeId() bool {
	if o != nil && !IsNil(o.InitialModeId) {
		return true
	}

	return false
}

// SetInitialModeId gets a reference to the given string and assigns it to the InitialModeId field.
func (o *VariableCollectionCreate) SetInitialModeId(v string) {
	o.InitialModeId = &v
}

// GetHiddenFromPublishing returns the HiddenFromPublishing field value if set, zero value otherwise.
func (o *VariableCollectionCreate) GetHiddenFromPublishing() bool {
	if o == nil || IsNil(o.HiddenFromPublishing) {
		var ret bool
		return ret
	}
	return *o.HiddenFromPublishing
}

// GetHiddenFromPublishingOk returns a tuple with the HiddenFromPublishing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableCollectionCreate) GetHiddenFromPublishingOk() (*bool, bool) {
	if o == nil || IsNil(o.HiddenFromPublishing) {
		return nil, false
	}
	return o.HiddenFromPublishing, true
}

// HasHiddenFromPublishing returns a boolean if a field has been set.
func (o *VariableCollectionCreate) HasHiddenFromPublishing() bool {
	if o != nil && !IsNil(o.HiddenFromPublishing) {
		return true
	}

	return false
}

// SetHiddenFromPublishing gets a reference to the given bool and assigns it to the HiddenFromPublishing field.
func (o *VariableCollectionCreate) SetHiddenFromPublishing(v bool) {
	o.HiddenFromPublishing = &v
}

func (o VariableCollectionCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableCollectionCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.InitialModeId) {
		toSerialize["initialModeId"] = o.InitialModeId
	}
	if !IsNil(o.HiddenFromPublishing) {
		toSerialize["hiddenFromPublishing"] = o.HiddenFromPublishing
	}
	return toSerialize, nil
}

func (o *VariableCollectionCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"name",
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

	varVariableCollectionCreate := _VariableCollectionCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariableCollectionCreate)

	if err != nil {
		return err
	}

	*o = VariableCollectionCreate(varVariableCollectionCreate)

	return err
}

type NullableVariableCollectionCreate struct {
	value *VariableCollectionCreate
	isSet bool
}

func (v NullableVariableCollectionCreate) Get() *VariableCollectionCreate {
	return v.value
}

func (v *NullableVariableCollectionCreate) Set(val *VariableCollectionCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableCollectionCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableCollectionCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableCollectionCreate(val *VariableCollectionCreate) *NullableVariableCollectionCreate {
	return &NullableVariableCollectionCreate{value: val, isSet: true}
}

func (v NullableVariableCollectionCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableCollectionCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


