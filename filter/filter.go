package filter

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/inabajunmr/kevaf"
	"github.com/inabajunmr/treview/github"
)

type readRepository struct {
	repo     github.Repository
	readDate time.Time
}

// OnlyNewComer filter only new comer(user never see)
func OnlyNewComer(repos []github.Repository) {
	// TODO  check dir for storage for treview
	path := "."
	kvs, _ := kevaf.NewMap(path)
	// TODO check err

	for _, repo := range repos {
		key := createKey(repo.Name)

		v, err := kvs.Get(key)
		if err != nil {
			json.Marshal(readRepository{repo, time.Now()})
			continue
		}

		var repo readRepository
		err = json.Unmarshal(v, repo)
		// TODO err

		// check first show date
		isSameDate := isSameDate(repo.readDate, time.Now())

		// if not same day
		if isSameDate {
			// TODO add to repos
			continue
		}
	}
}

func createKey(v string) string {
	return strings.Replace(strings.Replace(v, " ", "", -1), "-", "", -1)
}

func isSameDate(time1 time.Time, time2 time.Time) bool {
	if time1.Year() != time2.Year() {
		return false
	}

	if time1.YearDay() != time2.YearDay() {
		return false
	}

	return true
}
