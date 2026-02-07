# Benchmark


| Test Name         | Size        | Iterations | Time per Operation | Bytes per Operation | Allocations per Operation |
| ----------------- | ----------- | ---------: | ------------------ | ------------------- | ------------------------- |
| VectorEncodeUpper | Size(1)     |   30344331 | 44.94 ns/op        | 16 B/op             | 2 allocs/op               |
| VectorEncodeUpper | Size(2)     |   25402738 | 51.47 ns/op        | 16 B/op             | 2 allocs/op               |
| VectorEncodeUpper | Size(4)     |   21461568 | 49.64 ns/op        | 16 B/op             | 2 allocs/op               |
| VectorEncodeUpper | Size(8)     |   16794937 | 69.92 ns/op        | 32 B/op             | 2 allocs/op               |
| VectorEncodeUpper | Size(16)    |   20769904 | 54.96 ns/op        | 32 B/op             | 1 allocs/op               |
| VectorEncodeUpper | Size(32)    |   16074241 | 72.60 ns/op        | 64 B/op             | 1 allocs/op               |
| VectorEncodeUpper | Size(64)    |    9401102 | 128.8 ns/op        | 128 B/op            | 1 allocs/op               |
| VectorEncodeUpper | Size(128)   |    4839016 | 249.8 ns/op        | 256 B/op            | 1 allocs/op               |
| VectorEncodeUpper | Size(256)   |    2456257 | 481.8 ns/op        | 512 B/op            | 1 allocs/op               |
| VectorEncodeUpper | Size(512)   |    1242352 | 961.5 ns/op        | 1024 B/op           | 1 allocs/op               |
| VectorEncodeUpper | Size(1024)  |     680259 | 2247 ns/op         | 2048 B/op           | 1 allocs/op               |
| VectorEncodeUpper | Size(2048)  |     268635 | 4667 ns/op         | 4096 B/op           | 1 allocs/op               |
| VectorEncodeUpper | Size(4096)  |     155910 | 7816 ns/op         | 8192 B/op           | 1 allocs/op               |
| VectorEncodeUpper | Size(8192)  |      80754 | 14943 ns/op        | 16384 B/op          | 1 allocs/op               |
| VectorEncodeUpper | Size(16384) |      40042 | 28431 ns/op        | 32768 B/op          | 1 allocs/op               |
| VectorEncodeLower | Size(1)     |   91359660 | 13.95 ns/op        | 2 B/op              | 1 allocs/op               |
| VectorEncodeLower | Size(2)     |   77622174 | 15.10 ns/op        | 4 B/op              | 1 allocs/op               |
| VectorEncodeLower | Size(4)     |   63530046 | 19.85 ns/op        | 8 B/op              | 1 allocs/op               |
| VectorEncodeLower | Size(8)     |   42306999 | 28.43 ns/op        | 16 B/op             | 1 allocs/op               |
| VectorEncodeLower | Size(16)    |   21450097 | 57.05 ns/op        | 32 B/op             | 1 allocs/op               |
| VectorEncodeLower | Size(32)    |   14291804 | 79.00 ns/op        | 64 B/op             | 1 allocs/op               |
| VectorEncodeLower | Size(64)    |    8272261 | 143.7 ns/op        | 128 B/op            | 1 allocs/op               |
| VectorEncodeLower | Size(128)   |    4432632 | 270.4 ns/op        | 256 B/op            | 1 allocs/op               |
| VectorEncodeLower | Size(256)   |    2343235 | 534.3 ns/op        | 512 B/op            | 1 allocs/op               |
| VectorEncodeLower | Size(512)   |    1000000 | 1051 ns/op         | 1024 B/op           | 1 allocs/op               |
| VectorEncodeLower | Size(1024)  |     614527 | 2258 ns/op         | 2048 B/op           | 1 allocs/op               |
| VectorEncodeLower | Size(2048)  |     294907 | 4170 ns/op         | 4096 B/op           | 1 allocs/op               |
| VectorEncodeLower | Size(4096)  |     128986 | 9171 ns/op         | 8192 B/op           | 1 allocs/op               |
| VectorEncodeLower | Size(8192)  |      80143 | 15615 ns/op        | 16384 B/op          | 1 allocs/op               |
| VectorEncodeLower | Size(16384) |      38552 | 30919 ns/op        | 32768 B/op          | 1 allocs/op               |
| StdEncode         | Size(1)     |  393246772 | 3.069 ns/op        | 0 B/op              | 0 allocs/op               |
| StdEncode         | Size(2)     |  319578326 | 3.775 ns/op        | 0 B/op              | 0 allocs/op               |
| StdEncode         | Size(4)     |  197071905 | 5.940 ns/op        | 0 B/op              | 0 allocs/op               |
| StdEncode         | Size(8)     |  123898438 | 9.746 ns/op        | 0 B/op              | 0 allocs/op               |
| StdEncode         | Size(16)    |   68903740 | 17.77 ns/op        | 0 B/op              | 0 allocs/op               |
| StdEncode         | Size(32)    |   17773090 | 64.85 ns/op        | 64 B/op             | 1 allocs/op               |
| StdEncode         | Size(64)    |    9775496 | 121.8 ns/op        | 128 B/op            | 1 allocs/op               |
| StdEncode         | Size(2048)  |     345085 | 3503 ns/op         | 4096 B/op           | 1 allocs/op               |
| StdEncode         | Size(4096)  |     187275 | 9226 ns/op         | 8194 B/op           | 1 allocs/op               |
| StdEncode         | Size(8192)  |      69871 | 14700 ns/op        | 16384 B/op          | 1 allocs/op               |
| StdEncode         | Size(16384) |      44418 | 29030 ns/op        | 32768 B/op          | 1 allocs/op               |