package gopiv

import (
	"testing"
	"gonum.org/v1/gonum/floats"
	"sort"
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
		TextColumn: []TextColumn{
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
	for i, test := range []struct{
		Data []float64
		Unique []float64
	}{
		{
			Data: []float64{1.0, 2.0},
			Unique: []float64{1.0, 2.0},
		},
		{
			Data: []float64{1.0, 2.0, 3.0, 2.0},
			Unique: []float64{1.0, 2.0, 3.0},
		},
		{
			Data: []float64{1.0, 1.0, 1.0, 1.0001},
			Unique: []float64{1.0, 1.0001},
		},
		{
			Data: []float64{1.0, 1.0, -1.0, 1.0001},
			Unique: []float64{1.0, 1.0001, -1.0},
		},
	}{
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
	for i, test := range []struct{
		Data []string
		Unique []string
	}{
		{
			Data: []string{"one", "two"},
			Unique: []string{"one", "two"},
		},
		{
			Data: []string{"two", "one", "two", "two"},
			Unique: []string{"one", "two"},
		},
		{
			Data: []string{"one", "onee", "a", "b", "a"},
			Unique: []string{"one", "onee", "a", "b"},
		},
	}{
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