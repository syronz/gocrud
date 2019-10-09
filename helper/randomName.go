package helper

import (
	"github.com/goombaio/namegenerator"
	"time"
)

// RandomName generate random names with dash betweens
func RandomName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	return nameGenerator.Generate()
}
