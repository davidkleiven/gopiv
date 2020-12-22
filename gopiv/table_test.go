package gopiv

import (
	"sort"
	"testing"

	"gonum.org/v1/gonum/floats"
)

func DefaultTestTable() Table {
	return Table{
		NumericColumns: []NumericColumn{
			{
				Name: "firstColumn",
				Data: []float64{1.0, -1.0, 2.0},
			},
			{
				Name: "secondColumn",
				Data: []float64{-1.0, -1.0, 20.0},
			},
		},
		TextColumns: []TextColumn{
			{
				Name: "firstTextColumn",
				Data: []string{"item", "menu", "item"},
			},
			{
				Name: "secondTextColumn",
				Data: []string{"item2", "menu5", "item1"},
			},
		},
	}
}

func WikipediaExample() Table {
	return Table{
		NumericColumns: []NumericColumn{
			{
				Name: "Units",
				Data: []float64{12.0, 12.0, 12.0, 10.0, 10.0, 10.0, 11.0, 11.0, 11.0, 15.0, 15.0},
			},
			{
				Name: "Price",
				Data: []float64{11.04, 13.0, 11.96, 11.27, 12.12, 13.74, 11.44, 12.63, 12.06, 13.48, 11.08},
			},
			{
				Name: "Cost",
				Data: []float64{10.42, 12.60, 11.74, 10.56, 11.95, 13.33, 10.94, 11.73, 11.51, 13.29, 10.67},
			},
		},
		TextColumns: []TextColumn{
			{
				Name: "Style",
				Data: []string{"Tee", "Golf", "Fancy", "Tee", "Golf", "Fancy", "Tee", "Golf", "Fancy", "Tee", "Golf"},
			},
			{
				Name: "Gender",
				Data: []string{"Boy", "Boy", "Boy", "Girl", "Girl", "Girl", "Girl", "Boy", "Boy", "Girl", "Girl"},
			},
			{
				Name: "Region",
				Data: []string{"East", "East", "East", "East", "East", "East", "West", "North", "West", "West", "South"},
			},
		},
	}
}

func TestNumericHeaders(t *testing.T) {
	tab := DefaultTestTable()
	num := tab.NumericHeaders()
	expect := []string{"firstColumn", "secondColumn"}

	if !stringSliceEqual(num, expect) {
		t.Errorf("Expected %v got %v\n", expect, num)
	}
}

func TestTextHeaders(t *testing.T) {
	tab := DefaultTestTable()
	txt := tab.TextHeaders()
	expect := []string{"firstTextColumn", "secondTextColumn"}

	if !stringSliceEqual(txt, expect) {
		t.Errorf("Expected %v got %v\n", expect, txt)
	}
}

func TestHeaders(t *testing.T) {
	tab := DefaultTestTable()
	headers := tab.Headers()
	expect := []string{"firstColumn", "secondColumn", "firstTextColumn", "secondTextColumn"}

	if !stringSliceEqual(headers, expect) {
		t.Errorf("Expected %v got %v\n", expect, headers)
	}
}

func TestNumericDistinct(t *testing.T) {
	for i, test := range []struct {
		Data   []float64
		Unique []float64
	}{
		{
			Data:   []float64{1.0, 2.0},
			Unique: []float64{1.0, 2.0},
		},
		{
			Data:   []float64{1.0, 2.0, 3.0, 2.0},
			Unique: []float64{1.0, 2.0, 3.0},
		},
		{
			Data:   []float64{1.0, 1.0, 1.0, 1.0001},
			Unique: []float64{1.0, 1.0001},
		},
		{
			Data:   []float64{1.0, 1.0, -1.0, 1.0001},
			Unique: []float64{1.0, 1.0001, -1.0},
		},
	} {
		col := NumericColumn{
			Data: test.Data,
		}
		sort.Float64s(test.Unique)
		distinct := col.Distinct()
		sort.Float64s(distinct)
		if !floats.EqualApprox(test.Unique, distinct, 1e-6) {
			t.Errorf("Test #%d: Expected\n%v\ngot\n%v\n", i, test.Unique, distinct)
		}
	}
}

func TestTextDistinct(t *testing.T) {
	for i, test := range []struct {
		Data   []string
		Unique []string
	}{
		{
			Data:   []string{"one", "two"},
			Unique: []string{"one", "two"},
		},
		{
			Data:   []string{"two", "one", "two", "two"},
			Unique: []string{"one", "two"},
		},
		{
			Data:   []string{"one", "onee", "a", "b", "a"},
			Unique: []string{"one", "onee", "a", "b"},
		},
	} {
		col := TextColumn{
			Data: test.Data,
		}
		sort.Strings(test.Unique)
		distinct := col.Distinct()
		sort.Strings(distinct)
		if !stringSliceEqual(test.Unique, distinct) {
			t.Errorf("Test #%d: Expected\n%v\ngot\n%v\n", i, test.Unique, distinct)
		}
	}
}

func TestExpressions(t *testing.T) {
	tab := WikipediaExample()
	if !tab.IsConsistent() {
		t.Errorf("Inconsistent number of columns in wikipedia example\n")
	}

	for i, test := range []struct {
		Expression string
		Expect     int
	}{
		{
			Expression: "Region=='East'",
			Expect:     6,
		},
		{
			Expression: "Region=='West'",
			Expect:     3,
		},
		{
			Expression: "Region=='North'",
			Expect:     1,
		},
		{
			Expression: "Region=='South'",
			Expect:     1,
		},
		{
			Expression: "Region=='South' || Region=='North'",
			Expect:     2,
		},
		{
			Expression: "Region!='South'",
			Expect:     10,
		},
		{
			Expression: "Region!='South' && Price>12.0",
			Expect:     6,
		},
		{
			Expression: "Region=='West' && Price>12.0 && Price<13.0",
			Expect:     1,
		},
		{
			Expression: "Units>12.0",
			Expect:     2,
		},
	} {
		filtered, err := tab.Filter(test.Expression)
		if err != nil {
			t.Errorf("Test #%d: %s\n", i, err)
		}

		if filtered.Len() != test.Expect {
			t.Errorf("Test #%d: Expected %d rows, got %d\n", i, test.Expect, filtered.Len())
		}
	}
}

func TestGetNumericRow(t *testing.T) {
	table := DefaultTestTable()
	row := table.GetNumericRow(1)
	expect := []float64{-1.0, -1.0}
	if !floats.EqualApprox(expect, row, 1e-6) {
		t.Errorf("Expected\n%v\ngot\n%v\n", expect, row)
	}
}

func TestGetTextRow(t *testing.T) {
	table := DefaultTestTable()
	row := table.GetTextRow(1)
	expect := []string{"menu", "menu5"}
	if !stringSliceEqual(expect, row) {
		t.Errorf("Expected\n%v\ngot\n%v\n", expect, row)
	}
}

func TestGetRow(t *testing.T) {
	table := DefaultTestTable()
	row := table.GetRow(1)
	expect := []string{"-1.000000", "-1.000000", "menu", "menu5"}
	if !stringSliceEqual(expect, row) {
		t.Errorf("Expected\n%v\ngot\n%v\n", expect, row)
	}
}
