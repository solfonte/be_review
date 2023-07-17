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

func (m *Match) ApplySpecialRule(rule ParticularRule) {

	for _, outcome := range m.teams {
		rule.Apply(outcome.GetEvents())
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

func (m *Match) DefineFinalResult() {
	var teamsAreEven bool 
	var teamWithMoreScorePoints string 
	var maxScorePoints int = 0

	for team, outcome := range m.teams {
		events := outcome.GetEvents()
		scorePoints := 0
		scores, hasScores := events["score"]

		if hasScores {
			amountScores := len(scores)
			for _, score := range scores {
				scorePoints += score.GetFinalPoints()
			}
			outcome.SetAmountScores(amountScores)
		}

		
		if maxScorePoints < scorePoints {
			maxScorePoints = scorePoints
			teamsAreEven = false
			teamWithMoreScorePoints = team
		} else if maxScorePoints == scorePoints {
			teamsAreEven = true
		}
    }

	if !teamsAreEven {
		outcome := m.teams[teamWithMoreScorePoints]
		outcome.WinMatch()
	} else {
		for _, outcome := range m.teams {
			outcome.DrawMatch()
		}
	}
}

func (m *Match) GetResults() map[string]Result {
	results := make(map[string]Result)
	for team, outcome := range m.teams {
		points, bonusPoints, amountScores := outcome.GetResults()
		result := Result{Total_points: points, Bonus_points: bonusPoints, Scores_in_favor_amount: amountScores}
		results[team] = result
    }
	return results
}


func (m *Match) AssignPointsAccordingFinalResult() {
	for _, outcome := range m.teams {
		outcome.AssignPointsIfWinner()
		outcome.AssignPointsIfDraw()
    }
}