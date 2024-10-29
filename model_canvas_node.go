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

// checks if the CanvasNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CanvasNode{}

// CanvasNode struct for CanvasNode
type CanvasNode struct {
	// A string uniquely identifying this node within the document.
	Id string `json:"id"`
	// The name given to the node by the user in the tool.
	Name string `json:"name"`
	Type string `json:"type"`
	// Whether or not the node is visible on the canvas.
	Visible *bool `json:"visible,omitempty"`
	// If true, layer is locked and cannot be edited
	Locked *bool `json:"locked,omitempty"`
	// Whether the layer is fixed while the parent is scrolling
	// Deprecated
	IsFixed *bool `json:"isFixed,omitempty"`
	// How layer should be treated when the frame is resized
	ScrollBehavior string `json:"scrollBehavior"`
	// The rotation of the node, if not 0.
	Rotation *float32 `json:"rotation,omitempty"`
	// A mapping of a layer's property to component property name of component properties attached to this node. The component property name can be used to look up more information on the corresponding component's or component set's componentPropertyDefinitions.
	ComponentPropertyReferences map[string]string `json:"componentPropertyReferences,omitempty"`
	PluginData interface{} `json:"pluginData,omitempty"`
	SharedPluginData interface{} `json:"sharedPluginData,omitempty"`
	BoundVariables *IsLayerTraitBoundVariables `json:"boundVariables,omitempty"`
	// A mapping of variable collection ID to mode ID representing the explicitly set modes for this node.
	ExplicitVariableModes map[string]string `json:"explicitVariableModes,omitempty"`
	// An array of export settings representing images to export from the node.
	ExportSettings []ExportSetting `json:"exportSettings,omitempty"`
	Children []SubcanvasNode `json:"children"`
	// Background color of the canvas.
	BackgroundColor RGBA `json:"backgroundColor"`
	// Node ID that corresponds to the start frame for prototypes. This is deprecated with the introduction of multiple flows. Please use the `flowStartingPoints` field.
	// Deprecated
	PrototypeStartNodeID NullableString `json:"prototypeStartNodeID"`
	// An array of flow starting points sorted by its position in the prototype settings panel.
	FlowStartingPoints []FlowStartingPoint `json:"flowStartingPoints"`
	// The device used to view a prototype.
	PrototypeDevice PrototypeDevice `json:"prototypeDevice"`
	Measurements []Measurement `json:"measurements,omitempty"`
}

type _CanvasNode CanvasNode

// NewCanvasNode instantiates a new CanvasNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCanvasNode(id string, name string, type_ string, scrollBehavior string, children []SubcanvasNode, backgroundColor RGBA, prototypeStartNodeID NullableString, flowStartingPoints []FlowStartingPoint, prototypeDevice PrototypeDevice) *CanvasNode {
	this := CanvasNode{}
	this.Id = id
	this.Name = name
	this.Type = type_
	var visible bool = true
	this.Visible = &visible
	var locked bool = false
	this.Locked = &locked
	var isFixed bool = false
	this.IsFixed = &isFixed
	this.ScrollBehavior = scrollBehavior
	var rotation float32 = 0
	this.Rotation = &rotation
	this.Children = children
	this.BackgroundColor = backgroundColor
	this.PrototypeStartNodeID = prototypeStartNodeID
	this.FlowStartingPoints = flowStartingPoints
	this.PrototypeDevice = prototypeDevice
	return &this
}

// NewCanvasNodeWithDefaults instantiates a new CanvasNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCanvasNodeWithDefaults() *CanvasNode {
	this := CanvasNode{}
	var visible bool = true
	this.Visible = &visible
	var locked bool = false
	this.Locked = &locked
	var isFixed bool = false
	this.IsFixed = &isFixed
	var scrollBehavior string = "SCROLLS"
	this.ScrollBehavior = scrollBehavior
	var rotation float32 = 0
	this.Rotation = &rotation
	return &this
}

