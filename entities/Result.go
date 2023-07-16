package entities

type Result struct {
	Total_points           int
	Bonus_points           int
	Played_matches_amount  int
	Scores_in_favor_amount int
}


func (r *Result) AddTotalPoints(points int) {
	r.Total_points += points
}
