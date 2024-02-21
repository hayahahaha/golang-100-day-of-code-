package greatings

import (
	"regexp"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Say hello", func(t *testing.T) {
		name := "Gladys"
		want := regexp.MustCompile(`\b` + name + `\b`)
		msg, err := Hello(name)

		if !want.MatchString(msg) || err != nil {
			t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want)
		}
	})

	t.Run("Say hello empty", func(t *testing.T) {
		msg, err := Hello("")
		if err == nil {
			t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
		}

	})

}
