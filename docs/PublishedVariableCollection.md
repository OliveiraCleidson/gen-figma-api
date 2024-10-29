# PublishedVariableCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The unique identifier of this variable collection. | 
**SubscribedId** | **string** | The ID of the variable collection that is used by subscribing files. This ID changes every time the variable collection is modified and published. | 
**Name** | **string** | The name of this variable collection. | 
**Key** | **string** | The key of this variable collection. | 
**UpdatedAt** | **time.Time** | The UTC ISO 8601 time at which the variable collection was last updated.  This timestamp will change any time a variable in the collection is changed. | 

## Methods

### NewPublishedVariableCollection

`func NewPublishedVariableCollection(id string, subscribedId string, name string, key string, updatedAt time.Time, ) *PublishedVariableCollection`

NewPublishedVariableCollection instantiates a new PublishedVariableCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishedVariableCollectionWithDefaults

`func NewPublishedVariableCollectionWithDefaults() *PublishedVariableCollection`

NewPublishedVariableCollectionWithDefaults instantiates a new PublishedVariableCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PublishedVariableCollection) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PublishedVariableCollection) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PublishedVariableCollection) SetId(v string)`

SetId sets Id field to given value.


### GetSubscribedId

`func (o *PublishedVariableCollection) GetSubscribedId() string`

GetSubscribedId returns the SubscribedId field if non-nil, zero value otherwise.

### GetSubscribedIdOk

`func (o *PublishedVariableCollection) GetSubscribedIdOk() (*string, bool)`

GetSubscribedIdOk returns a tuple with the SubscribedId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribedId

`func (o *PublishedVariableCollection) SetSubscribedId(v string)`

SetSubscribedId sets SubscribedId field to given value.


### GetName

`func (o *PublishedVariableCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PublishedVariableCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PublishedVariableCollection) SetName(v string)`

SetName sets Name field to given value.


### GetKey

`func (o *PublishedVariableCollection) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PublishedVariableCollection) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PublishedVariableCollection) SetKey(v string)`

SetKey sets Key field to given value.


### GetUpdatedAt

`func (o *PublishedVariableCollection) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PublishedVariableCollection) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PublishedVariableCollection) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


