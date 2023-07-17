package matchTest

import (
	"testing"
	"fifa-review/entities"
)

func TestMatchRuleAppliesToEvent(t *testing.T) {
	m := make(map[string]string)
	event := entities.NewEvent("score", "0", &m)
	brasilEvents := []*entities.Event{
		event,
	}

	rule := entities.NewMatchRule("win", 3)
	eventMap := make(map[string][]*entities.Event)
	eventMap["win"] = brasilEvents

	rule.Apply(eventMap)
	points := event.GetFinalPoints()

	if points != 3 {
		t.Errorf("Result error. Expected 3 point and received: %d", points)

	}
}