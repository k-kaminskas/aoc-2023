package internal

import "sort"

/* Transformation --------------------------------------------------------------------------------------------------- */

type Transformation struct {
	destSt      int
	sourceSt    int
	rangeLength int
}

// TransformValue - Transform input value based on dest & source range start diffs
func (t *Transformation) TransformValue(value int) int {
	return value + t.destSt - t.sourceSt
}

// TransformRange - Transform Range value based on dest & source range start diffs
func (t *Transformation) TransformRange(rn *Range) *Range {
	op := t.destSt - t.sourceSt
	return &Range{rn.Start + op, rn.End + op}
}

/* Transformation Function ------------------------------------------------------------------------------------------ */

type TransformationFunction []*Transformation

// PerformTransformations - Performs set of transformations for a passed value
func (tr TransformationFunction) PerformTransformations(value int) int {
	for _, t := range tr {
		if t.sourceSt <= value && value <= t.sourceSt+t.rangeLength {
			val := t.TransformValue(value)
			return val
		}
	}
	return value
}

// PerformRangeTransformations - Performs set of transformations for a passed Range
func (tr TransformationFunction) PerformRangeTransformations(ranges ...*Range) []*Range {
	// No more transformations left, return the ranges
	if len(tr) == 0 {
		return ranges
	}
	ct := tr[0]

	// Transformed & remaining transformation ranges
	transformedRanges, remainingTrRanges := make([]*Range, 0), make([]*Range, 0)

	for _, rng := range ranges {
		// Produce a new range based on the current transformation
		trRange := &Range{ct.sourceSt, ct.sourceSt + ct.rangeLength}

		// Compare & get intersecting/remaining ranges
		intersectingRange, remainingRanges := rng.CompareRanges(trRange)
		if intersectingRange != nil {
			transformedRange := ct.TransformRange(intersectingRange)
			transformedRanges = append(transformedRanges, transformedRange)
		}
		remainingTrRanges = append(remainingTrRanges, remainingRanges...)
	}

	// Recursively process the remaining ranges with the next transformations
	return tr[1:].PerformRangeTransformations(
		append(transformedRanges, remainingTrRanges...)...,
	)
}

/* Function Set ----------------------------------------------------------------------------------------------------- */

type FunctionSet map[int]TransformationFunction

func (fs FunctionSet) GetSortedKeys() []int {
	keys := make([]int, 0, len(fs))
	for k := range fs {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

/* Range ------------------------------------------------------------------------------------------------------------ */

type Range struct {
	Start int
	End   int
}

// CompareRanges - Compares *comp_* range against current range & returns two
// new range sets - intersecting range & all other remaining ranges
func (r *Range) CompareRanges(comp_ *Range) (intersecting *Range, remaining []*Range) {
	intersectingStart := max(r.Start, comp_.Start)
	intersectingEnd := min(r.End, comp_.End)

	remaining = make([]*Range, 0)
	if intersectingStart <= intersectingEnd {
		intersecting = &Range{intersectingStart, intersectingEnd}
		if r.Start < intersectingStart {
			remaining = append(remaining, &Range{r.Start, intersectingStart})
		}
		if r.End > intersectingEnd {
			remaining = append(remaining, &Range{intersectingEnd, r.End})
		}
		return intersecting, remaining
	}

	return intersecting, append(remaining, r)
}
