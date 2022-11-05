package loader

type Loader struct {
	bas string
	bat int
	git bool
	org string
	rep string
}

func New(con Config) *Loader {
	{
		con = con.Ensure()
	}

	{
		con.Verify()
	}

	return &Loader{
		bas: con.Bas,
		bat: 64,
		git: con.Git,
		org: con.Org,
		rep: con.Rep,
	}
}
