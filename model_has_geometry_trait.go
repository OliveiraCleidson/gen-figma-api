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

// checks if the HasGeometryTrait type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HasGeometryTrait{}

// HasGeometryTrait struct for HasGeometryTrait
type HasGeometryTrait struct {
	// An array of fill paints applied to the node.
	Fills []Paint `json:"fills"`
	// A mapping of a StyleType to style ID (see Style) of styles present on this node. The style ID can be used to look up more information about the style in the top-level styles field.
	Styles map[string]string `json:"styles,omitempty"`
	// An array of stroke paints applied to the node.
	Strokes []Paint `json:"strokes,omitempty"`
	// The weight of strokes on the node.
	StrokeWeight *float32 `json:"strokeWeight,omitempty"`
	// Position of stroke relative to vector outline, as a string enum  - `INSIDE`: stroke drawn inside the shape boundary - `OUTSIDE`: stroke drawn outside the shape boundary - `CENTER`: stroke drawn centered along the shape boundary
	StrokeAlign *string `json:"strokeAlign,omitempty"`
	// A string enum with value of \"MITER\", \"BEVEL\", or \"ROUND\", describing how corners in vector paths are rendered.
	StrokeJoin *string `json:"strokeJoin,omitempty"`
	// An array of floating point numbers describing the pattern of dash length and gap lengths that the vector stroke will use when drawn.  For example a value of [1, 2] indicates that the stroke will be drawn with a dash of length 1 followed by a gap of length 2, repeated.
	StrokeDashes []float32 `json:"strokeDashes,omitempty"`
	// Map from ID to PaintOverride for looking up fill overrides. To see which regions are overriden, you must use the `geometry=paths` option. Each path returned may have an `overrideID` which maps to this table.
	FillOverrideTable map[string]HasGeometryTraitAllOfFillOverrideTable `json:"fillOverrideTable,omitempty"`
	// Only specified if parameter `geometry=paths` is used. An array of paths representing the object fill.
	FillGeometry []Path `json:"fillGeometry,omitempty"`
	// Only specified if parameter `geometry=paths` is used. An array of paths representing the object stroke.
	StrokeGeometry []Path `json:"strokeGeometry,omitempty"`
	// A string enum describing the end caps of vector paths.
	StrokeCap *string `json:"strokeCap,omitempty"`
	// Only valid if `strokeJoin` is \"MITER\". The corner angle, in degrees, below which `strokeJoin` will be set to \"BEVEL\" to avoid super sharp corners. By default this is 28.96 degrees.
	StrokeMiterAngle *float32 `json:"strokeMiterAngle,omitempty"`
}

type _HasGeometryTrait HasGeometryTrait

// NewHasGeometryTrait instantiates a new HasGeometryTrait object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasGeometryTrait(fills []Paint) *HasGeometryTrait {
	this := HasGeometryTrait{}
	this.Fills = fills
	var strokeWeight float32 = 1
	this.StrokeWeight = &strokeWeight
	var strokeJoin string = "MITER"
	this.StrokeJoin = &strokeJoin
	var strokeCap string = "NONE"
	this.StrokeCap = &strokeCap
	var strokeMiterAngle float32 = 28.96
	this.StrokeMiterAngle = &strokeMiterAngle
	return &this
}

// NewHasGeometryTraitWithDefaults instantiates a new HasGeometryTrait object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasGeometryTraitWithDefaults() *HasGeometryTrait {
	this := HasGeometryTrait{}
	var strokeWeight float32 = 1
	this.StrokeWeight = &strokeWeight
	var strokeJoin string = "MITER"
	this.StrokeJoin = &strokeJoin
	var strokeCap string = "NONE"
	this.StrokeCap = &strokeCap
	var strokeMiterAngle float32 = 28.96
	this.StrokeMiterAngle = &strokeMiterAngle
	return &this
}

// GetFills returns the Fills field value
func (o *HasGeometryTrait) GetFills() []Paint {
	if o == nil {
		var ret []Paint
		return ret
	}

	return o.Fills
}

// GetFillsOk returns a tuple with the Fills field value
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetFillsOk() ([]Paint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fills, true
}

// SetFills sets field value
func (o *HasGeometryTrait) SetFills(v []Paint) {
	o.Fills = v
}

// GetStyles returns the Styles field value if set, zero value otherwise.
func (o *HasGeometryTrait) GetStyles() map[string]string {
	if o == nil || IsNil(o.Styles) {
		var ret map[string]string
		return ret
	}
	return o.Styles
}

