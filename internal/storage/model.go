package storage

import "time"

type UID int

type Data map[UID]map[time.Time]string
