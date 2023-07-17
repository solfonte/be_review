package matchTest

import (
	"testing"
	"fifa-review/entities"
)

func TestApplyingNoSpecialRulesOnMatch_BrazilShouldWin(t *testing.T) {

	m := make(map[string]string)
	brasilEvents := []*entities.Event{
		entities.NewEvent("score", "0", &m),
	}

	spainEvents := []*entities.Event{}

	match := entities.NewMatch("Brazil", brasilEvents, "Spain", spainEvents)

	match.DefineFinalResult()
	match.AssignPointsAccordingFinalResult()
	results := match.GetResults()

	brazilResults, hasBrazilResults := results["Brazil"]
	spainResults, hasSpainResults := results["Spain"]

	if !hasBrazilResults {
		t.Errorf("Missing Brazil results.")
	}

	if !hasSpainResults {
		t.Errorf("Missing Spain results.")
	}

	if brazilResults.Total_points != 3 {
		t.Errorf("Result error. Expected 3 points and Brazil results points are: %d", brazilResults.Total_points,)
	}

	if spainResults.Total_points != 0 {
		t.Errorf("Result error. Expected 0 points and Spain results points are: %d", spainResults.Total_points,)

	}
}


func TestApplyingNoSpecialRulesOnMatch_BothTeamsShouldBeEven(t *testing.T) {

	m := make(map[string]string)
	brasilEvents := []*entities.Event{
		entities.NewEvent("score", "0", &m),
	}

	spainEvents := []*entities.Event{
		entities.NewEvent("score", "90", &m),
	}

	match := entities.NewMatch("Brazil", brasilEvents, "Spain", spainEvents)

	match.DefineFinalResult()
	match.AssignPointsAccordingFinalResult()
	results := match.GetResults()

	brazilResults, hasBrazilResults := results["Brazil"]
	spainResults, hasSpainResults := results["Spain"]

	if !hasBrazilResults {
		t.Errorf("Missing Brazil results.")
	}

	if !hasSpainResults {
		t.Errorf("Missing Spain results.")
	}

	if brazilResults.Total_points != 1 {
		t.Errorf("Result error. Expected 1 points and Brazil results points are: %d", brazilResults.Total_points,)
	}

	if spainResults.Total_points != 1 {
		t.Errorf("Result error. Expected 1 points and Spain results points are: %d", spainResults.Total_points,)

	}
}



func TestApplyingParticularRulesOnMatch_BothTeamsShouldBeEven(t *testing.T) {

	m := make(map[string]string)
	brasilEvents := []*entities.Event{
		entities.NewEvent("score", "0", &m),
	}

	spainEvents := []*entities.Event{
		entities.NewEvent("score", "90", &m),
	}

	match := entities.NewMatch("Brazil", brasilEvents, "Spain", spainEvents)

	particularRule := entities.NewParticularRule("particular", 1, 2, "", []string{}, "")

	match.ApplySpecialRule(particularRule)
	match.DefineFinalResult()
	match.AssignPointsAccordingFinalResult()
	results := match.GetResults()

	brazilResults, hasBrazilResults := results["Brazil"]
	spainResults, hasSpainResults := results["Spain"]

	if !hasBrazilResults {
		t.Errorf("Missing Brazil results.")
	}

	if !hasSpainResults {
		t.Errorf("Missing Spain results.")
	}

	if brazilResults.Total_points != 1 {
		t.Errorf("Result error. Expected 1 points and Brazil results points are: %d", brazilResults.Total_points,)
	}

	if spainResults.Total_points != 1 {
		t.Errorf("Result error. Expected 1 points and Spain results points are: %d", spainResults.Total_points,)

	}
}

