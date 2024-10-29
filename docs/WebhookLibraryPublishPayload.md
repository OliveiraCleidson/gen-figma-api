# WebhookLibraryPublishPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Passcode** | **string** | The passcode specified when the webhook was created, should match what was initially provided | 
**Timestamp** | **time.Time** | UTC ISO 8601 timestamp of when the event was triggered. | 
**WebhookId** | **string** | The id of the webhook that caused the callback | 
**EventType** | **string** |  | 
**CreatedComponents** | [**[]LibraryItemData**](LibraryItemData.md) | Components that were created by the library publish | 
**CreatedStyles** | [**[]LibraryItemData**](LibraryItemData.md) | Styles that were created by the library publish | 
**CreatedVariables** | [**[]LibraryItemData**](LibraryItemData.md) | Variables that were created by the library publish | 
**ModifiedComponents** | [**[]LibraryItemData**](LibraryItemData.md) | Components that were modified by the library publish | 
**ModifiedStyles** | [**[]LibraryItemData**](LibraryItemData.md) | Styles that were modified by the library publish | 
**ModifiedVariables** | [**[]LibraryItemData**](LibraryItemData.md) | Variables that were modified by the library publish | 
**DeletedComponents** | [**[]LibraryItemData**](LibraryItemData.md) | Components that were deleted by the library publish | 
**DeletedStyles** | [**[]LibraryItemData**](LibraryItemData.md) | Styles that were deleted by the library publish | 
**DeletedVariables** | [**[]LibraryItemData**](LibraryItemData.md) | Variables that were deleted by the library publish | 
**Description** | Pointer to **string** | Description of the library publish | [optional] 
**FileKey** | **string** | The key of the file that was published | 
**FileName** | **string** | The name of the file that was published | 
**LibraryItem** | [**LibraryItemData**](LibraryItemData.md) | The library item that was published | 
**TriggeredBy** | [**User**](User.md) | The user that published the library and triggered this event | 

## Methods

### NewWebhookLibraryPublishPayload

`func NewWebhookLibraryPublishPayload(passcode string, timestamp time.Time, webhookId string, eventType string, createdComponents []LibraryItemData, createdStyles []LibraryItemData, createdVariables []LibraryItemData, modifiedComponents []LibraryItemData, modifiedStyles []LibraryItemData, modifiedVariables []LibraryItemData, deletedComponents []LibraryItemData, deletedStyles []LibraryItemData, deletedVariables []LibraryItemData, fileKey string, fileName string, libraryItem LibraryItemData, triggeredBy User, ) *WebhookLibraryPublishPayload`

NewWebhookLibraryPublishPayload instantiates a new WebhookLibraryPublishPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookLibraryPublishPayloadWithDefaults

`func NewWebhookLibraryPublishPayloadWithDefaults() *WebhookLibraryPublishPayload`

NewWebhookLibraryPublishPayloadWithDefaults instantiates a new WebhookLibraryPublishPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPasscode

`func (o *WebhookLibraryPublishPayload) GetPasscode() string`

GetPasscode returns the Passcode field if non-nil, zero value otherwise.

### GetPasscodeOk

`func (o *WebhookLibraryPublishPayload) GetPasscodeOk() (*string, bool)`

GetPasscodeOk returns a tuple with the Passcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasscode

`func (o *WebhookLibraryPublishPayload) SetPasscode(v string)`

SetPasscode sets Passcode field to given value.


### GetTimestamp

`func (o *WebhookLibraryPublishPayload) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookLibraryPublishPayload) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookLibraryPublishPayload) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.


### GetWebhookId

`func (o *WebhookLibraryPublishPayload) GetWebhookId() string`

GetWebhookId returns the WebhookId field if non-nil, zero value otherwise.

### GetWebhookIdOk

`func (o *WebhookLibraryPublishPayload) GetWebhookIdOk() (*string, bool)`

GetWebhookIdOk returns a tuple with the WebhookId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookId

`func (o *WebhookLibraryPublishPayload) SetWebhookId(v string)`

SetWebhookId sets WebhookId field to given value.


### GetEventType

`func (o *WebhookLibraryPublishPayload) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *WebhookLibraryPublishPayload) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *WebhookLibraryPublishPayload) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetCreatedComponents

