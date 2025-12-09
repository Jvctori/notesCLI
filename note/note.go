package note

import (
	"fmt"
	"time"
)

type Note struct {
	// json tags
	// serve para setar nomes dos campos para o json
	// a func Marshal() irá olhar as tags antes de nomear os campos do JSON
	// `json:"campo"`
	//
	User          string    `json:"user"`
	NoteTitle     string    `json:"title"`
	NoteContent   string    `json:"content"`
	NoteCreatedAt time.Time `json:"created_at"`
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
	noteDContent := fmt.Sprintf("Conteúdo: \n%s\n", note.NoteContent)
	userNcreater := fmt.Sprintf("Criada por: %s\n", note.User)

	fmt.Println("*****************************************************")
	fmt.Print(noteDTitle, "\n", noteDContent, "\n", userNcreater, "\n")
}
