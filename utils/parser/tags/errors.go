package tags

import (
	"errors"
)

var (
	NilStructFieldsError   = errors.New("nil struct fields")
	NilGoStructFieldsError = errors.New("nil go struct fields")
	NilStructJSONTagError  = errors.New("nil struct json tag")
)
