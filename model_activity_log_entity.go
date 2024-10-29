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
	"gopkg.in/validator.v2"
	"fmt"
)

// ActivityLogEntity - The resource the actor took the action on. It can be a user, file, project or other resource types.
type ActivityLogEntity struct {
	ActivityLogFileEntity *ActivityLogFileEntity
	ActivityLogFileRepoEntity *ActivityLogFileRepoEntity
	ActivityLogOrgEntity *ActivityLogOrgEntity
	ActivityLogPluginEntity *ActivityLogPluginEntity
	ActivityLogProjectEntity *ActivityLogProjectEntity
	ActivityLogTeamEntity *ActivityLogTeamEntity
	ActivityLogUserEntity *ActivityLogUserEntity
	ActivityLogWidgetEntity *ActivityLogWidgetEntity
	ActivityLogWorkspaceEntity *ActivityLogWorkspaceEntity
}

// ActivityLogFileEntityAsActivityLogEntity is a convenience function that returns ActivityLogFileEntity wrapped in ActivityLogEntity
func ActivityLogFileEntityAsActivityLogEntity(v *ActivityLogFileEntity) ActivityLogEntity {
	return ActivityLogEntity{
		ActivityLogFileEntity: v,
	}
}

// ActivityLogFileRepoEntityAsActivityLogEntity is a convenience function that returns ActivityLogFileRepoEntity wrapped in ActivityLogEntity
func ActivityLogFileRepoEntityAsActivityLogEntity(v *ActivityLogFileRepoEntity) ActivityLogEntity {
	return ActivityLogEntity{
		ActivityLogFileRepoEntity: v,
	}
}

// ActivityLogOrgEntityAsActivityLogEntity is a convenience function that returns ActivityLogOrgEntity wrapped in ActivityLogEntity
func ActivityLogOrgEntityAsActivityLogEntity(v *ActivityLogOrgEntity) ActivityLogEntity {
	return ActivityLogEntity{
		ActivityLogOrgEntity: v,
	}
}

// ActivityLogPluginEntityAsActivityLogEntity is a convenience function that returns ActivityLogPluginEntity wrapped in ActivityLogEntity
func ActivityLogPluginEntityAsActivityLogEntity(v *ActivityLogPluginEntity) ActivityLogEntity {
	return ActivityLogEntity{
		ActivityLogPluginEntity: v,
	}
}

// ActivityLogProjectEntityAsActivityLogEntity is a convenience function that returns ActivityLogProjectEntity wrapped in ActivityLogEntity
func ActivityLogProjectEntityAsActivityLogEntity(v *ActivityLogProjectEntity) ActivityLogEntity {
	return ActivityLogEntity{
		ActivityLogProjectEntity: v,
	}
}

// ActivityLogTeamEntityAsActivityLogEntity is a convenience function that returns ActivityLogTeamEntity wrapped in ActivityLogEntity
func ActivityLogTeamEntityAsActivityLogEntity(v *ActivityLogTeamEntity) ActivityLogEntity {
	return ActivityLogEntity{
		ActivityLogTeamEntity: v,
	}
}

// ActivityLogUserEntityAsActivityLogEntity is a convenience function that returns ActivityLogUserEntity wrapped in ActivityLogEntity
func ActivityLogUserEntityAsActivityLogEntity(v *ActivityLogUserEntity) ActivityLogEntity {
	return ActivityLogEntity{
		ActivityLogUserEntity: v,
	}
}

// ActivityLogWidgetEntityAsActivityLogEntity is a convenience function that returns ActivityLogWidgetEntity wrapped in ActivityLogEntity
func ActivityLogWidgetEntityAsActivityLogEntity(v *ActivityLogWidgetEntity) ActivityLogEntity {
	return ActivityLogEntity{
		ActivityLogWidgetEntity: v,
	}
}

