package mai3

type MaiProtocolVersion int64

const (
	ProtocolV1 MaiProtocolVersion = 1
	ProtocolV2 MaiProtocolVersion = 2
	ProtocolV3 MaiProtocolVersion = 3
)

const (
	MaiV3GasLimit      int64 = 5000000
	MaiV3GasPerTrade   int64 = 635000 // launcher use this
	MaiV3MaxMatchGroup       = int(MaiV3GasLimit/MaiV3GasPerTrade - 2)
)
