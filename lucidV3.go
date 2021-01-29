package lucid

import (
	"math/rand"
	"strconv"
	"sync"
	"time"
)

//GeneratorV3 id generator
type GeneratorV3 struct {
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

func NewGeneratorV3(machineID int64) *GeneratorV3 {
	if machineID < 0 || machineID > 9 {
		panic("machineID is wrong")
	}
	length := int64(1e7) - 1
	sliceCh := make(chan []int64, 5)
	go func() {
		for {
			sliceCh <- ShuffledSlice(int(length))
		}
	}()
	return &GeneratorV3{
		machineID: machineID,
		maxNum:    length,
		nowFunc: func() time.Time {
			return time.Now().Local()
		},
		randNumSlice: ShuffledSlice(int(length)),
		sliceCh:      sliceCh,
	}
}

func (g *GeneratorV3) ID() int64 {
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

func (g *GeneratorV3) getLeftNum() int64 {
	return g.dateNumber() + g.secondNumber() + g.machineIDNumber()
}

func (g *GeneratorV3) dateNumber() int64 {
	dateStr := g.nowFunc().Format("060102")
	n, _ := strconv.ParseInt(dateStr, 10, 64)
	return n * 1e13
}

func (g *GeneratorV3) secondNumber() int64 {
	h, m, s := g.nowFunc().Clock()
	return (int64(h)*3600 + int64(m)*60 + int64(s)) * 1e8
}

func (g *GeneratorV3) machineIDNumber() int64 {
	return g.machineID * (1e7)
}

func ShuffledSlice(length int) []int64 {
	slice1 := make([]int64, 0, length)
	slice2 := make([]int64, length)
	var i int64
	for i = 0; i < int64(length); i++ {
		slice1 = append(slice1, i+1)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i, j := range r.Perm(length) {
		slice2[i] = slice1[j]
	}
	return slice2
}
