// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model802

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/core/typelabel"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block802 - Battery Base Model -

const (
	ModelID = 802
)

const (
	A                = "A"
	AChaMax          = "AChaMax"
	ADisChaMax       = "ADisChaMax"
	AHRtg            = "AHRtg"
	AHRtg_SF         = "AHRtg_SF"
	AMax_SF          = "AMax_SF"
	A_SF             = "A_SF"
	AlmRst           = "AlmRst"
	CellVAvg         = "CellVAvg"
	CellVMax         = "CellVMax"
	CellVMaxMod      = "CellVMaxMod"
	CellVMaxStr      = "CellVMaxStr"
	CellVMin         = "CellVMin"
	CellVMinMod      = "CellVMinMod"
	CellVMinStr      = "CellVMinStr"
	CellV_SF         = "CellV_SF"
	ChaSt            = "ChaSt"
	CtrlHb           = "CtrlHb"
	DisChaRte        = "DisChaRte"
	DisChaRte_SF     = "DisChaRte_SF"
	DoD              = "DoD"
	DoD_SF           = "DoD_SF"
	Evt1             = "Evt1"
	Evt2             = "Evt2"
	EvtVnd1          = "EvtVnd1"
	EvtVnd2          = "EvtVnd2"
	Hb               = "Hb"
	LocRemCtl        = "LocRemCtl"
	NCyc             = "NCyc"
	Pad1             = "Pad1"
	ReqInvState      = "ReqInvState"
	ReqW             = "ReqW"
	SetInvState      = "SetInvState"
	SetOp            = "SetOp"
	SoC              = "SoC"
	SoCMax           = "SoCMax"
	SoCMin           = "SoCMin"
	SoCRsvMin        = "SoCRsvMin"
	SoC_SF           = "SoC_SF"
	SoH              = "SoH"
	SoH_SF           = "SoH_SF"
	SocRsvMax        = "SocRsvMax"
	State            = "State"
	Typ              = "Typ"
	V                = "V"
	VMax             = "VMax"
	VMin             = "VMin"
	V_SF             = "V_SF"
	W                = "W"
	WChaDisChaMax_SF = "WChaDisChaMax_SF"
	WChaRteMax       = "WChaRteMax"
	WDisChaRteMax    = "WDisChaRteMax"
	WHRtg            = "WHRtg"
	WHRtg_SF         = "WHRtg_SF"
	W_SF             = "W_SF"
	WarrDt           = "WarrDt"
)

