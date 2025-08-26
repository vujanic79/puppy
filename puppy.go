package puppy

import (
	"github.com/vujanic79/dog"
)

func Bark() string {
	return dog.WhenGrownUp("Woof!")
}

func Barks() string {
	return "Woof! Woof! Woof!"
}
