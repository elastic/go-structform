package gotype

import "errors"

var (
	errUnsupported          = errors.New("unsupported")
	errTODO                 = errors.New("TODO")
	errMapRequiresStringKey = errors.New("map requires string key")
	errSquashNeedObject     = errors.New("require map or struct when using squash/inline")
)

func visitErrTODO(V visitor, v interface{}) error {
	return errTODO
}