type Block802 struct {
	AHRtg            uint16           `sunspec:"offset=0,sf=AHRtg_SF"`
	WHRtg            uint16           `sunspec:"offset=1,sf=WHRtg_SF"`
	WChaRteMax       uint16           `sunspec:"offset=2,sf=WChaDisChaMax_SF"`
	WDisChaRteMax    uint16           `sunspec:"offset=3,sf=WChaDisChaMax_SF"`
	DisChaRte        uint16           `sunspec:"offset=4,sf=DisChaRte_SF"`
	SoCMax           uint16           `sunspec:"offset=5,sf=SoC_SF"`
	SoCMin           uint16           `sunspec:"offset=6,sf=SoC_SF"`
	SocRsvMax        uint16           `sunspec:"offset=7,sf=SoC_SF,access=rw"`
	SoCRsvMin        uint16           `sunspec:"offset=8,sf=SoC_SF,access=rw"`
	SoC              uint16           `sunspec:"offset=9,sf=SoC_SF"`
	DoD              uint16           `sunspec:"offset=10,sf=DoD_SF"`
	SoH              uint16           `sunspec:"offset=11,sf=SoH_SF"`
	NCyc             uint32           `sunspec:"offset=12"`
	ChaSt            core.Enum16      `sunspec:"offset=14"`
	LocRemCtl        core.Enum16      `sunspec:"offset=15"`
	Hb               uint16           `sunspec:"offset=16"`
	CtrlHb           uint16           `sunspec:"offset=17,access=rw"`
	AlmRst           uint16           `sunspec:"offset=18,access=rw"`
	Typ              core.Enum16      `sunspec:"offset=19"`
	State            core.Enum16      `sunspec:"offset=20"`
	Pad1             core.Pad         `sunspec:"offset=21"`
	WarrDt           uint32           `sunspec:"offset=22"`
	Evt1             core.Bitfield32  `sunspec:"offset=24"`
	Evt2             core.Bitfield32  `sunspec:"offset=26"`
	EvtVnd1          core.Bitfield32  `sunspec:"offset=28"`
	EvtVnd2          core.Bitfield32  `sunspec:"offset=30"`
	V                uint16           `sunspec:"offset=32,sf=V_SF"`
	VMax             uint16           `sunspec:"offset=33,sf=V_SF"`
	VMin             uint16           `sunspec:"offset=34,sf=V_SF"`
	CellVMax         uint16           `sunspec:"offset=35,sf=CellV_SF"`
	CellVMaxStr      uint16           `sunspec:"offset=36"`
	CellVMaxMod      uint16           `sunspec:"offset=37"`
	CellVMin         uint16           `sunspec:"offset=38,sf=CellV_SF"`
	CellVMinStr      uint16           `sunspec:"offset=39"`
	CellVMinMod      uint16           `sunspec:"offset=40"`
	CellVAvg         uint16           `sunspec:"offset=41,sf=CellV_SF"`
	A                int16            `sunspec:"offset=42,sf=A_SF"`
	AChaMax          uint16           `sunspec:"offset=43,sf=AMax_SF"`
	ADisChaMax       uint16           `sunspec:"offset=44,sf=AMax_SF"`
	W                int16            `sunspec:"offset=45,sf=W_SF"`
	ReqInvState      core.Enum16      `sunspec:"offset=46"`
	ReqW             int16            `sunspec:"offset=47,sf=W_SF"`
	SetOp            core.Enum16      `sunspec:"offset=48,access=rw"`
	SetInvState      core.Enum16      `sunspec:"offset=49,access=rw"`
	AHRtg_SF         core.ScaleFactor `sunspec:"offset=50"`
	WHRtg_SF         core.ScaleFactor `sunspec:"offset=51"`
	WChaDisChaMax_SF core.ScaleFactor `sunspec:"offset=52"`
	DisChaRte_SF     core.ScaleFactor `sunspec:"offset=53"`
	SoC_SF           core.ScaleFactor `sunspec:"offset=54"`
	DoD_SF           core.ScaleFactor `sunspec:"offset=55"`
	SoH_SF           core.ScaleFactor `sunspec:"offset=56"`
	V_SF             core.ScaleFactor `sunspec:"offset=57"`
	CellV_SF         core.ScaleFactor `sunspec:"offset=58"`
	A_SF             core.ScaleFactor `sunspec:"offset=59"`
	AMax_SF          core.ScaleFactor `sunspec:"offset=60"`
	W_SF             core.ScaleFactor `sunspec:"offset=61"`
}

