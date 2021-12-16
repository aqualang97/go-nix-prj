package main

import (
	"fmt"
	"my-uuid/repositories/filesystem"
)

func main() {
	repo := filesystem.UserFileRepository{}

	user := repo.GetByEmail("")
	fmt.Println(user)
}
