# ComponentPropertiesTrait

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ComponentPropertyDefinitions** | Pointer to [**map[string]ComponentPropertyDefinition**](ComponentPropertyDefinition.md) | A mapping of name to &#x60;ComponentPropertyDefinition&#x60; for every component property on this component. Each property has a type, defaultValue, and other optional values. | [optional] 

## Methods

### NewComponentPropertiesTrait

`func NewComponentPropertiesTrait() *ComponentPropertiesTrait`

NewComponentPropertiesTrait instantiates a new ComponentPropertiesTrait object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentPropertiesTraitWithDefaults

`func NewComponentPropertiesTraitWithDefaults() *ComponentPropertiesTrait`

NewComponentPropertiesTraitWithDefaults instantiates a new ComponentPropertiesTrait object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComponentPropertyDefinitions

`func (o *ComponentPropertiesTrait) GetComponentPropertyDefinitions() map[string]ComponentPropertyDefinition`

GetComponentPropertyDefinitions returns the ComponentPropertyDefinitions field if non-nil, zero value otherwise.

### GetComponentPropertyDefinitionsOk

`func (o *ComponentPropertiesTrait) GetComponentPropertyDefinitionsOk() (*map[string]ComponentPropertyDefinition, bool)`

GetComponentPropertyDefinitionsOk returns a tuple with the ComponentPropertyDefinitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentPropertyDefinitions

`func (o *ComponentPropertiesTrait) SetComponentPropertyDefinitions(v map[string]ComponentPropertyDefinition)`

SetComponentPropertyDefinitions sets ComponentPropertyDefinitions field to given value.

### HasComponentPropertyDefinitions

`func (o *ComponentPropertiesTrait) HasComponentPropertyDefinitions() bool`

HasComponentPropertyDefinitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