// GetId returns the Id field value
func (o *CanvasNode) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CanvasNode) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *CanvasNode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CanvasNode) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *CanvasNode) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CanvasNode) SetType(v string) {
	o.Type = v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *CanvasNode) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// HasVisible returns a boolean if a field has been set.
func (o *CanvasNode) HasVisible() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *CanvasNode) SetVisible(v bool) {
	o.Visible = &v
}

// GetLocked returns the Locked field value if set, zero value otherwise.
func (o *CanvasNode) GetLocked() bool {
	if o == nil || IsNil(o.Locked) {
		var ret bool
		return ret
	}
	return *o.Locked
}

// GetLockedOk returns a tuple with the Locked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetLockedOk() (*bool, bool) {
	if o == nil || IsNil(o.Locked) {
		return nil, false
	}
	return o.Locked, true
}

// HasLocked returns a boolean if a field has been set.
func (o *CanvasNode) HasLocked() bool {
	if o != nil && !IsNil(o.Locked) {
		return true
	}

	return false
}

// SetLocked gets a reference to the given bool and assigns it to the Locked field.
func (o *CanvasNode) SetLocked(v bool) {
	o.Locked = &v
}

// GetIsFixed returns the IsFixed field value if set, zero value otherwise.
// Deprecated
func (o *CanvasNode) GetIsFixed() bool {
	if o == nil || IsNil(o.IsFixed) {
		var ret bool
		return ret
	}
	return *o.IsFixed
}

// GetIsFixedOk returns a tuple with the IsFixed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *CanvasNode) GetIsFixedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFixed) {
		return nil, false
	}
	return o.IsFixed, true
}

// HasIsFixed returns a boolean if a field has been set.
func (o *CanvasNode) HasIsFixed() bool {
	if o != nil && !IsNil(o.IsFixed) {
		return true
	}

	return false
}

// SetIsFixed gets a reference to the given bool and assigns it to the IsFixed field.
// Deprecated
func (o *CanvasNode) SetIsFixed(v bool) {
	o.IsFixed = &v
}

// GetScrollBehavior returns the ScrollBehavior field value
func (o *CanvasNode) GetScrollBehavior() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScrollBehavior
}

// GetScrollBehaviorOk returns a tuple with the ScrollBehavior field value
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetScrollBehaviorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScrollBehavior, true
}

// SetScrollBehavior sets field value
func (o *CanvasNode) SetScrollBehavior(v string) {
	o.ScrollBehavior = v
}

// GetRotation returns the Rotation field value if set, zero value otherwise.
func (o *CanvasNode) GetRotation() float32 {
	if o == nil || IsNil(o.Rotation) {
		var ret float32
		return ret
	}
	return *o.Rotation
}

// GetRotationOk returns a tuple with the Rotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetRotationOk() (*float32, bool) {
	if o == nil || IsNil(o.Rotation) {
		return nil, false
	}
	return o.Rotation, true
}

// HasRotation returns a boolean if a field has been set.
func (o *CanvasNode) HasRotation() bool {
	if o != nil && !IsNil(o.Rotation) {
		return true
	}

	return false
}

// SetRotation gets a reference to the given float32 and assigns it to the Rotation field.
func (o *CanvasNode) SetRotation(v float32) {
	o.Rotation = &v
}

// GetComponentPropertyReferences returns the ComponentPropertyReferences field value if set, zero value otherwise.
func (o *CanvasNode) GetComponentPropertyReferences() map[string]string {
	if o == nil || IsNil(o.ComponentPropertyReferences) {
		var ret map[string]string
		return ret
	}
	return o.ComponentPropertyReferences
}

// GetComponentPropertyReferencesOk returns a tuple with the ComponentPropertyReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetComponentPropertyReferencesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.ComponentPropertyReferences) {
		return map[string]string{}, false
	}
	return o.ComponentPropertyReferences, true
}

