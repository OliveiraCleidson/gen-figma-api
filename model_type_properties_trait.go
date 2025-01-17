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

// checks if the TypePropertiesTrait type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TypePropertiesTrait{}

// TypePropertiesTrait struct for TypePropertiesTrait
type TypePropertiesTrait struct {
	// The raw characters in the text node.
	Characters string `json:"characters"`
	// Style of text including font family and weight.
	Style TypeStyle `json:"style"`
	// The array corresponds to characters in the text box, where each element references the 'styleOverrideTable' to apply specific styles to each character. The array's length can be less than or equal to the number of characters due to the removal of trailing zeros. Elements with a value of 0 indicate characters that use the default type style. If the array is shorter than the total number of characters, the characters beyond the array's length also use the default style.
	CharacterStyleOverrides []float32 `json:"characterStyleOverrides"`
	// Internal property, preserved for backward compatibility. Avoid using this value.
	LayoutVersion *float32 `json:"layoutVersion,omitempty"`
	// Map from ID to TypeStyle for looking up style overrides.
	StyleOverrideTable map[string]TypeStyle `json:"styleOverrideTable"`
	// An array with the same number of elements as lines in the text node, where lines are delimited by newline or paragraph separator characters. Each element in the array corresponds to the list type of a specific line. List types are represented as string enums with one of these possible values:  - `NONE`: Not a list item. - `ORDERED`: Text is an ordered list (numbered). - `UNORDERED`: Text is an unordered list (bulleted).
	LineTypes []string `json:"lineTypes"`
	// An array with the same number of elements as lines in the text node, where lines are delimited by newline or paragraph separator characters. Each element in the array corresponds to the indentation level of a specific line.
	LineIndentations []float32 `json:"lineIndentations"`
}

type _TypePropertiesTrait TypePropertiesTrait

// NewTypePropertiesTrait instantiates a new TypePropertiesTrait object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTypePropertiesTrait(characters string, style TypeStyle, characterStyleOverrides []float32, styleOverrideTable map[string]TypeStyle, lineTypes []string, lineIndentations []float32) *TypePropertiesTrait {
	this := TypePropertiesTrait{}
	this.Characters = characters
	this.Style = style
	this.CharacterStyleOverrides = characterStyleOverrides
	this.StyleOverrideTable = styleOverrideTable
	this.LineTypes = lineTypes
	this.LineIndentations = lineIndentations
	return &this
}

// NewTypePropertiesTraitWithDefaults instantiates a new TypePropertiesTrait object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTypePropertiesTraitWithDefaults() *TypePropertiesTrait {
	this := TypePropertiesTrait{}
	return &this
}

// GetCharacters returns the Characters field value
func (o *TypePropertiesTrait) GetCharacters() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Characters
}

// GetCharactersOk returns a tuple with the Characters field value
// and a boolean to check if the value has been set.
func (o *TypePropertiesTrait) GetCharactersOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Characters, true
}

// SetCharacters sets field value
func (o *TypePropertiesTrait) SetCharacters(v string) {
	o.Characters = v
}

// GetStyle returns the Style field value
func (o *TypePropertiesTrait) GetStyle() TypeStyle {
	if o == nil {
		var ret TypeStyle
		return ret
	}

	return o.Style
}

// GetStyleOk returns a tuple with the Style field value
// and a boolean to check if the value has been set.
func (o *TypePropertiesTrait) GetStyleOk() (*TypeStyle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Style, true
}

// SetStyle sets field value
func (o *TypePropertiesTrait) SetStyle(v TypeStyle) {
	o.Style = v
}

// GetCharacterStyleOverrides returns the CharacterStyleOverrides field value
func (o *TypePropertiesTrait) GetCharacterStyleOverrides() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.CharacterStyleOverrides
}

// GetCharacterStyleOverridesOk returns a tuple with the CharacterStyleOverrides field value
// and a boolean to check if the value has been set.
func (o *TypePropertiesTrait) GetCharacterStyleOverridesOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CharacterStyleOverrides, true
}

// SetCharacterStyleOverrides sets field value
func (o *TypePropertiesTrait) SetCharacterStyleOverrides(v []float32) {
	o.CharacterStyleOverrides = v
}

// GetLayoutVersion returns the LayoutVersion field value if set, zero value otherwise.
func (o *TypePropertiesTrait) GetLayoutVersion() float32 {
	if o == nil || IsNil(o.LayoutVersion) {
		var ret float32
		return ret
	}
	return *o.LayoutVersion
}

