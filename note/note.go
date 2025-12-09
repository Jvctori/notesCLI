package note

import (
	"fmt"
	"time"
)

type Note struct {
	User          string    `json:"User"`
	NoteTitle     string    `json:"noteTitle"`
	NoteContent   string    `json:"NoteContent"`
	NoteCreatedAt time.Time `json:"NoteCreatedAt"`
}

// Create new Note
func NewNote(user, nt, nc string) Note {
	ct := time.Now()

	return Note{
		User:          user,
		NoteTitle:     nt,
		NoteContent:   nc,
		NoteCreatedAt: ct,
	}
}

func (note Note) DisplayNote() {
	noteDTitle := fmt.Sprintf("Nota: %s\n\n", note.NoteTitle)
	noteDContent := fmt.Sprintf("Conteúdo: %s\n", note.NoteContent)
	userNcreater := fmt.Sprintf("Criada por: %s\n", note.User)

	fmt.Println("*****************************************************")
	fmt.Print(noteDTitle, "\n", noteDContent, "\n", userNcreater, "\n")
	fmt.Println("*****************************************************")
}
