package lucid

import (
	"testing"
)

func TestLucID(t *testing.T) {
	g := NewGenerator(1)
	m := make(map[int64]uint8)
	for i := 0; i < 1e6; i++ {
		m[g.ID()] = 1
	}
	if len(m) != 1e6 {
		t.Fail()
	}
}

func TestLucIDV2(t *testing.T) {
	g := NewGeneratorV2(1)
	m := make(map[int64]uint8)
	for i := 0; i < 1e6; i++ {
		m[g.ID()] = 1
	}
	if len(m) != 1e6 {
		t.Fail()
	}
}

func TestLucIDV3(t *testing.T) {
	g := NewGeneratorV3(1)
	m := make(map[int64]uint8)
	for i := 0; i < 1e6; i++ {
		m[g.ID()] = 1
	}
	if len(m) != 1e6 {
		t.Fail()
	}
}

func TestLucIDV4(t *testing.T) {
	g := NewGeneratorV4(1)
	m := make(map[int64]uint8)
	for i := 0; i < 1e6; i++ {
		m[g.ID()] = 1
	}
	if len(m) != 1e6 {
		t.Fail()
	}
}

func BenchmarkLucID(b *testing.B) {
	g := NewGenerator(1)
	for i := 0; i < b.N; i++ {
		g.ID()
	}
}

func BenchmarkLucIDV2(b *testing.B) {
	g := NewGeneratorV2(1)
	for i := 0; i < b.N; i++ {
		g.ID()
	}
}

func BenchmarkLucIDV3(b *testing.B) {
	g := NewGeneratorV3(1)
	for i := 0; i < b.N; i++ {
		g.ID()
	}
}

func BenchmarkLucIDV4(b *testing.B) {
	g := NewGeneratorV4(1)
	for i := 0; i < b.N; i++ {
		g.ID()
	}
}
