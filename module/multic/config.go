package multic

import (
	"github.com/phoebetron/getlin"
)

type Config struct {
	// Sin is the list of Single Modules used for training on the desired number
	// of classes. If there were 10 classes to be classified, the Multic Module
	// would have to be configured with 10 Single Modules in order to train each
	// Module on its own class. The provided Modules are expected to be
	// homogeneous in config and shape.
	Sin []getlin.Module
}

func (c Config) Verify() {
	if len(c.Sin) == 0 {
		panic("Config.Sin must not be empty")
	}
}
