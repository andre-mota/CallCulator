package main

import "testing"

const (
	testdata1 = "data/testdata.csv"
	testdata2 = "data/testdata2.csv"
)

// func TestHandleErr(t *testing.T) {
// 	err := errors.New("Error!!!")
// 	e := handleErr(err)
// 	if e != "Error!!!" {
// 		t.Fatalf("Error on handleErr(err), should be %s, got %s", "Error!!!", err)
// 	}
//}

func TestString(t *testing.T) {
	var m Money = 1575
	if m.String() != "15.75" {
		t.Fatalf("Error on m.String(), should be %s, got %s", "15.75", m.String())
	}
}

func TestFileParser(t *testing.T) {
	tCallerCounter := FileParser(testdata1)

	_, ok := tCallerCounter["+351217538222"]
	if !ok {
		t.Fatalf("Error on FileParser(testdata1), should have key %s", "+351217538222")
	}

	tCallerCounter = FileParser(testdata2)

	_, ok = tCallerCounter["+351215355312"]
	if !ok {
		t.Fatalf("Error on FileParser(testdata1), should have key %s", "+351215355312")
	}
}

func TestExec(t *testing.T) {
	if exec(testdata1) != 51 {
		t.Fatalf("Error on exec(testdata1), result should be %s, got %s", "0.51", exec(testdata1))
	}
	if exec(testdata2) != 203 {
		t.Fatalf("Error on exec(testdata2), result should be %s, got %s", "2.03", exec(testdata2))
	}
}
