# ART Node16 SIMD Search Benchmarking

This repo is an experiment to try to get SSE2 accelerated search of 16 byte
array as required for the Node16 of an [Adaptive Radix
Tree](http://www-db.in.tum.de/~leis/papers/ART.pdf) or ART.

C code is based on github.com/armon/libart which is essentially the same as in
the paper.

Uses github.com/minio/c2goasm to generate Go ASM for the SIMD instructions.

## Conclusion (so far)

On a 2013 Macbook Pro. It turns out that SIMD is slower than some other methods.
It may be I didn't find the optimal Assembly code and it could be some other
effects - the overhead of jumping to ASM which can't be inlined for example. I
did attempt to mitigate that by forcing `go:noinline` on the competing methods
with no change though.

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

