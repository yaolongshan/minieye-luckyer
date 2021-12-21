package main

import (
	"code/minieye-luckyer/models/db"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	db.InitDB()
	app.Run(":8080")
}



func currPath() string {
	exec, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	path, _ := filepath.EvalSymlinks(filepath.Dir(exec))
	fmt.Println(path)
	return path
}
