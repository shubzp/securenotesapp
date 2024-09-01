package database

import (
	"errors"

	"securenotesapp.com/internal/notes"
	"securenotesapp.com/internal/users"
)

type Database struct {
	noteStore []notes.Note
	userStore []users.User
}

type IDatabase interface {
	SaveNote(note notes.Note) error
	GetNote(noteId string) (notes.Note, error)
	AddUser(user users.User) error
	VerifyUserCredentials(userName, password string) (bool, error)
	UpdateUserInfo(userInfo users.User) error
}

func InitializeDatabase() IDatabase {
	return &Database{
		noteStore: []notes.Note{},
		userStore: []users.User{},
	}
}

func (dB *Database) SaveNote(note notes.Note) error {
	dB.noteStore = append(dB.noteStore, note)
	return nil
}

func (dB *Database) GetNote(noteId string) (notes.Note, error) {
	for _, note := range dB.noteStore {
		if note.Id == noteId {
			return note, nil
		}
	}
	return notes.Note{}, errors.New("note not found")
}

func (dB *Database) AddUser(user users.User) error {
	dB.userStore = append(dB.userStore, user)
	return nil
}

func (dB *Database) VerifyUserCredentials(userName, password string) (bool, error) {
	for _, user := range dB.userStore {
		if user.UserName == userName {
			if user.Password == password {
				return true, nil
			} else {
				return false, nil
			}
		}
	}
	return false, errors.New("user not found")
}

func (dB *Database) UpdateUserInfo(userInfo users.User) error {
	for i, user := range dB.userStore {
		if user.UserId == userInfo.UserId {
			dB.userStore[i] = userInfo
			return nil
		}
	}

	return errors.New("user not found")
}
