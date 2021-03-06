// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model401

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block401 - String Combiner (Current) - A basic string combiner

const (
	ModelID = 401
)

const (
	DCA      = "DCA"
	DCAMax   = "DCAMax"
	DCA_SF   = "DCA_SF"
	DCAhr    = "DCAhr"
	DCAhr_SF = "DCAhr_SF"
	DCV      = "DCV"
	DCV_SF   = "DCV_SF"
	Evt      = "Evt"
	EvtVnd   = "EvtVnd"
	InDCA    = "InDCA"
	InDCAhr  = "InDCAhr"
	InEvt    = "InEvt"
	InEvtVnd = "InEvtVnd"
	InID     = "InID"
	N        = "N"
	Tmp      = "Tmp"
)

type Block401Repeat struct {
	InID     uint16             `sunspec:"offset=0"`
	InEvt    sunspec.Bitfield32 `sunspec:"offset=1"`
	InEvtVnd sunspec.Bitfield32 `sunspec:"offset=3"`
	InDCA    int16              `sunspec:"offset=5,sf=DCA_SF"`
	InDCAhr  uint32             `sunspec:"offset=6,sf=DCAhr_SF"`
}

type Block401 struct {
	DCA_SF   sunspec.ScaleFactor `sunspec:"offset=0"`
	DCAhr_SF sunspec.ScaleFactor `sunspec:"offset=1"`
	DCV_SF   sunspec.ScaleFactor `sunspec:"offset=2"`
	DCAMax   uint16              `sunspec:"offset=3,sf=DCA_SF"`
	N        sunspec.Count       `sunspec:"offset=4"`
	Evt      sunspec.Bitfield32  `sunspec:"offset=5"`
	EvtVnd   sunspec.Bitfield32  `sunspec:"offset=7"`
	DCA      int16               `sunspec:"offset=9,sf=DCA_SF"`
	DCAhr    uint32              `sunspec:"offset=10,sf=DCAhr_SF"`
	DCV      uint16              `sunspec:"offset=12,sf=DCV_SF"`
	Tmp      int16               `sunspec:"offset=13"`

	Repeats []Block401Repeat
}

func (self *Block401) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "string_combiner",
		Length: 22,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 14,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: DCA_SF, Offset: 0, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: DCAhr_SF, Offset: 1, Type: typelabel.Sunssf},
					smdx.PointElement{Id: DCV_SF, Offset: 2, Type: typelabel.Sunssf},
					smdx.PointElement{Id: DCAMax, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "DCA_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: N, Offset: 4, Type: typelabel.Count, Mandatory: true},
					smdx.PointElement{Id: Evt, Offset: 5, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: EvtVnd, Offset: 7, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: DCA, Offset: 9, Type: typelabel.Int16, ScaleFactor: "DCA_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: DCAhr, Offset: 10, Type: typelabel.Uint32, ScaleFactor: "DCAhr_SF", Units: "Ah"},
					smdx.PointElement{Id: DCV, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "DCV_SF", Units: "V"},
					smdx.PointElement{Id: Tmp, Offset: 13, Type: typelabel.Int16, Units: "C"},
				},
			},
			smdx.BlockElement{Name: "string",
				Length: 8,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: InID, Offset: 0, Type: typelabel.Uint16, Mandatory: true},
					smdx.PointElement{Id: InEvt, Offset: 1, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: InEvtVnd, Offset: 3, Type: typelabel.Bitfield32},
					smdx.PointElement{Id: InDCA, Offset: 5, Type: typelabel.Int16, ScaleFactor: "DCA_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: InDCAhr, Offset: 6, Type: typelabel.Uint32, ScaleFactor: "DCAhr_SF", Units: "Ah"},
				},
			},
		}})
}
