package flavor

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/charlie4284/github-runner-manager-server/internal/job"
)

type Manager struct {
	flavors       map[string]*job.Flavor
	sortedFlavors []*job.Flavor
}

func New() (*Manager, error) {
	m := &Manager{
		flavors:       map[string]*job.Flavor{},
		sortedFlavors: []*job.Flavor{},
	}
	if err := m.loadFlavors(); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *Manager) loadFlavors() error {
	return nil
}

var ErrFlavorNotFoundError = errors.New("matching flavor not found")

func (m *Manager) FindFlavor(flavorName string) (*job.Flavor, error) {
	flavor, err := parseFlavor(flavorName)
	if err != nil {
		return nil, err
	}
	if foundFlavor, ok := m.flavors[convertFlavorToString(flavor)]; !ok {
		return nil, ErrFlavorNotFoundError
	} else {
		return foundFlavor, nil
	}
}

var ErrInvalidFlavor = errors.New("invalid flavor name")

func parseFlavor(flavorName string) (*job.Flavor, error) {
	parts := strings.Split(flavorName, "-")
	flavor := job.Flavor{}
	for _, part := range parts {
		if strings.Contains(part, "cpu") {
			disk, err := parseCPUPart(part)
			if err != nil {
				return nil, ErrInvalidFlavor
			}
			flavor.Disk = disk
		} else if strings.Contains(part, "ram") {
			ram, err := parseRAMPart(part)
			if err != nil {
				return nil, ErrInvalidFlavor
			}
			flavor.Ram = ram
		} else if strings.Contains(part, "disk") {
			disk, err := parseDiskPart(part)
			if err != nil {
				return nil, ErrInvalidFlavor
			}
			flavor.Ram = disk
		}
	}
	if flavor.Cores == 0 || flavor.Disk == 0 || flavor.Ram == 0 {
		return nil, ErrInvalidFlavor
	}
	return &flavor, nil
}

func parseCPUPart(part string) (int, error) {
	coreNum := strings.ReplaceAll(part, "cpu", "")
	return strconv.Atoi(coreNum)
}

func parseRAMPart(part string) (int, error) {
	ramNum := strings.ReplaceAll(part, "ram", "")
	return strconv.Atoi(ramNum)
}

func parseDiskPart(part string) (int, error) {
	diskNum := strings.ReplaceAll(part, "disk", "")
	return strconv.Atoi(diskNum)
}

func convertFlavorToString(flavor *job.Flavor) string {
	return fmt.Sprintf("cpu%d-ram%d-disk%d", flavor.Cores, flavor.Ram, flavor.Disk)
}
