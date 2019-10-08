package helper

import (
	"go_crud/models"
	"strings"
)

func EnvReplace(str string, env models.Env) string {
	for _, v := range env.Dict {
		str = strings.Replace(str, v.Key, v.Value, -1)
	}

	str = strings.Replace(str, "_RANDOM_NAME_", RandomName(), -1)
	str = strings.Replace(str, "_RANDOM_NUMBER_", RandomNumber(1, 10000000), -1)

	return str

}
