package note

import (
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
