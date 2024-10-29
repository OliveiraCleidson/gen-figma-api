# SliceNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string uniquely identifying this node within the document. | 
**Name** | **string** | The name given to the node by the user in the tool. | 
**Type** | **string** | The type of this node, represented by the string literal \&quot;SLICE\&quot; | 
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

## Methods

### NewSliceNode

`func NewSliceNode(id string, name string, type_ string, scrollBehavior string, ) *SliceNode`

NewSliceNode instantiates a new SliceNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSliceNodeWithDefaults

`func NewSliceNodeWithDefaults() *SliceNode`

NewSliceNodeWithDefaults instantiates a new SliceNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SliceNode) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SliceNode) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SliceNode) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SliceNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SliceNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SliceNode) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *SliceNode) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SliceNode) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SliceNode) SetType(v string)`

SetType sets Type field to given value.


### GetVisible

`func (o *SliceNode) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *SliceNode) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *SliceNode) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *SliceNode) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetLocked

`func (o *SliceNode) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *SliceNode) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *SliceNode) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *SliceNode) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetIsFixed

`func (o *SliceNode) GetIsFixed() bool`

GetIsFixed returns the IsFixed field if non-nil, zero value otherwise.

### GetIsFixedOk

`func (o *SliceNode) GetIsFixedOk() (*bool, bool)`

GetIsFixedOk returns a tuple with the IsFixed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFixed

`func (o *SliceNode) SetIsFixed(v bool)`

SetIsFixed sets IsFixed field to given value.

### HasIsFixed

`func (o *SliceNode) HasIsFixed() bool`

HasIsFixed returns a boolean if a field has been set.

### GetScrollBehavior

`func (o *SliceNode) GetScrollBehavior() string`

GetScrollBehavior returns the ScrollBehavior field if non-nil, zero value otherwise.

### GetScrollBehaviorOk

`func (o *SliceNode) GetScrollBehaviorOk() (*string, bool)`

GetScrollBehaviorOk returns a tuple with the ScrollBehavior field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrollBehavior

`func (o *SliceNode) SetScrollBehavior(v string)`

SetScrollBehavior sets ScrollBehavior field to given value.


### GetRotation

`func (o *SliceNode) GetRotation() float32`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *SliceNode) GetRotationOk() (*float32, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *SliceNode) SetRotation(v float32)`

SetRotation sets Rotation field to given value.

### HasRotation

`func (o *SliceNode) HasRotation() bool`

HasRotation returns a boolean if a field has been set.

### GetComponentPropertyReferences

`func (o *SliceNode) GetComponentPropertyReferences() map[string]string`

GetComponentPropertyReferences returns the ComponentPropertyReferences field if non-nil, zero value otherwise.

### GetComponentPropertyReferencesOk

`func (o *SliceNode) GetComponentPropertyReferencesOk() (*map[string]string, bool)`

GetComponentPropertyReferencesOk returns a tuple with the ComponentPropertyReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPropertyReferences

`func (o *SliceNode) SetComponentPropertyReferences(v map[string]string)`

SetComponentPropertyReferences sets ComponentPropertyReferences field to given value.

### HasComponentPropertyReferences

`func (o *SliceNode) HasComponentPropertyReferences() bool`

HasComponentPropertyReferences returns a boolean if a field has been set.

### GetPluginData

`func (o *SliceNode) GetPluginData() interface{}`

GetPluginData returns the PluginData field if non-nil, zero value otherwise.

### GetPluginDataOk

`func (o *SliceNode) GetPluginDataOk() (*interface{}, bool)`

GetPluginDataOk returns a tuple with the PluginData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPluginData

`func (o *SliceNode) SetPluginData(v interface{})`

SetPluginData sets PluginData field to given value.

### HasPluginData

`func (o *SliceNode) HasPluginData() bool`

HasPluginData returns a boolean if a field has been set.

### SetPluginDataNil

`func (o *SliceNode) SetPluginDataNil(b bool)`

 SetPluginDataNil sets the value for PluginData to be an explicit nil

### UnsetPluginData
`func (o *SliceNode) UnsetPluginData()`

UnsetPluginData ensures that no value is present for PluginData, not even an explicit nil
### GetSharedPluginData

`func (o *SliceNode) GetSharedPluginData() interface{}`

GetSharedPluginData returns the SharedPluginData field if non-nil, zero value otherwise.

### GetSharedPluginDataOk

`func (o *SliceNode) GetSharedPluginDataOk() (*interface{}, bool)`

GetSharedPluginDataOk returns a tuple with the SharedPluginData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedPluginData

`func (o *SliceNode) SetSharedPluginData(v interface{})`

SetSharedPluginData sets SharedPluginData field to given value.

### HasSharedPluginData

`func (o *SliceNode) HasSharedPluginData() bool`

HasSharedPluginData returns a boolean if a field has been set.

### SetSharedPluginDataNil

`func (o *SliceNode) SetSharedPluginDataNil(b bool)`

 SetSharedPluginDataNil sets the value for SharedPluginData to be an explicit nil

### UnsetSharedPluginData
`func (o *SliceNode) UnsetSharedPluginData()`

UnsetSharedPluginData ensures that no value is present for SharedPluginData, not even an explicit nil
### GetBoundVariables

`func (o *SliceNode) GetBoundVariables() IsLayerTraitBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *SliceNode) GetBoundVariablesOk() (*IsLayerTraitBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *SliceNode) SetBoundVariables(v IsLayerTraitBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *SliceNode) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.

### GetExplicitVariableModes

`func (o *SliceNode) GetExplicitVariableModes() map[string]string`

GetExplicitVariableModes returns the ExplicitVariableModes field if non-nil, zero value otherwise.

### GetExplicitVariableModesOk

`func (o *SliceNode) GetExplicitVariableModesOk() (*map[string]string, bool)`

GetExplicitVariableModesOk returns a tuple with the ExplicitVariableModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitVariableModes

`func (o *SliceNode) SetExplicitVariableModes(v map[string]string)`

SetExplicitVariableModes sets ExplicitVariableModes field to given value.

### HasExplicitVariableModes

`func (o *SliceNode) HasExplicitVariableModes() bool`

HasExplicitVariableModes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


