# LocalVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of this variable. | 
**Name** | **string** | The name of this variable. | 
**Key** | **string** | The key of this variable. | 
**VariableCollectionId** | **string** | The id of the variable collection that contains this variable. | 
**ResolvedType** | **string** | The resolved type of the variable. | 
**ValuesByMode** | [**map[string]LocalVariableValuesByModeValue**](LocalVariableValuesByModeValue.md) | The values for each mode of this variable. | 
**Remote** | **bool** | Whether this variable is remote. | 
**Description** | **string** | The description of this variable. | 
**HiddenFromPublishing** | **bool** | Whether this variable is hidden when publishing the current file as a library.  If the parent &#x60;VariableCollection&#x60; is marked as &#x60;hiddenFromPublishing&#x60;, then this variable will also be hidden from publishing via the UI. &#x60;hiddenFromPublishing&#x60; is independently toggled for a variable and collection. However, both must be true for a given variable to be publishable. | 
**Scopes** | [**[]VariableScope**](VariableScope.md) | An array of scopes in the UI where this variable is shown. Setting this property will show/hide this variable in the variable picker UI for different fields.  Setting scopes for a variable does not prevent that variable from being bound in other scopes (for example, via the Plugin API). This only limits the variables that are shown in pickers within the Figma UI. | 
**CodeSyntax** | [**VariableCodeSyntax**](VariableCodeSyntax.md) |  | 
**DeletedButReferenced** | Pointer to **bool** | Indicates that the variable was deleted in the editor, but the document may still contain references to the variable. References to the variable may exist through bound values or variable aliases. | [optional] [default to false]

## Methods

### NewLocalVariable

`func NewLocalVariable(id string, name string, key string, variableCollectionId string, resolvedType string, valuesByMode map[string]LocalVariableValuesByModeValue, remote bool, description string, hiddenFromPublishing bool, scopes []VariableScope, codeSyntax VariableCodeSyntax, ) *LocalVariable`

NewLocalVariable instantiates a new LocalVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalVariableWithDefaults

`func NewLocalVariableWithDefaults() *LocalVariable`

NewLocalVariableWithDefaults instantiates a new LocalVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LocalVariable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalVariable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalVariable) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *LocalVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LocalVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LocalVariable) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *LocalVariable) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *LocalVariable) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *LocalVariable) SetKey(v string)`

SetKey sets Key field to given value.


### GetVariableCollectionId

`func (o *LocalVariable) GetVariableCollectionId() string`

GetVariableCollectionId returns the VariableCollectionId field if non-nil, zero value otherwise.

### GetVariableCollectionIdOk

`func (o *LocalVariable) GetVariableCollectionIdOk() (*string, bool)`

GetVariableCollectionIdOk returns a tuple with the VariableCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableCollectionId

`func (o *LocalVariable) SetVariableCollectionId(v string)`

SetVariableCollectionId sets VariableCollectionId field to given value.


### GetResolvedType

`func (o *LocalVariable) GetResolvedType() string`

GetResolvedType returns the ResolvedType field if non-nil, zero value otherwise.

### GetResolvedTypeOk

`func (o *LocalVariable) GetResolvedTypeOk() (*string, bool)`

GetResolvedTypeOk returns a tuple with the ResolvedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedType

`func (o *LocalVariable) SetResolvedType(v string)`

SetResolvedType sets ResolvedType field to given value.


### GetValuesByMode

`func (o *LocalVariable) GetValuesByMode() map[string]LocalVariableValuesByModeValue`

GetValuesByMode returns the ValuesByMode field if non-nil, zero value otherwise.

### GetValuesByModeOk

`func (o *LocalVariable) GetValuesByModeOk() (*map[string]LocalVariableValuesByModeValue, bool)`

GetValuesByModeOk returns a tuple with the ValuesByMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesByMode

`func (o *LocalVariable) SetValuesByMode(v map[string]LocalVariableValuesByModeValue)`

SetValuesByMode sets ValuesByMode field to given value.


### GetRemote

`func (o *LocalVariable) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *LocalVariable) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *LocalVariable) SetRemote(v bool)`

SetRemote sets Remote field to given value.


### GetDescription

`func (o *LocalVariable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LocalVariable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LocalVariable) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetHiddenFromPublishing

`func (o *LocalVariable) GetHiddenFromPublishing() bool`

GetHiddenFromPublishing returns the HiddenFromPublishing field if non-nil, zero value otherwise.

### GetHiddenFromPublishingOk

`func (o *LocalVariable) GetHiddenFromPublishingOk() (*bool, bool)`

GetHiddenFromPublishingOk returns a tuple with the HiddenFromPublishing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHiddenFromPublishing

`func (o *LocalVariable) SetHiddenFromPublishing(v bool)`

SetHiddenFromPublishing sets HiddenFromPublishing field to given value.


### GetScopes

`func (o *LocalVariable) GetScopes() []VariableScope`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *LocalVariable) GetScopesOk() (*[]VariableScope, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *LocalVariable) SetScopes(v []VariableScope)`

SetScopes sets Scopes field to given value.


### GetCodeSyntax

`func (o *LocalVariable) GetCodeSyntax() VariableCodeSyntax`

GetCodeSyntax returns the CodeSyntax field if non-nil, zero value otherwise.

### GetCodeSyntaxOk

`func (o *LocalVariable) GetCodeSyntaxOk() (*VariableCodeSyntax, bool)`

GetCodeSyntaxOk returns a tuple with the CodeSyntax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeSyntax

`func (o *LocalVariable) SetCodeSyntax(v VariableCodeSyntax)`

SetCodeSyntax sets CodeSyntax field to given value.


### GetDeletedButReferenced

`func (o *LocalVariable) GetDeletedButReferenced() bool`

GetDeletedButReferenced returns the DeletedButReferenced field if non-nil, zero value otherwise.

### GetDeletedButReferencedOk

`func (o *LocalVariable) GetDeletedButReferencedOk() (*bool, bool)`

GetDeletedButReferencedOk returns a tuple with the DeletedButReferenced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedButReferenced

`func (o *LocalVariable) SetDeletedButReferenced(v bool)`

SetDeletedButReferenced sets DeletedButReferenced field to given value.

### HasDeletedButReferenced

`func (o *LocalVariable) HasDeletedButReferenced() bool`

HasDeletedButReferenced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


