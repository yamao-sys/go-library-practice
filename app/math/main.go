package main

import (
	"fmt"
	"math"
)

func main() {
	// // 数学的な定義
	// // 円周率
	// fmt.Println(math.Pi)

	// // 2の平方根
	// fmt.Println(math.Sqrt2)

	// // 数値型に関する定数
	// // float32で表現可能な最大値
	// fmt.Println(math.MaxFloat32)

	// // float32で表現可能な0でない最小値
	// fmt.Println(math.SmallestNonzeroFloat32)

	// // float64で表現可能な最大値
	// fmt.Println(math.MaxFloat64)

	// // float64で表現可能な0でない最小値
	// fmt.Println(math.SmallestNonzeroFloat64)

	// // int64 ver
	// fmt.Println(math.MaxInt64)
	// fmt.Println(math.MinInt64)

	// // 絶対値
	// fmt.Println(math.Abs(100))
	// fmt.Println(math.Abs(-100))

	// // 累乗を求める
	// fmt.Println(math.Pow(0, 2)) // 0の2乗
	// fmt.Println(math.Pow(2, 2)) // 2の2乗

	// // 平方根, 立方根
	// fmt.Println(math.Sqrt2)
	// fmt.Println(math.Cbrt(8)) // 8の立方根

	// // 最大値, 最小値
	// fmt.Println(math.Max(1, 2))
	// fmt.Println(math.Min(1, 2))

	// 小数点以下の切り捨て、切り上げ

	// math.Trunc 数値の正負を問わず、小数点以下を切り捨てる
	fmt.Println(math.Trunc(3.14))
	fmt.Println(math.Trunc(-3.14))

	// math.Floor 引数の数値より小さい最大の整数を返す
	fmt.Println(math.Floor(3.5))
	fmt.Println(math.Floor(-3.5))

	// math.Ceil 引数の数値より大きい最小の整数を返す
	fmt.Println(math.Ceil(3.5))
	fmt.Println(math.Ceil(-3.5))
}
