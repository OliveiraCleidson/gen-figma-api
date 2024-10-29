# ActivityLog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The ID of the event. | 
**Timestamp** | **float32** | The timestamp of the event in seconds since the Unix epoch. | 
**Actor** | [**ActivityLogActor**](ActivityLogActor.md) |  | 
**Action** | [**ActivityLogAction**](ActivityLogAction.md) |  | 
**Entity** | [**ActivityLogEntity**](ActivityLogEntity.md) |  | 
**Context** | [**ActivityLogContext**](ActivityLogContext.md) |  | 

## Methods

### NewActivityLog

`func NewActivityLog(id string, timestamp float32, actor ActivityLogActor, action ActivityLogAction, entity ActivityLogEntity, context ActivityLogContext, ) *ActivityLog`

NewActivityLog instantiates a new ActivityLog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogWithDefaults

`func NewActivityLogWithDefaults() *ActivityLog`

NewActivityLogWithDefaults instantiates a new ActivityLog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ActivityLog) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityLog) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityLog) SetId(v string)`

SetId sets Id field to given value.


### GetTimestamp

`func (o *ActivityLog) GetTimestamp() float32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ActivityLog) GetTimestampOk() (*float32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ActivityLog) SetTimestamp(v float32)`

SetTimestamp sets Timestamp field to given value.


### GetActor

`func (o *ActivityLog) GetActor() ActivityLogActor`

GetActor returns the Actor field if non-nil, zero value otherwise.

### GetActorOk

`func (o *ActivityLog) GetActorOk() (*ActivityLogActor, bool)`

GetActorOk returns a tuple with the Actor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActor

`func (o *ActivityLog) SetActor(v ActivityLogActor)`

SetActor sets Actor field to given value.


### GetAction

`func (o *ActivityLog) GetAction() ActivityLogAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ActivityLog) GetActionOk() (*ActivityLogAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ActivityLog) SetAction(v ActivityLogAction)`

SetAction sets Action field to given value.


### GetEntity

`func (o *ActivityLog) GetEntity() ActivityLogEntity`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *ActivityLog) GetEntityOk() (*ActivityLogEntity, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *ActivityLog) SetEntity(v ActivityLogEntity)`

SetEntity sets Entity field to given value.


### GetContext

`func (o *ActivityLog) GetContext() ActivityLogContext`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ActivityLog) GetContextOk() (*ActivityLogContext, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ActivityLog) SetContext(v ActivityLogContext)`

SetContext sets Context field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


