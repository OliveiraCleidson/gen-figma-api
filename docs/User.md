# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique stable id of the user. | 
**Handle** | **string** | Name of the user. | 
**ImgUrl** | **string** | URL link to the user&#39;s profile image. | 

## Methods

### NewUser

`func NewUser(id string, handle string, imgUrl string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetHandle

`func (o *User) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *User) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *User) SetHandle(v string)`

SetHandle sets Handle field to given value.


### GetImgUrl

`func (o *User) GetImgUrl() string`

GetImgUrl returns the ImgUrl field if non-nil, zero value otherwise.

### GetImgUrlOk

`func (o *User) GetImgUrlOk() (*string, bool)`

GetImgUrlOk returns a tuple with the ImgUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImgUrl

`func (o *User) SetImgUrl(v string)`

SetImgUrl sets ImgUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


