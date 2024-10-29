# HasFramePropertiesTrait

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClipsContent** | **bool** | Whether or not this node clip content outside of its bounds | 
**Background** | Pointer to [**[]Paint**](Paint.md) | Background of the node. This is deprecated, as backgrounds for frames are now in the &#x60;fills&#x60; field. | [optional] 
**BackgroundColor** | Pointer to [**RGBA**](RGBA.md) | Background color of the node. This is deprecated, as frames now support more than a solid color as a background. Please use the &#x60;fills&#x60; field instead. | [optional] 
**LayoutGrids** | Pointer to [**[]LayoutGrid**](LayoutGrid.md) | An array of layout grids attached to this node (see layout grids section for more details). GROUP nodes do not have this attribute | [optional] 
**OverflowDirection** | Pointer to **string** | Whether a node has primary axis scrolling, horizontal or vertical. | [optional] [default to "NONE"]
**LayoutMode** | Pointer to **string** | Whether this layer uses auto-layout to position its children. | [optional] [default to "NONE"]
**PrimaryAxisSizingMode** | Pointer to **string** | Whether the primary axis has a fixed length (determined by the user) or an automatic length (determined by the layout engine). This property is only applicable for auto-layout frames. | [optional] [default to "AUTO"]
**CounterAxisSizingMode** | Pointer to **string** | Whether the counter axis has a fixed length (determined by the user) or an automatic length (determined by the layout engine). This property is only applicable for auto-layout frames. | [optional] [default to "AUTO"]
**PrimaryAxisAlignItems** | Pointer to **string** | Determines how the auto-layout frame&#39;s children should be aligned in the primary axis direction. This property is only applicable for auto-layout frames. | [optional] [default to "MIN"]
**CounterAxisAlignItems** | Pointer to **string** | Determines how the auto-layout frame&#39;s children should be aligned in the counter axis direction. This property is only applicable for auto-layout frames. | [optional] [default to "MIN"]
**PaddingLeft** | Pointer to **float32** | The padding between the left border of the frame and its children. This property is only applicable for auto-layout frames. | [optional] [default to 0]
**PaddingRight** | Pointer to **float32** | The padding between the right border of the frame and its children. This property is only applicable for auto-layout frames. | [optional] [default to 0]
**PaddingTop** | Pointer to **float32** | The padding between the top border of the frame and its children. This property is only applicable for auto-layout frames. | [optional] [default to 0]
**PaddingBottom** | Pointer to **float32** | The padding between the bottom border of the frame and its children. This property is only applicable for auto-layout frames. | [optional] [default to 0]
**ItemSpacing** | Pointer to **float32** | The distance between children of the frame. Can be negative. This property is only applicable for auto-layout frames. | [optional] [default to 0]
**ItemReverseZIndex** | Pointer to **bool** | Determines the canvas stacking order of layers in this frame. When true, the first layer will be draw on top. This property is only applicable for auto-layout frames. | [optional] [default to false]
**StrokesIncludedInLayout** | Pointer to **bool** | Determines whether strokes are included in layout calculations. When true, auto-layout frames behave like css \&quot;box-sizing: border-box\&quot;. This property is only applicable for auto-layout frames. | [optional] [default to false]
**LayoutWrap** | Pointer to **string** | Whether this auto-layout frame has wrapping enabled. | [optional] 
**CounterAxisSpacing** | Pointer to **float32** | The distance between wrapped tracks of an auto-layout frame. This property is only applicable for auto-layout frames with &#x60;layoutWrap: \&quot;WRAP\&quot;&#x60; | [optional] 
**CounterAxisAlignContent** | Pointer to **string** | Determines how the auto-layout frame’s wrapped tracks should be aligned in the counter axis direction. This property is only applicable for auto-layout frames with &#x60;layoutWrap: \&quot;WRAP\&quot;&#x60;. | [optional] [default to "AUTO"]

## Methods

### NewHasFramePropertiesTrait

`func NewHasFramePropertiesTrait(clipsContent bool, ) *HasFramePropertiesTrait`

NewHasFramePropertiesTrait instantiates a new HasFramePropertiesTrait object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHasFramePropertiesTraitWithDefaults

`func NewHasFramePropertiesTraitWithDefaults() *HasFramePropertiesTrait`

NewHasFramePropertiesTraitWithDefaults instantiates a new HasFramePropertiesTrait object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClipsContent

`func (o *HasFramePropertiesTrait) GetClipsContent() bool`

GetClipsContent returns the ClipsContent field if non-nil, zero value otherwise.

### GetClipsContentOk

`func (o *HasFramePropertiesTrait) GetClipsContentOk() (*bool, bool)`

GetClipsContentOk returns a tuple with the ClipsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipsContent

`func (o *HasFramePropertiesTrait) SetClipsContent(v bool)`

