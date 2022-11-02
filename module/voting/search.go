package voting

import (
	"github.com/phoebetron/getlin"
	"github.com/phoebetron/getlin/vector"
)

func (m *Module) Search(vec getlin.Vector) getlin.Vector {
	if m.wei == nil {
		return vector.New(vector.Config{Out: []bool{vec.Inp().Maj()}})
	}

	// The positive input Vector bits have to cross the threshold learned by the
	// weights Module. If it does, the vote is true. Reinforcement of true
	// patterns will keep the threshold high, while reinforcement of false
	// patterns will help find a balance. In case a pattern cannot reach the
	// required threshold, the vote is false.
	if vec.Inp().Wei(true) > m.wei.Search(vec).Out().Wei(true) {
		return vector.New(vector.Config{Out: []bool{true}})
	}

	return vector.New(vector.Config{Out: []bool{false}})
}
