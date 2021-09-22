package utils

import (
	"fmt"
	sf "github.com/bwmarrin/snowflake"
	"time"
)

var node *sf.Node

func Init(startTime string, machineID int64) (err error) { //传入起始时间，机器的id值
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		return
	}
	sf.Epoch = st.UnixNano() / 1000000 //毫秒值
	node, err = sf.NewNode(machineID)
	return
}

//获取id值
func GetId() int64 {
	return node.Generate().Int64() //可获取string
}

func GetCommunityID() *sf.Node {
	err := Init("2020-07-01", 2)
	if err != nil {
	}
	return node
}

func use() {
	err := Init("2020-07-01", 2)
	if err != nil {
		fmt.Println("init failed ,err:", err)
		return
	}
	//获取id值
	id := GetId()
	fmt.Println(id)
}
