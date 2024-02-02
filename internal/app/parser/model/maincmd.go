package model

type MainCommandParserFn func([]string) (SubCmdParserFn, UseCaseCliHandlerFn)
type MainCmd struct {
	Name          string
	Description   string
	GroupParserFn MainCommandParserFn
}

var AvailGroupCmd []MainCmd