`func (o *WebhookLibraryPublishPayload) GetCreatedComponents() []LibraryItemData`

GetCreatedComponents returns the CreatedComponents field if non-nil, zero value otherwise.

### GetCreatedComponentsOk

`func (o *WebhookLibraryPublishPayload) GetCreatedComponentsOk() (*[]LibraryItemData, bool)`

GetCreatedComponentsOk returns a tuple with the CreatedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedComponents

`func (o *WebhookLibraryPublishPayload) SetCreatedComponents(v []LibraryItemData)`

SetCreatedComponents sets CreatedComponents field to given value.


### GetCreatedStyles

`func (o *WebhookLibraryPublishPayload) GetCreatedStyles() []LibraryItemData`

GetCreatedStyles returns the CreatedStyles field if non-nil, zero value otherwise.

### GetCreatedStylesOk

`func (o *WebhookLibraryPublishPayload) GetCreatedStylesOk() (*[]LibraryItemData, bool)`

GetCreatedStylesOk returns a tuple with the CreatedStyles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedStyles

`func (o *WebhookLibraryPublishPayload) SetCreatedStyles(v []LibraryItemData)`

SetCreatedStyles sets CreatedStyles field to given value.


### GetCreatedVariables

`func (o *WebhookLibraryPublishPayload) GetCreatedVariables() []LibraryItemData`

GetCreatedVariables returns the CreatedVariables field if non-nil, zero value otherwise.

### GetCreatedVariablesOk

`func (o *WebhookLibraryPublishPayload) GetCreatedVariablesOk() (*[]LibraryItemData, bool)`

GetCreatedVariablesOk returns a tuple with the CreatedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedVariables

`func (o *WebhookLibraryPublishPayload) SetCreatedVariables(v []LibraryItemData)`

SetCreatedVariables sets CreatedVariables field to given value.


### GetModifiedComponents

`func (o *WebhookLibraryPublishPayload) GetModifiedComponents() []LibraryItemData`

GetModifiedComponents returns the ModifiedComponents field if non-nil, zero value otherwise.

### GetModifiedComponentsOk

`func (o *WebhookLibraryPublishPayload) GetModifiedComponentsOk() (*[]LibraryItemData, bool)`

GetModifiedComponentsOk returns a tuple with the ModifiedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedComponents

`func (o *WebhookLibraryPublishPayload) SetModifiedComponents(v []LibraryItemData)`

SetModifiedComponents sets ModifiedComponents field to given value.


### GetModifiedStyles

`func (o *WebhookLibraryPublishPayload) GetModifiedStyles() []LibraryItemData`

GetModifiedStyles returns the ModifiedStyles field if non-nil, zero value otherwise.

### GetModifiedStylesOk

`func (o *WebhookLibraryPublishPayload) GetModifiedStylesOk() (*[]LibraryItemData, bool)`

GetModifiedStylesOk returns a tuple with the ModifiedStyles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedStyles

`func (o *WebhookLibraryPublishPayload) SetModifiedStyles(v []LibraryItemData)`

SetModifiedStyles sets ModifiedStyles field to given value.


### GetModifiedVariables

`func (o *WebhookLibraryPublishPayload) GetModifiedVariables() []LibraryItemData`

GetModifiedVariables returns the ModifiedVariables field if non-nil, zero value otherwise.

### GetModifiedVariablesOk

`func (o *WebhookLibraryPublishPayload) GetModifiedVariablesOk() (*[]LibraryItemData, bool)`

GetModifiedVariablesOk returns a tuple with the ModifiedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModifiedVariables

