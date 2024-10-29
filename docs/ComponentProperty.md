# ComponentProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ComponentPropertyType**](ComponentPropertyType.md) | Type of this component property. | 
**Value** | [**ComponentPropertyValue**](ComponentPropertyValue.md) |  | 
**PreferredValues** | Pointer to [**[]InstanceSwapPreferredValue**](InstanceSwapPreferredValue.md) | Preferred values for this property. Only applicable if type is &#x60;INSTANCE_SWAP&#x60;. | [optional] 
**BoundVariables** | Pointer to [**ComponentPropertyBoundVariables**](ComponentPropertyBoundVariables.md) |  | [optional] 

## Methods

### NewComponentProperty

`func NewComponentProperty(type_ ComponentPropertyType, value ComponentPropertyValue, ) *ComponentProperty`

NewComponentProperty instantiates a new ComponentProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentPropertyWithDefaults

`func NewComponentPropertyWithDefaults() *ComponentProperty`

NewComponentPropertyWithDefaults instantiates a new ComponentProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ComponentProperty) GetType() ComponentPropertyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ComponentProperty) GetTypeOk() (*ComponentPropertyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ComponentProperty) SetType(v ComponentPropertyType)`

SetType sets Type field to given value.


### GetValue

`func (o *ComponentProperty) GetValue() ComponentPropertyValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ComponentProperty) GetValueOk() (*ComponentPropertyValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ComponentProperty) SetValue(v ComponentPropertyValue)`

SetValue sets Value field to given value.


### GetPreferredValues

`func (o *ComponentProperty) GetPreferredValues() []InstanceSwapPreferredValue`

GetPreferredValues returns the PreferredValues field if non-nil, zero value otherwise.

### GetPreferredValuesOk

`func (o *ComponentProperty) GetPreferredValuesOk() (*[]InstanceSwapPreferredValue, bool)`

GetPreferredValuesOk returns a tuple with the PreferredValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredValues

`func (o *ComponentProperty) SetPreferredValues(v []InstanceSwapPreferredValue)`

SetPreferredValues sets PreferredValues field to given value.

### HasPreferredValues

`func (o *ComponentProperty) HasPreferredValues() bool`

HasPreferredValues returns a boolean if a field has been set.

### GetBoundVariables

`func (o *ComponentProperty) GetBoundVariables() ComponentPropertyBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *ComponentProperty) GetBoundVariablesOk() (*ComponentPropertyBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *ComponentProperty) SetBoundVariables(v ComponentPropertyBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *ComponentProperty) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


