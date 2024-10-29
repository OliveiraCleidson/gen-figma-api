# LayoutConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vertical** | **string** | Vertical constraint (relative to containing frame) as an enum:  - &#x60;TOP&#x60;: Node is laid out relative to top of the containing frame - &#x60;BOTTOM&#x60;: Node is laid out relative to bottom of the containing frame - &#x60;CENTER&#x60;: Node is vertically centered relative to containing frame - &#x60;TOP_BOTTOM&#x60;: Both top and bottom of node are constrained relative to containing frame (node stretches with frame) - &#x60;SCALE&#x60;: Node scales vertically with containing frame | 
**Horizontal** | **string** | Horizontal constraint (relative to containing frame) as an enum:  - &#x60;LEFT&#x60;: Node is laid out relative to left of the containing frame - &#x60;RIGHT&#x60;: Node is laid out relative to right of the containing frame - &#x60;CENTER&#x60;: Node is horizontally centered relative to containing frame - &#x60;LEFT_RIGHT&#x60;: Both left and right of node are constrained relative to containing frame (node stretches with frame) - &#x60;SCALE&#x60;: Node scales horizontally with containing frame | 

## Methods

### NewLayoutConstraint

`func NewLayoutConstraint(vertical string, horizontal string, ) *LayoutConstraint`

NewLayoutConstraint instantiates a new LayoutConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLayoutConstraintWithDefaults

`func NewLayoutConstraintWithDefaults() *LayoutConstraint`

NewLayoutConstraintWithDefaults instantiates a new LayoutConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVertical

`func (o *LayoutConstraint) GetVertical() string`

GetVertical returns the Vertical field if non-nil, zero value otherwise.

### GetVerticalOk

`func (o *LayoutConstraint) GetVerticalOk() (*string, bool)`

GetVerticalOk returns a tuple with the Vertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertical

`func (o *LayoutConstraint) SetVertical(v string)`

SetVertical sets Vertical field to given value.


### GetHorizontal

`func (o *LayoutConstraint) GetHorizontal() string`

GetHorizontal returns the Horizontal field if non-nil, zero value otherwise.

### GetHorizontalOk

`func (o *LayoutConstraint) GetHorizontalOk() (*string, bool)`

GetHorizontalOk returns a tuple with the Horizontal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorizontal

`func (o *LayoutConstraint) SetHorizontal(v string)`

SetHorizontal sets Horizontal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


