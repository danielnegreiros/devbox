package ports

type CommandOutputDTO struct {
	OutMessage string
	ErrMessage string
	Success    bool
	Error      error
}

func NewCommandOutputDTO(outMessage string, errMessage string, success bool, err error) CommandOutputDTO {
	return CommandOutputDTO{
		OutMessage: outMessage,
		ErrMessage: errMessage,
		Success:    success,
		Error:      err,
	}
}

type InfraRepo interface {
	Close()
	ExecuteCommand(command string) CommandOutputDTO
}
