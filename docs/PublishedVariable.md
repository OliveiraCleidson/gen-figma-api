# PublishedVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of this variable. | 
**SubscribedId** | **string** | The ID of the variable that is used by subscribing files. This ID changes every time the variable is modified and published. | 
**Name** | **string** | The name of this variable. | 
**Key** | **string** | The key of this variable. | 
**VariableCollectionId** | **string** | The id of the variable collection that contains this variable. | 
**ResolvedDataType** | **string** | The resolved type of the variable. | 
**UpdatedAt** | **time.Time** | The UTC ISO 8601 time at which the variable was last updated. | 

## Methods

### NewPublishedVariable

`func NewPublishedVariable(id string, subscribedId string, name string, key string, variableCollectionId string, resolvedDataType string, updatedAt time.Time, ) *PublishedVariable`

NewPublishedVariable instantiates a new PublishedVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedVariableWithDefaults

`func NewPublishedVariableWithDefaults() *PublishedVariable`

NewPublishedVariableWithDefaults instantiates a new PublishedVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublishedVariable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublishedVariable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublishedVariable) SetId(v string)`

SetId sets Id field to given value.


### GetSubscribedId

`func (o *PublishedVariable) GetSubscribedId() string`

GetSubscribedId returns the SubscribedId field if non-nil, zero value otherwise.

### GetSubscribedIdOk

`func (o *PublishedVariable) GetSubscribedIdOk() (*string, bool)`

GetSubscribedIdOk returns a tuple with the SubscribedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedId

`func (o *PublishedVariable) SetSubscribedId(v string)`

SetSubscribedId sets SubscribedId field to given value.


### GetName

`func (o *PublishedVariable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublishedVariable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublishedVariable) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *PublishedVariable) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PublishedVariable) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PublishedVariable) SetKey(v string)`

SetKey sets Key field to given value.


### GetVariableCollectionId

`func (o *PublishedVariable) GetVariableCollectionId() string`

GetVariableCollectionId returns the VariableCollectionId field if non-nil, zero value otherwise.

### GetVariableCollectionIdOk

`func (o *PublishedVariable) GetVariableCollectionIdOk() (*string, bool)`

GetVariableCollectionIdOk returns a tuple with the VariableCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableCollectionId

`func (o *PublishedVariable) SetVariableCollectionId(v string)`

SetVariableCollectionId sets VariableCollectionId field to given value.


### GetResolvedDataType

`func (o *PublishedVariable) GetResolvedDataType() string`

GetResolvedDataType returns the ResolvedDataType field if non-nil, zero value otherwise.

### GetResolvedDataTypeOk

`func (o *PublishedVariable) GetResolvedDataTypeOk() (*string, bool)`

GetResolvedDataTypeOk returns a tuple with the ResolvedDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedDataType

`func (o *PublishedVariable) SetResolvedDataType(v string)`

SetResolvedDataType sets ResolvedDataType field to given value.


### GetUpdatedAt

`func (o *PublishedVariable) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublishedVariable) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublishedVariable) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


