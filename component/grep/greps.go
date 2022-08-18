package grep

import (
	"regexp"
)

var KnownExt = []string{".doc", ".docx", ".eml", ".log", ".msg", ".odt", ".pages", ".rtf", ".tex", ".txt", ".bin", ".dat", ".csv", ".tar", ".ppt", ".pptx", ".xml", ".html", ".js", ".css", ".zip", ".wav", ".mp3", ".ogg", ".mp4", ".mpg", ".mov", ".blend", ".bmp", ".gif", ".png", ".jpg", ".jpeg", ".epub", ".pub", ".xls", ".xlsx", ".xlr", ".odb", ".sql", ".sqlite", ".db", ".mdb", ".bat", ".jar", ".apk", ".app", ".pkg", ".json", ".php", ".yml", ".asc", ".pgp", ".rar", ".tar.gz", ".7z", ".rpm", ".iso", ".img", ".c", ".py", ".config", ".lua", ".swift", ".java", ".bak", ".tmp", ".part", ".pdf"}

func IsFileFormat(url string, ext ...string) bool {
	if ext == nil {
		for i := range KnownExt {
			pattern := regexp.MustCompile("([a-zA-Z0-9\\s_\\.\\-\\(\\):])+(" + regexp.QuoteMeta(KnownExt[i]) + ")$")
			if pattern.MatchString(url) {
				return true
			}
		}
	} else {
		for i := range ext {
			pattern := regexp.MustCompile("([a-zA-Z0-9\\s_\\.\\-\\(\\):])+(" + regexp.QuoteMeta(ext[i]) + ")$")
			if pattern.MatchString(url) {
				return true
			}
		}
	}
	return false
}

func FileNameFromPath(path string) string {
	pattern := regexp.MustCompile("[^/]+$")
	name := pattern.FindString(path)
	return name
}

func DomainNameFromURL(url string) string {
	pattern := regexp.MustCompile("[^(https?\\:\\/\\/)].*[^(\\/*)]")
	name := pattern.FindString(url)
	return name
}
