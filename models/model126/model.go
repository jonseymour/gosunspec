// NOTICE
// This file was automatically generated by ../../../generators/core.go. Do not edit it!
// You can regenerate it by running 'go generate ./core' from the directory above.

package model126

import (
	"github.com/crabmusket/gosunspec/core"
	"github.com/crabmusket/gosunspec/core/typelabel"
	"github.com/crabmusket/gosunspec/smdx"
)

// Block126 - Static Volt-VAR - Static Volt-VAR Arrays

const (
	ModelID = 126
)

const (
	ActCrv       = "ActCrv"
	ActPt        = "ActPt"
	CrvNam       = "CrvNam"
	DeptRef      = "DeptRef"
	DeptRef_SF   = "DeptRef_SF"
	ModEna       = "ModEna"
	NCrv         = "NCrv"
	NPt          = "NPt"
	ReadOnly     = "ReadOnly"
	RmpDecTmm    = "RmpDecTmm"
	RmpIncDec_SF = "RmpIncDec_SF"
	RmpIncTmm    = "RmpIncTmm"
	RmpTms       = "RmpTms"
	RvrtTms      = "RvrtTms"
	V1           = "V1"
	V10          = "V10"
	V11          = "V11"
	V12          = "V12"
	V13          = "V13"
	V14          = "V14"
	V15          = "V15"
	V16          = "V16"
	V17          = "V17"
	V18          = "V18"
	V19          = "V19"
	V2           = "V2"
	V20          = "V20"
	V3           = "V3"
	V4           = "V4"
	V5           = "V5"
	V6           = "V6"
	V7           = "V7"
	V8           = "V8"
	V9           = "V9"
	VAr1         = "VAr1"
	VAr10        = "VAr10"
	VAr11        = "VAr11"
	VAr12        = "VAr12"
	VAr13        = "VAr13"
	VAr14        = "VAr14"
	VAr15        = "VAr15"
	VAr16        = "VAr16"
	VAr17        = "VAr17"
	VAr18        = "VAr18"
	VAr19        = "VAr19"
	VAr2         = "VAr2"
	VAr20        = "VAr20"
	VAr3         = "VAr3"
	VAr4         = "VAr4"
	VAr5         = "VAr5"
	VAr6         = "VAr6"
	VAr7         = "VAr7"
	VAr8         = "VAr8"
	VAr9         = "VAr9"
	V_SF         = "V_SF"
	WinTms       = "WinTms"
)

type Block126Repeat struct {
	ActPt     uint16      `sunspec:"offset=0,len=1,access=rw"`
	DeptRef   core.Enum16 `sunspec:"offset=1,len=1,access=rw"`
	V1        uint16      `sunspec:"offset=2,len=1,sf=V_SF,access=rw"`
	VAr1      int16       `sunspec:"offset=3,len=1,sf=DeptRef_SF,access=rw"`
	V2        uint16      `sunspec:"offset=4,len=1,sf=V_SF,access=rw"`
	VAr2      int16       `sunspec:"offset=5,len=1,sf=DeptRef_SF,access=rw"`
	V3        uint16      `sunspec:"offset=6,len=1,sf=V_SF,access=rw"`
	VAr3      int16       `sunspec:"offset=7,len=1,sf=DeptRef_SF,access=rw"`
	V4        uint16      `sunspec:"offset=8,len=1,sf=V_SF,access=rw"`
	VAr4      int16       `sunspec:"offset=9,len=1,sf=DeptRef_SF,access=rw"`
	V5        uint16      `sunspec:"offset=10,len=1,sf=V_SF,access=rw"`
	VAr5      int16       `sunspec:"offset=11,len=1,sf=DeptRef_SF,access=rw"`
	V6        uint16      `sunspec:"offset=12,len=1,sf=V_SF,access=rw"`
	VAr6      int16       `sunspec:"offset=13,len=1,sf=DeptRef_SF,access=rw"`
	V7        uint16      `sunspec:"offset=14,len=1,sf=V_SF,access=rw"`
	VAr7      int16       `sunspec:"offset=15,len=1,sf=DeptRef_SF,access=rw"`
	V8        uint16      `sunspec:"offset=16,len=1,sf=V_SF,access=rw"`
	VAr8      int16       `sunspec:"offset=17,len=1,sf=DeptRef_SF,access=rw"`
	V9        uint16      `sunspec:"offset=18,len=1,sf=V_SF,access=rw"`
	VAr9      int16       `sunspec:"offset=19,len=1,sf=DeptRef_SF,access=rw"`
	V10       uint16      `sunspec:"offset=20,len=1,sf=V_SF,access=rw"`
	VAr10     int16       `sunspec:"offset=21,len=1,sf=DeptRef_SF,access=rw"`
	V11       uint16      `sunspec:"offset=22,len=1,sf=V_SF,access=rw"`
	VAr11     int16       `sunspec:"offset=23,len=1,sf=DeptRef_SF,access=rw"`
	V12       uint16      `sunspec:"offset=24,len=1,sf=V_SF,access=rw"`
	VAr12     int16       `sunspec:"offset=25,len=1,sf=DeptRef_SF,access=rw"`
	V13       uint16      `sunspec:"offset=26,len=1,sf=V_SF,access=rw"`
	VAr13     int16       `sunspec:"offset=27,len=1,sf=DeptRef_SF,access=rw"`
	V14       uint16      `sunspec:"offset=28,len=1,sf=V_SF,access=rw"`
	VAr14     int16       `sunspec:"offset=29,len=1,sf=DeptRef_SF,access=rw"`
	V15       uint16      `sunspec:"offset=30,len=1,sf=V_SF,access=rw"`
	VAr15     int16       `sunspec:"offset=31,len=1,sf=DeptRef_SF,access=rw"`
	V16       uint16      `sunspec:"offset=32,len=1,sf=V_SF,access=rw"`
	VAr16     int16       `sunspec:"offset=33,len=1,sf=DeptRef_SF,access=rw"`
	V17       uint16      `sunspec:"offset=34,len=1,sf=V_SF,access=rw"`
	VAr17     int16       `sunspec:"offset=35,len=1,sf=DeptRef_SF,access=rw"`
	V18       uint16      `sunspec:"offset=36,len=1,sf=V_SF,access=rw"`
	VAr18     int16       `sunspec:"offset=37,len=1,sf=DeptRef_SF,access=rw"`
	V19       uint16      `sunspec:"offset=38,len=1,sf=V_SF,access=rw"`
	VAr19     int16       `sunspec:"offset=39,len=1,sf=DeptRef_SF,access=rw"`
	V20       uint16      `sunspec:"offset=40,len=1,sf=V_SF,access=rw"`
	VAr20     int16       `sunspec:"offset=41,len=1,sf=DeptRef_SF,access=rw"`
	CrvNam    string      `sunspec:"offset=42,len=8,access=rw"`
	RmpTms    uint16      `sunspec:"offset=50,len=1,access=rw"`
	RmpDecTmm uint16      `sunspec:"offset=51,len=1,sf=RmpIncDec_SF,access=rw"`
	RmpIncTmm uint16      `sunspec:"offset=52,len=1,sf=RmpIncDec_SF,access=rw"`
	ReadOnly  core.Enum16 `sunspec:"offset=53,len=1,access=r"`
}

