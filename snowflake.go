package glib

import (
	"encoding/base64"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"
)

/* ================================================================================
 * snowflake id
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

type (
	Snowflake struct {
		Epoch           int64
		NodeBits        uint8
		SeqBits         uint8
		encodeBase32Map string
		encodeBase58Map string
		decodeBase32Map [256]byte
		decodeBase58Map [256]byte

		nodeMax   int64
		nodeMask  int64
		seqMask   int64
		timeShift uint8
		nodeShift uint8
	}

	SnowflakeNode struct {
		mu   sync.Mutex
		time int64
		node int64
		seq  int64
	}

	SnowflakeId int64

	SnowflakeJSONError struct{ original []byte }
)

var snowflake *Snowflake

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化Snowflake
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewSnowflake() *Snowflake {
	snowflake = &Snowflake{
		Epoch:           1551377400197,
		NodeBits:        10,
		SeqBits:         12,
		encodeBase32Map: "ybndrfg8ejkmcpqxot1uwisza345h769",
		encodeBase58Map: "123456789abcdefghijkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ",
	}

	snowflake.init()

	return snowflake
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化Snowflake
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Snowflake) init() {
	s.nodeMax = -1 ^ (-1 << s.NodeBits)
	s.nodeMask = s.nodeMax << s.SeqBits
	s.seqMask = -1 ^ (-1 << s.SeqBits)
	s.timeShift = s.NodeBits + s.SeqBits
	s.nodeShift = s.SeqBits

	for i := 0; i < len(s.encodeBase58Map); i++ {
		s.decodeBase58Map[i] = 0xFF
	}

	for i := 0; i < len(s.encodeBase58Map); i++ {
		s.decodeBase58Map[s.encodeBase58Map[i]] = byte(i)
	}

	for i := 0; i < len(s.encodeBase32Map); i++ {
		s.decodeBase32Map[i] = 0xFF
	}

	for i := 0; i < len(s.encodeBase32Map); i++ {
		s.decodeBase32Map[s.encodeBase32Map[i]] = byte(i)
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取节点对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Snowflake) Node(node int64) (*SnowflakeNode, error) {
	// re-calc in case custom NodeBits or SeqBits were set
	s.nodeMax = -1 ^ (-1 << s.NodeBits)
	s.nodeMask = s.nodeMax << s.SeqBits
	s.seqMask = -1 ^ (-1 << s.SeqBits)
	s.timeShift = s.NodeBits + s.SeqBits
	s.nodeShift = s.SeqBits

	if node < 0 || node > s.nodeMax {
		return nil, errors.New("SnowflakeNode number must be between 0 and " + strconv.FormatInt(s.nodeMax, 10))
	}

	return &SnowflakeNode{
		time: 0,
		node: node,
		seq:  0,
	}, nil
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取唯一Id
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (n *SnowflakeNode) Id() SnowflakeId {
	n.mu.Lock()

	now := time.Now().UnixNano() / 1000000

	if n.time == now {
		n.seq = (n.seq + 1) & snowflake.seqMask

		if n.seq == 0 {
			for now <= n.time {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		n.seq = 0
	}

	n.time = now

	r := SnowflakeId((now-snowflake.Epoch)<<snowflake.timeShift |
		(n.node << snowflake.nodeShift) |
		(n.seq),
	)

	n.mu.Unlock()
	return r
}

// Int64 returns an int64 of the snowflake Id
func (f SnowflakeId) Int64() int64 {
	return int64(f)
}

// String returns a string of the snowflake Id
func (f SnowflakeId) String() string {
	return strconv.FormatInt(int64(f), 10)
}

// Bytes returns a byte slice of the snowflake Id
func (f SnowflakeId) Bytes() []byte {
	return []byte(f.String())
}

// IntBytes returns an array of bytes of the snowflake Id, encoded as a
// big endian integer.
func (f SnowflakeId) IntBytes() [8]byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(f))
	return b
}

// Base2 returns a string base2 of the snowflake Id
func (f SnowflakeId) Base2() string {
	return strconv.FormatInt(int64(f), 2)
}

// Base32 uses the z-base-32 character set but encodes and decodes similar
// to base58, allowing it to create an even smaller result string.
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation interop with other packages.
func (f SnowflakeId) Base32() string {
	if f < 32 {
		return string(snowflake.encodeBase32Map[f])
	}

	b := make([]byte, 0, 12)
	for f >= 32 {
		b = append(b, snowflake.encodeBase32Map[f%32])
		f /= 32
	}
	b = append(b, snowflake.encodeBase32Map[f])

	for x, y := 0, len(b)-1; x < y; x, y = x+1, y-1 {
		b[x], b[y] = b[y], b[x]
	}

	return string(b)
}

// Base36 returns a base36 string of the snowflake Id
func (f SnowflakeId) Base36() string {
	return strconv.FormatInt(int64(f), 36)
}

// Base58 returns a base58 string of the snowflake Id
func (f SnowflakeId) Base58() string {
	if f < 58 {
		return string(snowflake.encodeBase58Map[f])
	}

	b := make([]byte, 0, 11)
	for f >= 58 {
		b = append(b, snowflake.encodeBase58Map[f%58])
		f /= 58
	}
	b = append(b, snowflake.encodeBase58Map[f])

	for x, y := 0, len(b)-1; x < y; x, y = x+1, y-1 {
		b[x], b[y] = b[y], b[x]
	}

	return string(b)
}

// Base64 returns a base64 string of the snowflake Id
func (f SnowflakeId) Base64() string {
	return base64.StdEncoding.EncodeToString(f.Bytes())
}

// Time returns an int64 unix timestamp of the snowflake Id time
func (f SnowflakeId) Time() int64 {
	return (int64(f) >> snowflake.timeShift) + snowflake.Epoch
}

// Node returns an int64 of the snowflake Id node number
func (f SnowflakeId) Node() int64 {
	return int64(f) & snowflake.nodeMask >> snowflake.nodeShift
}

// Step returns an int64 of the snowflake seq (or sequence) number
func (f SnowflakeId) Seq() int64 {
	return int64(f) & snowflake.seqMask
}

// MarshalJSON returns a json byte array string of the snowflake Id.
func (f SnowflakeId) MarshalJSON() ([]byte, error) {
	buff := make([]byte, 0, 22)
	buff = append(buff, '"')
	buff = strconv.AppendInt(buff, int64(f), 10)
	buff = append(buff, '"')
	return buff, nil
}

// UnmarshalJSON converts a json byte array of a snowflake Id into an Id type.
func (f *SnowflakeId) UnmarshalJSON(b []byte) error {
	if len(b) < 3 || b[0] != '"' || b[len(b)-1] != '"' {
		return SnowflakeJSONError{b}
	}

	i, err := strconv.ParseInt(string(b[1:len(b)-1]), 10, 64)
	if err != nil {
		return err
	}

	*f = SnowflakeId(i)
	return nil
}

func (s SnowflakeJSONError) Error() string {
	return fmt.Sprintf("invalid snowflake id %q", string(s.original))
}

// ParseBase32 parses a base32 []byte into a snowflake Id
// NOTE: There are many different base32 implementations so becareful when
// doing any interoperation interop with other packages.
func ParseBase32(b []byte) (SnowflakeId, error) {
	var id int64

	for i := range b {
		if snowflake.decodeBase32Map[b[i]] == 0xFF {
			return -1, errors.New("invalid base32")
		}
		id = id*32 + int64(snowflake.decodeBase32Map[b[i]])
	}

	return SnowflakeId(id), nil
}

// ParseBase58 parses a base58 []byte into a snowflake Id
func ParseBase58(b []byte) (SnowflakeId, error) {
	var id int64

	for i := range b {
		if snowflake.decodeBase58Map[b[i]] == 0xFF {
			return -1, errors.New("invalid base58")
		}
		id = id*58 + int64(snowflake.decodeBase58Map[b[i]])
	}

	return SnowflakeId(id), nil
}
