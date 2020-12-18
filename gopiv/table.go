package gopiv

// NumericColumn represents a column in a pivot table
type NumericColumn struct {
	Data []float64
	Name string
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

// TextColumn represents a column where the values are text
type TextColumn struct {
	Data []string
	Name string
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

// Table represents a flat pivot table
type Table struct {
	NumericColumns []NumericColumn
	TextColumn []TextColumn
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
	res := make([]string, len(t.TextColumn))
	for i, item := range t.TextColumn {
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