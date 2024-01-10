package hunt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSharkHuntsSuccessfully(t *testing.T) {
	//Given
	shark := Shark{hungry: true, speed: 2}
	prey := Prey{speed: 1}

	// When
	resultado := shark.Hunt(&prey)

	//then
	assert.Nil(t, resultado)
	assert.False(t, shark.hungry)
	assert.True(t, shark.tired)
}

func TestSharkCannotHuntBecauseIsTired(t *testing.T) {
	//Given
	shark := Shark{tired: true}
	prey := Prey{}

	// When
	resultado := shark.Hunt(&prey)

	//then
	assert.Equal(t, resultado.Error(), "cannot hunt, i am really tired")
}

func TestSharkCannotHuntBecaisIsNotHungry(t *testing.T) {
	//Given
	shark := Shark{}
	prey := Prey{}

	// When
	resultado := shark.Hunt(&prey)

	//then
	assert.Equal(t, resultado.Error(), "cannot hunt, i am not hungry")
}

func TestSharkCannotReachThePrey(t *testing.T) {
	//Given
	shark := Shark{speed: 1, hungry: true}
	prey := Prey{speed: 2}

	// When
	resultado := shark.Hunt(&prey)

	//then
	assert.Equal(t, resultado.Error(), "could not catch it")
}

func TestSharkHuntNilPrey(t *testing.T) {
	//Given
	shark := Shark{speed: 1, hungry: true}

	// When
	resultado := shark.Hunt(nil)

	//then
	assert.Equal(t, resultado.Error(), "cannot hunt, i cannot find my prey")
}
