package parallel

import (
	"strings"
	"testing"
)

func Test1(t *testing.T) { runSuite(t) }
func Test2(t *testing.T) { runSuite(t) }

func runSuite(t *testing.T) {
	t.Parallel() // run top-level tests in parallel

	for _, test := range testCases {
		t.Run(test.input, test.execute)
	}
}

var testCases = []TestCase{
	{
		input:    "this passes",
		expected: "THIS PASSES",
	},
	{
		input:    "this will fail",
		expected: "because this is wrong",
	},
}

type TestCase struct {
	input    string
	expected string
}

func (test TestCase) execute(t *testing.T) {
	t.Parallel() // run each sub-test in parallel

	actual := strings.ToUpper(test.input)
	if actual != test.expected {
		t.Errorf("\ngot:  [%s]\nwant: [%s]", actual, test.expected)
	}
}
