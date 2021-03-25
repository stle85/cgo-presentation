package main

import "fmt"

func main() {
	fmt.Println("[INFO] test app started...")

	h, err := OpenLib("lib/libTest.so")
	if err != nil {
		fmt.Printf("[ERROR] OpenLib error: %v", err)
	}

	err = CloseLib(h)
	if err != nil {
		fmt.Printf("[ERROR] CloseLib error: %v", err)
	}

	fmt.Println("[INFO] test app finished...")
}
