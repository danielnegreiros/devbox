package munix

import (
	"log"
	"testing"

	"github.com/danielnegreiros/go-proxmox-cli/internal/infrastructure/adapters"
)

func TestChangeWithRegex(t *testing.T) {
	path := "/tmp/hosts"
	regex := "^192.168.0.106.*"
	line := "192.168.0.106\tmy-windows-3"
	state := "present"
	owner := ""
	group := ""
	mode := ""
	validate := ""
	shouldCreate := true

	adapter := adapters.NewLocalAdapter()
	lineInfile := NewLineInFile(path, regex, line, state, owner, group, mode, validate, shouldCreate)
	result := lineInfile.Execute(adapter)
	log.Println(result)
}

// func TestAppendWithRegex(t *testing.T) {
// 	path := "/tmp/hosts"
// 	regex := "^sometest.*"
// 	line := "someconf=3"
// 	state := "present"
// 	owner := ""
// 	group := ""
// 	mode := ""
// 	validate := ""
// 	shouldCreate := true

// 	lineInfile := NewLineInFile(path, regex, line, state, owner, group, mode, validate, shouldCreate)
// 	result := lineInfile.Execute()
// 	log.Println(result)
// }

// func TestNoFileShouldCreateFalse(t *testing.T) {
// 	path := "/tmp/hosts2"
// 	regex := "^sometest.*"
// 	line := "someconf=3"
// 	state := "present"
// 	owner := ""
// 	group := ""
// 	mode := ""
// 	validate := ""
// 	shouldCreate := false

// 	lineInfile := NewLineInFile(path, regex, line, state, owner, group, mode, validate, shouldCreate)
// 	result := lineInfile.Execute()
// 	log.Println(result)
// }

// func TestNoFileShouldCreateTrue(t *testing.T) {
// 	path := "/tmp/hosts2"
// 	regex := "^sometest.*"
// 	line := "someconf=3"
// 	state := "present"
// 	owner := ""
// 	group := ""
// 	mode := ""
// 	validate := ""
// 	shouldCreate := true

// 	lineInfile := NewLineInFile(path, regex, line, state, owner, group, mode, validate, shouldCreate)
// 	result := lineInfile.Execute()
// 	log.Println(result)
// }
