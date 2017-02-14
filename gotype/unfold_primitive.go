package gotype

import "reflect"

var (
	invalidValue = reflect.Value{}
	trueValue    = reflect.ValueOf(true)
	falseValue   = reflect.ValueOf(false)
)

func (u *Unfolder) OnNil() error {
	return u.onValue(invalidValue)
}

func (u *Unfolder) OnBool(v bool) error {
	if v == true {
		return u.onValue(trueValue)
	}
	return u.onValue(falseValue)
}

func (u *Unfolder) OnString(v string) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnInt8(v int8) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnInt16(v int16) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnInt32(v int32) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnInt64(v int64) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnInt(v int) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnByte(v byte) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnUint8(v uint8) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnUint16(v uint16) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnUint32(v uint32) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnUint64(v uint64) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnUint(v uint) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnFloat32(v float32) error {
	return u.onValue(reflect.ValueOf(v))
}

func (u *Unfolder) OnFloat64(v float64) error {
	return u.onValue(reflect.ValueOf(v))
}
