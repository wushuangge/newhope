package _struct

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSession struct {
	Account			string 	//账户id
}

type AccountInfo struct {
	Id 				string `bson:"_id"`
	Account			string `bson:"account"`			//账户id
	Password		string `bson:"password"`		//账户密码
	Group  			string `bson:"group"`			//账户组
	Level			string `bson:"level"`			//账户级别，0：管理员账户 1：普通账户
}

type EntryInfo struct {
	Id				primitive.ObjectID `bson:"_id"`
	Company 		string `bson:"company"`			//公司
	Jobs 			string `bson:"jobs"`			//岗位
	Working			string `bson:"working"`			//在岗人员
	Leader			string `bson:"leader"`			//整改责任人
	Date			string `bson:"date"`			//整改期限
	Problem 		string `bson:"problem"`			//是否存在问题
	Type 			string `bson:"type"`			//问题类型
	Account         string `bson:"account"`			//操作人
}

type TermsInfo struct {
	Id				primitive.ObjectID `bson:"_id"`
	Terms			string `bson:"terms"`			//违规条款
}

type ResponseMessage struct {
	ErrCode 			int							//错误id
	ErrMessage			string						//错误信息
	Body    			string						//消息体
}
