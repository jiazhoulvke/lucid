package lucid

import (
	"strconv"
	"sync"
	"time"
)

//Generator id generator
type Generator struct {
	machineID int64
	nowFunc   func() time.Time
	now       int64
	leftNum   int64
	n         int64
	maxNum    int64
	sync.Mutex
}

func NewGenerator(machineID int64) *Generator {
	if machineID < 0 || machineID > 9 {
		panic("machineID is wrong")
	}
	return &Generator{
		machineID: machineID,
		maxNum:    1e7 - 1,
		nowFunc: func() time.Time {
			return time.Now().Local()
		},
	}
}

func (g *Generator) ID() int64 {
	now := g.nowFunc().Unix()
	g.Lock()
	defer g.Unlock()
	if g.n >= g.maxNum {
		for ; g.now == now; now = g.nowFunc().Unix() {
			g.now = now
			time.Sleep(time.Duration(10) * time.Millisecond)
		}
		g.n = 1
	}
	if now == g.now {
		g.n++
		return g.leftNum + g.n
	}
	g.now = now
	g.leftNum = g.getLeftNum()
	g.n = 1
	return g.leftNum + 1
}

func (g *Generator) getLeftNum() int64 {
	return g.dateNumber() + g.secondNumber() + g.machineIDNumber()
}

func (g *Generator) dateNumber() int64 {
	dateStr := g.nowFunc().Format("060102")
	n, _ := strconv.ParseInt(dateStr, 10, 64)
	return n * 1e13
}

func (g *Generator) secondNumber() int64 {
	h, m, s := g.nowFunc().Clock()
	return (int64(h)*3600 + int64(m)*60 + int64(s)) * 1e8
}

func (g *Generator) machineIDNumber() int64 {
	return g.machineID * (1e7)
}
