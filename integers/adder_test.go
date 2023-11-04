package integers
import (
	"testing"
	"github.com/stretchr/testify/require"
)


func TestAdder_Simple(t *testing.T){
	sum := Add(2,2)
	expected := 4

	require.Equal(t,sum,expected)


}

func TestAdder_withCases_WithHelperFunctions(t *testing.T){
	t.Run("Negative_Nums", func(t *testing.T) {
		sum := Add(-500,-100)
		expected := -600
		require.Equal(t,sum,expected)


	})

	t.Run("mixed numbers",func(t *testing.T) {
		sum := Add(-10,90)
		expected:=80
		require.Equal(t,sum,expected)


	})
}

func TestAdder_withTableTesting(t *testing.T){

	addTests := []struct{
		num_array []int 
		want int
	}{
		{ []int{2,2} ,4},
		{ []int{100,2},102},
		{ []int{-100,-900},-1000},
	}

	for _, tt := range addTests{
		get:= Add(tt.num_array[0],tt.num_array[1])
		want:=tt.want

		if get!= want{
			t.Errorf("%#v Expected '%d' but got '%d' ",tt.num_array,want,get)
		}
	}

}


