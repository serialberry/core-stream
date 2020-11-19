// this package implements;
// 1) subscription to a camera feed.
// 2) write frame(s) to disk as image(s).

package capture

import (
	"fmt"
	"time"

	"gocv.io/x/gocv"
)

type Device struct {
	Id   string
	Name string
	Dir  string
}

// start reading from camera stream
func (d *Device) Read() {
	d.MakeDirIfNotExists()

	device, err := gocv.OpenVideoCapture(d.Id)
	if nil != err {
		fmt.Printf("Error reading device: %v\n", d.Id)
	}
	defer device.Close()

	image := gocv.NewMat()
	defer image.Close()

	for {
		if success := device.Read(&image); !success {
			fmt.Printf("Error reading device: %v\n", d.Id)
			return
		}

		if image.Empty() {
			fmt.Println("frame empty")
			continue
		}

		if success := d.Save(image); !success {
			fmt.Printf("Error saving frame to disk: %v\n", err)
			return
		}
	}
}

// create directory location if not exists
func (d *Device) MakeDirIfNotExists() {

}

// write frame to disk
func (d *Device) Save(frame gocv.Mat) bool {
	image := fmt.Sprintf("%s/%s-%d.jpeg", d.Dir, d.Name, time.Now().UnixNano())
	return gocv.IMWrite(image, frame)
}
