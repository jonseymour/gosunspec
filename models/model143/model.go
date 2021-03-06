// NOTICE
// This file was automatically generated by ../../generators/models.go. Do not edit it!
// You can regenerate it by running 'go generate ./models' from the directory above.

package model143

import (
	"github.com/crabmusket/gosunspec"
	"github.com/crabmusket/gosunspec/smdx"
	"github.com/crabmusket/gosunspec/typelabel"
)

// Block143 - LFRTX - LFRT extended curve

const (
	ModelID = 143
)

const (
	ActCrv   = "ActCrv"
	ActPt    = "ActPt"
	CrvNam   = "CrvNam"
	CrvType  = "CrvType"
	Hz1      = "Hz1"
	Hz10     = "Hz10"
	Hz11     = "Hz11"
	Hz12     = "Hz12"
	Hz13     = "Hz13"
	Hz14     = "Hz14"
	Hz15     = "Hz15"
	Hz16     = "Hz16"
	Hz17     = "Hz17"
	Hz18     = "Hz18"
	Hz19     = "Hz19"
	Hz2      = "Hz2"
	Hz20     = "Hz20"
	Hz3      = "Hz3"
	Hz4      = "Hz4"
	Hz5      = "Hz5"
	Hz6      = "Hz6"
	Hz7      = "Hz7"
	Hz8      = "Hz8"
	Hz9      = "Hz9"
	Hz_SF    = "Hz_SF"
	ModEna   = "ModEna"
	NCrv     = "NCrv"
	NPt      = "NPt"
	ReadOnly = "ReadOnly"
	RmpTms   = "RmpTms"
	RvrtTms  = "RvrtTms"
	Tms1     = "Tms1"
	Tms10    = "Tms10"
	Tms11    = "Tms11"
	Tms12    = "Tms12"
	Tms13    = "Tms13"
	Tms14    = "Tms14"
	Tms15    = "Tms15"
	Tms16    = "Tms16"
	Tms17    = "Tms17"
	Tms18    = "Tms18"
	Tms19    = "Tms19"
	Tms2     = "Tms2"
	Tms20    = "Tms20"
	Tms3     = "Tms3"
	Tms4     = "Tms4"
	Tms5     = "Tms5"
	Tms6     = "Tms6"
	Tms7     = "Tms7"
	Tms8     = "Tms8"
	Tms9     = "Tms9"
	Tms_SF   = "Tms_SF"
	WinTms   = "WinTms"
)

type Block143Repeat struct {
	ActPt    uint16         `sunspec:"offset=0,len=1,access=rw"`
	Tms1     uint16         `sunspec:"offset=1,len=1,sf=Tms_SF,access=rw"`
	Hz1      uint16         `sunspec:"offset=2,len=1,sf=Hz_SF,access=rw"`
	Tms2     uint16         `sunspec:"offset=3,len=1,sf=Tms_SF,access=rw"`
	Hz2      uint16         `sunspec:"offset=4,len=1,sf=Hz_SF,access=rw"`
	Tms3     uint16         `sunspec:"offset=5,len=1,sf=Tms_SF,access=rw"`
	Hz3      uint16         `sunspec:"offset=6,len=1,sf=Hz_SF,access=rw"`
	Tms4     uint16         `sunspec:"offset=7,len=1,sf=Tms_SF,access=rw"`
	Hz4      uint16         `sunspec:"offset=8,len=1,sf=Hz_SF,access=rw"`
	Tms5     uint16         `sunspec:"offset=9,len=1,sf=Tms_SF,access=rw"`
	Hz5      uint16         `sunspec:"offset=10,len=1,sf=Hz_SF,access=rw"`
	Tms6     uint16         `sunspec:"offset=11,len=1,sf=Tms_SF,access=rw"`
	Hz6      uint16         `sunspec:"offset=12,len=1,sf=Hz_SF,access=rw"`
	Tms7     uint16         `sunspec:"offset=13,len=1,sf=Tms_SF,access=rw"`
	Hz7      uint16         `sunspec:"offset=14,len=1,sf=Hz_SF,access=rw"`
	Tms8     uint16         `sunspec:"offset=15,len=1,sf=Tms_SF,access=rw"`
	Hz8      uint16         `sunspec:"offset=16,len=1,sf=Hz_SF,access=rw"`
	Tms9     uint16         `sunspec:"offset=17,len=1,sf=Tms_SF,access=rw"`
	Hz9      uint16         `sunspec:"offset=18,len=1,sf=Hz_SF,access=rw"`
	Tms10    uint16         `sunspec:"offset=19,len=1,sf=Tms_SF,access=rw"`
	Hz10     uint16         `sunspec:"offset=20,len=1,sf=Hz_SF,access=rw"`
	Tms11    uint16         `sunspec:"offset=21,len=1,sf=Tms_SF,access=rw"`
	Hz11     uint16         `sunspec:"offset=22,len=1,sf=Hz_SF,access=rw"`
	Tms12    uint16         `sunspec:"offset=23,len=1,sf=Tms_SF,access=rw"`
	Hz12     uint16         `sunspec:"offset=24,len=1,sf=Hz_SF,access=rw"`
	Tms13    uint16         `sunspec:"offset=25,len=1,sf=Tms_SF,access=rw"`
	Hz13     uint16         `sunspec:"offset=26,len=1,sf=Hz_SF,access=rw"`
	Tms14    uint16         `sunspec:"offset=27,len=1,sf=Tms_SF,access=rw"`
	Hz14     uint16         `sunspec:"offset=28,len=1,sf=Hz_SF,access=rw"`
	Tms15    uint16         `sunspec:"offset=29,len=1,sf=Tms_SF,access=rw"`
	Hz15     uint16         `sunspec:"offset=30,len=1,sf=Hz_SF,access=rw"`
	Tms16    uint16         `sunspec:"offset=31,len=1,sf=Tms_SF,access=rw"`
	Hz16     uint16         `sunspec:"offset=32,len=1,sf=Hz_SF,access=rw"`
	Tms17    uint16         `sunspec:"offset=33,len=1,sf=Tms_SF,access=rw"`
	Hz17     uint16         `sunspec:"offset=34,len=1,sf=Hz_SF,access=rw"`
	Tms18    uint16         `sunspec:"offset=35,len=1,sf=Tms_SF,access=rw"`
	Hz18     uint16         `sunspec:"offset=36,len=1,sf=Hz_SF,access=rw"`
	Tms19    uint16         `sunspec:"offset=37,len=1,sf=Tms_SF,access=rw"`
	Hz19     uint16         `sunspec:"offset=38,len=1,sf=Hz_SF,access=rw"`
	Tms20    uint16         `sunspec:"offset=39,len=1,sf=Tms_SF,access=rw"`
	Hz20     uint16         `sunspec:"offset=40,len=1,sf=Hz_SF,access=rw"`
	CrvNam   string         `sunspec:"offset=41,len=8,access=rw"`
	ReadOnly sunspec.Enum16 `sunspec:"offset=49,len=1,access=r"`
}

