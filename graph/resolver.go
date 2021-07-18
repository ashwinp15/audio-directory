package graph

import (
	"github.com/ashwinp15/audio-directory/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	nooble *model.Nooble
}

func (r Resolver) uploadAudio() {
	UploadObject(r.nooble.Audio)
}

func (r Resolver) newEntry() {

}

func (r Resolver) updateEntry() {

}
