package env

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetEnv(t *testing.T) {
	SetEnv("dev")
	assert.Equal(t, IsDev(), true)
	SetEnv("test")
	assert.Equal(t, IsTest(), true)
	SetEnv("stage")
	assert.Equal(t, IsStage(), true)
	SetEnv("prod")
	assert.Equal(t, IsProd(), true)
}

func TestGetEnv(t *testing.T) {
	assert.Equal(t, GetEnv(), env)
}

func TestIsDev(t *testing.T) {
	SetEnv("dev")
	assert.Equal(t, IsDev(), true)
}

func TestIsTest(t *testing.T) {
	SetEnv("test")
	assert.Equal(t, IsTest(), true)
}

func TestIsStage(t *testing.T) {
	SetEnv("stage")
	assert.Equal(t, IsStage(), true)
}

func TestIsProd(t *testing.T) {
	SetEnv("prod")
	assert.Equal(t, IsProd(), true)
}

func TestInitEnv(t *testing.T) {
	// default env should be DEV when no value is provided
	InitEnv()
	assert.Equal(t, IsDev(), true)
}
