package lucid

import (
	"strconv"
	"sync"
	"time"
)

//GeneratorV2 id generator
type GeneratorV2 struct {
	machineID int64
	nowFunc   func() time.Time
	now       int64
	leftNum   int64
	n         int64
	maxNum    int64
	sync.Mutex
}

func NewGeneratorV2(machineID int64) *GeneratorV2 {
	if machineID < 0 || machineID > 9 {
		panic("machineID is wrong")
	}
	return &GeneratorV2{
		machineID: machineID,
		maxNum:    1e6 - 1,
		nowFunc: func() time.Time {
			return time.Now().Local()
		},
	}
}

func (g *GeneratorV2) ID() int64 {
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

func (g *GeneratorV2) getLeftNum() int64 {
	return g.datetimeNumber() + g.machineIDNumber()
}

func (g *GeneratorV2) datetimeNumber() int64 {
	dateStr := g.nowFunc().Format("060102150405")
	n, _ := strconv.ParseInt(dateStr, 10, 64)
	return n * 1e7
}

func (g *GeneratorV2) machineIDNumber() int64 {
	return g.machineID * (1e6)
}
