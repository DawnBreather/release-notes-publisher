package html

import (
  "github.com/tdewolff/minify/v2"
  minifyhtml "github.com/tdewolff/minify/v2/html"
  . "html_to_xhtml_converter/config"
  "log"
)

func ProcessHTMLContent(htmlContent string, config Config) string {
  if config.ShouldMinify {
    return minifyHTML(htmlContent)
  }
  return htmlContent
}

func minifyHTML(htmlContent string) string {
  m := minify.New()
  m.AddFunc("text/html", minifyhtml.Minify)
  minifiedHTML, err := m.String("text/html", htmlContent)
  if err != nil {
    log.Fatalf("Failed to minify: %v", err)
  }
  return minifiedHTML
}
