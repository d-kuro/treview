# treview
GitHub Trending viewer

# Install
```
$ go get github.com/inabajunmr/treview
```

# Usage
```
$ treview -h
Usage:
  treview is cli viewer for GitHub Trending. [flags]

Flags:
  -h, --help          help for treview
  -l, --lang string   filter by lang
  -s, --span string   trending span (default "Today")
```

# Example
```
$ treview -l go
------------------------
sourcegraph / sourcegraph	https://github.com/sourcegraph/sourcegraph
Lang:Go	Fork:40	⭐️1221	⭐️1137 stars today
Code search and intelligence, self-hosted and scalable
------------------------
spiral / roadrunner	https://github.com/spiral/roadrunner
Lang:Go	Fork:40	⭐️1129	⭐️185 stars today
High-performance PHP application server, load-balancer and process manager written in Golang
------------------------
micromdm / micromdm	https://github.com/micromdm/micromdm
Lang:Go	Fork:76	⭐️587	⭐️140 stars today
Mobile Device Management server
------------------------
```
