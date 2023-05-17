package main

import "log"

type ApplicationUsage struct {
	Name  string  `json:"name"`
	Usage float64 `json:"usage"`
	Unit  string  `json:"unit"`
}
type UsageStatistic struct {
	DailyUsage       float64            `json:"daily_usage"`
	MonthlyUsage     float64            `json:"monthly_usage"`
	Unit             string             `json:"unit"`
	LastUpdate       float64            `json:"last_update"`
	ApplicationUsage []ApplicationUsage `json:"application_usage"`
}

type Rank struct {
	Level string
	Point int
}

func main() {
	ml := []Rank{}

	ml = append(ml, Rank{Level: "B", Point: 20})
	ml = append(ml, Rank{Level: "A", Point: 30})
	ml = append(ml, Rank{Level: "A", Point: 10})
	ml = append(ml, Rank{Level: "A", Point: 10})
	ml = append(ml, Rank{Level: "B", Point: 30})

	ary := make(map[string]Rank)

	for _, v := range ml {
		if v.Level == "A" {
			ary[v.Level] = Rank{Level: v.Level, Point: ary[v.Level].Point + v.Point}
		} else if v.Level == "B" {
			ary[v.Level] = Rank{Level: v.Level, Point: ary[v.Level].Point + v.Point}
		}
	}

	log.Println("LOG-SB", ary)
}
