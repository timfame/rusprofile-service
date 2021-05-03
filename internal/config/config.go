package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Rusprofile *Rusprofile `json:"rusprofile"`
	Redis      *Redis      `json:"redis"`
	Cache      *Cache      `json:"cache"`
	Grpc       *Grpc       `json:"grpc"`
}

func New() Config {
	return Config{
		Rusprofile: &Rusprofile{},
		Redis:      &Redis{},
		Cache:      &Cache{},
		Grpc:       &Grpc{},
	}
}

func FromJSONFile(filepath string) (Config, error) {
	jsonFile, err := os.Open(filepath)
	if err != nil {
		return Config{}, err
	}
	defer jsonFile.Close()

	data, _ := ioutil.ReadAll(jsonFile)
	result := New()
	err = json.Unmarshal(data, &result)
	return result, err
}

func (c *Config) Init(withRequiredRedis bool) error {
	if err := c.Rusprofile.Init(); err != nil {
		return err
	}
	if err := c.Redis.Init(); err != nil && withRequiredRedis {
		return err
	}
	if err := c.Cache.Init(); err != nil {
		return err
	}
	if err := c.Grpc.Init(); err != nil {
		return err
	}
	return nil
}
