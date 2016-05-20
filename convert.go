package mediaserver

import (
    "os"
    "io"
    "os/exec"
	"fmt"
)

func convert(filename, pathname, videoBitrate, bufsize, resolution string) error {
    cmd := exec.Command("ffmpeg", 
        "-y", 
        "-i " + filename, 
        "-c:a aac", 
        "-ac 2", 
        "-b:a 128k", 
        "-c:v libx264",
        "-x264opts 'keyint=24:min-keyint=24:no-scenecut'",
        "-b:v " + videoBitrate,
        "-maxrate " + videoBitrate,
        "-bufsize " + bufsize,
        "-vf \"scale=-1:" + resolution + "\"",
        pathname + "/" + resolution + "p.mp4",)
        cmd.Stderr = os.Stderr
        cmd.Stdout = os.Stdout
    return cmd.Run()
}

func convertVideoToStream(filename, pathname, descdir, descname, details string) {
    err := os.MkdirAll(descdir, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    err = os.MkdirAll(descdir + "/dash/", 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    fd, err := os.OpenFile(descdir + "/" + descname, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer fd.Close()
    _, err = io.WriteString(fd, details)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    err = convert(filename, pathname, "400k", "400k", "360")
    if err != nil {
        fmt.Println(err)
        return
    }
    
    err = convert(filename, pathname, "800k", "500k", "560")
    if err != nil {
        fmt.Println(err)
        return
    }
    
    err = convert(filename, pathname, "1500k", "1000k", "720")
    if err != nil {
        fmt.Println(err)
    }
    
    err = convert(filename, pathname, "3000k", "2500k", "1080")
    if err != nil {
        fmt.Println(err)
    }
    
    cmd := exec.Command("MP4Box",
        "-dash 5000",
        "-rap",
        "-frag-rap",
        "-profile onDemand",
        "-out " + pathname + "/dash/info.mpd",
        pathname + "/360p.mp4",
        pathname + "/560p.mp4",
        pathname + "/720p.mp4",
        pathname + "/1080p.mp4",
        )
    err = cmd.Run()
    if err != nil {
        fmt.Println(err)
    }
    return
}