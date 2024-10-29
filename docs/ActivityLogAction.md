# ActivityLogAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of the action. | 
**Details** | **map[string]interface{}** | Metadata of the action. Each action type supports its own metadata attributes. | 

## Methods

### NewActivityLogAction

`func NewActivityLogAction(type_ string, details map[string]interface{}, ) *ActivityLogAction`

NewActivityLogAction instantiates a new ActivityLogAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityLogActionWithDefaults

`func NewActivityLogActionWithDefaults() *ActivityLogAction`

NewActivityLogActionWithDefaults instantiates a new ActivityLogAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ActivityLogAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ActivityLogAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ActivityLogAction) SetType(v string)`

SetType sets Type field to given value.


### GetDetails

`func (o *ActivityLogAction) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ActivityLogAction) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ActivityLogAction) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.


### SetDetailsNil

`func (o *ActivityLogAction) SetDetailsNil(b bool)`

 SetDetailsNil sets the value for Details to be an explicit nil

### UnsetDetails
`func (o *ActivityLogAction) UnsetDetails()`

UnsetDetails ensures that no value is present for Details, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


