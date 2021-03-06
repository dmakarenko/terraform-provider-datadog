/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// RoleResponseRelationships Relationships of the role object returned by the API.
type RoleResponseRelationships struct {
	Org        *RelationshipToOrganization  `json:"org,omitempty"`
	OtherOrgs  *RelationshipToOrganizations `json:"other_orgs,omitempty"`
	OtherRoles *RelationshipToRoles         `json:"other_roles,omitempty"`
	Roles      *RelationshipToRoles         `json:"roles,omitempty"`
}

// NewRoleResponseRelationships instantiates a new RoleResponseRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleResponseRelationships() *RoleResponseRelationships {
	this := RoleResponseRelationships{}
	return &this
}

// NewRoleResponseRelationshipsWithDefaults instantiates a new RoleResponseRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleResponseRelationshipsWithDefaults() *RoleResponseRelationships {
	this := RoleResponseRelationships{}
	return &this
}

// GetOrg returns the Org field value if set, zero value otherwise.
func (o *RoleResponseRelationships) GetOrg() RelationshipToOrganization {
	if o == nil || o.Org == nil {
		var ret RelationshipToOrganization
		return ret
	}
	return *o.Org
}

// GetOrgOk returns a tuple with the Org field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleResponseRelationships) GetOrgOk() (*RelationshipToOrganization, bool) {
	if o == nil || o.Org == nil {
		return nil, false
	}
	return o.Org, true
}

// HasOrg returns a boolean if a field has been set.
func (o *RoleResponseRelationships) HasOrg() bool {
	if o != nil && o.Org != nil {
		return true
	}

	return false
}

// SetOrg gets a reference to the given RelationshipToOrganization and assigns it to the Org field.
func (o *RoleResponseRelationships) SetOrg(v RelationshipToOrganization) {
	o.Org = &v
}

// GetOtherOrgs returns the OtherOrgs field value if set, zero value otherwise.
func (o *RoleResponseRelationships) GetOtherOrgs() RelationshipToOrganizations {
	if o == nil || o.OtherOrgs == nil {
		var ret RelationshipToOrganizations
		return ret
	}
	return *o.OtherOrgs
}

// GetOtherOrgsOk returns a tuple with the OtherOrgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleResponseRelationships) GetOtherOrgsOk() (*RelationshipToOrganizations, bool) {
	if o == nil || o.OtherOrgs == nil {
		return nil, false
	}
	return o.OtherOrgs, true
}

// HasOtherOrgs returns a boolean if a field has been set.
func (o *RoleResponseRelationships) HasOtherOrgs() bool {
	if o != nil && o.OtherOrgs != nil {
		return true
	}

	return false
}

// SetOtherOrgs gets a reference to the given RelationshipToOrganizations and assigns it to the OtherOrgs field.
func (o *RoleResponseRelationships) SetOtherOrgs(v RelationshipToOrganizations) {
	o.OtherOrgs = &v
}

// GetOtherRoles returns the OtherRoles field value if set, zero value otherwise.
func (o *RoleResponseRelationships) GetOtherRoles() RelationshipToRoles {
	if o == nil || o.OtherRoles == nil {
		var ret RelationshipToRoles
		return ret
	}
	return *o.OtherRoles
}

// GetOtherRolesOk returns a tuple with the OtherRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleResponseRelationships) GetOtherRolesOk() (*RelationshipToRoles, bool) {
	if o == nil || o.OtherRoles == nil {
		return nil, false
	}
	return o.OtherRoles, true
}

// HasOtherRoles returns a boolean if a field has been set.
func (o *RoleResponseRelationships) HasOtherRoles() bool {
	if o != nil && o.OtherRoles != nil {
		return true
	}

	return false
}

// SetOtherRoles gets a reference to the given RelationshipToRoles and assigns it to the OtherRoles field.
func (o *RoleResponseRelationships) SetOtherRoles(v RelationshipToRoles) {
	o.OtherRoles = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *RoleResponseRelationships) GetRoles() RelationshipToRoles {
	if o == nil || o.Roles == nil {
		var ret RelationshipToRoles
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleResponseRelationships) GetRolesOk() (*RelationshipToRoles, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *RoleResponseRelationships) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given RelationshipToRoles and assigns it to the Roles field.
func (o *RoleResponseRelationships) SetRoles(v RelationshipToRoles) {
	o.Roles = &v
}

func (o RoleResponseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Org != nil {
		toSerialize["org"] = o.Org
	}
	if o.OtherOrgs != nil {
		toSerialize["other_orgs"] = o.OtherOrgs
	}
	if o.OtherRoles != nil {
		toSerialize["other_roles"] = o.OtherRoles
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	return json.Marshal(toSerialize)
}

type NullableRoleResponseRelationships struct {
	value *RoleResponseRelationships
	isSet bool
}

func (v NullableRoleResponseRelationships) Get() *RoleResponseRelationships {
	return v.value
}

func (v *NullableRoleResponseRelationships) Set(val *RoleResponseRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleResponseRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleResponseRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleResponseRelationships(val *RoleResponseRelationships) *NullableRoleResponseRelationships {
	return &NullableRoleResponseRelationships{value: val, isSet: true}
}

func (v NullableRoleResponseRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleResponseRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
