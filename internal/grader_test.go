package internal

import "testing"

func TestPercentageGrader(t *testing.T) {
	e := Enrollment{score: 0.875, Grader: PercentageGrader{}}
	got, _ := e.Grade(e)
	want := "87.5%"
	if got != want {
		t.Fatalf("got %s want %s", got, want)
	}
}
func TestPassFailGrader(t *testing.T) {
	p := PassFailGrader{PassMark: 0.6}
	if g, _ := p.Grade(Enrollment{score: 0.7}); g != "PASS" {
		t.Fatal("expected PASS")
	}
	if g, _ := p.Grade(Enrollment{score: 0.2}); g != "FAIL" {
		t.Fatal("expected FAIL")
	}
}
