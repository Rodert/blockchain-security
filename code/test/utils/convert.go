package utils

import (
	"fmt"
	"sort"
	"strconv"
	"time"
)

// Convert string to specify type.
type StrTo string

func (f StrTo) Uint8() (uint8, error) {
	v, err := strconv.ParseUint(f.String(), 10, 8)
	return uint8(v), err
}

func (f StrTo) Int() (int, error) {
	v, err := strconv.ParseInt(f.String(), 10, 0)
	return int(v), err
}

func (f StrTo) Int64() (int64, error) {
	v, err := strconv.ParseInt(f.String(), 10, 64)
	return int64(v), err
}

func (f StrTo) MustUint8() uint8 {
	v, _ := f.Uint8()
	return v
}

func (f StrTo) MustInt() int {
	v, _ := f.Int()
	return v
}

func (f StrTo) MustInt64() int64 {
	v, _ := f.Int64()
	return v
}

func (f StrTo) Float64() (float64, error) {
	return strconv.ParseFloat(f.String(), 10)
}

func (f StrTo) MustFloat64() float64 {
	v, _ := f.Float64()
	return v
}

func (f StrTo) String() string {
	return string(f)
}

// Convert any type to string.
func ToStr(value interface{}, args ...int) (s string) {
	switch v := value.(type) {
	case bool:
		s = strconv.FormatBool(v)
	case float32:
		s = strconv.FormatFloat(float64(v), 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 32))
	case float64:
		s = strconv.FormatFloat(v, 'f', argInt(args).Get(0, -1), argInt(args).Get(1, 64))
	case int:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int8:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int16:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int32:
		s = strconv.FormatInt(int64(v), argInt(args).Get(0, 10))
	case int64:
		s = strconv.FormatInt(v, argInt(args).Get(0, 10))
	case uint:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint8:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint16:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint32:
		s = strconv.FormatUint(uint64(v), argInt(args).Get(0, 10))
	case uint64:
		s = strconv.FormatUint(v, argInt(args).Get(0, 10))
	case string:
		s = v
	case []byte:
		s = string(v)
	default:
		s = fmt.Sprintf("%v", v)
	}
	return s
}

type argInt []int

func (a argInt) Get(i int, args ...int) (r int) {
	if i >= 0 && i < len(a) {
		r = a[i]
	} else if len(args) > 0 {
		r = args[0]
	}
	return
}

// HexStr2int converts hex format string to decimal number.
func HexStr2int(hexStr string) (int, error) {
	num := 0
	length := len(hexStr)
	for i := 0; i < length; i++ {
		char := hexStr[length-i-1]
		factor := -1

		switch {
		case char >= '0' && char <= '9':
			factor = int(char) - '0'
		case char >= 'a' && char <= 'f':
			factor = int(char) - 'a' + 10
		default:
			return -1, fmt.Errorf("invalid hex: %s", string(char))
		}

		num += factor * PowInt(16, i)
	}
	return num, nil
}

// Int2HexStr converts decimal number to hex format string.
func Int2HexStr(num int) (hex string) {
	if num == 0 {
		return "0"
	}

	for num > 0 {
		r := num % 16

		c := "?"
		if r >= 0 && r <= 9 {
			c = fmt.Sprint(r + '0')
		} else {
			c = fmt.Sprint(r + 'a' - 10)
		}
		hex = c + hex
		num /= 16
	}
	return hex
}

// PowInt is int type of math.Pow function.
func PowInt(x int, y int) int {
	if y <= 0 {
		return 1
	} else {
		if y%2 == 0 {
			sqrt := PowInt(x, y/2)
			return sqrt * sqrt
		} else {
			return PowInt(x, y-1) * x
		}
	}
}

func RuneAccumulation(str string) rune {
	var val rune
	for _, s := range str {
		val += rune(s)
	}
	return val
}

func StringSort(str1, str2 string) (small, big string) {
	s := []string{str1, str2}
	sort.Strings(s)
	return s[0], s[1]
}

func StringSortByRune(str1, str2 string) (small, big string) {
	if RuneAccumulation(str1) > RuneAccumulation(str2) {
		return str2, str1
	} else {
		return str1, str2
	}
}

func StringP(v string) *string { return &v }

func BoolP(v bool) *bool { return &v }

func IntP(v int) *int { return &v }

func TimeP(v time.Time) *time.Time { return &v }
