package assignment2

import "testing"

func TestNewPerson(t *testing.T) {
    name := "Alice"
    age := 25
    gender := "female"
    p := NewPerson(name, age, gender)
    expected := Person{Name: name, Age: age, Gender: gender}
    if p != expected {
        t.Errorf("NewPerson(%q, %d, %q) = %v; expected %v", name, age, gender, p, expected)
    }
    
    name = "Bob"
    age = 30
    gender = "male"
    p = NewPerson(name, age, gender)
    expected = Person{Name: name, Age: age, Gender: gender}
    if p != expected {
        t.Errorf("NewPerson(%q, %d, %q) = %v; expected %v", name, age, gender, p, expected)
    }
    
}
