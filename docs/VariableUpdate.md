# VariableUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform for the variable. | 
**Id** | **string** | The id of the variable to update. | 
**Name** | Pointer to **string** | The name of this variable. | [optional] 
**Description** | Pointer to **string** | The description of this variable. | [optional] 
**HiddenFromPublishing** | Pointer to **bool** | Whether this variable is hidden when publishing the current file as a library. | [optional] [default to false]
**Scopes** | Pointer to [**[]VariableScope**](VariableScope.md) | An array of scopes in the UI where this variable is shown. Setting this property will show/hide this variable in the variable picker UI for different fields. | [optional] 
**CodeSyntax** | Pointer to [**VariableCodeSyntax**](VariableCodeSyntax.md) |  | [optional] 

## Methods

### NewVariableUpdate

`func NewVariableUpdate(action string, id string, ) *VariableUpdate`

NewVariableUpdate instantiates a new VariableUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableUpdateWithDefaults

`func NewVariableUpdateWithDefaults() *VariableUpdate`

NewVariableUpdateWithDefaults instantiates a new VariableUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VariableUpdate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VariableUpdate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VariableUpdate) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *VariableUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariableUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariableUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *VariableUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VariableUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VariableUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VariableUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VariableUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VariableUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHiddenFromPublishing

`func (o *VariableUpdate) GetHiddenFromPublishing() bool`

GetHiddenFromPublishing returns the HiddenFromPublishing field if non-nil, zero value otherwise.

### GetHiddenFromPublishingOk

`func (o *VariableUpdate) GetHiddenFromPublishingOk() (*bool, bool)`

GetHiddenFromPublishingOk returns a tuple with the HiddenFromPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromPublishing

`func (o *VariableUpdate) SetHiddenFromPublishing(v bool)`

SetHiddenFromPublishing sets HiddenFromPublishing field to given value.

### HasHiddenFromPublishing

`func (o *VariableUpdate) HasHiddenFromPublishing() bool`

HasHiddenFromPublishing returns a boolean if a field has been set.

### GetScopes

`func (o *VariableUpdate) GetScopes() []VariableScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *VariableUpdate) GetScopesOk() (*[]VariableScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *VariableUpdate) SetScopes(v []VariableScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *VariableUpdate) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetCodeSyntax

`func (o *VariableUpdate) GetCodeSyntax() VariableCodeSyntax`

GetCodeSyntax returns the CodeSyntax field if non-nil, zero value otherwise.

### GetCodeSyntaxOk

`func (o *VariableUpdate) GetCodeSyntaxOk() (*VariableCodeSyntax, bool)`

GetCodeSyntaxOk returns a tuple with the CodeSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSyntax

`func (o *VariableUpdate) SetCodeSyntax(v VariableCodeSyntax)`

SetCodeSyntax sets CodeSyntax field to given value.

### HasCodeSyntax

`func (o *VariableUpdate) HasCodeSyntax() bool`

HasCodeSyntax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