// HasComponentPropertyReferences returns a boolean if a field has been set.
func (o *CanvasNode) HasComponentPropertyReferences() bool {
	if o != nil && !IsNil(o.ComponentPropertyReferences) {
		return true
	}

	return false
}

// SetComponentPropertyReferences gets a reference to the given map[string]string and assigns it to the ComponentPropertyReferences field.
func (o *CanvasNode) SetComponentPropertyReferences(v map[string]string) {
	o.ComponentPropertyReferences = v
}

// GetPluginData returns the PluginData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CanvasNode) GetPluginData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.PluginData
}

// GetPluginDataOk returns a tuple with the PluginData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CanvasNode) GetPluginDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.PluginData) {
		return nil, false
	}
	return &o.PluginData, true
}

// HasPluginData returns a boolean if a field has been set.
func (o *CanvasNode) HasPluginData() bool {
	if o != nil && !IsNil(o.PluginData) {
		return true
	}

	return false
}

// SetPluginData gets a reference to the given interface{} and assigns it to the PluginData field.
func (o *CanvasNode) SetPluginData(v interface{}) {
	o.PluginData = v
}

// GetSharedPluginData returns the SharedPluginData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CanvasNode) GetSharedPluginData() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.SharedPluginData
}

// GetSharedPluginDataOk returns a tuple with the SharedPluginData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CanvasNode) GetSharedPluginDataOk() (*interface{}, bool) {
	if o == nil || IsNil(o.SharedPluginData) {
		return nil, false
	}
	return &o.SharedPluginData, true
}

// HasSharedPluginData returns a boolean if a field has been set.
func (o *CanvasNode) HasSharedPluginData() bool {
	if o != nil && !IsNil(o.SharedPluginData) {
		return true
	}

	return false
}

// SetSharedPluginData gets a reference to the given interface{} and assigns it to the SharedPluginData field.
func (o *CanvasNode) SetSharedPluginData(v interface{}) {
	o.SharedPluginData = v
}

// GetBoundVariables returns the BoundVariables field value if set, zero value otherwise.
func (o *CanvasNode) GetBoundVariables() IsLayerTraitBoundVariables {
	if o == nil || IsNil(o.BoundVariables) {
		var ret IsLayerTraitBoundVariables
		return ret
	}
	return *o.BoundVariables
}

// GetBoundVariablesOk returns a tuple with the BoundVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetBoundVariablesOk() (*IsLayerTraitBoundVariables, bool) {
	if o == nil || IsNil(o.BoundVariables) {
		return nil, false
	}
	return o.BoundVariables, true
}

// HasBoundVariables returns a boolean if a field has been set.
func (o *CanvasNode) HasBoundVariables() bool {
	if o != nil && !IsNil(o.BoundVariables) {
		return true
	}

	return false
}

// SetBoundVariables gets a reference to the given IsLayerTraitBoundVariables and assigns it to the BoundVariables field.
func (o *CanvasNode) SetBoundVariables(v IsLayerTraitBoundVariables) {
	o.BoundVariables = &v
}

// GetExplicitVariableModes returns the ExplicitVariableModes field value if set, zero value otherwise.
func (o *CanvasNode) GetExplicitVariableModes() map[string]string {
	if o == nil || IsNil(o.ExplicitVariableModes) {
		var ret map[string]string
		return ret
	}
	return o.ExplicitVariableModes
}

// GetExplicitVariableModesOk returns a tuple with the ExplicitVariableModes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetExplicitVariableModesOk() (map[string]string, bool) {
	if o == nil || IsNil(o.ExplicitVariableModes) {
		return map[string]string{}, false
	}
	return o.ExplicitVariableModes, true
}

// HasExplicitVariableModes returns a boolean if a field has been set.
func (o *CanvasNode) HasExplicitVariableModes() bool {
	if o != nil && !IsNil(o.ExplicitVariableModes) {
		return true
	}

	return false
}

