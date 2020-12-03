// this package implements;
// 1) subscription to a camera feed.
// 2) write frame(s) to disk as image(s).

package capture

import (
	"fmt"

	"github.com/serialberry/core-stream/pkg/storage"
	"gocv.io/x/gocv"
)

type Device struct {
	Id   string
	Name string
	Dir  string
}

// start reading from camera stream
func (d *Device) Read() {
	directory := &storage.Directory{Path: d.Dir}
	storage.CreateIfDirNotExists(directory, directory.Path)

	device, err := gocv.OpenVideoCapture(d.Id)
	if nil != err {
		fmt.Printf("Error reading device: %v\n", d.Id)
	}
	defer device.Close()

	image := &storage.File{Data: gocv.NewMat(), Path: d.Name}
	defer image.Data.Close()

	nextId := storage.NewSequenceId()

	for {
		if success := device.Read(&image.Data); !success {
			fmt.Printf("Error reading device: %v\n", d.Id)
			return
		}

		if image.Data.Empty() {
			fmt.Println("frame empty")
			continue
		}

		if success := storage.SaveImageToDisk(image, nextId, directory.Path, image.Path); !success {
			fmt.Printf("Error saving frame to disk: %v\n", err)
			return
		}
	}
}
