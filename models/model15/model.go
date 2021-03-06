// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model15

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block15 - Interface Counters Model - Interface counters

const (
	ModelID = 15
)

const (
	Clr       = "Clr"
	InCnt     = "InCnt"
	InDscCnt  = "InDscCnt"
	InErrCnt  = "InErrCnt"
	InNUcCnt  = "InNUcCnt"
	InUcCnt   = "InUcCnt"
	InUnkCnt  = "InUnkCnt"
	OutCnt    = "OutCnt"
	OutDscCnt = "OutDscCnt"
	OutErrCnt = "OutErrCnt"
	OutNUcCnt = "OutNUcCnt"
	OutUcCnt  = "OutUcCnt"
	Pad       = "Pad"
)

type Block15 struct {
	Clr       uint16        `sunspec:"offset=0,access=rw"`
	InCnt     sunspec.Acc32 `sunspec:"offset=1"`
	InUcCnt   sunspec.Acc32 `sunspec:"offset=3"`
	InNUcCnt  sunspec.Acc32 `sunspec:"offset=5"`
	InDscCnt  sunspec.Acc32 `sunspec:"offset=7"`
	InErrCnt  sunspec.Acc32 `sunspec:"offset=9"`
	InUnkCnt  sunspec.Acc32 `sunspec:"offset=11"`
	OutCnt    sunspec.Acc32 `sunspec:"offset=13"`
	OutUcCnt  sunspec.Acc32 `sunspec:"offset=15"`
	OutNUcCnt sunspec.Acc32 `sunspec:"offset=17"`
	OutDscCnt sunspec.Acc32 `sunspec:"offset=19"`
	OutErrCnt sunspec.Acc32 `sunspec:"offset=21"`
	Pad       sunspec.Pad   `sunspec:"offset=23"`
}

func (self *Block15) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "",
		Length: 24,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 24,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: Clr, Offset: 0, Type: typelabel.Uint16, Access: "rw"},
					smdx.PointElement{Id: InCnt, Offset: 1, Type: typelabel.Acc32},
					smdx.PointElement{Id: InUcCnt, Offset: 3, Type: typelabel.Acc32},
					smdx.PointElement{Id: InNUcCnt, Offset: 5, Type: typelabel.Acc32},
					smdx.PointElement{Id: InDscCnt, Offset: 7, Type: typelabel.Acc32},
					smdx.PointElement{Id: InErrCnt, Offset: 9, Type: typelabel.Acc32},
					smdx.PointElement{Id: InUnkCnt, Offset: 11, Type: typelabel.Acc32},
					smdx.PointElement{Id: OutCnt, Offset: 13, Type: typelabel.Acc32},
					smdx.PointElement{Id: OutUcCnt, Offset: 15, Type: typelabel.Acc32},
					smdx.PointElement{Id: OutNUcCnt, Offset: 17, Type: typelabel.Acc32},
					smdx.PointElement{Id: OutDscCnt, Offset: 19, Type: typelabel.Acc32},
					smdx.PointElement{Id: OutErrCnt, Offset: 21, Type: typelabel.Acc32},
					smdx.PointElement{Id: Pad, Offset: 23, Type: typelabel.Pad},
				},
			},
		}})
}
