# ART Node16 SIMD Search Benchmarking

This repo is an experiment to try to get SSE2 accelerated search of 16 byte
array as required for the Node16 of an [Adaptive Radix
Tree](http://www-db.in.tum.de/~leis/papers/ART.pdf) or ART.

C code is based on github.com/armon/libart which is essentially the same as in
the paper.

Uses github.com/minio/c2goasm to generate Go ASM for the SIMD instructions.

## Sample results on 2013 MBP (i7-4850HQ)

```
BenchmarkIndexOfByteIn16Bytes_SIMD_5-8             	300000000	         5.32 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_SortBS_5-8           	200000000	         9.52 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_UnrolledBS_5-8       	300000000	         5.00 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_Brute_5-8            	500000000	         3.38 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_BruteUnrolled_5-8    	500000000	         3.29 ns/op	       0 B/op	       0 allocs/op

BenchmarkIndexOfByteIn16Bytes_SIMD_16-8            	200000000	         6.58 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_SortBS_16-8          	100000000	        16.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_UnrolledBS_16-8      	300000000	         5.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_Brute_16-8           	200000000	         8.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_BruteUnrolled_16-8   	200000000	         6.79 ns/op	       0 B/op	       0 allocs/op
```
## Sample results on 2017 MBP (i7-7700HQ)

```
BenchmarkIndexOfByteIn16Bytes_SIMD_5-8             	300000000	         5.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_SortBS_5-8           	200000000	         9.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_UnrolledBS_5-8       	300000000	         4.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_Brute_5-8            	500000000	         3.53 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_BruteUnrolled_5-8    	500000000	         3.13 ns/op	       0 B/op	       0 allocs/op

BenchmarkIndexOfByteIn16Bytes_SIMD_16-8            	300000000	         5.69 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_SortBS_16-8          	100000000	        14.2 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_UnrolledBS_16-8      	300000000	         4.50 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_Brute_16-8           	200000000	         8.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_BruteUnrolled_16-8   	200000000	         6.26 ns/op	       0 B/op	       0 allocs/op
```

But **with `go ... -gcflags="-l=4"`** which since go 1.10 allows inlining some intermediate functions we get:

```
BenchmarkIndexOfByteIn16Bytes_SIMD_5-8             	500000000	         3.28 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_SortBS_5-8           	200000000	         8.98 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_UnrolledBS_5-8       	300000000	         4.45 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_Brute_5-8            	500000000	         3.16 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_BruteUnrolled_5-8    	500000000	         3.12 ns/op	       0 B/op	       0 allocs/op

BenchmarkIndexOfByteIn16Bytes_SIMD_16-8            	300000000	         4.11 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_SortBS_16-8          	100000000	        14.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_UnrolledBS_16-8      	300000000	         4.46 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_Brute_16-8           	200000000	         7.74 ns/op	       0 B/op	       0 allocs/op
BenchmarkIndexOfByteIn16Bytes_BruteUnrolled_16-8   	200000000	         6.23 ns/op	       0 B/op	       0 allocs/op
```

## Conclusion (so far)

On a 2013 Macbook Pro. It turns out that SIMD is slower than some other methods, 
**unless `-gcflags="-l=4"` is used to compile the Go program**.
It may be I didn't find the optimal Assembly code and it could be some other
effects too though.

It also seems that brute force is much faster than `sort.Search` binary search 
in all cases, and that Go (as of 1.11.5) is not smart enough to unroll the 
simplest brute-force for loop since manually unrolling it into 16 `if` 
statements is still always quicker.

The fastest for 5 children (the smallest a Node16 in an ART should be) is 
unrolled brute force. For all 16 nodes full, unrolled binary search is quicker
(only ever requires 5 `if` branches).

I've not looked at the ASM from these yet to determine if there is a way to get
a totally branchless version the [equivalent of C ternary operator on some platforms](https://blog.demofox.org/2017/06/20/simd-gpu-friendly-branchless-binary-search/).

## Recommendation

For go-lang ART, the extra complexity of SSE instructions seems to not be 
worth it for Node16 search and a manually unrolled binary search is likely 
quicker unless non-standard go compiler optimizations are also enabled, 
even then it's borderline.

## Compiling

The Go code should compile with no extra deps since the ASM is committed.

BUilding the ASM is only tested/compiled on macOS 10.14.1 using default build 
tools `clang`. Compiler flags used can be seen in the makefile. I have no idea 
how portable the generated Go ASM code is in practice although my understanding 
is that it should work for any `amd64`/`x86_64` CPU that supports SSE2.

Run `make install-deps` to download some of the deps for rebuilding the ASM. 
You'll also need YASM (`brew instal yasm`). Then `make` will rebuild the Go ASM 
or `make bench` will rebuild and run benchmarks.

In practice a cpuid runtime check would be needed to select this 
implementation or a pure-go one, something like 
[the one in internal/cpu](https://golang.org/src/internal/cpu/cpu_x86.go#L74).
According to Go's internal/cpu package, SSE2 support is a requirement for the 
amd64 arch so that may even be unnecessary?
