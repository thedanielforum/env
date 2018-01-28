package env

import (
	"log"
	"os"
)

// Env instance
type Env int

const (
	DEV = iota
	TEST
	STAGE
	PROD
)

var envString = [...]string{
	"dev",
	"test",
	"stage",
	"prod",
}

// String returns env as string
func (e Env) String() string {
	return envString[int(e)]
}

var env Env

// SetEnv sets current env
func SetEnv(k string) {
	found := false
	for i, v := range envString {
		if k == v {
			found = true
			env = Env(i)
		}
	}

	if !found {
		log.Fatalf("env error: provided value [%s] is not a valid env", k)
	}
}

// GetEnv returns current env
func GetEnv() Env {
	return env
}

// IsDev env
func IsDev() bool {
	return GetEnv() == DEV
}

// IsTest env
func IsTest() bool {
	return GetEnv() == TEST
}

// IsStage env
func IsStage() bool {
	return GetEnv() == STAGE
}

// IsProd env
func IsProd() bool {
	return GetEnv() == PROD
}

// InitEnv sets the env from RUN_ENV environment variable
func InitEnv() {
	e := os.Getenv("RUN_ENV")
	if e != "" {
		SetEnv(e)
	} else {
		// set the default env to DEV
		SetEnv(envString[0])
	}
}