// SetExplicitVariableModes gets a reference to the given map[string]string and assigns it to the ExplicitVariableModes field.
func (o *CanvasNode) SetExplicitVariableModes(v map[string]string) {
	o.ExplicitVariableModes = v
}

// GetExportSettings returns the ExportSettings field value if set, zero value otherwise.
func (o *CanvasNode) GetExportSettings() []ExportSetting {
	if o == nil || IsNil(o.ExportSettings) {
		var ret []ExportSetting
		return ret
	}
	return o.ExportSettings
}

// GetExportSettingsOk returns a tuple with the ExportSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetExportSettingsOk() ([]ExportSetting, bool) {
	if o == nil || IsNil(o.ExportSettings) {
		return nil, false
	}
	return o.ExportSettings, true
}

// HasExportSettings returns a boolean if a field has been set.
func (o *CanvasNode) HasExportSettings() bool {
	if o != nil && !IsNil(o.ExportSettings) {
		return true
	}

	return false
}

// SetExportSettings gets a reference to the given []ExportSetting and assigns it to the ExportSettings field.
func (o *CanvasNode) SetExportSettings(v []ExportSetting) {
	o.ExportSettings = v
}

// GetChildren returns the Children field value
func (o *CanvasNode) GetChildren() []SubcanvasNode {
	if o == nil {
		var ret []SubcanvasNode
		return ret
	}

	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetChildrenOk() ([]SubcanvasNode, bool) {
	if o == nil {
		return nil, false
	}
	return o.Children, true
}

// SetChildren sets field value
func (o *CanvasNode) SetChildren(v []SubcanvasNode) {
	o.Children = v
}

// GetBackgroundColor returns the BackgroundColor field value
func (o *CanvasNode) GetBackgroundColor() RGBA {
	if o == nil {
		var ret RGBA
		return ret
	}

	return o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetBackgroundColorOk() (*RGBA, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BackgroundColor, true
}

// SetBackgroundColor sets field value
func (o *CanvasNode) SetBackgroundColor(v RGBA) {
	o.BackgroundColor = v
}

// GetPrototypeStartNodeID returns the PrototypeStartNodeID field value
// If the value is explicit nil, the zero value for string will be returned
// Deprecated
func (o *CanvasNode) GetPrototypeStartNodeID() string {
	if o == nil || o.PrototypeStartNodeID.Get() == nil {
		var ret string
		return ret
	}

	return *o.PrototypeStartNodeID.Get()
}

// GetPrototypeStartNodeIDOk returns a tuple with the PrototypeStartNodeID field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
// Deprecated
func (o *CanvasNode) GetPrototypeStartNodeIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrototypeStartNodeID.Get(), o.PrototypeStartNodeID.IsSet()
}

// SetPrototypeStartNodeID sets field value
// Deprecated
func (o *CanvasNode) SetPrototypeStartNodeID(v string) {
	o.PrototypeStartNodeID.Set(&v)
}

// GetFlowStartingPoints returns the FlowStartingPoints field value
func (o *CanvasNode) GetFlowStartingPoints() []FlowStartingPoint {
	if o == nil {
		var ret []FlowStartingPoint
		return ret
	}

	return o.FlowStartingPoints
}

// GetFlowStartingPointsOk returns a tuple with the FlowStartingPoints field value
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetFlowStartingPointsOk() ([]FlowStartingPoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.FlowStartingPoints, true
}

// SetFlowStartingPoints sets field value
func (o *CanvasNode) SetFlowStartingPoints(v []FlowStartingPoint) {
	o.FlowStartingPoints = v
}

// GetPrototypeDevice returns the PrototypeDevice field value
func (o *CanvasNode) GetPrototypeDevice() PrototypeDevice {
	if o == nil {
		var ret PrototypeDevice
		return ret
	}

	return o.PrototypeDevice
}

// GetPrototypeDeviceOk returns a tuple with the PrototypeDevice field value
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetPrototypeDeviceOk() (*PrototypeDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrototypeDevice, true
}

// SetPrototypeDevice sets field value
func (o *CanvasNode) SetPrototypeDevice(v PrototypeDevice) {
	o.PrototypeDevice = v
}

// GetMeasurements returns the Measurements field value if set, zero value otherwise.
func (o *CanvasNode) GetMeasurements() []Measurement {
	if o == nil || IsNil(o.Measurements) {
		var ret []Measurement
		return ret
	}
	return o.Measurements
}

// GetMeasurementsOk returns a tuple with the Measurements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CanvasNode) GetMeasurementsOk() ([]Measurement, bool) {
	if o == nil || IsNil(o.Measurements) {
		return nil, false
	}
	return o.Measurements, true
}

