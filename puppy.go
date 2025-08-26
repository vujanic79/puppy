package puppy

import (
	"github.com/vujanic79/dog"
)

func Bark() string {
	return dog.WhenGrownUp("Woof!")
}

func Barks() string {
	return dog.WhenGrownUp("Woof! Woof! Woof!")
}

func From11() string {
	return "Hello, I'm from version 1.0.0"
}

func From12() string {
	return "Hello, I'm from version 1.2.0"
}
