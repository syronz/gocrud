package helper

import (
	"github.com/goombaio/namegenerator"
	"time"
)

func RandomName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	return nameGenerator.Generate()
}
