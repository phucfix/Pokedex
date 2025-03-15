package main

import "testing"

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input string
        expected []string
    }{
        {
            input:    "   hello  world   ",
            expected: []string{"hello", "world"},
        },
        {
            input:    "   ",
            expected: []string{},
        },
        {
            input:    "  hello  ",
            expected: []string{"hello"},
        },
        {
            input:    " HellO WoRLD  ",
            expected: []string{"hello", "world"},
        },
    }

    for _, c := range cases {
        actual := cleanInput(c.input)

        actualLength := len(actual)
        expectedLength := len(c.expected)
        if actualLength != expectedLength {
            t.Errorf("Length not match: Actual(%d) - Expected(%d)", actualLength, expectedLength)
            continue
        }

        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("Word not match at %dth word: Actual(%s) - Expected(%s)",
                        i, word, expectedWord)
            }
        }
    }
}
