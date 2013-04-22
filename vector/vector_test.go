package vector

import "testing"

func BenchmarkAdd(b *testing.B) {
	v1 := Vec{1, 2, 3}
	v2 := Vec{4, 5, 6}

	for i := 0; i < b.N; i++ {
		_ = v1.Add(&v2)
	}
}

func BenchmarkAddInPlace(b *testing.B) {
	v1 := Vec{1, 2, 3}
	v2 := Vec{4, 5, 6}

	for i := 0; i < b.N; i++ {
		v1.AddInPlace(&v2)
	}
}

func BenchmarkDist(b *testing.B) {
	v1 := Vec{1, 2, 3}
	v2 := Vec{4, 5, 6}

	for i := 0; i < b.N; i++ {
		_ = v1.Dist(&v2)
	}
}

func BenchmarkDistSq(b *testing.B) {
	v1 := Vec{1, 2, 3}
	v2 := Vec{4, 5, 6}

	for i := 0; i < b.N; i++ {
		_ = v1.DistSq(&v2)
	}
}

func BenchmarkDot(b *testing.B) {
	v1 := Vec{1, 2, 3}
	v2 := Vec{4, 5, 6}

	for i := 0; i < b.N; i++ {
		_ = v1.Dot(&v2)
	}
}

func BenchmarkLen(b *testing.B) {
	v := Vec{1, 2, 3}

	for i := 0; i < b.N; i++ {
		_ = v.Len()
	}
}

func BenchmarkLenSq(b *testing.B) {
	v := Vec{1, 2, 3}

	for i := 0; i < b.N; i++ {
		_ = v.LenSq()
	}
}

func BenchmarkNorm(b *testing.B) {
	v := Vec{1, 2, 3}

	for i := 0; i < b.N; i++ {
		_ = v.Norm()
	}
}

func BenchmarkNormInPlace(b *testing.B) {
	v := Vec{1, 2, 3}

	for i := 0; i < b.N; i++ {
		v.NormInPlace()
	}
}

func BenchmarkSub(b *testing.B) {
	v1 := Vec{1, 2, 3}
	v2 := Vec{4, 5, 6}

	for i := 0; i < b.N; i++ {
		_ = v1.Sub(&v2)
	}
}

func BenchmarkSubInPlace(b *testing.B) {
	v1 := Vec{1, 2, 3}
	v2 := Vec{4, 5, 6}

	for i := 0; i < b.N; i++ {
		v1.SubInPlace(&v2)
	}
}
