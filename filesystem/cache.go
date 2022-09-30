package filesystem

import (
	"common"
	"encoding/base64"
	filesystem_structs "filesystem/structs"
	"fmt"
	"net/url"
	"network"
	"strings"
)

var resourceCache = &filesystem_structs.ResourceCache{
	ResourceMap: make(map[string]*filesystem_structs.Resource, 99),
}
var imageCache = &filesystem_structs.ImageCache{
	ImageMap: make(map[string]*filesystem_structs.Image, 99),
}

func ParseURL(u string) *url.URL {
	URL, err := url.Parse(u)
	if err != nil {
		URL = ParseURL(common.WebBrowserName + "://error?err=failedToParseURL")
	}
	return URL
}

func AddResource(resource *filesystem_structs.Resource) {
	resourceCache.AddResource(resource)
}

func GetResource(url string) *filesystem_structs.Resource {
	URL := ParseURL(url)
	return GetResourceByURL(URL)
}

func GetResourceByURL(URL *url.URL) *filesystem_structs.Resource {
	switch URL.Scheme {
	case strings.ToLower(common.WebBrowserName):
		return getInternalPage(URL)
	case "file":
		return getLocalPage(URL)
	default:
		URL.Scheme = "http"
	}
	return GetExternalPage(URL)
}

func GetHomepage() string {
	return `<html>
		<head>
			<title>HomePage</title>
		</head>
		<body>
			` + common.WebBrowserName + ` Home Page
		</body>
	</html>`
}

func GetHistoryPage() string {
	return ``
}

func GetAboutPage() string {
	return ``
}

func getInternalPage(URL *url.URL) *filesystem_structs.Resource {
	var body string
	switch URL.Host {
	case "HomePage":
		body = GetHomepage()
	case "History":
		body = GetHistoryPage()
	case "About":
		body = GetAboutPage()
	default:
		body = ""
	}
	resource := &filesystem_structs.Resource{
		Body: body,
		URL:  URL,
	}
	return resource
}

func getLocalPage(URL *url.URL) *filesystem_structs.Resource {
	resource := &filesystem_structs.Resource{}
	return resource
}

func GetExternalPage(URL *url.URL) *filesystem_structs.Resource {
	url := URL.String()
	cache := resourceCache.GetResource(url)
	if cache != nil {
		return cache
	}
	resource := &filesystem_structs.Resource{
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

func AddImage(image *filesystem_structs.Image) {
	imageCache.AddImage(image)
}

func GetImage(URL *url.URL) (*filesystem_structs.Image, error) {
	url := URL.String()
	cache := imageCache.GetImage(url)
	if cache != nil {
		return cache, nil
	}
	image := &filesystem_structs.Image{
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
