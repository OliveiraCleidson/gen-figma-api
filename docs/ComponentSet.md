# ComponentSet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the component set | 
**Name** | **string** | Name of the component set | 
**Description** | **string** | The description of the component set as entered in the editor | 
**DocumentationLinks** | Pointer to [**[]DocumentationLink**](DocumentationLink.md) | An array of documentation links attached to this component set | [optional] 
**Remote** | Pointer to **bool** | Whether this component set is a remote component set that doesn&#39;t live in this file | [optional] 

## Methods

### NewComponentSet

`func NewComponentSet(key string, name string, description string, ) *ComponentSet`

NewComponentSet instantiates a new ComponentSet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentSetWithDefaults

`func NewComponentSetWithDefaults() *ComponentSet`

NewComponentSetWithDefaults instantiates a new ComponentSet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ComponentSet) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ComponentSet) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ComponentSet) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *ComponentSet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ComponentSet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ComponentSet) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ComponentSet) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ComponentSet) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ComponentSet) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDocumentationLinks

`func (o *ComponentSet) GetDocumentationLinks() []DocumentationLink`

GetDocumentationLinks returns the DocumentationLinks field if non-nil, zero value otherwise.

### GetDocumentationLinksOk

`func (o *ComponentSet) GetDocumentationLinksOk() (*[]DocumentationLink, bool)`

GetDocumentationLinksOk returns a tuple with the DocumentationLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLinks

`func (o *ComponentSet) SetDocumentationLinks(v []DocumentationLink)`

SetDocumentationLinks sets DocumentationLinks field to given value.

### HasDocumentationLinks

`func (o *ComponentSet) HasDocumentationLinks() bool`

HasDocumentationLinks returns a boolean if a field has been set.

### GetRemote

`func (o *ComponentSet) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *ComponentSet) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *ComponentSet) SetRemote(v bool)`

SetRemote sets Remote field to given value.

### HasRemote

`func (o *ComponentSet) HasRemote() bool`

HasRemote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


