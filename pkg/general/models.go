package sheeters

type LLCel struct {
	key  *any
	next *LLCel
}

type LL struct {
	first *LLCel
	last  *LLCel
	count int
}

type G_sheet interface {
	id() int
	render()
	data() any
}
