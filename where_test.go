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

// Quality assurance: Verify (and measure the performance) of the public API of the "generic" package.
package generic_test

import (
	"testing"

	"github.com/go-essentials/assert"
	"github.com/go-essentials/generic"
)

// UT: Return elements from a slice that matches a given predicate.
func TestWhere(t *testing.T) {
	t.Parallel() // Enable parallel execution.

	t.Run("When preallocating the slice that returns the matching elements.", func(t *testing.T) {
		// ARRANGE.
		data := generic.Slice[int]{1, 2, 3, 4, 5, 6}
		want := generic.Slice[int]{2, 4, 6}

		// ACT.
		got := data.Where(func(n int) bool {
			return n%2 == 0
		}, true)

		// ASSERT.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  (Pre-allocate) - Return elements from a slice that matches a given predicate.\n"+
			"\033[32mExpected: %v\033[0m\n"+
			"\033[31mActual:   %v\033[0m\n\n", want, got)
	})

	t.Run("When NOT preallocating the slice that returns the matching elements.", func(t *testing.T) {
		// ARRANGE.
		data := generic.Slice[int]{1, 2, 3, 4, 5, 6}
		want := generic.Slice[int]{2, 4, 6}

		// ACT.
		got := data.Where(func(n int) bool {
			return n%2 == 0
		}, false)

		// ASSERT.
		assert.EqualSf(t, got, want, "\n\n"+
			"UT Name:  (NO Pre-allocate) - Return elements from a slice that matches a given predicate.\n"+
			"\033[32mExpected: %v\033[0m\n"+
			"\033[31mActual:   %v\033[0m\n\n", want, got)
	})
}

// BENCHMARK: Measure the performance of the "Where" function.
func benchmarkWhere(b *testing.B, size int, prealloc bool) {
	// UTILITY FUNCTION: Populate some test data.
	populateData := func(amount int) generic.Slice[int] {
		data := make(generic.Slice[int], amount)

		for idx := range amount {
			data[idx] = idx
		}

		return data
	}

	// ARRANGE.
	data := populateData(size)
	b.ResetTimer()

	// RUN.
	for b.Loop() {
		data.Where(func(n int) bool {
			return n%2 == 0
		}, prealloc)
	}
}

// BENCHMARK COLLECTION: Measure the performance of the "Where" function.
func BenchmarkWhere_SmallWithPrealloc(b *testing.B)     { benchmarkWhere(b, 10, true) }
func BenchmarkWhere_MediumWithPrealloc(b *testing.B)    { benchmarkWhere(b, 10_000, true) }
func BenchmarkWhere_LargeWithPrealloc(b *testing.B)     { benchmarkWhere(b, 1_000_000, true) }
func BenchmarkWhere_SmallWithoutPrealloc(b *testing.B)  { benchmarkWhere(b, 10, false) }
func BenchmarkWhere_MediumWithoutPrealloc(b *testing.B) { benchmarkWhere(b, 10_000, false) }
func BenchmarkWhere_LargeWithoutPrealloc(b *testing.B)  { benchmarkWhere(b, 1_000_000, false) }
