package str

import (
	"strconv"
)

// " 10 "
// -> 10, nil
// ------------------
// "#"
// 0, error
func ConvertString(num string) (int, error) {
	return strconv.Atoi(num)
}
