package mai3

type MaiProtocolVersion int64

const (
	ProtocolV1 MaiProtocolVersion = 1
	ProtocolV2 MaiProtocolVersion = 2
	ProtocolV3 MaiProtocolVersion = 3
)

const (
	MaiV2GasLimit      int64 = 5000000
	MaiV2GasPerTrade   int64 = 635000 // launcher use this
	MaiV2MaxMatchGroup       = int(MaiV2GasLimit/MaiV2GasPerTrade - 2)
)
