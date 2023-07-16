package entities

import ("fmt")

type Match struct {
	teams map[string]*Outcome
}

func NewMatch(awayPlayer string, awayEvents []*Event, homePlayer string, homeEvents []*Event) *Match {

	teams := make(map[string]*Outcome)
	teams[awayPlayer] = NewOutcome(awayEvents)
	teams[homePlayer] = NewOutcome(homeEvents)
	return &Match{teams: teams}
}

func (m *Match) ApplySpecialRule(rule ParticularRule) {

	for _, outcome := range m.teams {
		rule.Apply(outcome.GetEvents())
		fmt.Println(outcome)
    } 
}

func (m *Match) ApplyBonusPointsRule(rule BonusPointsRule) {

	for _, outcome := range m.teams {
		bonusPoints := rule.Apply(outcome.GetEvents())
		outcome.AddBonusPoints(bonusPoints)
	}
}

func (m *Match) ApplyRuleToWinner(rule MatchRule) {

	for _, outcome := range m.teams {
		rule.Apply(outcome.GetEvents())
	}
	
}

func (m *Match) DefineWinner() {
	var teamsAreEven bool 
	var teamWithMorePoints string 
	var maxPoints int = 0

	for team, outcome := range m.teams {
		events := outcome.GetEvents()
		points := 0
		scores, hasScores := events["score"]

		if hasScores {
			for _, score := range scores{
				points += score.GetFinalPoints()
			}
		}

		outcome.AddTotalPoints(points)

		if maxPoints < points {
			maxPoints = points
			teamsAreEven = false
			teamWithMorePoints = team
		} else if maxPoints == points {
			teamsAreEven = true
		}
    }

	if !teamsAreEven {
		outcome := m.teams[teamWithMorePoints]
		outcome.WinMatch()
	}
}

func (m *Match) GetResults() map[string]Result {
	results := make(map[string]Result)
	for team, outcome := range m.teams {
		points, bonusPoints := outcome.GetResults()
		result := Result{Total_points: points, Bonus_points: bonusPoints}
		results[team] = result
    }
	return results
}


func (m *Match) AssignPointsToWinner() {
	for _, outcome := range m.teams {
		outcome.AssignPointsIfWinner()
    }
}