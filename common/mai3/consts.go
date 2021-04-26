package mai3

type MaiProtocolVersion int64

const (
	ProtocolV1 MaiProtocolVersion = 1
	ProtocolV2 MaiProtocolVersion = 2
	ProtocolV3 MaiProtocolVersion = 3
)

const (
	MaiV3BaseGas        int64 = 170000
	MaiV3GasForEachPerp int64 = 100000
)
