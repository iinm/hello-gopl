実行方法:

```
go test -v -bench=. -benchmem
```


実行結果:

```
=== RUN   TestMakeTestArgs
--- PASS: TestMakeTestArgs (0.00s)
BenchmarkSlowEcho10-4            2000000               652 ns/op             389 B/op         10 allocs/op
BenchmarkSlowEcho100-4            200000              8682 ns/op           22162 B/op        100 allocs/op
BenchmarkSlowEcho1000-4             2000            554486 ns/op         2160639 B/op       1002 allocs/op
BenchmarkSlowEcho10000-4              30          52800198 ns/op        217046375 B/op     10004 allocs/op
BenchmarkEcho10-4                5000000               288 ns/op             199 B/op          3 allocs/op
BenchmarkEcho100-4               1000000              1420 ns/op            1686 B/op          3 allocs/op
BenchmarkEcho1000-4               100000             12004 ns/op           18695 B/op          3 allocs/op
BenchmarkEcho10000-4               10000            116476 ns/op          213085 B/op          3 allocs/op
PASS
ok      _/hello-gopl/ch01/ex03      12.519s
```

- `SlowEcho<N>` が非効率なバージョン、 `Echo<N>` が `strings.Join` を使ったバージョンの結果。`<N>` は引数の数
- `Echo` は入力を10倍にしたら実行時間もだいたい10倍程度になるが、 `SlowEcho` の実行時間の増加率はもっと大きい。1000から10000に増やしたときは実行時間が100倍程となっている (554486 ns/op → 52800198 ns/op)。
