/*
 * Copyright (c) 2012, Yang Junyong <yanunon@gmail.com>
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *     * Redistributions of source code must retain the above copyright
 *       notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above copyright
 *       notice, this list of conditions and the following disclaimer in the
 *       documentation and/or other materials provided with the distribution.
 *     * Neither the name of the Google Inc. nor the
 *       names of its contributors may be used to endorse or promote products
 *       derived from this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
 * ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
 * WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
 * DISCLAIMED. IN NO EVENT SHALL <COPYRIGHT HOLDER> BE LIABLE FOR ANY
 * DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
 * (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
 * LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
 * ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
 * SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

// Package Aliyun OSS API.
//
package oss

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	ACL_PUBLIC_RW = "public-read-write"
	ACL_PUBLIC_R  = "public-read"
	ACL_PRIVATE   = "private"
)

type AccessControlList struct {
	Grant string
}
type AccessControlPolicy struct {
	Owner             Owner
	AccessControlList AccessControlList
}

type Bucket struct {
	Name         string
	CreationDate string
}

type CompleteMultipartUpload struct {
	Part []Multipart
}

type Multipart struct {
	PartNumber int
	ETag       string
}

type CompleteMultipartUploadResult struct {
	Location string
	Bucket   string
	ETag     string
	Key      string
}

type FileGroup struct {
	Bucket     string
	Key        string
	ETag       string
	FileLength int
	FilePart   CreateFileGroup
}

type CreateFileGroup struct {
	Part []GroupPart
}

type CompleteFileGroup struct {
	Bucket string
	Key    string
	Size   int
	ETag   string
}

//GroupPart's partname should be the object's key and ETag is the same with the object's md5sum.
type GroupPart struct {
	PartNumber int
	PartName   string
	PartSize   int
	ETag       string
}

type Client struct {
	AccessID     string
	AccessKey    string
	Host         string
	HttpClient   *http.Client
	FileIOLocker sync.Mutex
	ChanNum      int
}

type Buckets struct {
	Bucket []Bucket
}

type initMultipartUploadResult struct {
	Bucket   string
	Key      string
	UploadId string
}

type ListAllMyBucketsResult struct {
	Owner   Owner
	Buckets Buckets
}

type ListBucketResult struct {
	Name        string
	Prefix      string
	Marker      string
	NextMarker  string
	MaxKeys     int
	Delimiter   string
	IsTruncated bool
	Contents    []Object
}

type Object struct {
	Key          string
	LastModified string
	ETag         string
	Type         string
	Size         int
	StorageClass string
	Owner        Owner
}

type Owner struct {
	ID          string
	DisplayName string
}

type ListMultipartUploadResult struct {
	Bucket             string
	KeyMarker          string
	UploadIdMarker     string
	NextKeyMarker      string
	NextUploadIdMarker string
	Delimiter          string
	Prefix             string
	MaxUploads         int
	IsTruncated        bool
	Upload             []UploadPart
}

type UploadPart struct {
	Key       string
	UploadId  string
	Initiated string
}

type UploadedPart struct {
	PartNumber   int
	LastModified string
	ETag         string
	Size         int
}

type ListPartsResult struct {
	Bucket               string
	Key                  string
	UploadId             string
	NextPartNumberMarker string
	MaxParts             int
	IsTruncated          bool
	Part                 []UploadedPart
}

type valSorter struct {
	Keys []string
	Vals []string
}

type multipartSorter struct {
	Parts []Multipart
}

//NewClient returns a new Client given a Host, AccessID and AccessKey.
//ChanNum is the number of goroutines used in multipart upload.
func NewClient(host, accessId, accessKey string, channum int) *Client {
	client := Client{
		Host:       host,
		AccessID:   accessId,
		AccessKey:  accessKey,
		HttpClient: http.DefaultClient,
		ChanNum:    channum,
	}
	return &client
}

func (c *Client) signHeader(req *http.Request, canonicalizedResource string) {
	//format x-oss-
	tmpParams := make(map[string]string)

	for k, v := range req.Header {
		if strings.HasPrefix(strings.ToLower(k), "x-oss-") {
			tmpParams[strings.ToLower(k)] = v[0]
		}
	}
	//sort
	vs := newValSorter(tmpParams)
	vs.Sort()

	canonicalizedOSSHeaders := ""
	for i := range vs.Keys {
		canonicalizedOSSHeaders += vs.Keys[i] + ":" + vs.Vals[i] + "\n"
	}

	date := req.Header.Get("Date")
	contentType := req.Header.Get("Content-Type")
	contentMd5 := req.Header.Get("Content-Md5")

	signStr := req.Method + "\n" + contentMd5 + "\n" + contentType + "\n" + date + "\n" + canonicalizedOSSHeaders + canonicalizedResource
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(c.AccessKey)) //sha1.New()
	io.WriteString(h, signStr)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	authorizationStr := "OSS " + c.AccessID + ":" + signedStr
	//fmt.Println(authorizationStr)
	req.Header.Set("Authorization", authorizationStr)
}

func (c *Client) doRequest(method, path, canonicalizedResource string, params map[string]string, data io.Reader) (resp *http.Response, err error) {
	reqUrl := "http://" + c.Host + path
	req, _ := http.NewRequest(method, reqUrl, data)
	date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")
	req.Header.Set("Date", date)
	req.Header.Set("Host", c.Host)

	if params != nil {
		for k, v := range params {
			req.Header.Set(k, v)
		}
	}

	if data != nil {
		req.Header.Set("Content-Length", strconv.Itoa(int(req.ContentLength)))
	}
	c.signHeader(req, canonicalizedResource)
	resp, err = c.HttpClient.Do(req)
	return
}

//Get bucket list. Return a ListAllMyBucketsResult object.
func (c *Client) GetService() (lar ListAllMyBucketsResult, err error) {
	resp, err := c.doRequest("GET", "/", "/", nil, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}

	err = xml.Unmarshal(body, &lar)
	return
}

//Create a new bucket with a name.
func (c *Client) PutBucket(bname string) (err error) {
	reqStr := "/" + bname
	resp, err := c.doRequest("PUT", reqStr, reqStr, nil, nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
	}
	return
}

//Set bucket access control list.
func (c *Client) PutBucketACL(bname, acl string) (err error) {
	params := map[string]string{"x-oss-acl": acl}
	reqStr := "/" + bname
	resp, err := c.doRequest("PUT", reqStr, reqStr, params, nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
	}
	return
}

//Get bucket's object list.
func (c *Client) GetBucket(bname, prefix, marker, delimiter, maxkeys string) (lbr ListBucketResult, err error) {
	reqStr := "/" + bname
	resStr := reqStr
	query := map[string]string{}
	if prefix != "" {
		query["prefix"] = prefix
	}

	if marker != "" {
		query["marker"] = marker
	}

	if delimiter != "" {
		query["delimiter"] = delimiter
	}

	if maxkeys != "" {
		query["max-keys"] = maxkeys
	}

	if len(query) > 0 {
		reqStr += "?"
		for k, v := range query {
			reqStr += k + "=" + v + "&"
		}
	}

	resp, err := c.doRequest("GET", reqStr, resStr, nil, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}
	err = xml.Unmarshal(body, &lbr)
	return
}

//Get bucket's access control list.
func (c *Client) GetBucketACL(bname string) (acl AccessControlPolicy, err error) {
	reqStr := "/" + bname + "?acl"
	resp, err := c.doRequest("GET", reqStr, reqStr, nil, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}

	err = xml.Unmarshal(body, &acl)
	return
}

//Delete bucket by name.
func (c *Client) DeleteBucket(bname string) (err error) {
	return c.DeleteObject(bname)
}

//Copy object from path src to dst. The format of path is "/bucketName/objectName".
func (c *Client) CopyObject(dst, src string) (err error) {
	if strings.HasPrefix(src, "/") == false {
		src = "/" + src
	}
	if strings.HasPrefix(dst, "/") == false {
		dst = "/" + dst
	}
	params := map[string]string{"x-oss-copy-source": src}
	resp, err := c.doRequest("PUT", dst, dst, params, nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
	}
	return
}

//Delete object by its path. The format of path is "/bucketName/objectName".
func (c *Client) DeleteObject(opath string) (err error) {
	if strings.HasPrefix(opath, "/") == false {
		opath = "/" + opath
	}
	resp, err := c.doRequest("DELETE", opath, opath, nil, nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 204 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
	}
	return
}

//Download object by its path. The format of path is "/bucketName/objectName".
//If rangeStart > -1 and rangeEnd > -1, download the object partially.
//Return the object's byte array.
func (c *Client) GetObject(opath string, rangeStart, rangeEnd int) (obytes []byte, err error) {
	if strings.HasPrefix(opath, "/") == false {
		opath = "/" + opath
	}

	params := map[string]string{}
	if rangeStart > -1 && rangeEnd > -1 {
		params["range"] = "bytes=" + strconv.Itoa(rangeStart) + "-" + strconv.Itoa(rangeEnd)
	}

	resp, err := c.doRequest("GET", opath, opath, params, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 && resp.StatusCode != 206 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}
	//fmt.Println(string(body))
	obytes = body
	return
}

//Upload object by its remote path and local file path. The format of remote path is "/bucketName/objectName".
func (c *Client) PutObject(opath string, filepath string) (err error) {
	if strings.HasPrefix(opath, "/") == false {
		opath = "/" + opath
	}

	//reqUrl := "http://" + c.Host + opath
	buffer := new(bytes.Buffer)

	fh, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer fh.Close()
	io.Copy(buffer, fh)

	contentType := http.DetectContentType(buffer.Bytes())
	params := map[string]string{}
	params["Content-Type"] = contentType

	resp, err := c.doRequest("PUT", opath, opath, params, buffer)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}
	return
}

//Get object's meta information by its path. The format of remote path is "/bucketName/objectName".
func (c *Client) HeadObject(opath string) (header http.Header, err error) {
	if strings.HasPrefix(opath, "/") == false {
		opath = "/" + opath
	}
	resp, err := c.doRequest("HEAD", opath, opath, nil, nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		return
	}
	header = resp.Header
	return
}

type deleteObj struct {
	//XMLName xml.Name	`xml:"Object"`
	Key string
}

type deleteList struct {
	XMLName xml.Name `xml:"Delete"`
	Object  []deleteObj
	Quiet   bool
}

//Delete multiple objects by bucket name and keys.
func (c *Client) DeleteMultipleObject(bname string, keys []string) (err error) {
	dl := deleteList{}
	for _, v := range keys {
		dl.Object = append(dl.Object, deleteObj{v})
	}
	dl.Quiet = true

	bs, err := xml.Marshal(dl)
	if err != nil {
		return
	}

	reqStr := "/" + bname + "?delete"
	buffer := new(bytes.Buffer)
	buffer.Write(bs)

	h := md5.New()
	h.Write(bs)
	md5sum := base64.StdEncoding.EncodeToString(h.Sum(nil))
	params := map[string]string{}
	params["Content-MD5"] = md5sum

	resp, err := c.doRequest("POST", reqStr, reqStr, params, buffer)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}
	return
}

type muJob struct {
	File     *os.File
	Start    int
	Length   int
	Idx      int
	Opath    string
	UploadId string
}

func (c *Client) initMultipartUpload(opath string) (imur initMultipartUploadResult, err error) {
	resp, err := c.doRequest("POST", opath+"?uploads", opath+"?uploads", nil, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}

	err = xml.Unmarshal(body, &imur)
	return
}

func (c *Client) uploadDoWork(job *muJob) (part Multipart, err error) {
	buffer := new(bytes.Buffer)
	c.FileIOLocker.Lock()
	job.File.Seek(int64(job.Start), 0)
	io.CopyN(buffer, job.File, int64(job.Length))
	c.FileIOLocker.Unlock()
	h := md5.New()
	h.Write(buffer.Bytes())
	md5sum := fmt.Sprintf("%x", h.Sum(nil))
	md5sum = "\"" + strings.ToUpper(md5sum) + "\""

	reqStr := job.Opath + "?partNumber=" + strconv.Itoa(job.Idx) + "&uploadId=" + job.UploadId

	resp, err := c.doRequest("PUT", reqStr, reqStr, nil, buffer)
	if err != nil {
		fmt.Println(err)
		return
	}

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Printf("resp status:%s\n", err)
		fmt.Println(string(body))
		return
	}

	ETag := resp.Header.Get("ETag")
	if ETag != md5sum {
		fmt.Printf("ETag:%s != md5sum %s\n", ETag, md5sum)
	}
	part.ETag = ETag
	part.PartNumber = job.Idx
	return
}

func (c *Client) uploadWorker(jobs chan muJob, finishes chan Multipart, endWorker chan int) {
	var job muJob
	for {
		select {
		case job = <-jobs:
			part, _ := c.uploadDoWork(&job)
			finishes <- part
		case <-endWorker:
			break
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (c *Client) uploadPart(imur initMultipartUploadResult, opath, filepath string) (cmu CompleteMultipartUpload, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}

	buffer_len := 5 << 20
	fi, err := file.Stat()
	if err != nil {
		return
	}
	file_len := int(fi.Size())
	jobs_num := (file_len + buffer_len - 1) / buffer_len
	//fmt.Printf("jobs_num:%d\n", jobs_num)

	jobs := make(chan muJob, jobs_num)
	finishes := make(chan Multipart, jobs_num)
	endWorker := make(chan int, c.ChanNum)
	//start go
	for i := 0; i < c.ChanNum; i++ {
		go c.uploadWorker(jobs, finishes, endWorker)
	}

	//add job
	for i := 0; i < jobs_num; i++ {
		var job muJob
		if i == jobs_num-1 {
			last_len := file_len - buffer_len*i
			job = muJob{file, i * buffer_len, last_len, i + 1, opath, imur.UploadId}
		} else {
			job = muJob{file, i * buffer_len, buffer_len, i + 1, opath, imur.UploadId}
		}
		jobs <- job
	}

	//get finished
	for i := 0; i < jobs_num; i++ {
		var part Multipart
		part = <-finishes
		cmu.Part = append(cmu.Part, part)
	}
	mps := multipartSorter{cmu.Part}
	mps.Sort()
	cmu.Part = mps.Parts
	//end go
	for i := 0; i < c.ChanNum; i++ {
		endWorker <- i
	}
	return

}

func (c *Client) completeMultipartUpload(cmu CompleteMultipartUpload, opath, uploadId string) (cmur CompleteMultipartUploadResult, err error) {
	bs, err := xml.Marshal(cmu)
	if err != nil {
		return
	}

	reqStr := opath + "?uploadId=" + uploadId

	buffer := new(bytes.Buffer)
	buffer.Write(bs)
	//fmt.Println(string(bs))

	resp, err := c.doRequest("POST", reqStr, reqStr, nil, buffer)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}
	err = xml.Unmarshal(body, &cmur)
	return
}

//Upload large object.
func (c *Client) PutLargeObject(opath string, filepath string) (err error) {
	if strings.HasPrefix(opath, "/") == false {
		opath = "/" + opath
	}

	imur, err := c.initMultipartUpload(opath)
	//fmt.Printf("%+v\n", imur)
	imu, err := c.uploadPart(imur, opath, filepath)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = c.completeMultipartUpload(imu, opath, imur.UploadId)
	return

}

//Abort multipart upload by its path and uploadID.
func (c *Client) AbortMultipartUpload(opath, uploadId string) (err error) {
	if strings.HasPrefix(opath, "/") == false {
		opath = "/" + opath
	}

	reqStr := opath + "?uploadId=" + uploadId
	resp, err := c.doRequest("DELETE", reqStr, reqStr, nil, nil)
	if err != nil {
		return
	}

	if resp.StatusCode != 204 {
		err = errors.New(resp.Status)
		body, _ := ioutil.ReadAll(resp.Body)
		defer resp.Body.Close()
		fmt.Println(string(body))
	}
	return
}

//List uncompleted multipart upload in a bucket which named bname.
//The params is additional, it can be:"prefix","marker","delimiter","upload-id-marker","max-keys".
//Return an object of ListMultipartUploadResult.
func (c *Client) ListMultipartUpload(bname string, params map[string]string) (lmur ListMultipartUploadResult, err error) {
	if strings.HasPrefix(bname, "/") == false {
		bname = "/" + bname
	}

	reqStr := bname + "?uploads"
	if params != nil {
		for k, v := range params {
			reqStr += "&" + k + "=" + v
		}
	}

	resp, err := c.doRequest("GET", reqStr, reqStr, nil, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}
	//fmt.Println(string(body))
	err = xml.Unmarshal(body, &lmur)
	return
}

//List uploaded parts for a multipart uploading.
func (c *Client) ListParts(opath, uploadId string) (lpr ListPartsResult, err error) {
	if strings.HasPrefix(opath, "/") == false {
		opath = "/" + opath
	}

	reqStr := opath + "?uploadId=" + uploadId
	resp, err := c.doRequest("GET", reqStr, reqStr, nil, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}

	err = xml.Unmarshal(body, &lpr)
	return
}

//Put the CreateFileGroup to the server.
//The format of path is "/bucketName/objectGroupName".
//Return an object of CompleteFileGroup.
//Notice:objects and objectgroup should in the same bucket.
func (c *Client) PostObjectGroup(cfg CreateFileGroup, gpath string) (completefg CompleteFileGroup, err error) {
	//part := []GroupPart{{1, "11", "111"}, {2, "22", "222"}, {3, "33", "333"}}
	//fg := CreateFileGroup{Part:part}
	bs, err := xml.Marshal(cfg)
	if err != nil {
		return
	}

	if strings.HasPrefix(gpath, "/") == false {
		gpath = "/" + gpath
	}

	reqStr := gpath + "?group"
	buffer := new(bytes.Buffer)
	buffer.Write(bs)

	resp, err := c.doRequest("POST", reqStr, reqStr, nil, buffer)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}
	err = xml.Unmarshal(body, &completefg)
	return
}

//Get objectgroup's index by its path.
//The format of path is "/bucketName/objectGroupName".
//Return an object of FileGroup
func (c *Client) GetObjectGroupIndex(gpath string) (fg FileGroup, err error) {
	params := map[string]string{"x-oss-file-group": ""}
	if strings.HasPrefix(gpath, "/") == false {
		gpath = "/" + gpath
	}
	resp, err := c.doRequest("GET", gpath, gpath, params, nil)
	if err != nil {
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
		fmt.Println(string(body))
		return
	}

	err = xml.Unmarshal(body, &fg)
	return
}

//Get objectgroup by its path. The format of path is "/bucketName/objectGroupName".
//The usage is the same as GetObject.
func (c *Client) GetObjectGroup(gpath string, rangeStart, rangeEnd int) (obytes []byte, err error) {
	return c.GetObject(gpath, rangeStart, rangeEnd)
}

//Get objectgroup's meta information by its path. The format of path is "/bucketName/objectGroupName".
func (c *Client) HeadObjectGroup(gpath string) (header http.Header, err error) {
	return c.HeadObject(gpath)
}

//Delete objectgroup by its path. The format of path is "/bucketName/objectGroupName".
func (c *Client) DeleteObjectGroup(gpath string) (err error) {
	return c.DeleteObject(gpath)
}

func newValSorter(m map[string]string) *valSorter {
	vs := &valSorter{
		Keys: make([]string, 0, len(m)),
		Vals: make([]string, 0, len(m)),
	}

	for k, v := range m {
		vs.Keys = append(vs.Keys, k)
		vs.Vals = append(vs.Vals, v)
	}
	return vs
}

func (vs *valSorter) Sort() {
	sort.Sort(vs)
}

func (vs *valSorter) Len() int {
	return len(vs.Vals)
}

func (vs *valSorter) Less(i, j int) bool {
	return bytes.Compare([]byte(vs.Keys[i]), []byte(vs.Keys[j])) < 0
}

func (vs *valSorter) Swap(i, j int) {
	vs.Vals[i], vs.Vals[j] = vs.Vals[j], vs.Vals[i]
	vs.Keys[i], vs.Keys[j] = vs.Keys[j], vs.Keys[i]
}

func (ms *multipartSorter) Sort() {
	sort.Sort(ms)
}

func (ms *multipartSorter) Len() int {
	return len(ms.Parts)
}

func (ms *multipartSorter) Less(i, j int) bool {
	return ms.Parts[i].PartNumber < ms.Parts[j].PartNumber
}

func (ms *multipartSorter) Swap(i, j int) {
	ms.Parts[i], ms.Parts[j] = ms.Parts[j], ms.Parts[i]
}
