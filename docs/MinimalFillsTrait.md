# MinimalFillsTrait

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fills** | [**[]Paint**](Paint.md) | An array of fill paints applied to the node. | 
**Styles** | Pointer to **map[string]string** | A mapping of a StyleType to style ID (see Style) of styles present on this node. The style ID can be used to look up more information about the style in the top-level styles field. | [optional] 

## Methods

### NewMinimalFillsTrait

`func NewMinimalFillsTrait(fills []Paint, ) *MinimalFillsTrait`

NewMinimalFillsTrait instantiates a new MinimalFillsTrait object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMinimalFillsTraitWithDefaults

`func NewMinimalFillsTraitWithDefaults() *MinimalFillsTrait`

NewMinimalFillsTraitWithDefaults instantiates a new MinimalFillsTrait object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFills

`func (o *MinimalFillsTrait) GetFills() []Paint`

GetFills returns the Fills field if non-nil, zero value otherwise.

### GetFillsOk

`func (o *MinimalFillsTrait) GetFillsOk() (*[]Paint, bool)`

GetFillsOk returns a tuple with the Fills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFills

`func (o *MinimalFillsTrait) SetFills(v []Paint)`

SetFills sets Fills field to given value.


### GetStyles

`func (o *MinimalFillsTrait) GetStyles() map[string]string`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *MinimalFillsTrait) GetStylesOk() (*map[string]string, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *MinimalFillsTrait) SetStyles(v map[string]string)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *MinimalFillsTrait) HasStyles() bool`

HasStyles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


