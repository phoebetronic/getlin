package clause

func (c *Clause) Automa() int {
	return len(c.neg) + len(c.pos)
}
