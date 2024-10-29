# PostVariablesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VariableCollections** | Pointer to [**[]VariableCollectionChange**](VariableCollectionChange.md) | For creating, updating, and deleting variable collections. | [optional] 
**VariableModes** | Pointer to [**[]VariableModeChange**](VariableModeChange.md) | For creating, updating, and deleting modes within variable collections. | [optional] 
**Variables** | Pointer to [**[]VariableChange**](VariableChange.md) | For creating, updating, and deleting variables. | [optional] 
**VariableModeValues** | Pointer to [**[]VariableModeValue**](VariableModeValue.md) | For setting a specific value, given a variable and a mode. | [optional] 

## Methods

### NewPostVariablesRequest

`func NewPostVariablesRequest() *PostVariablesRequest`

NewPostVariablesRequest instantiates a new PostVariablesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostVariablesRequestWithDefaults

`func NewPostVariablesRequestWithDefaults() *PostVariablesRequest`

NewPostVariablesRequestWithDefaults instantiates a new PostVariablesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVariableCollections

`func (o *PostVariablesRequest) GetVariableCollections() []VariableCollectionChange`

GetVariableCollections returns the VariableCollections field if non-nil, zero value otherwise.

### GetVariableCollectionsOk

`func (o *PostVariablesRequest) GetVariableCollectionsOk() (*[]VariableCollectionChange, bool)`

GetVariableCollectionsOk returns a tuple with the VariableCollections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableCollections

`func (o *PostVariablesRequest) SetVariableCollections(v []VariableCollectionChange)`

SetVariableCollections sets VariableCollections field to given value.

### HasVariableCollections

`func (o *PostVariablesRequest) HasVariableCollections() bool`

HasVariableCollections returns a boolean if a field has been set.

### GetVariableModes

`func (o *PostVariablesRequest) GetVariableModes() []VariableModeChange`

GetVariableModes returns the VariableModes field if non-nil, zero value otherwise.

### GetVariableModesOk

`func (o *PostVariablesRequest) GetVariableModesOk() (*[]VariableModeChange, bool)`

GetVariableModesOk returns a tuple with the VariableModes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableModes

`func (o *PostVariablesRequest) SetVariableModes(v []VariableModeChange)`

SetVariableModes sets VariableModes field to given value.

### HasVariableModes

`func (o *PostVariablesRequest) HasVariableModes() bool`

HasVariableModes returns a boolean if a field has been set.

### GetVariables

`func (o *PostVariablesRequest) GetVariables() []VariableChange`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *PostVariablesRequest) GetVariablesOk() (*[]VariableChange, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *PostVariablesRequest) SetVariables(v []VariableChange)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *PostVariablesRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetVariableModeValues

`func (o *PostVariablesRequest) GetVariableModeValues() []VariableModeValue`

GetVariableModeValues returns the VariableModeValues field if non-nil, zero value otherwise.

### GetVariableModeValuesOk

`func (o *PostVariablesRequest) GetVariableModeValuesOk() (*[]VariableModeValue, bool)`

GetVariableModeValuesOk returns a tuple with the VariableModeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableModeValues

`func (o *PostVariablesRequest) SetVariableModeValues(v []VariableModeValue)`

SetVariableModeValues sets VariableModeValues field to given value.

### HasVariableModeValues

`func (o *PostVariablesRequest) HasVariableModeValues() bool`

HasVariableModeValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


