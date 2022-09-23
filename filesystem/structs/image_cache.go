package structs

type Image struct {
	Key  string
	Data []byte
}

type ImageCache struct {
	Images   []*Image
	ImageMap map[string]*Image
}

func (cache *ImageCache) AddImage(image *Image) {
	cache.Images = append(cache.Images, image)
	cache.ImageMap[image.Key] = image
}

func (cache *ImageCache) GetImage(key string) *Image {
	return cache.ImageMap[key]
}
