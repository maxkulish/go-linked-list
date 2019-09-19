package main

import "testing"

func TestList_Size(t *testing.T)  {

	var tests = []struct{
		vals []int
		want *List
	}{
		{[]int{1}, &List{
			head:  nil,
			count: 1,
		}},
	}

	for _, test := range tests {
		lst := &List{}
		for val := range test.vals {
			lst.AddHead(val)
		}

		if lst != test.want {
			t.Errorf("Expected: %+v. Got: %+v", test.want, lst)
		}
	}

}

func TestList_AddHead(t *testing.T) {

	f := func(n int, resExp string) {
		t.Helper()

		res := NewLinkedList(n).Print()

		if res != resExp {
			t.Fatalf("unexpected result for NewLinkedList(%d)\ngot:%s\nexpected:%s", n, res, resExp)
		}
	}

	f(0, "")
	f(1, "1 ")
	f(10, "1 2 3 4 5 6 7 8 9 10 ")
	f(20, "1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 ")
}

func BenchmarkIsLoop(b *testing.B)  {
	b.ReportAllocs()

	lst := NewLinkedList(b.N)

	for i := 0; i <= b.N; i++ {
		lst.IsLoop()
	}
}

func BenchmarkList_HasLoop(b *testing.B) {
	b.ReportAllocs()

	lst := NewLinkedList(b.N)

	for i:=0; i <= b.N; i++ {
		lst.HasLoop()
	}
}