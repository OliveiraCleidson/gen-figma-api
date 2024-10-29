# FrameOffsetRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | **string** | Unique id specifying the frame. | 
**NodeOffset** | [**Vector**](Vector.md) | 2D vector offset within the frame from the top-left corner. | 
**RegionHeight** | **float32** | The height of the comment region. Must be greater than 0. | 
**RegionWidth** | **float32** | The width of the comment region. Must be greater than 0. | 
**CommentPinCorner** | Pointer to **string** | The corner of the comment region to pin to the node&#39;s corner as a string enum. | [optional] [default to "bottom-right"]

## Methods

### NewFrameOffsetRegion

`func NewFrameOffsetRegion(nodeId string, nodeOffset Vector, regionHeight float32, regionWidth float32, ) *FrameOffsetRegion`

NewFrameOffsetRegion instantiates a new FrameOffsetRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameOffsetRegionWithDefaults

`func NewFrameOffsetRegionWithDefaults() *FrameOffsetRegion`

NewFrameOffsetRegionWithDefaults instantiates a new FrameOffsetRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *FrameOffsetRegion) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *FrameOffsetRegion) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *FrameOffsetRegion) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeOffset

`func (o *FrameOffsetRegion) GetNodeOffset() Vector`

GetNodeOffset returns the NodeOffset field if non-nil, zero value otherwise.

### GetNodeOffsetOk

`func (o *FrameOffsetRegion) GetNodeOffsetOk() (*Vector, bool)`

GetNodeOffsetOk returns a tuple with the NodeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeOffset

`func (o *FrameOffsetRegion) SetNodeOffset(v Vector)`

SetNodeOffset sets NodeOffset field to given value.


### GetRegionHeight

`func (o *FrameOffsetRegion) GetRegionHeight() float32`

GetRegionHeight returns the RegionHeight field if non-nil, zero value otherwise.

### GetRegionHeightOk

`func (o *FrameOffsetRegion) GetRegionHeightOk() (*float32, bool)`

GetRegionHeightOk returns a tuple with the RegionHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionHeight

`func (o *FrameOffsetRegion) SetRegionHeight(v float32)`

SetRegionHeight sets RegionHeight field to given value.


### GetRegionWidth

`func (o *FrameOffsetRegion) GetRegionWidth() float32`

GetRegionWidth returns the RegionWidth field if non-nil, zero value otherwise.

### GetRegionWidthOk

`func (o *FrameOffsetRegion) GetRegionWidthOk() (*float32, bool)`

GetRegionWidthOk returns a tuple with the RegionWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionWidth

`func (o *FrameOffsetRegion) SetRegionWidth(v float32)`

SetRegionWidth sets RegionWidth field to given value.


### GetCommentPinCorner

`func (o *FrameOffsetRegion) GetCommentPinCorner() string`

GetCommentPinCorner returns the CommentPinCorner field if non-nil, zero value otherwise.

### GetCommentPinCornerOk

`func (o *FrameOffsetRegion) GetCommentPinCornerOk() (*string, bool)`

GetCommentPinCornerOk returns a tuple with the CommentPinCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentPinCorner

`func (o *FrameOffsetRegion) SetCommentPinCorner(v string)`

SetCommentPinCorner sets CommentPinCorner field to given value.

### HasCommentPinCorner

`func (o *FrameOffsetRegion) HasCommentPinCorner() bool`

HasCommentPinCorner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


