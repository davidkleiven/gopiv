package gopiv

// NumericColumn represents a column in a pivot table
type NumericColumn struct {
	Data []float64
	Name string
}

// TextColumn represents a column where the values are text
type TextColumn struct {
	Data []string
	Name string
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