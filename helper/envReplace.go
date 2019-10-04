package helper

import (
	"fmt"
	"go_crud/models"
	"strings"
)

func EnvReplace(str string, env models.Env) string {
	for i, v := range env.Dict {
		fmt.Printf("..........  %+v ... %+v\n", i, v)
		fmt.Println("∫∫∫∫∫∫∫∫∫∫∫∫∫∫∫∫∫∫∫ ", i, v)
		str = strings.Replace(str, v.Key, v.Value, -1)
	}

	return str

}
