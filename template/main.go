package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// 1行の文字列がデフォルトを超えることがあるのでサイズを明示的に指定
var rdr = bufio.NewReaderSize(os.Stdin, 1000000)

func readline() string {
	buf := make([]byte, 0, 16)
	for {
		l, p, err := rdr.ReadLine()
		if err != nil {
			panic(err)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}

func readIntSlice() []int {
	slice := make([]int, 0)
	lines := strings.Split(readline(), " ")
	for _, v := range lines {
		slice = append(slice, toInt(v))
	}
	return slice
}

func toInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func toStr(n int) string {
	return strconv.Itoa(n)
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func abs(n int) int {
	return int(math.Abs(float64(n)))
}

func max(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func min(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func isOdd(n int) bool {
	return n&1 == 1
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// 1文字
func isUpper(s string) bool {
	return unicode.IsUpper(rune(s[0]))
}

// 1文字
func isLower(s string) bool {
	return unicode.IsUpper(rune(s[0]))
}

var (
	dr = []int{1, -1, 0, 0}
	dc = []int{0, 0, 1, -1}
)

func main() {
	N := toInt(readline())
	NS := readIntSlice()

	Println(N, NS)
}