SetClipsContent sets ClipsContent field to given value.


### GetBackground

`func (o *HasFramePropertiesTrait) GetBackground() []Paint`

GetBackground returns the Background field if non-nil, zero value otherwise.

### GetBackgroundOk

`func (o *HasFramePropertiesTrait) GetBackgroundOk() (*[]Paint, bool)`

GetBackgroundOk returns a tuple with the Background field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackground

`func (o *HasFramePropertiesTrait) SetBackground(v []Paint)`

SetBackground sets Background field to given value.

### HasBackground

`func (o *HasFramePropertiesTrait) HasBackground() bool`

HasBackground returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *HasFramePropertiesTrait) GetBackgroundColor() RGBA`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *HasFramePropertiesTrait) GetBackgroundColorOk() (*RGBA, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *HasFramePropertiesTrait) SetBackgroundColor(v RGBA)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *HasFramePropertiesTrait) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetLayoutGrids

`func (o *HasFramePropertiesTrait) GetLayoutGrids() []LayoutGrid`

GetLayoutGrids returns the LayoutGrids field if non-nil, zero value otherwise.

### GetLayoutGridsOk

`func (o *HasFramePropertiesTrait) GetLayoutGridsOk() (*[]LayoutGrid, bool)`

GetLayoutGridsOk returns a tuple with the LayoutGrids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutGrids

`func (o *HasFramePropertiesTrait) SetLayoutGrids(v []LayoutGrid)`

SetLayoutGrids sets LayoutGrids field to given value.

### HasLayoutGrids

`func (o *HasFramePropertiesTrait) HasLayoutGrids() bool`

HasLayoutGrids returns a boolean if a field has been set.

### GetOverflowDirection

`func (o *HasFramePropertiesTrait) GetOverflowDirection() string`

GetOverflowDirection returns the OverflowDirection field if non-nil, zero value otherwise.

### GetOverflowDirectionOk

`func (o *HasFramePropertiesTrait) GetOverflowDirectionOk() (*string, bool)`

GetOverflowDirectionOk returns a tuple with the OverflowDirection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverflowDirection

`func (o *HasFramePropertiesTrait) SetOverflowDirection(v string)`

SetOverflowDirection sets OverflowDirection field to given value.

### HasOverflowDirection

`func (o *HasFramePropertiesTrait) HasOverflowDirection() bool`

HasOverflowDirection returns a boolean if a field has been set.

### GetLayoutMode

`func (o *HasFramePropertiesTrait) GetLayoutMode() string`

GetLayoutMode returns the LayoutMode field if non-nil, zero value otherwise.

### GetLayoutModeOk

`func (o *HasFramePropertiesTrait) GetLayoutModeOk() (*string, bool)`

GetLayoutModeOk returns a tuple with the LayoutMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutMode

`func (o *HasFramePropertiesTrait) SetLayoutMode(v string)`

SetLayoutMode sets LayoutMode field to given value.

### HasLayoutMode

`func (o *HasFramePropertiesTrait) HasLayoutMode() bool`

HasLayoutMode returns a boolean if a field has been set.

### GetPrimaryAxisSizingMode

`func (o *HasFramePropertiesTrait) GetPrimaryAxisSizingMode() string`

GetPrimaryAxisSizingMode returns the PrimaryAxisSizingMode field if non-nil, zero value otherwise.

### GetPrimaryAxisSizingModeOk

`func (o *HasFramePropertiesTrait) GetPrimaryAxisSizingModeOk() (*string, bool)`

GetPrimaryAxisSizingModeOk returns a tuple with the PrimaryAxisSizingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAxisSizingMode

`func (o *HasFramePropertiesTrait) SetPrimaryAxisSizingMode(v string)`

SetPrimaryAxisSizingMode sets PrimaryAxisSizingMode field to given value.

### HasPrimaryAxisSizingMode

`func (o *HasFramePropertiesTrait) HasPrimaryAxisSizingMode() bool`

HasPrimaryAxisSizingMode returns a boolean if a field has been set.

### GetCounterAxisSizingMode

`func (o *HasFramePropertiesTrait) GetCounterAxisSizingMode() string`

GetCounterAxisSizingMode returns the CounterAxisSizingMode field if non-nil, zero value otherwise.

### GetCounterAxisSizingModeOk

`func (o *HasFramePropertiesTrait) GetCounterAxisSizingModeOk() (*string, bool)`

GetCounterAxisSizingModeOk returns a tuple with the CounterAxisSizingMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisSizingMode

`func (o *HasFramePropertiesTrait) SetCounterAxisSizingMode(v string)`

SetCounterAxisSizingMode sets CounterAxisSizingMode field to given value.

### HasCounterAxisSizingMode

`func (o *HasFramePropertiesTrait) HasCounterAxisSizingMode() bool`

HasCounterAxisSizingMode returns a boolean if a field has been set.

### GetPrimaryAxisAlignItems

`func (o *HasFramePropertiesTrait) GetPrimaryAxisAlignItems() string`

GetPrimaryAxisAlignItems returns the PrimaryAxisAlignItems field if non-nil, zero value otherwise.

### GetPrimaryAxisAlignItemsOk

`func (o *HasFramePropertiesTrait) GetPrimaryAxisAlignItemsOk() (*string, bool)`

GetPrimaryAxisAlignItemsOk returns a tuple with the PrimaryAxisAlignItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryAxisAlignItems

`func (o *HasFramePropertiesTrait) SetPrimaryAxisAlignItems(v string)`

SetPrimaryAxisAlignItems sets PrimaryAxisAlignItems field to given value.

### HasPrimaryAxisAlignItems

`func (o *HasFramePropertiesTrait) HasPrimaryAxisAlignItems() bool`

HasPrimaryAxisAlignItems returns a boolean if a field has been set.

### GetCounterAxisAlignItems

`func (o *HasFramePropertiesTrait) GetCounterAxisAlignItems() string`

GetCounterAxisAlignItems returns the CounterAxisAlignItems field if non-nil, zero value otherwise.

### GetCounterAxisAlignItemsOk

`func (o *HasFramePropertiesTrait) GetCounterAxisAlignItemsOk() (*string, bool)`

GetCounterAxisAlignItemsOk returns a tuple with the CounterAxisAlignItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisAlignItems

`func (o *HasFramePropertiesTrait) SetCounterAxisAlignItems(v string)`

SetCounterAxisAlignItems sets CounterAxisAlignItems field to given value.

### HasCounterAxisAlignItems

`func (o *HasFramePropertiesTrait) HasCounterAxisAlignItems() bool`

HasCounterAxisAlignItems returns a boolean if a field has been set.

### GetPaddingLeft

`func (o *HasFramePropertiesTrait) GetPaddingLeft() float32`

GetPaddingLeft returns the PaddingLeft field if non-nil, zero value otherwise.

### GetPaddingLeftOk

`func (o *HasFramePropertiesTrait) GetPaddingLeftOk() (*float32, bool)`

GetPaddingLeftOk returns a tuple with the PaddingLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingLeft

`func (o *HasFramePropertiesTrait) SetPaddingLeft(v float32)`

SetPaddingLeft sets PaddingLeft field to given value.

### HasPaddingLeft

`func (o *HasFramePropertiesTrait) HasPaddingLeft() bool`

HasPaddingLeft returns a boolean if a field has been set.

### GetPaddingRight

`func (o *HasFramePropertiesTrait) GetPaddingRight() float32`

GetPaddingRight returns the PaddingRight field if non-nil, zero value otherwise.

### GetPaddingRightOk

`func (o *HasFramePropertiesTrait) GetPaddingRightOk() (*float32, bool)`

GetPaddingRightOk returns a tuple with the PaddingRight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingRight

`func (o *HasFramePropertiesTrait) SetPaddingRight(v float32)`

SetPaddingRight sets PaddingRight field to given value.

### HasPaddingRight

`func (o *HasFramePropertiesTrait) HasPaddingRight() bool`

HasPaddingRight returns a boolean if a field has been set.

### GetPaddingTop

`func (o *HasFramePropertiesTrait) GetPaddingTop() float32`

GetPaddingTop returns the PaddingTop field if non-nil, zero value otherwise.

### GetPaddingTopOk

`func (o *HasFramePropertiesTrait) GetPaddingTopOk() (*float32, bool)`

GetPaddingTopOk returns a tuple with the PaddingTop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingTop

`func (o *HasFramePropertiesTrait) SetPaddingTop(v float32)`

SetPaddingTop sets PaddingTop field to given value.

### HasPaddingTop

`func (o *HasFramePropertiesTrait) HasPaddingTop() bool`

HasPaddingTop returns a boolean if a field has been set.

### GetPaddingBottom

`func (o *HasFramePropertiesTrait) GetPaddingBottom() float32`

GetPaddingBottom returns the PaddingBottom field if non-nil, zero value otherwise.

### GetPaddingBottomOk

`func (o *HasFramePropertiesTrait) GetPaddingBottomOk() (*float32, bool)`

GetPaddingBottomOk returns a tuple with the PaddingBottom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaddingBottom

`func (o *HasFramePropertiesTrait) SetPaddingBottom(v float32)`

SetPaddingBottom sets PaddingBottom field to given value.

### HasPaddingBottom

`func (o *HasFramePropertiesTrait) HasPaddingBottom() bool`

HasPaddingBottom returns a boolean if a field has been set.

### GetItemSpacing

`func (o *HasFramePropertiesTrait) GetItemSpacing() float32`

GetItemSpacing returns the ItemSpacing field if non-nil, zero value otherwise.

### GetItemSpacingOk

`func (o *HasFramePropertiesTrait) GetItemSpacingOk() (*float32, bool)`

GetItemSpacingOk returns a tuple with the ItemSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSpacing

`func (o *HasFramePropertiesTrait) SetItemSpacing(v float32)`

SetItemSpacing sets ItemSpacing field to given value.

### HasItemSpacing

`func (o *HasFramePropertiesTrait) HasItemSpacing() bool`

HasItemSpacing returns a boolean if a field has been set.

### GetItemReverseZIndex

`func (o *HasFramePropertiesTrait) GetItemReverseZIndex() bool`

GetItemReverseZIndex returns the ItemReverseZIndex field if non-nil, zero value otherwise.

### GetItemReverseZIndexOk

`func (o *HasFramePropertiesTrait) GetItemReverseZIndexOk() (*bool, bool)`

GetItemReverseZIndexOk returns a tuple with the ItemReverseZIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemReverseZIndex

`func (o *HasFramePropertiesTrait) SetItemReverseZIndex(v bool)`

SetItemReverseZIndex sets ItemReverseZIndex field to given value.

### HasItemReverseZIndex

`func (o *HasFramePropertiesTrait) HasItemReverseZIndex() bool`

HasItemReverseZIndex returns a boolean if a field has been set.

### GetStrokesIncludedInLayout

`func (o *HasFramePropertiesTrait) GetStrokesIncludedInLayout() bool`

GetStrokesIncludedInLayout returns the StrokesIncludedInLayout field if non-nil, zero value otherwise.

### GetStrokesIncludedInLayoutOk

`func (o *HasFramePropertiesTrait) GetStrokesIncludedInLayoutOk() (*bool, bool)`

GetStrokesIncludedInLayoutOk returns a tuple with the StrokesIncludedInLayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrokesIncludedInLayout

`func (o *HasFramePropertiesTrait) SetStrokesIncludedInLayout(v bool)`

SetStrokesIncludedInLayout sets StrokesIncludedInLayout field to given value.

### HasStrokesIncludedInLayout

`func (o *HasFramePropertiesTrait) HasStrokesIncludedInLayout() bool`

HasStrokesIncludedInLayout returns a boolean if a field has been set.

### GetLayoutWrap

`func (o *HasFramePropertiesTrait) GetLayoutWrap() string`

GetLayoutWrap returns the LayoutWrap field if non-nil, zero value otherwise.

### GetLayoutWrapOk

`func (o *HasFramePropertiesTrait) GetLayoutWrapOk() (*string, bool)`

GetLayoutWrapOk returns a tuple with the LayoutWrap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutWrap

`func (o *HasFramePropertiesTrait) SetLayoutWrap(v string)`

SetLayoutWrap sets LayoutWrap field to given value.

### HasLayoutWrap

`func (o *HasFramePropertiesTrait) HasLayoutWrap() bool`

HasLayoutWrap returns a boolean if a field has been set.

### GetCounterAxisSpacing

`func (o *HasFramePropertiesTrait) GetCounterAxisSpacing() float32`

GetCounterAxisSpacing returns the CounterAxisSpacing field if non-nil, zero value otherwise.

### GetCounterAxisSpacingOk

`func (o *HasFramePropertiesTrait) GetCounterAxisSpacingOk() (*float32, bool)`

GetCounterAxisSpacingOk returns a tuple with the CounterAxisSpacing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisSpacing

`func (o *HasFramePropertiesTrait) SetCounterAxisSpacing(v float32)`

SetCounterAxisSpacing sets CounterAxisSpacing field to given value.

### HasCounterAxisSpacing

`func (o *HasFramePropertiesTrait) HasCounterAxisSpacing() bool`

HasCounterAxisSpacing returns a boolean if a field has been set.

### GetCounterAxisAlignContent

`func (o *HasFramePropertiesTrait) GetCounterAxisAlignContent() string`

GetCounterAxisAlignContent returns the CounterAxisAlignContent field if non-nil, zero value otherwise.

### GetCounterAxisAlignContentOk

`func (o *HasFramePropertiesTrait) GetCounterAxisAlignContentOk() (*string, bool)`

GetCounterAxisAlignContentOk returns a tuple with the CounterAxisAlignContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterAxisAlignContent

`func (o *HasFramePropertiesTrait) SetCounterAxisAlignContent(v string)`

SetCounterAxisAlignContent sets CounterAxisAlignContent field to given value.

### HasCounterAxisAlignContent

`func (o *HasFramePropertiesTrait) HasCounterAxisAlignContent() bool`

HasCounterAxisAlignContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


