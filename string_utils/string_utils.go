package string_utils

import (
	crand "crypto/rand"
	"github.com/google/uuid"
	"math"
	"math/big"
	mrand "math/rand"
	"time"
)

//const GlobalGenBaseStr = "ABCDEFGHIGKLMNOPQRSTUVWXYZabcdefghigklmnopqrstuvwxyz0123456789"

const GlobalGenBaseStr = "abcdefghigklmnopqrstuvwxyz0123456789"

//var GlobalGenBaseStr = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

const GlobalGenBaseStrArrayLen = 16

var GlobalGenBaseStrArray = make([]string, GlobalGenBaseStrArrayLen)

func init() {
	for i := 0; i < GlobalGenBaseStrArrayLen; i++ {
		GlobalGenBaseStrArray[i] = ShuffleStr(GlobalGenBaseStr)
	}
}

func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func ShuffleStr(baseStr string) string {

	baseRuneArray := []rune(baseStr)

	maxRandomIndex := len(baseRuneArray)

	currentNanoTime := time.Now().UnixNano()

	currentRandNum, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))

	currentMRandSeed := currentNanoTime ^ currentRandNum.Int64()

	mrand.Seed(currentMRandSeed)

	var currentBaseRuneArray []rune

	for j := 0; j < maxRandomIndex; j++ {

		var baseRuneArrayIndex = 0

		if len(baseRuneArray) == 1 {
			baseRuneArrayIndex = 0
		} else {
			baseRuneArrayIndex = mrand.Intn(len(baseRuneArray))
		}

		//fmt.Println(len(baseRuneArray), baseRuneArrayIndex)

		currentBaseRuneArray = append(currentBaseRuneArray, baseRuneArray[baseRuneArrayIndex])

		if baseRuneArrayIndex == 0 {
			baseRuneArray = baseRuneArray[1:]
		} else if baseRuneArrayIndex == len(baseRuneArray)-1 {
			baseRuneArray = baseRuneArray[0 : len(baseRuneArray)-1]
		} else {
			baseRuneArray = append(baseRuneArray[:baseRuneArrayIndex], baseRuneArray[baseRuneArrayIndex+1:]...)
		}
	}

	return string(currentBaseRuneArray)
}

func GetRandomShuffleBaseStr() string {

	currentNanoTime := time.Now().UnixNano()

	currentRandNum, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))

	currentMRandSeed := currentNanoTime ^ currentRandNum.Int64()

	//fmt.Println(currentMRandSeed)
	mrand.Seed(currentMRandSeed)

	return GlobalGenBaseStrArray[mrand.Intn(GlobalGenBaseStrArrayLen)]
}

func GenRandomStr(rLen int, baseStr string) string {

	if len(baseStr) < 1 {
		baseStr = GetRandomShuffleBaseStr()
	}

	resultRunes := make([]rune, rLen)

	maxRandomIndex := len(baseStr)

	currentNanoTime := time.Now().UnixNano()

	currentRandNum, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))

	currentMRandSeed := currentNanoTime ^ currentRandNum.Int64()

	//fmt.Println(currentMRandSeed)
	mrand.Seed(currentMRandSeed)

	for i := 0; i < rLen; i++ {
		resultRunes[i] = rune(baseStr[mrand.Intn(maxRandomIndex)])
	}

	return string(resultRunes)
}

func GenUUIDStr(rLen int, baseStr string) string {
	return uuid.NewSHA1(uuid.NameSpaceX500, []byte(GenRandomStr(rLen, baseStr))).String()
}
