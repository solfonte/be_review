package schemas

type RuleConditionSchema struct {
	Distance string
	After_time string
	At_least int
	Player string

}

type RuleSchema struct {
	Name string
	Type string
	Event string
	Points int
	Condition RuleConditionSchema
	Bonus_points int
	Value_factor string
}