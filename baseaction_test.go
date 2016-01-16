package cdtest

import "testing"

func TestSayAction(t *testing.T) {
	a, _ := NewAction("SAY", "HI")
	if a == nil {
		t.Error("you miss to allocate new action")
	}
	x := a.Run()
	if x != "HI" {
		t.Error("Not working action")
	}
	a, _ = NewAction("sAy", "HI")
	if a == nil {
		t.Error("you miss to allocate new action")
	}
	x = a.Run()
	if x != "HI" {
		t.Error("Not working action")
	}
	a = &SayAction{Params: []string{"HI"}}
	if a == nil {
		t.Error("you miss to allocate new action")
	}
	x = a.Run()
	if x != "HI" {
		t.Error("Not working action")
	}

}

func TestCalcAction(t *testing.T) {
	a, _ := NewAction("cAlc", "1", "2", "3")
	if a == nil {
		t.Error("you miss to allocate new action")
	}
	x := a.Run()
	if x != "6" {
		t.Error("Not working action:", x)
	}
	var err error
	a, err = NewAction("cAlc", "1", "2s", "3")
	if a != nil {
		t.Error("missing expected error here")
	}
	if err == nil {
		t.Error("missing error report")
	}
}
