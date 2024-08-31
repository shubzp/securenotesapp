package notes

import (
	uuid "github.com/hashicorp/go-uuid"
)

type Notes struct {
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

func CreateNote(content string) INotes {
	id, _ := uuid.GenerateUUID()
	return &Notes{
		Id:       id,
		Content:  content,
		Metadata: NoteMetadata{},
	}
}

func (note *Notes) Save() error {
	return nil
}

func (note *Notes) Read() (string, error) {
	return note.Content, nil
}

func (note *Notes) Edit(updatedContent string) error {
	return nil
}

func (note *Notes) Delete() error {
	return nil
}

func (note *Notes) UpdateMetadata(metadata NoteMetadata) error {
	note.Metadata = metadata
	return nil
}
