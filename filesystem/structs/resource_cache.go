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
	resources   []*Resource
	resourceMap map[string]*Resource
}

func (cache *ResourceCache) AddResource(resource *Resource) {
	cache.resources = append(cache.resources, resource)
	cache.resourceMap[resource.Key] = resource
}

func (cache *ResourceCache) GetResource(key string) *Resource {
	return cache.resourceMap[key]
}
