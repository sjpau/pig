# pig
Pig is a simple web crawler to download href elements in bulk.
### Usage
-url, --url [string] \
	Specify absolute URL to visit. \
-ask, --ask [bool] \
	Ask before downloading a file. \
-ext, --ext [string] \
	Look only for files with specified extension. \
-lazy, --lazy [int]	\
	Crawl with delay. \
-target, --target [string] \
	Specify directory for downloading. \
-v, --v [bool] \
	Verbose mode. \
### Example
	pig --url https://example.com --v=true --ask=false --timer 30 --lazy 2 --target /home/user/Downloads
### Installation
`go install github.com/sjpau/pig@latest`
