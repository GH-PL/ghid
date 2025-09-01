package decode

import "strings"

func HashFromString(nameHash string) (Hash, bool) {
	nameHash = strings.ToUpper(nameHash)
	for hash := Hash(1); hash < maxHash; hash++ {
		if hash.String() == nameHash {
			return hash, true
		}
	}
	return 0, false
}
