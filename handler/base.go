package handler

import (
	"github.com/rumyantseva/vk-authorization-example/config"
	"gopkg.in/reform.v1"
)

type Handler struct {
	DB *reform.DB
	VK config.VKConf
}

func NewHandler(db *reform.DB, vk config.VKConf) *Handler {
	h := &Handler{DB: db, VK: vk}
	return h
}
