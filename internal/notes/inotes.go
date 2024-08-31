package notes

type INotes interface {
	GetId() string
	GetNote() Notes
	Read() (string, error)
	Edit(content string) error
	Delete() error
	UpdateMetadata(NoteMetadata) error
}
