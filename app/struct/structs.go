package _struct

const (
	LevelEditor  = 0x001
	LevelChecker = 0x010
	LevelManager = 0x100
)

//Service
type TaskService struct {
	ID       string            `bson:"_id"`      //id
	Network  string            `bson:"network"`  //网络类型
	Address  string            `bson:"address"`  //地址
	Name     string            `bson:"name"`     //名称
	Path     map[string]string `bson:"path"`     //路径
	Enable   bool              `bson:"enable"`   //是否启用
	Reserved string            `bson:"reserved"` //预留
}

//HeartBeat
type HeartBeat struct {
	ID       string `bson:"_id"`      //id
	Network  string `bson:"network"`  //网络类型
	Address  string `bson:"address"`  //地址
	Reserved string `bson:"reserved"` //预留
}

//Manager
type TaskManagement struct {
	ID          string `bson:"_id"`         //联合id(唯一标识)
	ProjectID   string `bson:"project_id"`  //项目id
	InstanceID  string `bson:"instance_id"` //实例id
	TaskID      string `bson:"task_id"`     //任务id
	DataType    string `bson:"data_type"`   //数据类别
	JobType     string `bson:"job_type"`    //任务类别
	ToolName    string `bson:"tool_name"`   //工具名称
	Status      string `bson:"status"`      //任务状态
	CreateTime  int64  `bson:"time"`        //创建时间
	User        string `bson:"user"`        //用户
	Distributor string `bson:"distributor"` //分配者
	Checker     string `bson:"checker"`     //校验
	Group       string `bson:"group"`       //组
	Reserved    string `bson:"reserved"`    //预留
}

type TaskFromService struct {
	ID         string `bson:"_id"`
	ProjectID  string
	InstanceID string
	TaskID     string
	DataType   string
	JobType    string
	Status     string
	ToolName   string
	Req        interface{}
}

type ServiceTimer struct {
	Enable  bool
	Counter int
}

// Paging 分页结构
type Paging struct {
	Pos        int64         `json:"pos"`
	TotalCount int64         `json:"total_count"`
	Data       []interface{} `json:"data"`
}

type UserSession struct {
	ID string `json:"_id"` //用户代码
}

type UserInfo struct {
	ID    string `json:"_id"`
	User  string `json:"user"`  //用户代码
	Group string `json:"group"` //组代码
}
