package algo

import (
	"github.com/jphacks/TK_2310_1/entity"
	"sort"
	"time"
)

type Job struct {
	start, finish time.Time
	weight        int
}

func binarySearch(job []entity.Event, index int) int {
	lo, hi := 0, index-1
	for lo <= hi {
		mid := (lo + hi) / 2
		if job[mid].WillCompleteAt.Before(job[index].WillStartAt) {
			if job[mid+1].WillCompleteAt.Before(job[index].WillStartAt) {
				lo = mid + 1
			} else {
				return mid
			}
		} else {
			hi = mid - 1
		}
	}
	return -1
}

/*
func schedule(job []Job) (int, []Job) {
	sort.Slice(job, func(i, j int) bool {
		return job[i].finish.Before(job[j].finish)
	})
	n := len(job)
	dp := make([]int, n)
	dp[0] = job[0].weight
	for i := 1; i < n; i++ {
		incl := job[i].weight
		l := binarySearch(job, i)
		if l != -1 {
			incl += dp[l]
		}
		dp[i] = max(dp[i-1], incl)
	}
	selected := []Job{}
	i := n - 1
	for i >= 0 {
		if i == 0 || dp[i-1] != dp[i] {
			selected = append(selected, job[i])
			i = binarySearch(job, i) // Move to the job that doesn't conflict with job[i]
		} else {
			i--
		}
	}

	return dp[n-1], selected
}
*/

func Optimalplan(job []entity.Event) (int, []entity.Event) {
	sort.Slice(job, func(i, j int) bool {
		return job[i].WillCompleteAt.Before(job[j].WillCompleteAt)
	})
	n := len(job)
	dp := make([]int, n)
	dp[0] = job[0].UnitPrice
	for i := 1; i < n; i++ {
		incl := job[i].UnitPrice
		l := binarySearch(job, i)
		if l != -1 {
			incl += dp[l]
		}
		dp[i] = max(dp[i-1], incl)
	}
	selected := []entity.Event{}
	i := n - 1
	for i >= 0 {
		if i == 0 || dp[i-1] != dp[i] {
			selected = append(selected, job[i])
			i = binarySearch(job, i) // Move to the job that doesn't conflict with job[i]
		} else {
			i--
		}
	}

	return dp[n-1], selected
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
