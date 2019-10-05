package helper

import (
	"go_crud/models"
	"strings"
)

func EnvReplace(str string, env models.Env) string {
	for _, v := range env.Dict {
		str = strings.Replace(str, v.Key, v.Value, -1)
	}

	return str

}