func (self *Block802) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "battery",
		Length: 54,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 62,

				Points: []smdx.PointElement{
					smdx.PointElement{Id: AHRtg, Offset: 0, Type: typelabel.Uint16, ScaleFactor: "AHRtg_SF", Units: "Ah", Mandatory: true},
					smdx.PointElement{Id: WHRtg, Offset: 1, Type: typelabel.Uint16, ScaleFactor: "WHRtg_SF", Units: "Wh", Mandatory: true},
					smdx.PointElement{Id: WChaRteMax, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "WChaDisChaMax_SF", Units: "W", Mandatory: true},
					smdx.PointElement{Id: WDisChaRteMax, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "WChaDisChaMax_SF", Units: "W", Mandatory: true},
					smdx.PointElement{Id: DisChaRte, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "DisChaRte_SF", Units: "%WHRtg"},
					smdx.PointElement{Id: SoCMax, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%WHRtg"},
					smdx.PointElement{Id: SoCMin, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%WHRtg"},
					smdx.PointElement{Id: SocRsvMax, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%WHRtg", Access: "rw"},
					smdx.PointElement{Id: SoCRsvMin, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%WHRtg", Access: "rw"},
					smdx.PointElement{Id: SoC, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "SoC_SF", Units: "%WHRtg", Mandatory: true},
					smdx.PointElement{Id: DoD, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "DoD_SF", Units: "%"},
					smdx.PointElement{Id: SoH, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "SoH_SF", Units: "%"},
					smdx.PointElement{Id: NCyc, Offset: 12, Type: typelabel.Uint32},
					smdx.PointElement{Id: ChaSt, Offset: 14, Type: typelabel.Enum16},
					smdx.PointElement{Id: LocRemCtl, Offset: 15, Type: typelabel.Enum16, Mandatory: true},
					smdx.PointElement{Id: Hb, Offset: 16, Type: typelabel.Uint16},
					smdx.PointElement{Id: CtrlHb, Offset: 17, Type: typelabel.Uint16, Access: "rw"},
					smdx.PointElement{Id: AlmRst, Offset: 18, Type: typelabel.Uint16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: Typ, Offset: 19, Type: typelabel.Enum16, Mandatory: true},
					smdx.PointElement{Id: State, Offset: 20, Type: typelabel.Enum16, Mandatory: true},
					smdx.PointElement{Id: Pad1, Offset: 21, Type: typelabel.Pad, Mandatory: true},
					smdx.PointElement{Id: WarrDt, Offset: 22, Type: typelabel.Uint32},
					smdx.PointElement{Id: Evt1, Offset: 24, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: Evt2, Offset: 26, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: EvtVnd1, Offset: 28, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: EvtVnd2, Offset: 30, Type: typelabel.Bitfield32, Mandatory: true},
					smdx.PointElement{Id: V, Offset: 32, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V", Mandatory: true},
					smdx.PointElement{Id: VMax, Offset: 33, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: VMin, Offset: 34, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "V"},
					smdx.PointElement{Id: CellVMax, Offset: 35, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V"},
					smdx.PointElement{Id: CellVMaxStr, Offset: 36, Type: typelabel.Uint16},
					smdx.PointElement{Id: CellVMaxMod, Offset: 37, Type: typelabel.Uint16},
					smdx.PointElement{Id: CellVMin, Offset: 38, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V"},
					smdx.PointElement{Id: CellVMinStr, Offset: 39, Type: typelabel.Uint16},
					smdx.PointElement{Id: CellVMinMod, Offset: 40, Type: typelabel.Uint16},
					smdx.PointElement{Id: CellVAvg, Offset: 41, Type: typelabel.Uint16, ScaleFactor: "CellV_SF", Units: "V"},
					smdx.PointElement{Id: A, Offset: 42, Type: typelabel.Int16, ScaleFactor: "A_SF", Units: "A", Mandatory: true},
					smdx.PointElement{Id: AChaMax, Offset: 43, Type: typelabel.Uint16, ScaleFactor: "AMax_SF", Units: "A"},
					smdx.PointElement{Id: ADisChaMax, Offset: 44, Type: typelabel.Uint16, ScaleFactor: "AMax_SF", Units: "A"},
					smdx.PointElement{Id: W, Offset: 45, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "W", Mandatory: true},
					smdx.PointElement{Id: ReqInvState, Offset: 46, Type: typelabel.Enum16},
					smdx.PointElement{Id: ReqW, Offset: 47, Type: typelabel.Int16, ScaleFactor: "W_SF", Units: "W"},
					smdx.PointElement{Id: SetOp, Offset: 48, Type: typelabel.Enum16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: SetInvState, Offset: 49, Type: typelabel.Enum16, Access: "rw", Mandatory: true},
					smdx.PointElement{Id: AHRtg_SF, Offset: 50, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: WHRtg_SF, Offset: 51, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: WChaDisChaMax_SF, Offset: 52, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: DisChaRte_SF, Offset: 53, Type: typelabel.Sunssf},
					smdx.PointElement{Id: SoC_SF, Offset: 54, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: DoD_SF, Offset: 55, Type: typelabel.Sunssf},
					smdx.PointElement{Id: SoH_SF, Offset: 56, Type: typelabel.Sunssf},
					smdx.PointElement{Id: V_SF, Offset: 57, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: CellV_SF, Offset: 58, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: A_SF, Offset: 59, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: AMax_SF, Offset: 60, Type: typelabel.Sunssf, Mandatory: true},
					smdx.PointElement{Id: W_SF, Offset: 61, Type: typelabel.Sunssf},
				},
			},
		}})
}