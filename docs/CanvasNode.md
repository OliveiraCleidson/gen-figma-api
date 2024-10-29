# CanvasNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string uniquely identifying this node within the document. | 
**Name** | **string** | The name given to the node by the user in the tool. | 
**Type** | **string** |  | 
**Visible** | Pointer to **bool** | Whether or not the node is visible on the canvas. | [optional] [default to true]
**Locked** | Pointer to **bool** | If true, layer is locked and cannot be edited | [optional] [default to false]
**IsFixed** | Pointer to **bool** | Whether the layer is fixed while the parent is scrolling | [optional] [default to false]
**ScrollBehavior** | **string** | How layer should be treated when the frame is resized | [default to "SCROLLS"]
**Rotation** | Pointer to **float32** | The rotation of the node, if not 0. | [optional] [default to 0]
**ComponentPropertyReferences** | Pointer to **map[string]string** | A mapping of a layer&#39;s property to component property name of component properties attached to this node. The component property name can be used to look up more information on the corresponding component&#39;s or component set&#39;s componentPropertyDefinitions. | [optional] 
**PluginData** | Pointer to **interface{}** |  | [optional] 
**SharedPluginData** | Pointer to **interface{}** |  | [optional] 
**BoundVariables** | Pointer to [**IsLayerTraitBoundVariables**](IsLayerTraitBoundVariables.md) |  | [optional] 
**ExplicitVariableModes** | Pointer to **map[string]string** | A mapping of variable collection ID to mode ID representing the explicitly set modes for this node. | [optional] 
**ExportSettings** | Pointer to [**[]ExportSetting**](ExportSetting.md) | An array of export settings representing images to export from the node. | [optional] 
**Children** | [**[]SubcanvasNode**](SubcanvasNode.md) |  | 
**BackgroundColor** | [**RGBA**](RGBA.md) | Background color of the canvas. | 
**PrototypeStartNodeID** | **NullableString** | Node ID that corresponds to the start frame for prototypes. This is deprecated with the introduction of multiple flows. Please use the &#x60;flowStartingPoints&#x60; field. | 
**FlowStartingPoints** | [**[]FlowStartingPoint**](FlowStartingPoint.md) | An array of flow starting points sorted by its position in the prototype settings panel. | 
**PrototypeDevice** | [**PrototypeDevice**](PrototypeDevice.md) | The device used to view a prototype. | 
**Measurements** | Pointer to [**[]Measurement**](Measurement.md) |  | [optional] 

## Methods

### NewCanvasNode

`func NewCanvasNode(id string, name string, type_ string, scrollBehavior string, children []SubcanvasNode, backgroundColor RGBA, prototypeStartNodeID NullableString, flowStartingPoints []FlowStartingPoint, prototypeDevice PrototypeDevice, ) *CanvasNode`

NewCanvasNode instantiates a new CanvasNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCanvasNodeWithDefaults

`func NewCanvasNodeWithDefaults() *CanvasNode`

NewCanvasNodeWithDefaults instantiates a new CanvasNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CanvasNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CanvasNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CanvasNode) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CanvasNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CanvasNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CanvasNode) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CanvasNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CanvasNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CanvasNode) SetType(v string)`

SetType sets Type field to given value.


### GetVisible

`func (o *CanvasNode) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *CanvasNode) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *CanvasNode) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *CanvasNode) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetLocked

`func (o *CanvasNode) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *CanvasNode) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *CanvasNode) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *CanvasNode) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetIsFixed

`func (o *CanvasNode) GetIsFixed() bool`

GetIsFixed returns the IsFixed field if non-nil, zero value otherwise.

### GetIsFixedOk

`func (o *CanvasNode) GetIsFixedOk() (*bool, bool)`

GetIsFixedOk returns a tuple with the IsFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFixed

`func (o *CanvasNode) SetIsFixed(v bool)`

SetIsFixed sets IsFixed field to given value.

### HasIsFixed

`func (o *CanvasNode) HasIsFixed() bool`

HasIsFixed returns a boolean if a field has been set.

### GetScrollBehavior

`func (o *CanvasNode) GetScrollBehavior() string`

GetScrollBehavior returns the ScrollBehavior field if non-nil, zero value otherwise.

### GetScrollBehaviorOk

`func (o *CanvasNode) GetScrollBehaviorOk() (*string, bool)`

GetScrollBehaviorOk returns a tuple with the ScrollBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollBehavior

`func (o *CanvasNode) SetScrollBehavior(v string)`

SetScrollBehavior sets ScrollBehavior field to given value.


### GetRotation

`func (o *CanvasNode) GetRotation() float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *CanvasNode) GetRotationOk() (*float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *CanvasNode) SetRotation(v float32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *CanvasNode) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetComponentPropertyReferences

