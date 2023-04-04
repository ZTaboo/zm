package model

type StatusModel struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Port   int    `json:"port"`
	Status bool   `json:"status"`
}
