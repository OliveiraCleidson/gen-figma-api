# TriggerOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Delay** | **float32** |  | 
**DeprecatedVersion** | Pointer to **bool** | Whether this is a [deprecated version](https://help.figma.com/hc/en-us/articles/360040035834-Prototype-triggers#h_01HHN04REHJNP168R26P1CMP0A) of the trigger that was left unchanged for backwards compatibility. If not present, the trigger is the latest version. | [optional] 

## Methods

### NewTriggerOneOf1

`func NewTriggerOneOf1(type_ string, delay float32, ) *TriggerOneOf1`

NewTriggerOneOf1 instantiates a new TriggerOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerOneOf1WithDefaults

`func NewTriggerOneOf1WithDefaults() *TriggerOneOf1`

NewTriggerOneOf1WithDefaults instantiates a new TriggerOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TriggerOneOf1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TriggerOneOf1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TriggerOneOf1) SetType(v string)`

SetType sets Type field to given value.


### GetDelay

`func (o *TriggerOneOf1) GetDelay() float32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *TriggerOneOf1) GetDelayOk() (*float32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *TriggerOneOf1) SetDelay(v float32)`

SetDelay sets Delay field to given value.


### GetDeprecatedVersion

`func (o *TriggerOneOf1) GetDeprecatedVersion() bool`

GetDeprecatedVersion returns the DeprecatedVersion field if non-nil, zero value otherwise.

### GetDeprecatedVersionOk

`func (o *TriggerOneOf1) GetDeprecatedVersionOk() (*bool, bool)`

GetDeprecatedVersionOk returns a tuple with the DeprecatedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedVersion

`func (o *TriggerOneOf1) SetDeprecatedVersion(v bool)`

SetDeprecatedVersion sets DeprecatedVersion field to given value.

### HasDeprecatedVersion

`func (o *TriggerOneOf1) HasDeprecatedVersion() bool`

HasDeprecatedVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


