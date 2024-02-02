package adapters

// import (
// 	"log"
// 	"testing"
// )

// func TestLinux(t *testing.T) {

// 	adapter := NewLinuxAdapter("192.168.0.51", 22, "daniel", "xxxxxx")
// 	command := "ls -lhtr"

// 	outputDTO := adapter.ExecuteCommand(command)

// 	if !outputDTO.Success {
// 		log.Panic(outputDTO.ErrMessage)
// 	}

// 	if outputDTO.OutMessage != "abc" {
// 		t.Errorf("\n Error. Expected: %s, Found %s", "abc", outputDTO.OutMessage)
// 	}

// }
