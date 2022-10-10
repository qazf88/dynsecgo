package dynsecgo

import (
	"encoding/json"
	"fmt"
)

// ExistGroup
func (ds *DynSec) ExistGroup(groupName string) (bool, error) {

	jsonCommand, err := ds.command.GetGroup(groupName)
	if err != nil {
		return false, err
	}

	result, err := ds.sendCommand(jsonCommand)
	if err != nil {
		return false, err
	}

	var newResponse response

	if err := json.Unmarshal(result, &newResponse); err != nil {
		return false, err
	}

	if newResponse.Responses[0].Error != nil {
		if *newResponse.Responses[0].Error == "Group already exists" {
			return true, nil
		} else if *newResponse.Responses[0].Error == "Invalid/missing groupname" {
			return false, nil
		}
	}
	return false, fmt.Errorf("undefinid error for response: %s", string(result))
}

// AddBaseGroup
func (ds *DynSec) AddBaseGroup(groupName string) error {
	return ds.AddGroup(groupName, []string{}, []int{})
}

// AddGroup
func (ds *DynSec) AddGroup(groupName string, roleName []string, rolePriority []int) error {

	jsonCommand, err := ds.command.AddGroup(groupName, roleName, rolePriority)
	if err != nil {
		return err
	}

	result, err := ds.sendCommand(jsonCommand)
	if err != nil {
		return err
	}

	var newResponse response
	if err := json.Unmarshal(result, &newResponse); err != nil {
		return err
	}

	if newResponse.Responses[0].Error != nil {
		return fmt.Errorf(*newResponse.Responses[0].Error)
	}

	return nil
}

// DeleteGroup
func (ds *DynSec) DeleteGroup(groupName string) error {

	jsonCommand, err := ds.command.DeleteGroup(groupName)
	if err != nil {
		return err
	}

	result, err := ds.sendCommand(jsonCommand)
	if err != nil {
		return err
	}

	var newResponse response
	if err := json.Unmarshal(result, &newResponse); err != nil {
		return err
	}

	if newResponse.Responses[0].Error != nil {
		return fmt.Errorf(*newResponse.Responses[0].Error)
	}

	return nil
}
