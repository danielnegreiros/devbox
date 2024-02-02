package adapters

import (
	"log"
	"testing"
)

func TestLinuxLocal(t *testing.T) {

	adapter := NewLocalAdapter()
	command := "ls -lhtr"

	outputDTO := adapter.ExecuteCommand(command)

	if !outputDTO.Success {
		log.Panic(outputDTO.ErrMessage)
	}

	if outputDTO.Error != nil{
		t.Error(outputDTO.Error)
	}

}