// GetStylesOk returns a tuple with the Styles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetStylesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.Styles) {
		return map[string]string{}, false
	}
	return o.Styles, true
}

// HasStyles returns a boolean if a field has been set.
func (o *HasGeometryTrait) HasStyles() bool {
	if o != nil && !IsNil(o.Styles) {
		return true
	}

	return false
}

// SetStyles gets a reference to the given map[string]string and assigns it to the Styles field.
func (o *HasGeometryTrait) SetStyles(v map[string]string) {
	o.Styles = v
}

// GetStrokes returns the Strokes field value if set, zero value otherwise.
func (o *HasGeometryTrait) GetStrokes() []Paint {
	if o == nil || IsNil(o.Strokes) {
		var ret []Paint
		return ret
	}
	return o.Strokes
}

// GetStrokesOk returns a tuple with the Strokes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetStrokesOk() ([]Paint, bool) {
	if o == nil || IsNil(o.Strokes) {
		return nil, false
	}
	return o.Strokes, true
}

// HasStrokes returns a boolean if a field has been set.
func (o *HasGeometryTrait) HasStrokes() bool {
	if o != nil && !IsNil(o.Strokes) {
		return true
	}

	return false
}

// SetStrokes gets a reference to the given []Paint and assigns it to the Strokes field.
func (o *HasGeometryTrait) SetStrokes(v []Paint) {
	o.Strokes = v
}

// GetStrokeWeight returns the StrokeWeight field value if set, zero value otherwise.
func (o *HasGeometryTrait) GetStrokeWeight() float32 {
	if o == nil || IsNil(o.StrokeWeight) {
		var ret float32
		return ret
	}
	return *o.StrokeWeight
}

// GetStrokeWeightOk returns a tuple with the StrokeWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetStrokeWeightOk() (*float32, bool) {
	if o == nil || IsNil(o.StrokeWeight) {
		return nil, false
	}
	return o.StrokeWeight, true
}

// HasStrokeWeight returns a boolean if a field has been set.
func (o *HasGeometryTrait) HasStrokeWeight() bool {
	if o != nil && !IsNil(o.StrokeWeight) {
		return true
	}

	return false
}

// SetStrokeWeight gets a reference to the given float32 and assigns it to the StrokeWeight field.
func (o *HasGeometryTrait) SetStrokeWeight(v float32) {
	o.StrokeWeight = &v
}

// GetStrokeAlign returns the StrokeAlign field value if set, zero value otherwise.
func (o *HasGeometryTrait) GetStrokeAlign() string {
	if o == nil || IsNil(o.StrokeAlign) {
		var ret string
		return ret
	}
	return *o.StrokeAlign
}

// GetStrokeAlignOk returns a tuple with the StrokeAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetStrokeAlignOk() (*string, bool) {
	if o == nil || IsNil(o.StrokeAlign) {
		return nil, false
	}
	return o.StrokeAlign, true
}

// HasStrokeAlign returns a boolean if a field has been set.
func (o *HasGeometryTrait) HasStrokeAlign() bool {
	if o != nil && !IsNil(o.StrokeAlign) {
		return true
	}

	return false
}

// SetStrokeAlign gets a reference to the given string and assigns it to the StrokeAlign field.
func (o *HasGeometryTrait) SetStrokeAlign(v string) {
	o.StrokeAlign = &v
}

// GetStrokeJoin returns the StrokeJoin field value if set, zero value otherwise.
func (o *HasGeometryTrait) GetStrokeJoin() string {
	if o == nil || IsNil(o.StrokeJoin) {
		var ret string
		return ret
	}
	return *o.StrokeJoin
}

// GetStrokeJoinOk returns a tuple with the StrokeJoin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetStrokeJoinOk() (*string, bool) {
	if o == nil || IsNil(o.StrokeJoin) {
		return nil, false
	}
	return o.StrokeJoin, true
}

// HasStrokeJoin returns a boolean if a field has been set.
func (o *HasGeometryTrait) HasStrokeJoin() bool {
	if o != nil && !IsNil(o.StrokeJoin) {
		return true
	}

	return false
}

// SetStrokeJoin gets a reference to the given string and assigns it to the StrokeJoin field.
func (o *HasGeometryTrait) SetStrokeJoin(v string) {
	o.StrokeJoin = &v
}

