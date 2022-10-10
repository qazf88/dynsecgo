package dynsecgo

import (
	"encoding/json"
	"fmt"
)

// ExistRole
func (ds *DynSec) ExistRole(roleName string) (bool, error) {

	jsonCommand, err := ds.command.GetRole(roleName)
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
		if *newResponse.Responses[0].Error == "Role already exists" {
			return true, nil
		} else if *newResponse.Responses[0].Error == "Role not found" {
			return false, nil
		}
	}
	return false, fmt.Errorf("undefinid error for response: %s", string(result))
}

// GetRole
func (ds *DynSec) GetRoleJson(roleName string) ([]byte, error) {

	jsonCommand, err := ds.command.GetRole(roleName)
	if err != nil {
		return nil, err
	}

	result, err := ds.sendCommand(jsonCommand)
	if err != nil {
		return nil, err
	}

	var newResponse response

	if err := json.Unmarshal(result, &newResponse); err != nil {
		return nil, err
	}

	if newResponse.Responses[0].Error != nil {
		return nil, fmt.Errorf(*newResponse.Responses[0].Error)
	}

	return result, nil
}

// AddRole
func (ds *DynSec) AddRole(roleName, textName, textdescription string, aclType []string, topic []string, priority []int, allow []bool) error {

	if len(aclType) != len(topic) || len(aclType) != len(priority) || len(aclType) != len(allow) {
		return fmt.Errorf("the number of acls is not equal to the number of topic, priority or allow")
	}

	for _, _type := range aclType {
		switch _type {
		case AclType.PublishClientReceive:
			goto next
		case AclType.PublishClientSend:
			goto next
		case AclType.SubscribeLiteral:
			goto next
		case AclType.SubscribePattern:
			goto next
		case AclType.UnsubscribeLiteral:
			goto next
		case AclType.UnsubscribePattern:
			goto next
		default:
			return fmt.Errorf("incorrect acl type")

		}

	}

next:

	var acls []Acl

	for index, _type := range aclType {

		_acl := ds.command.Acl(_type, topic[index], priority[index], allow[index])
		acls = append(acls, _acl)
	}

	jsonCommand, err := ds.command.AddRole(roleName, textName, textdescription, acls)
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
