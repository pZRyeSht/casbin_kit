package config

type Conf struct {
	Server   Server   `json:"server" yaml:"server"`
	Database Database `json:"database" yaml:"database"`
}

type Server struct {
	RunMode      string `json:"run_mode" yaml:"run_mode"`           // 运行模式
	HttpPort     int    `json:"http_port" yaml:"http_port"`         // 服务端口
	ReadTimeout  int    `json:"read_timeout" yaml:"read_timeout"`   // 读取超时时间
	WriteTimeout int    `json:"write_timeout" yaml:"write_timeout"` // 写入超时时间
}

type Database struct {
	DBType      string `json:"db_type" yaml:"db_type"`             // 数据库类型
	Username    string `json:"username" yaml:"username"`           // 用户名
	Password    string `json:"password" yaml:"password"`           // 密码
	Host        string `json:"host" yaml:"host"`                   // 主机
	DBName      string `json:"db_name" yaml:"db_name"`             // 数据库名
	TablePrefix string `json:"table_prefix" yaml:"table_prefix"`   // 表前缀
	Charset     string `json:"charset" yaml:"charset"`             // 字符集
	ParseTime   bool   `json:"parse_time" yaml:"parse_time"`       // 解析
	MaxIdleTime int    `json:"max_idle_time" yaml:"max_idle_time"` // idle时间
	MaxOpenConn int    `json:"max_open_conn" yaml:"max_open_conn"` // 最大连接数
}
