// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model10

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block10 - Communication Interface Header - To be included first for a complete interface description

const (
	ModelID = 10
)

const (
	Ctl = "Ctl"
	Pad = "Pad"
	St  = "St"
	Typ = "Typ"
)

type Block10 struct {
	St  sunspec.Enum16 `sunspec:"offset=0"`
	Ctl uint16         `sunspec:"offset=1,access=rw"`
	Typ sunspec.Enum16 `sunspec:"offset=2"`
	Pad sunspec.Pad    `sunspec:"offset=3"`
}

func (self *Block10) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 4,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 4,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: St, Offset: 0, Type: typelabel.Enum16, Mandatory: true},
					smdx.PointElement{Id: Ctl, Offset: 1, Type: typelabel.Uint16, Access: "rw"},
					smdx.PointElement{Id: Typ, Offset: 2, Type: typelabel.Enum16},
					smdx.PointElement{Id: Pad, Offset: 3, Type: typelabel.Pad},
				},
			},
		}})
}
