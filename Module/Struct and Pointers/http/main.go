package main

type Calculator struct {
	a int
	b int
}

func (c *Calculator) add(a int, b int) int {
	return a + b
}

func (c *Calculator) subtract(a int, b int) int {
	return a - b
}

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Hello, ESD SOFTWARE DEVELOPMENT!")
	// })

	// // "/" is endpoint to welcome message

	// http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Please Enter your email address!")
	// })

	// // "/login" is endpoint to login

	// http.ListenAndServe(":8080", nil)

	// fmt.Println("Server Running on: localhost:8080")

}
