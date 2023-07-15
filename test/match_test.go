package matchTest

import (
	"testing"
	"fifa-review/entities"
)

func TestApplyingParticularRulesOnMatch_BothTeamsShouldBeEven(t *testing.T) {

	brasilEvents := []*entities.Event{
		entities.NewEvent("score", "0", make(map[string]string)),
	}

	spainEvents := []*entities.Event{
		entities.NewEvent("score", "90", make(map[string]string)),
	}

	match := entities.NewMatch("Brazil", brasilEvents, "Spain", spainEvents)

	particularRule := entities.NewParticularRule("particular", 0, 2, "", "")

	match.ApplyRules(particularRule)
	match.DefineWinner()
	results := match.GetResults()

	brazilResults, hasBrazilResults := results["Brazil"]
	spainResults, hasSpainResults := results["Spain"]

	if !hasBrazilResults {
		t.Errorf("Missing Brazil results.")
	}

	if !hasSpainResults {
		t.Errorf("Missing Spain results.")
	}

	if brazilResults.Total_points != 2 {
		t.Errorf("Result error. Expected 2 points and Brazil results points are: %d", brazilResults.Total_points,)
	}

	if spainResults.Total_points != 2 {
		t.Errorf("Result error. Expected 2 points and Spain results points are: %d", spainResults.Total_points,)

	}
}