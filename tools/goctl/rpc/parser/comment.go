package parser

import "github.com/emicklei/proto"

// GetComment returns content with prefix //
func GetComment(comment *proto.Comment) string {
	if comment == nil {
		return ""
	}

	if len(comment.Lines) > 1 {
		return "// " + comment.Lines[1]
	} else {
		return "// " + comment.Message()
	}
}
