package inicfg

import "strconv"

type Value struct {
	data string
}

func (value *Value) AsString() string {
	return value.data
}

func (value *Value) AsInt() (int, error) {
	return strconv.Atoi(value.data)
}

func (value *Value) AsFloat32() (float64, error) {
	return strconv.ParseFloat(value.data, 32)
}

func (value *Value) AsFloat64() (float64, error) {
	return strconv.ParseFloat(value.data, 64)
}

func (value *Value) AsBool() (bool, error) {
	return strconv.ParseBool(value.data)
}
