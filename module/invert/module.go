package invert

// Module is an invert module only used for testing. It enables us to track
// vector modifications inside module layers during inference.
type Module struct{}

func New(con Config) *Module {
	{
		con.Verify()
	}

	return &Module{}
}
