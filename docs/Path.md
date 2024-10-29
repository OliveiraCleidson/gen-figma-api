# Path

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | **string** | A series of path commands that encodes how to draw the path. | 
**WindingRule** | **string** | The winding rule for the path (same as in SVGs). This determines whether a given point in space is inside or outside the path. | 
**OverrideID** | Pointer to **float32** | If there is a per-region fill, this refers to an ID in the &#x60;fillOverrideTable&#x60;. | [optional] 

## Methods

### NewPath

`func NewPath(path string, windingRule string, ) *Path`

NewPath instantiates a new Path object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPathWithDefaults

`func NewPathWithDefaults() *Path`

NewPathWithDefaults instantiates a new Path object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *Path) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Path) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Path) SetPath(v string)`

SetPath sets Path field to given value.


### GetWindingRule

`func (o *Path) GetWindingRule() string`

GetWindingRule returns the WindingRule field if non-nil, zero value otherwise.

### GetWindingRuleOk

`func (o *Path) GetWindingRuleOk() (*string, bool)`

GetWindingRuleOk returns a tuple with the WindingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindingRule

`func (o *Path) SetWindingRule(v string)`

SetWindingRule sets WindingRule field to given value.


### GetOverrideID

`func (o *Path) GetOverrideID() float32`

GetOverrideID returns the OverrideID field if non-nil, zero value otherwise.

### GetOverrideIDOk

`func (o *Path) GetOverrideIDOk() (*float32, bool)`

GetOverrideIDOk returns a tuple with the OverrideID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideID

`func (o *Path) SetOverrideID(v float32)`

SetOverrideID sets OverrideID field to given value.

### HasOverrideID

`func (o *Path) HasOverrideID() bool`

HasOverrideID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


