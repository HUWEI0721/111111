package chaincode

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// 定义合约结构体
type SmartContract struct {
	contractapi.Contract
}

// 注册用户
func (s *SmartContract) RegisterUser(ctx contractapi.TransactionContextInterface, userID string, userType string, realInfoHash string) error {
	user := User{
		UserID:       userID,
		UserType:     userType,
		RealInfoHash: realInfoHash,
		MoutaiList:  []*Moutai{},
	}
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 茅台酒上链，传入用户ID、茅台酒上链信息和图片信息
func (s *SmartContract) Uplink(ctx contractapi.TransactionContextInterface, userID string, traceability_code string, arg1 string, arg2 string, arg3 string, arg4 string, arg5 string, arg6 string) (string, error) {
	// 获取用户类型
	userType, err := s.GetUserType(ctx, userID)
	if err != nil {
		return "", fmt.Errorf("failed to get user type: %v", err)
	}

	// 通过溯源码获取茅台酒的上链信息
	MoutaiAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	// 将茅台酒的信息转换为Moutai结构体
	var seafood Moutai
	if MoutaiAsBytes != nil {
		err = json.Unmarshal(MoutaiAsBytes, &seafood)
		if err != nil {
			return "", fmt.Errorf("failed to unmarshal seafood: %v", err)
		}
	}

	// 获取时间戳,修正时区
	txtime, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return "", fmt.Errorf("failed to read TxTimestamp: %v", err)
	}
	timeLocation, _ := time.LoadLocation("Asia/Shanghai") // 选择你所在的时区
	timeStr := time.Unix(txtime.Seconds, 0).In(timeLocation).Format("2006-01-02 15:04:05")

	// 获取交易ID
	txid := ctx.GetStub().GetTxID()
	// 给茅台酒信息中加上溯源码
	seafood.Traceability_code = traceability_code
	// 设置图片字段
	seafood.Image = arg6 // 假设 arg6 是 Image 字段

	// 不同用户类型的上链的参数不一致
	switch userType {
	// 渔民
	case "渔民":
		// 将传入的茅台酒上链信息转换为Brewer_input结构体
		seafood.Brewer_input.Moutai_name = arg1
		seafood.Brewer_input.Moutai_origin = arg2
		seafood.Brewer_input.Moutai_start = arg3
		seafood.Brewer_input.Moutai_end = arg4
		seafood.Brewer_input.Brewer_name = arg5
		seafood.Brewer_input.Brewer_Txid = txid
		seafood.Brewer_input.Brewer_Timestamp = timeStr
		seafood.Brewer_input.Image = arg6 // 设置子结构体的 Image 字段
	// 酿造厂
	case "酿造厂":
		// 将传入的茅台酒上链信息转换为Distillery_input结构体
		seafood.Distillery_input.Moutai_productName = arg1
		seafood.Distillery_input.Moutai_batch = arg2
		seafood.Distillery_input.Moutai_productionTime = arg3
		seafood.Distillery_input.Distillery_name = arg4
		seafood.Distillery_input.Distillery_phone = arg5
		seafood.Distillery_input.Distillery_Txid = txid
		seafood.Distillery_input.Distillery_Timestamp = timeStr
		seafood.Distillery_input.Image = arg6 // 设置子结构体的 Image 字段
	// 物流司机信息
	case "物流司机信息":
		// 将传入的茅台酒上链信息转换为Driver_input结构体
		seafood.Driver_input.Driver_name = arg1
		seafood.Driver_input.Driver_age = arg2
		seafood.Driver_input.Driver_phone = arg3
		seafood.Driver_input.Driver_carNumber = arg4
		seafood.Driver_input.Driver_transport = arg5
		seafood.Driver_input.Driver_Txid = txid
		seafood.Driver_input.Driver_Timestamp = timeStr
		seafood.Driver_input.Image = arg6 // 设置子结构体的 Image 字段
	// 销售终端
	case "销售终端":
		// 将传入的茅台酒上链信息转换为Shop_input结构体
		seafood.Shop_input.Shop_storeTime = arg1
		seafood.Shop_input.Shop_sellTime = arg2
		seafood.Shop_input.Shop_shopName = arg3
		seafood.Shop_input.Shop_shopAddress = arg4
		seafood.Shop_input.Shop_shopPhone = arg5
		seafood.Shop_input.Shop_Txid = txid
		seafood.Shop_input.Shop_Timestamp = timeStr
		seafood.Shop_input.Image = arg6 // 设置子结构体的 Image 字段
	}

	// 将茅台酒的信息转换为json格式
	seafoodAsBytes, err := json.Marshal(seafood)
	if err != nil {
		return "", fmt.Errorf("failed to marshal seafood: %v", err)
	}
	// 将茅台酒的信息存入区块链
	err = ctx.GetStub().PutState(traceability_code, seafoodAsBytes)
	if err != nil {
		return "", fmt.Errorf("failed to put seafood: %v", err)
	}
	// 将茅台酒存入用户的茅台酒列表
	err = s.AddMoutai(ctx, userID, &seafood)
	if err != nil {
		return "", fmt.Errorf("failed to add seafood to user: %v", err)
	}

	return txid, nil
}

