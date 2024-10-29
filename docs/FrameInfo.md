# FrameInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeId** | Pointer to **string** | The ID of the frame node within the file. | [optional] 
**Name** | Pointer to **string** | The name of the frame node. | [optional] 
**BackgroundColor** | Pointer to **string** | The background color of the frame node. | [optional] 
**PageId** | **string** | The ID of the page containing the frame node. | 
**PageName** | **string** | The name of the page containing the frame node. | 

## Methods

### NewFrameInfo

`func NewFrameInfo(pageId string, pageName string, ) *FrameInfo`

NewFrameInfo instantiates a new FrameInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFrameInfoWithDefaults

`func NewFrameInfoWithDefaults() *FrameInfo`

NewFrameInfoWithDefaults instantiates a new FrameInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeId

`func (o *FrameInfo) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *FrameInfo) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *FrameInfo) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *FrameInfo) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetName

`func (o *FrameInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FrameInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FrameInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FrameInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *FrameInfo) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *FrameInfo) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *FrameInfo) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *FrameInfo) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetPageId

`func (o *FrameInfo) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *FrameInfo) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *FrameInfo) SetPageId(v string)`

SetPageId sets PageId field to given value.


### GetPageName

`func (o *FrameInfo) GetPageName() string`

GetPageName returns the PageName field if non-nil, zero value otherwise.

### GetPageNameOk

`func (o *FrameInfo) GetPageNameOk() (*string, bool)`

GetPageNameOk returns a tuple with the PageName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageName

`func (o *FrameInfo) SetPageName(v string)`

SetPageName sets PageName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


