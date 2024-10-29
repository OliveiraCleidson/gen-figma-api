# ActivityLogUserEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of entity. | 
**Id** | **string** | Unique stable id of the user. | 
**Name** | **string** | Name of the user. | 
**Email** | **string** | Email associated with the user&#39;s account. | 

## Methods

### NewActivityLogUserEntity

`func NewActivityLogUserEntity(type_ string, id string, name string, email string, ) *ActivityLogUserEntity`

NewActivityLogUserEntity instantiates a new ActivityLogUserEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogUserEntityWithDefaults

`func NewActivityLogUserEntityWithDefaults() *ActivityLogUserEntity`

NewActivityLogUserEntityWithDefaults instantiates a new ActivityLogUserEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActivityLogUserEntity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityLogUserEntity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityLogUserEntity) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ActivityLogUserEntity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityLogUserEntity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityLogUserEntity) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ActivityLogUserEntity) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityLogUserEntity) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityLogUserEntity) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *ActivityLogUserEntity) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ActivityLogUserEntity) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ActivityLogUserEntity) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


