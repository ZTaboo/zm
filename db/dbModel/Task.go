package dbModel

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	Name            string `gorm:"type:varchar(255);not null"` // 任务名
	Path            string `gorm:"type:varchar(255);not null"` // 执行路径
	ExecutableFiles string `gorm:"type:varchar(255);not null"` // 执行文件
	Port            int    `gorm:"type:int;not null"`          // 任务端口
	Pid             int    // 任务进程id
}
