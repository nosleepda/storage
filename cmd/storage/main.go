package main

import (
	"fmt"
	"github.com/nosleepda/storage/internal/storage"
	"log"
)

func main() {
	fmt.Println("it works")
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it uploaded", file)

	//otherUuid, err := uuid.NewUUID()
	//if err != nil {
	//	log.Fatal(err)
	//}

	gotFile, err := st.GetById(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("it received", gotFile)

}
