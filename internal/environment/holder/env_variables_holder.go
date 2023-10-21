package holder

import (
	"os"
	"strings"
)

/*
 * Структура для задания и получения переменных окружения
 */
type envVariablesHolder struct {
}

func (holder *envVariablesHolder) Store(key, value string) error {
	return os.Setenv(key, value)
}

func (holder *envVariablesHolder) Get(key string) string {
	return os.Getenv(key)
}

func (holder *envVariablesHolder) GetAll() map[string]string {
	environ := os.Environ()
	m := map[string]string{}

	for _, s := range environ {
		splitArr := strings.Split(s, "=")
		m[splitArr[0]] = splitArr[1]
	}

	return m
}

var EnvVariablesHolder = &envVariablesHolder{}
