package glib

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"sync/atomic"
	"time"
)

/* ================================================================================
 * ObjectId
 * ================================================================================ */

var objectIdCounter uint32 = 0
var objectIdMachineId = objectMachineId()

type ObjectId string

func (id ObjectId) Hex() string {
	return hex.EncodeToString([]byte(id))
}

func NewObjectId() ObjectId {
	var b [12]byte
	binary.BigEndian.PutUint32(b[:], uint32(time.Now().Unix()))

	b[4] = objectIdMachineId[0]
	b[5] = objectIdMachineId[1]
	b[6] = objectIdMachineId[2]

	pid := os.Getpid()
	b[7] = byte(pid >> 8)
	b[8] = byte(pid)

	i := atomic.AddUint32(&objectIdCounter, 1)
	b[9] = byte(i >> 16)
	b[10] = byte(i >> 8)
	b[11] = byte(i)

	return ObjectId(b[:])
}

func objectMachineId() []byte {
	var sum [3]byte
	id := sum[:]
	hostname, err1 := os.Hostname()
	if err1 != nil {
		_, err2 := io.ReadFull(rand.Reader, id)
		if err2 != nil {
			panic(fmt.Errorf("cannot get hostname: %v; %v", err1, err2))
		}
		return id
	}

	hw := md5.New()
	hw.Write([]byte(hostname))
	copy(id, hw.Sum(nil))

	return id
}
