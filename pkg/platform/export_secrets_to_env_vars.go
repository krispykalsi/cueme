//go:build kubernetes

package platform

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ExportSecretsToEnvVars(secretsDir string) error {
	files, err := ioutil.ReadDir(secretsDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() && (file.Mode()&os.ModeSymlink) == os.ModeSymlink {
			continue
		}
		key := file.Name()
		path := filepath.Join(secretsDir, key)
		bytes, readErr := ioutil.ReadFile(path)
		if readErr != nil {
			if !strings.Contains(readErr.Error(), "is a directory") {
				return readErr
			}
		}
		value := string(bytes)
		envErr := os.Setenv(key, value)
		if envErr != nil {
			return envErr
		}
	}
	return nil
}
