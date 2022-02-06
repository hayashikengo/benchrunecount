# Run bench

## clone repo
```bash
$ git clone git@github.com:hayashikengo/benchrunecount.git
$ cd benchrunecount
```

## Run bench 
```bash
$ go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: benchrunecount
cpu: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
BenchmarkRunCountInString16ascii-8   	88374483	        13.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkRunCountInString16multi-8   	17920575	        62.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkRunCountInString32ascii-8   	44421601	        25.03 ns/op	       0 B/op	       0 allocs/op
BenchmarkRunCountInString32multi-8   	10113091	       116.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkCastToRuneArray16ascii-8    	100000000	        10.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkCastToRuneArray16multi-8    	19589365	        57.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkCastToRuneArray32ascii-8    	58416907	        20.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkCastToRuneArray32multi-8    	 9205329	       127.1 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	benchrunecount	9.988s
```

# 解説記事
https://zenn.dev/kenghaya/articles/1f6935419e0f21

# 参考記事
https://qiita.com/reiki4040/items/b82bf5056ee747dcf713
https://gist.github.com/reiki4040/6f225993b68805125e0f