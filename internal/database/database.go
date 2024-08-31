package database

import (
	"errors"

	"securenotesapp.com/internal/notes"
)

type Database struct {
	notesStore []notes.Notes
}

type IDatabase interface {
	SaveNote(note notes.Notes) error
	GetNote(noteId string) (notes.Notes, error)
}

func InitializeDatabase() IDatabase {
	return &Database{
		notesStore: []notes.Notes{},
	}
}

func (dB *Database) SaveNote(note notes.Notes) error {
	dB.notesStore = append(dB.notesStore, note)
	return nil
}

func (dB *Database) GetNote(noteId string) (notes.Notes, error) {
	for _, note := range dB.notesStore {
		if note.Id == noteId {
			return note, nil
		}
	}
	return notes.Notes{}, errors.New("note not found")
}
