package d

type Row struct {
	values []Value
	schema Schema
}

func NewRow(schema Schema, values []Value) Row {
	err := schema.Validate(values)
	if err != nil {
		panic(struct {
			SchemaEntries []SchemaEntry
			Values        []Value
			Error         error
		}{
			schema.Entries(), values, err,
		})
	}
	return Row{
		schema: schema,
		values: values,
	}
}

func (r Row) Schema() Schema {
	return r.schema
}

func (r Row) Keys() S {
	return r.schema.Keys()
}

func (r Row) Length() int {
	return r.schema.Length()
}

func (r Row) Value(key string) Value {
	index := r.schema.Index(key)
	if index == NOT_FOUND {
		panic("key not found: " + key)
	}

	return r.values[index]
}

func (r Row) SetValue(key string, value Value) Row {
	index := r.schema.Index(key)
	if index == NOT_FOUND {
		panic("key not found: " + key)
	}

	if r.schema.Type(key) != value.Type {
		panic("value type mismatch")
	}

	r.values[index] = value

	return r
}
