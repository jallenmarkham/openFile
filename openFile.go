package openFile

import (
	"fmt"

	. "modernc.org/tk9.0"
	_ "modernc.org/tk9.0/themes/azure"
)

func openfile(type_name string, ext []string) []string {
	ActivateTheme("azure light")
	// menubtn := App.Menubutton(Width(10), Height(1), Txt(`Test`))
	// Pack(Label(Image(NewPhoto(Data(gopher)))),
	// 	TExit(),
	// 	menubtn,
	// 	Padx("1m"), Pady("2m"), Ipadx("1m"), Ipady("1m"))
	// App.Center().Wait()
	files := GetOpenFile(Filetypes([]FileType{{TypeName: "Go files", Extensions: []string{".go"}, MacType: ""}}))
	fmt.Println(files)
	return files
}
