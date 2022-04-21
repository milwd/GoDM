# GoDM
### basic Download manager with parallel downloads, implemented in GOlang 

a simple download manager, written in GO
parallel downloads with the help of concurrency

# how does it work?
* gets head of http response
* gets size and availability to download in sections (instead of the whole file) with respect to http range header response
* initializes go routines for partial downloads
* run them simultaneously until all is Done
* merge them

# run
- install <a href=https://go.dev/doc/install>Go</a>
- in the program folder run `go run .`
- open the output file with related applications

# note
the merge results may have flaws due to file compressing!

<table>
  <tr>
    <td>Original <a href=https://www.industrialempathy.com/img/remote/ZiClJf-1920w.jpg>(source)</a></td>
     <td>Downloaded and Compressed</td>
  </tr>
  <tr>
    <td><img src="imgs/original.jpg" width=533 height=300></td>
    <td><img src="imgs/output.jpg" width=533 height=300></td>
  </tr>
 </table>

memory management is neccessary for rather large files, not yet implemented!
