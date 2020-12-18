package gopiv

import (
	"testing"
	"sort"
)

func TestIdentifyParameters(t *testing.T) {
	for i, test := range []struct{
		Expression string
		Params []string
		Candidates []string
	}{
		{
			Expression: "x*y^4",
			Params: []string{"x", "y"},
			Candidates: []string{"z", "x", "xx", "y", "yy"},
		},
		{
			Expression: "xx*y^4",
			Params: []string{"x", "y", "xx"},
			Candidates: []string{"z", "x", "xx", "y", "yy"},
		},
	}{
		sort.Strings(test.Params)
		found := identifyParameters(test.Candidates, test.Expression)
		sort.Strings(found)

		if !stringSliceEqual(test.Params, found) {
			t.Errorf("Test #%d: Expected\n%v\nGot\n%v\n", i, test.Params, found)
		}
	}
}