package main

func main() {

}

func aVeryBigSum(ar []int64) int64 {
	var sum int64

	for _, v := range ar {
		sum += v
	}

	return sum
}