type Block143 struct {
	ActCrv  uint16              `sunspec:"offset=0,len=1,access=rw"`
	ModEna  sunspec.Bitfield16  `sunspec:"offset=1,len=1,access=rw"`
	WinTms  uint16              `sunspec:"offset=2,len=1,access=rw"`
	RvrtTms uint16              `sunspec:"offset=3,len=1,access=rw"`
	RmpTms  uint16              `sunspec:"offset=4,len=1,access=rw"`
	NCrv    uint16              `sunspec:"offset=5,len=1,access=r"`
	NPt     uint16              `sunspec:"offset=6,len=1,access=r"`
	Tms_SF  sunspec.ScaleFactor `sunspec:"offset=7,len=1,access=r"`
	Hz_SF   sunspec.ScaleFactor `sunspec:"offset=8,len=1,access=r"`
	CrvType sunspec.Enum16      `sunspec:"offset=9,len=1,access=r"`

	Repeats []Block143Repeat
}

func (self *Block143) GetId() sunspec.ModelId {
	return ModelID
}

func init() {
	smdx.RegisterModel(&smdx.ModelElement{
		Id:     ModelID,
		Name:   "lfrtx",
		Length: 60,
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
					smdx.PointElement{Id: Tms_SF, Offset: 7, Type: typelabel.Sunssf, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: Hz_SF, Offset: 8, Type: typelabel.Sunssf, Access: "r", Length: 1, Mandatory: true},
					smdx.PointElement{Id: CrvType, Offset: 9, Type: typelabel.Enum16, Access: "r", Length: 1, Mandatory: true},
				},
			},
			smdx.BlockElement{Name: "curve",
				Length: 50,
				Type:   "repeating",
				Points: []smdx.PointElement{
					smdx.PointElement{Id: ActPt, Offset: 0, Type: typelabel.Uint16, Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: Tms1, Offset: 1, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: Hz1, Offset: 2, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1, Mandatory: true},
					smdx.PointElement{Id: Tms2, Offset: 3, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz2, Offset: 4, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms3, Offset: 5, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz3, Offset: 6, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms4, Offset: 7, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz4, Offset: 8, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms5, Offset: 9, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz5, Offset: 10, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms6, Offset: 11, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz6, Offset: 12, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms7, Offset: 13, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz7, Offset: 14, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms8, Offset: 15, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz8, Offset: 16, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms9, Offset: 17, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz9, Offset: 18, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms10, Offset: 19, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz10, Offset: 20, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms11, Offset: 21, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz11, Offset: 22, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms12, Offset: 23, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz12, Offset: 24, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms13, Offset: 25, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz13, Offset: 26, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms14, Offset: 27, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz14, Offset: 28, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms15, Offset: 29, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz15, Offset: 30, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms16, Offset: 31, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz16, Offset: 32, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms17, Offset: 33, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz17, Offset: 34, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms18, Offset: 35, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz18, Offset: 36, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms19, Offset: 37, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz19, Offset: 38, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: Tms20, Offset: 39, Type: typelabel.Uint16, ScaleFactor: "Tms_SF", Units: "Secs", Access: "rw", Length: 1},
					smdx.PointElement{Id: Hz20, Offset: 40, Type: typelabel.Uint16, ScaleFactor: "Hz_SF", Units: "Hz", Access: "rw", Length: 1},
					smdx.PointElement{Id: CrvNam, Offset: 41, Type: typelabel.String, Access: "rw", Length: 8},
					smdx.PointElement{Id: ReadOnly, Offset: 49, Type: typelabel.Enum16, Access: "r", Length: 1, Mandatory: true},
				},
			},
		}})
}
