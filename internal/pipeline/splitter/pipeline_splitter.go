package splitter

type PipelineSplitter interface {
	Split(cmd string) []string
}
