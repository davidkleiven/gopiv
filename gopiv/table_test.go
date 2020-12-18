package gopiv

import "testing"

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