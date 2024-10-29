# Trigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Timeout** | **float32** |  | 
**Delay** | **float32** |  | 
**DeprecatedVersion** | Pointer to **bool** | Whether this is a [deprecated version](https://help.figma.com/hc/en-us/articles/360040035834-Prototype-triggers#h_01HHN04REHJNP168R26P1CMP0A) of the trigger that was left unchanged for backwards compatibility. If not present, the trigger is the latest version. | [optional] 
**Device** | **string** |  | 
**KeyCodes** | **[]float32** |  | 
**MediaHitTime** | **float32** |  | 

## Methods

### NewTrigger

`func NewTrigger(type_ string, timeout float32, delay float32, device string, keyCodes []float32, mediaHitTime float32, ) *Trigger`

NewTrigger instantiates a new Trigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWithDefaults

`func NewTriggerWithDefaults() *Trigger`

NewTriggerWithDefaults instantiates a new Trigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Trigger) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Trigger) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Trigger) SetType(v string)`

SetType sets Type field to given value.


### GetTimeout

`func (o *Trigger) GetTimeout() float32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Trigger) GetTimeoutOk() (*float32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Trigger) SetTimeout(v float32)`

SetTimeout sets Timeout field to given value.


### GetDelay

`func (o *Trigger) GetDelay() float32`

GetDelay returns the Delay field if non-nil, zero value otherwise.

### GetDelayOk

`func (o *Trigger) GetDelayOk() (*float32, bool)`

GetDelayOk returns a tuple with the Delay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelay

`func (o *Trigger) SetDelay(v float32)`

SetDelay sets Delay field to given value.


### GetDeprecatedVersion

`func (o *Trigger) GetDeprecatedVersion() bool`

GetDeprecatedVersion returns the DeprecatedVersion field if non-nil, zero value otherwise.

### GetDeprecatedVersionOk

`func (o *Trigger) GetDeprecatedVersionOk() (*bool, bool)`

GetDeprecatedVersionOk returns a tuple with the DeprecatedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedVersion

`func (o *Trigger) SetDeprecatedVersion(v bool)`

SetDeprecatedVersion sets DeprecatedVersion field to given value.

### HasDeprecatedVersion

`func (o *Trigger) HasDeprecatedVersion() bool`

HasDeprecatedVersion returns a boolean if a field has been set.

### GetDevice

`func (o *Trigger) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *Trigger) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *Trigger) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetKeyCodes

`func (o *Trigger) GetKeyCodes() []float32`

GetKeyCodes returns the KeyCodes field if non-nil, zero value otherwise.

### GetKeyCodesOk

`func (o *Trigger) GetKeyCodesOk() (*[]float32, bool)`

GetKeyCodesOk returns a tuple with the KeyCodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyCodes

`func (o *Trigger) SetKeyCodes(v []float32)`

SetKeyCodes sets KeyCodes field to given value.


### GetMediaHitTime

`func (o *Trigger) GetMediaHitTime() float32`

GetMediaHitTime returns the MediaHitTime field if non-nil, zero value otherwise.

### GetMediaHitTimeOk

`func (o *Trigger) GetMediaHitTimeOk() (*float32, bool)`

GetMediaHitTimeOk returns a tuple with the MediaHitTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaHitTime

`func (o *Trigger) SetMediaHitTime(v float32)`

SetMediaHitTime sets MediaHitTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


