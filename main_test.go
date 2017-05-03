package main

import "testing"

func TestMain(t *testing.T) {
	testdata1 := "data/testdata.csv"
	testdata2 := "data/testdata2.csv"

	if exec(testdata1) != 51 {
		t.Fatalf("Error on exec(testdata1), result should be %s, got %s", "0.51", exec(testdata1))
	}
	if exec(testdata2) != 203 {
		t.Fatalf("Error on exec(testdata2), result should be %s, got %s", "2.03", exec(testdata2))
	}
}
