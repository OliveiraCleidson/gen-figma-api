# Action

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Url** | **string** |  | 
**DestinationId** | **string** |  | 
**MediaAction** | **string** |  | 
**AmountToSkip** | **float32** |  | 
**NewTimestamp** | **float32** |  | 
**VariableId** | **string** |  | 
**VariableValue** | Pointer to [**VariableData**](VariableData.md) |  | [optional] 
**VariableCollectionId** | Pointer to **string** |  | [optional] 
**VariableModeId** | Pointer to **string** |  | [optional] 
**ConditionalBlocks** | [**[]ConditionalBlock**](ConditionalBlock.md) |  | 
**Navigation** | [**Navigation**](Navigation.md) |  | 
**Transition** | [**Transition**](Transition.md) |  | 
**PreserveScrollPosition** | Pointer to **bool** | Whether the scroll offsets of any scrollable elements in the current screen or overlay are preserved when navigating to the destination. This is applicable only if the layout of both the current frame and its destination are the same. | [optional] 
**OverlayRelativePosition** | Pointer to [**Vector**](Vector.md) | Applicable only when &#x60;navigation&#x60; is &#x60;\&quot;OVERLAY\&quot;&#x60; and the destination is a frame with &#x60;overlayPosition&#x60; equal to &#x60;\&quot;MANUAL\&quot;&#x60;. This value represents the offset by which the overlay is opened relative to this node. | [optional] 
**ResetVideoPosition** | Pointer to **bool** | When true, all videos within the destination frame will reset their memorized playback position to 00:00 before starting to play. | [optional] 
**ResetScrollPosition** | Pointer to **bool** | Whether the scroll offsets of any scrollable elements in the current screen or overlay reset when navigating to the destination. This is applicable only if the layout of both the current frame and its destination are the same. | [optional] 
**ResetInteractiveComponents** | Pointer to **bool** | Whether the state of any interactive components in the current screen or overlay reset when navigating to the destination. This is applicable if there are interactive components in the destination frame. | [optional] 

## Methods

### NewAction

`func NewAction(type_ string, url string, destinationId string, mediaAction string, amountToSkip float32, newTimestamp float32, variableId string, conditionalBlocks []ConditionalBlock, navigation Navigation, transition Transition, ) *Action`

NewAction instantiates a new Action object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionWithDefaults

`func NewActionWithDefaults() *Action`

NewActionWithDefaults instantiates a new Action object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Action) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Action) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Action) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *Action) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Action) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Action) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetDestinationId

`func (o *Action) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *Action) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *Action) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetMediaAction

`func (o *Action) GetMediaAction() string`

GetMediaAction returns the MediaAction field if non-nil, zero value otherwise.

### GetMediaActionOk

`func (o *Action) GetMediaActionOk() (*string, bool)`

GetMediaActionOk returns a tuple with the MediaAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaAction

`func (o *Action) SetMediaAction(v string)`

SetMediaAction sets MediaAction field to given value.


### GetAmountToSkip

`func (o *Action) GetAmountToSkip() float32`

GetAmountToSkip returns the AmountToSkip field if non-nil, zero value otherwise.

### GetAmountToSkipOk

`func (o *Action) GetAmountToSkipOk() (*float32, bool)`

GetAmountToSkipOk returns a tuple with the AmountToSkip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountToSkip

`func (o *Action) SetAmountToSkip(v float32)`

SetAmountToSkip sets AmountToSkip field to given value.


### GetNewTimestamp

`func (o *Action) GetNewTimestamp() float32`

GetNewTimestamp returns the NewTimestamp field if non-nil, zero value otherwise.

### GetNewTimestampOk

`func (o *Action) GetNewTimestampOk() (*float32, bool)`

GetNewTimestampOk returns a tuple with the NewTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewTimestamp

`func (o *Action) SetNewTimestamp(v float32)`

SetNewTimestamp sets NewTimestamp field to given value.


### GetVariableId

`func (o *Action) GetVariableId() string`

GetVariableId returns the VariableId field if non-nil, zero value otherwise.

### GetVariableIdOk

`func (o *Action) GetVariableIdOk() (*string, bool)`

GetVariableIdOk returns a tuple with the VariableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableId

`func (o *Action) SetVariableId(v string)`

SetVariableId sets VariableId field to given value.


### GetVariableValue

`func (o *Action) GetVariableValue() VariableData`

GetVariableValue returns the VariableValue field if non-nil, zero value otherwise.

### GetVariableValueOk

`func (o *Action) GetVariableValueOk() (*VariableData, bool)`

GetVariableValueOk returns a tuple with the VariableValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableValue

`func (o *Action) SetVariableValue(v VariableData)`

SetVariableValue sets VariableValue field to given value.

### HasVariableValue

`func (o *Action) HasVariableValue() bool`

HasVariableValue returns a boolean if a field has been set.

### GetVariableCollectionId

`func (o *Action) GetVariableCollectionId() string`

GetVariableCollectionId returns the VariableCollectionId field if non-nil, zero value otherwise.

### GetVariableCollectionIdOk

`func (o *Action) GetVariableCollectionIdOk() (*string, bool)`

GetVariableCollectionIdOk returns a tuple with the VariableCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableCollectionId

`func (o *Action) SetVariableCollectionId(v string)`

SetVariableCollectionId sets VariableCollectionId field to given value.

### HasVariableCollectionId

`func (o *Action) HasVariableCollectionId() bool`

HasVariableCollectionId returns a boolean if a field has been set.

### GetVariableModeId

`func (o *Action) GetVariableModeId() string`

GetVariableModeId returns the VariableModeId field if non-nil, zero value otherwise.

### GetVariableModeIdOk

`func (o *Action) GetVariableModeIdOk() (*string, bool)`

GetVariableModeIdOk returns a tuple with the VariableModeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableModeId

`func (o *Action) SetVariableModeId(v string)`

SetVariableModeId sets VariableModeId field to given value.

### HasVariableModeId

`func (o *Action) HasVariableModeId() bool`

HasVariableModeId returns a boolean if a field has been set.

### GetConditionalBlocks

`func (o *Action) GetConditionalBlocks() []ConditionalBlock`

GetConditionalBlocks returns the ConditionalBlocks field if non-nil, zero value otherwise.

### GetConditionalBlocksOk

`func (o *Action) GetConditionalBlocksOk() (*[]ConditionalBlock, bool)`

GetConditionalBlocksOk returns a tuple with the ConditionalBlocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionalBlocks

`func (o *Action) SetConditionalBlocks(v []ConditionalBlock)`

SetConditionalBlocks sets ConditionalBlocks field to given value.


### GetNavigation

`func (o *Action) GetNavigation() Navigation`

GetNavigation returns the Navigation field if non-nil, zero value otherwise.

### GetNavigationOk

`func (o *Action) GetNavigationOk() (*Navigation, bool)`

GetNavigationOk returns a tuple with the Navigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigation

`func (o *Action) SetNavigation(v Navigation)`

SetNavigation sets Navigation field to given value.


### GetTransition

`func (o *Action) GetTransition() Transition`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *Action) GetTransitionOk() (*Transition, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *Action) SetTransition(v Transition)`

