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

// checks if the LibraryAnalyticsUsagesByFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LibraryAnalyticsUsagesByFile{}

// LibraryAnalyticsUsagesByFile Library analytics usage data broken down by file.
type LibraryAnalyticsUsagesByFile struct {
	// The name of the file using the library.
	FileName string `json:"file_name"`
	// The name of the team the file belongs to.
	TeamName string `json:"team_name"`
	// The name of the workspace that the file belongs to.
	WorkspaceName *string `json:"workspace_name,omitempty"`
	// The number of component instances from the library used within the file.
	NumInstances float32 `json:"num_instances"`
}

type _LibraryAnalyticsUsagesByFile LibraryAnalyticsUsagesByFile

// NewLibraryAnalyticsUsagesByFile instantiates a new LibraryAnalyticsUsagesByFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLibraryAnalyticsUsagesByFile(fileName string, teamName string, numInstances float32) *LibraryAnalyticsUsagesByFile {
	this := LibraryAnalyticsUsagesByFile{}
	this.FileName = fileName
	this.TeamName = teamName
	this.NumInstances = numInstances
	return &this
}

// NewLibraryAnalyticsUsagesByFileWithDefaults instantiates a new LibraryAnalyticsUsagesByFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLibraryAnalyticsUsagesByFileWithDefaults() *LibraryAnalyticsUsagesByFile {
	this := LibraryAnalyticsUsagesByFile{}
	return &this
}

// GetFileName returns the FileName field value
func (o *LibraryAnalyticsUsagesByFile) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *LibraryAnalyticsUsagesByFile) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *LibraryAnalyticsUsagesByFile) SetFileName(v string) {
	o.FileName = v
}

// GetTeamName returns the TeamName field value
func (o *LibraryAnalyticsUsagesByFile) GetTeamName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamName
}

// GetTeamNameOk returns a tuple with the TeamName field value
// and a boolean to check if the value has been set.
func (o *LibraryAnalyticsUsagesByFile) GetTeamNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamName, true
}

// SetTeamName sets field value
func (o *LibraryAnalyticsUsagesByFile) SetTeamName(v string) {
	o.TeamName = v
}

// GetWorkspaceName returns the WorkspaceName field value if set, zero value otherwise.
func (o *LibraryAnalyticsUsagesByFile) GetWorkspaceName() string {
	if o == nil || IsNil(o.WorkspaceName) {
		var ret string
		return ret
	}
	return *o.WorkspaceName
}

// GetWorkspaceNameOk returns a tuple with the WorkspaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LibraryAnalyticsUsagesByFile) GetWorkspaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.WorkspaceName) {
		return nil, false
	}
	return o.WorkspaceName, true
}

// HasWorkspaceName returns a boolean if a field has been set.
func (o *LibraryAnalyticsUsagesByFile) HasWorkspaceName() bool {
	if o != nil && !IsNil(o.WorkspaceName) {
		return true
	}

	return false
}

// SetWorkspaceName gets a reference to the given string and assigns it to the WorkspaceName field.
func (o *LibraryAnalyticsUsagesByFile) SetWorkspaceName(v string) {
	o.WorkspaceName = &v
}

// GetNumInstances returns the NumInstances field value
func (o *LibraryAnalyticsUsagesByFile) GetNumInstances() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NumInstances
}

// GetNumInstancesOk returns a tuple with the NumInstances field value
// and a boolean to check if the value has been set.
func (o *LibraryAnalyticsUsagesByFile) GetNumInstancesOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumInstances, true
}

// SetNumInstances sets field value
func (o *LibraryAnalyticsUsagesByFile) SetNumInstances(v float32) {
	o.NumInstances = v
}

func (o LibraryAnalyticsUsagesByFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LibraryAnalyticsUsagesByFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file_name"] = o.FileName
	toSerialize["team_name"] = o.TeamName
	if !IsNil(o.WorkspaceName) {
		toSerialize["workspace_name"] = o.WorkspaceName
	}
	toSerialize["num_instances"] = o.NumInstances
	return toSerialize, nil
}

func (o *LibraryAnalyticsUsagesByFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file_name",
		"team_name",
		"num_instances",
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

	varLibraryAnalyticsUsagesByFile := _LibraryAnalyticsUsagesByFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varLibraryAnalyticsUsagesByFile)

	if err != nil {
		return err
	}

	*o = LibraryAnalyticsUsagesByFile(varLibraryAnalyticsUsagesByFile)

	return err
}

type NullableLibraryAnalyticsUsagesByFile struct {
	value *LibraryAnalyticsUsagesByFile
	isSet bool
}

func (v NullableLibraryAnalyticsUsagesByFile) Get() *LibraryAnalyticsUsagesByFile {
	return v.value
}

func (v *NullableLibraryAnalyticsUsagesByFile) Set(val *LibraryAnalyticsUsagesByFile) {
	v.value = val
	v.isSet = true
}

func (v NullableLibraryAnalyticsUsagesByFile) IsSet() bool {
	return v.isSet
}

func (v *NullableLibraryAnalyticsUsagesByFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLibraryAnalyticsUsagesByFile(val *LibraryAnalyticsUsagesByFile) *NullableLibraryAnalyticsUsagesByFile {
	return &NullableLibraryAnalyticsUsagesByFile{value: val, isSet: true}
}

func (v NullableLibraryAnalyticsUsagesByFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLibraryAnalyticsUsagesByFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

