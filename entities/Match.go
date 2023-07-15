package entities

type Match struct {
	teams map[string]Outcome
}

func NewMatch(awayPlayer string, awayEvents []Event, homePlayer string, homeEvents []Event) Match {

	teams := make(map[string]Outcome)
    teams[awayPlayer] = *NewOutcome(awayEvents)
    teams[homePlayer] = *NewOutcome(homeEvents)
	return Match{teams: teams}
}