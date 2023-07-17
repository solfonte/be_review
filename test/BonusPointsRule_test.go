package matchTest

import (
	"testing"
	"fifa-review/entities"
)

func TestBonusPointsRuleAppliesToEventBasedOnOccurrences(t *testing.T) {
	m := make(map[string]string)
	brasilEvents := []*entities.Event{
		entities.NewEvent("score", "0", &m),
	}

	rule := entities.NewBonusPointsRule("score", 1, 1, "", []string{}, "")
	eventMap := make(map[string][]*entities.Event)
	eventMap["score"] = brasilEvents

	points := rule.Apply(eventMap)

	if points != 1 {
		t.Errorf("Result error. Expected 1 point and received: %d", points)

	}
}

func TestBonusPointsRuleAppliesToEventBasedOnAfterTime(t *testing.T) {
	m := make(map[string]string)
	brasilEvents := []*entities.Event{
		entities.NewEvent("score", "90", &m),
	}

	rule := entities.NewBonusPointsRule("score", 0, 2, "", []string{"45 +0"}, "")
	eventMap := make(map[string][]*entities.Event)
	eventMap["score"] = brasilEvents

	points := rule.Apply(eventMap)

	if points != 2 {
		t.Errorf("Result error. Expected 2 point and received: %d", points)

	}

}

func TestBonusPointsRuleAppliesToEventBasedOnPlayer(t *testing.T) {
	eventDetails := make(map[string]string)
	eventDetails["player"] = "goalkeeper"
	brasilEvents := []*entities.Event{
		entities.NewEvent("score", "90", &eventDetails),
	}

	rule := entities.NewBonusPointsRule("score", 0, 3, "", []string{}, "goalkeeper")
	eventMap := make(map[string][]*entities.Event)
	eventMap["score"] = brasilEvents

	points := rule.Apply(eventMap)

	if points != 3 {
		t.Errorf("Result error. Expected 3 point and received: %d", points)

	}
}

func TestBonusPointsRuleAppliesToEventBasedOnDistance(t *testing.T) {
	eventDetails := make(map[string]string)
	eventDetails["distance"] = "20"
	brasilEvents := []*entities.Event{
		entities.NewEvent("score", "0", &eventDetails),
	}

	rule := entities.NewBonusPointsRule("score", 0, 3, "", []string{}, "+25")
	eventMap := make(map[string][]*entities.Event)
	eventMap["score"] = brasilEvents

	points := rule.Apply(eventMap)

	if points != 0 {
		t.Errorf("Result error. Expected 0 point and received: %d", points)

	}
}