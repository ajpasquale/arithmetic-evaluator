package evaluator

import "testing"

func TestNewQueue(t *testing.T) {
	q := newQueue()

	if q == nil {
		t.Fatal("queue should be created")
	}
}

func TestQueueEnqueue(t *testing.T) {
	q := newQueue()

	if err := q.Enqueue("100"); err != nil {
		t.Fatal("enqueue should add one hundred")
	}
	if err := q.Enqueue("+"); err != nil {
		t.Fatal("enqueue should add addition token")
	}
	if err := q.Enqueue("-"); err != nil {
		t.Fatal("enqueue should add subtraction token")
	}
	if err := q.Enqueue("*"); err != nil {
		t.Fatal("enqueue should add multiplication token")
	}
	if err := q.Enqueue("/"); err != nil {
		t.Fatal("enqueue should add division token")
	}
	if err := q.Enqueue("A"); err == nil {
		t.Fatal("enqueue should not allow invalid characters")
	}
	if err := q.Enqueue("&"); err == nil {
		t.Fatal("enqueue should not allow invalid characters")
	}
}

func TestQueueDequeue(t *testing.T) {
	// evaluate "54*32*+1-"
	q := newQueue()
	q.Enqueue("5")
	q.Enqueue("4")
	q.Enqueue("*")
	q.Enqueue("3")
	q.Enqueue("2")
	q.Enqueue("*")
	q.Enqueue("+")
	q.Enqueue("1")
	q.Enqueue("-")

	if tk, _ := q.Dequeue(); tk.val != "5" || tk.typ != tokenNumeric {
		t.Fatal("dequeue should return 5 and a token of tokenNumeric")
	}
	if tk, _ := q.Dequeue(); tk.val != "4" || tk.typ != tokenNumeric {
		t.Fatal("dequeue should return 4 and a token of tokenNumeric")
	}
	if tk, _ := q.Dequeue(); tk.val != "*" || tk.typ != tokenOperatorMulti {
		t.Fatal("dequeue should return * and a token of tokenOperatorMulti")
	}
	if tk, _ := q.Dequeue(); tk.val != "3" || tk.typ != tokenNumeric {
		t.Fatal("dequeue should return 3 and a token of tokenNumeric")
	}
	if tk, _ := q.Dequeue(); tk.val != "2" || tk.typ != tokenNumeric {
		t.Fatal("dequeue should return 2 and a token of tokenNumeric")
	}
	if tk, _ := q.Dequeue(); tk.val != "*" || tk.typ != tokenOperatorMulti {
		t.Fatal("dequeue should return * and a token of tokenOperatorMulti")
	}
	if tk, _ := q.Dequeue(); tk.val != "+" || tk.typ != tokenOperatorAdd {
		t.Fatal("dequeue should return + and a token of tokenOperatorAdd")
	}
	if tk, _ := q.Dequeue(); tk.val != "1" || tk.typ != tokenNumeric {
		t.Fatal("dequeue should return 1 and a token of tokenNumeric")
	}
	if tk, _ := q.Dequeue(); tk.val != "-" || tk.typ != tokenOperatorSub {
		t.Fatal("dequeue should return + and a token of tokenOperatorSub")
	}
	if _, err := q.Dequeue(); err == nil {
		t.Fatal("dequeue should return an error")
	}
}
