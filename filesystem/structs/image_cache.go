package structs

type Image struct {
	Key  string
	Data []byte
}

type ImageCache struct {
	images   []*Image
	imageMap map[string]*Image
}

func (cache *ImageCache) AddImage(image *Image) {
	cache.images = append(cache.images, image)
	cache.imageMap[image.Key] = image
}

func (cache *ImageCache) GetImage(key string) *Image {
	return cache.imageMap[key]
}
