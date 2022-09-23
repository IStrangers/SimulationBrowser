package filesystem

import (
	"filesystem/structs"
	"net/url"
)

var resourceCache = &structs.ResourceCache{}
var imageCache = &structs.ImageCache{}

func AddResource(resource *structs.Resource) {
	resourceCache.AddResource(resource)
}

func GetResource(URL *url.URL) *structs.Resource {
	url := URL.String()
	return resourceCache.GetResource(url)
}

func AddImage(image *structs.Image) {
	imageCache.AddImage(image)
}

func GetImage(URL *url.URL) *structs.Image {
	url := URL.String()
	return imageCache.GetImage(url)
}
