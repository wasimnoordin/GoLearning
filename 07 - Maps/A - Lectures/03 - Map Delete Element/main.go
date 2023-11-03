package main

func main() {
	am := make(map[string]int)
	am["Ligma"] = 28
	am["Deez"] = 37
	am["Updog"] = 99

	delete(am, "Deez")
	delete(am, "Deiez") // won't panic
}