// HasMeasurements returns a boolean if a field has been set.
func (o *CanvasNode) HasMeasurements() bool {
	if o != nil && !IsNil(o.Measurements) {
		return true
	}

	return false
}

// SetMeasurements gets a reference to the given []Measurement and assigns it to the Measurements field.
func (o *CanvasNode) SetMeasurements(v []Measurement) {
	o.Measurements = v
}

func (o CanvasNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CanvasNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["type"] = o.Type
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	if !IsNil(o.Locked) {
		toSerialize["locked"] = o.Locked
	}
	if !IsNil(o.IsFixed) {
		toSerialize["isFixed"] = o.IsFixed
	}
	toSerialize["scrollBehavior"] = o.ScrollBehavior
	if !IsNil(o.Rotation) {
		toSerialize["rotation"] = o.Rotation
	}
	if !IsNil(o.ComponentPropertyReferences) {
		toSerialize["componentPropertyReferences"] = o.ComponentPropertyReferences
	}
	if o.PluginData != nil {
		toSerialize["pluginData"] = o.PluginData
	}
	if o.SharedPluginData != nil {
		toSerialize["sharedPluginData"] = o.SharedPluginData
	}
	if !IsNil(o.BoundVariables) {
		toSerialize["boundVariables"] = o.BoundVariables
	}
	if !IsNil(o.ExplicitVariableModes) {
		toSerialize["explicitVariableModes"] = o.ExplicitVariableModes
	}
	if !IsNil(o.ExportSettings) {
		toSerialize["exportSettings"] = o.ExportSettings
	}
	toSerialize["children"] = o.Children
	toSerialize["backgroundColor"] = o.BackgroundColor
	toSerialize["prototypeStartNodeID"] = o.PrototypeStartNodeID.Get()
	toSerialize["flowStartingPoints"] = o.FlowStartingPoints
	toSerialize["prototypeDevice"] = o.PrototypeDevice
	if !IsNil(o.Measurements) {
		toSerialize["measurements"] = o.Measurements
	}
	return toSerialize, nil
}

func (o *CanvasNode) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"type",
		"scrollBehavior",
		"children",
		"backgroundColor",
		"prototypeStartNodeID",
		"flowStartingPoints",
		"prototypeDevice",
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

	varCanvasNode := _CanvasNode{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCanvasNode)

	if err != nil {
		return err
	}

	*o = CanvasNode(varCanvasNode)

	return err
}

type NullableCanvasNode struct {
	value *CanvasNode
	isSet bool
}

func (v NullableCanvasNode) Get() *CanvasNode {
	return v.value
}

func (v *NullableCanvasNode) Set(val *CanvasNode) {
	v.value = val
	v.isSet = true
}

func (v NullableCanvasNode) IsSet() bool {
	return v.isSet
}

func (v *NullableCanvasNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCanvasNode(val *CanvasNode) *NullableCanvasNode {
	return &NullableCanvasNode{value: val, isSet: true}
}

func (v NullableCanvasNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCanvasNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

