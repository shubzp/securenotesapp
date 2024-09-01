package notes

type INote interface {
	GetId() string
	GetNote() Note
	Read() (string, error)
	Edit(content string) error
	Delete() error
	UpdateMetadata(NoteMetadata) error
}
