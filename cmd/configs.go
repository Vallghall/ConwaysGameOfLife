package cmd

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"gopkg.in/yaml.v3"
	"image/png"
	"log"
)

//go:embed configs/*
var configs embed.FS

type Configs struct {
	*Cell   `yaml:"cell"`
	*Window `yaml:"window"`
}

type Cell struct {
	Size            int    `yaml:"size"`
	DeadSpritePath  string `yaml:"dead-sprite"`
	AliveSpritePath string `yaml:"alive-sprite"`
	deadSprite      *ebiten.Image
	aliveSprite     *ebiten.Image
}

func (c *Cell) DeadSprite() *ebiten.Image {
	return c.deadSprite
}

func (c *Cell) AliveSprite() *ebiten.Image {
	return c.aliveSprite
}

type Window struct {
	Title  string `yaml:"title"`
	Width  int    `yaml:"width"`
	Height int    `yaml:"height"`
}

func NewConfigs() *Configs {
	c := &Configs{}

	data, err := configs.ReadFile("configs/configs.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		log.Fatalln(err)
	}

	deadI, err := configs.Open(c.DeadSpritePath)
	if err != nil {
		log.Fatalln(err)
	}

	aliveI, err := configs.Open(c.AliveSpritePath)
	if err != nil {
		log.Fatalln(err)
	}

	dead, err := png.Decode(deadI)
	if err != nil {
		log.Fatalln(err)
	}

	alive, err := png.Decode(aliveI)
	if err != nil {
		log.Fatalln(err)
	}

	c.deadSprite = ebiten.NewImageFromImage(dead)
	c.aliveSprite = ebiten.NewImageFromImage(alive)

	return c
}
