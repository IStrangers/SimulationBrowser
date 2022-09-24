package filesystem

import (
	browser "browser/structs"
	"common"
	"encoding/base64"
	filesystem "filesystem/structs"
	"fmt"
	"net/url"
	"network"
	"strings"
)

var resourceCache = &filesystem.ResourceCache{
	ResourceMap: make(map[string]*filesystem.Resource, 99),
}
var imageCache = &filesystem.ImageCache{
	ImageMap: make(map[string]*filesystem.Image, 99),
}

func ParseURL(u string) *url.URL {
	URL, err := url.Parse(u)
	if err != nil {
		URL = ParseURL(browser.WebBrowserName + "://error?err=failedToParseURL")
	}
	return URL
}

func AddResource(resource *filesystem.Resource) {
	resourceCache.AddResource(resource)
}

func GetResource(url string) *filesystem.Resource {
	URL := ParseURL(url)
	return GetResourceByURL(URL)
}

func GetResourceByURL(URL *url.URL) *filesystem.Resource {
	switch URL.Scheme {
	case browser.WebBrowserName:
		return getInternalPage(URL)
	case "file":
		return getLocalPage(URL)
	default:
		URL.Scheme = "http"
	}
	return GetExternalPage(URL)
}

func getInternalPage(URL *url.URL) *filesystem.Resource {
	resource := &filesystem.Resource{}
	return resource
}

func getLocalPage(URL *url.URL) *filesystem.Resource {
	resource := &filesystem.Resource{}
	return resource
}

func GetExternalPage(URL *url.URL) *filesystem.Resource {
	url := URL.String()
	cache := resourceCache.GetResource(url)
	if cache != nil {
		return cache
	}
	resource := &filesystem.Resource{
		Key: url,
		URL: URL,
	}
	result, err := network.SendGetRequest(url)
	if err != nil {
		resource.Body = loaderErrorPage(err.Error())
		return resource
	}
	resource.Body = string(result.Body)
	resource.ContentType = result.ContentType
	resource.Code = result.StatusCode
	AddResource(resource)
	return resource
}

func AddImage(image *filesystem.Image) {
	imageCache.AddImage(image)
}

func GetImage(URL *url.URL) (*filesystem.Image, error) {
	url := URL.String()
	cache := imageCache.GetImage(url)
	if cache != nil {
		return cache, nil
	}
	image := &filesystem.Image{
		Key: url,
	}
	if len(url) >= 22 && common.IsBase64Image(url) {
		imgData := url[strings.IndexByte(url, ',')+1:]
		decodedData, err := base64.RawStdEncoding.DecodeString(imgData)
		if err != nil {
			return nil, fmt.Errorf("Failed to decode base64 data (%s)", err)
		}
		image.Data = decodedData
	} else {
		result, err := network.SendGetRequest(url)
		if err != nil {
			return nil, err
		}
		image.Data = result.Body
	}
	AddImage(image)
	return image, nil
}
