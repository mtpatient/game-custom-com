package main

import (
	"context"
	"fmt"
	"game-custom-com/internal/model/entity"
	"game-custom-com/internal/service"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func TestGetSign(t *testing.T) {
	signature, err := service.Img().GetSignatures(context.Background(), 2)
	if err != nil {
		return
	}
	fmt.Println(signature)
}

func TestSaveImages(t *testing.T) {
	images := make([]entity.Image, 2)
	images[0] = entity.Image{
		Type: 0,
		Url:  "test1",
	}

	images[1] = entity.Image{
		Type: 1,
		Url:  "test2",
	}

	service.Img().Save(context.Background(), images)
}

func TestGetAvatar(t *testing.T) {
	avatar, err := service.Img().GetAllAvatar(context.Background())
	if err != nil {
		return
	}
	fmt.Println(avatar)
}

func TestUpdate(t *testing.T) {
	err := service.Img().Update(context.Background(), entity.Image{
		Type: 0,
		Id:   7,
		Url:  "https://game-custom-1312933264.cos.ap-guangzhou.myqcloud.com/img/2024-02-18/1a0j290irg0cz87m3lpj594740d6cl59",
		Name: "test",
	})
	if err != nil {
		g.Log(err.Error())
	}
	return
}
