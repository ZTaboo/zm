package db

import (
	"ZM/db/dbModel"
	"ZM/model"
	"log"
	"net"
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
		var stat bool
		listen, err := net.Listen("tcp4", ":"+strconv.Itoa(item.Port))
		if err != nil {
			stat = true
		} else {
			stat = false
			err = listen.Close()
			if err != nil {
				log.Println("Listen close failed: ", err.Error())
				return nil, err
			}
		}

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
	return task, nil
}
