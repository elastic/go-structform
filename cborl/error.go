package cborl

import "errors"

func errTODO() error {
	err := errors.New("TODO")
	panic(err)
}

var errInvalidCode = errors.New("invalid type code")
var errTextKeyRequired = errors.New("only text keys supported")
var errIndefByteSeq = errors.New("text/bytes of indefinite length not supported")
