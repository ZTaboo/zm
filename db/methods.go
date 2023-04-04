package db

import (
	"ZM/db/dbModel"
	"ZM/model"
	"ZM/utils"
	"log"
	"strconv"
)

// Add 添加任务
func Add(path string, name string, port int) error {
	if err := Db.Create(&dbModel.Task{
		Name: name,
		Path: path,
		Port: port,
	}).Error; err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetAllTask 获取task表全部信息
func GetAllTask() ([]model.StatusModel, error) {
	var status []model.StatusModel
	var tasks []dbModel.Task
	if err := Db.Find(&tasks).Error; err != nil {
		log.Println(err)
		return nil, err
	}
	for _, item := range tasks {
		stat := utils.PortCheck(strconv.Itoa(item.Port))
		status = append(status, model.StatusModel{
			Name:   item.Name,
			Path:   item.Path,
			Port:   item.Port,
			Status: stat,
		})
	}
	return status, nil
}

// TaskExist 查询任务是否存在
func TaskExist(name string) (dbModel.Task, error) {
	var task dbModel.Task
	if err := Db.Where("name = ?", name).First(&task).Error; err != nil {
		return dbModel.Task{}, err
	}
	// 判断端口是否占用

	return task, nil
}

func UpdatePid(taskName string, pid int) {
	if err := Db.Model(&dbModel.Task{}).Where("name = ?", taskName).Update("pid", pid).Error; err != nil {
		log.Println(err)
	}
}

func StopTask(name string) (int, error) {
	//	查找pid
	var task dbModel.Task
	if err := Db.Where("name = ?", name).First(&task).Error; err != nil {
		log.Println(err)
		return 0, err
	}
	return task.Pid, nil
}
