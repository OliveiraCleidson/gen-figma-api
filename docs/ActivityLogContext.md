# ActivityLogContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | **NullableString** | The third-party application that triggered the event, if applicable. | 
**IpAddress** | **string** | The IP address from of the client that sent the event request. | 
**IsFigmaSupportTeamAction** | **bool** | If Figma&#39;s Support team triggered the event. This is either true or false. | 
**OrgId** | **string** | The id of the organization where the event took place. | 
**TeamId** | **NullableString** | The id of the team where the event took place -- if this took place in a specific team. | 

## Methods

### NewActivityLogContext

`func NewActivityLogContext(clientName NullableString, ipAddress string, isFigmaSupportTeamAction bool, orgId string, teamId NullableString, ) *ActivityLogContext`

NewActivityLogContext instantiates a new ActivityLogContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogContextWithDefaults

`func NewActivityLogContextWithDefaults() *ActivityLogContext`

NewActivityLogContextWithDefaults instantiates a new ActivityLogContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *ActivityLogContext) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ActivityLogContext) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ActivityLogContext) SetClientName(v string)`

SetClientName sets ClientName field to given value.


### SetClientNameNil

`func (o *ActivityLogContext) SetClientNameNil(b bool)`

 SetClientNameNil sets the value for ClientName to be an explicit nil

### UnsetClientName
`func (o *ActivityLogContext) UnsetClientName()`

UnsetClientName ensures that no value is present for ClientName, not even an explicit nil
### GetIpAddress

`func (o *ActivityLogContext) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ActivityLogContext) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ActivityLogContext) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.


### GetIsFigmaSupportTeamAction

`func (o *ActivityLogContext) GetIsFigmaSupportTeamAction() bool`

GetIsFigmaSupportTeamAction returns the IsFigmaSupportTeamAction field if non-nil, zero value otherwise.

### GetIsFigmaSupportTeamActionOk

`func (o *ActivityLogContext) GetIsFigmaSupportTeamActionOk() (*bool, bool)`

GetIsFigmaSupportTeamActionOk returns a tuple with the IsFigmaSupportTeamAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFigmaSupportTeamAction

`func (o *ActivityLogContext) SetIsFigmaSupportTeamAction(v bool)`

SetIsFigmaSupportTeamAction sets IsFigmaSupportTeamAction field to given value.


### GetOrgId

`func (o *ActivityLogContext) GetOrgId() string`

GetOrgId returns the OrgId field if non-nil, zero value otherwise.

### GetOrgIdOk

`func (o *ActivityLogContext) GetOrgIdOk() (*string, bool)`

GetOrgIdOk returns a tuple with the OrgId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrgId

`func (o *ActivityLogContext) SetOrgId(v string)`

SetOrgId sets OrgId field to given value.


### GetTeamId

`func (o *ActivityLogContext) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *ActivityLogContext) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *ActivityLogContext) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### SetTeamIdNil

`func (o *ActivityLogContext) SetTeamIdNil(b bool)`

 SetTeamIdNil sets the value for TeamId to be an explicit nil

### UnsetTeamId
`func (o *ActivityLogContext) UnsetTeamId()`

UnsetTeamId ensures that no value is present for TeamId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


