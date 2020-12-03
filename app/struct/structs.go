package _struct

type UserSession struct {
	Account			string 	//账户id
}

type AccountInfo struct {
	Account			string `json:"account"`			//账户id
	Password		string `json:"password"`		//账户密码
	Group  			string `json:"group"`			//账户组
	Level			string `json:"level"`			//账户级别，0：管理员账户 1：普通账户
}

type EntryInfo struct {
	Company 		string `json:"company"`			//公司
	Jobs 			string `json:"jobs"`			//岗位
	Working			string `json:"working"`			//在岗人员
	Leader			string `json:"leader"`			//整改责任人
	Date			string `json:"date"`			//整改期限
	Problem 		string `json:"problem"`			//是否存在问题
	Type 			string `json:"type"`			//问题类型
	Account         string `json:"account"`			//操作人
}
