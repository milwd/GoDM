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

<<<<<<< HEAD
- just open the output with a proper application

# note
the merge results may have flaws due to file compressing!

<p float="left">
  <img src="imgs/original.jpg" width="100" />
  <img src="imgs/output.jpg" width="100" /> 
</p>

memory management is neccessary for rather large files, not yet implemented!
=======
# note
the merge results may have flaws due to file compressing!

memory management is neccessary for rather large files, not yet implemented!
>>>>>>> 27332434c7f6806cb6cbf0be591dca97f1cd3879
