#include <emmintrin.h>
#include <cstring>

void IndexOfByteIn16Bytes(int* n, unsigned char* arr, unsigned char* key, char* idx) {

    int mask, bitfield;
    __m128i cmp;

    // Compare the key to all 16 stored keys
    cmp = _mm_cmpeq_epi8(_mm_set1_epi8(*key),
        _mm_loadu_si128((__m128i*)arr));

    // Use a mask to ignore children that don't exist
    mask = (1 << *n) - 1;
    bitfield = _mm_movemask_epi8(cmp) & mask;
    if (bitfield) {
        *idx = char(__builtin_ctz(bitfield));
    }
}