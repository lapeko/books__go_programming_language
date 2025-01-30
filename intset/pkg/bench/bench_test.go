package bench

import (
	"math/rand"
	"testing"

	"github.com/lapeko/books__go_programming_language/intset/pkg/intset"
	binaryEngine "github.com/lapeko/books__go_programming_language/intset/pkg/intset/engine/binary-engine"
	mapEngine "github.com/lapeko/books__go_programming_language/intset/pkg/intset/engine/map-engine"
)

const lowValueLimit = 1000000

// Hight values
func BenchmarkMapAddHighValues(b *testing.B) {
	s := intset.New(mapEngine.New())
	for i := 0; i < b.N; i++ {
		s.Add(rand.Uint64())
	}
}

// TODO optimise. This tests fails. Seems an array is too big
func BenchmarkByteAddHighValues(b *testing.B) {
	s := intset.New(binaryEngine.New())
	for i := 0; i < b.N; i++ {
		s.Add(rand.Uint64())
	}
}

func BenchmarkMapHasHighValues(b *testing.B) {
	s := intset.New(mapEngine.New())
	randFill(s, b.N, false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Has(rand.Uint64())
	}
}

func BenchmarkByteHasHighValues(b *testing.B) {
	s := intset.New(binaryEngine.New())
	randFill(s, b.N, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Has(rand.Uint64())
	}
}

func BenchmarkMapDeleteHighValues(b *testing.B) {
	s := intset.New(mapEngine.New())
	randFill(s, b.N, false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Delete(rand.Uint64())
	}
}

func BenchmarkByteDeleteHighValues(b *testing.B) {
	s := intset.New(binaryEngine.New())
	randFill(s, b.N, false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Delete(rand.Uint64())
	}
}

func BenchmarkMapStringHighValues(b *testing.B) {
	s := intset.New(mapEngine.New())
	randFill(s, b.N, false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s.String()
	}
}

func BenchmarkByteStringHighValues(b *testing.B) {
	s := intset.New(binaryEngine.New())
	randFill(s, b.N, false)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s.String()
	}
}

// Low values
func BenchmarkMapAddLowValues(b *testing.B) {
	s := intset.New(mapEngine.New())
	for i := 0; i < b.N; i++ {
		s.Add(lowValueLimitRand())
	}
}

func BenchmarkByteAddLowValues(b *testing.B) {
	s := intset.New(binaryEngine.New())
	for i := 0; i < b.N; i++ {
		s.Add(lowValueLimitRand())
	}
}

func BenchmarkMapHasLowValues(b *testing.B) {
	s := intset.New(mapEngine.New())
	randFill(s, b.N, true)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Has(lowValueLimitRand())
	}
}

func BenchmarkByteHasLowValues(b *testing.B) {
	s := intset.New(binaryEngine.New())
	randFill(s, b.N, true)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Has(lowValueLimitRand())
	}
}

func BenchmarkMapDeleteLowValues(b *testing.B) {
	s := intset.New(mapEngine.New())
	randFill(s, b.N, true)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Delete(lowValueLimitRand())
	}
}

func BenchmarkByteDeleteLowValues(b *testing.B) {
	s := intset.New(binaryEngine.New())
	randFill(s, b.N, true)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s.Delete(lowValueLimitRand())
	}
}

func BenchmarkMapStringLowValues(b *testing.B) {
	s := intset.New(mapEngine.New())
	randFill(s, b.N, true)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s.String()
	}
}

func BenchmarkByteStringLowValues(b *testing.B) {
	s := intset.New(binaryEngine.New())
	randFill(s, b.N, true)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = s.String()
	}
}

func lowValueLimitRand() uint64 {
	return uint64(rand.Int63n(lowValueLimit))
}

func randFill(s intset.IntSet, sine int, low bool) {
	var rFunc func() uint64
	if low {
		rFunc = lowValueLimitRand
	} else {
		rFunc = rand.Uint64
	}
	for i := 0; i < sine; i++ {
		s.Add(rFunc())
	}
}
