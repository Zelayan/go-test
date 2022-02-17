package fib

import "testing"

func TestFib(t *testing.T) {
	var (
		in = 7
		excepted = 13
	)

	actual := Fib(7)

	if actual != excepted {
		t.Errorf("Fib(%d) == %d, excepted is %d", in, actual, excepted)
	}
}

func Test(t *testing.T) {
	testCases := []struct {
		desc	string
		
	}{
		{
			desc: "",
			
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			
		})
	}
}