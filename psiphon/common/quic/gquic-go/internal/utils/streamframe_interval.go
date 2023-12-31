package utils

import "github.com/astaguna/popon-core/psiphon/common/quic/gquic-go/internal/protocol"

// ByteInterval is an interval from one ByteCount to the other
type ByteInterval struct {
	Start protocol.ByteCount
	End   protocol.ByteCount
}
