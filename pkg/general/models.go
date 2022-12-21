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
