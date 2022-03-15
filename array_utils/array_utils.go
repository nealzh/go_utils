package array_utils

type NumericInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type NumericUint interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type NumericFloat interface {
	~float32 | ~float64
}

func SumIntArray[V NumericInt](valueArray []V) int64 {

	if valueArray == nil || len(valueArray) < 1 {
		panic("输入的数组不能为空！")
	}

	var res int64 = 0
	for i := 0; i < len(valueArray); i++ {
		res = res + int64(valueArray[i])
	}
	return res
}

func SumUintArray[V NumericUint](valueArray []V) uint64 {

	if valueArray == nil || len(valueArray) < 1 {
		panic("输入的数组不能为空！")
	}

	var res uint64 = 0
	for i := 0; i < len(valueArray); i++ {
		res = res + uint64(valueArray[i])
	}
	return res
}

func SumFloatArray[V NumericFloat](valueArray []V) float64 {

	if valueArray == nil || len(valueArray) < 1 {
		panic("输入的数组不能为空！")
	}

	var res float64 = 0.0
	for i := 0; i < len(valueArray); i++ {
		res = res + float64(valueArray[i])
	}
	return res
}

func MaxArray[V NumericInt | NumericUint | NumericFloat](valueArray []V) V {

	if valueArray == nil || len(valueArray) < 1 {
		panic("输入的数组不能为空！")
	}

	res := valueArray[0]

	for i := 0; i < len(valueArray); i++ {
		if valueArray[i] > res {
			res = valueArray[i]
		}
	}

	return res
}

func MinArray[V NumericInt | NumericUint | NumericFloat](valueArray []V) V {

	if valueArray == nil || len(valueArray) < 1 {
		panic("输入的数组不能为空！")
	}

	res := valueArray[0]

	for i := 0; i < len(valueArray); i++ {
		if valueArray[i] < res {
			res = valueArray[i]
		}
	}

	return res
}

func MeanArray[V NumericInt | NumericUint | NumericFloat](valueArray []V) float64 {

	if valueArray == nil || len(valueArray) < 1 {
		panic("输入的数组不能为空！")
	}

	return SumFloatArray(valueArray) / float64(len(valueArray))
}
