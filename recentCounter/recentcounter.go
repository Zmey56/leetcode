package main


type RecentCounter struct {
    requests []int
}

func Constructor() RecentCounter {
    return RecentCounter{}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.requests = append(rc.requests, t)

	for rc.requests[0] < t - 3000{
		rc.requests = rc.requests[1:]
	}

    return len(rc.requests)
}


