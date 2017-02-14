package gotype

import "errors"

var (
	errInvalidState             = errors.New("invalid state")
	errUnsupported              = errors.New("unsupported")
	errMapRequiresStringKey     = errors.New("map requires string key")
	errSquashNeedObject         = errors.New("require map or struct when using squash/inline")
	errNilInput                 = errors.New("nil input")
	errRequiresPointer          = errors.New("requires pointer")
	errKeyIntoNonStruct         = errors.New("key for non-structure target")
	errUnexpectedObjectKey      = errors.New("unexpected object key")
	errRequiresPrimitive        = errors.New("requires primitive target to set a boolean value")
	errRequiresBoolReceiver     = errors.New("requires bool receiver")
	errIncompatibleTypes        = errors.New("can not assign to incompatible go type")
	errStartArrayWaitingForKey  = errors.New("start array while waiting for object field name")
	errStartObjectWaitingForKey = errors.New("start object while waiting for object field name")
	errExpectedArrayNotObject   = errors.New("expected array but received object")
	errExpectedObjectNotArray   = errors.New("expected object but received array")
)

func errTODO() error {
	panic(errors.New("TODO"))
}

func visitErrTODO(V visitor, v interface{}) error {
	return errTODO()
}
