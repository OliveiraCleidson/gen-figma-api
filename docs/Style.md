# Style

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | The key of the style | 
**Name** | **string** | Name of the style | 
**Description** | **string** | Description of the style | 
**Remote** | **bool** | Whether this style is a remote style that doesn&#39;t live in this file | 
**StyleType** | [**StyleType**](StyleType.md) |  | 

## Methods

### NewStyle

`func NewStyle(key string, name string, description string, remote bool, styleType StyleType, ) *Style`

NewStyle instantiates a new Style object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStyleWithDefaults

`func NewStyleWithDefaults() *Style`

NewStyleWithDefaults instantiates a new Style object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *Style) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Style) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Style) SetKey(v string)`

SetKey sets Key field to given value.


### GetName

`func (o *Style) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Style) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Style) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Style) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Style) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Style) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetRemote

`func (o *Style) GetRemote() bool`

GetRemote returns the Remote field if non-nil, zero value otherwise.

### GetRemoteOk

`func (o *Style) GetRemoteOk() (*bool, bool)`

GetRemoteOk returns a tuple with the Remote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemote

`func (o *Style) SetRemote(v bool)`

SetRemote sets Remote field to given value.


### GetStyleType

`func (o *Style) GetStyleType() StyleType`

GetStyleType returns the StyleType field if non-nil, zero value otherwise.

### GetStyleTypeOk

`func (o *Style) GetStyleTypeOk() (*StyleType, bool)`

GetStyleTypeOk returns a tuple with the StyleType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleType

`func (o *Style) SetStyleType(v StyleType)`

SetStyleType sets StyleType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


