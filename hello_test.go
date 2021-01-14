package main


import "testing"


func TestHello(t *testing.T) {
	emptyResult := hello("")

	if emptyResult != "Hello, my dude!" {
		t.Errorf("hello(\"\") failed, got %v, expected %v", "Hello, my dude!", emptyResult)
	} else {
		t.Logf("hello(\"\") success!")
	}

	result := hello("Benny Bwah")

	if result != "Hello, Benny Bwah!" {
		t.Errorf("hello(\"\") failed, got %v, expected %v", "Hello, Benny Bwah!", result)
	} else {
		t.Logf("hello(\"Benny Bwah\") success!")
	}
}