// 添加茅台酒到用户的茅台酒列表
func (s *SmartContract) AddMoutai(ctx contractapi.TransactionContextInterface, userID string, seafood *Moutai) error {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return err
	}
	user.MoutaiList = append(user.MoutaiList, seafood)
	userAsBytes, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(userID, userAsBytes)
	if err != nil {
		return err
	}
	return nil
}

// 获取用户类型
func (s *SmartContract) GetUserType(ctx contractapi.TransactionContextInterface, userID string) (string, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return "", fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return "", fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return "", err
	}
	return user.UserType, nil
}

// 获取用户信息
func (s *SmartContract) GetUserInfo(ctx contractapi.TransactionContextInterface, userID string) (*User, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return &User{}, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return &User{}, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return &User{}, err
	}
	return &user, nil
}

// 获取茅台酒的上链信息
func (s *SmartContract) GetMoutaiInfo(ctx contractapi.TransactionContextInterface, traceability_code string) (*Moutai, error) {
	MoutaiAsBytes, err := ctx.GetStub().GetState(traceability_code)
	if err != nil {
		return &Moutai{}, fmt.Errorf("failed to read from world state: %v", err)
	}

	// 将返回的结果转换为Moutai结构体
	var seafood Moutai
	if MoutaiAsBytes != nil {
		err = json.Unmarshal(MoutaiAsBytes, &seafood)
		if err != nil {
			return &Moutai{}, fmt.Errorf("failed to unmarshal seafood: %v", err)
		}
		if seafood.Traceability_code != "" {
			return &seafood, nil
		}
	}
	return &Moutai{}, fmt.Errorf("the seafood %s does not exist", traceability_code)
}

// 获取用户的茅台酒ID列表
func (s *SmartContract) GetMoutaiList(ctx contractapi.TransactionContextInterface, userID string) ([]*Moutai, error) {
	userBytes, err := ctx.GetStub().GetState(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if userBytes == nil {
		return nil, fmt.Errorf("the user %s does not exist", userID)
	}
	// 将返回的结果转换为User结构体
	var user User
	err = json.Unmarshal(userBytes, &user)
	if err != nil {
		return nil, err
	}
	return user.MoutaiList, nil
}

// 获取所有的茅台酒信息
func (s *SmartContract) GetAllMoutaiInfo(ctx contractapi.TransactionContextInterface) ([]Moutai, error) {
	moutaiListAsBytes, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	defer moutaiListAsBytes.Close()
	var seafoods []Moutai
	for moutaiListAsBytes.HasNext() {
		queryResponse, err := moutaiListAsBytes.Next()
		if err != nil {
			return nil, err
		}
		var seafood Moutai
		err = json.Unmarshal(queryResponse.Value, &seafood)
		if err != nil {
			return nil, err
		}
		// 过滤非茅台酒的信息
		if seafood.Traceability_code != "" {
			seafoods = append(seafoods, seafood)
		}
	}
	return seafoods, nil
}

// 获取茅台酒上链历史
func (s *SmartContract) GetMoutaiHistory(ctx contractapi.TransactionContextInterface, traceability_code string) ([]HistoryQueryResult, error) {
	log.Printf("GetAssetHistory: ID %v", traceability_code)

	resultsIterator, err := ctx.GetStub().GetHistoryForKey(traceability_code)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var records []HistoryQueryResult
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		var seafood Moutai
		if len(response.Value) > 0 {
			err = json.Unmarshal(response.Value, &seafood)
			if err != nil {
				return nil, err
			}
		} else {
			seafood = Moutai{
				Traceability_code: traceability_code,
			}
		}

		timestamp, err := ptypes.Timestamp(response.Timestamp)
		if err != nil {
			return nil, err
		}
		// 指定目标时区
		targetLocation, err := time.LoadLocation("Asia/Shanghai")
		if err != nil {
			return nil, err
		}

		// 将时间戳转换到目标时区
		timestamp = timestamp.In(targetLocation)
		// 格式化时间戳为指定格式的字符串
		formattedTime := timestamp.Format("2024-12-01 22:11:00")

		record := HistoryQueryResult{
			TxId:      response.TxId,
			Timestamp: formattedTime,
			Record:    &seafood,
			IsDelete:  response.IsDelete,
		}
		records = append(records, record)
	}

	return records, nil
}
