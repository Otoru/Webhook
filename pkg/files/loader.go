package files

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

func GetYamlFiles() ([]string, error) {
	result := make([]string, 0)

	workdir := viper.GetString("workdir")

	if _, err := os.Stat(workdir); err != nil {
		return result, err
	}

	handler := func(path string, spec os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !spec.IsDir() {
			switch filepath.Ext(path) {
			case ".yml", ".yaml":
				result = append(result, path)
				return nil
			default:
				return nil
			}
		}

		return nil
	}

	err := filepath.WalkDir(workdir, handler)

	if err != nil {
		return result, err
	}

	return result, nil
}