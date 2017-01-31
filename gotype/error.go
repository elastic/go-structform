package gotype

import "errors"

var (
	errUnsupported          = errors.New("unsupported")
	errTODO                 = errors.New("TODO")
	errMapRequiresStringKey = errors.New("map requires string key")
)

func visitErrTODO(V visitor, v interface{}) error {
	return errTODO
}
