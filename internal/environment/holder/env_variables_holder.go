package holder

import (
	"os"

	"golang.org/x/exp/maps"
)

type envVariablesHolder struct {
	envVarMap map[string]string
}

func (holder *envVariablesHolder) Store(key, value string) {
	holder.envVarMap[key] = value
}

func (holder *envVariablesHolder) Get(key string) string {
	if val, exist := holder.envVarMap[key]; exist {
		return val
	} else {
		return os.Getenv(key)
	}
}

func (holder *envVariablesHolder) GetAll() map[string]string {
	return maps.Clone(holder.envVarMap)
}

var EnvVariablesHolder = &envVariablesHolder{map[string]string{}}
