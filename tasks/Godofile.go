package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/godo.v1"
)

func tasks(p *godo.Project) {
	godo.Env = `GOPATH=.vendor::$GOPATH`

	p.Task("getframes", func(c *godo.Context) {
		var videosDir os.FileInfo

		dataDir := "data/"
		dataDirList, _ := ioutil.ReadDir(dataDir)

		for i := range dataDirList {
			if dataDirList[i].IsDir() && dataDirList[i].Name() == "videos" {
				videosDir = dataDirList[i]
				break
			}
		}

		var seasonsList []string
		videosDirList, _ := ioutil.ReadDir(dataDir + videosDir.Name())
		for i := range videosDirList {
			if videosDirList[i].IsDir() {
				d := dataDir + "frames/" + videosDirList[i].Name()
				err := os.MkdirAll(d, 0777)
				if err != nil {
					log.Fatal(err)
				}
				seasonsList = append(seasonsList, videosDirList[i].Name())
			}
		}

		fmt.Println("Found " + string(len(seasonsList)) + " seasons to convert to thumbnails. This will take some time.")

		for _, seasonDir := range seasonsList {
			seasonVideosList, _ := ioutil.ReadDir(dataDir + videosDir.Name() + "/" + seasonDir)
			// var episodesList []string

			// Make dirs in data/frames for all of the episodes
			for j := range seasonVideosList {
				epName := strings.Replace(seasonVideosList[j].Name(), ".mp4", "", -1)
				epInPath := dataDir + "videos/" + seasonDir + "/" + seasonVideosList[j].Name()
				epOutPath := dataDir + "frames/" + seasonDir + "/" + epName
				os.MkdirAll(epOutPath, 0777)
				// episodesList = append(episodesList, epPath)
				c := exec.Command(
					"/root/bin/ffmpeg",
					"-i",
					epInPath,
					"-vf",
					"fps=1",
					epOutPath+"/out%04d.png",
				)
				_, err := c.CombinedOutput()
				if err == nil {
					fmt.Println("Converted " + epName + " with no errors")
				} else {
					fmt.Println(err)
				}
			}
		}
	})
}

func main() {
	godo.Godo(tasks)
}