SetTransition sets Transition field to given value.


### GetPreserveScrollPosition

`func (o *Action) GetPreserveScrollPosition() bool`

GetPreserveScrollPosition returns the PreserveScrollPosition field if non-nil, zero value otherwise.

### GetPreserveScrollPositionOk

`func (o *Action) GetPreserveScrollPositionOk() (*bool, bool)`

GetPreserveScrollPositionOk returns a tuple with the PreserveScrollPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveScrollPosition

`func (o *Action) SetPreserveScrollPosition(v bool)`

SetPreserveScrollPosition sets PreserveScrollPosition field to given value.

### HasPreserveScrollPosition

`func (o *Action) HasPreserveScrollPosition() bool`

HasPreserveScrollPosition returns a boolean if a field has been set.

### GetOverlayRelativePosition

`func (o *Action) GetOverlayRelativePosition() Vector`

GetOverlayRelativePosition returns the OverlayRelativePosition field if non-nil, zero value otherwise.

### GetOverlayRelativePositionOk

`func (o *Action) GetOverlayRelativePositionOk() (*Vector, bool)`

GetOverlayRelativePositionOk returns a tuple with the OverlayRelativePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayRelativePosition

`func (o *Action) SetOverlayRelativePosition(v Vector)`

SetOverlayRelativePosition sets OverlayRelativePosition field to given value.

### HasOverlayRelativePosition

`func (o *Action) HasOverlayRelativePosition() bool`

HasOverlayRelativePosition returns a boolean if a field has been set.

### GetResetVideoPosition

`func (o *Action) GetResetVideoPosition() bool`

GetResetVideoPosition returns the ResetVideoPosition field if non-nil, zero value otherwise.

### GetResetVideoPositionOk

`func (o *Action) GetResetVideoPositionOk() (*bool, bool)`

GetResetVideoPositionOk returns a tuple with the ResetVideoPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetVideoPosition

`func (o *Action) SetResetVideoPosition(v bool)`

SetResetVideoPosition sets ResetVideoPosition field to given value.

### HasResetVideoPosition

`func (o *Action) HasResetVideoPosition() bool`

HasResetVideoPosition returns a boolean if a field has been set.

### GetResetScrollPosition

`func (o *Action) GetResetScrollPosition() bool`

GetResetScrollPosition returns the ResetScrollPosition field if non-nil, zero value otherwise.

### GetResetScrollPositionOk

`func (o *Action) GetResetScrollPositionOk() (*bool, bool)`

GetResetScrollPositionOk returns a tuple with the ResetScrollPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetScrollPosition

`func (o *Action) SetResetScrollPosition(v bool)`

SetResetScrollPosition sets ResetScrollPosition field to given value.

### HasResetScrollPosition

`func (o *Action) HasResetScrollPosition() bool`

HasResetScrollPosition returns a boolean if a field has been set.

### GetResetInteractiveComponents

`func (o *Action) GetResetInteractiveComponents() bool`

GetResetInteractiveComponents returns the ResetInteractiveComponents field if non-nil, zero value otherwise.

### GetResetInteractiveComponentsOk

`func (o *Action) GetResetInteractiveComponentsOk() (*bool, bool)`

GetResetInteractiveComponentsOk returns a tuple with the ResetInteractiveComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetInteractiveComponents

`func (o *Action) SetResetInteractiveComponents(v bool)`

SetResetInteractiveComponents sets ResetInteractiveComponents field to given value.

### HasResetInteractiveComponents

`func (o *Action) HasResetInteractiveComponents() bool`

HasResetInteractiveComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


