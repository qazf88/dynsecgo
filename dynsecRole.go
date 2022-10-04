package dynsecgo

import (
	"encoding/json"
	"fmt"
)

// ExistRole
func (ds *DynSec) ExistRole(roleName string) (bool, error) {

	role := &commands{Commands: []command{{
		Command: "getRole", Rolename: roleName}}}

	jsonRequest, err := json.Marshal(role)
	if err != nil {
		return false, err
	}

	result, err := publishCommand(ds.mc, jsonRequest)
	if err != nil {
		return false, err
	}

	var newResponse response

	if err := json.Unmarshal(result, &newResponse); err != nil {
		return false, err
	}

	if newResponse.Responses[0].Error == "Role already exists" {
		return true, nil
	} else if newResponse.Responses[0].Error == "Role not found" {
		return false, nil
	}

	return false, fmt.Errorf("undefinid error for response: %s", string(result))
}

// AddRole
func (ds *DynSec) AddRole(roleName, name, description, subTopic, pubTopic string) {

}
