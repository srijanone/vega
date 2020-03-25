package downloader

import (
	"context"
	"fmt"
	"os"
	"sync"

	getter "github.com/hashicorp/go-getter"
)

type Downloader struct{}

func (d *Downloader) Download(src string, dest string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error in getting working directory: %s", err)
	}

	opts := []getter.ClientOption{}

	client := &getter.Client{
		Ctx:     context.Background(),
		Src:     src,
		Dst:     dest,
		Pwd:     pwd,
		Mode:    getter.ClientModeAny,
		Options: opts,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	errChannel := make(chan error, 2)

	go func() {
		defer wg.Done()
		if err := client.Get(); err != nil {
			fmt.Println("Error in downloading", err)
			errChannel <- err
		}
	}()

	wg.Wait()
}

// For testing now, will be removed
//func main() {
//	fmt.Println("Downloading")
//	d := &Downloader{}
//	d.Download("git@github.com:Azure/draft.git//packs", "packs")
//}
