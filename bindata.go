// Code generated by go-bindata.
// sources:
// templates/jenkins/multi-job.xml
// templates/jenkins/normal-job.xml
// templates/jenkins/pipeline.xml
// DO NOT EDIT!

package pipeline

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _templatesJenkinsMultiJobXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x18\x69\x6f\xdb\x38\xf6\x73\xfd\x2b\x8c\xa0\x40\x80\x62\x23\x25\xd9\xdd\xb6\x0b\x28\x6e\x1b\x27\xed\x7a\x90\xc3\xe3\xa3\xf9\x58\xd0\x12\x2d\xb3\xa6\x44\x81\xa4\xd2\x78\x82\xfe\xf7\x79\xbc\x24\xea\x48\x93\x74\x8a\xc1\x24\x1f\x2c\xbe\x8b\xef\x7e\x4f\x8a\xde\xdd\x65\x74\x78\x8b\xb9\x20\x2c\x3f\xd9\x3f\x0a\x0e\xf7\x87\x38\x8f\x59\x42\xf2\xf4\x64\x7f\xb9\xf8\x78\xf0\x76\xff\xdd\x68\x10\xc5\x2c\x0b\x24\xd9\x22\x1a\x7c\xc5\xf9\x96\xe4\x22\x28\x68\x99\xaa\xdf\xac\xa4\x92\x7c\x65\xab\xe0\x52\x3d\xfc\xc6\x56\x53\xce\xbe\xe2\x58\x0e\x0d\xc1\xc9\x9e\x65\x38\x70\x84\x07\x06\xf1\xfe\x28\x38\xfa\xef\xde\x68\x30\x1c\x46\x28\x96\x70\xbb\x08\xf5\x21\xc1\x22\xe6\xa4\x50\x90\x51\x14\xfa\x27\x85\xdd\x62\x5c\x9c\xe1\x02\xe7\x09\x68\x49\xb0\x18\xad\x11\x15\x38\x0a\x3b\x70\x45\x5c\x70\x56\x60\x2e\xed\x11\x00\x02\x07\x09\x41\x2b\xe6\x59\x41\x0a\x4c\x49\x8e\x83\xa9\x7d\x98\x1a\x9e\x5d\xa5\x7e\x02\x60\xf0\xcf\xee\xc0\x91\x3a\xfd\x0f\x83\xb7\xc1\xeb\x3d\x23\x19\x64\x4b\x24\xb6\x57\x28\xc3\xa3\x02\x71\x44\x29\xa6\x43\x7c\x87\xe3\x52\x69\x1e\x85\x15\xd2\x51\x0b\x89\x52\xac\x21\xf7\xf7\xc3\x60\xee\x4e\xc3\xef\xdf\xa3\xb0\xc6\x19\xad\xc3\x67\xa8\x6d\x59\x54\xb8\x62\xc6\xd1\xed\x2e\xd8\x94\x89\x60\x79\x15\xae\x94\xc8\x4d\xb9\x0a\x3e\xe9\x1f\x1b\xaa\x8e\xcd\x86\x48\x45\xe8\x38\x38\xac\x4d\x2c\x0c\xf9\x92\x53\xad\xb5\x91\x01\x27\xad\xb5\x87\xb4\x6a\xff\x9c\x12\x3a\x72\x61\x33\x74\x91\x88\xb3\x61\x4c\x91\x10\x27\x7b\x5d\x51\x4a\xce\x7c\x7c\xb9\xe7\xab\xff\xfe\x38\xf8\x8f\x56\xfd\xc5\x0b\x70\x46\xbe\x26\xe9\x67\x93\xe3\xa3\x63\xa5\x98\x0f\x50\x24\xa5\xc0\x7c\x86\x33\x26\xf1\x58\xe3\x84\x82\xbe\x88\x7a\xee\x5a\xb6\x28\x35\x21\x08\xa8\x7d\xb2\x9c\x5d\x68\x87\x28\x90\x96\x12\x3e\x51\x0c\xb0\xf4\xa9\x11\xad\x38\xca\xe3\x0d\x7e\x58\xa7\x53\x4d\x30\x2f\x70\x6c\xb5\xc9\x55\xf2\xbc\x0a\x33\x24\x24\xe6\x51\xa8\x8f\x83\x07\x34\x69\x32\x47\xa1\x7f\x5b\x94\x10\x81\x56\x14\xcf\xcb\x55\xc6\x92\x92\xd6\xf5\xd6\x45\x98\xa0\x73\x48\x79\x70\xeb\xad\x8f\x91\xbc\x04\x8e\x3e\x8c\xe1\x49\xd8\x27\x9c\x63\x8e\x64\x8d\x32\xe6\x97\x00\x53\x8d\xa1\xba\xf4\x51\x42\x23\x10\x95\x72\xc3\xf8\x35\x1f\xb3\x2c\x23\x12\x5c\xe0\x04\x74\x11\xb6\x5c\x28\x46\xb9\x8e\xdf\x58\x3d\xdd\x30\xbe\x15\x05\x8a\x4d\x39\x1a\xa4\x21\xfc\x06\xe5\x76\x5d\xca\x8a\xc0\x09\xee\xc0\x07\xb6\x60\xca\x1c\x9f\x3a\x87\x5a\xda\x26\xd0\xb9\x4d\x05\x7d\xca\x28\x75\x54\x1e\xc4\x90\x90\x34\x67\x1c\x5f\x31\x49\xd6\x3b\xa3\xbf\x23\xed\xc1\x18\x16\x48\xa7\xf9\x06\x7a\x11\xfb\x36\xa6\x2c\xaf\x94\x6d\x83\x0d\xf1\xaa\x24\x34\x19\x6f\x18\x83\x1c\xfc\x41\xad\x41\x43\xa3\xc1\x19\x5e\x23\xe8\xe6\xa7\x1e\xcb\x5e\x68\xe5\x00\xd1\x82\x31\x3a\xb2\x34\x51\xe8\x00\xb6\x05\x57\x91\x5b\xa7\xee\x1a\x4a\x84\xac\xf8\x39\xa6\x10\xcc\x5b\xbc\x40\x3c\xc5\xf2\x8c\xf0\x1a\xb1\xc6\x1c\xba\x3b\x76\x00\x7c\x17\xd3\x32\xc1\xc9\x0c\xa7\xd5\xf8\xf0\xc0\xaa\xc4\x84\xa7\x95\x49\x15\xd5\x57\x3b\xc0\xf3\x0c\x11\xea\xa0\x62\x4b\x8a\x05\x4a\x6d\xda\xba\x93\xc5\xc5\x99\x2f\xe0\x97\x67\xee\xa3\xde\xc1\x77\x12\xe7\xc2\x9a\x0b\xd5\x0a\x0a\x8d\x06\xaa\x43\xc6\x28\x9f\x31\x94\x59\xad\xdd\x49\x0f\x54\x53\xaa\x49\xab\x74\x13\x8d\x5c\x51\x16\x6f\x75\x1c\x6f\x36\x38\x3f\x63\xdf\x72\x21\x39\x46\x99\x06\xc1\x0e\xe0\x98\x1e\xa5\xeb\x0a\x5b\x16\x4f\x11\xd5\xa1\x52\x82\x24\x27\x69\x0a\xd1\x33\x46\x43\x5d\x92\xf5\x30\x98\x88\x49\x4e\x24\x41\x14\x76\x0c\x28\x4c\xe3\x0f\x9b\xa1\x8e\x21\x80\x39\xb0\x30\xcf\xf5\xa8\x55\xad\xed\xd5\xd0\xfe\x83\xcb\x74\xab\xb3\x48\x53\x3b\x53\x26\xa4\xa9\x9c\xff\x33\xb6\x15\xcd\xc2\x6a\x23\xed\x74\x7b\xec\x66\xd0\x1a\xb6\x11\xa3\x28\x2c\x00\x9e\x45\x6a\x20\x41\x2f\x84\x5c\x36\x15\xe4\xae\x6b\x83\x07\xae\x2a\x2b\x4f\x3c\x67\x0d\x3b\x35\x9c\xf5\xf0\xde\x20\x61\xb6\x0a\xfd\x74\x08\x8d\xa8\x82\x34\x68\x80\x57\x38\x88\x36\x03\x5a\x55\x8a\x61\x49\x29\x57\x0a\xe5\x5c\xff\x54\x7d\xa6\x4e\xa6\x1b\x74\xc3\xea\x2f\x02\x7c\xb5\x03\xe9\x5e\xeb\x00\x3e\x91\xf2\xc9\x14\x76\xaa\xcc\x4d\x12\x0f\xe0\xd3\xe1\xbb\x02\x1a\x51\x02\x71\x70\x0e\xf5\x20\x3e\xa1\x2d\x00\xd0\xa9\x55\x12\x0a\xe2\x13\xc2\x26\x27\x20\x25\x67\x6a\x56\x4d\x91\xdc\xc0\x3e\xda\x01\xf9\xf4\x19\xba\x9b\x61\x88\x34\xb4\x75\x70\xaf\x77\x6a\xa8\x99\xab\xab\x14\x66\x37\x97\xaa\x23\xa4\xbb\x4a\xdf\x1e\x54\x97\x15\xfc\x98\x10\xbd\x10\x37\xd8\x6a\xb0\xcf\x02\x5b\x23\x97\x1f\xa8\xaa\x19\xeb\x3d\x1f\xd2\x70\x73\xc5\xaf\x33\xb1\x47\xd6\x96\x50\xaa\xa3\x79\x9d\x03\xf3\x0c\x0b\x08\x71\x7d\xeb\xc7\x0f\x93\x8b\xe5\xec\x1c\x56\xf1\x1f\x92\xf9\x02\x75\x72\x5f\xe7\x74\x37\x59\x43\x8c\xc6\x1b\x95\x67\x55\xf1\xf5\x23\x5b\x1a\x9b\x1d\x69\xe8\xfd\xf5\xed\x47\xb0\x94\xcd\xf0\x2d\x51\x6d\x53\x57\x85\x4e\x1e\x0c\xe3\x5f\xf4\xae\x8c\x0d\x71\x90\xe1\x2b\xd8\xaf\x7f\x2f\x71\x89\x13\xd3\x06\x44\x5d\xb0\x3d\xb8\x26\x7b\xdf\xc6\xf5\xb0\x3a\x0d\xf3\xc2\x8e\x7d\x66\xa5\xfe\xd9\x7a\xf3\x1b\x92\x91\xd6\x29\x77\xe5\x52\x49\xf2\x52\xcf\xa4\x3a\x66\xf3\xe5\x78\x7c\x3e\x9f\x7f\x5c\x5e\x68\xa5\x7a\x28\x06\x4f\x56\xaf\xa7\x3d\xd9\x58\xbb\xee\x58\x94\x2b\x18\x79\x9b\x76\xfb\xbf\x82\xd1\x77\x89\xe0\x66\xea\x77\x21\xd8\xf5\x02\xfd\x9e\x01\x5d\x93\x97\x99\xd8\xc1\xc6\x9b\x89\xe6\xfb\x46\xa0\xc5\x57\xaf\x4c\xb6\x15\x07\xc6\xf7\x16\x6a\x5b\x77\x95\x0f\x9a\xa5\xf3\xca\x77\x04\x19\xf2\xef\x3d\xdf\x5b\xad\x00\xfd\xca\xe4\xfb\x0b\xa9\xf7\x93\x69\xd7\x49\x39\x58\x70\xdc\xa0\xb7\x6f\x69\xaa\x3b\x0b\xdd\xaf\x3b\xf1\x50\x5b\x4e\x2f\xb5\x4d\x8e\x5f\x1b\xa9\xee\x90\x6d\x64\x4a\x23\x47\x5a\xce\x28\x9c\xe1\xe4\x0f\x9c\x34\xee\x68\x67\x41\x83\xf2\xc0\x92\x42\xa8\x8e\xff\xf7\x8c\x24\x78\xf4\xba\x9e\xc1\xf8\x77\xb4\xb6\x7f\x58\x73\xab\xbf\x2f\xd4\xf9\x55\x65\x56\x85\xe9\x9f\x58\xb6\x41\x3d\x34\xb8\xac\xd3\x6f\x88\xdc\x5c\x31\x4f\x21\x6b\xe9\x43\xe8\x87\xcb\xe9\x99\x21\x6d\x59\xfb\x7c\x71\x7d\x1b\x65\xb3\x4d\x9a\x59\x7a\xc3\x51\x51\xd4\xdb\x62\xeb\x1a\x04\x2f\x0e\x31\xa3\x8c\x07\x1f\xe0\x69\xac\x9e\x4e\x3d\xae\x2a\x61\x2a\xba\xf7\x87\x8d\xb4\x89\x34\xf0\x12\x15\x7a\x45\x83\x17\x11\x9e\x29\xcb\x3c\x58\xbf\x79\x8f\xdc\xfb\xf8\x92\xff\x84\xef\x5f\xf6\x63\x0e\xf8\x1d\xfc\xb7\x2a\x25\xe3\x4f\xfb\x80\xf7\xa6\xb6\xce\x7e\x04\x5d\xe0\xac\x80\x37\x50\x3c\x7a\x79\x7f\xba\x9c\x5c\x9c\x7d\xb9\x5a\x5e\x9e\x9e\xcf\xbe\x1f\xbc\xbc\xff\x34\x59\x7c\x99\x9d\x7f\x9e\xcc\x27\xd7\x57\xff\xa2\x38\x4f\xe5\xe6\xe4\x0d\x64\x67\x9b\xd3\x49\x2c\x8b\x04\x8e\x67\x44\x00\x78\xa7\x1d\x64\x56\xb0\x2e\xfc\xe9\x1f\xfa\xba\x86\xf6\x75\xc2\x76\xe8\x25\x81\x46\x2c\x51\x06\xee\x0e\x16\xf5\x73\x6f\xf8\x3d\x5a\x18\x78\xaf\x83\x63\xf3\xf2\x69\xa7\x74\x9d\x62\xcf\x1a\xf7\x76\x20\x8c\x06\x7f\x06\x00\x00\xff\xff\x96\x02\xac\xa2\x71\x16\x00\x00")

func templatesJenkinsMultiJobXmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesJenkinsMultiJobXml,
		"templates/jenkins/multi-job.xml",
	)
}

func templatesJenkinsMultiJobXml() (*asset, error) {
	bytes, err := templatesJenkinsMultiJobXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jenkins/multi-job.xml", size: 5745, mode: os.FileMode(420), modTime: time.Unix(1444386462, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJenkinsNormalJobXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5b\x7b\x6f\xe3\x36\xb6\xff\x3f\x9f\xc2\xc8\x1d\x60\xee\x2d\x6a\x79\x92\x99\xdb\x69\x01\xc7\x4d\x62\x3b\x33\x69\x1d\xc7\xd7\x76\xa6\x77\xb1\x58\x04\xb4\x44\xdb\x6a\x28\x51\x10\xa9\x24\xde\xd9\xf9\xee\x7b\xf8\x14\x45\xc9\x76\xd2\x29\x0a\x0c\x36\x2d\xd0\x5a\xe7\x45\xf2\xf0\x9c\x1f\x0f\x1f\xe9\xfe\xfc\x98\x90\xd6\x3d\xce\x59\x4c\xd3\x93\xd7\x47\xc1\x9b\xd7\x2d\x9c\x86\x34\x8a\xd3\xd5\xc9\xeb\x9b\xf9\x45\xfb\xc7\xd7\x3f\xf7\x0e\xba\x59\x4e\x7f\xc7\x21\xef\x1d\xb4\x5a\x5d\x14\x72\x10\x66\x1d\xf9\x11\x61\x16\xe6\x71\x26\x28\x8a\x40\xe8\x6a\x4a\x39\xe2\x34\x6f\x85\x04\x31\x76\x72\xb8\x2e\x22\x46\xd3\x80\x23\x76\xc7\x82\x91\x65\x1f\x0a\x71\x61\x01\x6d\xd8\x9c\xfe\x8a\x71\xd6\x6b\x1f\x75\x3b\xce\xa7\xe2\xa7\x45\xe2\xb0\xcb\x2f\xc5\x45\x39\x8f\x97\xd0\xa3\x41\xd5\x4a\x03\xb9\x2a\x3f\xb6\x76\x1c\xe9\xb1\x6b\xbc\xdb\x29\x87\x22\xbf\xef\x80\x31\xc0\x19\x4e\x23\xf0\x50\x8c\x59\x6f\x89\x08\xc3\xdd\x4e\x8d\x2e\x84\xc1\x61\x19\x06\xb3\xea\x13\x08\x0c\x07\x51\x8c\x16\x94\x04\xbf\xe3\xf4\x2e\x4e\x59\x90\xc5\x19\x26\x71\x8a\x83\x89\xfe\x31\x51\x3a\x9b\x56\x46\x8a\x55\x9c\x9e\x1c\x46\x40\x86\xb9\xd9\xb4\x8d\x68\x5b\x71\x4e\xdf\x04\x3f\x06\xef\xb5\x03\xc1\xb6\x70\xed\x18\x25\xb8\xf7\xf9\x73\x2b\x98\xeb\x8f\xd6\x97\x2f\xdd\x8e\xe5\x18\x51\xc6\xd1\x0a\x5b\xd9\x99\xf9\x92\xc2\x25\x4f\x75\xb9\xf3\x8c\x3e\x6b\x95\x90\x26\x41\x48\x73\x74\xbf\x09\xf4\xb4\xab\x1e\xb3\x60\x15\xf3\x75\xb1\x08\x3e\xc8\xff\x4d\x54\x38\xd5\x06\xac\x84\x4e\x8f\x82\xa3\xe3\xe0\x4d\x39\x3e\x1d\x7d\x37\x39\x91\xbd\x56\x36\xe0\x4b\xf6\xda\x61\xea\x6e\xff\xb1\x4e\xa8\x39\xaf\xce\x1b\xb4\x16\x2f\x5b\xc1\x25\xbb\x4c\x63\x1e\x23\xf2\x0b\x5d\x40\x9b\x42\x90\x85\x89\x17\xdf\x4e\x1b\xa2\x81\x59\xff\xea\xd0\x1d\xd7\xe9\x71\x70\x1c\xfc\x74\x68\xfd\x94\x2e\xe3\xd5\x27\x95\x77\xbd\x63\xd1\x67\x97\xa0\x84\x0a\x86\xf3\x29\x4e\x28\xc7\x7d\xc9\x65\xd6\x21\x0d\x4d\xde\x78\xc2\x46\x56\xd8\x29\xfd\x76\x33\x1d\x49\xa7\x15\xc6\x5b\x62\xcc\x4f\xb6\x06\x7a\xcd\x5d\xea\x2e\x72\x94\x86\x6b\xbc\xb3\x87\xe7\x52\x66\x96\xe1\xd0\xe9\x5b\x2a\xc2\xed\xbb\x8e\xe8\x5e\xc9\x8f\x97\x31\xce\x65\x3f\x53\x37\x76\x9b\x3a\xea\x1b\xed\x76\xaa\x5d\xe9\x46\x31\x43\x0b\x82\x67\xc5\x22\xa1\x51\x41\xca\xb4\xad\x33\x94\x42\x8e\xc3\x02\x66\xe1\xde\xe5\xf0\xbc\x00\x8d\x26\x8e\x6e\x84\x7e\xc0\x29\xce\x11\x2f\x59\xca\x3f\x05\xd0\x04\x58\xda\x46\xf7\x0a\x6a\x98\x2a\xf8\x9a\xe6\xd7\x79\x9f\x26\x49\xcc\x39\xce\x8d\x81\x3a\x43\x07\x14\xc1\x28\x95\xb3\xdc\x17\xbf\x7e\xa3\xf9\x1d\xcb\x50\xa8\x12\x5b\x31\x95\xe0\x03\x24\xee\x75\xc1\xad\x80\x31\x5c\xa3\x1f\xe8\xd4\x2b\x52\x7c\x6e\x5c\xaa\x65\xab\x44\xe3\x36\x11\x15\x13\x4a\x88\x91\x72\x28\x4a\x24\x5e\xa5\x34\xc7\x63\x0a\x60\xbb\x51\xfd\x37\xa2\x0d\x1c\x9b\x02\xb3\x35\x22\x84\x3e\xf4\x09\x4d\x6d\x67\x7d\xb2\x0e\xc2\x22\x26\x51\x7f\x4d\x29\x04\xe9\x8e\xe4\x2c\x78\x4c\x82\x01\x5e\xa2\x82\xf0\x73\x47\xe5\xb0\xa3\xed\x80\xd0\x9c\x52\xd2\xd3\x32\xdd\x8e\x21\x68\x24\xb7\x33\xb7\x5c\x99\x66\x48\xcc\xb8\xd5\xcf\x31\x81\xc9\xbc\xc7\x73\x94\xaf\x30\x1f\xc4\x79\xc9\x58\xe2\x1c\x16\x09\x6c\x08\xf8\x31\x24\x45\x84\xa3\x29\x5e\xd9\x25\xd5\x21\x8b\x2c\x64\x4e\xaf\x54\xa8\x08\x84\xae\x11\x87\x09\x8a\x89\xa1\xb2\xbb\x38\x9b\xa3\x95\x0e\x5b\xf3\xa5\x27\x21\xad\x34\xd9\xd3\x10\x27\xe6\x1d\xd6\x7c\xe8\x2c\x44\x8c\x08\x23\x9f\x00\xcb\x9b\x8c\x25\x5f\x5f\xb7\x18\x26\x6e\xb7\xfe\xf4\x7c\xd8\xeb\x73\xfc\xc8\x71\xca\xca\xba\xa4\x03\x3d\x92\x3f\x78\x1e\xaf\x56\xe0\x46\x25\xa7\x47\x7b\x93\x31\x9e\x63\x94\x00\x9e\x33\x05\xe8\x20\x69\xd6\x38\xa3\x11\x4c\xb1\x28\x8c\xb0\x8c\x91\xb9\x22\x96\xeb\x28\x20\x4e\xc7\x7e\x15\xda\x9e\x5e\x52\xa4\x57\xfd\x46\x20\x6a\x7d\x29\xbb\x7e\xaf\x73\xcc\xd6\x94\x44\x3e\x30\xce\x6e\xfa\xfd\xe1\x6c\x56\xc5\x41\x60\xd2\x1c\x0a\x34\x44\x7a\x6f\xba\x1d\xf3\xb3\x64\x86\x94\x40\xc5\x72\x3e\xba\x19\x8a\x65\x85\xd0\xbc\xc2\x4b\x32\x82\xb9\x1a\x93\x8e\x8f\x2a\xcd\x82\xad\xd7\xa9\x6e\xe7\xe9\xfe\xb1\xe1\xe2\xfa\x5c\x4b\xe0\xe8\x0a\xa5\x05\xe4\xee\x46\xc7\x15\xc4\x81\x9d\x02\x53\x2a\x9a\x16\x60\x11\x6d\x72\x7c\xef\xbb\x96\xfe\x17\xe6\xd9\x59\x4f\x34\xc2\x4c\x28\xe3\x0a\x45\x3e\x52\x7a\xc7\xaa\x20\xe3\x33\x0f\x2a\x4b\xcb\xd6\x96\xdd\x21\x81\x6f\x9c\xa0\x52\xb0\xf3\x5b\x8e\xb2\xcc\x86\x99\xbf\xfc\x21\x08\x4d\x39\x15\xc1\x19\xfc\xea\x8b\x5f\xe7\x8e\x96\xad\x13\xac\x1c\x54\x78\xef\xdc\x0a\x48\x12\xaf\x50\x26\xab\x33\x08\xf5\x3c\xd1\x73\x6b\x68\x07\x8d\x2b\xe4\x9e\x76\x9f\x5e\x9a\xea\xc2\x04\x72\x13\x86\xbe\x28\x44\x75\xff\xdc\x22\x55\x6f\x32\xe6\x18\xa2\x0d\xf2\xbd\xf7\xea\xf3\xf9\xcd\xe5\x68\x70\x3b\xbe\xb9\x3a\x1f\x4e\xbf\xb4\x5f\x7d\xfe\x70\x39\xbf\x9d\x0e\x3f\x5d\xce\x2e\xaf\xc7\xdf\x13\x9c\xae\xf8\xfa\xe4\x3d\xe4\x8d\xaf\x59\x66\x5d\x04\x9f\x83\x98\x01\x79\x23\xbd\xa0\x02\xba\x4e\x7f\x7a\x3d\x5b\x1f\x68\xf3\x8c\xf2\x38\xc1\x50\x2f\x27\xe0\xc5\x60\x5e\xfe\x6e\x9c\x55\x47\x16\x4a\xdb\x1f\x82\xe3\x43\x8d\x51\xb5\xc8\xa9\xe4\xc3\x37\x55\x63\x4a\x80\xaa\xe1\x14\x2c\x78\x32\x3f\x45\x31\xb0\x64\xd5\xca\xef\xa5\x2a\x7d\xa9\x4a\x5f\xaa\xd2\x97\xaa\x54\x4d\x42\xda\xd8\xa4\x5b\x56\x96\x25\xdd\xb7\xbd\xe0\x6e\x81\x7d\x55\x5c\x94\x67\x0e\x33\x82\xee\xf1\x08\x2d\x30\xd1\xab\x01\xcc\x28\x04\x24\x8e\xc6\x34\xd2\x87\x37\xae\x04\xe4\xa0\xcb\x3f\xf0\x2a\x96\x10\xa5\x53\x8a\x12\x13\xa9\xe6\xf3\xa0\x84\xa2\xc8\x43\xa0\x48\xb9\x9a\xd0\xf0\x4e\x8d\x60\x8d\xd3\x01\x7d\x48\x55\x09\x2b\x49\xb0\x3d\x30\x4a\x7b\xe5\xea\xc6\x4c\x79\xbc\xdb\x54\x4d\xea\x40\xad\x6c\x80\x79\x10\xb3\x2a\x53\xec\xa8\x3c\xf2\xc1\x8e\x7a\xff\x19\xd5\xbe\x5b\xeb\x7f\x55\xa5\x5f\xaf\xf3\x77\x54\xf9\x3b\x6a\xfc\xed\x15\xfe\x53\xeb\x7b\xaf\xba\x7f\x7a\x6d\x5f\x09\x2a\x19\xc7\xee\xd6\x0a\xb0\x77\x85\x5b\xc1\x99\x39\x7c\xc5\x99\x5f\xd9\x9b\x3c\x09\x69\xb6\x31\xa7\xae\x41\x1f\x3e\x8c\x8e\xcd\x4c\x57\x02\xaa\xa6\xb7\xc7\xc1\x51\xed\x40\x50\x3a\x5f\xfb\x58\x24\xa5\xac\x10\xdd\x43\x41\xab\xb0\x8c\x89\x58\x95\x84\xbc\x6d\x49\xc8\x69\xba\xdd\x88\x49\x6c\x2c\xf7\x75\x1a\xf2\x58\x49\x61\x98\x80\xdd\xda\xe1\x76\xe3\xb8\xec\x8e\x47\x7a\x71\xa6\x35\x0f\x9d\x3a\x03\x02\x97\x2c\x50\x78\x37\xa7\x23\xc4\xf8\xac\x08\x43\xcc\xd8\xb2\x20\x7a\xe6\xb6\xb2\x9d\x2a\x4a\x47\xda\x85\x1c\xc8\x8c\x8b\x45\x7c\xb5\xe9\x01\x46\x7f\x20\x74\x81\xc8\x0c\x73\x0e\x69\x53\x86\xa4\x27\x68\xab\x18\x33\x30\x4b\x89\x28\xac\x7c\x17\xa0\x8b\xf3\x2c\x8f\x53\x6e\xfc\xe6\x54\x0d\xdb\x04\x9a\xb1\x71\xeb\x9c\x37\x6d\x1a\xbb\x95\x5b\x83\xd9\x1a\x13\x37\xfc\x93\x04\xa5\x10\xb9\xff\xd5\x32\xf5\x7b\x2b\x84\xb1\x89\x69\x49\xd0\x2a\x0e\x0f\xf0\x63\x46\x73\xde\x9a\x5c\x4e\x86\xa3\xcb\xf1\x50\xef\x31\x4e\x5e\xfd\x37\x0e\xd7\xb4\x75\xf8\xea\xb3\xe5\x7c\x1a\x4e\xc5\x76\xe3\xcb\x61\xeb\x5f\xad\xb0\xe0\xad\xf6\xf2\xa8\xd5\x8e\x5e\xb7\x5f\xff\x8f\x31\x22\x76\x25\xb3\x8f\x67\x4f\x53\x3e\xd6\xca\xcd\x07\xc7\x4e\x79\x0f\x4b\x64\x0b\x32\x10\x83\x52\x7b\x8d\xf2\xa8\xa5\xb6\x3f\xd0\x50\x79\xd8\x72\x20\x0b\x27\x35\x58\xe3\x97\x96\x3a\xdf\x56\xe3\xaf\xb8\xd9\xf3\x94\x5e\x66\xcc\x26\x35\x2b\x16\x50\x0d\xac\xfd\x83\x90\x8f\x3c\x21\x53\x2c\xc6\x59\x9e\x83\xac\x81\x66\xc5\xa5\xc4\xc4\x7c\xd9\xec\xac\xc8\x40\x7a\xfe\x6f\x99\x9b\xb9\x34\xa7\xaa\x0c\x56\x46\x6a\x09\x0f\x4e\x9b\x81\xd7\xf6\xbe\xf6\x95\xd5\xd2\xa8\x6d\xce\xde\x68\x4c\xed\xa7\xcc\x6f\x87\x5b\x57\x82\x1a\xc8\xd1\x51\x27\x5d\x46\x45\xf0\xea\x1a\x90\x3c\x98\x39\x3a\xf2\xdb\xd1\x52\x7c\x57\x0f\x91\x07\xb4\x61\xa3\x38\xd5\x49\x5c\x59\xb0\x9a\x99\xae\xba\xb8\x5b\x3a\x23\x06\x0f\xcc\x57\xb5\x01\xa8\x37\xaf\x62\x58\xfc\x53\x53\x5a\x55\x48\xa5\x5f\x3b\xcf\x70\x6c\x35\x19\x85\x72\xc3\xb4\xee\xb4\xb8\xe3\x24\xe8\x2c\x8d\x72\x1a\x47\x30\x6e\x1e\x4c\x90\xd8\x2e\xa4\x36\xf8\x68\xbe\x32\x9b\xf3\x30\x76\xca\x2a\xa9\x71\x7b\x4b\x84\x8e\x50\xac\x87\xa4\x96\x69\x0b\x11\xb1\x11\x2e\x23\x72\x8d\x11\xe1\xeb\x8d\x73\x5c\x97\x7e\xf4\x49\x76\x3d\x1c\xc5\xa2\xce\x07\x07\x3a\x6b\xa4\xa2\xd9\xc5\x47\x36\x28\x63\xea\xef\x6e\xa3\xff\x80\x65\xa7\x64\x59\x18\x55\xe5\xfa\x50\x5f\xaa\x96\x2d\x8a\x1a\xac\x48\xaf\xd3\x0b\x28\x91\xcb\xea\xcb\xa3\xda\x1e\x33\x3c\xc9\xf1\x7d\x4c\x0b\x26\x83\xe4\x8c\x4d\x4d\xd5\xee\xec\x3e\xb6\x8a\x38\x66\x66\x5c\x54\x77\x3b\x8c\x6c\x11\x70\x4c\x0c\x30\xe1\xe8\x13\x22\x45\xb9\xf7\xf2\xa8\x35\xaf\x32\x67\x9a\x10\xd9\xb0\x98\xb5\x43\xd8\x5b\x01\x72\xbc\x77\x26\xaa\x80\xaa\x51\x34\x3e\xa7\x1c\x11\x08\xf4\x4e\x33\xe7\x63\xbc\x5a\x6f\x61\x8d\x69\x9e\xa0\x6d\x7a\x23\xfa\x50\x72\x96\xd2\xc1\xf5\x96\x1c\xba\x68\x47\x95\x0c\x4e\xc0\x5e\x54\xf9\xaa\x8a\xf0\x74\x1a\x6c\xa9\x8e\xed\xb2\xa6\x24\x7c\x7b\x5a\xaf\xc1\x62\x65\x34\x65\xac\x96\xce\x87\xcf\x82\x44\x03\xa8\xfc\x42\x7e\x55\x3d\x65\x68\x62\x95\xeb\x7e\x2a\x0e\x48\xb3\x82\xe3\x31\x7e\xd0\x90\xe2\x11\x2b\x35\xc2\x14\x33\x4a\xee\xf1\x54\xef\x2f\x21\xa3\xd7\xd5\x1a\xa1\x51\xc0\x66\x93\x02\x80\x9a\x67\x4a\x60\x80\xb4\xd2\x32\x1a\x75\x9e\x8d\x11\x3b\xb0\x08\xea\x97\x68\x51\xac\x98\x69\x6f\x4b\xbd\xba\x34\x62\x42\xfe\x1c\x7e\xd4\xf1\xc7\x88\x9c\xbe\x0b\x7e\x38\xfa\x4b\xc0\xe7\xe2\x72\x3c\x38\xbf\xf9\x30\x7b\x01\x9e\xaf\x04\x9e\x5d\xd0\xb3\x13\x7c\xf6\xc0\xcf\x2e\x00\xda\x0e\x41\x75\x10\xda\xc2\xaa\xb7\xf6\xcd\xc0\x83\x51\x7b\x0e\x3a\xd4\x53\xd5\x43\x06\x71\x8e\xc4\xa6\x28\xbd\x3b\x0b\xc1\x12\xec\x6f\x6c\x70\xfb\x64\x6f\x7f\xa7\x0d\x96\xfe\xd2\xc7\x51\x55\x7a\x6d\x43\xb3\x1d\x14\x76\x00\xce\x24\x89\x6a\x45\x8f\x67\x37\x13\x22\x49\x54\x87\x18\x60\x9c\xbe\x0d\xde\xbd\xf9\x4b\xd0\x65\x72\x35\x78\x01\x96\x17\x60\xf9\x0f\x00\x96\x6a\x4a\xd6\xca\x8d\x3d\xd9\xb9\x23\xd7\xc5\x53\x3c\x2b\xb7\x2f\xeb\xd5\xfe\x5d\xa8\x34\x14\x17\x92\x09\x95\xc5\xbb\x77\x7f\x49\xee\xcf\xcf\x66\xbf\xbe\x94\x15\x2f\xd9\xff\xad\x65\xff\x93\x77\x1d\x6b\xb1\x4d\xbb\xb8\xfc\xff\xab\x21\xe4\xb7\xbb\x65\x4b\xd5\x7e\x6b\x7e\x3d\xb8\xee\x76\xd2\xea\xe6\x8b\xb8\x63\x57\x97\x83\x7d\xc4\xcc\xab\x03\x87\x60\x64\x10\xc4\xdb\x0a\x3f\x66\xf6\xbc\xc7\x7c\x37\xa1\xd0\x36\xb0\xa8\x17\x39\xd5\xaa\xa5\x41\x79\x58\x11\x90\x36\x3c\x9d\x66\x68\x6b\x82\xa0\x1d\xf0\xf6\x29\xa6\x44\xdd\x0f\x07\xfd\x35\x0e\xef\x18\xdf\x10\xbc\x0f\xe6\xee\x4b\xa5\x52\xbf\x0e\x78\xa5\xd8\xe9\x9b\xe0\x7d\x70\x74\xe4\x5e\xcd\x79\x6f\x0f\x58\x91\x65\x10\x87\xf2\xf5\x97\x39\x97\x67\x34\x87\x42\xaf\xcd\xb0\xbd\x00\x55\x09\xbc\xc9\xfc\x87\x0e\xca\xa5\x29\xcf\x37\x2e\x45\x98\xe5\xb9\x38\x40\x0b\xed\xd0\xc4\x0b\xe9\xbc\x72\xa6\xb6\x67\x80\x73\xdb\x5a\x55\x45\x77\xa4\x62\x5a\x12\x7c\xa9\x24\x4e\x7b\x47\x6f\xc4\x3f\xdd\x8e\xf8\x5d\xe3\xa3\x47\xc5\x3f\x02\x3e\xfc\xf6\xf9\x06\x26\xac\x11\x4b\xa8\x49\x32\x1b\x1b\x0e\xfe\x56\x03\xcf\x0a\xbb\x61\xbb\x2f\x08\x6a\xe1\xab\x6d\xf8\xb1\xb7\xd7\x6d\x10\xc3\xd5\x49\x52\x3e\xab\xcd\x66\x97\xc8\x65\xed\x48\x0c\x97\xb8\x2b\x9c\x98\x52\x5a\xe4\xa1\x84\x02\xbf\xe4\x97\x28\x57\x3c\x9a\x37\xe1\x20\xe1\xb2\xcc\x9f\x44\xf4\x22\x73\xf9\x6e\x29\x16\x0b\xc3\xca\x93\x94\x67\x45\xfd\xae\x12\x02\x33\x5e\x3b\xa0\x77\x0f\xfc\x7f\x2f\xd2\x98\x07\xbf\xdc\xc0\x7f\x01\xf1\xa0\x6f\x67\x79\xb8\x16\x0f\xb2\x6c\x2a\x49\x09\x58\x99\x9c\x6d\x03\x97\x56\x85\xb4\x3a\xcd\xae\xb6\x02\x8e\x75\xf8\x46\x47\x1c\x3e\x8f\x68\xba\x9a\xf1\x28\xa6\xee\xdf\x3e\x94\x44\xd7\xfc\x00\x71\x64\x07\xe8\x5c\xa0\xa9\xaa\x65\x16\x22\x82\x2f\x90\xbc\x70\x82\x9e\x81\xbf\x6a\xe4\xaa\x23\xf7\x0c\x76\x87\x07\xd5\xfb\x8d\x6d\x70\x84\xc5\x0b\x05\xfc\xc8\x01\x33\xb9\xf8\x0b\x8e\x48\x3e\x59\xa8\xe3\x91\x94\x6b\x83\xe0\xe9\x71\xf0\xf6\x27\xf7\xb6\x23\x8c\xb3\x18\x22\x73\x14\x33\x75\x1f\x79\xfd\x90\xe2\x5c\x5a\x31\x37\x03\xae\x44\x15\xc4\x8a\x1c\x9b\xdb\x55\x37\x82\xb7\xf5\xd1\xe2\xb4\x52\x91\xa7\x87\x60\xc2\x7b\x22\xa9\x22\x56\xe8\x78\x29\x57\xe9\x49\xc7\x47\xbb\x62\x21\x6f\x4c\x5f\x4d\xa6\xd7\xbf\x0c\xfb\xf3\xdb\xc1\xf0\xe2\xec\x66\x34\xbf\x9d\xdd\x9c\x8b\x6f\x00\x3f\x2d\x51\xd5\x5b\xd0\x68\x53\x57\xea\x5f\x8f\xe7\xc3\x31\x28\x49\xf6\x96\x6e\x40\xb2\xdd\xc7\x51\x65\xe8\x4f\x74\x80\x35\xc1\x02\x31\x96\xa9\x6f\xd1\x1f\x5c\x67\x5f\x9b\x5d\x80\x03\x14\xae\x13\x61\xb2\x0e\x0d\x8e\x84\xac\x23\x47\xd4\xbe\x56\xf0\xa8\x55\x15\x71\xf3\x2e\x16\x26\x5f\xa9\x46\xf7\x1d\x94\x91\xcd\x9c\xd6\xbd\x3a\x1d\x4e\x46\x7f\x9b\x5f\xcb\xcb\x18\x29\xe1\x37\x97\x42\x0c\x73\x01\x9f\x3d\x7d\x01\x2e\x61\xc9\x12\xab\x68\x5a\x8d\x90\x1a\x64\x3d\x33\xea\x0c\x00\x36\x85\x74\xa5\x67\x1c\x6c\xca\x8b\xa3\xc6\xbe\x99\xfd\xc5\xcc\x84\x63\x2d\x0c\x3d\x01\x4f\xaf\xaf\x4c\xf6\x5e\x7d\x9e\xf5\xa7\x97\x93\xf9\xf7\x2d\xae\x9f\x8d\x9e\x1c\xae\x72\x4a\xef\x37\x6d\xd1\x76\x60\xa8\x87\x5f\xac\x45\xa3\x7a\xb0\x3f\x24\xba\x62\xfa\x00\x2e\x66\xf2\x6f\xdf\xca\x4e\x4e\xa6\xc3\xd9\x70\x3c\xb8\x55\x6d\x8b\x67\x08\xae\xd8\xc1\x1f\x88\xa4\x67\xc7\x50\x19\x3d\xfb\xa2\xa6\xcb\xd0\xbd\x78\x5c\x07\xe5\xb5\xad\xd0\x4b\x8a\xf5\xeb\x8e\xe7\x41\xbb\xa2\xa6\x19\x4f\x77\xe0\xb4\xf3\x2e\xa3\x61\x99\x33\xdc\x2a\xe0\xb7\xca\x3f\xe8\x63\xf5\xc7\x1d\xa8\xfa\x20\x01\x84\x69\x4a\x36\x97\xcb\xda\x3b\x8b\x1a\xd9\x56\xfc\x48\x2c\x65\xd7\xc0\xd6\x92\x0e\xc1\x76\x40\xdc\xb6\x0e\x93\x8c\x6f\x74\xdf\xca\xdb\x5d\x9f\xd1\xb4\xa4\x35\x0f\xac\x79\x25\x03\xbf\xaa\xc7\xf3\x95\xbf\x5b\x40\x45\x20\xff\xe2\x4d\x94\x47\x45\xc2\x36\x0c\x62\x9b\x55\xff\xf2\x2d\x90\x6f\x01\xec\x63\x67\x93\xc9\x32\x76\xcc\xcb\x67\x9d\xb4\x76\xc5\x93\x2a\xb5\x47\xdd\x47\xc1\xbb\xe0\xad\x5f\x8b\xef\x58\xbb\xf4\x4b\xe5\x29\x6c\xe4\x45\x7d\xae\x5a\x44\x39\x4a\x60\x17\x98\x33\xef\xf1\xb2\xfb\x02\xcf\x04\xff\x02\x1a\xff\xbf\x02\x17\x38\x52\x4f\x30\x99\x93\x00\x75\xde\x76\x38\xdb\xdd\x11\x1f\xc5\xdc\x8d\xa9\x79\xcb\xe6\x3c\x34\x52\xc1\x56\x9b\x0f\xb1\xf3\x6c\x94\xd6\x13\xff\xe7\xce\xd4\x9e\x48\xa9\xc4\x88\x7f\x86\x65\x06\x1e\xff\x13\x47\x95\x36\xfc\x28\xa8\x48\xb6\xb5\xa8\xb8\x5e\xff\xe9\x19\x41\xb0\xb7\xb9\x86\x6a\xbf\x66\xf4\x4f\x8f\xae\xaf\x8c\xb0\xaf\x88\xb2\xc6\x48\x93\xc4\xcc\x7d\x56\xe8\xcc\xa2\x7d\xd1\xc6\x7c\x27\x45\xb1\xd8\x4b\x94\x6f\x08\x4b\x92\x2b\xa8\x9d\xfe\x5b\xcc\xd7\x63\xea\x74\x48\x8f\x74\x1b\x7b\x7b\x3a\x3d\x73\x4a\xbd\xd1\x3e\xdf\x5c\x3d\xd8\xc1\x25\xce\x9b\x26\xe7\xc9\xdf\xbf\x03\x00\x00\xff\xff\x78\xd0\xf7\x0b\xd4\x3e\x00\x00")

func templatesJenkinsNormalJobXmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesJenkinsNormalJobXml,
		"templates/jenkins/normal-job.xml",
	)
}

func templatesJenkinsNormalJobXml() (*asset, error) {
	bytes, err := templatesJenkinsNormalJobXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jenkins/normal-job.xml", size: 16084, mode: os.FileMode(420), modTime: time.Unix(1444386461, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesJenkinsPipelineXml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x93\x4f\x6f\x13\x31\x10\xc5\xef\xf9\x14\xd6\x8a\x6b\xbc\x29\x07\xd4\x83\xe3\x82\x02\x95\x40\x05\x8a\x28\x5c\x91\xbb\x3b\x71\x0c\xde\xb1\xe5\x3f\x49\xaa\xaa\xdf\x9d\xd9\xdd\x38\x69\xca\x1f\x29\xea\x29\xf1\x7b\x3f\x3f\xcf\x78\xbc\xe2\x62\xdb\x59\xb6\x86\x10\x8d\xc3\x79\x75\xc6\x67\x15\x03\x6c\x5c\x6b\x50\xcf\xab\x6f\x37\x97\xd3\xf3\xea\x42\x4e\x44\x04\xde\x1a\x75\xeb\x2c\xff\x09\xf8\xcb\x60\xe4\xde\x78\xb0\x06\x81\xbf\xa5\x1f\x0a\xb8\xbb\xde\x09\xdf\x0d\x6c\x98\xb7\x59\x1b\x0a\x6c\x77\xe6\xb4\xe0\xd3\xd1\x79\x3d\xe3\xe7\xfc\x55\x25\x27\x8c\x09\x54\x1d\xc8\xfb\x7b\xc6\x3f\xd1\x1f\xf6\xf0\x20\xea\x41\xe9\xad\xa5\xb1\x09\xc2\xbb\x2d\x34\x39\xb9\x10\xe5\x52\xd9\x08\xa2\x7e\x2a\x1f\xd0\x2f\x19\x32\x1c\x63\xa3\xd4\x23\x3e\x38\x0f\x21\x19\x88\xac\xb1\x2a\xc6\x79\xb5\xca\x6d\x74\xc8\x3b\x47\x75\xf2\xbe\xf0\x17\xd7\x23\x73\x77\x65\x62\xaa\xea\x61\x5b\xe3\x3a\xef\x10\x30\x7d\xf5\xd0\x0c\x87\x91\x78\xe2\x85\xfc\x98\x2e\x1e\xa7\x8c\x21\xff\xed\x7d\xb0\x97\x26\xc4\xf4\xc1\xdd\x0e\xc8\xe5\x6e\x31\x60\x7b\xa7\xa0\xd4\xd0\x9e\xbc\x52\x07\xb0\xe8\x63\xd9\xf5\xf3\xeb\x16\xf5\x9f\x17\x22\xd0\x7d\x5e\x96\x7d\x51\x9e\xcd\xa8\x8f\x23\xa5\x67\xe2\xca\x6d\xde\x68\x1d\x40\xab\x04\x6d\xf1\xca\xb0\xfe\xe1\x96\xf0\x85\xb3\xb9\x43\x8a\x1e\x93\xcb\x72\xc8\x75\x34\x53\xd4\x12\xa9\x26\xca\xd9\xad\xf6\x27\xae\x55\x52\x87\xa7\xf3\x58\xea\x91\xec\x5b\x3a\xef\x3d\xd2\x43\x59\x2b\x2b\x5f\x8a\xfa\x89\x52\x72\x16\x2b\x85\x1a\x8e\x72\x8a\xd4\x23\xca\x5a\xb7\xf9\xa8\x30\x2b\x7b\x13\x8c\xd6\xf4\x45\xc9\x14\x32\x91\x7f\x73\xfa\x1d\xd4\x29\x6c\x7d\x19\x6a\xa4\xb7\x76\xf2\x74\xe4\xe4\x77\x00\x00\x00\xff\xff\x31\x82\x93\x82\xc0\x03\x00\x00")

func templatesJenkinsPipelineXmlBytes() ([]byte, error) {
	return bindataRead(
		_templatesJenkinsPipelineXml,
		"templates/jenkins/pipeline.xml",
	)
}

func templatesJenkinsPipelineXml() (*asset, error) {
	bytes, err := templatesJenkinsPipelineXmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates/jenkins/pipeline.xml", size: 960, mode: os.FileMode(420), modTime: time.Unix(1444285629, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"templates/jenkins/multi-job.xml":  templatesJenkinsMultiJobXml,
	"templates/jenkins/normal-job.xml": templatesJenkinsNormalJobXml,
	"templates/jenkins/pipeline.xml":   templatesJenkinsPipelineXml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"templates": &bintree{nil, map[string]*bintree{
		"jenkins": &bintree{nil, map[string]*bintree{
			"multi-job.xml":  &bintree{templatesJenkinsMultiJobXml, map[string]*bintree{}},
			"normal-job.xml": &bintree{templatesJenkinsNormalJobXml, map[string]*bintree{}},
			"pipeline.xml":   &bintree{templatesJenkinsPipelineXml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
