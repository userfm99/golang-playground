package main

import (
	"fmt"
)

// FCList ---
type FCList struct {
	FCType string `json:"fcType"`
	FCCode string `json:"fcCode"`
}

func main() {
	mStore := "K427"
	mBranch := "KZ01"

	wl := []FCList{{"BRANCH", "KZ01"}}
	bl := []FCList{{"BRANCH", "TZ01"}, {"STORE", "K446"}}

	fmt.Printf("FC status: %v\n", isFcOK(wl, bl, mBranch, mStore))

}

func isFcOK(wl, bl []FCList, currentBranch, currentStore string) bool {
	if len(wl) == 0 && len(bl) == 0 {
		return true
	}

	branchStat := false
	// check whitelist
	for _, w := range wl {
		if w.FCCode == currentStore {
			return true
		}
		if w.FCCode == currentBranch {
			branchStat = true
		}
	}

	if !branchStat {
		return false
	}

	branchBlacklisted := len(bl) == 0
	// check blacklist
	for _, w := range bl {
		if w.FCCode == currentStore {
			return false
		}
		if w.FCCode == currentBranch {
			branchBlacklisted = true
		}
	}

	if branchStat == false && branchBlacklisted == false {
		return true
	}

	return branchStat
}
