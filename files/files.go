package main

import (
	"fmt"
	"log"
	"os"
)

func main()  {
	count := 1000000
	for i := 0; i < count; i++ {
		err := createNewFile(i)
		if err != nil {
			log.Fatal("Do not create new file: ", err)
		}
	}	
}

func createNewFile(index int) error {
	f, err := os.Create(fmt.Sprintf("data/file_%d.file", index))
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}