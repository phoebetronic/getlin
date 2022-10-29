package static

import "github.com/phoebetron/getlin"

// Module is a static module only used for testing. It enables us to inject
// statically configured vectors to be returned during inference.
type Module struct {
	out getlin.Vector
}

func New(con Config) *Module {
	{
		con.Verify()
	}

	return &Module{
		out: con.Out,
	}
}
