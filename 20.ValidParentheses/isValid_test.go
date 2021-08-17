package valid

import "testing"

func TestIsValid(t *testing.T) {
	s1 := "()"
	expect1 := IsValid(s1)
	if !expect1 {
		t.Errorf("'%s' failed", s1)
	}
	s2 := ")"
	expect2 := IsValid(s2)
	if expect2 {
		t.Errorf("'%s' failed", s2)
	}
	s3 := "[])"
	expect3 := IsValid(s3)
	if expect3 {
		t.Errorf("'%s' failed", s3)
	}
}

func TestFisValid(t *testing.T) {
	s1 := "()"
	expect1 := IsValid(s1)
	if !expect1 {
		t.Errorf("'%s' failed", s1)
	}
	s2 := ")"
	expect2 := IsValid(s2)
	if expect2 {
		t.Errorf("'%s' failed", s2)
	}
	s3 := "[])"
	expect3 := IsValid(s3)
	if expect3 {
		t.Errorf("'%s' failed", s3)
	}
}
