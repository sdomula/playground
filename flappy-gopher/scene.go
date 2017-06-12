package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	img "github.com/veandco/go-sdl2/sdl_image"
)

type scene struct {
	bg   *sdl.Texture
	bird *sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene, error) {
	bg, err := img.LoadTexture(r, "res/imgs/background.png")
	if err != nil {
		return nil, fmt.Errorf("could not load background image: %v", err)
	}

	bird, err := img.LoadTexture(r, "res/imgs/bird_frame_1.png")
	if err != nil {
		return nil, fmt.Errorf("could not load bird image_frame_1: %v", err)
	}

	return &scene{bg: bg, bird: bird}, nil
}

func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()

	if err := r.Copy(s.bg, nil, nil); err != nil {
		return fmt.Errorf("could not copy background: %v", err)
	}

	rect := &sdl.Rect{X: 10, Y: 300 - 43/2, W: 50, H: 43}
	if err := r.Copy(s.bird, nil, rect); err != nil {
		return fmt.Errorf("could not copy bird: %v", err)
	}

	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()

	s.bird.Destroy()
}
