package matchTest

import (
	"testing"
	"fifa-review/entities"
)

func TestParticularRuleAppliesToEventBasedOnOccurrences(t *testing.T) {
	m := make(map[string]string)
	event := entities.NewEvent("score", "0", &m)
	brasilEvents := []*entities.Event{ event}

	rule := entities.NewParticularRule("score", 1, 3, "", []string{}, "")
	eventMap := make(map[string][]*entities.Event)
	eventMap["score"] = brasilEvents

	rule.Apply(eventMap)
	points := event.GetFinalPoints()

	if points != 3 {
		t.Errorf("Result error. Expected 3 point and received: %d", points)
	}
}

func TestParticularRuleAppliesToEventBasedOnAfterTime(t *testing.T) {
	m := make(map[string]string)
	event := entities.NewEvent("score", "45", &m)
	brasilEvents := []*entities.Event{ event}

	rule := entities.NewParticularRule("score", 0, 2, "", []string{"90 +0"}, "")
	eventMap := make(map[string][]*entities.Event)
	eventMap["score"] = brasilEvents

	rule.Apply(eventMap)
	points := event.GetFinalPoints()

	if points != 1 {
		t.Errorf("Result error. Expected 1 point and received: %d", points)
	}
}

func TestParticularRuleAppliesToEventBasedOnPlayer(t *testing.T) {
	eventDetails := make(map[string]string)
	eventDetails["player"] = "goalkeeper"
	event := entities.NewEvent("score", "45", &eventDetails)
	brasilEvents := []*entities.Event{ event}

	rule := entities.NewParticularRule("score", 0, 3, "goalkeeper", []string{}, "")
	eventMap := make(map[string][]*entities.Event)
	eventMap["score"] = brasilEvents

	rule.Apply(eventMap)
	points := event.GetFinalPoints()

	if points != 3 {
		t.Errorf("Result error. Expected 3 point and received: %d", points)

	}
}

func TestParticularRuleAppliesToEventBasedOnDistance(t *testing.T) {
	eventDetails := make(map[string]string)
	eventDetails["distance"] = "20"
	event := entities.NewEvent("score", "0", &eventDetails)
	brasilEvents := []*entities.Event{ event}
	
	rule := entities.NewParticularRule("score", 0, 2, "", []string{}, "+25")
	eventMap := make(map[string][]*entities.Event)
	eventMap["score"] = brasilEvents

	rule.Apply(eventMap)
	points := event.GetFinalPoints()

	if points != 1 {
		t.Errorf("Result error. Expected 1 point and received: %d", points)

	}
} 