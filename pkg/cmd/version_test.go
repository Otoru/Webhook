package cmd

import (
	"bytes"
	"io/ioutil"
	"strings"
	"testing"

	"github.com/otoru/webhook/internal/version"
)

func TestCreateVersionCommand(t *testing.T) {
	t.Run("The command generates the expected output", func(t *testing.T) {
		buffer := bytes.NewBufferString("")
		cmd := CreateVersionCommand(buffer)

		cmd.Execute()

		result, err := ioutil.ReadAll(buffer)
		got := strings.TrimSpace(string(result))

		if err != nil {
			t.Fatal(err)
		}

		expected := version.Get(false)

		if got != expected {
			t.Errorf("Expected '%s' and got '%s'", expected, got)
		}
	})

	t.Run("The short flag shortens the command output", func(t *testing.T) {
		buffer := bytes.NewBufferString("")
		cmd := CreateVersionCommand(buffer)
		cmd.SetArgs([]string{"--short"})

		cmd.Execute()

		result, err := ioutil.ReadAll(buffer)
		got := strings.TrimSpace(string(result))

		if err != nil {
			t.Fatal(err)
		}

		expected := version.Get(true)

		if got != expected {
			t.Errorf("Expected '%s' and got '%s'", expected, got)
		}
	})
}
