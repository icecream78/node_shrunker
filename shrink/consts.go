package shrink

var DefaultRemoveDirNames []string = []string{
	"test",
	"tests",
	"example",
	"examples",
}

var DefaultRemoveFileNames []string = []string{
	"package.json",
}

var DefaultRemoveFileExt []string = []string{
	".ts",
	".d.ts",
	".coffee",
}

var NodeModulesDirname string = "node_modules"
var (
	progressChar string = "├───"
	lastChar     string = "└───"
	tabChar      string = "	"
)
