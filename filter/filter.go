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
func OnlyNewComer(repos []github.Repository) []github.Repository {
	// TODO  check dir for storage for treview
	path := "."
	kvs, _ := kevaf.NewMap(path)
	// TODO check err

	filteredRepos := make([]github.Repository, 0)

	for _, repo := range repos {
		key := createKey(repo.Name)

		v, err := kvs.Get(key)
		if err != nil {
			json, _ := json.Marshal(readRepository{repo, time.Now()})
			// TODO err
			kvs.Put(key, json)
			// TODO err
			filteredRepos = append(filteredRepos, repo)
			continue
		}

		var readRepo readRepository
		json.Unmarshal(v, readRepo)
		// TODO err

		// check first show date
		isSameDate := isSameDate(readRepo.readDate, time.Now())

		// if not same day
		if isSameDate {
			filteredRepos = append(filteredRepos, repo)
			continue
		}
	}

	return filteredRepos
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
