# GoDM
Download manager with parallel download, implemented in GOlang (project)

a simple download manager, written in GO

parallel downloads with the help of concurrency

* gets head of http response
* gets size and availability to download in sections (instead of the whole file) with respect to http range header response
* initializes go routines for partial downloads
* run them simultaneously until all is Done
* merge them

(the merge results may have flaws due to file compressing)