`func (o *CanvasNode) GetComponentPropertyReferences() map[string]string`

GetComponentPropertyReferences returns the ComponentPropertyReferences field if non-nil, zero value otherwise.

### GetComponentPropertyReferencesOk

`func (o *CanvasNode) GetComponentPropertyReferencesOk() (*map[string]string, bool)`

GetComponentPropertyReferencesOk returns a tuple with the ComponentPropertyReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPropertyReferences

`func (o *CanvasNode) SetComponentPropertyReferences(v map[string]string)`

SetComponentPropertyReferences sets ComponentPropertyReferences field to given value.

### HasComponentPropertyReferences

`func (o *CanvasNode) HasComponentPropertyReferences() bool`

HasComponentPropertyReferences returns a boolean if a field has been set.

### GetPluginData

`func (o *CanvasNode) GetPluginData() interface{}`

GetPluginData returns the PluginData field if non-nil, zero value otherwise.

### GetPluginDataOk

`func (o *CanvasNode) GetPluginDataOk() (*interface{}, bool)`

GetPluginDataOk returns a tuple with the PluginData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginData

`func (o *CanvasNode) SetPluginData(v interface{})`

SetPluginData sets PluginData field to given value.

### HasPluginData

`func (o *CanvasNode) HasPluginData() bool`

HasPluginData returns a boolean if a field has been set.

### SetPluginDataNil

`func (o *CanvasNode) SetPluginDataNil(b bool)`

 SetPluginDataNil sets the value for PluginData to be an explicit nil

### UnsetPluginData
`func (o *CanvasNode) UnsetPluginData()`

UnsetPluginData ensures that no value is present for PluginData, not even an explicit nil
### GetSharedPluginData

`func (o *CanvasNode) GetSharedPluginData() interface{}`

GetSharedPluginData returns the SharedPluginData field if non-nil, zero value otherwise.

### GetSharedPluginDataOk

`func (o *CanvasNode) GetSharedPluginDataOk() (*interface{}, bool)`

GetSharedPluginDataOk returns a tuple with the SharedPluginData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPluginData

`func (o *CanvasNode) SetSharedPluginData(v interface{})`

SetSharedPluginData sets SharedPluginData field to given value.

### HasSharedPluginData

`func (o *CanvasNode) HasSharedPluginData() bool`

HasSharedPluginData returns a boolean if a field has been set.

### SetSharedPluginDataNil

`func (o *CanvasNode) SetSharedPluginDataNil(b bool)`

 SetSharedPluginDataNil sets the value for SharedPluginData to be an explicit nil

### UnsetSharedPluginData
`func (o *CanvasNode) UnsetSharedPluginData()`

UnsetSharedPluginData ensures that no value is present for SharedPluginData, not even an explicit nil
### GetBoundVariables

