package assignment2

/*
Person struct [WEIGHT = 1]
Define a struct type named "Person" with the following fields:

"Name": a string representing the person's name
"Age": an integer representing the person's age
"Gender": a string representing the person's gender
Then, define a function named "NewPerson" that takes three arguments (name string, age int, gender string) and returns a new instance of the "Person" struct.

Examples
Input: name = "Alice", age = 25, gender = "female"
Output: Person{Name: "Alice", Age: 25, Gender: "female"}

Input: name = "Bob", age = 30, gender = "male"
Output: Person{Name: "Bob", Age: 30, Gender: "male"}
*/

// Define the Person struct here
type Person struct {
	// TODO: Your Code here
	Name  string
	Age   int
	Gender string
	
}

// Define the NewPerson function here
func NewPerson(name string, age int, gender string) Person {
	return Person{
		// TODO: Your Code here
		Name: name,
		Age: age,
		Gender: gender,
	}
}
