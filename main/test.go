package main

import "fmt"

func testNoStackErrors() error {
	return fmt.Errorf("找不到github")
}
