package bastet_test

import (
	"strings"
	"testing"

	"github.com/andreyvit/diff"
	"github.com/gdey/bastet"
)

func TestProcess(t *testing.T) {
	type testCase struct {
		Templates []bastet.Template
		Values    map[string]string
		Expected  string
	}

	fn := func(tc testCase) func(t *testing.T) {
		return func(t *testing.T) {

			var w strings.Builder
			bastet.Process(&w, tc.Templates, tc.Values)
			got := w.String()
			if tc.Expected != got {
				t.Errorf("invalid output, expected:\n%v", diff.LineDiff(tc.Expected, got))
			}

		}
	}

	tests := map[string]testCase{
		"Example": {
			Values: map[string]string{
				"name":     "Gautam Dey",
				"greeting": "Hello",
				"date":     "25 of July",
			},
			Templates: []bastet.Template{
				{
					Name:   "file:header.tpl",
					Reader: strings.NewReader("{{.greeting}} {{.name}},\n\n"),
				},
				{
					Name: "file:body.tpl",
					Reader: strings.NewReader(`Thank you for comming to our event on {{.date}}, {{.name}}.
We hope you had fun, and if you have any questions please feel
free to reach out to us.

Thank you.`),
				},
			},
			Expected: `Hello Gautam Dey,

Thank you for comming to our event on 25 of July, Gautam Dey.
We hope you had fun, and if you have any questions please feel
free to reach out to us.

Thank you.`,
		},
	}

	for name, tc := range tests {
		t.Run(name, fn(tc))
	}
}
