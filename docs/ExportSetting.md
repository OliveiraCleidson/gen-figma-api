# ExportSetting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suffix** | **string** |  | 
**Format** | **string** |  | 
**Constraint** | [**Constraint**](Constraint.md) |  | 

## Methods

### NewExportSetting

`func NewExportSetting(suffix string, format string, constraint Constraint, ) *ExportSetting`

NewExportSetting instantiates a new ExportSetting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportSettingWithDefaults

`func NewExportSettingWithDefaults() *ExportSetting`

NewExportSettingWithDefaults instantiates a new ExportSetting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuffix

`func (o *ExportSetting) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *ExportSetting) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *ExportSetting) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.


### GetFormat

`func (o *ExportSetting) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ExportSetting) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ExportSetting) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetConstraint

`func (o *ExportSetting) GetConstraint() Constraint`

GetConstraint returns the Constraint field if non-nil, zero value otherwise.

### GetConstraintOk

`func (o *ExportSetting) GetConstraintOk() (*Constraint, bool)`

GetConstraintOk returns a tuple with the Constraint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraint

`func (o *ExportSetting) SetConstraint(v Constraint)`

SetConstraint sets Constraint field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


