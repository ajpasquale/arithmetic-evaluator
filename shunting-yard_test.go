package evaluator

import (
	"testing"
)

func TestParse(t *testing.T) {
	//tests := []struct {
	//		in   string
	//	want string
	//}{
	//{"1+1+1", "11+1+"},
	//}
	input := newQueue()

	input.enqueue("1")
	input.enqueue("+")
	input.enqueue("1")
	input.enqueue("+")
	input.enqueue("1")

	output, err := parse(*input)
	if err != nil {
		t.Fatal(err)
	}
	if output.String() != "11+1+" {
		t.Errorf("evaluator(%q)\nhave %v \nwant %v", "1+1+1", output.String(), "11+1+")
	}

	input = newQueue()
	input.enqueue("5")
	input.enqueue("+")
	input.enqueue("3")
	input.enqueue("*")
	input.enqueue("1")
	input.enqueue("-")
	input.enqueue("9")

	output, err = parse(*input)
	if err != nil {
		t.Fatal(err)
	}
	if output.String() != "531*+9-" {
		t.Errorf("evaluator(%q)\nhave %v \nwant %v", "5+3*1-9", output.String(), "531*+9-")
	}

	input = newQueue()
	input.enqueue("10")
	input.enqueue("-")
	input.enqueue("2")
	input.enqueue("*")
	input.enqueue("3")
	input.enqueue("/")
	input.enqueue("1")

	output, err = parse(*input)
	if err != nil {
		t.Fatal(err)
	}
	if output.String() != "1023*1/-" {
		t.Errorf("evaluator(%q)\nhave %v \nwant %v", "10-2*3/1", output.String(), "1023*1/-")
	}
}
