package textextractor

type FolderExtractor struct {
	TextExtractor
	name  string
	files []string
}
