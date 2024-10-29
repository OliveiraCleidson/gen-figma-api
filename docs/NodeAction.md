# NodeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**DestinationId** | **NullableString** |  | 
**Navigation** | [**Navigation**](Navigation.md) |  | 
**Transition** | [**NullableTransition**](Transition.md) |  | 
**PreserveScrollPosition** | Pointer to **bool** | Whether the scroll offsets of any scrollable elements in the current screen or overlay are preserved when navigating to the destination. This is applicable only if the layout of both the current frame and its destination are the same. | [optional] 
**OverlayRelativePosition** | Pointer to [**Vector**](Vector.md) | Applicable only when &#x60;navigation&#x60; is &#x60;\&quot;OVERLAY\&quot;&#x60; and the destination is a frame with &#x60;overlayPosition&#x60; equal to &#x60;\&quot;MANUAL\&quot;&#x60;. This value represents the offset by which the overlay is opened relative to this node. | [optional] 
**ResetVideoPosition** | Pointer to **bool** | When true, all videos within the destination frame will reset their memorized playback position to 00:00 before starting to play. | [optional] 
**ResetScrollPosition** | Pointer to **bool** | Whether the scroll offsets of any scrollable elements in the current screen or overlay reset when navigating to the destination. This is applicable only if the layout of both the current frame and its destination are the same. | [optional] 
**ResetInteractiveComponents** | Pointer to **bool** | Whether the state of any interactive components in the current screen or overlay reset when navigating to the destination. This is applicable if there are interactive components in the destination frame. | [optional] 

## Methods

### NewNodeAction

`func NewNodeAction(type_ string, destinationId NullableString, navigation Navigation, transition NullableTransition, ) *NodeAction`

NewNodeAction instantiates a new NodeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNodeActionWithDefaults

`func NewNodeActionWithDefaults() *NodeAction`

NewNodeActionWithDefaults instantiates a new NodeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NodeAction) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NodeAction) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NodeAction) SetType(v string)`

SetType sets Type field to given value.


### GetDestinationId

`func (o *NodeAction) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *NodeAction) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *NodeAction) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### SetDestinationIdNil

`func (o *NodeAction) SetDestinationIdNil(b bool)`

 SetDestinationIdNil sets the value for DestinationId to be an explicit nil

### UnsetDestinationId
`func (o *NodeAction) UnsetDestinationId()`

UnsetDestinationId ensures that no value is present for DestinationId, not even an explicit nil
### GetNavigation

`func (o *NodeAction) GetNavigation() Navigation`

GetNavigation returns the Navigation field if non-nil, zero value otherwise.

### GetNavigationOk

`func (o *NodeAction) GetNavigationOk() (*Navigation, bool)`

GetNavigationOk returns a tuple with the Navigation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigation

`func (o *NodeAction) SetNavigation(v Navigation)`

SetNavigation sets Navigation field to given value.


### GetTransition

`func (o *NodeAction) GetTransition() Transition`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *NodeAction) GetTransitionOk() (*Transition, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *NodeAction) SetTransition(v Transition)`

SetTransition sets Transition field to given value.


### SetTransitionNil

`func (o *NodeAction) SetTransitionNil(b bool)`

 SetTransitionNil sets the value for Transition to be an explicit nil

### UnsetTransition
`func (o *NodeAction) UnsetTransition()`

UnsetTransition ensures that no value is present for Transition, not even an explicit nil
### GetPreserveScrollPosition

`func (o *NodeAction) GetPreserveScrollPosition() bool`

GetPreserveScrollPosition returns the PreserveScrollPosition field if non-nil, zero value otherwise.

### GetPreserveScrollPositionOk

`func (o *NodeAction) GetPreserveScrollPositionOk() (*bool, bool)`

GetPreserveScrollPositionOk returns a tuple with the PreserveScrollPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreserveScrollPosition

`func (o *NodeAction) SetPreserveScrollPosition(v bool)`

SetPreserveScrollPosition sets PreserveScrollPosition field to given value.

### HasPreserveScrollPosition

`func (o *NodeAction) HasPreserveScrollPosition() bool`

HasPreserveScrollPosition returns a boolean if a field has been set.

### GetOverlayRelativePosition

`func (o *NodeAction) GetOverlayRelativePosition() Vector`

GetOverlayRelativePosition returns the OverlayRelativePosition field if non-nil, zero value otherwise.

### GetOverlayRelativePositionOk

`func (o *NodeAction) GetOverlayRelativePositionOk() (*Vector, bool)`

GetOverlayRelativePositionOk returns a tuple with the OverlayRelativePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverlayRelativePosition

`func (o *NodeAction) SetOverlayRelativePosition(v Vector)`

SetOverlayRelativePosition sets OverlayRelativePosition field to given value.

### HasOverlayRelativePosition

`func (o *NodeAction) HasOverlayRelativePosition() bool`

HasOverlayRelativePosition returns a boolean if a field has been set.

### GetResetVideoPosition

`func (o *NodeAction) GetResetVideoPosition() bool`

GetResetVideoPosition returns the ResetVideoPosition field if non-nil, zero value otherwise.

### GetResetVideoPositionOk

`func (o *NodeAction) GetResetVideoPositionOk() (*bool, bool)`

GetResetVideoPositionOk returns a tuple with the ResetVideoPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetVideoPosition

`func (o *NodeAction) SetResetVideoPosition(v bool)`

SetResetVideoPosition sets ResetVideoPosition field to given value.

### HasResetVideoPosition

`func (o *NodeAction) HasResetVideoPosition() bool`

HasResetVideoPosition returns a boolean if a field has been set.

### GetResetScrollPosition

`func (o *NodeAction) GetResetScrollPosition() bool`

GetResetScrollPosition returns the ResetScrollPosition field if non-nil, zero value otherwise.

### GetResetScrollPositionOk

`func (o *NodeAction) GetResetScrollPositionOk() (*bool, bool)`

GetResetScrollPositionOk returns a tuple with the ResetScrollPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetScrollPosition

`func (o *NodeAction) SetResetScrollPosition(v bool)`

SetResetScrollPosition sets ResetScrollPosition field to given value.

### HasResetScrollPosition

`func (o *NodeAction) HasResetScrollPosition() bool`

HasResetScrollPosition returns a boolean if a field has been set.

### GetResetInteractiveComponents

`func (o *NodeAction) GetResetInteractiveComponents() bool`

GetResetInteractiveComponents returns the ResetInteractiveComponents field if non-nil, zero value otherwise.

### GetResetInteractiveComponentsOk

`func (o *NodeAction) GetResetInteractiveComponentsOk() (*bool, bool)`

GetResetInteractiveComponentsOk returns a tuple with the ResetInteractiveComponents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetInteractiveComponents

`func (o *NodeAction) SetResetInteractiveComponents(v bool)`

SetResetInteractiveComponents sets ResetInteractiveComponents field to given value.

### HasResetInteractiveComponents

`func (o *NodeAction) HasResetInteractiveComponents() bool`

HasResetInteractiveComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


