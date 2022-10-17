package files

import "testing"

func TestSpecification(t *testing.T) {
	t.Run("An invalid kind generates an error", func(t *testing.T) {
		spec := new(Specification)
		spec.Kind = "Invalid"

		if err := spec.Validate(); err.Error() != "The kind 'Invalid' is invalid" {
			t.Error("We expected a specific error here")
		}
	})

	t.Run("An invalid ApiVersion generates an error", func(t *testing.T) {
		spec := new(Specification)
		spec.Kind = "Service"
		spec.ApiVersion = "invalid/v1"

		if err := spec.Validate(); err.Error() != "The ApiVersion 'invalid/v1' is invalid" {
			t.Error("We expected a specific error here")
		}
	})

	t.Run("An unnamed specification throws an error", func(t *testing.T) {
		spec := new(Specification)
		spec.Kind = "Service"
		spec.ApiVersion = "webhook/v1"

		if err := spec.Validate(); err.Error() != "Every specification needs a name" {
			t.Error("We expected a specific error here")
		}
	})

	t.Run("Every specification without definition generates an error", func(t *testing.T) {
		spec := new(Specification)
		spec.Kind = "Service"
		spec.ApiVersion = "webhook/v1"
		spec.Metadata.Name = "Test example service"

		if err := spec.Validate(); err.Error() != "Every specification that needs at least one definition" {
			t.Error("We expected a specific error here")
		}
	})
}
