# VariableData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VariableDataType**](VariableDataType.md) |  | [optional] 
**ResolvedType** | Pointer to [**VariableResolvedDataType**](VariableResolvedDataType.md) |  | [optional] 
**Value** | Pointer to [**VariableDataValue**](VariableDataValue.md) |  | [optional] 

## Methods

### NewVariableData

`func NewVariableData() *VariableData`

NewVariableData instantiates a new VariableData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVariableDataWithDefaults

`func NewVariableDataWithDefaults() *VariableData`

NewVariableDataWithDefaults instantiates a new VariableData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VariableData) GetType() VariableDataType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VariableData) GetTypeOk() (*VariableDataType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VariableData) SetType(v VariableDataType)`

SetType sets Type field to given value.

### HasType

`func (o *VariableData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResolvedType

`func (o *VariableData) GetResolvedType() VariableResolvedDataType`

GetResolvedType returns the ResolvedType field if non-nil, zero value otherwise.

### GetResolvedTypeOk

`func (o *VariableData) GetResolvedTypeOk() (*VariableResolvedDataType, bool)`

GetResolvedTypeOk returns a tuple with the ResolvedType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedType

`func (o *VariableData) SetResolvedType(v VariableResolvedDataType)`

SetResolvedType sets ResolvedType field to given value.

### HasResolvedType

`func (o *VariableData) HasResolvedType() bool`

HasResolvedType returns a boolean if a field has been set.

### GetValue

`func (o *VariableData) GetValue() VariableDataValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VariableData) GetValueOk() (*VariableDataValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VariableData) SetValue(v VariableDataValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *VariableData) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


