# Go Fast Hex

Attempts to make a fast hex encoder/decoder in Go.

in pkg/hexvec we have a SIMD-optimized hex encoder/decoder that can encode/decode.
On my machine, it can only encode/decode with x32 intructions. Which is about as fast as the standard library.