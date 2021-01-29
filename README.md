受twitter的Snowflake启发的分布式ID生成器。相比原版生成的ID更直观。

目前有4个版本(下面说的位都是指10进制的):

- 第一种: 年(2位),月(2位),日(2位),距今日0点秒数(5位),机器ID(1位),累加数(7位)

- 第二种: 年(2位),月(2位),日(2位),时(2位),分(2位),秒(2位),机器ID(1位),累加数(6位)

- 第三种: 年(2位),月(2位),日(2位),距今日0点秒数(5位),机器ID(1位),随机数(7位)

- 第四种: 年(2位),月(2位),日(2位),时(2位),分(2位),秒(2位),机器ID(1位),随机数(6位)


示例:

```go
package main

import (
	"fmt"

	"github.com/jiazhoulvke/lucid"
)

func main() {
	g := lucid.NewGenerator(1)
	fmt.Println(g.ID())
	// 2101301018310000001
	g2 := lucid.NewGeneratorV2(1)
	fmt.Println(g2.ID())
	// 2101300249431000001
	g3 := lucid.NewGeneratorV3(1)
	fmt.Println(g3.ID())
	// 2101301018311104222
	g4 := lucid.NewGeneratorV4(1)
	fmt.Println(g4.ID())
	// 2101300249441290760
}
```

CPU: Intel(R) Core(TM) i7-7700HQ CPU @ 2.80GHz

Benchmark:

```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/jiazhoulvke/lucid
BenchmarkLucID-8        17433301                73.4 ns/op
BenchmarkLucIDV2-8       2463409               834 ns/op
BenchmarkLucIDV3-8             7         186879004 ns/op
BenchmarkLucIDV4-8       3398319               871 ns/op
PASS
ok      github.com/jiazhoulvke/lucid    16.806s
```

```
$ go test -bench . -benchtime=5s
goos: linux
goarch: amd64
pkg: github.com/jiazhoulvke/lucid
BenchmarkLucID-8        123520425               99.2 ns/op
BenchmarkLucIDV2-8      15943622               947 ns/op
BenchmarkLucIDV3-8      112661851              109 ns/op
BenchmarkLucIDV4-8      12714645               949 ns/op
PASS
ok      github.com/jiazhoulvke/lucid    86.138s
```
