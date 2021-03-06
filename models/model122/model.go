// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model122

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block122 - Measurements_Status - Inverter Controls Extended Measurements and Status

const (
	ModelID = 122
)

const (
	ActVAh      = "ActVAh"
	ActVArhQ1   = "ActVArhQ1"
	ActVArhQ2   = "ActVArhQ2"
	ActVArhQ3   = "ActVArhQ3"
	ActVArhQ4   = "ActVArhQ4"
	ActWh       = "ActWh"
	ECPConn     = "ECPConn"
	PVConn      = "PVConn"
	Ris         = "Ris"
	Ris_SF      = "Ris_SF"
	RtSt        = "RtSt"
	StActCtl    = "StActCtl"
	StSetLimMsk = "StSetLimMsk"
	StorConn    = "StorConn"
	TmSrc       = "TmSrc"
	Tms         = "Tms"
	VArAval     = "VArAval"
	VArAval_SF  = "VArAval_SF"
	WAval       = "WAval"
	WAval_SF    = "WAval_SF"
)

type Block122 struct {
	PVConn      sunspec.Bitfield16  `sunspec:"offset=0,len=1,access=r"`
	StorConn    sunspec.Bitfield16  `sunspec:"offset=1,len=1,access=r"`
	ECPConn     sunspec.Bitfield16  `sunspec:"offset=2,len=1,access=r"`
	ActWh       sunspec.Acc64       `sunspec:"offset=3,len=4,access=r"`
	ActVAh      sunspec.Acc64       `sunspec:"offset=7,len=4,access=r"`
	ActVArhQ1   sunspec.Acc64       `sunspec:"offset=11,len=4,access=r"`
	ActVArhQ2   sunspec.Acc64       `sunspec:"offset=15,len=4,access=r"`
	ActVArhQ3   sunspec.Acc64       `sunspec:"offset=19,len=4,access=r"`
	ActVArhQ4   sunspec.Acc64       `sunspec:"offset=23,len=4,access=r"`
	VArAval     int16               `sunspec:"offset=27,len=1,sf=VArAval_SF,access=r"`
	VArAval_SF  sunspec.ScaleFactor `sunspec:"offset=28,len=1,access=r"`
	WAval       uint16              `sunspec:"offset=29,len=1,sf=WAval_SF,access=r"`
	WAval_SF    sunspec.ScaleFactor `sunspec:"offset=30,len=1,access=r"`
	StSetLimMsk sunspec.Bitfield32  `sunspec:"offset=31,len=2,access=r"`
	StActCtl    sunspec.Bitfield32  `sunspec:"offset=33,len=2,access=r"`
	TmSrc       string              `sunspec:"offset=35,len=4,access=r"`
	Tms         uint32              `sunspec:"offset=39,len=2,access=r"`
	RtSt        sunspec.Bitfield16  `sunspec:"offset=41,len=1,access=r"`
	Ris         uint16              `sunspec:"offset=42,len=1,sf=Ris_SF,access=r"`
	Ris_SF      sunspec.ScaleFactor `sunspec:"offset=43,len=1,access=r"`
}

func (self *Block122) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "status",
		Length: 44,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 44,
				Type:   "fixed",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: PVConn, Offset: 0, Type: typelabel.Bitfield16, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: StorConn, Offset: 1, Type: typelabel.Bitfield16, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: ECPConn, Offset: 2, Type: typelabel.Bitfield16, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: ActWh, Offset: 3, Type: typelabel.Acc64, Units: "Wh", Access: "r", Length: 4},
					smdx.PointElement{Id: ActVAh, Offset: 7, Type: typelabel.Acc64, Units: "VAh", Access: "r", Length: 4},
					smdx.PointElement{Id: ActVArhQ1, Offset: 11, Type: typelabel.Acc64, Units: "varh", Access: "r", Length: 4},
					smdx.PointElement{Id: ActVArhQ2, Offset: 15, Type: typelabel.Acc64, Units: "varh", Access: "r", Length: 4},
					smdx.PointElement{Id: ActVArhQ3, Offset: 19, Type: typelabel.Acc64, Units: "varh", Access: "r", Length: 4},
					smdx.PointElement{Id: ActVArhQ4, Offset: 23, Type: typelabel.Acc64, Units: "varh", Access: "r", Length: 4},
					smdx.PointElement{Id: VArAval, Offset: 27, Type: typelabel.Int16, ScaleFactor: "VArAval_SF", Units: "var", Access: "r", Length: 1},
					smdx.PointElement{Id: VArAval_SF, Offset: 28, Type: typelabel.Sunssf, Access: "r", Length: 1},
					smdx.PointElement{Id: WAval, Offset: 29, Type: typelabel.Uint16, ScaleFactor: "WAval_SF", Units: "var", Access: "r", Length: 1},
					smdx.PointElement{Id: WAval_SF, Offset: 30, Type: typelabel.Sunssf, Access: "r", Length: 1},
					smdx.PointElement{Id: StSetLimMsk, Offset: 31, Type: typelabel.Bitfield32, Access: "r", Length: 2},
					smdx.PointElement{Id: StActCtl, Offset: 33, Type: typelabel.Bitfield32, Access: "r", Length: 2},
					smdx.PointElement{Id: TmSrc, Offset: 35, Type: typelabel.String, Access: "r", Length: 4},
					smdx.PointElement{Id: Tms, Offset: 39, Type: typelabel.Uint32, Units: "Secs", Access: "r", Length: 2},
					smdx.PointElement{Id: RtSt, Offset: 41, Type: typelabel.Bitfield16, Access: "r", Length: 1},
					smdx.PointElement{Id: Ris, Offset: 42, Type: typelabel.Uint16, ScaleFactor: "Ris_SF", Units: "ohms", Access: "r", Length: 1},
					smdx.PointElement{Id: Ris_SF, Offset: 43, Type: typelabel.Sunssf, Access: "r", Length: 1},
				},
			},
		}})
}
