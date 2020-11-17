// Subscribe to a camera feed. Save captured frame to disk as image.

package main

import (
	"fmt"
	"os"
	"time"

	"gocv.io/x/gocv"
)

type Frame struct {
	image gocv.Mat
}

var (
	deviceName    string
	deviceId      string
	saveDirectory string
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("go run main.go [camera name] [camera id or url] [save directory]")
		return
	}

	deviceName = os.Args[1]
	deviceId = os.Args[2]
	saveDirectory = os.Args[3]

	device, err := gocv.OpenVideoCapture(deviceId)
	if nil != err {
		fmt.Printf("Error reading device: %v\n", deviceId)
	}
	defer device.Close()

	frame := Frame{image: gocv.NewMat()}
	defer frame.image.Close()

	for {
		if success := device.Read(&frame.image); !success {
			fmt.Printf("Error reading device: %v\n", deviceId)
			return
		}

		if frame.image.Empty() {
			fmt.Println("frame empty")
			continue
		}

		if success := frame.dump(deviceName, saveDirectory); !success {
			fmt.Printf("Error saving frame to disk: %v\n", err)
			return
		}
	}
}

// saving frame (image) to disk
func (f *Frame) dump(name string, dir string) bool {
	location := fmt.Sprintf("%s/%s-%d.jpeg", dir, name, time.Now().UnixNano())
	return gocv.IMWrite(location, f.image)
}
