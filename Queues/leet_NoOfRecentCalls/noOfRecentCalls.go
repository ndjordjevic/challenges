package main

import "log"

func main() {
	rc := Constructor()

	log.Println(rc.Ping(1))
	log.Println(rc.Ping(2))
	log.Println(rc.Ping(12))
	log.Println(rc.Ping(6000))
	log.Println(rc.Ping(6008))
}

type RecentCounter struct {
	times []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		times: make([]int, 0, 10000),
	}
}

func (rc *RecentCounter) Ping(t int) int {

	rc.times = append(rc.times, t)

	for len(rc.times) > 0 && rc.times[0]+3000 < t {
		rc.times = rc.times[1:]
	}

	return len(rc.times)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
