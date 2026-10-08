package collector

import (
	"encoding/json"
	"testing"
)

// Shape of an iLO 5 (v3.20) Thermal entry: missing thresholds are null.
func TestThresholdLevels(t *testing.T) {
	var temp Temperature
	body := `{"Name": "01-Inlet Ambient", "PhysicalContext": "Intake", "ReadingCelsius": 28,
		"UpperThresholdNonCritical": null, "UpperThresholdCritical": 42, "UpperThresholdFatal": 47,
		"LowerThresholdCritical": "N/A", "LowerThresholdFatal": 0}`
	if err := json.Unmarshal([]byte(body), &temp); err != nil {
		t.Fatal(err)
	}

	got := temp.Levels()
	want := []ThresholdLevel{
		{"upper", "critical", 42},
		{"upper", "fatal", 47},
		{"lower", "fatal", 0},
	}
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("level %d: got %v, want %v", i, got[i], want[i])
		}
	}
	if temp.PhysicalContext != "Intake" {
		t.Errorf("context: got %q", temp.PhysicalContext)
	}
}