// GetStrokeDashes returns the StrokeDashes field value if set, zero value otherwise.
func (o *HasGeometryTrait) GetStrokeDashes() []float32 {
	if o == nil || IsNil(o.StrokeDashes) {
		var ret []float32
		return ret
	}
	return o.StrokeDashes
}

// GetStrokeDashesOk returns a tuple with the StrokeDashes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetStrokeDashesOk() ([]float32, bool) {
	if o == nil || IsNil(o.StrokeDashes) {
		return nil, false
	}
	return o.StrokeDashes, true
}

// HasStrokeDashes returns a boolean if a field has been set.
func (o *HasGeometryTrait) HasStrokeDashes() bool {
	if o != nil && !IsNil(o.StrokeDashes) {
		return true
	}

	return false
}

// SetStrokeDashes gets a reference to the given []float32 and assigns it to the StrokeDashes field.
func (o *HasGeometryTrait) SetStrokeDashes(v []float32) {
	o.StrokeDashes = v
}

// GetFillOverrideTable returns the FillOverrideTable field value if set, zero value otherwise.
func (o *HasGeometryTrait) GetFillOverrideTable() map[string]HasGeometryTraitAllOfFillOverrideTable {
	if o == nil || IsNil(o.FillOverrideTable) {
		var ret map[string]HasGeometryTraitAllOfFillOverrideTable
		return ret
	}
	return o.FillOverrideTable
}

// GetFillOverrideTableOk returns a tuple with the FillOverrideTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetFillOverrideTableOk() (map[string]HasGeometryTraitAllOfFillOverrideTable, bool) {
	if o == nil || IsNil(o.FillOverrideTable) {
		return map[string]HasGeometryTraitAllOfFillOverrideTable{}, false
	}
	return o.FillOverrideTable, true
}

// HasFillOverrideTable returns a boolean if a field has been set.
func (o *HasGeometryTrait) HasFillOverrideTable() bool {
	if o != nil && !IsNil(o.FillOverrideTable) {
		return true
	}

	return false
}

// SetFillOverrideTable gets a reference to the given map[string]HasGeometryTraitAllOfFillOverrideTable and assigns it to the FillOverrideTable field.
func (o *HasGeometryTrait) SetFillOverrideTable(v map[string]HasGeometryTraitAllOfFillOverrideTable) {
	o.FillOverrideTable = v
}

// GetFillGeometry returns the FillGeometry field value if set, zero value otherwise.
func (o *HasGeometryTrait) GetFillGeometry() []Path {
	if o == nil || IsNil(o.FillGeometry) {
		var ret []Path
		return ret
	}
	return o.FillGeometry
}

// GetFillGeometryOk returns a tuple with the FillGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetFillGeometryOk() ([]Path, bool) {
	if o == nil || IsNil(o.FillGeometry) {
		return nil, false
	}
	return o.FillGeometry, true
}

// HasFillGeometry returns a boolean if a field has been set.
func (o *HasGeometryTrait) HasFillGeometry() bool {
	if o != nil && !IsNil(o.FillGeometry) {
		return true
	}

	return false
}

// SetFillGeometry gets a reference to the given []Path and assigns it to the FillGeometry field.
func (o *HasGeometryTrait) SetFillGeometry(v []Path) {
	o.FillGeometry = v
}

// GetStrokeGeometry returns the StrokeGeometry field value if set, zero value otherwise.
func (o *HasGeometryTrait) GetStrokeGeometry() []Path {
	if o == nil || IsNil(o.StrokeGeometry) {
		var ret []Path
		return ret
	}
	return o.StrokeGeometry
}

// GetStrokeGeometryOk returns a tuple with the StrokeGeometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetStrokeGeometryOk() ([]Path, bool) {
	if o == nil || IsNil(o.StrokeGeometry) {
		return nil, false
	}
	return o.StrokeGeometry, true
}

// HasStrokeGeometry returns a boolean if a field has been set.
func (o *HasGeometryTrait) HasStrokeGeometry() bool {
	if o != nil && !IsNil(o.StrokeGeometry) {
		return true
	}

	return false
}

// SetStrokeGeometry gets a reference to the given []Path and assigns it to the StrokeGeometry field.
func (o *HasGeometryTrait) SetStrokeGeometry(v []Path) {
	o.StrokeGeometry = v
}

