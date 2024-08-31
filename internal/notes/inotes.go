package notes

type INotes interface {
	Save() error
	Read() (string, error)
	Edit(content string) error
	Delete() error
	UpdateMetadata(NoteMetadata) error
}
