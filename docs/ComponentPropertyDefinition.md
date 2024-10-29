# ComponentPropertyDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ComponentPropertyType**](ComponentPropertyType.md) | Type of this component property. | 
**DefaultValue** | [**ComponentPropertyDefinitionDefaultValue**](ComponentPropertyDefinitionDefaultValue.md) |  | 
**VariantOptions** | Pointer to **[]string** | All possible values for this property. Only exists on VARIANT properties. | [optional] 
**PreferredValues** | Pointer to [**[]InstanceSwapPreferredValue**](InstanceSwapPreferredValue.md) | Preferred values for this property. Only applicable if type is &#x60;INSTANCE_SWAP&#x60;. | [optional] 

## Methods

### NewComponentPropertyDefinition

`func NewComponentPropertyDefinition(type_ ComponentPropertyType, defaultValue ComponentPropertyDefinitionDefaultValue, ) *ComponentPropertyDefinition`

NewComponentPropertyDefinition instantiates a new ComponentPropertyDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentPropertyDefinitionWithDefaults

`func NewComponentPropertyDefinitionWithDefaults() *ComponentPropertyDefinition`

NewComponentPropertyDefinitionWithDefaults instantiates a new ComponentPropertyDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ComponentPropertyDefinition) GetType() ComponentPropertyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComponentPropertyDefinition) GetTypeOk() (*ComponentPropertyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComponentPropertyDefinition) SetType(v ComponentPropertyType)`

SetType sets Type field to given value.


### GetDefaultValue

`func (o *ComponentPropertyDefinition) GetDefaultValue() ComponentPropertyDefinitionDefaultValue`

GetDefaultValue returns the DefaultValue field if non-nil, zero value otherwise.

### GetDefaultValueOk

`func (o *ComponentPropertyDefinition) GetDefaultValueOk() (*ComponentPropertyDefinitionDefaultValue, bool)`

GetDefaultValueOk returns a tuple with the DefaultValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultValue

`func (o *ComponentPropertyDefinition) SetDefaultValue(v ComponentPropertyDefinitionDefaultValue)`

SetDefaultValue sets DefaultValue field to given value.


### GetVariantOptions

`func (o *ComponentPropertyDefinition) GetVariantOptions() []string`

GetVariantOptions returns the VariantOptions field if non-nil, zero value otherwise.

### GetVariantOptionsOk

`func (o *ComponentPropertyDefinition) GetVariantOptionsOk() (*[]string, bool)`

GetVariantOptionsOk returns a tuple with the VariantOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariantOptions

`func (o *ComponentPropertyDefinition) SetVariantOptions(v []string)`

SetVariantOptions sets VariantOptions field to given value.

### HasVariantOptions

`func (o *ComponentPropertyDefinition) HasVariantOptions() bool`

HasVariantOptions returns a boolean if a field has been set.

### GetPreferredValues

`func (o *ComponentPropertyDefinition) GetPreferredValues() []InstanceSwapPreferredValue`

GetPreferredValues returns the PreferredValues field if non-nil, zero value otherwise.

### GetPreferredValuesOk

`func (o *ComponentPropertyDefinition) GetPreferredValuesOk() (*[]InstanceSwapPreferredValue, bool)`

GetPreferredValuesOk returns a tuple with the PreferredValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredValues

`func (o *ComponentPropertyDefinition) SetPreferredValues(v []InstanceSwapPreferredValue)`

SetPreferredValues sets PreferredValues field to given value.

### HasPreferredValues

`func (o *ComponentPropertyDefinition) HasPreferredValues() bool`

HasPreferredValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


