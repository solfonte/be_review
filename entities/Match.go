package entities

type Match struct {
	teams map[string]*Outcome
}

func NewMatch(awayPlayer string, awayEvents []*Event, homePlayer string, homeEvents []*Event) *Match {

	teams := make(map[string]*Outcome)
	teams[awayPlayer] = NewOutcome(awayEvents)
	teams[homePlayer] = NewOutcome(homeEvents)
	return &Match{teams: teams}
}

func (m *Match) ApplyRules(rule Rule) {

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

		outcome.SetTotalPoints(points)

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
		results[team] = *outcome.GetResults()
    }
	return results
}


func (m *Match) AssignPointsToWinner() {
	for _, outcome := range m.teams {
		outcome.AssignPointsIfWinner()
    }
}