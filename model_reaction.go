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
	"time"
	"bytes"
	"fmt"
)

// checks if the Reaction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Reaction{}

// Reaction A reaction left by a user.
type Reaction struct {
	// The user who left the reaction.
	User User `json:"user"`
	// The emoji type of reaction as shortcode (e.g. `:heart:`, `:+1::skin-tone-2:`). The list of accepted emoji shortcodes can be found in [this file](https://raw.githubusercontent.com/missive/emoji-mart/main/packages/emoji-mart-data/sets/14/native.json) under the top-level emojis and aliases fields, with optional skin tone modifiers when applicable.
	Emoji string `json:"emoji"`
	// The UTC ISO 8601 time at which the reaction was left.
	CreatedAt time.Time `json:"created_at"`
}

type _Reaction Reaction

// NewReaction instantiates a new Reaction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReaction(user User, emoji string, createdAt time.Time) *Reaction {
	this := Reaction{}
	this.User = user
	this.Emoji = emoji
	this.CreatedAt = createdAt
	return &this
}

// NewReactionWithDefaults instantiates a new Reaction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReactionWithDefaults() *Reaction {
	this := Reaction{}
	return &this
}

// GetUser returns the User field value
func (o *Reaction) GetUser() User {
	if o == nil {
		var ret User
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *Reaction) GetUserOk() (*User, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *Reaction) SetUser(v User) {
	o.User = v
}

// GetEmoji returns the Emoji field value
func (o *Reaction) GetEmoji() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Emoji
}

// GetEmojiOk returns a tuple with the Emoji field value
// and a boolean to check if the value has been set.
func (o *Reaction) GetEmojiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Emoji, true
}

// SetEmoji sets field value
func (o *Reaction) SetEmoji(v string) {
	o.Emoji = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Reaction) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Reaction) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Reaction) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o Reaction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Reaction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	toSerialize["emoji"] = o.Emoji
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *Reaction) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
		"emoji",
		"created_at",
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

	varReaction := _Reaction{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReaction)

	if err != nil {
		return err
	}

	*o = Reaction(varReaction)

	return err
}

type NullableReaction struct {
	value *Reaction
	isSet bool
}

func (v NullableReaction) Get() *Reaction {
	return v.value
}

func (v *NullableReaction) Set(val *Reaction) {
	v.value = val
	v.isSet = true
}

func (v NullableReaction) IsSet() bool {
	return v.isSet
}

func (v *NullableReaction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReaction(val *Reaction) *NullableReaction {
	return &NullableReaction{value: val, isSet: true}
}

func (v NullableReaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReaction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