// GetLayoutVersionOk returns a tuple with the LayoutVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TypePropertiesTrait) GetLayoutVersionOk() (*float32, bool) {
	if o == nil || IsNil(o.LayoutVersion) {
		return nil, false
	}
	return o.LayoutVersion, true
}

// HasLayoutVersion returns a boolean if a field has been set.
func (o *TypePropertiesTrait) HasLayoutVersion() bool {
	if o != nil && !IsNil(o.LayoutVersion) {
		return true
	}

	return false
}

// SetLayoutVersion gets a reference to the given float32 and assigns it to the LayoutVersion field.
func (o *TypePropertiesTrait) SetLayoutVersion(v float32) {
	o.LayoutVersion = &v
}

// GetStyleOverrideTable returns the StyleOverrideTable field value
func (o *TypePropertiesTrait) GetStyleOverrideTable() map[string]TypeStyle {
	if o == nil {
		var ret map[string]TypeStyle
		return ret
	}

	return o.StyleOverrideTable
}

// GetStyleOverrideTableOk returns a tuple with the StyleOverrideTable field value
// and a boolean to check if the value has been set.
func (o *TypePropertiesTrait) GetStyleOverrideTableOk() (map[string]TypeStyle, bool) {
	if o == nil {
		return map[string]TypeStyle{}, false
	}
	return o.StyleOverrideTable, true
}

// SetStyleOverrideTable sets field value
func (o *TypePropertiesTrait) SetStyleOverrideTable(v map[string]TypeStyle) {
	o.StyleOverrideTable = v
}

// GetLineTypes returns the LineTypes field value
func (o *TypePropertiesTrait) GetLineTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.LineTypes
}

// GetLineTypesOk returns a tuple with the LineTypes field value
// and a boolean to check if the value has been set.
func (o *TypePropertiesTrait) GetLineTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineTypes, true
}

// SetLineTypes sets field value
func (o *TypePropertiesTrait) SetLineTypes(v []string) {
	o.LineTypes = v
}

// GetLineIndentations returns the LineIndentations field value
func (o *TypePropertiesTrait) GetLineIndentations() []float32 {
	if o == nil {
		var ret []float32
		return ret
	}

	return o.LineIndentations
}

// GetLineIndentationsOk returns a tuple with the LineIndentations field value
// and a boolean to check if the value has been set.
func (o *TypePropertiesTrait) GetLineIndentationsOk() ([]float32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LineIndentations, true
}

// SetLineIndentations sets field value
func (o *TypePropertiesTrait) SetLineIndentations(v []float32) {
	o.LineIndentations = v
}

func (o TypePropertiesTrait) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TypePropertiesTrait) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["characters"] = o.Characters
	toSerialize["style"] = o.Style
	toSerialize["characterStyleOverrides"] = o.CharacterStyleOverrides
	if !IsNil(o.LayoutVersion) {
		toSerialize["layoutVersion"] = o.LayoutVersion
	}
	toSerialize["styleOverrideTable"] = o.StyleOverrideTable
	toSerialize["lineTypes"] = o.LineTypes
	toSerialize["lineIndentations"] = o.LineIndentations
	return toSerialize, nil
}

func (o *TypePropertiesTrait) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"characters",
		"style",
		"characterStyleOverrides",
		"styleOverrideTable",
		"lineTypes",
		"lineIndentations",
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

	varTypePropertiesTrait := _TypePropertiesTrait{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTypePropertiesTrait)

	if err != nil {
		return err
	}

	*o = TypePropertiesTrait(varTypePropertiesTrait)

	return err
}

type NullableTypePropertiesTrait struct {
	value *TypePropertiesTrait
	isSet bool
}

func (v NullableTypePropertiesTrait) Get() *TypePropertiesTrait {
	return v.value
}

func (v *NullableTypePropertiesTrait) Set(val *TypePropertiesTrait) {
	v.value = val
	v.isSet = true
}

func (v NullableTypePropertiesTrait) IsSet() bool {
	return v.isSet
}

func (v *NullableTypePropertiesTrait) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTypePropertiesTrait(val *TypePropertiesTrait) *NullableTypePropertiesTrait {
	return &NullableTypePropertiesTrait{value: val, isSet: true}
}

func (v NullableTypePropertiesTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTypePropertiesTrait) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


