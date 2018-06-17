Id package generates true random byte sequence, using "crypto/rand" package, and it also can encode it to hex. You have to set a number of bytes, which should be generated. After hex encoding, the number of the symbols is doubled.

GetRandomBytes(n int) ([]byte, error)
---
GetRandomHex(n int) ([]byte, error)
---
GetRandomHexString(n int) (string, error)
