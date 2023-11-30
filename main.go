package main

import (
	. "html_to_xhtml_converter/config"
	"html_to_xhtml_converter/html"
	. "html_to_xhtml_converter/input"
	. "html_to_xhtml_converter/output"
	"html_to_xhtml_converter/xhtml"
)

func main() {
  ConfigStorage = ParseFlags()
  htmlContent := ReadInputContent(ConfigStorage)
  finalHTML := html.ProcessHTMLContent(htmlContent, ConfigStorage)
  output := xhtml.ConvertHTMLToXHTML(finalHTML)
  WriteOutput(output, ConfigStorage)
}
