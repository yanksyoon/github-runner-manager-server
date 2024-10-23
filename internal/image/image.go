package image

import "github.com/charlie4284/github-runner-manager-server/internal/job"

type Manager struct {
	imageIDToImage   map[string]*job.Image
	imageNameToImage map[string]*job.Image
}

func New() (*Manager, error) {
	m := &Manager{}
	if err := m.loadImages(); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Manager) loadImages() error {
	return nil
}

func (m *Manager) FindImage(imageIDOrName string) (*job.Image, error) {
	if image, ok := m.imageIDToImage[imageIDOrName]; ok {
		return image, nil
	}
	if image, ok := m.imageNameToImage[imageIDOrName]; ok {
		return image, nil
	}
	return m.searchImage(imageIDOrName)
}

func (m *Manager) searchImage(imageIDOrName string) (*job.Image, error) {
	return &job.Image{}, nil
}
