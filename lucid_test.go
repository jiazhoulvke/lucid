package lucid

import (
	"testing"
)

func TestIntuitiveID(t *testing.T) {
	g := NewGenerator(1)
	m := make(map[int64]uint8)
	for i := 0; i < 1e8; i++ {
		m[g.ID()] = 1
	}
	if len(m) != 1e8 {
		t.Fail()
	}
}

func BenchmarkIntuitiveID(b *testing.B) {
	g := NewGenerator(1)
	for i := 0; i < b.N; i++ {
		g.ID()
	}
}
