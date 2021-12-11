//go:build kubernetes

package main

import (
	"cueme/pkg/platform"
	"log"
)

func init() {
	const secretsDir = "/cueme-secrets"
	err := platform.ExportSecretsToEnvVars(secretsDir)
	if err != nil {
		log.Fatalln(err)
	}
}
