package conf

import (
	"fmt"
)

func GetImageUrl(imageName string) (string, error) {
	images_url := map[string]string{
		"ubuntu-latest": "https://cloud-images.ubuntu.com/jammy/current/jammy-server-cloudimg-amd64.img",
		"fedora-38":     "https://download.fedoraproject.org/pub/fedora/linux/releases/38/Cloud/x86_64/images/Fedora-Cloud-Base-38-1.6.x86_64.qcow2",
	}

	image, ok := images_url[imageName]
	if ok {
		return image, nil
	}

	return "", fmt.Errorf("image %s not found", imageName)
}
