package _struct

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserSession struct {
	Account			string 	//账户id
}

type AccountInfo struct {
	ID    			string `json:"_id"`
	Account			string `json:"account"`			//账户id
	Password		string `json:"password"`		//账户密码
	Group  			string `json:"group"`			//账户组
	Level			int32 `json:"level"`			//账户级别，0：管理员账户 1：普通账户
}

type EntryInfo struct {
	ID         		primitive.ObjectID `bson:"_id"`
	Company 		string `json:"company"`			//公司
	Jobs 			string `json:"jobs"`			//岗位
	Working			string `json:"working"`			//在岗人员
	Leader			string `json:"leader"`			//整改责任人
	TimeLimit		string `json:"time_limit"`		//整改期限
	ExistProblem 	string `json:"exist_problem"`	//是否存在问题
	ProblemType 	string `json:"problem_type"`	//问题类型
}
