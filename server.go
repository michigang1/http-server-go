package main

func main() {
	if err := runApplication(); err != nil {
	 log.Panicf("Failed to run application: %v", err)
	}
 }
