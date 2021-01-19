受twitter的Snowflake启发的分布式ID生成器。相比原版生成的ID更直观。

Example:

```go
package main

import (
	"fmt"

	"github.com/jiazhoulvke/lucid"
)

func main() {
	g := lucid.NewGenerator(1)
	fmt.Println(g.ID())
	//2101201335310000001
	// 年（2位）,月（2位）,日（2位）,距今日0点秒数（5位）,机器ID（1位），累加数（7位）
}
```
Benchmark:
```
$ go test -bench .
goos: linux
goarch: amd64
pkg: github.com/jiazhoulvke/lucid
BenchmarkLucidID-2          12711658                91.2 ns/op
PASS
ok      github.com/jiazhoulvke/lucid    39.734s
```
