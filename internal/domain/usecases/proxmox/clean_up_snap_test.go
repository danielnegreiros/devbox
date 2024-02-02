package proxmox

import (
	"log"
	"testing"
)

func initSnap() []SnapShotData {

	return []SnapShotData{
		{
			Name:     "permanent",
			SnapTime: 1702571265,
			Vmid:     1000,
		},
		{
			Name:     "current",
			SnapTime: 0,
			Vmid:     1000,
		},
		{
			Name:     "origin",
			SnapTime: 1701796340,
			Vmid:     102,
		},
	}
}

func TestCleanUpSnaps(t *testing.T) {

	snaps := initSnap()
	days := 2
	include := ".*"
	exclude := "per.*"

	resSnaps := getSnapsOlderThan(snaps, days)
	resSnaps = getIncludedOnlySnaps(resSnaps, include)
	resSnaps = getWithExcludedSnaps(resSnaps, exclude)

	log.Println(resSnaps)

}
