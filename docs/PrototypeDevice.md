# PrototypeDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Size** | Pointer to [**Size**](Size.md) |  | [optional] 
**PresetIdentifier** | Pointer to **string** |  | [optional] 
**Rotation** | **string** |  | 

## Methods

### NewPrototypeDevice

`func NewPrototypeDevice(type_ string, rotation string, ) *PrototypeDevice`

NewPrototypeDevice instantiates a new PrototypeDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrototypeDeviceWithDefaults

`func NewPrototypeDeviceWithDefaults() *PrototypeDevice`

NewPrototypeDeviceWithDefaults instantiates a new PrototypeDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PrototypeDevice) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrototypeDevice) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrototypeDevice) SetType(v string)`

SetType sets Type field to given value.


### GetSize

`func (o *PrototypeDevice) GetSize() Size`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PrototypeDevice) GetSizeOk() (*Size, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PrototypeDevice) SetSize(v Size)`

SetSize sets Size field to given value.

### HasSize

`func (o *PrototypeDevice) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetPresetIdentifier

`func (o *PrototypeDevice) GetPresetIdentifier() string`

GetPresetIdentifier returns the PresetIdentifier field if non-nil, zero value otherwise.

### GetPresetIdentifierOk

`func (o *PrototypeDevice) GetPresetIdentifierOk() (*string, bool)`

GetPresetIdentifierOk returns a tuple with the PresetIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresetIdentifier

`func (o *PrototypeDevice) SetPresetIdentifier(v string)`

SetPresetIdentifier sets PresetIdentifier field to given value.

### HasPresetIdentifier

`func (o *PrototypeDevice) HasPresetIdentifier() bool`

HasPresetIdentifier returns a boolean if a field has been set.

### GetRotation

`func (o *PrototypeDevice) GetRotation() string`

GetRotation returns the Rotation field if non-nil, zero value otherwise.

### GetRotationOk

`func (o *PrototypeDevice) GetRotationOk() (*string, bool)`

GetRotationOk returns a tuple with the Rotation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRotation

`func (o *PrototypeDevice) SetRotation(v string)`

SetRotation sets Rotation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