type Block126 struct {
	ActCrv       uint16           `sunspec:"offset=0,len=1,access=rw"`
	ModEna       core.Bitfield16  `sunspec:"offset=1,len=1,access=rw"`
	WinTms       uint16           `sunspec:"offset=2,len=1,access=rw"`
	RvrtTms      uint16           `sunspec:"offset=3,len=1,access=rw"`
	RmpTms       uint16           `sunspec:"offset=4,len=1,access=rw"`
	NCrv         uint16           `sunspec:"offset=5,len=1,access=r"`
	NPt          uint16           `sunspec:"offset=6,len=1,access=r"`
	V_SF         core.ScaleFactor `sunspec:"offset=7,len=1,access=r"`
	DeptRef_SF   core.ScaleFactor `sunspec:"offset=8,len=1,access=r"`
	RmpIncDec_SF core.ScaleFactor `sunspec:"offset=9,len=1,access=r"`

	Repeats []Block126Repeat
}

func (self *Block126) GetId() core.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "volt_var",
		Length: 64,
		Blocks: []smdx.BlockElement{
			smdx.BlockElement{
				Length: 10,
				Type:   "fixed",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ActCrv, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: ModEna, Offset: 1, Type: typelabel.Bitfield16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: WinTms, Offset: 2, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: RvrtTms, Offset: 3, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: RmpTms, Offset: 4, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: NCrv, Offset: 5, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: NPt, Offset: 6, Type: typelabel.Uint16, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: V_SF, Offset: 7, Type: typelabel.Sunssf, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: DeptRef_SF, Offset: 8, Type: typelabel.Sunssf, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: RmpIncDec_SF, Offset: 9, Type: typelabel.Sunssf, Access: "r", Length: 1},
				},
			},
			smdx.BlockElement{Name: "curve",
				Length: 54,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ActPt, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: DeptRef, Offset: 1, Type: typelabel.Enum16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: V1, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: VAr1, Offset: 3, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: V2, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr2, Offset: 5, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V3, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr3, Offset: 7, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Units: "Various", Access: "rw", Length: 1},
					smdx.PointElement{Id: V4, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "V_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr4, Offset: 9, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V5, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr5, Offset: 11, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V6, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr6, Offset: 13, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V7, Offset: 14, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr7, Offset: 15, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V8, Offset: 16, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr8, Offset: 17, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V9, Offset: 18, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr9, Offset: 19, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V10, Offset: 20, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr10, Offset: 21, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V11, Offset: 22, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr11, Offset: 23, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V12, Offset: 24, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr12, Offset: 25, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V13, Offset: 26, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr13, Offset: 27, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V14, Offset: 28, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr14, Offset: 29, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V15, Offset: 30, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr15, Offset: 31, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V16, Offset: 32, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr16, Offset: 33, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V17, Offset: 34, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr17, Offset: 35, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V18, Offset: 36, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr18, Offset: 37, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V19, Offset: 38, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr19, Offset: 39, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: V20, Offset: 40, Type: typelabel.Uint16, ScaleFactor: "V_SF", Units: "% VRef", Access: "rw", Length: 1},
					smdx.PointElement{Id: VAr20, Offset: 41, Type: typelabel.Int16, ScaleFactor: "DeptRef_SF", Access: "rw", Length: 1},
					smdx.PointElement{Id: CrvNam, Offset: 42, Type: typelabel.String, Access: "rw", Length: 8},
					smdx.PointElement{Id: RmpTms, Offset: 50, Type: typelabel.Uint16, Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: RmpDecTmm, Offset: 51, Type: typelabel.Uint16, ScaleFactor: "RmpIncDec_SF", Units: "% ref_value/min", Access: "rw", Length: 1},
					smdx.PointElement{Id: RmpIncTmm, Offset: 52, Type: typelabel.Uint16, ScaleFactor: "RmpIncDec_SF", Units: "% ref_value/min", Access: "rw", Length: 1},
					smdx.PointElement{Id: ReadOnly, Offset: 53, Type: typelabel.Enum16, Access: "r", Length: 1, Mandatory: true},
				},
			},
		}})
}
