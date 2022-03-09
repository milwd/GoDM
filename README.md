# GoDM
### Download manager with parallel downloads, implemented in GOlang (project)

a simple download manager, written in GO
parallel downloads with the help of concurrency

# how does it work?
* gets head of http response
* gets size and availability to download in sections (instead of the whole file) with respect to http range header response
* initializes go routines for partial downloads
* run them simultaneously until all is Done
* merge them

# note
the merge results may have flaws due to file compressing!
memory management is neccessary for rather large files, not yet implemented!
