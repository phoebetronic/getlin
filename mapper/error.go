package mapper

import (
	"errors"

	"github.com/xh3b4sd/tracer"
)

var invalidMappingError = &tracer.Error{
	Kind: "invalidMappingError",
}

func IsInvalidMapping(err error) bool {
	return errors.Is(err, invalidMappingError)
}
