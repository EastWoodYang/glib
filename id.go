package glib

/* ================================================================================
 * snowflake id
 * qq group: 582452342
 * email   : 2091938785@qq.com
 * author  : 美丽的地球啊 - mliu
 * ================================================================================ */

var (
	snowflakeNodes []*SnowflakeNode
)

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 初始化节点集合
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func init() {
	snowflakeNodes = make([]*SnowflakeNode, 0, 16)
	for i := 0; i < 16; i++ {
		snowflakeNodes = append(snowflakeNodes, NewSnowflake().Node(int64(i)))
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取唯一id
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetId(nodeId int) int64 {
	if nodeId < 0 || nodeId > 16 {
		nodeId = 0
	}

	uniquId := snowflakeNodes[nodeId].Id()

	return uniquId.Int64()
}
