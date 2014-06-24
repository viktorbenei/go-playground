package mclass

import (
	"testing"
)

func TestName(t *testing.T) {
	desired_name := "Xyz"
	p := Person{Name: desired_name}

	t.Log("Name should be %v\n", desired_name)
	if p.Name != desired_name {
		t.Fail()
	}
}

func TestShouldFail(t *testing.T) {
	t.Error("This should fail - fail example")
}
