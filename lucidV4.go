package lucid

import (
	"strconv"
	"sync"
	"time"
)

//GeneratorV4 id generator
type GeneratorV4 struct {
	machineID int64
	nowFunc   func() time.Time
	now       int64
	leftNum   int64
	maxNum    int64
	sync.Mutex
	randNumSlice []int64
	sliceIndex   int
	sliceCh      chan []int64
}

func NewGeneratorV4(machineID int64) *GeneratorV4 {
	if machineID < 0 || machineID > 9 {
		panic("machineID is wrong")
	}
	length := int64(1e6) - 1
	sliceCh := make(chan []int64, 5)
	go func() {
		for {
			sliceCh <- ShuffledSlice(int(length))
		}
	}()
	return &GeneratorV4{
		machineID: machineID,
		maxNum:    length,
		nowFunc: func() time.Time {
			return time.Now().Local()
		},
		randNumSlice: ShuffledSlice(int(length)),
		sliceCh:      sliceCh,
	}
}

func (g *GeneratorV4) ID() int64 {
	now := g.nowFunc().Unix()
	g.Lock()
	defer g.Unlock()
	if g.sliceIndex >= int(g.maxNum-1) {
		for ; g.now == now; now = g.nowFunc().Unix() {
			g.now = now
			time.Sleep(time.Duration(10) * time.Millisecond)
		}
		g.randNumSlice = <-g.sliceCh
		g.sliceIndex = 0
	}
	if now == g.now {
		g.sliceIndex++
		return g.leftNum + g.randNumSlice[g.sliceIndex]
	}
	g.randNumSlice = <-g.sliceCh
	g.sliceIndex = 0
	g.now = now
	g.leftNum = g.getLeftNum()
	return g.leftNum + g.randNumSlice[0]
}

func (g *GeneratorV4) getLeftNum() int64 {
	return g.datetimeNumber() + g.machineIDNumber()
}

func (g *GeneratorV4) datetimeNumber() int64 {
	dateStr := g.nowFunc().Format("060102150405")
	n, _ := strconv.ParseInt(dateStr, 10, 64)
	return n * 1e7
}

func (g *GeneratorV4) machineIDNumber() int64 {
	return g.machineID * (1e6)
}
