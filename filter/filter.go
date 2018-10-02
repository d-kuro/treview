package filter

import (
	"strings"

	"github.com/inabajunmr/kevaf"
	"github.com/inabajunmr/treview/github"
)

// OnlyNewComer filter only new comer(user never see)
func OnlyNewComer(repos []github.Repository) {
	// TODO  check dir for storage for treview
	path := "."
	kvs, _ := kevaf.NewMap(path)
	// TODO check err

	for _, repo := range repos {
		key := createKey(repo.Name)

		// TODO create json for save

		// get value from map
		v, err := kvs.Get(key)
		if err != nil {
			continue
		}

		// if exist
		// check first show date

		// if not same day
		// filter

		// if same day
		// not filter

		// if not exist
		// Put
	}
}

func createKey(v string) string {
	return strings.Replace(strings.Replace(v, " ", "", -1), "-", "", -1)
}
