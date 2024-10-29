/*
Figma API

This is the OpenAPI specification for the [Figma REST API](https://www.figma.com/developers/api).  Note: we are releasing the OpenAPI specification as a beta given the large surface area and complexity of the REST API. If you notice any inaccuracies with the specification, please [file an issue](https://github.com/figma/rest-api-spec/issues).

API version: 0.20.0
Contact: support@figma.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the HasFramePropertiesTrait type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HasFramePropertiesTrait{}

// HasFramePropertiesTrait struct for HasFramePropertiesTrait
type HasFramePropertiesTrait struct {
	// Whether or not this node clip content outside of its bounds
	ClipsContent bool `json:"clipsContent"`
	// Background of the node. This is deprecated, as backgrounds for frames are now in the `fills` field.
	Background []Paint `json:"background,omitempty"`
	// Background color of the node. This is deprecated, as frames now support more than a solid color as a background. Please use the `fills` field instead.
	// Deprecated
	BackgroundColor *RGBA `json:"backgroundColor,omitempty"`
	// An array of layout grids attached to this node (see layout grids section for more details). GROUP nodes do not have this attribute
	LayoutGrids []LayoutGrid `json:"layoutGrids,omitempty"`
	// Whether a node has primary axis scrolling, horizontal or vertical.
	OverflowDirection *string `json:"overflowDirection,omitempty"`
	// Whether this layer uses auto-layout to position its children.
	LayoutMode *string `json:"layoutMode,omitempty"`
	// Whether the primary axis has a fixed length (determined by the user) or an automatic length (determined by the layout engine). This property is only applicable for auto-layout frames.
	PrimaryAxisSizingMode *string `json:"primaryAxisSizingMode,omitempty"`
	// Whether the counter axis has a fixed length (determined by the user) or an automatic length (determined by the layout engine). This property is only applicable for auto-layout frames.
	CounterAxisSizingMode *string `json:"counterAxisSizingMode,omitempty"`
	// Determines how the auto-layout frame's children should be aligned in the primary axis direction. This property is only applicable for auto-layout frames.
	PrimaryAxisAlignItems *string `json:"primaryAxisAlignItems,omitempty"`
	// Determines how the auto-layout frame's children should be aligned in the counter axis direction. This property is only applicable for auto-layout frames.
	CounterAxisAlignItems *string `json:"counterAxisAlignItems,omitempty"`
	// The padding between the left border of the frame and its children. This property is only applicable for auto-layout frames.
	PaddingLeft *float32 `json:"paddingLeft,omitempty"`
	// The padding between the right border of the frame and its children. This property is only applicable for auto-layout frames.
	PaddingRight *float32 `json:"paddingRight,omitempty"`
	// The padding between the top border of the frame and its children. This property is only applicable for auto-layout frames.
	PaddingTop *float32 `json:"paddingTop,omitempty"`
	// The padding between the bottom border of the frame and its children. This property is only applicable for auto-layout frames.
	PaddingBottom *float32 `json:"paddingBottom,omitempty"`
	// The distance between children of the frame. Can be negative. This property is only applicable for auto-layout frames.
	ItemSpacing *float32 `json:"itemSpacing,omitempty"`
	// Determines the canvas stacking order of layers in this frame. When true, the first layer will be draw on top. This property is only applicable for auto-layout frames.
	ItemReverseZIndex *bool `json:"itemReverseZIndex,omitempty"`
	// Determines whether strokes are included in layout calculations. When true, auto-layout frames behave like css \"box-sizing: border-box\". This property is only applicable for auto-layout frames.
	StrokesIncludedInLayout *bool `json:"strokesIncludedInLayout,omitempty"`
	// Whether this auto-layout frame has wrapping enabled.
	LayoutWrap *string `json:"layoutWrap,omitempty"`
	// The distance between wrapped tracks of an auto-layout frame. This property is only applicable for auto-layout frames with `layoutWrap: \"WRAP\"`
	CounterAxisSpacing *float32 `json:"counterAxisSpacing,omitempty"`
	// Determines how the auto-layout frame’s wrapped tracks should be aligned in the counter axis direction. This property is only applicable for auto-layout frames with `layoutWrap: \"WRAP\"`.
	CounterAxisAlignContent *string `json:"counterAxisAlignContent,omitempty"`
}

type _HasFramePropertiesTrait HasFramePropertiesTrait

// NewHasFramePropertiesTrait instantiates a new HasFramePropertiesTrait object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHasFramePropertiesTrait(clipsContent bool) *HasFramePropertiesTrait {
	this := HasFramePropertiesTrait{}
	this.ClipsContent = clipsContent
	var overflowDirection string = "NONE"
	this.OverflowDirection = &overflowDirection
	var layoutMode string = "NONE"
	this.LayoutMode = &layoutMode
	var primaryAxisSizingMode string = "AUTO"
	this.PrimaryAxisSizingMode = &primaryAxisSizingMode
	var counterAxisSizingMode string = "AUTO"
	this.CounterAxisSizingMode = &counterAxisSizingMode
	var primaryAxisAlignItems string = "MIN"
	this.PrimaryAxisAlignItems = &primaryAxisAlignItems
	var counterAxisAlignItems string = "MIN"
	this.CounterAxisAlignItems = &counterAxisAlignItems
	var paddingLeft float32 = 0
	this.PaddingLeft = &paddingLeft
	var paddingRight float32 = 0
	this.PaddingRight = &paddingRight
	var paddingTop float32 = 0
	this.PaddingTop = &paddingTop
	var paddingBottom float32 = 0
	this.PaddingBottom = &paddingBottom
	var itemSpacing float32 = 0
	this.ItemSpacing = &itemSpacing
	var itemReverseZIndex bool = false
	this.ItemReverseZIndex = &itemReverseZIndex
	var strokesIncludedInLayout bool = false
	this.StrokesIncludedInLayout = &strokesIncludedInLayout
	var counterAxisAlignContent string = "AUTO"
	this.CounterAxisAlignContent = &counterAxisAlignContent
	return &this
}

// NewHasFramePropertiesTraitWithDefaults instantiates a new HasFramePropertiesTrait object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHasFramePropertiesTraitWithDefaults() *HasFramePropertiesTrait {
	this := HasFramePropertiesTrait{}
	var overflowDirection string = "NONE"
	this.OverflowDirection = &overflowDirection
	var layoutMode string = "NONE"
	this.LayoutMode = &layoutMode
	var primaryAxisSizingMode string = "AUTO"
	this.PrimaryAxisSizingMode = &primaryAxisSizingMode
	var counterAxisSizingMode string = "AUTO"
	this.CounterAxisSizingMode = &counterAxisSizingMode
	var primaryAxisAlignItems string = "MIN"
	this.PrimaryAxisAlignItems = &primaryAxisAlignItems
	var counterAxisAlignItems string = "MIN"
	this.CounterAxisAlignItems = &counterAxisAlignItems
	var paddingLeft float32 = 0
	this.PaddingLeft = &paddingLeft
	var paddingRight float32 = 0
	this.PaddingRight = &paddingRight
	var paddingTop float32 = 0
	this.PaddingTop = &paddingTop
	var paddingBottom float32 = 0
	this.PaddingBottom = &paddingBottom
	var itemSpacing float32 = 0
	this.ItemSpacing = &itemSpacing
	var itemReverseZIndex bool = false
	this.ItemReverseZIndex = &itemReverseZIndex
	var strokesIncludedInLayout bool = false
	this.StrokesIncludedInLayout = &strokesIncludedInLayout
	var counterAxisAlignContent string = "AUTO"
	this.CounterAxisAlignContent = &counterAxisAlignContent
	return &this
}

// GetClipsContent returns the ClipsContent field value
func (o *HasFramePropertiesTrait) GetClipsContent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ClipsContent
}

// GetClipsContentOk returns a tuple with the ClipsContent field value
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetClipsContentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClipsContent, true
}

// SetClipsContent sets field value
func (o *HasFramePropertiesTrait) SetClipsContent(v bool) {
	o.ClipsContent = v
}

// GetBackground returns the Background field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetBackground() []Paint {
	if o == nil || IsNil(o.Background) {
		var ret []Paint
		return ret
	}
	return o.Background
}

// GetBackgroundOk returns a tuple with the Background field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetBackgroundOk() ([]Paint, bool) {
	if o == nil || IsNil(o.Background) {
		return nil, false
	}
	return o.Background, true
}

// HasBackground returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasBackground() bool {
	if o != nil && !IsNil(o.Background) {
		return true
	}

	return false
}

// SetBackground gets a reference to the given []Paint and assigns it to the Background field.
func (o *HasFramePropertiesTrait) SetBackground(v []Paint) {
	o.Background = v
}

// GetBackgroundColor returns the BackgroundColor field value if set, zero value otherwise.
// Deprecated
func (o *HasFramePropertiesTrait) GetBackgroundColor() RGBA {
	if o == nil || IsNil(o.BackgroundColor) {
		var ret RGBA
		return ret
	}
	return *o.BackgroundColor
}

// GetBackgroundColorOk returns a tuple with the BackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *HasFramePropertiesTrait) GetBackgroundColorOk() (*RGBA, bool) {
	if o == nil || IsNil(o.BackgroundColor) {
		return nil, false
	}
	return o.BackgroundColor, true
}

// HasBackgroundColor returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasBackgroundColor() bool {
	if o != nil && !IsNil(o.BackgroundColor) {
		return true
	}

	return false
}

// SetBackgroundColor gets a reference to the given RGBA and assigns it to the BackgroundColor field.
// Deprecated
func (o *HasFramePropertiesTrait) SetBackgroundColor(v RGBA) {
	o.BackgroundColor = &v
}

// GetLayoutGrids returns the LayoutGrids field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetLayoutGrids() []LayoutGrid {
	if o == nil || IsNil(o.LayoutGrids) {
		var ret []LayoutGrid
		return ret
	}
	return o.LayoutGrids
}

// GetLayoutGridsOk returns a tuple with the LayoutGrids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetLayoutGridsOk() ([]LayoutGrid, bool) {
	if o == nil || IsNil(o.LayoutGrids) {
		return nil, false
	}
	return o.LayoutGrids, true
}

// HasLayoutGrids returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasLayoutGrids() bool {
	if o != nil && !IsNil(o.LayoutGrids) {
		return true
	}

	return false
}

// SetLayoutGrids gets a reference to the given []LayoutGrid and assigns it to the LayoutGrids field.
func (o *HasFramePropertiesTrait) SetLayoutGrids(v []LayoutGrid) {
	o.LayoutGrids = v
}

// GetOverflowDirection returns the OverflowDirection field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetOverflowDirection() string {
	if o == nil || IsNil(o.OverflowDirection) {
		var ret string
		return ret
	}
	return *o.OverflowDirection
}

// GetOverflowDirectionOk returns a tuple with the OverflowDirection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetOverflowDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.OverflowDirection) {
		return nil, false
	}
	return o.OverflowDirection, true
}

// HasOverflowDirection returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasOverflowDirection() bool {
	if o != nil && !IsNil(o.OverflowDirection) {
		return true
	}

	return false
}

// SetOverflowDirection gets a reference to the given string and assigns it to the OverflowDirection field.
func (o *HasFramePropertiesTrait) SetOverflowDirection(v string) {
	o.OverflowDirection = &v
}

// GetLayoutMode returns the LayoutMode field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetLayoutMode() string {
	if o == nil || IsNil(o.LayoutMode) {
		var ret string
		return ret
	}
	return *o.LayoutMode
}

// GetLayoutModeOk returns a tuple with the LayoutMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetLayoutModeOk() (*string, bool) {
	if o == nil || IsNil(o.LayoutMode) {
		return nil, false
	}
	return o.LayoutMode, true
}

// HasLayoutMode returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasLayoutMode() bool {
	if o != nil && !IsNil(o.LayoutMode) {
		return true
	}

	return false
}

// SetLayoutMode gets a reference to the given string and assigns it to the LayoutMode field.
func (o *HasFramePropertiesTrait) SetLayoutMode(v string) {
	o.LayoutMode = &v
}

// GetPrimaryAxisSizingMode returns the PrimaryAxisSizingMode field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetPrimaryAxisSizingMode() string {
	if o == nil || IsNil(o.PrimaryAxisSizingMode) {
		var ret string
		return ret
	}
	return *o.PrimaryAxisSizingMode
}

// GetPrimaryAxisSizingModeOk returns a tuple with the PrimaryAxisSizingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetPrimaryAxisSizingModeOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryAxisSizingMode) {
		return nil, false
	}
	return o.PrimaryAxisSizingMode, true
}

// HasPrimaryAxisSizingMode returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasPrimaryAxisSizingMode() bool {
	if o != nil && !IsNil(o.PrimaryAxisSizingMode) {
		return true
	}

	return false
}

// SetPrimaryAxisSizingMode gets a reference to the given string and assigns it to the PrimaryAxisSizingMode field.
func (o *HasFramePropertiesTrait) SetPrimaryAxisSizingMode(v string) {
	o.PrimaryAxisSizingMode = &v
}

// GetCounterAxisSizingMode returns the CounterAxisSizingMode field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetCounterAxisSizingMode() string {
	if o == nil || IsNil(o.CounterAxisSizingMode) {
		var ret string
		return ret
	}
	return *o.CounterAxisSizingMode
}

// GetCounterAxisSizingModeOk returns a tuple with the CounterAxisSizingMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetCounterAxisSizingModeOk() (*string, bool) {
	if o == nil || IsNil(o.CounterAxisSizingMode) {
		return nil, false
	}
	return o.CounterAxisSizingMode, true
}

// HasCounterAxisSizingMode returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasCounterAxisSizingMode() bool {
	if o != nil && !IsNil(o.CounterAxisSizingMode) {
		return true
	}

	return false
}

// SetCounterAxisSizingMode gets a reference to the given string and assigns it to the CounterAxisSizingMode field.
func (o *HasFramePropertiesTrait) SetCounterAxisSizingMode(v string) {
	o.CounterAxisSizingMode = &v
}

// GetPrimaryAxisAlignItems returns the PrimaryAxisAlignItems field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetPrimaryAxisAlignItems() string {
	if o == nil || IsNil(o.PrimaryAxisAlignItems) {
		var ret string
		return ret
	}
	return *o.PrimaryAxisAlignItems
}

// GetPrimaryAxisAlignItemsOk returns a tuple with the PrimaryAxisAlignItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetPrimaryAxisAlignItemsOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryAxisAlignItems) {
		return nil, false
	}
	return o.PrimaryAxisAlignItems, true
}

// HasPrimaryAxisAlignItems returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasPrimaryAxisAlignItems() bool {
	if o != nil && !IsNil(o.PrimaryAxisAlignItems) {
		return true
	}

	return false
}

// SetPrimaryAxisAlignItems gets a reference to the given string and assigns it to the PrimaryAxisAlignItems field.
func (o *HasFramePropertiesTrait) SetPrimaryAxisAlignItems(v string) {
	o.PrimaryAxisAlignItems = &v
}

// GetCounterAxisAlignItems returns the CounterAxisAlignItems field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetCounterAxisAlignItems() string {
	if o == nil || IsNil(o.CounterAxisAlignItems) {
		var ret string
		return ret
	}
	return *o.CounterAxisAlignItems
}

// GetCounterAxisAlignItemsOk returns a tuple with the CounterAxisAlignItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetCounterAxisAlignItemsOk() (*string, bool) {
	if o == nil || IsNil(o.CounterAxisAlignItems) {
		return nil, false
	}
	return o.CounterAxisAlignItems, true
}

// HasCounterAxisAlignItems returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasCounterAxisAlignItems() bool {
	if o != nil && !IsNil(o.CounterAxisAlignItems) {
		return true
	}

	return false
}

// SetCounterAxisAlignItems gets a reference to the given string and assigns it to the CounterAxisAlignItems field.
func (o *HasFramePropertiesTrait) SetCounterAxisAlignItems(v string) {
	o.CounterAxisAlignItems = &v
}

// GetPaddingLeft returns the PaddingLeft field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetPaddingLeft() float32 {
	if o == nil || IsNil(o.PaddingLeft) {
		var ret float32
		return ret
	}
	return *o.PaddingLeft
}

// GetPaddingLeftOk returns a tuple with the PaddingLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetPaddingLeftOk() (*float32, bool) {
	if o == nil || IsNil(o.PaddingLeft) {
		return nil, false
	}
	return o.PaddingLeft, true
}

// HasPaddingLeft returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasPaddingLeft() bool {
	if o != nil && !IsNil(o.PaddingLeft) {
		return true
	}

	return false
}

// SetPaddingLeft gets a reference to the given float32 and assigns it to the PaddingLeft field.
func (o *HasFramePropertiesTrait) SetPaddingLeft(v float32) {
	o.PaddingLeft = &v
}

// GetPaddingRight returns the PaddingRight field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetPaddingRight() float32 {
	if o == nil || IsNil(o.PaddingRight) {
		var ret float32
		return ret
	}
	return *o.PaddingRight
}

// GetPaddingRightOk returns a tuple with the PaddingRight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetPaddingRightOk() (*float32, bool) {
	if o == nil || IsNil(o.PaddingRight) {
		return nil, false
	}
	return o.PaddingRight, true
}

// HasPaddingRight returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasPaddingRight() bool {
	if o != nil && !IsNil(o.PaddingRight) {
		return true
	}

	return false
}

// SetPaddingRight gets a reference to the given float32 and assigns it to the PaddingRight field.
func (o *HasFramePropertiesTrait) SetPaddingRight(v float32) {
	o.PaddingRight = &v
}

// GetPaddingTop returns the PaddingTop field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetPaddingTop() float32 {
	if o == nil || IsNil(o.PaddingTop) {
		var ret float32
		return ret
	}
	return *o.PaddingTop
}

// GetPaddingTopOk returns a tuple with the PaddingTop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetPaddingTopOk() (*float32, bool) {
	if o == nil || IsNil(o.PaddingTop) {
		return nil, false
	}
	return o.PaddingTop, true
}

// HasPaddingTop returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasPaddingTop() bool {
	if o != nil && !IsNil(o.PaddingTop) {
		return true
	}

	return false
}

// SetPaddingTop gets a reference to the given float32 and assigns it to the PaddingTop field.
func (o *HasFramePropertiesTrait) SetPaddingTop(v float32) {
	o.PaddingTop = &v
}

// GetPaddingBottom returns the PaddingBottom field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetPaddingBottom() float32 {
	if o == nil || IsNil(o.PaddingBottom) {
		var ret float32
		return ret
	}
	return *o.PaddingBottom
}

// GetPaddingBottomOk returns a tuple with the PaddingBottom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetPaddingBottomOk() (*float32, bool) {
	if o == nil || IsNil(o.PaddingBottom) {
		return nil, false
	}
	return o.PaddingBottom, true
}

// HasPaddingBottom returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasPaddingBottom() bool {
	if o != nil && !IsNil(o.PaddingBottom) {
		return true
	}

	return false
}

// SetPaddingBottom gets a reference to the given float32 and assigns it to the PaddingBottom field.
func (o *HasFramePropertiesTrait) SetPaddingBottom(v float32) {
	o.PaddingBottom = &v
}

// GetItemSpacing returns the ItemSpacing field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetItemSpacing() float32 {
	if o == nil || IsNil(o.ItemSpacing) {
		var ret float32
		return ret
	}
	return *o.ItemSpacing
}

// GetItemSpacingOk returns a tuple with the ItemSpacing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetItemSpacingOk() (*float32, bool) {
	if o == nil || IsNil(o.ItemSpacing) {
		return nil, false
	}
	return o.ItemSpacing, true
}

// HasItemSpacing returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasItemSpacing() bool {
	if o != nil && !IsNil(o.ItemSpacing) {
		return true
	}

	return false
}

// SetItemSpacing gets a reference to the given float32 and assigns it to the ItemSpacing field.
func (o *HasFramePropertiesTrait) SetItemSpacing(v float32) {
	o.ItemSpacing = &v
}

// GetItemReverseZIndex returns the ItemReverseZIndex field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetItemReverseZIndex() bool {
	if o == nil || IsNil(o.ItemReverseZIndex) {
		var ret bool
		return ret
	}
	return *o.ItemReverseZIndex
}

// GetItemReverseZIndexOk returns a tuple with the ItemReverseZIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetItemReverseZIndexOk() (*bool, bool) {
	if o == nil || IsNil(o.ItemReverseZIndex) {
		return nil, false
	}
	return o.ItemReverseZIndex, true
}

// HasItemReverseZIndex returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasItemReverseZIndex() bool {
	if o != nil && !IsNil(o.ItemReverseZIndex) {
		return true
	}

	return false
}

// SetItemReverseZIndex gets a reference to the given bool and assigns it to the ItemReverseZIndex field.
func (o *HasFramePropertiesTrait) SetItemReverseZIndex(v bool) {
	o.ItemReverseZIndex = &v
}

// GetStrokesIncludedInLayout returns the StrokesIncludedInLayout field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetStrokesIncludedInLayout() bool {
	if o == nil || IsNil(o.StrokesIncludedInLayout) {
		var ret bool
		return ret
	}
	return *o.StrokesIncludedInLayout
}

// GetStrokesIncludedInLayoutOk returns a tuple with the StrokesIncludedInLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetStrokesIncludedInLayoutOk() (*bool, bool) {
	if o == nil || IsNil(o.StrokesIncludedInLayout) {
		return nil, false
	}
	return o.StrokesIncludedInLayout, true
}

// HasStrokesIncludedInLayout returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasStrokesIncludedInLayout() bool {
	if o != nil && !IsNil(o.StrokesIncludedInLayout) {
		return true
	}

	return false
}

// SetStrokesIncludedInLayout gets a reference to the given bool and assigns it to the StrokesIncludedInLayout field.
func (o *HasFramePropertiesTrait) SetStrokesIncludedInLayout(v bool) {
	o.StrokesIncludedInLayout = &v
}

// GetLayoutWrap returns the LayoutWrap field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetLayoutWrap() string {
	if o == nil || IsNil(o.LayoutWrap) {
		var ret string
		return ret
	}
	return *o.LayoutWrap
}

// GetLayoutWrapOk returns a tuple with the LayoutWrap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetLayoutWrapOk() (*string, bool) {
	if o == nil || IsNil(o.LayoutWrap) {
		return nil, false
	}
	return o.LayoutWrap, true
}

// HasLayoutWrap returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasLayoutWrap() bool {
	if o != nil && !IsNil(o.LayoutWrap) {
		return true
	}

	return false
}

// SetLayoutWrap gets a reference to the given string and assigns it to the LayoutWrap field.
func (o *HasFramePropertiesTrait) SetLayoutWrap(v string) {
	o.LayoutWrap = &v
}

// GetCounterAxisSpacing returns the CounterAxisSpacing field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetCounterAxisSpacing() float32 {
	if o == nil || IsNil(o.CounterAxisSpacing) {
		var ret float32
		return ret
	}
	return *o.CounterAxisSpacing
}

// GetCounterAxisSpacingOk returns a tuple with the CounterAxisSpacing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetCounterAxisSpacingOk() (*float32, bool) {
	if o == nil || IsNil(o.CounterAxisSpacing) {
		return nil, false
	}
	return o.CounterAxisSpacing, true
}

// HasCounterAxisSpacing returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasCounterAxisSpacing() bool {
	if o != nil && !IsNil(o.CounterAxisSpacing) {
		return true
	}

	return false
}

// SetCounterAxisSpacing gets a reference to the given float32 and assigns it to the CounterAxisSpacing field.
func (o *HasFramePropertiesTrait) SetCounterAxisSpacing(v float32) {
	o.CounterAxisSpacing = &v
}

// GetCounterAxisAlignContent returns the CounterAxisAlignContent field value if set, zero value otherwise.
func (o *HasFramePropertiesTrait) GetCounterAxisAlignContent() string {
	if o == nil || IsNil(o.CounterAxisAlignContent) {
		var ret string
		return ret
	}
	return *o.CounterAxisAlignContent
}

// GetCounterAxisAlignContentOk returns a tuple with the CounterAxisAlignContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HasFramePropertiesTrait) GetCounterAxisAlignContentOk() (*string, bool) {
	if o == nil || IsNil(o.CounterAxisAlignContent) {
		return nil, false
	}
	return o.CounterAxisAlignContent, true
}

// HasCounterAxisAlignContent returns a boolean if a field has been set.
func (o *HasFramePropertiesTrait) HasCounterAxisAlignContent() bool {
	if o != nil && !IsNil(o.CounterAxisAlignContent) {
		return true
	}

	return false
}

// SetCounterAxisAlignContent gets a reference to the given string and assigns it to the CounterAxisAlignContent field.
func (o *HasFramePropertiesTrait) SetCounterAxisAlignContent(v string) {
	o.CounterAxisAlignContent = &v
}

func (o HasFramePropertiesTrait) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HasFramePropertiesTrait) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clipsContent"] = o.ClipsContent
	if !IsNil(o.Background) {
		toSerialize["background"] = o.Background
	}
	if !IsNil(o.BackgroundColor) {
		toSerialize["backgroundColor"] = o.BackgroundColor
	}
	if !IsNil(o.LayoutGrids) {
		toSerialize["layoutGrids"] = o.LayoutGrids
	}
	if !IsNil(o.OverflowDirection) {
		toSerialize["overflowDirection"] = o.OverflowDirection
	}
	if !IsNil(o.LayoutMode) {
		toSerialize["layoutMode"] = o.LayoutMode
	}
	if !IsNil(o.PrimaryAxisSizingMode) {
		toSerialize["primaryAxisSizingMode"] = o.PrimaryAxisSizingMode
	}
	if !IsNil(o.CounterAxisSizingMode) {
		toSerialize["counterAxisSizingMode"] = o.CounterAxisSizingMode
	}
	if !IsNil(o.PrimaryAxisAlignItems) {
		toSerialize["primaryAxisAlignItems"] = o.PrimaryAxisAlignItems
	}
	if !IsNil(o.CounterAxisAlignItems) {
		toSerialize["counterAxisAlignItems"] = o.CounterAxisAlignItems
	}
	if !IsNil(o.PaddingLeft) {
		toSerialize["paddingLeft"] = o.PaddingLeft
	}
	if !IsNil(o.PaddingRight) {
		toSerialize["paddingRight"] = o.PaddingRight
	}
	if !IsNil(o.PaddingTop) {
		toSerialize["paddingTop"] = o.PaddingTop
	}
	if !IsNil(o.PaddingBottom) {
		toSerialize["paddingBottom"] = o.PaddingBottom
	}
	if !IsNil(o.ItemSpacing) {
		toSerialize["itemSpacing"] = o.ItemSpacing
	}
	if !IsNil(o.ItemReverseZIndex) {
		toSerialize["itemReverseZIndex"] = o.ItemReverseZIndex
	}
	if !IsNil(o.StrokesIncludedInLayout) {
		toSerialize["strokesIncludedInLayout"] = o.StrokesIncludedInLayout
	}
	if !IsNil(o.LayoutWrap) {
		toSerialize["layoutWrap"] = o.LayoutWrap
	}
	if !IsNil(o.CounterAxisSpacing) {
		toSerialize["counterAxisSpacing"] = o.CounterAxisSpacing
	}
	if !IsNil(o.CounterAxisAlignContent) {
		toSerialize["counterAxisAlignContent"] = o.CounterAxisAlignContent
	}
	return toSerialize, nil
}

func (o *HasFramePropertiesTrait) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"clipsContent",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varHasFramePropertiesTrait := _HasFramePropertiesTrait{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHasFramePropertiesTrait)

	if err != nil {
		return err
	}

	*o = HasFramePropertiesTrait(varHasFramePropertiesTrait)

	return err
}

type NullableHasFramePropertiesTrait struct {
	value *HasFramePropertiesTrait
	isSet bool
}

func (v NullableHasFramePropertiesTrait) Get() *HasFramePropertiesTrait {
	return v.value
}

func (v *NullableHasFramePropertiesTrait) Set(val *HasFramePropertiesTrait) {
	v.value = val
	v.isSet = true
}

func (v NullableHasFramePropertiesTrait) IsSet() bool {
	return v.isSet
}

func (v *NullableHasFramePropertiesTrait) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHasFramePropertiesTrait(val *HasFramePropertiesTrait) *NullableHasFramePropertiesTrait {
	return &NullableHasFramePropertiesTrait{value: val, isSet: true}
}

func (v NullableHasFramePropertiesTrait) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHasFramePropertiesTrait) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

