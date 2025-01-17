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

// checks if the ActivityLogContext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActivityLogContext{}

// ActivityLogContext Contextual information about the event.
type ActivityLogContext struct {
	// The third-party application that triggered the event, if applicable.
	ClientName NullableString `json:"client_name"`
	// The IP address from of the client that sent the event request.
	IpAddress string `json:"ip_address"`
	// If Figma's Support team triggered the event. This is either true or false.
	IsFigmaSupportTeamAction bool `json:"is_figma_support_team_action"`
	// The id of the organization where the event took place.
	OrgId string `json:"org_id"`
	// The id of the team where the event took place -- if this took place in a specific team.
	TeamId NullableString `json:"team_id"`
}

type _ActivityLogContext ActivityLogContext

// NewActivityLogContext instantiates a new ActivityLogContext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityLogContext(clientName NullableString, ipAddress string, isFigmaSupportTeamAction bool, orgId string, teamId NullableString) *ActivityLogContext {
	this := ActivityLogContext{}
	this.ClientName = clientName
	this.IpAddress = ipAddress
	this.IsFigmaSupportTeamAction = isFigmaSupportTeamAction
	this.OrgId = orgId
	this.TeamId = teamId
	return &this
}

// NewActivityLogContextWithDefaults instantiates a new ActivityLogContext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityLogContextWithDefaults() *ActivityLogContext {
	this := ActivityLogContext{}
	return &this
}

// GetClientName returns the ClientName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ActivityLogContext) GetClientName() string {
	if o == nil || o.ClientName.Get() == nil {
		var ret string
		return ret
	}

	return *o.ClientName.Get()
}

// GetClientNameOk returns a tuple with the ClientName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivityLogContext) GetClientNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ClientName.Get(), o.ClientName.IsSet()
}

// SetClientName sets field value
func (o *ActivityLogContext) SetClientName(v string) {
	o.ClientName.Set(&v)
}

// GetIpAddress returns the IpAddress field value
func (o *ActivityLogContext) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *ActivityLogContext) GetIpAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *ActivityLogContext) SetIpAddress(v string) {
	o.IpAddress = v
}

// GetIsFigmaSupportTeamAction returns the IsFigmaSupportTeamAction field value
func (o *ActivityLogContext) GetIsFigmaSupportTeamAction() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFigmaSupportTeamAction
}

// GetIsFigmaSupportTeamActionOk returns a tuple with the IsFigmaSupportTeamAction field value
// and a boolean to check if the value has been set.
func (o *ActivityLogContext) GetIsFigmaSupportTeamActionOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFigmaSupportTeamAction, true
}

// SetIsFigmaSupportTeamAction sets field value
func (o *ActivityLogContext) SetIsFigmaSupportTeamAction(v bool) {
	o.IsFigmaSupportTeamAction = v
}

// GetOrgId returns the OrgId field value
func (o *ActivityLogContext) GetOrgId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrgId
}

// GetOrgIdOk returns a tuple with the OrgId field value
// and a boolean to check if the value has been set.
func (o *ActivityLogContext) GetOrgIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrgId, true
}

// SetOrgId sets field value
func (o *ActivityLogContext) SetOrgId(v string) {
	o.OrgId = v
}

// GetTeamId returns the TeamId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ActivityLogContext) GetTeamId() string {
	if o == nil || o.TeamId.Get() == nil {
		var ret string
		return ret
	}

	return *o.TeamId.Get()
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ActivityLogContext) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TeamId.Get(), o.TeamId.IsSet()
}

// SetTeamId sets field value
func (o *ActivityLogContext) SetTeamId(v string) {
	o.TeamId.Set(&v)
}

func (o ActivityLogContext) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActivityLogContext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client_name"] = o.ClientName.Get()
	toSerialize["ip_address"] = o.IpAddress
	toSerialize["is_figma_support_team_action"] = o.IsFigmaSupportTeamAction
	toSerialize["org_id"] = o.OrgId
	toSerialize["team_id"] = o.TeamId.Get()
	return toSerialize, nil
}

func (o *ActivityLogContext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"client_name",
		"ip_address",
		"is_figma_support_team_action",
		"org_id",
		"team_id",
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

	varActivityLogContext := _ActivityLogContext{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varActivityLogContext)

	if err != nil {
		return err
	}

	*o = ActivityLogContext(varActivityLogContext)

	return err
}

type NullableActivityLogContext struct {
	value *ActivityLogContext
	isSet bool
}

func (v NullableActivityLogContext) Get() *ActivityLogContext {
	return v.value
}

func (v *NullableActivityLogContext) Set(val *ActivityLogContext) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityLogContext) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityLogContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityLogContext(val *ActivityLogContext) *NullableActivityLogContext {
	return &NullableActivityLogContext{value: val, isSet: true}
}

func (v NullableActivityLogContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityLogContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


