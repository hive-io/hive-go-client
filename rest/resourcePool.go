package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

// ResourcePoolMember is a member entry in a resource pool
type ResourcePoolMember struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// ResourcePool describes a resource pool record from the rest api
type ResourcePool struct {
	ID          string               `json:"id,omitempty"`
	Name        string               `json:"name"`
	Type        string               `json:"type"`
	Members     []ResourcePoolMember `json:"members,omitempty"`
	Description string               `json:"description,omitempty"`
	Tags        []string             `json:"tags,omitempty"`
}

// ResourcePoolConfig holds the fields that can be updated via the configuration endpoint
type ResourcePoolConfig struct {
	Name        string   `json:"name,omitempty"`
	Description string   `json:"description,omitempty"`
	Tags        []string `json:"tags,omitempty"`
}

// ListResourcePools returns all resource pools
func (client *Client) ListResourcePools(query string) ([]ResourcePool, error) {
	var pools []ResourcePool
	path := "resourcePools"
	if query != "" {
		path += "?" + query
	}
	body, err := client.request("GET", path, nil)
	if err != nil {
		return pools, err
	}
	err = json.Unmarshal(body, &pools)
	return pools, err
}

// GetResourcePool requests a resource pool by id
func (client *Client) GetResourcePool(id string) (*ResourcePool, error) {
	var pool *ResourcePool
	if id == "" {
		return pool, errors.New("id cannot be empty")
	}
	body, err := client.request("GET", "resourcePool/"+id, nil)
	if err != nil {
		return pool, err
	}
	err = json.Unmarshal(body, &pool)
	return pool, err
}

// GetResourcePoolByName finds a resource pool by name
func (client *Client) GetResourcePoolByName(name string) (*ResourcePool, error) {
	pools, err := client.ListResourcePools("name=" + url.QueryEscape(name))
	if err != nil {
		return nil, err
	}
	for _, pool := range pools {
		if pool.Name == name {
			return &pool, nil
		}
	}
	return nil, errors.New("resource pool not found")
}

// Create creates a new resource pool
func (pool *ResourcePool) Create(client *Client) (string, error) {
	jsonValue, _ := json.Marshal(pool)
	body, err := client.request("POST", "resourcePools", jsonValue)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// Delete removes a resource pool
func (pool *ResourcePool) Delete(client *Client) error {
	if pool.ID == "" || client == nil {
		return errors.New("invalid resource pool")
	}
	_, err := client.request("DELETE", "resourcePool/"+pool.ID, nil)
	return err
}

// UpdateConfiguration updates the name, description, and tags of a resource pool
func (pool *ResourcePool) UpdateConfiguration(client *Client, config ResourcePoolConfig) (string, error) {
	if pool.ID == "" || client == nil {
		return "", errors.New("invalid resource pool")
	}
	jsonValue, _ := json.Marshal(config)
	body, err := client.request("PUT", "resourcePool/"+pool.ID+"/configuration", jsonValue)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// AddMember adds a member to the resource pool
func (pool *ResourcePool) AddMember(client *Client, memberID, source, name string) (string, error) {
	if pool.ID == "" || client == nil {
		return "", errors.New("invalid resource pool")
	}
	if memberID == "" {
		return "", errors.New("memberID cannot be empty")
	}
	payload := map[string]string{"memberId": memberID}
	if source != "" {
		payload["source"] = source
	}
	if name != "" {
		payload["name"] = name
	}
	jsonValue, _ := json.Marshal(payload)
	body, err := client.request("PUT", "resourcePool/"+pool.ID+"/addMember", jsonValue)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// RemoveMember removes a member from the resource pool
func (pool *ResourcePool) RemoveMember(client *Client, memberID string) (string, error) {
	if pool.ID == "" || client == nil {
		return "", errors.New("invalid resource pool")
	}
	if memberID == "" {
		return "", errors.New("memberID cannot be empty")
	}
	path := fmt.Sprintf("resourcePool/%s/removeMember?memberId=%s", pool.ID, url.QueryEscape(memberID))
	body, err := client.request("DELETE", path, nil)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
