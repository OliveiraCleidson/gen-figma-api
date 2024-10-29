# VariableCollectionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform for the variable collection. | 
**Id** | **string** | The id of the variable collection to update. | 
**Name** | Pointer to **string** | The name of this variable collection. | [optional] 
**HiddenFromPublishing** | Pointer to **bool** | Whether this variable collection is hidden when publishing the current file as a library. | [optional] [default to false]

## Methods

### NewVariableCollectionUpdate

`func NewVariableCollectionUpdate(action string, id string, ) *VariableCollectionUpdate`

NewVariableCollectionUpdate instantiates a new VariableCollectionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableCollectionUpdateWithDefaults

`func NewVariableCollectionUpdateWithDefaults() *VariableCollectionUpdate`

NewVariableCollectionUpdateWithDefaults instantiates a new VariableCollectionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VariableCollectionUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VariableCollectionUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VariableCollectionUpdate) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *VariableCollectionUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariableCollectionUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariableCollectionUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VariableCollectionUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableCollectionUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableCollectionUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VariableCollectionUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetHiddenFromPublishing

`func (o *VariableCollectionUpdate) GetHiddenFromPublishing() bool`

GetHiddenFromPublishing returns the HiddenFromPublishing field if non-nil, zero value otherwise.

### GetHiddenFromPublishingOk

`func (o *VariableCollectionUpdate) GetHiddenFromPublishingOk() (*bool, bool)`

GetHiddenFromPublishingOk returns a tuple with the HiddenFromPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromPublishing

`func (o *VariableCollectionUpdate) SetHiddenFromPublishing(v bool)`

SetHiddenFromPublishing sets HiddenFromPublishing field to given value.

### HasHiddenFromPublishing

`func (o *VariableCollectionUpdate) HasHiddenFromPublishing() bool`

HasHiddenFromPublishing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


