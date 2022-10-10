package dynsecgo

import (
	"encoding/json"
	"fmt"
)

type DynSecCommand struct {
}

// AddClient
func (dsc *DynSecCommand) AddClient(clientName, password, clientID, textName, textdescription string, roleName []string, rolePriority []int, groupName []string, groupPriority []int) ([]byte, error) {

	newCommand := []command{{
		Command:  "createClient",
		Username: clientName,
		Password: password,
	}}

	if len(roleName) > 0 {
		if len(roleName) != len(rolePriority) {
			return nil, fmt.Errorf("the number of roles is not equal to the number of priority roles")
		}
		var _roles []role
		for i, _role := range roleName {
			_roles = append(_roles, role{Rolename: _role, Priority: rolePriority[i]})
		}
		newCommand[0].Roles = _roles
	}

	if len(groupName) > 0 {
		if len(groupName) != len(groupPriority) {
			return nil, fmt.Errorf("the number of groups is not equal to the number of priority groups")
		}
		var _groups []group
		for i, _group := range groupName {
			_groups = append(_groups, group{Groupname: _group, Priority: groupPriority[i]})
		}
		newCommand[0].Groups = _groups
	}

	addClientCommand := &commands{Commands: newCommand}

	json, err := json.Marshal(addClientCommand)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// GetClient command
func (dsc *DynSecCommand) GetClient(clientName string) ([]byte, error) {

	newCommand := &commands{Commands: []command{{
		Command: "getClient", Username: clientName}}}

	json, err := json.Marshal(newCommand)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// DeleteClient command
func (dsc *DynSecCommand) DeleteClient(clientName string) ([]byte, error) {
	newCommand := &commands{Commands: []command{{
		Command: "deleteClient", Username: clientName}}}

	json, err := json.Marshal(newCommand)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// DisableClient
func (dsc *DynSecCommand) DisableClient(clientName string) ([]byte, error) {
	newCommand := &commands{Commands: []command{{
		Command: "disableClient", Username: clientName}}}

	json, err := json.Marshal(newCommand)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// EnableClient
func (dsc *DynSecCommand) EnableClient(clientName string) ([]byte, error) {
	newCommand := &commands{Commands: []command{{
		Command: "enableClient", Username: clientName}}}

	json, err := json.Marshal(newCommand)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// GetGroup
func (dsc *DynSecCommand) GetGroup(groupName string) ([]byte, error) {
	newCommand := &commands{Commands: []command{{
		Command: "getGroup", Username: groupName}}}

	json, err := json.Marshal(newCommand)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// AddGroup
func (dsc *DynSecCommand) AddGroup(groupName string, roleName []string, rolePriority []int) ([]byte, error) {

	newCommand := []command{{
		Command:   "createGroup",
		Groupname: groupName,
	}}

	if len(roleName) > 0 {
		if len(roleName) != len(rolePriority) {
			return nil, fmt.Errorf("the number of roles is not equal to the number of priority roles")
		}
		var _roles []role
		for i, _role := range roleName {
			_roles = append(_roles, role{Rolename: _role, Priority: rolePriority[i]})
		}
		newCommand[0].Roles = _roles
	}

	addGroupCommand := &commands{Commands: newCommand}

	json, err := json.Marshal(addGroupCommand)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// DeleteGroup command
func (dsc *DynSecCommand) DeleteGroup(groupName string) ([]byte, error) {
	newCommand := &commands{Commands: []command{{
		Command: "deleteGroup", Groupname: groupName}}}

	json, err := json.Marshal(newCommand)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// GetRole
func (dsc *DynSecCommand) GetRole(roleName string) ([]byte, error) {
	newCommand := &commands{Commands: []command{{
		Command: "getRole", Rolename: roleName}}}

	json, err := json.Marshal(newCommand)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// AddRole
func (dsc *DynSecCommand) AddRole(roleName, textName, textdescription string, acls []Acl) ([]byte, error) {

	_command := command{
		Command:  "createRole",
		Rolename: roleName,
		Acls:     acls,
	}

	if len(textName) > 0 {
		_command.Textname = textName
	}

	if len(textdescription) > 0 {
		_command.Textdescription = textdescription
	}

	newCommand := []command{_command}

	addGroupCommand := &commands{Commands: newCommand}

	json, err := json.Marshal(addGroupCommand)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// DeleteRole command
func (dsc *DynSecCommand) DeleteRole(roleName string) ([]byte, error) {
	newCommand := &commands{Commands: []command{{
		Command: "deleteRole", Rolename: roleName}}}

	json, err := json.Marshal(newCommand)
	if err != nil {
		return nil, err
	}

	return json, nil
}

// Acls
func (dsc *DynSecCommand) Acl(aclType, topic string, priority int, allow bool) Acl {

	newAcl := Acl{
		Acltype:  aclType,
		Topic:    topic,
		Priority: priority,
		Allow:    allow,
	}

	return newAcl
}
