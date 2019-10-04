package helper

import (
	"go_crud/models"
)

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
					//fmt.Println(" INSIDE term %%%%%%%%% ", j, k, word, w)
				}
				envs[j] = obj
			}
		}
	}
	//fmt.Println("RRRRRRRRRRR +++++++++ ", envs["local"].Dict[1].Value)
	return envs

}
