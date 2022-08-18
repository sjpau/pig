package main

import "flag"

var (
	FlagAsk     bool
	FlagURL     string
	FlagExt     string
	FlagDest    string
	FlagLazy    int
	FlagVerbose bool
)

func InitCmdOptions() {
	flag.BoolVar(&FlagAsk, "ask", true, "Ask before downloading.")
	flag.BoolVar(&FlagVerbose, "v", false, "Verbose mode.")
	flag.StringVar(&FlagURL, "url", "", "Visit specified absolute URL.")
	flag.StringVar(&FlagExt, "ext", "", "Look for files with specified extensions.")
	flag.IntVar(&FlagLazy, "lazy", 0, "Crawl with specified delay.")
	flag.StringVar(&FlagDest, "target", "", "Specify target directory for the downloads.")
	flag.Parse()
}
