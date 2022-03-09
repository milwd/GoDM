package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"sync"
	"time"
)

/*

HTTP HEAD => content length int > 0             != -1
HTTP HEAD => header["accept-ranges"] == [bytes] != []

resp is a pointer

*/

func downloadParti(wg *sync.WaitGroup, min, max, ind int, folder string, all *[]bool, url string) {
	defer wg.Done()
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
	range_header := "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max)
	req.Header.Add("Range", range_header)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln("download part "+strconv.Itoa(ind)+"failed!\nerr: ", err)
	}
	defer resp.Body.Close()
	out, err := os.Create(folder + "/" + strconv.Itoa(ind))
	if err != nil {
		log.Fatalln("can't create output file (probably permission denied)!\nerr: ", err)
	}
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		log.Fatalln("can't insert data to output file "+strconv.Itoa(ind)+"!\nerr: ", err)
	}
	(*all)[ind] = true
	fmt.Println("done part ", ind)
}

func Find(slice []bool, val bool) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func merge1(all []bool, files []string) {
	_, found := Find(all, false)
	if !found {
		var buf bytes.Buffer
		for _, file := range files {
			b, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatalln("can't read file "+file+"\nerr: ", err)
			}
			buf.Write(b)
			time.Sleep(1000 * time.Millisecond)
		}
		err := ioutil.WriteFile("output_file", buf.Bytes(), 0777) // 0644
		if err != nil {
			log.Fatalln("can't merge files\nerr: ", err)
		}
	} else {
		fmt.Println("not all parts are available!")
	}
}

func merge2(all []bool, files []string) {
	_, found := Find(all, false)
	if !found {
		out, err := os.OpenFile("output", os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatalln("can't merge files\nerr: ", err)
		}
		defer out.Close()
		for _, file := range files {
			n, _ := os.Open(file)
			defer n.Close()
			io.Copy(out, n)
			time.Sleep(1000 * time.Millisecond)
		}
	} else {
		fmt.Println("not all parts are available!")
	}
}

func main() {
	// TODO : DOWNLOAD PROGRESS

	var url string
	var nParts int
	fmt.Println("enter URL of the file : ")
	fmt.Scanln(&url)
	fmt.Println("how many partitions ? (recommended : 5)")
	fmt.Scanln(&nParts)

	allDone := make([]bool, nParts)
	files := make([]string, nParts)
	cwd, _ := os.Getwd()
	os.Mkdir("parts", 0755)
	for i := 0; i < nParts; i++ {
		files[i] = path.Join(cwd, "parts", strconv.Itoa(i))
	}
	fmt.Println("filenames are : ")
	fmt.Println(files)

	var wg sync.WaitGroup

	resp, err := http.Head(url)
	// TODO : FILE FORMAT FOR THE OUTPUT
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("file byte length : ", int(resp.ContentLength))
	fmt.Println("headers : ", resp.Header["Accept-Ranges"])

	if resp.Header["Accept-Ranges"] == nil || int(resp.ContentLength) == -1 {
		log.Fatalln("requested server does not allow Range Requests or the file is out of reach!")
	}

	rng := int(resp.ContentLength)
	for i := 0; i < nParts-1; i++ {
		mn := i * (rng / nParts)
		mx := (i + 1) * (rng / nParts)
		wg.Add(1)
		go downloadParti(&wg, mn, mx, i, "parts", &allDone, url)
	}
	mn := (nParts - 1) * (rng / nParts)
	mx := rng
	wg.Add(1)
	go downloadParti(&wg, mn, mx, nParts-1, "parts", &allDone, url)

	fmt.Println("Main: Waiting for workers to finish")
	wg.Wait()
	merge2(allDone, files)
	fmt.Println("Main: Completed")
}
