# LocalVariableValuesByModeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**R** | **float32** | Red channel value, between 0 and 1. | 
**G** | **float32** | Green channel value, between 0 and 1. | 
**B** | **float32** | Blue channel value, between 0 and 1. | 
**A** | **float32** | Alpha channel value, between 0 and 1. | 
**Type** | **string** |  | 
**Id** | **string** | The id of the variable that the current variable is aliased to. This variable can be a local or remote variable, and both can be retrieved via the GET /v1/files/:file_key/variables/local endpoint. | 

## Methods

### NewLocalVariableValuesByModeValue

`func NewLocalVariableValuesByModeValue(r float32, g float32, b float32, a float32, type_ string, id string, ) *LocalVariableValuesByModeValue`

NewLocalVariableValuesByModeValue instantiates a new LocalVariableValuesByModeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalVariableValuesByModeValueWithDefaults

`func NewLocalVariableValuesByModeValueWithDefaults() *LocalVariableValuesByModeValue`

NewLocalVariableValuesByModeValueWithDefaults instantiates a new LocalVariableValuesByModeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetR

`func (o *LocalVariableValuesByModeValue) GetR() float32`

GetR returns the R field if non-nil, zero value otherwise.

### GetROk

`func (o *LocalVariableValuesByModeValue) GetROk() (*float32, bool)`

GetROk returns a tuple with the R field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetR

`func (o *LocalVariableValuesByModeValue) SetR(v float32)`

SetR sets R field to given value.


### GetG

`func (o *LocalVariableValuesByModeValue) GetG() float32`

GetG returns the G field if non-nil, zero value otherwise.

### GetGOk

`func (o *LocalVariableValuesByModeValue) GetGOk() (*float32, bool)`

GetGOk returns a tuple with the G field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetG

`func (o *LocalVariableValuesByModeValue) SetG(v float32)`

SetG sets G field to given value.


### GetB

`func (o *LocalVariableValuesByModeValue) GetB() float32`

GetB returns the B field if non-nil, zero value otherwise.

### GetBOk

`func (o *LocalVariableValuesByModeValue) GetBOk() (*float32, bool)`

GetBOk returns a tuple with the B field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetB

`func (o *LocalVariableValuesByModeValue) SetB(v float32)`

SetB sets B field to given value.


### GetA

`func (o *LocalVariableValuesByModeValue) GetA() float32`

GetA returns the A field if non-nil, zero value otherwise.

### GetAOk

`func (o *LocalVariableValuesByModeValue) GetAOk() (*float32, bool)`

GetAOk returns a tuple with the A field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetA

`func (o *LocalVariableValuesByModeValue) SetA(v float32)`

SetA sets A field to given value.


### GetType

`func (o *LocalVariableValuesByModeValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LocalVariableValuesByModeValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LocalVariableValuesByModeValue) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *LocalVariableValuesByModeValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LocalVariableValuesByModeValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LocalVariableValuesByModeValue) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

