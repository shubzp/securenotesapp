package notes

import (
	uuid "github.com/hashicorp/go-uuid"
)

type Note struct {
	Id       string
	Metadata NoteMetadata
	Content  string
}

type NoteMetadata struct {
	Title          string
	CreationDate   string
	LastUpdateDate string
	Author         string
	Genre          string
}

func CreateNote(content string) INote {
	id, _ := uuid.GenerateUUID()
	return &Note{
		Id:       id,
		Content:  content,
		Metadata: NoteMetadata{},
	}
}

func (note *Note) GetId() string {
	return note.Id
}

func (note *Note) GetNote() Note {
	return *note
}

func (note *Note) Read() (string, error) {
	return note.Content, nil
}

func (note *Note) Edit(updatedContent string) error {
	return nil
}

func (note *Note) Delete() error {
	return nil
}

func (note *Note) UpdateMetadata(metadata NoteMetadata) error {
	note.Metadata = metadata
	return nil
}
