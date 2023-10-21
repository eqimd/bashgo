package splitter

import "strings"

var _ PipelineSplitter = &PipelineSplitterImpl{}

type PipelineSplitterImpl struct {
}

func (splitter *PipelineSplitterImpl) Split(cmd string) []string {
	curQuotes := byte(' ')
	splitCmds := []string{}
	curSplitPos := 0

	for i := 0; i < len(cmd); i++ {
		sym := cmd[i]

		switch sym {
		case '\'':
			if curQuotes == '\'' {
				curQuotes = ' '
			} else if curQuotes == ' ' {
				curQuotes = '\''
			}
		case '"':
			if curQuotes == '"' {
				curQuotes = ' '
			} else if curQuotes == ' ' {
				curQuotes = '"'
			}
		}

		if curQuotes != ' ' {
			continue
		}

		if sym == '|' {
			addCmd := cmd[curSplitPos:i]
			addCmd = strings.TrimSpace(addCmd)
			splitCmds = append(splitCmds, addCmd)

			curSplitPos = i + 1
		}
	}

	if curSplitPos == len(cmd) {
		splitCmds = append(splitCmds, "")
	} else {
		addCmd := cmd[curSplitPos:]
		addCmd = strings.TrimSpace(addCmd)
		splitCmds = append(splitCmds, addCmd)
	}

	return splitCmds
}

func NewPipelineSplitterImpl() *PipelineSplitterImpl {
	return &PipelineSplitterImpl{}
}