// GetStrokeCap returns the StrokeCap field value if set, zero value otherwise.
func (o *HasGeometryTrait) GetStrokeCap() string {
	if o == nil || IsNil(o.StrokeCap) {
		var ret string
		return ret
	}
	return *o.StrokeCap
}

// GetStrokeCapOk returns a tuple with the StrokeCap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetStrokeCapOk() (*string, bool) {
	if o == nil || IsNil(o.StrokeCap) {
		return nil, false
	}
	return o.StrokeCap, true
}

// HasStrokeCap returns a boolean if a field has been set.
func (o *HasGeometryTrait) HasStrokeCap() bool {
	if o != nil && !IsNil(o.StrokeCap) {
		return true
	}

	return false
}

// SetStrokeCap gets a reference to the given string and assigns it to the StrokeCap field.
func (o *HasGeometryTrait) SetStrokeCap(v string) {
	o.StrokeCap = &v
}

// GetStrokeMiterAngle returns the StrokeMiterAngle field value if set, zero value otherwise.
func (o *HasGeometryTrait) GetStrokeMiterAngle() float32 {
	if o == nil || IsNil(o.StrokeMiterAngle) {
		var ret float32
		return ret
	}
	return *o.StrokeMiterAngle
}

// GetStrokeMiterAngleOk returns a tuple with the StrokeMiterAngle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasGeometryTrait) GetStrokeMiterAngleOk() (*float32, bool) {
	if o == nil || IsNil(o.StrokeMiterAngle) {
		return nil, false
	}
	return o.StrokeMiterAngle, true
}

// HasStrokeMiterAngle returns a boolean if a field has been set.
func (o *HasGeometryTrait) HasStrokeMiterAngle() bool {
	if o != nil && !IsNil(o.StrokeMiterAngle) {
		return true
	}

	return false
}

// SetStrokeMiterAngle gets a reference to the given float32 and assigns it to the StrokeMiterAngle field.
func (o *HasGeometryTrait) SetStrokeMiterAngle(v float32) {
	o.StrokeMiterAngle = &v
}

func (o HasGeometryTrait) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HasGeometryTrait) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fills"] = o.Fills
	if !IsNil(o.Styles) {
		toSerialize["styles"] = o.Styles
	}
	if !IsNil(o.Strokes) {
		toSerialize["strokes"] = o.Strokes
	}
	if !IsNil(o.StrokeWeight) {
		toSerialize["strokeWeight"] = o.StrokeWeight
	}
	if !IsNil(o.StrokeAlign) {
		toSerialize["strokeAlign"] = o.StrokeAlign
	}
	if !IsNil(o.StrokeJoin) {
		toSerialize["strokeJoin"] = o.StrokeJoin
	}
	if !IsNil(o.StrokeDashes) {
		toSerialize["strokeDashes"] = o.StrokeDashes
	}
	if !IsNil(o.FillOverrideTable) {
		toSerialize["fillOverrideTable"] = o.FillOverrideTable
	}
	if !IsNil(o.FillGeometry) {
		toSerialize["fillGeometry"] = o.FillGeometry
	}
	if !IsNil(o.StrokeGeometry) {
		toSerialize["strokeGeometry"] = o.StrokeGeometry
	}
	if !IsNil(o.StrokeCap) {
		toSerialize["strokeCap"] = o.StrokeCap
	}
	if !IsNil(o.StrokeMiterAngle) {
		toSerialize["strokeMiterAngle"] = o.StrokeMiterAngle
	}
	return toSerialize, nil
}

func (o *HasGeometryTrait) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fills",
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

	varHasGeometryTrait := _HasGeometryTrait{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHasGeometryTrait)

	if err != nil {
		return err
	}

	*o = HasGeometryTrait(varHasGeometryTrait)

	return err
}

type NullableHasGeometryTrait struct {
	value *HasGeometryTrait
	isSet bool
}

func (v NullableHasGeometryTrait) Get() *HasGeometryTrait {
	return v.value
}

func (v *NullableHasGeometryTrait) Set(val *HasGeometryTrait) {
	v.value = val
	v.isSet = true
}

func (v NullableHasGeometryTrait) IsSet() bool {
	return v.isSet
}

func (v *NullableHasGeometryTrait) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasGeometryTrait(val *HasGeometryTrait) *NullableHasGeometryTrait {
	return &NullableHasGeometryTrait{value: val, isSet: true}
}

func (v NullableHasGeometryTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasGeometryTrait) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


