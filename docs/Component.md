# Component

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the component | 
**Name** | **string** | Name of the component | 
**Description** | **string** | The description of the component as entered in the editor | 
**ComponentSetId** | Pointer to **string** | The ID of the component set if the component belongs to one | [optional] 
**DocumentationLinks** | [**[]DocumentationLink**](DocumentationLink.md) | An array of documentation links attached to this component | 
**Remote** | **bool** | Whether this component is a remote component that doesn&#39;t live in this file | 

## Methods

### NewComponent

`func NewComponent(key string, name string, description string, documentationLinks []DocumentationLink, remote bool, ) *Component`

NewComponent instantiates a new Component object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentWithDefaults

`func NewComponentWithDefaults() *Component`

NewComponentWithDefaults instantiates a new Component object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Component) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Component) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Component) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *Component) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Component) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Component) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Component) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Component) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Component) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetComponentSetId

`func (o *Component) GetComponentSetId() string`

GetComponentSetId returns the ComponentSetId field if non-nil, zero value otherwise.

### GetComponentSetIdOk

`func (o *Component) GetComponentSetIdOk() (*string, bool)`

GetComponentSetIdOk returns a tuple with the ComponentSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentSetId

`func (o *Component) SetComponentSetId(v string)`

SetComponentSetId sets ComponentSetId field to given value.

### HasComponentSetId

`func (o *Component) HasComponentSetId() bool`

HasComponentSetId returns a boolean if a field has been set.

### GetDocumentationLinks

`func (o *Component) GetDocumentationLinks() []DocumentationLink`

GetDocumentationLinks returns the DocumentationLinks field if non-nil, zero value otherwise.

### GetDocumentationLinksOk

`func (o *Component) GetDocumentationLinksOk() (*[]DocumentationLink, bool)`

GetDocumentationLinksOk returns a tuple with the DocumentationLinks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentationLinks

`func (o *Component) SetDocumentationLinks(v []DocumentationLink)`

SetDocumentationLinks sets DocumentationLinks field to given value.


### GetRemote

`func (o *Component) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *Component) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *Component) SetRemote(v bool)`

SetRemote sets Remote field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


