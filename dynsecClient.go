package dynsecgo

import (
	"encoding/json"
	"fmt"
)

// ExistClient
func (ds *DynSec) ExistClient(clientName string) (bool, error) {

	jsonCommand, err := ds.command.GetClient(clientName)
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
		if *newResponse.Responses[0].Error == "Client not found" {
			return false, nil
		}
	}

	if newResponse.Responses[0].Data != nil {
		if newResponse.Responses[0].Data.Client.Username == clientName {
			return true, nil
		}
	}

	return false, fmt.Errorf("undefinid error for response: %s", string(result))
}

// AddBaseClient
func (ds *DynSec) AddBaseClient(clientName, password, clientId string) error {
	return ds.AddClient(clientName, password, clientId, "", "", []string{}, []int{}, []string{}, []int{})
}

// AddClient
func (ds *DynSec) AddClient(clientName, password, clientID, textName, textdescription string, roleName []string, rolePriority []int, groupName []string, groupPriority []int) error {

	jsonCommand, err := ds.command.AddClient(clientName, password, clientID, textName, textdescription, roleName, rolePriority, groupName, groupPriority)
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

// GetClientJson
func (ds *DynSec) GetClientJson(clientName string) (string, error) {

	jsonCommand, err := ds.command.GetClient(clientName)
	if err != nil {
		return "", err
	}

	result, err := ds.sendCommand(jsonCommand)
	if err != nil {
		return "", err
	}
	var newResponse response
	if err := json.Unmarshal(result, &newResponse); err != nil {
		return "", err
	}
	if newResponse.Responses[0].Error != nil {
		if len(*newResponse.Responses[0].Error) > 1 {
			return "", fmt.Errorf(*newResponse.Responses[0].Error)
		}
	}

	return string(result), nil
}

// GetClient
func (ds *DynSec) GetClient(clientName string) (Client, error) {

	jsonCommand, err := ds.command.GetClient(clientName)
	if err != nil {
		return Client{}, err
	}

	result, err := ds.sendCommand(jsonCommand)
	if err != nil {
		return Client{}, err
	}
	var newResponse response
	if err := json.Unmarshal(result, &newResponse); err != nil {
		return Client{}, err
	}
	if newResponse.Responses[0].Error != nil {
		if len(*newResponse.Responses[0].Error) > 1 {
			return Client{}, fmt.Errorf(*newResponse.Responses[0].Error)
		}
	}

	if len(newResponse.Responses) > 1 {
		return Client{}, fmt.Errorf("undefined error")
	}

	client := &Client{
		UserName:        newResponse.Responses[0].Data.Client.Username,
		Disabled:        newResponse.Responses[0].Data.Client.Disabled,
		Roles:           newResponse.Responses[0].Data.Client.Roles,
		Groups:          newResponse.Responses[0].Data.Client.Groups,
		Textname:        newResponse.Responses[0].Data.Client.Textname,
		Textdescription: newResponse.Responses[0].Data.Client.Textdescription,
	}

	return *client, nil
}

// DeleteClient
func (ds *DynSec) DeleteClient(clientName string) error {
	jsonCommand, err := ds.command.DeleteClient(clientName)
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
		if *newResponse.Responses[0].Error == "Client not found" {
			return fmt.Errorf(*newResponse.Responses[0].Error)
		}
	}

	return nil
}

// DisableClient
func (ds *DynSec) DisableClient(clientName string) error {
	jsonCommand, err := ds.command.DisableClient(clientName)
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
		if *newResponse.Responses[0].Error == "Client not found" {
			return fmt.Errorf(*newResponse.Responses[0].Error)
		}
	}

	return nil
}

// EnableClient
func (ds *DynSec) EnableClient(clientName string) error {
	jsonCommand, err := ds.command.EnableClient(clientName)
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
		if *newResponse.Responses[0].Error == "Client not found" {
			return fmt.Errorf(*newResponse.Responses[0].Error)
		}
	}

	return nil
}
