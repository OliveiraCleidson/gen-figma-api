# BlurEffect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A string literal representing the effect&#39;s type. Always check the type before reading other properties. | 
**Visible** | **bool** | Whether this blur is active. | 
**Radius** | **float32** | Radius of the blur effect | 
**BoundVariables** | Pointer to [**BlurEffectBoundVariables**](BlurEffectBoundVariables.md) |  | [optional] 

## Methods

### NewBlurEffect

`func NewBlurEffect(type_ string, visible bool, radius float32, ) *BlurEffect`

NewBlurEffect instantiates a new BlurEffect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlurEffectWithDefaults

`func NewBlurEffectWithDefaults() *BlurEffect`

NewBlurEffectWithDefaults instantiates a new BlurEffect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BlurEffect) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BlurEffect) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BlurEffect) SetType(v string)`

SetType sets Type field to given value.


### GetVisible

`func (o *BlurEffect) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *BlurEffect) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *BlurEffect) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetRadius

`func (o *BlurEffect) GetRadius() float32`

GetRadius returns the Radius field if non-nil, zero value otherwise.

### GetRadiusOk

`func (o *BlurEffect) GetRadiusOk() (*float32, bool)`

GetRadiusOk returns a tuple with the Radius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadius

`func (o *BlurEffect) SetRadius(v float32)`

SetRadius sets Radius field to given value.


### GetBoundVariables

`func (o *BlurEffect) GetBoundVariables() BlurEffectBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *BlurEffect) GetBoundVariablesOk() (*BlurEffectBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *BlurEffect) SetBoundVariables(v BlurEffectBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *BlurEffect) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


