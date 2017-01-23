package d

type ValueType int

const (
	NoneType ValueType = iota
	FloatType
	StringType
)

//What if this was a simple box?  with `value interface{}`?

type Value struct {
	S string
	F float64

	Type  ValueType
	Blank bool
}

func VF(f float64, blank ...bool) Value {
	if len(blank) > 0 {
		return Value{Type: FloatType, Blank: true}
	}
	return Value{
		F:    f,
		Type: FloatType,
	}
}

func VS(s string, blank ...bool) Value {
	if len(blank) > 0 {
		return Value{Type: StringType, Blank: true}
	}
	return Value{
		S:    s,
		Type: StringType,
	}
}