// ActivityLogWorkspaceEntityAsActivityLogEntity is a convenience function that returns ActivityLogWorkspaceEntity wrapped in ActivityLogEntity
func ActivityLogWorkspaceEntityAsActivityLogEntity(v *ActivityLogWorkspaceEntity) ActivityLogEntity {
	return ActivityLogEntity{
		ActivityLogWorkspaceEntity: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ActivityLogEntity) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActivityLogFileEntity
	err = newStrictDecoder(data).Decode(&dst.ActivityLogFileEntity)
	if err == nil {
		jsonActivityLogFileEntity, _ := json.Marshal(dst.ActivityLogFileEntity)
		if string(jsonActivityLogFileEntity) == "{}" { // empty struct
			dst.ActivityLogFileEntity = nil
		} else {
			if err = validator.Validate(dst.ActivityLogFileEntity); err != nil {
				dst.ActivityLogFileEntity = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogFileEntity = nil
	}

	// try to unmarshal data into ActivityLogFileRepoEntity
	err = newStrictDecoder(data).Decode(&dst.ActivityLogFileRepoEntity)
	if err == nil {
		jsonActivityLogFileRepoEntity, _ := json.Marshal(dst.ActivityLogFileRepoEntity)
		if string(jsonActivityLogFileRepoEntity) == "{}" { // empty struct
			dst.ActivityLogFileRepoEntity = nil
		} else {
			if err = validator.Validate(dst.ActivityLogFileRepoEntity); err != nil {
				dst.ActivityLogFileRepoEntity = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogFileRepoEntity = nil
	}

	// try to unmarshal data into ActivityLogOrgEntity
	err = newStrictDecoder(data).Decode(&dst.ActivityLogOrgEntity)
	if err == nil {
		jsonActivityLogOrgEntity, _ := json.Marshal(dst.ActivityLogOrgEntity)
		if string(jsonActivityLogOrgEntity) == "{}" { // empty struct
			dst.ActivityLogOrgEntity = nil
		} else {
			if err = validator.Validate(dst.ActivityLogOrgEntity); err != nil {
				dst.ActivityLogOrgEntity = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogOrgEntity = nil
	}

	// try to unmarshal data into ActivityLogPluginEntity
	err = newStrictDecoder(data).Decode(&dst.ActivityLogPluginEntity)
	if err == nil {
		jsonActivityLogPluginEntity, _ := json.Marshal(dst.ActivityLogPluginEntity)
		if string(jsonActivityLogPluginEntity) == "{}" { // empty struct
			dst.ActivityLogPluginEntity = nil
		} else {
			if err = validator.Validate(dst.ActivityLogPluginEntity); err != nil {
				dst.ActivityLogPluginEntity = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogPluginEntity = nil
	}

	// try to unmarshal data into ActivityLogProjectEntity
	err = newStrictDecoder(data).Decode(&dst.ActivityLogProjectEntity)
	if err == nil {
		jsonActivityLogProjectEntity, _ := json.Marshal(dst.ActivityLogProjectEntity)
		if string(jsonActivityLogProjectEntity) == "{}" { // empty struct
			dst.ActivityLogProjectEntity = nil
		} else {
			if err = validator.Validate(dst.ActivityLogProjectEntity); err != nil {
				dst.ActivityLogProjectEntity = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogProjectEntity = nil
	}

	// try to unmarshal data into ActivityLogTeamEntity
	err = newStrictDecoder(data).Decode(&dst.ActivityLogTeamEntity)
	if err == nil {
		jsonActivityLogTeamEntity, _ := json.Marshal(dst.ActivityLogTeamEntity)
		if string(jsonActivityLogTeamEntity) == "{}" { // empty struct
			dst.ActivityLogTeamEntity = nil
		} else {
			if err = validator.Validate(dst.ActivityLogTeamEntity); err != nil {
				dst.ActivityLogTeamEntity = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogTeamEntity = nil
	}

	// try to unmarshal data into ActivityLogUserEntity
	err = newStrictDecoder(data).Decode(&dst.ActivityLogUserEntity)
	if err == nil {
		jsonActivityLogUserEntity, _ := json.Marshal(dst.ActivityLogUserEntity)
		if string(jsonActivityLogUserEntity) == "{}" { // empty struct
			dst.ActivityLogUserEntity = nil
		} else {
			if err = validator.Validate(dst.ActivityLogUserEntity); err != nil {
				dst.ActivityLogUserEntity = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogUserEntity = nil
	}

	// try to unmarshal data into ActivityLogWidgetEntity
	err = newStrictDecoder(data).Decode(&dst.ActivityLogWidgetEntity)
	if err == nil {
		jsonActivityLogWidgetEntity, _ := json.Marshal(dst.ActivityLogWidgetEntity)
		if string(jsonActivityLogWidgetEntity) == "{}" { // empty struct
			dst.ActivityLogWidgetEntity = nil
		} else {
			if err = validator.Validate(dst.ActivityLogWidgetEntity); err != nil {
				dst.ActivityLogWidgetEntity = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogWidgetEntity = nil
	}

	// try to unmarshal data into ActivityLogWorkspaceEntity
	err = newStrictDecoder(data).Decode(&dst.ActivityLogWorkspaceEntity)
	if err == nil {
		jsonActivityLogWorkspaceEntity, _ := json.Marshal(dst.ActivityLogWorkspaceEntity)
		if string(jsonActivityLogWorkspaceEntity) == "{}" { // empty struct
			dst.ActivityLogWorkspaceEntity = nil
		} else {
			if err = validator.Validate(dst.ActivityLogWorkspaceEntity); err != nil {
				dst.ActivityLogWorkspaceEntity = nil
			} else {
				match++
			}
		}
	} else {
		dst.ActivityLogWorkspaceEntity = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActivityLogFileEntity = nil
		dst.ActivityLogFileRepoEntity = nil
		dst.ActivityLogOrgEntity = nil
		dst.ActivityLogPluginEntity = nil
		dst.ActivityLogProjectEntity = nil
		dst.ActivityLogTeamEntity = nil
		dst.ActivityLogUserEntity = nil
		dst.ActivityLogWidgetEntity = nil
		dst.ActivityLogWorkspaceEntity = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ActivityLogEntity)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ActivityLogEntity)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ActivityLogEntity) MarshalJSON() ([]byte, error) {
	if src.ActivityLogFileEntity != nil {
		return json.Marshal(&src.ActivityLogFileEntity)
	}

	if src.ActivityLogFileRepoEntity != nil {
		return json.Marshal(&src.ActivityLogFileRepoEntity)
	}

	if src.ActivityLogOrgEntity != nil {
		return json.Marshal(&src.ActivityLogOrgEntity)
	}

	if src.ActivityLogPluginEntity != nil {
		return json.Marshal(&src.ActivityLogPluginEntity)
	}

	if src.ActivityLogProjectEntity != nil {
		return json.Marshal(&src.ActivityLogProjectEntity)
	}

	if src.ActivityLogTeamEntity != nil {
		return json.Marshal(&src.ActivityLogTeamEntity)
	}

	if src.ActivityLogUserEntity != nil {
		return json.Marshal(&src.ActivityLogUserEntity)
	}

	if src.ActivityLogWidgetEntity != nil {
		return json.Marshal(&src.ActivityLogWidgetEntity)
	}

	if src.ActivityLogWorkspaceEntity != nil {
		return json.Marshal(&src.ActivityLogWorkspaceEntity)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ActivityLogEntity) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ActivityLogFileEntity != nil {
		return obj.ActivityLogFileEntity
	}

	if obj.ActivityLogFileRepoEntity != nil {
		return obj.ActivityLogFileRepoEntity
	}

	if obj.ActivityLogOrgEntity != nil {
		return obj.ActivityLogOrgEntity
	}

	if obj.ActivityLogPluginEntity != nil {
		return obj.ActivityLogPluginEntity
	}

	if obj.ActivityLogProjectEntity != nil {
		return obj.ActivityLogProjectEntity
	}

	if obj.ActivityLogTeamEntity != nil {
		return obj.ActivityLogTeamEntity
	}

	if obj.ActivityLogUserEntity != nil {
		return obj.ActivityLogUserEntity
	}

	if obj.ActivityLogWidgetEntity != nil {
		return obj.ActivityLogWidgetEntity
	}

	if obj.ActivityLogWorkspaceEntity != nil {
		return obj.ActivityLogWorkspaceEntity
	}

	// all schemas are nil
	return nil
}

type NullableActivityLogEntity struct {
	value *ActivityLogEntity
	isSet bool
}

func (v NullableActivityLogEntity) Get() *ActivityLogEntity {
	return v.value
}

func (v *NullableActivityLogEntity) Set(val *ActivityLogEntity) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityLogEntity) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityLogEntity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityLogEntity(val *ActivityLogEntity) *NullableActivityLogEntity {
	return &NullableActivityLogEntity{value: val, isSet: true}
}

func (v NullableActivityLogEntity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityLogEntity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


