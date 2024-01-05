package main

import (
	"log"
	"os"
	"os/exec"
	"path"
)

// find all sub folders in /Volumes/Klipp/twitch-recorder/videos
// /Users/figge/tmp/video-conversion/videos/

// handbrakecli convert to /Volumes/Klipp/twitch-vods/videos/{folder}/{file
// /Users/figge/tmp/video-conversion/processed/

func main() {
    unprocessedDir := "/Volumes/Klipp/Twitch VODs/Unprocessed/"
    processedDir := "/Volumes/Klipp/Twitch VODs/"
    // unprocessedDir := "/Volumes/Klipp/twitch-recorder/videos/"
    dirs, err := os.ReadDir(unprocessedDir)
    if err != nil {
        log.Fatal(err)
    }

    for i, dir := range dirs {
        if !dir.IsDir() {
            dirs = append(dirs[:i], dirs[i+1:]...)
        }
    }

    for _, dir := range dirs {
        log.Println(dir.Name())
        videosDir := path.Join(unprocessedDir, dir.Name())
        videos, err := os.ReadDir(videosDir)
        if err != nil {
            log.Fatal(err)
        }

        for _, video := range videos {
            log.Println(video.Name())

            //check if processed folder exists
            userPath := path.Join(processedDir, dir.Name())
            _, err := os.Stat(userPath)
            if err != nil {
                os.Mkdir(userPath, 0777)
            }
            //handbrakecli --preset-import-file /Users/figge/go/src/github.com/jawee/convert-twitch-vods/twitch-vod.json --preset twitch-vod -i /Users/figge/tmp/video-conversion/videos/Roxkstar74/20220227_224638_why_is_all_my_income_discord_bots_today.mp4 -o /Users/figge/tmp/video-conversion/processed/Roxkstar74/20220227_224638_why_is_all_my_income_discord_bots_today.mp4
            inputFilePath := path.Join(unprocessedDir, dir.Name(), video.Name())
            outputFilePath := path.Join(processedDir, dir.Name(), video.Name())
            cmd := exec.Command("handbrakecli", "--preset-import-file",
                    "./twitch-vod.json", "--preset", "twitch-vod",
                    "-i", inputFilePath, "-o", outputFilePath)
            log.Println(cmd.Args)
            cmd.Stdout = os.Stdout
            err = cmd.Run()
            if err != nil {
                log.Fatal(err)
            }
            err = os.Remove(inputFilePath)
            if err != nil {
                log.Printf("Error removing file: %s", err)
            }
        }
    }
}

