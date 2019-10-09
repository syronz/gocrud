package helper

import (
	"go_crud/models"
)

// ParseEnv can handle complicated environment design of variables inside basic.json and it uses interfaces for
// encoding the content
func ParseEnv(env interface{}) map[string]models.Env {
	envs := map[string]models.Env{}

	switch v := env.(type) {
	case []interface{}:
		for _, packs := range v {
			pack := packs.(map[string]interface{})
			for j, terms := range pack {
				var obj models.Env
				term := terms.(map[string]interface{})
				for k, word := range term {
					w := word.(string)
					dict := models.Dict{k, w}
					obj.Dict = append(obj.Dict, dict)
				}
				envs[j] = obj
			}
		}
	}
	return envs

}
