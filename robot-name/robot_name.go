// Package robotname solves the Robot Name problem from Exercism.
package robotname

import (
	"math/rand"
)

const maxNames = 26 * 26 * 10 * 10 * 10

var usedNames = make(map[string]bool)

// Robot represents a rather simple robot.
type Robot struct {
	name string
}

// Name returns the name of the robot.
func (r *Robot) Name() string {
	// not thread safe
	if r.name == "" {
		if len(usedNames) == maxNames {
			panic("all possible names were generated")
		}

		newName := randomName()
		for usedNames[newName] {
			newName = randomName()
		}

		r.name = newName
		usedNames[newName] = true
	}

	return r.name
}

// Reset re-initializes a robot.
func (r *Robot) Reset() {
	r.name = ""
}

func randomName() string {
	return letter() + letter() + digit() + digit() + digit()
}

func letter() string {
	return string('A' + rand.Int31n(26))
}

func digit() string {
	return string('0' + rand.Int31n(10))
}
