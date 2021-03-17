package main

func main() {
	people := new(map[string]string)
	p := *people
	p["name"] = "Kalan" // panic: assignment to entry in nil map
}
