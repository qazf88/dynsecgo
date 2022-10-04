package dynsecgo

import (
	"encoding/json"
	"fmt"
)

// ExistGroup
func (ds *DynSec) ExistGroup(groupName string) (bool, error) {

	group := &commands{Commands: []command{{
		Command: "getGroup", Groupname: groupName}}}

	jsonRequest, err := json.Marshal(group)
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

	if newResponse.Responses[0].Error == "Group already exists" {
		return true, nil
	} else if newResponse.Responses[0].Error == "Group not found" {
		return false, nil
	}

	return false, fmt.Errorf("undefinid error for response: %s", string(result))
}

// AddGroup
func (ds *DynSec) AddGroup() {

}
