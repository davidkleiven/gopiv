package gopiv

import (
	"github.com/Knetic/govaluate"
)

// NumericColumn represents a column in a pivot table
type NumericColumn struct {
	Data []float64
	Name string
}

// Append a value to the row
func (nc *NumericColumn) Append(v float64) {
	nc.Data = append(nc.Data, v)
}

// Len returns the length of the underlying data array
func (nc NumericColumn) Len() int {
	return len(nc.Data)
}

// Distinct returns the distinct values of column
func (nc NumericColumn) Distinct() []float64 {
	vals := make(map[float64]interface{})
	for _, v := range nc.Data {
		vals[v] = nil
	}

	res := []float64{}
	for k := range vals {
		res = append(res, k)
	}
	return res
}

// Clear deletes all data
func (nc *NumericColumn) Clear() {
	nc.Data = []float64{}
}

// TextColumn represents a column where the values are text
type TextColumn struct {
	Data []string
	Name string
}

// Clear deletes all data
func (tc *TextColumn) Clear() {
	tc.Data = []string{}
}

// Distinct returns all the distinct values of the column
func (tc TextColumn) Distinct() []string {
	vals := make(map[string]interface{})
	for _, v := range tc.Data {
		vals[v] = nil
	}

	res := []string{}
	for k := range vals {
		res = append(res, k)
	}
	return res
}

// Len returns the length of the underlying data array
func (tc TextColumn) Len() int {
	return len(tc.Data)
}

// Append a value to the row
func (tc *TextColumn) Append(v string ){
	tc.Data = append(tc.Data, v)
}

// Table represents a flat pivot table
type Table struct {
	NumericColumns []NumericColumn
	TextColumns []TextColumn
}

// Len returns the number of items in a table
func (t Table) Len() int {
	if len(t.NumericColumns) > 0 {
		return len(t.NumericColumns[0].Data)
	} else if len(t.TextColumns) > 0 {
		return len(t.TextColumns[0].Data)
	}
	return 0
}

// EmptyTable returns a new table with no fields
func EmptyTable() Table {
	return Table{
		NumericColumns: []NumericColumn{},
		TextColumns: []TextColumn{},
	}
}

// EmptyTableFromSchema returns an empty table, but the fields are initializes
func EmptyTableFromSchema(numHeaders []string, txtHeaders []string) Table {
	tab := EmptyTable()
	for _, v := range numHeaders {
		tab.NumericColumns = append(tab.NumericColumns, NumericColumn{
			Name: v,
			Data: []float64{},
		})
	}

	for _, v := range txtHeaders {
		tab.TextColumns = append(tab.TextColumns, TextColumn{
			Name: v,
			Data: []string{},
		})
	}
	return tab
}

// NumericHeaders returns the name of the header of all numeric columns
func (t Table) NumericHeaders() []string {
	res := make([]string, len(t.NumericColumns))
	for i, item := range t.NumericColumns {
		res[i] = item.Name
	}
	return res
}

// TextHeaders returns a list of all headers of text fields
func (t Table) TextHeaders() []string {
	res := make([]string, len(t.TextColumns))
	for i, item := range t.TextColumns {
		res[i] = item.Name
	}
	return res
}

// Headers returns all headers
func (t Table) Headers() []string {
	num := t.NumericHeaders()
	txt := t.TextHeaders()
	return append(num, txt...)
}


// Filter all fields satisfying the passed expression.
func (t Table) Filter(expression string) (Table, error) {
	parameters := make(map[string]interface{})

	filteredTable := EmptyTableFromSchema(t.NumericHeaders(), t.TextHeaders())
	expr, err := govaluate.NewEvaluableExpression(expression);
	if err != nil {
		return filteredTable, err
	}

	for i := 0;i<t.Len();i++ {
		// Extract the entries on this row and add it to the parameters map
		for _, item := range t.NumericColumns {
			parameters[item.Name] = item.Data[i]
		}
		for _, item := range t.TextColumns {
			parameters[item.Name] = item.Data[i]
		}

		// Evaluate the passed expression using the extracted parameters
		res, err := expr.Evaluate(parameters)
		if err != nil {
			return filteredTable, err
		}

		if res.(bool) {
			// Transfer all items
			for j := range t.NumericColumns {
				filteredTable.NumericColumns[j].Append(t.NumericColumns[j].Data[i])
			}
			for j := range t.TextColumns {
				filteredTable.TextColumns[j].Append(t.TextColumns[j].Data[i])
			}
		}
	}
	return filteredTable, nil
}

// IsConsistent returns true if all columns have the same number of records.
func (t Table) IsConsistent() bool {
	l := t.Len()
	for _, v := range t.NumericColumns {
		if v.Len() != l {
			return false
		}
	}

	for _, v := range t.TextColumns {
		if v.Len() != l {
			return false
		}
	}
	return true
}