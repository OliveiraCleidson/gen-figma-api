# Overrides

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique ID for a node. | 
**OverriddenFields** | **[]string** | An array of properties. | 

## Methods

### NewOverrides

`func NewOverrides(id string, overriddenFields []string, ) *Overrides`

NewOverrides instantiates a new Overrides object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverridesWithDefaults

`func NewOverridesWithDefaults() *Overrides`

NewOverridesWithDefaults instantiates a new Overrides object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Overrides) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Overrides) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Overrides) SetId(v string)`

SetId sets Id field to given value.


### GetOverriddenFields

`func (o *Overrides) GetOverriddenFields() []string`

GetOverriddenFields returns the OverriddenFields field if non-nil, zero value otherwise.

### GetOverriddenFieldsOk

`func (o *Overrides) GetOverriddenFieldsOk() (*[]string, bool)`

GetOverriddenFieldsOk returns a tuple with the OverriddenFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverriddenFields

`func (o *Overrides) SetOverriddenFields(v []string)`

SetOverriddenFields sets OverriddenFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


