package http

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"

	"golang.org/x/net/proxy"

	"os"
	"path/filepath"

	"io"

	"github.com/gogather/com"
	"github.com/gogather/com/log"
	"github.com/toolkits/file"
)

// Jar - cookie jar
type Jar struct {
	lk          sync.Mutex
	CookiesData map[string][]*http.Cookie
}

// NewJar - new a Jar
func NewJar() *Jar {
	jar := new(Jar)
	jar.CookiesData = make(map[string][]*http.Cookie)
	return jar
}

func (j *Jar) addCookie(host string, cookie *http.Cookie) {
	hostCookiesData, ok := j.CookiesData[host]
	if ok {
		finded := false
		for i := 0; i < len(hostCookiesData); i++ {
			c := hostCookiesData[i]
			if c.Name == cookie.Name {
				hostCookiesData[i] = cookie
				finded = true
			}
		}

		if !finded {
			hostCookiesData = append(hostCookiesData, cookie)
		}

		j.CookiesData[host] = hostCookiesData
	} else {
		j.CookiesData[host] = append(hostCookiesData, cookie)
	}

}

// SetCookies - set cookies
func (j *Jar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	j.lk.Lock()
	for i := 0; i < len(cookies); i++ {
		j.addCookie(u.Host, cookies[i])
	}
	j.lk.Unlock()
}

// Cookies - get cookie
func (j *Jar) Cookies(u *url.URL) []*http.Cookie {
	return j.CookiesData[u.Host]
}

// HTTPClient - http client
type HTTPClient struct {
	cookiePath string
	jar        *Jar
	client     *http.Client
}

// NewHTTPClientWithCookieFile - new an HTTPClient from cookiePath
func NewHTTPClientWithCookieFile(cookiePath string) *HTTPClient {
	return newHTTPClientWithCookieFileAndTransport(cookiePath, nil)
}

// NewProxyHTTPClientWithCookieFile - new a proxy HTTPClient from cookiePath
func NewProxyHTTPClientWithCookieFile(cookiePath, proxyAddr string) *HTTPClient {
	transport := newProxyTransport(proxyAddr)
	return newHTTPClientWithCookieFileAndTransport(cookiePath, transport)
}

func newHTTPClientWithCookieFileAndTransport(cookiePath string, transport *http.Transport) *HTTPClient {
	hc := &HTTPClient{}
	hc.cookiePath = cookiePath
	jar := NewJar()
	jsonData, err := com.ReadFileString(cookiePath)
	if err == nil {
		err = json.Unmarshal([]byte(jsonData), jar)
		if err != nil {
			log.Warnln("illeage cookies jar file")
		}
	}

	hc.jar = jar

	hc.client = &http.Client{Transport: transport, CheckRedirect: nil, Jar: hc.jar, Timeout: 0}

	return hc
}

// NewHTTPClient new an http client
func NewHTTPClient() *HTTPClient {
	return newHTTPClient(nil)
}

// NewProxyHTTPClient new a proxy http client
func NewProxyHTTPClient(proxyAddr string) *HTTPClient {
	transport := newProxyTransport(proxyAddr)
	return newHTTPClient(transport)
}

func newHTTPClient(transport *http.Transport) *HTTPClient {
	hc := &HTTPClient{}
	jar := NewJar()
	hc.jar = jar
	if transport == nil {
		hc.client = &http.Client{Transport: nil, CheckRedirect: nil, Jar: hc.jar, Timeout: 0}
	} else {
		hc.client = &http.Client{Transport: transport, CheckRedirect: nil, Jar: hc.jar, Timeout: 0}
	}
	return hc
}

func newProxyTransport(proxyAddr string) *http.Transport {
	tbProxyURL, err := url.Parse(proxyAddr)
	if err != nil {
		log.Warnf("Failed to parse proxy URL: %v\n", err)
	}

	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
	if err != nil {
		log.Warnf("Failed to obtain proxy dialer: %v\n", err)
	}

	return &http.Transport{Dial: tbDialer.Dial}
}

func (h *HTTPClient) serialze() {
	jar := h.jar
	jsonData, err := com.JsonEncode(jar)
	if err == nil && h.cookiePath != "" {
		com.WriteFile(h.cookiePath, string(jsonData))
	}
}

// Post - post method
func (h *HTTPClient) Post(urlstr string, parm url.Values) (string, error) {
	resp, err := h.client.PostForm(urlstr, parm)

	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	h.serialze()
	return string(b), err
}

// Get - get method
func (h *HTTPClient) Get(urlstr string) (string, error) {
	resp, err := h.client.Get(urlstr)
	if err != nil {
		return "", err
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return "", err
	}
	h.serialze()
	return string(b), err
}

// Download download file to path
func (h *HTTPClient) Download(urlstr, path string) error {
	var err error
	directory := filepath.Dir(path)
	if !file.IsExist(directory) {
		err = os.MkdirAll(directory, 0755)
		if err != nil {
			return err
		}
	}
	resp, err := h.client.Get(urlstr)
	if err != nil {
		return err
	}
	fmt.Println(urlstr, resp.Status)
	file, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, resp.Body)
	if err!=nil {
		fmt.Println("io copy error")
		return err
	}
	err = file.Sync()
	if err != nil {
		fmt.Println("io sync error")
	}
	return err
}

// Do http client do method
func (h *HTTPClient) Do(r *http.Request) (*http.Response, error) {
	return h.client.Do(r)
}
