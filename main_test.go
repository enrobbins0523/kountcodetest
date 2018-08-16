package main

import(
	"testing"
	"github.com/o1egl/govatar"
)

func TestCreateAvatar(t *testing.T){
	cases:= []struct {
		in1 govatar.Gender
		in2 string
		result string
	}{
		{govatar.MALE, "mark", "mark's govatar has been created! Check C:/Users/All Users/mark.jpg!"},
		{govatar.FEMALE, "amber", "amber's govatar has been created! Check C:/Users/All Users/amber.jpg!"},
		{govatar.MALE, "curt", "curt's govatar has been created! Check C:/Users/All Users/curt.jpg!"},
	}
	for _, c := range cases {
		got := createAvatar(c.in1, c.in2)
		if got != c.result {
			t.Errorf("got: %q , want %q", got, c.result)
		}
	}
}

func BenchmarkCreateAvatarFemale(b *testing.B){
	for n := 0; n < b.N; n++ {
		createAvatar(govatar.FEMALE, "kelly")
	}
}

func BenchmarkCreateAvatarMale(b *testing.B){
	for n := 0; n < b.N; n++ {
		createAvatar(govatar.MALE, "george")
	}
}