`func (o *CanvasNode) GetBoundVariables() IsLayerTraitBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *CanvasNode) GetBoundVariablesOk() (*IsLayerTraitBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *CanvasNode) SetBoundVariables(v IsLayerTraitBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *CanvasNode) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.

### GetExplicitVariableModes

`func (o *CanvasNode) GetExplicitVariableModes() map[string]string`

GetExplicitVariableModes returns the ExplicitVariableModes field if non-nil, zero value otherwise.

### GetExplicitVariableModesOk

`func (o *CanvasNode) GetExplicitVariableModesOk() (*map[string]string, bool)`

GetExplicitVariableModesOk returns a tuple with the ExplicitVariableModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitVariableModes

`func (o *CanvasNode) SetExplicitVariableModes(v map[string]string)`

SetExplicitVariableModes sets ExplicitVariableModes field to given value.

### HasExplicitVariableModes

`func (o *CanvasNode) HasExplicitVariableModes() bool`

HasExplicitVariableModes returns a boolean if a field has been set.

### GetExportSettings

`func (o *CanvasNode) GetExportSettings() []ExportSetting`

GetExportSettings returns the ExportSettings field if non-nil, zero value otherwise.

### GetExportSettingsOk

`func (o *CanvasNode) GetExportSettingsOk() (*[]ExportSetting, bool)`

GetExportSettingsOk returns a tuple with the ExportSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportSettings

`func (o *CanvasNode) SetExportSettings(v []ExportSetting)`

SetExportSettings sets ExportSettings field to given value.

### HasExportSettings

`func (o *CanvasNode) HasExportSettings() bool`

HasExportSettings returns a boolean if a field has been set.

### GetChildren

`func (o *CanvasNode) GetChildren() []SubcanvasNode`

GetChildren returns the Children field if non-nil, zero value otherwise.

### GetChildrenOk

`func (o *CanvasNode) GetChildrenOk() (*[]SubcanvasNode, bool)`

GetChildrenOk returns a tuple with the Children field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildren

`func (o *CanvasNode) SetChildren(v []SubcanvasNode)`

SetChildren sets Children field to given value.


### GetBackgroundColor

`func (o *CanvasNode) GetBackgroundColor() RGBA`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *CanvasNode) GetBackgroundColorOk() (*RGBA, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *CanvasNode) SetBackgroundColor(v RGBA)`

SetBackgroundColor sets BackgroundColor field to given value.


### GetPrototypeStartNodeID

`func (o *CanvasNode) GetPrototypeStartNodeID() string`

GetPrototypeStartNodeID returns the PrototypeStartNodeID field if non-nil, zero value otherwise.

### GetPrototypeStartNodeIDOk

`func (o *CanvasNode) GetPrototypeStartNodeIDOk() (*string, bool)`

GetPrototypeStartNodeIDOk returns a tuple with the PrototypeStartNodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrototypeStartNodeID

`func (o *CanvasNode) SetPrototypeStartNodeID(v string)`

SetPrototypeStartNodeID sets PrototypeStartNodeID field to given value.


### SetPrototypeStartNodeIDNil

`func (o *CanvasNode) SetPrototypeStartNodeIDNil(b bool)`

 SetPrototypeStartNodeIDNil sets the value for PrototypeStartNodeID to be an explicit nil

### UnsetPrototypeStartNodeID
`func (o *CanvasNode) UnsetPrototypeStartNodeID()`

UnsetPrototypeStartNodeID ensures that no value is present for PrototypeStartNodeID, not even an explicit nil
### GetFlowStartingPoints

`func (o *CanvasNode) GetFlowStartingPoints() []FlowStartingPoint`

GetFlowStartingPoints returns the FlowStartingPoints field if non-nil, zero value otherwise.

### GetFlowStartingPointsOk

`func (o *CanvasNode) GetFlowStartingPointsOk() (*[]FlowStartingPoint, bool)`

GetFlowStartingPointsOk returns a tuple with the FlowStartingPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowStartingPoints

`func (o *CanvasNode) SetFlowStartingPoints(v []FlowStartingPoint)`

SetFlowStartingPoints sets FlowStartingPoints field to given value.


### GetPrototypeDevice

`func (o *CanvasNode) GetPrototypeDevice() PrototypeDevice`

GetPrototypeDevice returns the PrototypeDevice field if non-nil, zero value otherwise.

### GetPrototypeDeviceOk

`func (o *CanvasNode) GetPrototypeDeviceOk() (*PrototypeDevice, bool)`

GetPrototypeDeviceOk returns a tuple with the PrototypeDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrototypeDevice

`func (o *CanvasNode) SetPrototypeDevice(v PrototypeDevice)`

SetPrototypeDevice sets PrototypeDevice field to given value.


### GetMeasurements

`func (o *CanvasNode) GetMeasurements() []Measurement`

GetMeasurements returns the Measurements field if non-nil, zero value otherwise.

### GetMeasurementsOk

`func (o *CanvasNode) GetMeasurementsOk() (*[]Measurement, bool)`

GetMeasurementsOk returns a tuple with the Measurements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeasurements

`func (o *CanvasNode) SetMeasurements(v []Measurement)`

SetMeasurements sets Measurements field to given value.

### HasMeasurements

`func (o *CanvasNode) HasMeasurements() bool`

HasMeasurements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