`func (o *WebhookLibraryPublishPayload) SetModifiedVariables(v []LibraryItemData)`

SetModifiedVariables sets ModifiedVariables field to given value.


### GetDeletedComponents

`func (o *WebhookLibraryPublishPayload) GetDeletedComponents() []LibraryItemData`

GetDeletedComponents returns the DeletedComponents field if non-nil, zero value otherwise.

### GetDeletedComponentsOk

`func (o *WebhookLibraryPublishPayload) GetDeletedComponentsOk() (*[]LibraryItemData, bool)`

GetDeletedComponentsOk returns a tuple with the DeletedComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedComponents

`func (o *WebhookLibraryPublishPayload) SetDeletedComponents(v []LibraryItemData)`

SetDeletedComponents sets DeletedComponents field to given value.


### GetDeletedStyles

`func (o *WebhookLibraryPublishPayload) GetDeletedStyles() []LibraryItemData`

GetDeletedStyles returns the DeletedStyles field if non-nil, zero value otherwise.

### GetDeletedStylesOk

`func (o *WebhookLibraryPublishPayload) GetDeletedStylesOk() (*[]LibraryItemData, bool)`

GetDeletedStylesOk returns a tuple with the DeletedStyles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedStyles

`func (o *WebhookLibraryPublishPayload) SetDeletedStyles(v []LibraryItemData)`

SetDeletedStyles sets DeletedStyles field to given value.


### GetDeletedVariables

`func (o *WebhookLibraryPublishPayload) GetDeletedVariables() []LibraryItemData`

GetDeletedVariables returns the DeletedVariables field if non-nil, zero value otherwise.

### GetDeletedVariablesOk

`func (o *WebhookLibraryPublishPayload) GetDeletedVariablesOk() (*[]LibraryItemData, bool)`

GetDeletedVariablesOk returns a tuple with the DeletedVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedVariables

`func (o *WebhookLibraryPublishPayload) SetDeletedVariables(v []LibraryItemData)`

SetDeletedVariables sets DeletedVariables field to given value.


### GetDescription

`func (o *WebhookLibraryPublishPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WebhookLibraryPublishPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WebhookLibraryPublishPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WebhookLibraryPublishPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFileKey

`func (o *WebhookLibraryPublishPayload) GetFileKey() string`

GetFileKey returns the FileKey field if non-nil, zero value otherwise.

### GetFileKeyOk

`func (o *WebhookLibraryPublishPayload) GetFileKeyOk() (*string, bool)`

GetFileKeyOk returns a tuple with the FileKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileKey

`func (o *WebhookLibraryPublishPayload) SetFileKey(v string)`

SetFileKey sets FileKey field to given value.


### GetFileName

`func (o *WebhookLibraryPublishPayload) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *WebhookLibraryPublishPayload) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *WebhookLibraryPublishPayload) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetLibraryItem

`func (o *WebhookLibraryPublishPayload) GetLibraryItem() LibraryItemData`

GetLibraryItem returns the LibraryItem field if non-nil, zero value otherwise.

### GetLibraryItemOk

`func (o *WebhookLibraryPublishPayload) GetLibraryItemOk() (*LibraryItemData, bool)`

GetLibraryItemOk returns a tuple with the LibraryItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibraryItem

`func (o *WebhookLibraryPublishPayload) SetLibraryItem(v LibraryItemData)`

SetLibraryItem sets LibraryItem field to given value.


### GetTriggeredBy

`func (o *WebhookLibraryPublishPayload) GetTriggeredBy() User`

GetTriggeredBy returns the TriggeredBy field if non-nil, zero value otherwise.

### GetTriggeredByOk

`func (o *WebhookLibraryPublishPayload) GetTriggeredByOk() (*User, bool)`

GetTriggeredByOk returns a tuple with the TriggeredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggeredBy

`func (o *WebhookLibraryPublishPayload) SetTriggeredBy(v User)`

SetTriggeredBy sets TriggeredBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


