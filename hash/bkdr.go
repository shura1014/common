package hash

// BKDRHash64 64
func BKDRHash64(str []byte) uint64 {
	var (
		seed uint64 = 131
		hash uint64 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint64(str[i])
	}
	return hash
}

// BKDRHash32 32
func BKDRHash32(str []byte) uint32 {
	//strByte := bytesconv.StringToBytes(str)

	var (
		seed uint32 = 131
		hash uint32 = 0
	)
	for i := 0; i < len(str); i++ {
		hash = hash*seed + uint32(str[i])
	}
	return hash
}
