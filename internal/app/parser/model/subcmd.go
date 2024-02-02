package model

type UseCaseCliHandlerFn func(map[string]string)
type SubCmdParserFn func([]string) map[string]string
type SubCmd struct {
	Parent      string
	Name        string
	Description string
	ParseFunc   SubCmdParserFn
	ExecFunc    UseCaseCliHandlerFn
}

var AvailSubCmds []SubCmd
