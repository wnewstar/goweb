package tool

import (
    "fmt"
    "time"
    "math"
    "math/rand"
)

func init() {
    // 设置随机数种子
    rand.Seed(time.Now().UnixNano())
}

func GetRandNumInt(n int64) (r int64) {
    if n < 1 {
        r = 0
    } else {
        n = int64(math.Pow(10, float64(n - 1)))
        r = n + int64(rand.Intn(int(n * 9 - 1)))
    }

    return
}

func GetRandNumStr(n int64) (r string) {
    if n >= 1 {
        n = int64(math.Pow(10, float64(n - 1)))
        r = fmt.Sprint(n + int64(rand.Intn(int(n * 9 - 1))))
    }

    return
}
