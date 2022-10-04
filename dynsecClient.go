package dynsecgo

import (
	"encoding/json"
	"fmt"
)

// ExistClient
func (ds *DynSec) ExistClient(clientName string) (bool, error) {

	client := &commands{Commands: []command{{
		Command: "getClient", Username: clientName}}}

	jsonRquest, err := json.Marshal(client)
	if err != nil {
		return false, err
	}

	result, err := publishCommand(ds.mc, jsonRquest)
	if err != nil {
		return false, err
	}

	var newResponse response

	if err := json.Unmarshal(result, &newResponse); err != nil {
		return false, err
	}

	if newResponse.Responses[0].Error == "Client already exists" {
		return true, nil
	} else if newResponse.Responses[0].Error == "Client not found" {
		return false, nil
	}

	return false, fmt.Errorf("undefinid error for response: %s", string(result))
}

// AddClient
func (ds *DynSec) AddClient(clientName string, password string, roleName []string, rolePriority []int, groupName []string, groupPriority []int) error {

	newCommand := []command{{
		Command:  "createClient",
		Username: clientName,
		Password: password,
	}}

	if len(roleName) > 0 {
		var _roles []role
		for i, _role := range roleName {
			_roles = append(_roles, role{Rolename: _role, Priority: rolePriority[i]})
		}
		newCommand[0].Roles = _roles
	}

	if len(groupName) > 0 {
		var _groups []group
		for i, _group := range groupName {
			_groups = append(_groups, group{Groupname: _group, Priority: groupPriority[i]})
		}
		newCommand[0].Groups = _groups
	}

	client := &commands{Commands: newCommand}

	jsonRequest, err := json.Marshal(client)
	if err != nil {
		return err

	}
	result, err := publishCommand(ds.mc, jsonRequest)
	if err != nil {
		return err
	}
	var newResponse response
	if err := json.Unmarshal(result, &newResponse); err != nil {
		return err
	}

	if len(newResponse.Responses[0].Error) > 1 {
		return fmt.Errorf(newResponse.Responses[0].Error)
	}
	return nil
}

// GetClient
func (ds *DynSec) GetClient(clientName string) (string, error) {

	newCommand := &commands{Commands: []command{{Command: "getClient", Username: clientName}}}

	jsonRequest, err := json.Marshal(newCommand)
	if err != nil {
		return "", err
	}

	result, err := publishCommand(ds.mc, jsonRequest)
	if err != nil {
		return "", err
	}
	var newResponse response
	if err := json.Unmarshal(result, &newResponse); err != nil {
		return "", err
	}

	if len(newResponse.Responses[0].Error) > 1 {
		return "", fmt.Errorf(newResponse.Responses[0].Error)
	}

	return string(result), nil
}
