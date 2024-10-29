# LayoutGrid

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | **string** | Orientation of the grid as a string enum  - &#x60;COLUMNS&#x60;: Vertical grid - &#x60;ROWS&#x60;: Horizontal grid - &#x60;GRID&#x60;: Square grid | 
**SectionSize** | **float32** | Width of column grid or height of row grid or square grid spacing. | 
**Visible** | **bool** | Is the grid currently visible? | 
**Color** | [**RGBA**](RGBA.md) | Color of the grid | 
**Alignment** | **string** | Positioning of grid as a string enum  - &#x60;MIN&#x60;: Grid starts at the left or top of the frame - &#x60;MAX&#x60;: Grid starts at the right or bottom of the frame - &#x60;STRETCH&#x60;: Grid is stretched to fit the frame - &#x60;CENTER&#x60;: Grid is center aligned | 
**GutterSize** | **float32** | Spacing in between columns and rows | 
**Offset** | **float32** | Spacing before the first column or row | 
**Count** | **float32** | Number of columns or rows | 
**BoundVariables** | Pointer to [**LayoutGridBoundVariables**](LayoutGridBoundVariables.md) |  | [optional] 

## Methods

### NewLayoutGrid

`func NewLayoutGrid(pattern string, sectionSize float32, visible bool, color RGBA, alignment string, gutterSize float32, offset float32, count float32, ) *LayoutGrid`

NewLayoutGrid instantiates a new LayoutGrid object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutGridWithDefaults

`func NewLayoutGridWithDefaults() *LayoutGrid`

NewLayoutGridWithDefaults instantiates a new LayoutGrid object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *LayoutGrid) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *LayoutGrid) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *LayoutGrid) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetSectionSize

`func (o *LayoutGrid) GetSectionSize() float32`

GetSectionSize returns the SectionSize field if non-nil, zero value otherwise.

### GetSectionSizeOk

`func (o *LayoutGrid) GetSectionSizeOk() (*float32, bool)`

GetSectionSizeOk returns a tuple with the SectionSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionSize

`func (o *LayoutGrid) SetSectionSize(v float32)`

SetSectionSize sets SectionSize field to given value.


### GetVisible

`func (o *LayoutGrid) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *LayoutGrid) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *LayoutGrid) SetVisible(v bool)`

SetVisible sets Visible field to given value.


### GetColor

`func (o *LayoutGrid) GetColor() RGBA`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *LayoutGrid) GetColorOk() (*RGBA, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *LayoutGrid) SetColor(v RGBA)`

SetColor sets Color field to given value.


### GetAlignment

`func (o *LayoutGrid) GetAlignment() string`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *LayoutGrid) GetAlignmentOk() (*string, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *LayoutGrid) SetAlignment(v string)`

SetAlignment sets Alignment field to given value.


### GetGutterSize

`func (o *LayoutGrid) GetGutterSize() float32`

GetGutterSize returns the GutterSize field if non-nil, zero value otherwise.

### GetGutterSizeOk

`func (o *LayoutGrid) GetGutterSizeOk() (*float32, bool)`

GetGutterSizeOk returns a tuple with the GutterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGutterSize

`func (o *LayoutGrid) SetGutterSize(v float32)`

SetGutterSize sets GutterSize field to given value.


### GetOffset

`func (o *LayoutGrid) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *LayoutGrid) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *LayoutGrid) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetCount

`func (o *LayoutGrid) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *LayoutGrid) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *LayoutGrid) SetCount(v float32)`

SetCount sets Count field to given value.


### GetBoundVariables

`func (o *LayoutGrid) GetBoundVariables() LayoutGridBoundVariables`

GetBoundVariables returns the BoundVariables field if non-nil, zero value otherwise.

### GetBoundVariablesOk

`func (o *LayoutGrid) GetBoundVariablesOk() (*LayoutGridBoundVariables, bool)`

GetBoundVariablesOk returns a tuple with the BoundVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundVariables

`func (o *LayoutGrid) SetBoundVariables(v LayoutGridBoundVariables)`

SetBoundVariables sets BoundVariables field to given value.

### HasBoundVariables

`func (o *LayoutGrid) HasBoundVariables() bool`

HasBoundVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


