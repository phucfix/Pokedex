package main

import (
    "testing"
    "fmt"
)

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

func TestParseCommadLine(t *testing.T) {
    cases := []struct {
        input             string
        expectedCommand   string
        expectedParameter string
    } {
        {
            input:             "",
            expectedCommand:   "",
            expectedParameter: "",
        },
        {
            input:             "  hello",
            expectedCommand:   "hello",
            expectedParameter: "",
        },
        {
            input:             " hello world   ",
            expectedCommand:   "hello",
            expectedParameter: "world",
        },
    }

    for i, c := range cases {
        t.Run(fmt.Sprintf("Test case %d", i), func(t *testing.T) {
            actualCommand, actualParameter := parseCommandLine(c.input)
            if actualCommand != c.expectedCommand {
                t.Errorf("Command not match: Actual(%s) - Expected(%s)",
                         actualCommand, c.expectedCommand)
                return
            }
            if actualParameter != c.expectedParameter {
                t.Errorf("Parameter not match: Actual(%s) - Expected(%s)",
                         actualParameter, c.expectedParameter)
                return
            }
        })
    }
}
