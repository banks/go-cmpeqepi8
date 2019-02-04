package cmpeqepi8

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIndexOfByteIn16Bytes(t *testing.T) {
	tt := []struct {
		name  string
		bytes []byte
		c     byte
		want  int
	}{
		{
			name:  "full,0",
			bytes: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'l', 'm', 'n', 'o', 'p', 'q'},
			c:     'a',
			want:  0,
		},
		{
			name:  "full,1",
			bytes: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'l', 'm', 'n', 'o', 'p', 'q'},
			c:     'b',
			want:  1,
		},
		{
			name:  "full,15",
			bytes: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'l', 'm', 'n', 'o', 'p', 'q'},
			c:     'q',
			want:  15,
		},
		{
			name:  "full,-1",
			bytes: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'l', 'm', 'n', 'o', 'p', 'q'},
			c:     'z',
			want:  -1,
		},
		{
			name:  "partial,0",
			bytes: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'},
			c:     'a',
			want:  0,
		},
		{
			name:  "partial,1",
			bytes: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'},
			c:     'b',
			want:  1,
		},
		{
			name:  "partial,8",
			bytes: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'},
			c:     'h',
			want:  7,
		},
		{
			name:  "partial,-1",
			bytes: []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h'},
			c:     'i',
			want:  -1,
		},
		{
			name:  "empty,-1",
			bytes: []byte{},
			c:     'a',
			want:  -1,
		},
		{
			name:  "signedbits,0",
			bytes: []byte{0xFF},
			c:     0xFF,
			want:  0,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			var a [16]byte

			copy(a[:], tc.bytes)

			idx := IndexOfByteIn16Bytes(uint8(len(tc.bytes)), &a, tc.c)

			urIdx := unrolledIndexOfByteIn16Bytes(uint8(len(tc.bytes)), &a, tc.c)
			burIdx := bruteUnrolledIndexOfByteIn16Bytes(uint8(len(tc.bytes)), &a, tc.c)

			require.Equal(t, tc.want, idx)
			require.Equal(t, tc.want, urIdx)
			require.Equal(t, tc.want, burIdx)
		})
	}
}

var result int

//go:noinline
func unrolledIndexOfByteIn16Bytes(n uint8, arr *[16]byte, c byte) int {
	// Hand unrolled binary search
	idx := 0
	if n > 8 && arr[8] <= c {
		idx += 8
	}
	if int(n) > idx+4 && arr[idx+4] <= c {
		idx += 4
	}
	if int(n) > idx+2 && arr[idx+2] <= c {
		idx += 2
	}
	if int(n) > idx+1 && arr[idx+1] <= c {
		idx += 1
	}
	if arr[idx] != c {
		idx = -1
	}
	return idx
}

//go:noinline
func bruteIndexOfByteIn16Bytes(n uint8, arr *[16]byte, c byte) int {
	for i := 0; i < int(n); i++ {
		if arr[i] == c {
			return i
		}
	}
	return -1
}

//go:noinline
func bruteUnrolledIndexOfByteIn16Bytes(n uint8, arr *[16]byte, c byte) int {
	if n > 0 && arr[0] == c {
		return 0
	}
	if n > 1 && arr[1] == c {
		return 1
	}
	if n > 2 && arr[2] == c {
		return 2
	}
	if n > 3 && arr[3] == c {
		return 3
	}
	if n > 4 && arr[4] == c {
		return 4
	}
	if n > 5 && arr[5] == c {
		return 5
	}
	if n > 6 && arr[6] == c {
		return 6
	}
	if n > 7 && arr[7] == c {
		return 7
	}
	if n > 8 && arr[8] == c {
		return 8
	}
	if n > 9 && arr[9] == c {
		return 9
	}
	if n > 10 && arr[10] == c {
		return 10
	}
	if n > 11 && arr[11] == c {
		return 11
	}
	if n > 12 && arr[12] == c {
		return 12
	}
	if n > 13 && arr[13] == c {
		return 13
	}
	if n > 14 && arr[14] == c {
		return 14
	}
	if n > 15 && arr[15] == c {
		return 15
	}
	return -1
}

func BenchmarkIndexOfByteIn16Bytes_SIMD_5(b *testing.B) {
	a := [16]byte{'a', 'b', 'c', 'd', 'e'}
	l := uint8(5)
	c := byte('j')
	for i := 0; i < b.N; i++ {
		result = IndexOfByteIn16Bytes(l, &a, c)
	}
}

func BenchmarkIndexOfByteIn16Bytes_SortBS_5(b *testing.B) {
	a := [16]byte{'a', 'b', 'c', 'd', 'e'}
	l := 5
	//c := byte('j')
	for i := 0; i < b.N; i++ {
		result = sort.Search(l, func(i int) bool {
			return a[i] >= 'j'
		})
	}
}
func BenchmarkIndexOfByteIn16Bytes_UnrolledBS_5(b *testing.B) {
	a := [16]byte{'a', 'b', 'c', 'd', 'e'}
	l := uint8(5)
	c := byte('b')
	for i := 0; i < b.N; i++ {
		result = unrolledIndexOfByteIn16Bytes(l, &a, c)
	}
}

func BenchmarkIndexOfByteIn16Bytes_Brute_5(b *testing.B) {
	a := [16]byte{'a', 'b', 'c', 'd', 'e'}
	l := uint8(5)
	c := byte('b')
	for i := 0; i < b.N; i++ {
		result = bruteIndexOfByteIn16Bytes(l, &a, c)
	}
}

func BenchmarkIndexOfByteIn16Bytes_BruteUnrolled_5(b *testing.B) {
	a := [16]byte{'a', 'b', 'c', 'd', 'e'}
	l := uint8(5)
	c := byte('e')
	for i := 0; i < b.N; i++ {
		result = bruteUnrolledIndexOfByteIn16Bytes(l, &a, c)
	}
}

func BenchmarkIndexOfByteIn16Bytes_SIMD_16(b *testing.B) {
	a := [16]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'l', 'm', 'n', 'o', 'p', 'q'}
	l := uint8(16)
	c := byte('j')
	for i := 0; i < b.N; i++ {
		result = IndexOfByteIn16Bytes(l, &a, c)
	}
}

func BenchmarkIndexOfByteIn16Bytes_SortBS_16(b *testing.B) {
	a := [16]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'l', 'm', 'n', 'o', 'p', 'q'}
	l := 16
	//c := byte('j')
	for i := 0; i < b.N; i++ {
		result = sort.Search(l, func(i int) bool {
			return a[i] >= 'j'
		})
	}
}

func BenchmarkIndexOfByteIn16Bytes_UnrolledBS_16(b *testing.B) {
	a := [16]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'l', 'm', 'n', 'o', 'p', 'q'}
	l := uint8(16)
	c := byte('j')
	for i := 0; i < b.N; i++ {
		result = unrolledIndexOfByteIn16Bytes(l, &a, c)
	}
}

func BenchmarkIndexOfByteIn16Bytes_Brute_16(b *testing.B) {
	a := [16]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'l', 'm', 'n', 'o', 'p', 'q'}
	l := uint8(16)
	c := byte('j')
	for i := 0; i < b.N; i++ {
		result = bruteIndexOfByteIn16Bytes(l, &a, c)
	}
}

func BenchmarkIndexOfByteIn16Bytes_BruteUnrolled_16(b *testing.B) {
	a := [16]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'l', 'm', 'n', 'o', 'p', 'q'}
	l := uint8(16)
	c := byte('q')
	for i := 0; i < b.N; i++ {
		result = bruteUnrolledIndexOfByteIn16Bytes(l, &a, c)
	}
}
