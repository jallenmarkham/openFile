package openFile

import (
	"fmt"
	"strings"

	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
)

func Openfile(type_name string, ext []string) string {
	ActivateTheme("azure light")

	files := GetOpenFile(Multiple(false), Filetypes([]FileType{{TypeName: type_name, Extensions: ext, MacType: ""}}))
	fmt.Println(files)
	fname := strings.Join(files, " ")
	return fname
}
