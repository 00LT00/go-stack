# stack的两种实现方式对比

### slice实现

```
$go-stack\go-stack-slice>go test -bench=".*" -benchmem -v
goos: windows
goarch: amd64
pkg: github.com/00LT00/go-stack/go-stack-slice
Benchmark_Push
Benchmark_Push-16       16216171                71.2 ns/op            94 B/op          0 allocs/op
Benchmark_Pop
Benchmark_Pop-16        57144217                21.5 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/00LT00/go-stack/go-stack-slice       6.351s

```

### nodelist实现

```
$go-stack\go-nodelist-stack>go test -bench=".*" -benchmem -v
goos: windows
goarch: amd64
pkg: github.com/00LT00/go-stack/go-nodelist-stack
Benchmark_Push
Benchmark_Push-16       14119174                71.0 ns/op            32 B/op          1 allocs/op
Benchmark_Pop
Benchmark_Pop-16        52737284                23.1 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/00LT00/go-stack/go-nodelist-stack    9.343s

```

|          |    push    |    pop     |
| :------: | :--------: | :--------: |
|  slice   | 71.2 ns/op | 21.5 ns/op |
| nodelist | 71.0 ns/op | 23.1 ns/op |