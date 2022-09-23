package structs

import "net/url"

type Resource struct {
	Body        string
	ContentType string
	Code        int
	URL         *url.URL
	Key         string
}

type ResourceCache struct {
	Resources   []*Resource
	ResourceMap map[string]*Resource
}

func (cache *ResourceCache) AddResource(resource *Resource) {
	cache.Resources = append(cache.Resources, resource)
	cache.ResourceMap[resource.Key] = resource
}

func (cache *ResourceCache) GetResource(key string) *Resource {
	return cache.ResourceMap[key]
}
