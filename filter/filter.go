package filter

import (
	"github.com/inabajunmr/kevaf"
	"github.com/inabajunmr/treview/github"
)

// FilterOnlyNewComer filter only new comer(user never see)
func FilterOnlyNewComer(repos []github.Repository) {
	// check dir for storage for treview
	path := "."
	kvs := kevaf.NewMap(path)
	for repo := range repos {
		// create key from repo
		// get value from map

		// if exist
		// check first show date

		// if not same day
		// filter

		// if same day
		// not filter

		// if not exist
		// Put
		a := kvs.Get("")
	}
}
