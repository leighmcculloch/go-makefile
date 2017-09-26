package makefile

import (
	"reflect"
	"strings"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	testCases := []struct {
		File             string
		ExpectedMakefile Makefile
		ExpectedError    error
	}{
		{
			File:             ``,
			ExpectedMakefile: Makefile{},
			ExpectedError:    nil,
		},
		{
			File: `
target: dep
	command
`,
			ExpectedMakefile: Makefile{
				Targets: []Target{
					{Name: "target"},
				},
			},
			ExpectedError: nil,
		},
		{
			File: `
target1: dep1 dep2
	command1

target2: 
	command2
`,
			ExpectedMakefile: Makefile{
				Targets: []Target{
					{Name: "target1"},
					{Name: "target2"},
				},
			},
			ExpectedError: nil,
		},
		{
			File: `
export VAR = helloworld

target1: dep1 dep2
	command1

target2: 
	command2
`,
			ExpectedMakefile: Makefile{
				Targets: []Target{
					{Name: "target1"},
					{Name: "target2"},
				},
			},
			ExpectedError: nil,
		},
	}

	for _, tc := range testCases {
		m, err := Unmarshal(strings.NewReader(tc.File))
		switch {
		case !reflect.DeepEqual(err, tc.ExpectedError):
			t.Errorf("Unmarshal(%q) got error %v, want error %v", tc.File, err, tc.ExpectedError)
		case !reflect.DeepEqual(m, tc.ExpectedMakefile):
			t.Errorf("Unmarshal(%q) got makefile %v, want makefile %v", tc.File, m, tc.ExpectedMakefile)
		default:
			t.Logf("Unmarshal(%q) got makefile %v, error %v", tc.File, m, err)
		}
	}
}
