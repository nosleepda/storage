package storage

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/nosleepda/storage/internal/file"
	"log"
)

type Storage struct {
	files map[uuid.UUID]*file.File
}

func NewStorage() *Storage {
	return &Storage{
		files: map[uuid.UUID]*file.File{},
	}
}
func (s *Storage) Upload(filename string, blob []byte) (*file.File, error) {
	newFile, err := file.NewFile(filename, blob)
	if err != nil {
		log.Fatal(err)
	}

	s.files[newFile.ID] = newFile

	return newFile, nil
}

func (s *Storage) GetById(ID uuid.UUID) (*file.File, error) {
	storedFile, ok := s.files[ID]
	if !ok {
		return nil, fmt.Errorf("file by id %v not found", ID)
	}

	return storedFile, nil
}
