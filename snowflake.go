package glib

import (
	"encoding/base64"
	"encoding/binary"
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
		Epoch    int64
		NodeBits uint8
		SeqBits  uint8

		nodeMax   int64
		nodeMask  int64
		seqMask   int64
		timeShift uint8
		nodeShift uint8
	}

	SnowflakeNode struct {
		mu        sync.Mutex
		timestamp int64
		nodeId    int64
		seq       int64
	}

	SnowflakeId int64
)

var snowflake *Snowflake

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 实例化Snowflake
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func NewSnowflake() *Snowflake {
	snowflake = &Snowflake{
		Epoch:    1551377400197,
		NodeBits: 10,
		SeqBits:  12,
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
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取节点对象
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (s *Snowflake) Node(nodeId int64) *SnowflakeNode {
	if nodeId < 0 || nodeId > s.nodeMax {
		nodeId = 0
	}

	return &SnowflakeNode{
		timestamp: 0,
		nodeId:    nodeId,
		seq:       0,
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取唯一Id
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func (n *SnowflakeNode) Id() SnowflakeId {
	n.mu.Lock()
	defer n.mu.Unlock()

	//纳秒转换成毫秒
	now := time.Now().UnixNano() / 1000000

	if n.timestamp == now {
		n.seq = (n.seq + 1) & snowflake.seqMask

		if n.seq == 0 {
			for now <= n.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		n.seq = 0
	}

	n.timestamp = now

	id := SnowflakeId((now-snowflake.Epoch)<<snowflake.timeShift |
		(n.nodeId << snowflake.nodeShift) |
		(n.seq),
	)

	return id
}

func (f SnowflakeId) Timestamp() int64 {
	return (int64(f) >> snowflake.timeShift) + snowflake.Epoch
}

func (f SnowflakeId) Node() int64 {
	return int64(f) & snowflake.nodeMask >> snowflake.nodeShift
}

func (f SnowflakeId) Seq() int64 {
	return int64(f) & snowflake.seqMask
}

func (f SnowflakeId) Int64() int64 {
	return int64(f)
}

func (f SnowflakeId) String() string {
	return strconv.FormatInt(int64(f), 10)
}

func (f SnowflakeId) Bytes() []byte {
	return []byte(f.String())
}

func (f SnowflakeId) IntBytes() [8]byte {
	var b [8]byte
	binary.BigEndian.PutUint64(b[:], uint64(f))
	return b
}

func (f SnowflakeId) Base2() string {
	return strconv.FormatInt(int64(f), 2)
}

func (f SnowflakeId) Base64() string {
	return base64.StdEncoding.EncodeToString(f.Bytes())
}
