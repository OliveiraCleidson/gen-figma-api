# Hyperlink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | The type of hyperlink. Can be either &#x60;URL&#x60; or &#x60;NODE&#x60;. | 
**Url** | Pointer to **string** | The URL that the hyperlink points to, if &#x60;type&#x60; is &#x60;URL&#x60;. | [optional] 
**NodeID** | Pointer to **string** | The ID of the node that the hyperlink points to, if &#x60;type&#x60; is &#x60;NODE&#x60;. | [optional] 

## Methods

### NewHyperlink

`func NewHyperlink(type_ string, ) *Hyperlink`

NewHyperlink instantiates a new Hyperlink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHyperlinkWithDefaults

`func NewHyperlinkWithDefaults() *Hyperlink`

NewHyperlinkWithDefaults instantiates a new Hyperlink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Hyperlink) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Hyperlink) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Hyperlink) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *Hyperlink) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Hyperlink) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Hyperlink) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Hyperlink) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetNodeID

`func (o *Hyperlink) GetNodeID() string`

GetNodeID returns the NodeID field if non-nil, zero value otherwise.

### GetNodeIDOk

`func (o *Hyperlink) GetNodeIDOk() (*string, bool)`

GetNodeIDOk returns a tuple with the NodeID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeID

`func (o *Hyperlink) SetNodeID(v string)`

SetNodeID sets NodeID field to given value.

### HasNodeID

`func (o *Hyperlink) HasNodeID() bool`

HasNodeID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


