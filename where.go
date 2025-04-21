// =====================================================================================================================
// == LICENSE:       Copyright (c) 2025 Kevin De Coninck
// ==
// ==                Permission is hereby granted, free of charge, to any person
// ==                obtaining a copy of this software and associated documentation
// ==                files (the "Software"), to deal in the Software without
// ==                restriction, including without limitation the rights to use,
// ==                copy, modify, merge, publish, distribute, sublicense, and/or sell
// ==                copies of the Software, and to permit persons to whom the
// ==                Software is furnished to do so, subject to the following
// ==                conditions:
// ==
// ==                The above copyright notice and this permission notice shall be
// ==                included in all copies or substantial portions of the Software.
// ==
// ==                THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// ==                EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
// ==                OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// ==                NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// ==                HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY,
// ==                WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
// ==                FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR
// ==                OTHER DEALINGS IN THE SOFTWARE.
// =====================================================================================================================

// Package generic provides a set of "generic" functions.
package generic

// Where returns a new slice containing all elements in data that satisfy the given predicate.
//
// If prealloc is true, the result slice is preallocated with the full capacity of data, which avoids additional
// allocations during filtering. This may use more memory than necessary if few elements match.
//
// If prealloc is false, the result slice starts with zero capacity and grows as needed, using only as much memory as
// required to store the matching elements.
func (data Slice[T]) Where(predicate func(T) bool, prealloc bool) Slice[T] {
	var matches Slice[T]

	if prealloc {
		matches = make(Slice[T], 0, len(data))
	} else {
		matches = make(Slice[T], 0)
	}

	for _, el := range data {
		if predicate(el) {
			matches = append(matches, el)
		}
	}

	return matches
}
