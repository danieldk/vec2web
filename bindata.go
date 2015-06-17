// Code generated by go-bindata.
// sources:
// static/about.html
// static/index.html
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
	"os"
	"time"
	"io/ioutil"
	"path"
	"path/filepath"
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
	name string
	size int64
	mode os.FileMode
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

var _staticAboutHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x5f\x73\xdb\xb8\x11\x7f\xf7\xa7\x40\x78\x99\xb1\x3c\x17\x92\xb1\x13\x37\x89\x43\xa9\x75\xec\xa4\xc9\xcc\xdd\x5c\x26\x56\x7b\xed\xdc\xdc\x03\x48\xac\x44\xc8\x20\xc1\x00\xa0\x64\x35\xe3\xaf\xd5\xa7\xbe\xdd\x17\xeb\x02\xa4\x24\x92\x12\x15\x5f\x33\x97\x71\x67\xc2\x07\x8b\x04\xf6\xff\x2e\x7e\xbb\x94\x1c\x3d\xb8\xfc\xe9\x62\xfc\xcf\xf7\xaf\x49\x6a\x32\x31\x3a\x88\x1e\xf8\x3e\xd1\x74\x0e\x8c\x4c\x94\xcc\x48\xa9\xc4\x70\xf0\xf8\xf1\xd3\xc7\x47\xa9\x31\xc5\x59\x18\x4e\xc1\xc4\x52\x1a\x6d\x14\x2d\x82\x44\x66\x21\xdc\xd0\xac\x10\xa0\xc3\x9c\xce\x63\xaa\x42\xe2\xfb\x28\xc6\x4a\x23\x82\xe6\xd3\xa1\x07\xb9\x37\x8a\x52\xa0\x6c\x14\x65\x60\x28\xb1\x82\x7c\xf8\x58\xf2\xf9\xd0\xbb\x90\xb9\x81\xdc\xf8\xe3\x65\x01\x1e\x49\xaa\xa7\xa1\x67\xe0\xc6\x84\x56\xc4\x4b\x92\xa4\x54\x69\x30\xc3\xbf\x8d\xdf\xf8\xcf\xbd\xd1\x01\xc1\xab\x92\xb3\xda\xf1\x4a\x33\xe9\x6c\xe5\x34\x83\xa1\x37\xe7\xb0\x28\xa4\x32\x0d\xc1\x0b\xce\x4c\x3a\x64\x30\xe7\x09\xf8\xee\xe1\x11\xe1\x39\x37\x9c\x0a\x5f\x27\x54\xc0\xf0\x18\x05\x55\x92\x0c\x37\x02\x46\x73\x48\x4e\x16\x10\x47\x61\xf5\x58\xef\xd9\x28\xbd\x5a\x85\x01\xc5\x2b\x20\x17\x57\x57\xce\x75\xb7\x2f\x78\x7e\x4d\x52\x05\x93\xa1\x17\x86\x19\xbd\x49\x58\x1e\xac\xc3\x66\x1f\x6c\xe4\xd6\x0b\xe1\x93\xe0\x49\x70\x12\x26\x5a\x6f\xd6\x82\x8c\x23\x95\xd6\x1e\x51\x20\x86\x9e\x36\x4b\x8c\x71\x0a\x60\xbc\xa6\x0d\x17\xa5\x36\x98\xa5\x6a\x97\x4c\xa4\x22\x26\xe5\x9a\x18\xc0\x94\x50\x03\x3b\x0d\xba\x6b\x22\xab\x8f\xfd\x36\x74\xc4\x6a\x94\x8b\xde\xcd\x74\x90\x08\x59\xb2\x89\xa0\x0a\x9c\x6c\x3a\xa3\x37\xa1\xe0\xb1\x0e\x67\x1f\x4b\x50\xcb\x4d\x30\xe6\x54\x70\x46\x8d\x54\xe1\xe3\xe0\x34\x78\xd2\x0e\xc2\xdf\x57\x9b\x77\x8b\xc6\xdb\xf1\x8f\x3f\x9c\x12\x9d\xf2\x8c\xd0\x9c\x91\x0f\xa0\x0b\x99\xb3\x60\x56\x45\xe6\xdd\xeb\xe7\x44\x97\x85\xad\x08\x22\x27\x35\x31\x08\xc8\xb0\x32\xb4\x63\xc8\x80\x71\x4a\xac\x81\x1c\xa3\xb9\x8e\x1d\x8a\xfe\x85\x4f\x88\x30\x28\x82\xbc\xf8\xb5\x5a\xc5\x75\x9d\x28\x5e\x18\xa2\x55\xb2\xf1\x5e\x6a\x1d\xd4\xf9\xb6\x7e\xdb\x1a\x3e\x45\x83\xe6\x98\xe2\x67\x98\xe2\xf5\xb3\x73\x68\xa6\xf1\x64\x84\x95\x98\xdf\x23\x55\x55\x8e\x85\xc7\xc1\x53\x94\x59\x3f\xf5\x48\x8c\x1e\xfc\x02\x39\xe3\x93\x5f\x2b\x77\xa2\xd0\x9d\x44\x1b\xb4\x28\x96\x6c\xb9\x0a\x1f\xe3\x73\x92\x08\xaa\xf5\xd0\xb3\x87\x85\xf2\x1c\xd4\x2a\xb6\x75\x74\xaf\x0c\x35\x3c\x21\x55\x59\xac\xa3\x83\x9b\xb8\xb2\xe2\xad\x37\xab\x0f\x9f\xc1\x84\x96\xc2\x78\x2b\xca\x1e\x3d\xfe\x44\x94\x9c\x35\xa8\xda\x74\xb5\x30\x6b\xb7\xb3\x89\x34\xae\x28\x2e\x8d\x91\x39\x31\x88\x1f\x43\xaf\x7a\xf0\x3a\x8c\x46\x4e\xa7\x02\xf0\x90\x0a\x41\x0b\x0d\xcc\x23\x58\x52\xb4\x5e\xb6\x66\x54\xeb\xab\x65\xaa\xa6\x16\x54\xbe\xab\xb8\x3d\x42\x15\xa7\x3e\xdc\x14\x58\x20\xc0\x86\xde\x84\x0a\x4b\xeb\x56\xad\x07\x4a\x8a\xb5\xaa\x8e\x71\x36\x99\xc8\xb6\x32\x47\x2b\x5f\xe6\x62\xe9\x8d\xc6\x95\x41\xc8\xc3\xa7\x18\x52\x99\x63\xc6\x90\x6e\x2f\x33\x47\x5d\xbe\x53\xf1\x75\x89\xa3\xb0\x0a\x6a\x67\x95\x76\x62\x1c\x2b\x0c\x8f\x57\x23\x41\x10\x7a\x1b\xd8\xa4\xad\xb4\x86\x98\xd7\xad\x3c\x73\xb6\x0e\x60\x47\xec\x2a\x37\xeb\xe4\x75\xd3\x5f\x8a\x06\xc7\xaa\xec\xf0\x63\x3b\x13\x82\x8f\xd0\xea\x8d\x81\x6f\x65\x06\xd6\xba\x08\x61\x69\x07\xf1\x4a\x2c\x4d\x0c\x9f\xa3\xda\x06\x2f\x8d\x65\x89\x35\x7d\x6e\x3f\x76\x4b\x88\xc2\x52\x6c\xbb\x6d\x0f\x51\x18\xa0\x71\x1b\xb7\x36\x87\xa8\x4d\xd4\x39\x1a\xcd\xc3\x66\x61\x79\x73\x2e\x1b\xc7\x44\xc9\x45\xef\x41\x13\x7e\xc6\xfc\x13\xaf\xab\x6c\x1f\xf5\xf3\xf6\x79\x4c\x4f\x56\x0e\xe3\xdd\x41\x73\xa7\x68\xfb\x3e\xb6\xfd\x07\x33\x4f\x68\x51\x08\x9e\xb8\xfa\x26\xa5\x46\x3c\xc5\x92\x40\x24\xd7\x16\x7b\xff\x0a\x2a\xc3\x0a\x5c\x48\xc5\xb0\x59\x49\x32\xe1\x39\x6b\x49\xd1\x3c\xe3\x19\x15\x35\x05\x82\x37\xcd\xf5\x02\xd4\x1a\x9b\x51\x86\x49\xc1\xc2\x7a\x46\x22\x3e\x3a\x27\xdc\xc9\x79\xd5\x49\x24\xd5\xe4\xc2\xfe\xc1\xad\x3f\x47\x21\x1f\x05\xe4\x0d\xca\xaa\x1b\xdd\x59\x2b\x43\xc5\x5e\xaf\x50\xc7\x1b\x2c\xf1\x6b\x05\x3c\x49\xad\xa4\x5a\x21\xae\xbf\x47\x28\xd0\x6e\x09\x15\xe1\xf3\x25\x94\x46\x27\x29\xce\x3e\xac\x26\x6c\x49\xaa\x98\x9c\x35\x51\xac\x48\xb8\xa5\xe7\x3c\xd3\x06\x14\xa3\x59\x5b\xcd\x39\x76\xd0\x86\x96\x0f\xd2\xec\xa2\xda\x23\xf7\x92\xe6\x1c\x04\x6d\xd3\x57\x8b\x4d\xeb\x7f\xfb\xb7\x32\xb0\x43\xe8\xef\x08\xd6\x18\x33\x63\x13\xb7\xce\xb8\x82\x02\x9b\x14\xb6\x59\x57\x0d\xb6\x3c\x14\xb4\x38\x12\x05\x38\xad\x30\x12\x2f\x5d\xd9\x2c\x79\x3e\xed\xe2\x4d\x67\xc8\x90\x0c\x82\xa9\x94\x88\xa3\xae\x23\x16\xa1\x55\x78\x82\x0a\xf1\x68\xaf\x6e\x3b\xe0\x43\x88\x6d\x14\x68\x5b\x55\x7d\xfd\x0a\x50\xfe\x62\xb1\x08\xf4\x44\x07\x65\xce\x7d\x53\x42\x8c\x06\x41\x1e\x30\x08\x21\x0f\xa9\x4e\x84\x6d\xba\xb2\x54\x09\x4e\x4b\x38\xfc\xe1\x48\x41\x43\x4b\x46\x7d\xb6\x08\x6c\x8f\x47\x90\xff\xed\x3f\xaf\xa8\x7f\x19\xfe\x6c\xad\xb0\x13\x62\x51\x6a\x2c\x3f\x00\x81\x83\x35\x80\x0d\xad\x3b\xe5\x89\xd9\x28\xcf\x28\x17\x46\x9e\x31\x97\x14\xd4\xe6\x5f\xcb\xeb\xbf\x74\x4d\xf0\x46\xfb\xf7\x9d\x3e\x3b\xef\x64\x76\x2a\xe5\xb9\x3d\x23\x2e\xee\x41\x37\x85\x9b\x87\xcf\x81\x41\x0f\x74\xd4\xb7\x8d\xfb\x6a\x52\x08\xd7\xf8\xe5\x90\xab\x26\x68\x0e\x36\xa1\x9b\x08\xeb\x0c\xd2\x82\xeb\x9d\x53\x22\x0e\x38\xc7\xc7\x38\xe1\xd4\x33\x63\xcf\x80\xd3\x96\x7b\xd7\x61\x7b\xd6\x9d\xb5\x3f\x2b\xf9\x0b\x07\xdb\x59\xef\x5c\xdb\xd6\xdc\x52\x5d\xcd\x35\xee\x75\x68\x46\xe7\xb4\x5a\x6d\x24\xe3\xe1\x80\xc9\xa4\xb4\x13\xec\x51\x80\x67\x88\x2d\x07\x93\x32\x4f\x1c\xe6\x0e\x8e\xc8\xa7\x56\x8d\x3f\x1c\x1c\x7e\x47\x73\x2a\xe4\x74\x89\x28\x98\x1d\x1e\x05\xdb\xf6\x0c\x3e\x75\x10\x94\xe0\x50\xac\x35\x9d\xc2\x19\x39\x74\xd0\x8e\x6e\x95\x60\xb1\x21\x97\x86\x38\x1f\x0f\x1f\x6d\xf1\x4c\x00\x58\x4c\x93\xeb\x77\x58\x08\xfa\x8c\x6c\x0b\xb5\x97\x63\x46\xb1\x53\xb1\x2c\x52\x3b\x85\x90\xf5\x9d\x2f\xaf\x77\x48\xb5\x17\xcf\xf7\xb1\x29\xc8\xe4\x1c\x7a\x58\xeb\x84\xe0\x39\xe9\xe5\xc6\x93\xa9\xd3\xc3\x2d\xee\xdb\x1d\x1e\xe2\x09\x64\xbd\xae\x59\x04\x3a\xee\xdb\xb4\x57\x33\xa8\xd8\xc6\xb8\xd2\xa6\x82\xcc\xcf\x04\xb6\xe3\x0b\x82\xeb\x3e\x2d\xf6\x42\x69\xaf\xb3\xc2\x2c\x3f\x47\xb7\xd7\x2a\x65\x5f\xd7\x15\x02\xb4\x7d\x4b\x4a\x68\x6e\x4d\x8c\x81\x80\x15\xbc\x1d\xae\xe6\x75\xdb\xbb\xbb\x7b\x67\x47\xa8\xed\xe5\x20\xfd\xee\x01\xd5\x80\xf9\x64\xf7\x2e\xa2\x1d\xb3\xee\x45\x54\x9f\xdc\x3d\xaa\x26\xe5\xea\xfe\x05\xb5\x6d\xd5\x57\x8c\xe9\x36\x4c\xb4\x56\x6e\x8f\x02\x99\x0f\x0e\x75\x99\xe0\x94\xa0\x03\xdb\x86\x83\x78\x7e\xf8\x88\xac\x10\x7a\x00\x5d\x84\xb6\x17\x04\x38\x2b\xcd\x11\xd0\x2f\xab\x97\xe7\xc1\xd1\xcb\x83\x2d\xa2\x87\x03\xef\x3b\x5d\xc6\x19\x37\xe7\x15\xa0\x7b\x47\xc8\x26\x8b\x81\xc7\xb8\xa6\xb1\xc0\x97\xdc\x47\xc4\xa8\x12\x90\x79\x17\x2f\xc2\x1c\x8a\x1e\x5b\x4a\xe4\xb4\x03\xcb\xc0\xf3\x76\xd2\x1e\xd6\xb4\xaf\x95\x92\x0a\x5b\x46\xca\x19\x0c\xf6\x52\xbe\x57\x72\x8a\x77\x1a\x89\x75\x2a\x17\x6b\xe2\x6d\x96\xc0\x36\xcf\x1d\x1d\xc7\x5e\xb6\xf1\x9d\x11\xef\xfd\x4f\x57\x63\x6f\x77\x7d\xd9\xb7\xf5\x33\xeb\x8d\x8d\x6c\xb3\xb1\xa1\x47\x1a\xdf\x12\xb0\xf4\xfe\x85\xa6\xf6\x33\x8f\x2b\x15\x33\x2d\xf3\x1e\x15\x09\x4d\x52\x24\x71\xaf\xfc\xbb\x29\x4a\x25\x50\x44\xad\xbc\x47\x4a\x5d\x01\x67\x64\xd3\x9a\xad\xfa\x5d\xc9\x5f\x5d\x3b\xa3\xd9\x17\xfa\xb5\x4f\x8a\x2e\x5c\x4a\x2b\xf1\xfd\x84\x77\xaa\x1e\xe7\x74\x8f\x90\x1e\x38\x01\x5b\x23\x4d\x37\x67\x1f\xff\xf1\xf6\xc3\x3e\x3f\xe7\x54\x59\x2e\x32\x24\x8e\x34\xa8\xbe\xca\xd2\x30\xc6\x81\x67\x9f\xfd\xff\x43\x74\x36\x65\xef\x4a\xd9\xeb\x54\xe7\x5e\x35\x8e\xe3\x47\x3d\x45\x35\x76\x14\x1b\xa0\xc9\x7f\x68\x78\xb7\x91\xa5\x43\x79\xdb\x05\x05\x6b\x2b\x4a\x37\x34\x4f\xe0\xdb\x70\xf7\xc7\x0f\x77\x77\x6f\x9a\xf7\xa9\x5d\x9e\x7f\xf5\x69\xee\x5e\xf4\xc8\x2b\x9e\x71\x7c\x61\xe3\xe6\x5b\x9b\x6c\x62\xc4\xd7\xef\x93\x2b\xed\xff\x7f\x8d\x72\x7f\x09\x7d\xeb\x95\x5f\xde\x2b\xbf\x20\xc2\x77\x6a\x97\x9b\x83\xd5\x6a\x9e\xeb\x04\x74\x8a\xa2\x95\x88\x87\x01\x60\x85\xbb\x8d\x06\x34\xf1\x9c\xc1\xcd\x23\xc2\x0d\x64\xdd\xbc\xd9\x5c\x29\xb9\xc0\x5c\xa1\x83\x91\xb1\x5f\xc7\x76\x71\x63\x1b\x5f\x68\x51\x40\xce\x06\xc8\xd7\x21\xc5\x95\xd5\xa6\x13\xc7\x46\x1e\xf9\xde\xe9\x0d\x7e\xb6\x78\xfe\x3d\xf1\xa2\xd0\xae\x1e\xdd\x9d\x71\x13\xee\x3e\xf6\xdb\xa3\x83\x76\x84\xbb\xbf\x54\x1e\x1c\x44\xee\x77\xd8\xd5\x77\x85\x65\x2c\x64\x72\xed\x17\x52\x1b\x21\x29\xf3\x8f\x29\x3c\xa3\xa7\x93\xe3\x27\x8f\x27\xc9\xb3\x17\xf1\xd3\x09\x8b\x9f\xd3\xa7\x27\xcf\x4e\xd8\x8b\xa7\x27\x7f\x8a\x4f\xbd\xd1\xe6\x77\x0e\x32\x22\x81\xfd\xaa\xf2\x95\xbc\xf9\x01\x41\xc2\x3e\xea\x94\x2a\x78\x57\xad\x1d\x7c\xc2\x5a\x28\x04\x5d\x9e\xe5\x32\x07\xf2\x80\x67\xf6\x57\x5c\x9a\x9b\x97\xb7\x68\x95\x35\x62\x14\x85\xee\xe7\xcc\x28\xac\xfe\x8f\xe1\xbf\x01\x00\x00\xff\xff\x18\x4c\xc3\x34\xd8\x20\x00\x00")

func staticAboutHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticAboutHtml,
		"static/about.html",
	)
}

func staticAboutHtml() (*asset, error) {
	bytes, err := staticAboutHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/about.html", size: 8408, mode: os.FileMode(420), modTime: time.Unix(1426238000, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _staticIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1a\x6b\x6f\xdb\x38\xf2\xfb\xfe\x0a\x96\xed\x21\x0e\xae\x92\x9a\x64\x8b\xeb\x65\x6d\x03\xdd\x6e\x7a\x5d\xe0\x16\x2d\x36\xb9\x17\x8a\x7e\xa0\x25\xda\x66\x4a\x89\x2a\x49\x39\xf1\x2d\xf2\xdf\x6f\x46\xd4\xfb\xe1\x38\x4d\xb7\x97\xc3\x95\x40\x6c\x8b\x9c\xf7\x0c\x67\x86\x54\xa6\x8f\x7e\x7a\xfb\xea\xe2\x5f\xef\xce\xc8\xda\xc6\x72\xfe\xdd\x14\xbf\x88\x64\xc9\x6a\x46\x79\x42\xe7\xd3\x35\x67\xd1\x7c\x1a\x73\xcb\x00\xc2\xa6\x1e\xff\x94\x89\xcd\x8c\xbe\x52\x89\xe5\x89\xf5\x2e\xb6\x29\xa7\x24\x74\x4f\x33\x6a\xf9\xb5\x0d\x90\xc4\x0f\x24\x5c\x33\x6d\xb8\x9d\xfd\xed\xe2\xb5\xf7\x82\xce\xbf\x23\x30\x1c\x9d\x72\x85\x66\x76\xd9\x59\x4a\x58\xcc\x67\x74\x23\xf8\x55\xaa\xb4\x6d\x10\xbe\x12\x91\x5d\xcf\x22\xbe\x11\x21\xf7\xf2\x87\xa7\x44\x24\xc2\x0a\x26\x3d\x13\x32\xc9\x67\x47\x40\xc8\x51\xb2\xc2\x4a\x3e\xdf\xf0\xf0\xf8\x8a\x2f\xa6\x81\x7b\x2c\xd6\x1e\x79\x1e\xf9\x51\x29\x6b\xac\x66\x29\x90\xd7\x9c\xbc\x3a\x3f\x27\x9e\x57\x48\x21\x45\xf2\x91\xac\x35\x5f\xce\x68\x10\xc4\xec\x3a\x8c\x12\x7f\x51\xc2\xe3\x43\xa8\xe2\xa0\x9a\x08\x4e\xfc\x13\xff\x38\x08\x8d\xa9\xe7\xfc\x58\x00\x94\x31\x94\x68\x2e\x67\xd4\xd8\xad\xe4\x66\xcd\xb9\xa5\x4d\x19\x5e\x65\xc6\xaa\x98\xb8\x55\xb2\x54\x9a\xd8\xb5\x30\xc4\xf2\x38\x95\xcc\xf2\x41\x81\xd0\xfe\xa7\x41\xb0\xe2\xb6\x66\x86\xe2\xf0\x6b\x06\x58\xdc\x04\x09\xdb\x2c\x98\x2e\xbe\x76\xcb\xd0\x21\x6b\x80\x2e\x68\x77\x69\xfc\x50\xaa\x2c\x5a\x4a\xa6\x79\x4e\x9b\x5d\xb2\xeb\x40\x8a\x85\x09\x2e\x3f\x65\x5c\x6f\x6b\x63\x6c\x98\x14\x11\xb3\x4a\x07\xcf\xfc\xe7\xfe\x49\xdb\x08\x7f\x2f\x17\xf7\xb3\xc6\x9b\x8b\x5f\xfe\xfa\x9c\x98\xb5\x88\x09\x4b\x22\xf2\x2b\x37\xa9\x4a\x22\xff\xd2\x59\xe6\xe7\xb3\x17\xc4\x64\x29\x46\x04\x51\xcb\x02\x98\x4b\x1e\x43\x64\x98\x1c\x21\xe6\x91\x60\x04\x05\x14\x60\xcd\xca\x76\x40\xfa\xbd\x58\x12\x69\x81\x04\xf9\xf3\x07\x37\x0b\xf3\x26\xd4\x22\xb5\xc4\xe8\xb0\xd6\x5e\x19\xe3\x17\xfe\x46\xbd\x31\x86\x9f\x83\x40\x1b\x70\xf1\x9f\xc0\xc5\xd5\x73\xae\xd0\xa5\x81\x9d\x11\x38\x32\x77\xa1\xaa\x9d\x62\xc1\x91\xff\x3d\xd0\x2c\x9e\x46\x28\x4e\x1f\xbd\xe7\x49\x24\x96\x1f\x9c\x3a\xd3\x20\xdf\x89\x68\xb4\xe9\x42\x45\xdb\xd2\x7c\x91\xd8\x90\x50\x32\x63\x66\x14\x37\x0b\x13\x09\xd7\xa5\x6d\x0b\xeb\x9e\x5b\x66\x45\x48\x5c\x58\x54\xd6\x81\x45\x98\x29\x71\x8b\x45\xf7\xe5\x45\x7c\xc9\x32\x69\x69\x09\x39\xc2\xc7\x5b\xca\x4c\x44\x0d\xa8\x36\x5c\x41\x0c\xe5\xce\x65\x22\x8d\x31\x5d\x64\xd6\xaa\x84\x58\xc8\x1f\x33\xea\x1e\x68\x07\xd1\xaa\xd5\x4a\x72\xd8\xa4\x52\xb2\xd4\xf0\x88\x12\x08\x29\x56\x4c\xa3\x18\x6e\xbe\x9c\x66\x7a\x85\x49\xe5\xb1\xc3\xa6\x84\x69\xc1\x3c\x7e\x9d\x42\x80\xf0\x68\x46\x97\x4c\x22\x6c\x3e\x8b\x1a\x68\x25\x2b\x56\x1d\xe1\xd0\x99\x80\x56\x8a\x63\xb4\xa7\x12\xb9\xa5\xf3\x0b\x27\x10\xe0\x88\x15\x98\x54\x25\xe0\x31\x80\xdb\x89\x2c\x80\x97\x97\xb3\xf8\xba\xc0\xd3\xc0\x19\xb5\x33\xcb\x3a\x36\x5e\x68\x30\x0f\x2d\x32\x81\x1f\xd0\x3a\x6d\xb2\x96\x5b\x03\xf0\x6b\xcf\xcf\x22\xaa\x0c\xd8\x21\x5b\xfa\xa6\x72\x5e\xd7\xfd\x99\x6c\x60\x94\x61\x07\x5f\x7d\x4f\x48\x51\x42\xb2\xd0\x8a\x0d\x50\x02\x25\x6a\x79\xdf\xa8\x98\xa3\xb0\x53\xc8\x52\x03\xb8\x4d\x60\xb6\x50\x19\xc4\xf4\x4b\xfc\x1a\x46\x99\x06\x99\xec\xab\x8d\x9b\x28\xf0\x41\xb8\x5a\xad\x7a\x13\xb5\x81\x3a\x5b\xa3\xb9\xd9\x30\x2d\xd7\xfb\xb2\x34\xdf\x4a\xaa\x05\x93\x67\x5a\x2b\xb0\x61\x9e\x1d\x67\x34\x12\x06\x8a\xc0\xf6\x94\x24\x2a\xe1\x95\x65\xa1\xc6\x41\xfe\xcb\x3f\xbd\x08\xea\x33\x6c\x29\x02\x31\xcc\x8b\x95\xe6\x5e\x4d\x5b\x3a\x2c\xe6\x39\xf9\x53\x88\x88\x79\x11\x41\x1d\xce\xbf\x98\x15\x0d\x9a\x0a\xa5\xb5\xd4\xb9\xdf\x9b\x52\x17\xe2\x68\x75\x35\x9a\x1e\xa4\x17\x47\xde\xf7\xed\xbc\xd0\x73\x38\x6c\xd8\x85\x29\x55\x80\xdf\x52\x18\xdb\x8d\x12\xf0\xbd\x5b\x4f\x21\x57\x42\xb6\xcf\x37\x1d\x1d\x8d\x87\xc7\x60\x39\xcb\x92\xb0\xbf\xcd\xeb\x85\x8a\x5f\x27\x9b\xe0\xcc\xfc\x5c\xc4\x02\x4a\x9f\xb0\xdb\x91\xf8\x18\x16\xa8\x21\x01\x4b\x98\x54\xab\x6d\x4f\x80\x6a\x7e\x27\xff\x97\x0e\x6a\x88\xb9\x0b\xcd\x91\x4c\x0b\xc8\x5e\xd1\x2a\x75\x4d\x88\x50\x15\x4f\xf0\x3d\x97\xb4\x89\x85\x33\xa4\x30\x64\x1e\x16\x95\xa5\x7a\x9b\x09\x4a\x71\x5c\x90\xc2\x9f\x6d\xf0\xd7\x38\xd3\x45\x69\x0b\x89\x48\xde\x4a\xab\x2c\x6d\xa3\x9e\xc5\x0b\x1e\x45\x22\x59\x99\x01\x02\x68\x73\xb6\xe0\x12\x1b\x01\xe8\x48\x1b\xa0\x35\x1a\x98\x0a\x41\x06\x91\x0d\xf4\x09\xa1\x6d\x89\x50\x38\x85\x16\xcd\x26\xbf\x85\x7d\xe0\x48\x0c\xe8\xd6\xcd\x89\x3b\x55\xbe\x4d\xb5\x2b\xa5\xa1\x92\xfe\x03\x3e\x77\xa9\x23\x92\x34\xb3\x45\xd1\xc4\x5e\x9b\xee\xd2\x2c\x27\x49\x20\x95\x84\x7c\xad\x24\xd4\xe0\x19\x7d\x07\x61\x39\xa4\xe7\x98\x2e\x45\x95\x46\x77\x99\x6c\x11\x0b\x5b\xef\x10\xda\x12\xc3\xfd\x76\x30\x95\x50\x0b\x9b\x10\xf8\xab\xdb\x89\xf3\x7c\x7d\xb8\x34\xa1\x10\xa8\x45\x37\x23\xf7\x04\xdb\x27\xa4\x5d\x84\x95\x9b\x6e\xbf\x50\x2e\xa0\x3f\x23\x92\x0b\xcc\x6f\x81\x5c\x06\xf2\x91\x8b\x64\x72\xf4\x45\x63\xf9\xe8\x9e\xc1\xfc\xf9\xfa\x1c\x17\xfa\x1c\x7f\x51\x7d\x8e\x3b\xfa\xbc\xd6\x83\x89\xf7\xf7\x50\xe8\xa4\x50\xe8\xe4\x8b\x2a\x74\xd2\x51\xe8\x47\xae\xe1\x9c\x79\x9f\x74\xf3\xb2\x2c\x9b\xff\x95\x5c\xd3\xee\x7e\x7a\x00\x03\x4d\xcf\x8b\x76\xd3\xb3\x3e\x9e\xc3\x59\x16\xc4\x81\xd3\xdb\x71\xbf\x7a\xa3\x9e\x3a\x5f\x7f\x9b\xf5\x0a\x37\xb6\x44\xbc\x01\x72\x81\xcf\xcd\x44\x07\xab\xf9\xa7\x07\x27\x6e\x91\xf2\xa8\x4b\x20\xc8\x57\x77\xe9\xd3\x56\x21\xd5\x6a\x05\xac\xcc\x58\x1f\x5a\x8b\xf2\xae\x84\x1c\x48\xcc\x1d\x62\x78\x40\x21\xcd\x87\x52\xda\xaa\xe9\x28\xfb\x29\x07\x52\x9f\xdc\x36\x4c\x66\x3c\x51\x57\x33\x7a\xf4\xec\x59\x73\x0e\x8e\xcb\x33\xda\x9e\x61\xd7\x05\x54\x21\x79\x7e\x41\x74\x4a\x60\xea\x0f\x7b\x1e\xed\x40\xa5\x10\xd8\x43\xf6\x24\x9a\x7f\xca\xb8\xb1\xc3\x87\xaa\x5b\x22\xa4\xe7\xd8\x2f\xdb\xd9\xe7\x0c\xd2\x9e\x46\x45\x87\x8f\x76\x2e\x95\xc0\xed\x68\xba\x2d\x7f\x43\xa4\x4e\xcb\xef\x74\x49\xf7\x0c\xfe\xe6\x92\xfb\xed\x2e\x1a\x82\xea\xf8\x93\x1f\x7c\x0a\x80\xe6\xbd\x48\x90\x5f\x28\xf9\x2b\xa5\xa0\xe9\x65\xa9\x30\x83\x97\x4c\xc1\x91\x7f\x74\xe4\x1f\x97\x57\x4e\x23\xf7\x23\x6d\xba\xfb\xde\xd5\x5d\x76\xaf\xea\x6e\xa5\x7c\xcf\x7b\xb1\xcb\xd1\x6b\xb1\x36\xe7\x16\xeb\x3a\xdd\x05\x97\x6c\xc3\xdc\x6c\x23\x10\x9e\x4c\x22\x15\x66\x78\x01\x76\xe8\x6b\xce\xa2\xed\x64\x99\x25\x21\x1e\x46\xc8\xe4\x90\xfc\xd6\xf2\xeb\x13\x1f\x85\x9d\xfc\xd6\x4b\xb7\xc8\xe4\x94\xd0\xbf\x9c\x5d\xd0\xa7\xbd\x45\x3c\x9c\x5c\x38\x80\x4b\x03\x67\x9c\x3e\x44\xc8\xc2\x35\x2c\xe7\xf7\x2a\xfd\xd5\x4c\x4b\x40\x6d\xf4\x24\x7d\x10\x93\x85\x18\xae\x40\xa2\x92\x1d\xb9\x76\xe5\x2f\x07\x8b\xa2\xba\x55\x72\x90\x3f\xf4\x00\x6f\xfa\x6c\x78\x7e\xfa\x6d\x30\xb9\xfc\xf4\xcf\x37\xbf\x8e\x71\xd9\x40\xbe\x02\x0c\x32\x23\x39\x98\xef\x2e\xe9\x0c\xbf\x00\x5f\xf4\xd9\xe1\x78\x32\xa1\x8f\x9b\xc7\xf8\x43\xdf\xac\xd5\xd5\x64\x40\x38\x07\x7d\xf0\xb8\x7d\xf4\x3e\x38\xf4\xd1\xd1\x13\xfa\x8a\x25\x89\xb2\x90\x7f\x20\x43\xf2\x0d\x1c\xcb\x36\x0c\x7a\x6d\xcc\xf3\xb5\x19\xc1\xa6\xe4\x8f\x28\xe0\x90\xee\xad\x99\x1b\x84\xf8\xae\xcb\xba\xd1\xe4\x02\xdf\x7e\x60\x0e\x84\x49\x0c\x3e\x62\x2b\xf0\xf4\xc1\x05\x5e\x51\xe7\x19\x97\xc0\x0f\x94\x35\x0f\xf6\x83\xbe\xcd\x97\x9c\x47\x0b\x16\x7e\xfc\x19\x32\x02\xc8\x3c\x66\x6a\x40\x06\xb2\x2b\xb9\x4d\xd7\x78\x9b\x45\xaa\x5f\x9e\xfa\x38\x40\x15\x87\x48\x76\xa1\x69\x1e\xab\x0d\x1f\x41\x2d\x76\x26\x98\x71\x14\x7b\x09\xee\x5e\x1f\xec\x13\x55\x4b\xc1\x65\x34\xaa\x5a\xde\xae\x8e\x2d\xe2\x68\x1a\x95\x03\x31\x6d\x6c\x8e\x74\x9b\x61\x3b\xba\x28\x3d\x2a\x42\x39\x80\xda\x59\x9c\xda\xed\x6d\x70\x3b\xa5\xc2\x9a\x28\x34\x56\xed\x24\x82\x8d\x9f\xc7\xe9\x02\xe3\x12\x08\xf7\xcd\xd5\x1c\x37\xa3\xab\xc3\x2b\x03\xa6\xc6\x91\xf7\xcb\xfb\x1b\xd4\x70\xf0\x67\xf4\xe0\x2c\xda\x11\xeb\x41\x58\xf5\x64\x7f\xab\xda\xb5\xd0\x0f\xcf\xa8\x6d\xa9\xbe\xa2\x4d\x6f\x4d\xc0\xbe\x4a\x26\x07\x45\x95\xf3\xf1\xbc\xe1\x2f\x36\x07\x4f\xab\x4a\x34\xe1\x43\x45\x88\xfb\xa9\x86\xec\x9f\xd8\x9f\xdc\x49\x06\x0b\x49\x0f\x08\x6b\x4e\xfb\x84\x74\x08\x68\x2a\x9d\x60\x8b\x89\x25\x23\xa2\x4f\x89\xd5\x19\x1f\x28\x13\x88\xdb\x3c\x52\x1c\xfa\xf8\x72\x6b\x42\xe9\x20\xec\xc1\xe3\x46\xd7\x08\x25\x63\x2d\x22\x3e\x54\xdb\x6a\xc8\xf2\x74\x70\xd0\x29\x84\x7d\x94\xb1\xc6\x04\x47\xd1\x9c\xbc\x7b\x7b\x3e\xd4\x9d\xe0\xc0\x0e\xe0\x14\xb5\x41\xcb\x36\x0b\x1b\x56\x60\x0e\x87\x03\x29\xfe\x0d\xa2\x8e\x23\xef\x6e\x6f\x70\xec\x6e\x71\x70\xb8\x36\xa7\xbc\x68\x1a\x86\xb9\x63\x9f\x83\x63\xd0\x9a\x63\xa6\xaf\x74\xd2\xec\x2a\x77\xe9\x58\x73\x54\x53\xdf\x23\x7a\x72\xa5\x47\x88\x8c\xa4\x93\xbb\x75\x5a\x38\xee\xde\x6d\x7d\xb6\x75\xea\xb0\xdf\xa7\x4d\x6b\xb3\xe9\xb6\x6a\xc3\xfd\x57\x93\xd3\xfd\xcc\xdb\xcf\x2c\x1d\xc8\x9b\x6e\x52\x40\x59\x9b\x77\xf1\xdf\x9a\xbb\xdf\xb7\xb9\xdb\xbf\x68\x3e\xa4\x72\xf9\xf2\xab\x77\x73\x0f\xa2\x46\x36\x5e\x5a\xfc\xbf\x97\xc9\xd6\xfb\xba\xaf\x5e\x27\xab\x97\x8b\xff\x73\x85\x72\x77\x08\x7d\xab\x95\xf7\xaf\x95\xf7\xb0\xf0\x5e\xe5\xb2\xde\x58\xad\xe2\x59\x39\x60\xe0\x7a\xa9\xe5\x0c\x34\x7f\xff\x95\x35\x78\x03\x55\x18\x58\x98\x13\xf7\xbe\xed\x3d\xbe\x1f\xf9\x40\x0f\x3b\xa4\x7a\xef\x0c\x0b\x4a\xfd\xf9\x1d\x84\xfa\x6c\x07\x93\x54\x8f\xe6\x20\xd4\x13\x9f\xc3\x36\xce\x35\x6f\xe4\x5f\x91\x44\xfc\xfa\x29\x11\x96\xc7\xfd\xe0\x1c\xe0\xcf\xd2\x94\x27\xd1\x84\x4e\x55\x8a\xf8\x73\xbc\xae\x42\x64\xf8\xa2\xd3\xa0\x9c\x3c\xec\x10\xea\x4b\x78\x77\x3a\x37\xf5\xc3\xcd\x80\x7f\x3b\x9b\xbe\xa5\xcb\x9d\x55\x47\x0f\x6a\x75\xe5\x7c\x36\xb5\x9a\x04\xf3\x6e\x5d\xe8\xd7\x8f\x42\x25\xc0\xeb\x80\xc2\x4c\xb9\x98\x93\x8b\x2a\x6d\xfd\xfc\x2d\x5d\xae\x32\xce\x1e\xee\x8f\x58\x6f\xa7\x31\xf4\x96\xbd\xf0\xb3\x79\x0b\x3e\x0d\xdc\x3f\x19\x4e\x03\xf7\x7f\xc1\xff\x09\x00\x00\xff\xff\xf3\x04\xfb\xe5\x28\x2c\x00\x00")

func staticIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_staticIndexHtml,
		"static/index.html",
	)
}

func staticIndexHtml() (*asset, error) {
	bytes, err := staticIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "static/index.html", size: 11304, mode: os.FileMode(420), modTime: time.Unix(1434533514, 0)}
	a := &asset{bytes: bytes, info:  info}
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
	if (err != nil) {
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
	"static/about.html": staticAboutHtml,
	"static/index.html": staticIndexHtml,
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
	Func func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"static": &bintree{nil, map[string]*bintree{
		"about.html": &bintree{staticAboutHtml, map[string]*bintree{
		}},
		"index.html": &bintree{staticIndexHtml, map[string]*bintree{
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
        err = os.MkdirAll(_filePath(dir, path.Dir(name)), os.FileMode(0755))
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
                err = RestoreAssets(dir, path.Join(name, child))
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

