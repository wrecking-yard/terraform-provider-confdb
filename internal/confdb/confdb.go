package confdb

import (
	"embed"
	"encoding/json"
	"reflect"

	"codeberg.org/wrecking-yard/terraform-provider-confdb/internal/confdb/data"
	"github.com/itchyny/gojq"
)

type ResourceGroup struct {
	ID   string
	Name string
}

type PrivateDNSZones struct {
	ID            string
	Name          string
	ResourceGroup *ResourceGroup
}

type Subnet struct {
	ID      string
	Name    string
	Range   string
	Default bool
}

type VNet struct {
	ID            string
	Name          string
	Subnets       map[string]Subnet
	ResourceGroup *ResourceGroup
}

type StorageAccount struct {
	ID            string
	Name          string
	ResourceGroup *ResourceGroup
}

type AKS struct {
	ID            string
	Name          string
	Default       bool
	ResourceGroup *ResourceGroup
}

type Region struct {
	ID              string
	Name            string
	Default         bool
	ResourceGroups  map[string]ResourceGroup
	PrivateDNSZones map[string]PrivateDNSZones
	VNets           map[string]VNet `json:"vnets"`
	StorageAccounts map[string]StorageAccount
	AKSes           map[string]AKS
}

type Environment struct {
	Name    string
	Regions map[string]Region `json:"regions"`
}

type Subscription struct {
	ID   string
	Name string
	Envs map[string]Environment `json:"envs"`
}

type ConfDB struct {
	subscription  string
	environment   string
	region        string
	obj           map[string]any
	Subscriptions map[string]Subscription `json:"subscriptions"`
}

func (confdb *ConfDB) Init(fs embed.FS, fileName, subscription, environment, region string) (bool, error) {
	_fs := data.RawSubscriptions
	if !reflect.DeepEqual(fs, embed.FS{}) {
		_fs = fs
	}

	_fileName := "subscriptions.json"
	if fileName != "" {
		_fileName = fileName
	}

	rawSubscriptions, _ := _fs.ReadFile(_fileName)
	err := json.Unmarshal(rawSubscriptions, confdb)
	if err != nil {
		return true, err
	}

	confdb.subscription = subscription
	confdb.environment = environment
	confdb.region = region

	_json, err := json.Marshal(confdb)
	if err != nil {
		return true, err
	}

	err = json.Unmarshal(_json, &confdb.obj)
	if err != nil {
		return true, err
	}

	return true, nil
}

func (confdb ConfDB) DefaultVNet() (map[string]any, error) {

	query, err := gojq.Parse(
		".subscriptions." + confdb.subscription + ".envs." + confdb.environment + ".regions." + confdb.region + ".vnets | map(.)[0]",
	)
	if err != nil {
		return map[string]any{}, err
	}

	// https://github.com/itchyny/gojq?tab=readme-ov-file#usage-as-a-library
	iter := query.Run(confdb.obj)
	vnet := map[string]any{}
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			if err, ok := err.(*gojq.HaltError); ok && err.Value() == nil {
				break
			}
			return map[string]any{}, err
		}
		vnet, _ = v.(map[string]any)
		break
	}
	return vnet, nil
}
