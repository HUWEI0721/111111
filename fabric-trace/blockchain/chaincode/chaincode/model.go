package chaincode

/*
定义用户结构体
用户ID
用户类型
实名认证信息哈希,包括用户注册的姓名、身份证号、手机号、注册平台同意协议签名的哈希
茅台酒列表
*/
type User struct {
	UserID       string     `json:"userID"`
	UserType     string     `json:"userType"`
	RealInfoHash string     `json:"realInfoHash"`
	MoutaiList  []*Moutai   `json:"moutaiList"`
}

/*
定义茅台酒溯源结构体
溯源码
酿酒师输入
酿造厂输入
物流司机信息输入
销售点信息输入
图片字段
*/
type Moutai struct {
	Traceability_code string          `json:"traceability_code"`
	Brewer_input   Brewer_input 	  `json:"brewer_input"`
	Distillery_input     Distillery_input   `json:"distillery_input"`
	Driver_input      Driver_input    `json:"driver_input"`
	Shop_input        Shop_input      `json:"shop_input"`
	Image             string          `json:"image"` // 添加的图片字段
}

// HistoryQueryResult 结构体用于处理历史查询结果
type HistoryQueryResult struct {
	Record    *Moutai `json:"record"`
	TxId      string   `json:"txId"`
	Timestamp string   `json:"timestamp"`
	IsDelete  bool     `json:"isDelete"`
}

/*
酿酒师输入
茅台酒的溯源码，一物一码，主打高端市场（自动生成）
酿酒师信息
酿造开始时间
酿造完成时间
酿酒师名称
*/
type Brewer_input struct {
	Moutai_name    string `json:"moutai_name"`
	Moutai_origin  string `json:"moutai_origin"`
	Moutai_start   string `json:"moutai_start"`
	Moutai_end 	   string `json:"moutai_end"`
	Brewer_name    string `json:"brewer_name"`
	Brewer_Txid    string `json:"brewer_txid"`
	Brewer_Timestamp      string `json:"brewer_timestamp"`
	Image          string `json:"image"` // 添加的图片字段
}

/*
酿造厂输入
茅台酒
酿造批次
出厂时间（可以防止黑心商家修改时间）
酿造厂名称与地址
物流司机联系电话
*/
type Distillery_input struct {
	Moutai_productName     string `json:"moutai_productName"`
	Moutai_batch string `json:"moutai_batch"`
	Moutai_productionTime  string `json:"moutai_productionTime"`
	Distillery_name     string `json:"distillery_name"`
	Distillery_phone   string `json:"distillery_phone"`
	Distillery_Txid            string `json:"distillery_txid"`
	Distillery_Timestamp       string `json:"distillery_timestamp"`
	Image               string `json:"image"` // 添加的图片字段
}

/*
物流司机信息
姓名
车厢车厢温度
电话
运输车辆车牌号
运输与车厢记录
*/
type Driver_input struct {
	Driver_name      string `json:"driver_name"`
	Driver_age       string `json:"driver_temperature"`
	Driver_phone     string `json:"driver_phone"`
	Driver_carNumber string `json:"driver_carNumber"`
	Driver_transport string `json:"driver_transport"`
	Driver_Txid      string `json:"driver_txid"`
	Driver_Timestamp string `json:"driver_timestamp"`
	Image        string `json:"image"` // 添加的图片字段
}

/*
销售点信息
存入时间
销售时间
销售点名称
销售点位置
销售点电话
*/
type Shop_input struct {
	Shop_storeTime   string `json:"shop_storeTime"`
	Shop_sellTime    string `json:"shop_sellTime"`
	Shop_shopName    string `json:"shop_shopName"`
	Shop_shopAddress string `json:"shop_shopAddress"`
	Shop_shopPhone   string `json:"shop_shopPhone"`
	Shop_Txid        string `json:"shop_txid"`
	Shop_Timestamp   string `json:"shop_timestamp"`
	Image          string `json:"image"` // 添加的图片字段
}
