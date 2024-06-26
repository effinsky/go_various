package ex

type Pair[A comparable] struct {
	Item  A
	Count int
}

// just more descriptive for the func
func NewPair[A comparable](it A) Pair[A] {
	return Pair[A]{it, 1}
}

func RLE_iter[A comparable](in []A) []Pair[A] {
	if len(in) == 0 {
		return []Pair[A]{}
	}

	out := make([]Pair[A], 0, len(in))
	currPair := NewPair(in[0])

	for _, it := range in[1:] {
		if it == currPair.Item {
			currPair.Count++
		} else {
			out = append(out, currPair)
			currPair = NewPair(it)
		}
	}

	out = append(out, currPair)
	return out
}

func RLErec[A comparable](in []A) []Pair[A] {
	if len(in) == 0 {
		return []Pair[A]{}
	}
	return rleAux(in, 0, make([]Pair[A], 0, len(in)))
}

func rleAux[A comparable](in []A, count int, acc []Pair[A]) []Pair[A] {
	if len(in) == 0 {
		return acc
	}
	if len(in) == 1 {
		it := in[0]
		acc = append(acc, Pair[A]{
			Item:  it,
			Count: count + 1,
		})
		return acc
	}
	// below len is > 1
	curr := in[0]
	next := in[1]
	// include next in the tail, we take it out only for comparison with curr
	tail := in[1:]

	if curr == next {
		return rleAux(tail, count+1, acc)
	} else {
		acc = append(acc, Pair[A]{
			Item:  curr,
			Count: count + 1,
		})
		return rleAux(tail, 0, acc)
	}
}
