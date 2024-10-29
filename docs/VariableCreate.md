# VariableCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action to perform for the variable. | 
**Id** | Pointer to **string** | A temporary id for this variable. | [optional] 
**Name** | **string** | The name of this variable. | 
**VariableCollectionId** | **string** | The variable collection that will contain the variable. You can use the temporary id of a variable collection. | 
**ResolvedType** | **string** | The resolved type of the variable. | 
**Description** | Pointer to **string** | The description of this variable. | [optional] 
**HiddenFromPublishing** | Pointer to **bool** | Whether this variable is hidden when publishing the current file as a library. | [optional] [default to false]
**Scopes** | Pointer to [**[]VariableScope**](VariableScope.md) | An array of scopes in the UI where this variable is shown. Setting this property will show/hide this variable in the variable picker UI for different fields. | [optional] 
**CodeSyntax** | Pointer to [**VariableCodeSyntax**](VariableCodeSyntax.md) |  | [optional] 

## Methods

### NewVariableCreate

`func NewVariableCreate(action string, name string, variableCollectionId string, resolvedType string, ) *VariableCreate`

NewVariableCreate instantiates a new VariableCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableCreateWithDefaults

`func NewVariableCreateWithDefaults() *VariableCreate`

NewVariableCreateWithDefaults instantiates a new VariableCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *VariableCreate) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *VariableCreate) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *VariableCreate) SetAction(v string)`

SetAction sets Action field to given value.


### GetId

`func (o *VariableCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VariableCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VariableCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VariableCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *VariableCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VariableCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VariableCreate) SetName(v string)`

SetName sets Name field to given value.


### GetVariableCollectionId

`func (o *VariableCreate) GetVariableCollectionId() string`

GetVariableCollectionId returns the VariableCollectionId field if non-nil, zero value otherwise.

### GetVariableCollectionIdOk

`func (o *VariableCreate) GetVariableCollectionIdOk() (*string, bool)`

GetVariableCollectionIdOk returns a tuple with the VariableCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableCollectionId

`func (o *VariableCreate) SetVariableCollectionId(v string)`

SetVariableCollectionId sets VariableCollectionId field to given value.


### GetResolvedType

`func (o *VariableCreate) GetResolvedType() string`

GetResolvedType returns the ResolvedType field if non-nil, zero value otherwise.

### GetResolvedTypeOk

`func (o *VariableCreate) GetResolvedTypeOk() (*string, bool)`

GetResolvedTypeOk returns a tuple with the ResolvedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedType

`func (o *VariableCreate) SetResolvedType(v string)`

SetResolvedType sets ResolvedType field to given value.


### GetDescription

`func (o *VariableCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VariableCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VariableCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VariableCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHiddenFromPublishing

`func (o *VariableCreate) GetHiddenFromPublishing() bool`

GetHiddenFromPublishing returns the HiddenFromPublishing field if non-nil, zero value otherwise.

### GetHiddenFromPublishingOk

`func (o *VariableCreate) GetHiddenFromPublishingOk() (*bool, bool)`

GetHiddenFromPublishingOk returns a tuple with the HiddenFromPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromPublishing

`func (o *VariableCreate) SetHiddenFromPublishing(v bool)`

SetHiddenFromPublishing sets HiddenFromPublishing field to given value.

### HasHiddenFromPublishing

`func (o *VariableCreate) HasHiddenFromPublishing() bool`

HasHiddenFromPublishing returns a boolean if a field has been set.

### GetScopes

`func (o *VariableCreate) GetScopes() []VariableScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *VariableCreate) GetScopesOk() (*[]VariableScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *VariableCreate) SetScopes(v []VariableScope)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *VariableCreate) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetCodeSyntax

`func (o *VariableCreate) GetCodeSyntax() VariableCodeSyntax`

GetCodeSyntax returns the CodeSyntax field if non-nil, zero value otherwise.

### GetCodeSyntaxOk

`func (o *VariableCreate) GetCodeSyntaxOk() (*VariableCodeSyntax, bool)`

GetCodeSyntaxOk returns a tuple with the CodeSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSyntax

`func (o *VariableCreate) SetCodeSyntax(v VariableCodeSyntax)`

SetCodeSyntax sets CodeSyntax field to given value.

### HasCodeSyntax

`func (o *VariableCreate) HasCodeSyntax() bool`

HasCodeSyntax returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


