package markgen

import "github.com/russross/blackfriday"

//MdConverter : Global Converter
var MdConverter = NewMarkdownConverter()

//MarkdownConverter : class
type MarkdownConverter struct {
	convert func([]byte) []byte
}

//NewMarkdownConverter : Returns a new MarkdownConverter
func NewMarkdownConverter() *MarkdownConverter {
	return &MarkdownConverter{blackfriday.MarkdownCommon}
}

//UseBasic : Sets the Converter used
func (md *MarkdownConverter) UseBasic() {
	md.convert = blackfriday.MarkdownBasic
}

//Convert : Returns the converted html
func (md *MarkdownConverter) Convert(raw []byte) []byte {
	return md.convert(raw)
}
