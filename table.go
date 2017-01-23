package d

type Table struct {
	values [][]Value
	schema Schema
}

// Constructors
func NewTable(schema Schema, values [][]Value) Table {
	return Table{}
}

// Size and Schema
func (t Table) Rows() int {
	return 0
}

func (t Table) NumKeys() int {
	return 0
}

func (t Table) Schema() Schema {
	return t.schema
}

func (t Table) Keys() []string {
	return nil
}

func (t Table) TypeForKey(key string) ValueType {
	return 0
}

// Getters and Setters

// Individual Cell Getters
func (t Table) Value(index int, key string) Value {
	return Value{}
}

//Individual Cell Setters
func (t Table) SetValue(index int, key string, v Value) Table {
	return t
}

//Row Getters
func (t Table) Row(index int) Row {
	return Row{}
}

//Row Setters
func (t Table) AppendRows(row ...Row) Table {
	return t
}

func (t Table) ReplaceRow(index int, row Row) Table {
	return t
}

func (t Table) RemoveRow(index int) Table {
	return t
}

func (t Table) RemoveRows(indices []int) Table {
	return t
}

//Key Getters
func (t Table) FloatsForKey(key string, removeBlanks ...bool) D {
	return D{}
}

func (t Table) StringsForKey(key string, removeBlanks ...bool) S {
	return S{}
}

func (t Table) ValuesForKey(key string) []Value {
	return nil
}

//Key Setters
func (t Table) AddKey(key string, values []Value) Table {
	return t
}

func (t Table) AddFloatKey(key string, f D) Table {
	return t
}

func (t Table) AddStringKey(key string, s S) Table {
	return t
}

func (t Table) AddComputedKey(key string, cb func(row Row) Value) Table {
	return t
}

func (t Table) RemoveKey(key string) Table {
	return t
}

func (t Table) RemoveKeys(keys []string) Table {
	return t
}

//Table Array Methods
func (t Table) AppendTable(in Table) Table {
	return t
}

func (t Table) Slice(index int, count int) Table {
	return t
}

// Duplicate
func (t Table) Dup() Table {
	return t
}

// Iterators and Filters
func (t Table) Any(cb func(Row) bool) bool {
	return false
}

func (t Table) None(cb func(Row) bool) bool {
	return false
}

func (t Table) All(cb func(Row) bool) bool {
	return false
}

func (t Table) Filter(cb func(Row) bool) Table {
	return t
}

func (t Table) Reject(cb func(Row) bool) Table {
	return t
}

func (t Table) Each(cb func(Row) bool) Table {
	return t
}

func (t Table) Transform(schema Schema, cb func(Row) Row) Table {
	return t
}

func (t Table) MapToFloats(cb func(Row) float64) D {
	return D{}
}

func (t Table) ReduceToFloat(cb func(memo float64, r Row) float64, s float64) float64 {
	return 0
}

// bool operations
func (t Table) TruFa(cb func(r Row) bool) []bool {
	return nil
}

func (t Table) ApplyTruFa(trufa []bool) Table {
	return t
}

// Indexing
func (t Table) Index(indices []int) Table {
	return t
}

func (t Table) WithKeySubset(keys []string) Table {
	return t
}

// Order Helpers
func (t Table) Reverse() Table {
	return t
}

func (t Table) SortByKey(k ...string) Table {
	return t
}

func (t Table) Sort(less func(rowi, rowj Row) bool) Table {
	return t
}

func (t Table) SortedIndicesByKey(k ...string) Table {
	return t
}

func (t Table) SortedIndices(less func(rowi, rowj Row) bool) Table {
	return t
}
