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
	snowflakeNodes = make([]*SnowflakeNode, 0, 1024)
	for i := 0; i < 1024; i++ {
		snowflakeNodes = append(snowflakeNodes, NewSnowflake().GetNode(int64(i)))
	}
}

/* ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 * 获取唯一id
 * ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ */
func GetId(workerNodeId int64, isReset bool) int64 {
	if workerNodeId < 0 || workerNodeId > 1024 {
		workerNodeId = 0
	}

	uniquId := snowflakeNodes[workerNodeId].GetId(isReset)

	return uniquId.Int64()
}
