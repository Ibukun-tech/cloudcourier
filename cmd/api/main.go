package main

import (
	"log/slog"
	"os"

	"github.com/Ibukun-tech/cloudcourier"
)

func main() {
	cbb := &cloudcourier.CloudCourierBridge{}
	st, err := cloudcourier.NewCloudCourier(cbb)
	if err != nil {
		slog.Error("%s", err)
	}
	fl, err := os.Open("./20171451.jpg")
	if err != nil {
		slog.Error("%s", err)
	}
	st.UploadFile("/", fl)
}
