package schemas

type MatchTeamPairSchema struct {
	Away string
	Home string

}

type EventSchema struct {
	Event string
	Time string
	Player string
}

type MatchSchema struct {
	Teams MatchTeamPairSchema
	Home_events []EventSchema
	Away_events []EventSchema
}