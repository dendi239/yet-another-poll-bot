package grammar

import "testing"

func TestCorrectNegateTokenParse(t *testing.T) {
	correctSamples := []string{
		"!2",
		"!1 & !2",
		"1 | !2",
		"!(2 | !4)",
	}

	for _, s := range correctSamples {
		ts, err := Tokenize(s)
		if err != nil {
			t.Fail()
		}

		if _, rest, err := Parse(ts); err != nil || len(rest) > 0 {
			t.Fail()
		}
	}
}

func TestIncorrectNegateTokenParse(t *testing.T) {
	incorrectSamples := []string{
		"!",
		"2!",
		"(1|2!)",
		"!((2))!",
	}

	for _, s := range incorrectSamples {
		ts, err := Tokenize(s)
		if err != nil {
			t.Fail()
		}

		if term, rest, err := Parse(ts); err == nil && len(rest) == 0 {
			t.Error(s, term, rest)
		}
	}
}
