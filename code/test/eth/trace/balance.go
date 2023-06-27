package trace

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/shopspring/decimal"
)

// 计算代币最小单位下，总数额
// amount 数额 precision 精度
func MulStringDecimals(amount string, precision int) string {
	aDecimal, err := decimal.NewFromString(amount)
	if err != nil {
		return ""
	}
	// digit, _ := strconv.Atoi(precision)
	power := decimal.NewFromInt(int64(precision)) // 10 的 N 次方
	ten := decimal.NewFromInt(10)
	res := aDecimal.Mul(ten.Pow(power))
	return res.String()
}

// 计算代币除以精度后的总额
func DivStringDecimals(tokenTotalSupply string, precision int) string {
	if tokenTotalSupply == "" || tokenTotalSupply == "0" || precision == 0 {
		return tokenTotalSupply
	}
	tokenTotalSupplyDecimal, err := decimal.NewFromString(tokenTotalSupply)
	if err != nil {
		log.Print(err)
	}
	divisor := decimal.NewFromInt(int64(math.Pow(10, float64(precision))))
	tokenTotalSupplyDecimalNumber := tokenTotalSupplyDecimal.Div(divisor)
	return tokenTotalSupplyDecimalNumber.String()
}

// Decimals 类型，基于字符串加法
func AddStringDecimals(s1, s2 string) string {
	d1, err := decimal.NewFromString(s1)
	if err != nil {
		return s1
	}
	d2, err := decimal.NewFromString(s2)
	if err != nil {
		return s1
	}
	result := d1.Add(d2)
	return result.String()
}

// 转换供应量等大额数值
// string(float64) 转 M
func ToM(fs string) float64 {
	f, err := strconv.ParseFloat(fs, 64)
	if err != nil {
		return 0
	}
	return Float64DivFmt(f, 1000000, 2)
}

func Float64Div(f float64, divisor int) float64 {
	return f / float64(divisor)
}

// 被除数、除数、保留小数点后位数
func Float64DivFmt(f float64, divisor int, digits int) float64 {
	res := Float64Div(f, divisor)
	fs := fmt.Sprintf("%."+strconv.Itoa(digits)+"f", res)
	resFloat64, err := strconv.ParseFloat(fs, 64)
	if err != nil {
		return 0
	}
	return resFloat64
}

// 计算占比
func ProportionDivStringDecimals(s1, s2 string) float64 {
	if s2 == "" || s2 == "0" {
		return 0
	}
	d1, err := decimal.NewFromString(s1)
	if err != nil {
		return 0
	}
	d2, err := decimal.NewFromString(s2)
	if err != nil {
		return 0
	}
	result, _ := d1.Div(d2).Float64()
	return result
}
