package grep

import (
	"regexp"
	"strings"
)

var KnownExt = []string{".doc", ".docx", ".eml", ".log", ".msg", ".odt", ".pages", ".rtf", ".tex", ".txt", ".bin", ".dat", ".csv", ".tar", ".ppt", ".pptx", ".xml", ".html", ".js", ".css", ".zip", ".wav", ".mp3", ".ogg", ".mp4", ".mpg", ".mov", ".blend", ".bmp", ".gif", ".png", ".jpg", ".jpeg", ".epub", ".pub", ".xls", ".xlsx", ".xlr", ".odb", ".sql", ".sqlite", ".db", ".mdb", ".bat", ".jar", ".apk", ".app", ".pkg", ".json", ".php", ".yml", ".asc", ".pgp", ".rar", ".tar.gz", ".7z", ".rpm", ".iso", ".img", ".c", ".py", ".config", ".lua", ".swift", ".java", ".bak", ".tmp", ".part", ".pdf"}

func ClearHrefs(hrefs []string) []string {
	cleared := regexp.MustCompile("\"(.*?)\"")
	for i := range hrefs {
		hrefs[i] = cleared.FindString(hrefs[i])
		hrefs[i] = strings.Replace(hrefs[i], "\"", "", -1)
	}
	return hrefs
}

func IsKnownFile(url string) bool {
	for i := range KnownExt {
		pattern := regexp.MustCompile("([a-zA-Z0-9\\s_\\.\\-\\(\\):])+(" + regexp.QuoteMeta(KnownExt[i]) + ")$")
		if pattern.MatchString(url) {
			return true
		}
	}
	return false
}

/*Grep all internal hrefs from HTML body*/
func InternalPaths(body string) []string {
	pattern := regexp.MustCompile("href=\"[^http.*](.*?)\"")
	dirs := pattern.FindAllString(body, -1)
	dirs = ClearHrefs(dirs)
	return dirs
}

/*Grep all internal hrefs from HTML body with a specified extension*/
func InternalDocumentPaths(body string, ext ...string) []string {
	docs := []string{}
	if ext != nil {
		for i := range ext {
			pattern := regexp.MustCompile("href=\"[^http.*](.*?)" + regexp.QuoteMeta(ext[i]) + "\"")
			docs = append(docs, pattern.FindAllString(body, -1)...)
		}
	} else {
		for i := range KnownExt {
			pattern := regexp.MustCompile("href=\"[^http.*](.*?)" + regexp.QuoteMeta(KnownExt[i]) + "\"")
			docs = append(docs, pattern.FindAllString(body, -1)...)
		}
	}
	docs = ClearHrefs(docs)
	return docs
}

//TODO: add searching for userspecified extension
func FileNameFromPath(path string) string {
	pattern := regexp.MustCompile("[^/]+$")
	name := pattern.FindString(path)
	return name
}
