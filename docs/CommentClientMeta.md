# CommentClientMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**X** | **float32** | X coordinate of the position. | 
**Y** | **float32** | Y coordinate of the position. | 
**NodeId** | **string** | Unique id specifying the frame. | 
**NodeOffset** | [**Vector**](Vector.md) | 2D vector offset within the frame from the top-left corner. | 
**RegionHeight** | **float32** | The height of the comment region. Must be greater than 0. | 
**RegionWidth** | **float32** | The width of the comment region. Must be greater than 0. | 
**CommentPinCorner** | Pointer to **string** | The corner of the comment region to pin to the node&#39;s corner as a string enum. | [optional] [default to "bottom-right"]

## Methods

### NewCommentClientMeta

`func NewCommentClientMeta(x float32, y float32, nodeId string, nodeOffset Vector, regionHeight float32, regionWidth float32, ) *CommentClientMeta`

NewCommentClientMeta instantiates a new CommentClientMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommentClientMetaWithDefaults

`func NewCommentClientMetaWithDefaults() *CommentClientMeta`

NewCommentClientMetaWithDefaults instantiates a new CommentClientMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetX

`func (o *CommentClientMeta) GetX() float32`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *CommentClientMeta) GetXOk() (*float32, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *CommentClientMeta) SetX(v float32)`

SetX sets X field to given value.


### GetY

`func (o *CommentClientMeta) GetY() float32`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *CommentClientMeta) GetYOk() (*float32, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *CommentClientMeta) SetY(v float32)`

SetY sets Y field to given value.


### GetNodeId

`func (o *CommentClientMeta) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *CommentClientMeta) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *CommentClientMeta) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.


### GetNodeOffset

`func (o *CommentClientMeta) GetNodeOffset() Vector`

GetNodeOffset returns the NodeOffset field if non-nil, zero value otherwise.

### GetNodeOffsetOk

`func (o *CommentClientMeta) GetNodeOffsetOk() (*Vector, bool)`

GetNodeOffsetOk returns a tuple with the NodeOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeOffset

`func (o *CommentClientMeta) SetNodeOffset(v Vector)`

SetNodeOffset sets NodeOffset field to given value.


### GetRegionHeight

`func (o *CommentClientMeta) GetRegionHeight() float32`

GetRegionHeight returns the RegionHeight field if non-nil, zero value otherwise.

### GetRegionHeightOk

`func (o *CommentClientMeta) GetRegionHeightOk() (*float32, bool)`

GetRegionHeightOk returns a tuple with the RegionHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionHeight

`func (o *CommentClientMeta) SetRegionHeight(v float32)`

SetRegionHeight sets RegionHeight field to given value.


### GetRegionWidth

`func (o *CommentClientMeta) GetRegionWidth() float32`

GetRegionWidth returns the RegionWidth field if non-nil, zero value otherwise.

### GetRegionWidthOk

`func (o *CommentClientMeta) GetRegionWidthOk() (*float32, bool)`

GetRegionWidthOk returns a tuple with the RegionWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionWidth

`func (o *CommentClientMeta) SetRegionWidth(v float32)`

SetRegionWidth sets RegionWidth field to given value.


### GetCommentPinCorner

`func (o *CommentClientMeta) GetCommentPinCorner() string`

GetCommentPinCorner returns the CommentPinCorner field if non-nil, zero value otherwise.

### GetCommentPinCornerOk

`func (o *CommentClientMeta) GetCommentPinCornerOk() (*string, bool)`

GetCommentPinCornerOk returns a tuple with the CommentPinCorner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentPinCorner

`func (o *CommentClientMeta) SetCommentPinCorner(v string)`

SetCommentPinCorner sets CommentPinCorner field to given value.

### HasCommentPinCorner

`func (o *CommentClientMeta) HasCommentPinCorner() bool`

HasCommentPinCorner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


