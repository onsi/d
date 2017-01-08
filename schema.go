package d

import "errors"

const NOT_FOUND = -1

type SchemaEntry struct {
	Name string
	Type ValueType
}

type Schema struct {
	lookup  map[string]int
	entries []SchemaEntry
}

func NewSchema(entries []SchemaEntry) Schema {
	lookup := map[string]int{}
	for i, entry := range entries {
		lookup[entry.Name] = i
	}

	return Schema{
		lookup:  lookup,
		entries: entries,
	}
}

func (s Schema) Row(values []Value) Row {
	return NewRow(s, values)
}

func (s Schema) Table(values []Row) Table {
	return NewTable(s, values)
}

func (s Schema) Validate(values []Value) error {
	if len(values) != len(s.entries) {
		return errors.New("length mismatch")
	}

	for i, v := range values {
		if v.Type != s.entries[i].Type {
			return errors.New("type mismatch")
		}
	}
	return nil
}

func (s Schema) Entries() []SchemaEntry {
	dup := make([]SchemaEntry, len(s.entries))
	copy(dup, s.entries)
	return dup
}

func (s Schema) Keys() S {
	keys := S{}
	for _, entry := range s.entries {
		keys = append(keys, entry.Name)
	}

	return keys
}

func (s Schema) Length() int {
	return len(s.entries)
}

func (s Schema) Index(key string) int {
	index, exists := s.lookup[key]
	if !exists {
		return NOT_FOUND
	}
	return index
}

func (s Schema) Type(key string) ValueType {
	index := s.Index(key)
	if index == NOT_FOUND {
		return NoneType
	}

	return s.entries[index].Type
}
