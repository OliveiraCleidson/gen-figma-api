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

// checks if the VariableCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VariableCreate{}

// VariableCreate An object that contains details about creating a `Variable`.
type VariableCreate struct {
	// The action to perform for the variable.
	Action string `json:"action"`
	// A temporary id for this variable.
	Id *string `json:"id,omitempty"`
	// The name of this variable.
	Name string `json:"name"`
	// The variable collection that will contain the variable. You can use the temporary id of a variable collection.
	VariableCollectionId string `json:"variableCollectionId"`
	// The resolved type of the variable.
	ResolvedType string `json:"resolvedType"`
	// The description of this variable.
	Description *string `json:"description,omitempty"`
	// Whether this variable is hidden when publishing the current file as a library.
	HiddenFromPublishing *bool `json:"hiddenFromPublishing,omitempty"`
	// An array of scopes in the UI where this variable is shown. Setting this property will show/hide this variable in the variable picker UI for different fields.
	Scopes []VariableScope `json:"scopes,omitempty"`
	CodeSyntax *VariableCodeSyntax `json:"codeSyntax,omitempty"`
}

type _VariableCreate VariableCreate

// NewVariableCreate instantiates a new VariableCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVariableCreate(action string, name string, variableCollectionId string, resolvedType string) *VariableCreate {
	this := VariableCreate{}
	this.Action = action
	this.Name = name
	this.VariableCollectionId = variableCollectionId
	this.ResolvedType = resolvedType
	var hiddenFromPublishing bool = false
	this.HiddenFromPublishing = &hiddenFromPublishing
	return &this
}

// NewVariableCreateWithDefaults instantiates a new VariableCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVariableCreateWithDefaults() *VariableCreate {
	this := VariableCreate{}
	var hiddenFromPublishing bool = false
	this.HiddenFromPublishing = &hiddenFromPublishing
	return &this
}

// GetAction returns the Action field value
func (o *VariableCreate) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *VariableCreate) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *VariableCreate) SetAction(v string) {
	o.Action = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *VariableCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *VariableCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *VariableCreate) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *VariableCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VariableCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VariableCreate) SetName(v string) {
	o.Name = v
}

// GetVariableCollectionId returns the VariableCollectionId field value
func (o *VariableCreate) GetVariableCollectionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VariableCollectionId
}

// GetVariableCollectionIdOk returns a tuple with the VariableCollectionId field value
// and a boolean to check if the value has been set.
func (o *VariableCreate) GetVariableCollectionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VariableCollectionId, true
}

// SetVariableCollectionId sets field value
func (o *VariableCreate) SetVariableCollectionId(v string) {
	o.VariableCollectionId = v
}

// GetResolvedType returns the ResolvedType field value
func (o *VariableCreate) GetResolvedType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResolvedType
}

// GetResolvedTypeOk returns a tuple with the ResolvedType field value
// and a boolean to check if the value has been set.
func (o *VariableCreate) GetResolvedTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResolvedType, true
}

// SetResolvedType sets field value
func (o *VariableCreate) SetResolvedType(v string) {
	o.ResolvedType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VariableCreate) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableCreate) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VariableCreate) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VariableCreate) SetDescription(v string) {
	o.Description = &v
}

// GetHiddenFromPublishing returns the HiddenFromPublishing field value if set, zero value otherwise.
func (o *VariableCreate) GetHiddenFromPublishing() bool {
	if o == nil || IsNil(o.HiddenFromPublishing) {
		var ret bool
		return ret
	}
	return *o.HiddenFromPublishing
}

// GetHiddenFromPublishingOk returns a tuple with the HiddenFromPublishing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableCreate) GetHiddenFromPublishingOk() (*bool, bool) {
	if o == nil || IsNil(o.HiddenFromPublishing) {
		return nil, false
	}
	return o.HiddenFromPublishing, true
}

// HasHiddenFromPublishing returns a boolean if a field has been set.
func (o *VariableCreate) HasHiddenFromPublishing() bool {
	if o != nil && !IsNil(o.HiddenFromPublishing) {
		return true
	}

	return false
}

// SetHiddenFromPublishing gets a reference to the given bool and assigns it to the HiddenFromPublishing field.
func (o *VariableCreate) SetHiddenFromPublishing(v bool) {
	o.HiddenFromPublishing = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *VariableCreate) GetScopes() []VariableScope {
	if o == nil || IsNil(o.Scopes) {
		var ret []VariableScope
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableCreate) GetScopesOk() ([]VariableScope, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *VariableCreate) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []VariableScope and assigns it to the Scopes field.
func (o *VariableCreate) SetScopes(v []VariableScope) {
	o.Scopes = v
}

// GetCodeSyntax returns the CodeSyntax field value if set, zero value otherwise.
func (o *VariableCreate) GetCodeSyntax() VariableCodeSyntax {
	if o == nil || IsNil(o.CodeSyntax) {
		var ret VariableCodeSyntax
		return ret
	}
	return *o.CodeSyntax
}

// GetCodeSyntaxOk returns a tuple with the CodeSyntax field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VariableCreate) GetCodeSyntaxOk() (*VariableCodeSyntax, bool) {
	if o == nil || IsNil(o.CodeSyntax) {
		return nil, false
	}
	return o.CodeSyntax, true
}

// HasCodeSyntax returns a boolean if a field has been set.
func (o *VariableCreate) HasCodeSyntax() bool {
	if o != nil && !IsNil(o.CodeSyntax) {
		return true
	}

	return false
}

// SetCodeSyntax gets a reference to the given VariableCodeSyntax and assigns it to the CodeSyntax field.
func (o *VariableCreate) SetCodeSyntax(v VariableCodeSyntax) {
	o.CodeSyntax = &v
}

func (o VariableCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VariableCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	toSerialize["variableCollectionId"] = o.VariableCollectionId
	toSerialize["resolvedType"] = o.ResolvedType
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.HiddenFromPublishing) {
		toSerialize["hiddenFromPublishing"] = o.HiddenFromPublishing
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	if !IsNil(o.CodeSyntax) {
		toSerialize["codeSyntax"] = o.CodeSyntax
	}
	return toSerialize, nil
}

func (o *VariableCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
		"name",
		"variableCollectionId",
		"resolvedType",
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

	varVariableCreate := _VariableCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varVariableCreate)

	if err != nil {
		return err
	}

	*o = VariableCreate(varVariableCreate)

	return err
}

type NullableVariableCreate struct {
	value *VariableCreate
	isSet bool
}

func (v NullableVariableCreate) Get() *VariableCreate {
	return v.value
}

func (v *NullableVariableCreate) Set(val *VariableCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableVariableCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableVariableCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariableCreate(val *VariableCreate) *NullableVariableCreate {
	return &NullableVariableCreate{value: val, isSet: true}
}

func (v NullableVariableCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariableCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

