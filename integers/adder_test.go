package integers
import "testing"

func TestAdder(t *testing.T){

	addTests := []struct{
		num_array []int 
		want int
	}{
		{ []int{2,2} ,4},
		{ []int{100,2},102},
		{ []int{-100,-900},-100},
	}

	for _, tt := range addTests{
		get:= Add(tt.num_array[0],tt.num_array[1])
		want:=tt.want

		if get!= want{
			t.Errorf("%#v Expected '%d' but got '%d' ",tt.num_array,want,get)
		}
	}

}