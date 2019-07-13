package client

import (
	"encoding/json"
	"errors"

	"github.com/ghodss/yaml"
)

type StoragePool struct {
	ID           string   `json:"id,omitempty"`
	Name         string   `json:"name"`
	Type         string   `json:"type"`
	Server       string   `json:"server"`
	Path         string   `json:"path"`
	Username     string   `json:"username,omitempty"`
	Password     string   `json:"password,omitempty"`
	Key          string   `json:"key,omitempty"`
	MountOptions []string `json:"mountOptions,omitempty"`
	Roles        []string `json:"roles,omitempty"`
	Tags         []string `json:"tags,omitempty"`
}

func (sp StoragePool) String() string {
	json, _ := json.MarshalIndent(sp, "", "  ")
	return string(json)
}

func (sp *StoragePool) ToJson() ([]byte, error) {
	return json.Marshal(sp)
}

func (sp *StoragePool) FromJson(data []byte) error {
	return json.Unmarshal(data, sp)
}

func (sp *StoragePool) ToYaml() ([]byte, error) {
	return yaml.Marshal(sp)
}

func (sp *StoragePool) FromYaml(data []byte) error {
	return yaml.Unmarshal(data, sp)
}

func (client *Client) ListStoragePools() ([]StoragePool, error) {
	var pools []StoragePool
	body, err := client.Request("GET", "storage/pools", nil)
	if err != nil {
		return pools, err
	}
	err = json.Unmarshal(body, &pools)
	return pools, err
}

func (client *Client) GetStoragePoolByName(name string) (*StoragePool, error) {
	var pools, err = client.ListStoragePools()
	if err != nil {
		return nil, err
	}
	for _, pool := range pools {
		if pool.Name == name {
			return &pool, nil
		}
	}
	return nil, errors.New("Storage Pool not found")
}

func (client *Client) GetStoragePool(id string) (StoragePool, error) {
	var pool StoragePool
	if id == "" {
		return pool, errors.New("id cannot be empty")
	}
	body, err := client.Request("GET", "storage/pool/"+id, nil)
	if err != nil {
		return pool, err
	}
	err = json.Unmarshal(body, &pool)
	return pool, err
}

func (pool *StoragePool) Create(client *Client) (string, error) {
	var result string
	jsonValue, _ := json.Marshal(pool)
	body, err := client.Request("POST", "storage/pools", jsonValue)
	if err == nil {
		result = string(body)
	}
	return result, err
}

func (pool *StoragePool) Delete(client *Client) error {
	if pool.ID == "" {
		return errors.New("id cannot be empty")
	}
	_, err := client.Request("DELETE", "storage/pool/"+pool.ID, nil)
	if err != nil {
		return err
	}
	return err
}