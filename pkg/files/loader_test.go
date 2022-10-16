package files

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"testing"

	"github.com/spf13/viper"
)

func TestGetSpecifications(t *testing.T) {
	t.Run("A valid document generates a correct specification", func(t *testing.T) {
		raw := []byte(`
apiVersion: webhook/v1
kind: Listener
metadata:
  name: example-listener
  labels:
    app: example
    version: v1
specs:
  - name: public-listener
    host: 0.0.0.0
    port: 3000
`,
		)

		doc := new(Document)
		doc.Raw = raw

		parameter := make([]*Document, 0)
		parameter = append(parameter, doc)

		specs, err := GetSpecifications(parameter)

		if err != nil {
			t.Fatal(err)
		}

		if len(specs) != 1 {
			t.Error("We expected a specification here")
		}
	})

	t.Run("An invalid specification generates an error", func(t *testing.T) {
		raw := []byte(`
invalidKey: anotherValue
  kind: API
`,
		)

		doc := new(Document)
		doc.Raw = raw

		parameter := make([]*Document, 0)
		parameter = append(parameter, doc)

		_, err := GetSpecifications(parameter)

		if err == nil {
			t.Error("This specification should throw an error")
		}
	})
}

func TestGetDocuments(t *testing.T) {
	t.Run("A blank file does not generate documents", func(t *testing.T) {
		workdir := t.TempDir()
		file, err := os.CreateTemp(workdir, "*.empty.yml")

		if err != nil {
			t.Fatal(err)
		}

		defer file.Close()

		documents, err := GetDocuments([]string{file.Name()})

		if err != nil {
			t.Fatal(err)
		}

		if len(documents) != 0 {
			t.Error("We expected an empty list here")
		}
	})

	t.Run("files with multiple documents returns a populated list", func(t *testing.T) {
		workdir := t.TempDir()
		first, err := os.CreateTemp(workdir, "*.first.yml")

		if err != nil {
			t.Fatal(err)
		}

		first.WriteString(`
key: value
---
another: value
---
end: other
`,
		)

		second, err := os.CreateTemp(workdir, "*.second.yml")

		if err != nil {
			t.Fatal(err)
		}

		second.WriteString(`
new: app
---
key: another
---
xp: to
`,
		)

		files := []string{first.Name(), second.Name()}

		documents, err := GetDocuments(files)

		if err != nil {
			t.Fatal(err)
		}

		if len(documents) != 6 {
			t.Error("We expected an empty list here")
		}
	})
}

func TestGetYamlFiles(t *testing.T) {
	t.Run("An invalid working directory throws an error", func(t *testing.T) {
		viper.Set("workdir", "/an-invalid/directory")

		_, err := GetYamlFiles()

		if err == nil {
			t.Error("We expect a error here")
		}
	})

	t.Run("A directory without files generates an empty list", func(t *testing.T) {
		directory := t.TempDir()
		viper.Set("workdir", directory)

		result, err := GetYamlFiles()

		if err != nil {
			t.Fatal(err)
		}

		if len(result) != 0 {
			t.Logf("Recived result: %s", result)
			t.Error("We should generate an empty list here")
		}
	})

	t.Run("A directory without yaml's generates an empty list", func(t *testing.T) {
		directory := t.TempDir()
		viper.Set("workdir", directory)
		files := []string{"first.json", "second.json", "third.json"}

		for _, file := range files {
			file, err := os.CreateTemp(directory, fmt.Sprintf("*.%s", file))

			if err != nil {
				t.Fatal(err)
			}

			defer file.Close()
		}

		result, err := GetYamlFiles()

		if err != nil {
			t.Fatal(err)
		}

		if len(result) != 0 {
			t.Logf("Recived result: %s", result)
			t.Error("We should generate an empty list here")
		}
	})

	t.Run("The command returns all yaml's within the directory", func(t *testing.T) {
		directory := t.TempDir()
		viper.Set("workdir", directory)

		expected := []string{}

		nested, err := os.MkdirTemp(directory, "nested")

		if err != nil {
			t.Fatal(err)
		}

		files := []string{"first.yaml", "second.yml", "third.yml"}
		nestedFiles := []string{"first.yml", "second.yaml", "third.yaml"}

		for _, file := range files {
			file, err := os.CreateTemp(directory, fmt.Sprintf("*.%s", file))

			if err != nil {
				t.Fatal(err)
			}

			expected = append(expected, file.Name())
			defer file.Close()
		}

		for _, file := range nestedFiles {
			file, err := os.CreateTemp(nested, fmt.Sprintf("*.%s", file))

			if err != nil {
				t.Fatal(err)
			}

			expected = append(expected, file.Name())
			defer file.Close()
		}

		result, err := GetYamlFiles()

		if err != nil {
			t.Fatal(err)
		}

		sort.Strings(result)
		sort.Strings(expected)

		if !reflect.DeepEqual(result, expected) {
			t.Logf("Recived result: %s", result)
			t.Logf("Expected result: %s", expected)
			t.Error("Generated file list is not as expected")
		}
	})
}
