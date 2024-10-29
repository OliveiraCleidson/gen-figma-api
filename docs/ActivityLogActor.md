# ActivityLogActor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | The type of the user. | [optional] 
**Id** | Pointer to **string** | The ID of the user. | [optional] 
**Name** | **string** | The name of the user. For SCIM events, the value is \&quot;SCIM Provider\&quot;. For official support actions, the value is \&quot;Figma Support\&quot;. | 
**Email** | Pointer to **string** | The email of the user. | [optional] 

## Methods

### NewActivityLogActor

`func NewActivityLogActor(name string, ) *ActivityLogActor`

NewActivityLogActor instantiates a new ActivityLogActor object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogActorWithDefaults

`func NewActivityLogActorWithDefaults() *ActivityLogActor`

NewActivityLogActorWithDefaults instantiates a new ActivityLogActor object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActivityLogActor) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityLogActor) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityLogActor) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ActivityLogActor) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *ActivityLogActor) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityLogActor) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityLogActor) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityLogActor) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActivityLogActor) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityLogActor) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityLogActor) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *ActivityLogActor) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ActivityLogActor) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ActivityLogActor) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ActivityLogActor) HasEmail() bool`

HasEmail returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


