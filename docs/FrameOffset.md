# FrameOffset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** | Unique id specifying the frame. | 
**NodeOffset** | [**Vector**](Vector.md) | 2D vector offset within the frame from the top-left corner. | 

## Methods

### NewFrameOffset

`func NewFrameOffset(nodeId string, nodeOffset Vector, ) *FrameOffset`

NewFrameOffset instantiates a new FrameOffset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameOffsetWithDefaults

`func NewFrameOffsetWithDefaults() *FrameOffset`

NewFrameOffsetWithDefaults instantiates a new FrameOffset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *FrameOffset) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *FrameOffset) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *FrameOffset) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeOffset

`func (o *FrameOffset) GetNodeOffset() Vector`

GetNodeOffset returns the NodeOffset field if non-nil, zero value otherwise.

### GetNodeOffsetOk

`func (o *FrameOffset) GetNodeOffsetOk() (*Vector, bool)`

GetNodeOffsetOk returns a tuple with the NodeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeOffset

`func (o *FrameOffset) SetNodeOffset(v Vector)`

SetNodeOffset sets NodeOffset field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


