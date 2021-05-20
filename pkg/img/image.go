package img

import (
	"image"
	"net/http"
)

func ImageConfigFromURL(url string) (image.Config, error) {
	res, err := http.Get(url)
	if err != nil {
		return image.Config{}, err
	}
	defer res.Body.Close()
	imconfig, _, err := image.DecodeConfig(res.Body)
	return imconfig, err

}
