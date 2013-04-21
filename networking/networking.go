package networking

import (
	"time"
)

const (
	// Packets over this many bytes will be truncated.
	MaxPacketSize = 8192

	// "Ensured" packets will be dropped if they are unsuccessfully sent
	// this many times.
	MaxRetries = 64

	// "Ensured" packets will be retried after this duration until there
	// is a notification of arrival or MaxRetries attempts have been made.
	EnsureInterval = time.Millisecond * 200
)
