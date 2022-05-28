package ops

import (
	"testing"
)


// "Normal" way of testing a function
func TestSum(t *testing.T){
	result := Sum(5, 7)
	expected := 12

	if result != expected{
		t.Errorf("Expected  %v and got %v", expected, result)
	}

	result = Sum(-7, -6)

	if result != 0{
		t.Errorf("Expected more than 0")
	}
}



//With cases table we can test several cases instead of just one
//and we can avoid reassigning the function values
func TestSumWithTable(t *testing.T){

	//We declare an"on the fly" slice of struct for testing several cases
	//We need to declare the values to be used as expected or testing data
	table := []struct{
			value1 int
			value2 int
			expected int
		}{
			{5, 7, 12},
			{6, 8, 14},
			{4, 3, 7},
			{-7, -6, 0},
			
		}

		//Here we iterate over that slice, first assignin the value
		//returned from our function and after that testing each
		//of our cases. i.e: first case test 5 and 7 as arguments
		//expecting a 12 as total. The same for the other cases
		for _, property := range table{
			total := Sum(property.value1, property.value2)

			if total != property.expected{
				t.Errorf("Wrong result")
			}
		}
	}
