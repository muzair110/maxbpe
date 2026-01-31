// This function is responsible for converting the slice of bytes into the int32 datatype
package utils

func ConvertSliceToInt32(bytes []byte) []int32 {
	var newBytes []int32

	for _, val := range bytes {
		newBytes = append(newBytes, int32(val))
	}

	return newBytes
}
