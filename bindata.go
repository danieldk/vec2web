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

func bindata_read(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindata_file_info struct {
	name string
	size int64
	mode os.FileMode
	modTime time.Time
}

func (fi bindata_file_info) Name() string {
	return fi.name
}
func (fi bindata_file_info) Size() int64 {
	return fi.size
}
func (fi bindata_file_info) Mode() os.FileMode {
	return fi.mode
}
func (fi bindata_file_info) ModTime() time.Time {
	return fi.modTime
}
func (fi bindata_file_info) IsDir() bool {
	return false
}
func (fi bindata_file_info) Sys() interface{} {
	return nil
}

var _static_about_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x5f\x73\xdb\xb8\x11\x7f\xf7\xa7\x40\x78\x99\xb1\x3c\x17\x92\xb1\x13\x37\x89\x43\xa9\x75\xec\xa4\xc9\xcc\xdd\x5c\x26\x56\x7b\xed\xdc\xdc\x03\x48\xac\x44\xc8\x20\xc1\x00\xa0\x64\x35\xe3\xaf\xd5\xa7\xbe\xdd\x17\xeb\x02\xa4\x24\x92\x12\x15\x5f\x33\x97\x71\x67\xc2\x07\x8b\x04\xf6\xff\x2e\x7e\xbb\x94\x1c\x3d\xb8\xfc\xe9\x62\xfc\xcf\xf7\xaf\x49\x6a\x32\x31\x3a\x88\x1e\xf8\x3e\xd1\x74\x0e\x8c\x4c\x94\xcc\x48\xa9\xc4\x70\xf0\xf8\xf1\xd3\xc7\x47\xa9\x31\xc5\x59\x18\x4e\xc1\xc4\x52\x1a\x6d\x14\x2d\x82\x44\x66\x21\xdc\xd0\xac\x10\xa0\xc3\x9c\xce\x63\xaa\x42\xe2\xfb\x28\xc6\x4a\x23\x82\xe6\xd3\xa1\x07\xb9\x37\x8a\x52\xa0\x6c\x14\x65\x60\x28\xb1\x82\x7c\xf8\x58\xf2\xf9\xd0\xbb\x90\xb9\x81\xdc\xf8\xe3\x65\x01\x1e\x49\xaa\xa7\xa1\x67\xe0\xc6\x84\x56\xc4\x4b\x92\xa4\x54\x69\x30\xc3\xbf\x8d\xdf\xf8\xcf\xbd\xd1\x01\xc1\xab\x92\xb3\xda\xf1\x4a\x33\xe9\x6c\xe5\x34\x83\xa1\x37\xe7\xb0\x28\xa4\x32\x0d\xc1\x0b\xce\x4c\x3a\x64\x30\xe7\x09\xf8\xee\xe1\x11\xe1\x39\x37\x9c\x0a\x5f\x27\x54\xc0\xf0\x18\x05\x55\x92\x0c\x37\x02\x46\x73\x48\x4e\x16\x10\x47\x61\xf5\x58\xef\xd9\x28\xbd\x5a\x85\x01\xc5\x2b\x20\x17\x57\x57\xce\x75\xb7\x2f\x78\x7e\x4d\x52\x05\x93\xa1\x17\x86\x19\xbd\x49\x58\x1e\xac\xc3\x66\x1f\x6c\xe4\xd6\x0b\xe1\x93\xe0\x49\x70\x12\x26\x5a\x6f\xd6\x82\x8c\x23\x95\xd6\x1e\x51\x20\x86\x9e\x36\x4b\x8c\x71\x0a\x60\xbc\xa6\x0d\x17\xa5\x36\x98\xa5\x6a\x97\x4c\xa4\x22\x26\xe5\x9a\x18\xc0\x94\x50\x03\x3b\x0d\xba\x6b\x22\xab\x8f\xfd\x36\x74\xc4\x6a\x94\x8b\xde\xcd\x74\x90\x08\x59\xb2\x89\xa0\x0a\x9c\x6c\x3a\xa3\x37\xa1\xe0\xb1\x0e\x67\x1f\x4b\x50\xcb\x4d\x30\xe6\x54\x70\x46\x8d\x54\xe1\xe3\xe0\x34\x78\xd2\x0e\xc2\xdf\x57\x9b\x77\x8b\xc6\xdb\xf1\x8f\x3f\x9c\x12\x9d\xf2\x8c\xd0\x9c\x91\x0f\xa0\x0b\x99\xb3\x60\x56\x45\xe6\xdd\xeb\xe7\x44\x97\x85\xad\x08\x22\x27\x35\x31\x08\xc8\xb0\x32\xb4\x63\xc8\x80\x71\x4a\xac\x81\x1c\xa3\xb9\x8e\x1d\x8a\xfe\x85\x4f\x88\x30\x28\x82\xbc\xf8\xb5\x5a\xc5\x75\x9d\x28\x5e\x18\xa2\x55\xb2\xf1\x5e\x6a\x1d\xd4\xf9\xb6\x7e\xdb\x1a\x3e\x45\x83\xe6\x98\xe2\x67\x98\xe2\xf5\xb3\x73\x68\xa6\xf1\x64\x84\x95\x98\xdf\x23\x55\x55\x8e\x85\xc7\xc1\x53\x94\x59\x3f\xf5\x48\x8c\x1e\xfc\x02\x39\xe3\x93\x5f\x2b\x77\xa2\xd0\x9d\x44\x1b\xb4\x28\x96\x6c\xb9\x0a\x1f\xe3\x73\x92\x08\xaa\xf5\xd0\xb3\x87\x85\xf2\x1c\xd4\x2a\xb6\x75\x74\xaf\x0c\x35\x3c\x21\x55\x59\xac\xa3\x83\x9b\xb8\xb2\xe2\xad\x37\xab\x0f\x9f\xc1\x84\x96\xc2\x78\x2b\xca\x1e\x3d\xfe\x44\x94\x9c\x35\xa8\xda\x74\xb5\x30\x6b\xb7\xb3\x89\x34\xae\x28\x2e\x8d\x91\x39\x31\x88\x1f\x43\xaf\x7a\xf0\x3a\x8c\x46\x4e\xa7\x02\xf0\x90\x0a\x41\x0b\x0d\xcc\x23\x58\x52\xb4\x5e\xb6\x66\x54\xeb\xab\x65\xaa\xa6\x16\x54\xbe\xab\xb8\x3d\x42\x15\xa7\x3e\xdc\x14\x58\x20\xc0\x86\xde\x84\x0a\x4b\xeb\x56\xad\x07\x4a\x8a\xb5\xaa\x8e\x71\x36\x99\xc8\xb6\x32\x47\x2b\x5f\xe6\x62\xe9\x8d\xc6\x95\x41\xc8\xc3\xa7\x18\x52\x99\x63\xc6\x90\x6e\x2f\x33\x47\x5d\xbe\x53\xf1\x75\x89\xa3\xb0\x0a\x6a\x67\x95\x76\x62\x1c\x2b\x0c\x8f\x57\x23\x41\x10\x7a\x1b\xd8\xa4\xad\xb4\x86\x98\xd7\xad\x3c\x73\xb6\x0e\x60\x47\xec\x2a\x37\xeb\xe4\x75\xd3\x5f\x8a\x06\xc7\xaa\xec\xf0\x63\x3b\x13\x82\x8f\xd0\xea\x8d\x81\x6f\x65\x06\xd6\xba\x08\x61\x69\x07\xf1\x4a\x2c\x4d\x0c\x9f\xa3\xda\x06\x2f\x8d\x65\x89\x35\x7d\x6e\x3f\x76\x4b\x88\xc2\x52\x6c\xbb\x6d\x0f\x51\x18\xa0\x71\x1b\xb7\x36\x87\xa8\x4d\xd4\x39\x1a\xcd\xc3\x66\x61\x79\x73\x2e\x1b\xc7\x44\xc9\x45\xef\x41\x13\x7e\xc6\xfc\x13\xaf\xab\x6c\x1f\xf5\xf3\xf6\x79\x4c\x4f\x56\x0e\xe3\xdd\x41\x73\xa7\x68\xfb\x3e\xb6\xfd\x07\x33\x4f\x68\x51\x08\x9e\xb8\xfa\x26\xa5\x46\x3c\xc5\x92\x40\x24\xd7\x16\x7b\xff\x0a\x2a\xc3\x0a\x5c\x48\xc5\xb0\x59\x49\x32\xe1\x39\x6b\x49\xd1\x3c\xe3\x19\x15\x35\x05\x82\x37\xcd\xf5\x02\xd4\x1a\x9b\x51\x86\x49\xc1\xc2\x7a\x46\x22\x3e\x3a\x27\xdc\xc9\x79\xd5\x49\x24\xd5\xe4\xc2\xfe\xc1\xad\x3f\x47\x21\x1f\x05\xe4\x0d\xca\xaa\x1b\xdd\x59\x2b\x43\xc5\x5e\xaf\x50\xc7\x1b\x2c\xf1\x6b\x05\x3c\x49\xad\xa4\x5a\x21\xae\xbf\x47\x28\xd0\x6e\x09\x15\xe1\xf3\x25\x94\x46\x27\x29\xce\x3e\xac\x26\x6c\x49\xaa\x98\x9c\x35\x51\xac\x48\xb8\xa5\xe7\x3c\xd3\x06\x14\xa3\x59\x5b\xcd\x39\x76\xd0\x86\x96\x0f\xd2\xec\xa2\xda\x23\xf7\x92\xe6\x1c\x04\x6d\xd3\x57\x8b\x4d\xeb\x7f\xfb\xb7\x32\xb0\x43\xe8\xef\x08\xd6\x18\x33\x63\x13\xb7\xce\xb8\x82\x02\x9b\x14\xb6\x59\x57\x0d\xb6\x3c\x14\xb4\x38\x12\x05\x38\xad\x30\x12\x2f\x5d\xd9\x2c\x79\x3e\xed\xe2\x4d\x67\xc8\x90\x0c\x82\xa9\x94\x88\xa3\xae\x23\x16\xa1\x55\x78\x82\x0a\xf1\x68\xaf\x6e\x3b\xe0\x43\x88\x6d\x14\x68\x5b\x55\x7d\xfd\x0a\x50\xfe\x62\xb1\x08\xf4\x44\x07\x65\xce\x7d\x53\x42\x8c\x06\x41\x1e\x30\x08\x21\x0f\xa9\x4e\x84\x6d\xba\xb2\x54\x09\x4e\x4b\x38\xfc\xe1\x48\x41\x43\x4b\x46\x7d\xb6\x08\x6c\x8f\x47\x90\xff\xed\x3f\xaf\xa8\x7f\x19\xfe\x6c\xad\xb0\x13\x62\x51\x6a\x2c\x3f\x00\x81\x83\x35\x80\x0d\xad\x3b\xe5\x89\xd9\x28\xcf\x28\x17\x46\x9e\x31\x97\x14\xd4\xe6\x5f\xcb\xeb\xbf\x74\x4d\xf0\x46\xfb\xf7\x9d\x3e\x3b\xef\x64\x76\x2a\xe5\xb9\x3d\x23\x2e\xee\x41\x37\x85\x9b\x87\xcf\x81\x41\x0f\x74\xd4\xb7\x8d\xfb\x6a\x52\x08\xd7\xf8\xe5\x90\xab\x26\x68\x0e\x36\xa1\x9b\x08\xeb\x0c\xd2\x82\xeb\x9d\x53\x22\x0e\x38\xc7\xc7\x38\xe1\xd4\x33\x63\xcf\x80\xd3\x96\x7b\xd7\x61\x7b\xd6\x9d\xb5\x3f\x2b\xf9\x0b\x07\xdb\x59\xef\x5c\xdb\xd6\xdc\x52\x5d\xcd\x35\xee\x75\x68\x46\xe7\xb4\x5a\x6d\x24\xe3\xe1\x80\xc9\xa4\xb4\x13\xec\x51\x80\x67\x88\x2d\x07\x93\x32\x4f\x1c\xe6\x0e\x8e\xc8\xa7\x56\x8d\x3f\x1c\x1c\x7e\x47\x73\x2a\xe4\x74\x89\x28\x98\x1d\x1e\x05\xdb\xf6\x0c\x3e\x75\x10\x94\xe0\x50\xac\x35\x9d\xc2\x19\x39\x74\xd0\x8e\x6e\x95\x60\xb1\x21\x97\x86\x38\x1f\x0f\x1f\x6d\xf1\x4c\x00\x58\x4c\x93\xeb\x77\x58\x08\xfa\x8c\x6c\x0b\xb5\x97\x63\x46\xb1\x53\xb1\x2c\x52\x3b\x85\x90\xf5\x9d\x2f\xaf\x77\x48\xb5\x17\xcf\xf7\xb1\x29\xc8\xe4\x1c\x7a\x58\xeb\x84\xe0\x39\xe9\xe5\xc6\x93\xa9\xd3\xc3\x2d\xee\xdb\x1d\x1e\xe2\x09\x64\xbd\xae\x59\x04\x3a\xee\xdb\xb4\x57\x33\xa8\xd8\xc6\xb8\xd2\xa6\x82\xcc\xcf\x04\xb6\xe3\x0b\x82\xeb\x3e\x2d\xf6\x42\x69\xaf\xb3\xc2\x2c\x3f\x47\xb7\xd7\x2a\x65\x5f\xd7\x15\x02\xb4\x7d\x4b\x4a\x68\x6e\x4d\x8c\x81\x80\x15\xbc\x1d\xae\xe6\x75\xdb\xbb\xbb\x7b\x67\x47\xa8\xed\xe5\x20\xfd\xee\x01\xd5\x80\xf9\x64\xf7\x2e\xa2\x1d\xb3\xee\x45\x54\x9f\xdc\x3d\xaa\x26\xe5\xea\xfe\x05\xb5\x6d\xd5\x57\x8c\xe9\x36\x4c\xb4\x56\x6e\x8f\x02\x99\x0f\x0e\x75\x99\xe0\x94\xa0\x03\xdb\x86\x83\x78\x7e\xf8\x88\xac\x10\x7a\x00\x5d\x84\xb6\x17\x04\x38\x2b\xcd\x11\xd0\x2f\xab\x97\xe7\xc1\xd1\xcb\x83\x2d\xa2\x87\x03\xef\x3b\x5d\xc6\x19\x37\xe7\x15\xa0\x7b\x47\xc8\x26\x8b\x81\xc7\xb8\xa6\xb1\xc0\x97\xdc\x47\xc4\xa8\x12\x90\x79\x17\x2f\xc2\x1c\x8a\x1e\x5b\x4a\xe4\xb4\x03\xcb\xc0\xf3\x76\xd2\x1e\xd6\xb4\xaf\x95\x92\x0a\x5b\x46\xca\x19\x0c\xf6\x52\xbe\x57\x72\x8a\x77\x1a\x89\x75\x2a\x17\x6b\xe2\x6d\x96\xc0\x36\xcf\x1d\x1d\xc7\x5e\xb6\xf1\x9d\x11\xef\xfd\x4f\x57\x63\x6f\x77\x7d\xd9\xb7\xf5\x33\xeb\x8d\x8d\x6c\xb3\xb1\xa1\x47\x1a\xdf\x12\xb0\xf4\xfe\x85\xa6\xf6\x33\x8f\x2b\x15\x33\x2d\xf3\x1e\x15\x09\x4d\x52\x24\x71\xaf\xfc\xbb\x29\x4a\x25\x50\x44\xad\xbc\x47\x4a\x5d\x01\x67\x64\xd3\x9a\xad\xfa\x5d\xc9\x5f\x5d\x3b\xa3\xd9\x17\xfa\xb5\x4f\x8a\x2e\x5c\x4a\x2b\xf1\xfd\x84\x77\xaa\x1e\xe7\x74\x8f\x90\x1e\x38\x01\x5b\x23\x4d\x37\x67\x1f\xff\xf1\xf6\xc3\x3e\x3f\xe7\x54\x59\x2e\x32\x24\x8e\x34\xa8\xbe\xca\xd2\x30\xc6\x81\x67\x9f\xfd\xff\x43\x74\x36\x65\xef\x4a\xd9\xeb\x54\xe7\x5e\x35\x8e\xe3\x47\x3d\x45\x35\x76\x14\x1b\xa0\xc9\x7f\x68\x78\xb7\x91\xa5\x43\x79\xdb\x05\x05\x6b\x2b\x4a\x37\x34\x4f\xe0\xdb\x70\xf7\xc7\x0f\x77\x77\x6f\x9a\xf7\xa9\x5d\x9e\x7f\xf5\x69\xee\x5e\xf4\xc8\x2b\x9e\x71\x7c\x61\xe3\xe6\x5b\x9b\x6c\x62\xc4\xd7\xef\x93\x2b\xed\xff\x7f\x8d\x72\x7f\x09\x7d\xeb\x95\x5f\xde\x2b\xbf\x20\xc2\x77\x6a\x97\x9b\x83\xd5\x6a\x9e\xeb\x04\x74\x8a\xa2\x95\x88\x87\x01\x60\x85\xbb\x8d\x06\x34\xf1\x9c\xc1\xcd\x23\xc2\x0d\x64\xdd\xbc\xd9\x5c\x29\xb9\xc0\x5c\xa1\x83\x91\xb1\x5f\xc7\x76\x71\x63\x1b\x5f\x68\x51\x40\xce\x06\xc8\xd7\x21\xc5\x95\xd5\xa6\x13\xc7\x46\x1e\xf9\xde\xe9\x0d\x7e\xb6\x78\xfe\x3d\xf1\xa2\xd0\xae\x1e\xdd\x9d\x71\x13\xee\x3e\xf6\xdb\xa3\x83\x76\x84\xbb\xbf\x54\x1e\x1c\x44\xee\x77\xd8\xd5\x77\x85\x65\x2c\x64\x72\xed\x17\x52\x1b\x21\x29\xf3\x8f\x29\x3c\xa3\xa7\x93\xe3\x27\x8f\x27\xc9\xb3\x17\xf1\xd3\x09\x8b\x9f\xd3\xa7\x27\xcf\x4e\xd8\x8b\xa7\x27\x7f\x8a\x4f\xbd\xd1\xe6\x77\x0e\x32\x22\x81\xfd\xaa\xf2\x95\xbc\xf9\x01\x41\xc2\x3e\xea\x94\x2a\x78\x57\xad\x1d\x7c\xc2\x5a\x28\x04\x5d\x9e\xe5\x32\x07\xf2\x80\x67\xf6\x57\x5c\x9a\x9b\x97\xb7\x68\x95\x35\x62\x14\x85\xee\xe7\xcc\x28\xac\xfe\x8f\xe1\xbf\x01\x00\x00\xff\xff\x18\x4c\xc3\x34\xd8\x20\x00\x00")

func static_about_html_bytes() ([]byte, error) {
	return bindata_read(
		_static_about_html,
		"static/about.html",
	)
}

func static_about_html() (*asset, error) {
	bytes, err := static_about_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/about.html", size: 8408, mode: os.FileMode(420), modTime: time.Unix(1426238000, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _static_index_html = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5a\x7b\x6f\xdb\x38\x12\xff\x3f\x9f\x82\x65\x7a\x88\x8d\x8d\xa4\xd8\x49\xae\x6d\x6a\x1b\x68\xbb\x2d\x5a\x60\x17\x2d\x36\xbe\x17\x16\xfb\x07\x25\xd1\x36\x5d\x4a\x54\x49\xca\x89\xaf\xc8\x77\xbf\x21\xf5\xb0\x9e\xae\x6f\xd3\xf6\x72\xd8\x15\xd0\xda\x22\x67\x86\x33\xbf\x19\xce\x0c\x19\x4f\x1e\xfd\xf8\xfe\xd5\xfc\x5f\x1f\x5e\xa3\x95\x8e\xf8\xec\x68\xf2\xc8\x71\x90\x22\x1b\x1a\xa2\x85\x14\x11\x4a\x25\x9f\x0e\xce\xce\x2e\xce\x86\x2b\xad\x93\x2b\xcf\x5b\x52\xed\x0b\xa1\x95\x96\x24\x71\x03\x11\x79\xf4\x96\x44\x09\xa7\xca\x8b\xc9\xc6\x27\xd2\x43\x8e\x03\x62\x8c\x34\xc4\x49\xbc\x9c\x62\x1a\xe3\xd9\x64\x45\x49\x38\x9b\x44\x54\x13\x64\x04\x39\xf4\x53\xca\x36\x53\xfc\x4a\xc4\x9a\xc6\xda\x99\x6f\x13\x8a\x51\x90\xbd\x4d\xb1\xa6\xb7\xda\x33\x22\x9e\xa3\x60\x45\xa4\xa2\x7a\xfa\xb7\xf9\x1b\xe7\x29\x9e\x1d\x21\x78\x32\x39\xc5\x0c\x4e\xf5\xa2\x31\x15\x93\x88\x4e\xf1\x86\xd1\x9b\x44\x48\x5d\x11\x7c\xc3\x42\xbd\x9a\x86\x74\xc3\x02\xea\xd8\x97\x53\xc4\x62\xa6\x19\xe1\x8e\x0a\x08\xa7\xd3\x11\x08\xca\x24\x69\xa6\x39\x9d\x6d\x68\x30\xbe\xa1\xfe\xc4\xcb\x5e\xf3\x39\x83\xd2\xcb\x02\x06\x10\x2f\x29\x7a\x75\x7d\x6d\x4d\xb7\xf3\x9c\xc5\x1f\xd1\x4a\xd2\xc5\x14\x7b\x5e\x44\x6e\x83\x30\x76\x4b\xd8\xcc\x8b\x41\xae\x1c\xf0\xce\xdd\x73\x77\xec\x05\x4a\xed\xc6\xdc\x88\x01\x95\x52\x18\x49\xca\xa7\x58\xe9\x2d\x60\xbc\xa2\x54\xe3\xaa\x0e\xaf\x52\xa5\xc1\x4b\xd9\x2c\x5a\x08\x89\xf4\x8a\x29\xa4\x29\xb8\x84\x68\xda\xa9\xd0\xa1\x8e\xcc\x3e\xf6\xeb\xd0\x10\xab\x40\x2e\x58\xb7\x56\x6e\xc0\x45\x1a\x2e\x38\x91\xd4\xca\x26\x6b\x72\xeb\x71\xe6\x2b\x6f\xfd\x29\xa5\x72\xbb\x03\x63\x43\x38\x0b\x89\x16\xd2\x3b\x73\x2f\xdd\xf3\x3a\x08\x7f\x2f\x26\x0f\x43\xe3\xed\xfc\xe7\x9f\x2e\x91\x5a\xb1\x08\x91\x38\x44\xbf\x50\x95\x88\x38\x74\xd7\x19\x32\xef\x5e\x3f\x45\x2a\x4d\x4c\x44\x20\xb1\xc8\x89\x29\xa7\x11\x44\x86\xb2\x0c\x11\x0d\x19\x41\x46\x41\x06\x68\x96\xd8\x81\xe8\x5f\xd9\x02\x71\x0d\x22\xd0\xb3\xdf\xb2\x51\x18\x57\x81\x64\x89\x46\x4a\x06\x3b\xeb\x85\x52\x6e\xee\x6f\x63\xb7\x89\xe1\x4b\x50\x68\x03\x2e\x7e\x02\x2e\x2e\xdf\xad\x41\x6b\x05\x3b\xc3\xcb\xc4\xfc\x37\x52\x65\x66\x98\x37\x72\x2f\x40\x66\xfe\xd6\x23\x71\xf2\xe8\x57\x1a\x87\x6c\xf1\x5b\x66\xce\xc4\xb3\x3b\xd1\x80\x36\xf1\x45\xb8\x2d\xe0\x0b\xd9\x06\x05\x9c\x28\x35\xc5\x66\xb3\x10\x16\x53\x59\x60\x9b\xa3\x7b\xad\x89\x66\x01\xca\xc2\xa2\x44\x07\x26\x61\xa4\xe0\xcd\x27\xb3\x0f\x27\xa4\x0b\x92\x72\x8d\x0b\xca\x9e\x75\x9c\x05\x4f\x59\x58\xa1\xaa\xd3\xe5\xc2\x8c\xde\x56\x27\x54\x79\x26\x7e\xaa\xb5\x88\x91\x86\xfc\x31\xc5\xd9\x0b\x6e\x30\x6a\xb1\x5c\x72\x0a\x9b\x94\x73\x92\x28\x1a\x62\x04\x21\x45\xf2\x61\xa3\x46\x36\x5e\x0c\x13\xb9\x34\x49\xe5\x38\xe3\xc6\x88\x48\x46\x1c\x7a\x9b\x40\x80\xd0\x70\x8a\x17\x84\x1b\x5a\x3b\x6a\x2c\x90\x82\x97\x4b\x35\x94\x33\xce\x04\xb6\x42\x1d\x25\x1d\x11\xf3\x2d\x9e\xcd\x33\x85\x80\x87\x2d\x01\x52\x11\x83\xc7\x80\x6e\x2f\x33\x83\xb5\x1c\xbb\xc4\xf7\x25\x9e\x78\x19\xa8\x8d\x51\xd2\xc0\xd8\x97\x00\x0f\xce\x33\x81\xeb\xe1\x5d\xda\x24\x35\xb7\x7a\xe0\xd7\x96\x9f\x59\x58\x02\xd8\x10\x5b\xf8\xa6\x74\x5e\xd3\xfd\x29\xaf\x70\x14\x61\x07\x1f\x6d\x4f\x70\x56\x50\x92\x40\xb3\x0d\x48\x02\x23\x76\xfa\xbe\x15\x11\x35\xca\x4e\x20\x4b\x75\xf0\x56\x89\x89\x2f\x52\x88\xe9\x17\xe6\xa3\x9b\x65\xe2\xa5\xbc\x6d\xb6\xd9\x44\x9e\x0b\xca\xed\xcc\xda\x6d\xa2\x3a\x51\x63\x6b\x54\x37\x9b\x49\xcb\xbb\x7d\x59\xd9\x26\x52\xdc\xf4\x6e\x34\xee\x44\xa1\x73\x51\xdf\x61\x2d\xe8\x20\xf4\x7d\x93\x60\x85\xd9\x15\xf0\x9d\x33\xa5\x9b\x78\x03\x8a\xd9\x7c\x02\x59\x07\xf2\xa6\x0d\x5f\xdc\x8b\xec\x71\x08\x32\x48\x1c\xb4\x37\xcc\x6e\xa2\x5c\xaf\xb1\x2f\xcd\xc8\xec\x9a\x45\x0c\x8a\x08\xd3\xdb\x1e\xa4\xbb\x15\xaa\x68\x40\x62\xc2\xc5\x72\xdb\x52\xa0\x1c\xdf\xbb\xfe\x8b\x8c\xaa\x6b\xf1\xcc\xc9\x3d\x39\x0b\x98\x9d\xbc\xe9\x68\x42\x68\xa8\xca\x35\x61\xbf\x51\x8e\xab\x5c\x66\x04\xe5\x40\xda\x9d\x51\x22\xd5\x0a\x4b\x28\x6a\x51\x2e\xca\x7c\xad\x93\xbf\x31\x23\x4d\x96\xba\x92\x86\xc9\x59\x4a\x91\x26\x1d\x84\x06\x5b\xe2\x53\x6e\x4a\x27\xf4\x4d\x42\x42\x82\xfe\x07\xfc\x0f\x30\x98\xe1\x4e\x06\x16\x27\xa9\xce\x73\xb1\x69\xe1\x70\x6d\xa5\x1c\x7b\x9c\x77\x67\x56\x24\x82\x36\x25\xa0\x2b\xc1\x21\xb5\x4f\xf1\x07\xf0\x91\xea\x52\xba\x99\x36\xf2\xe1\x3c\xf9\x1b\xb3\x55\xea\x47\x4c\xef\xc2\x05\xd7\xd4\xc8\xbe\x67\x34\xa5\x52\xbe\x8e\x11\xfc\xdb\x55\xa9\x6b\x3b\xdf\x9d\xf1\x8c\x12\xc6\x8a\xe6\x46\x6f\x29\x76\x88\x7f\x33\x4f\x15\x11\x78\x98\x5f\x73\xea\xaf\xef\xd6\x51\xe6\x57\x34\xfa\xaa\x9e\x1d\xdd\xd3\xb5\xbf\xdf\x9e\x71\x6e\xcf\xf8\xab\xda\x33\x6e\xd8\xf3\x46\x76\xee\xc9\x6f\x61\xd0\x79\x6e\xd0\xf9\x57\x35\xe8\xbc\x61\xd0\x4b\x2a\xa1\x99\xbf\xcf\xe6\x7b\x51\x64\xd4\xff\xc9\xce\xcb\x07\x8e\x7a\x08\x3a\xea\xe1\xd3\x7a\x3d\x5c\x8d\x67\x70\x60\x00\x75\xa0\x45\x1e\xb7\x13\xbb\xb1\x53\xda\xf9\xf7\x69\x2b\xa7\x9b\x6a\x49\x2b\x24\x73\xf3\x5e\xdd\xf6\x30\x6b\xff\x77\xe0\x58\xc3\x12\x1a\x36\x05\x78\x76\x76\x9f\x3d\x75\x13\x12\x29\x96\xb0\x14\x94\x6b\x7b\x14\xb2\x69\x1f\xdc\xb9\xbd\x42\xb1\x28\xf2\x4b\xa6\xca\x87\x82\xb2\x23\x4d\x35\x84\x99\x2e\x10\x55\x5f\x0a\x6d\xcb\x7a\x54\x94\xda\x8c\x64\xd7\x1e\xc3\x21\x2e\xa5\xb1\xb8\x99\xe2\xd1\xd9\x59\x75\x0c\xce\x24\x53\x5c\x1f\x21\xb7\x39\x55\xae\xb9\x3d\x85\x5f\x21\x18\xfa\xcb\x81\xfd\x33\x98\x14\xc0\xf2\x2c\x5e\xc2\x69\x10\x4e\x6b\x4a\x77\x77\xae\x5f\x88\x90\x96\x63\x5f\x4b\x29\x64\x1f\xa2\x45\x8b\xc3\x29\x1c\x1f\xed\xff\x4e\x48\xe2\x25\x9c\x48\x72\x54\xec\x58\x13\xe5\xa4\x65\x91\x3f\xb3\xcb\x18\x9c\x0b\x23\xcc\x76\x54\x57\x10\xfd\xb3\xdc\xe0\x86\x4a\x3f\xab\x25\xf6\x9a\xc6\x25\x07\x06\x7f\x75\x2a\xfb\x9e\x9d\xe6\xbc\xb2\xc7\xb4\xdd\x65\x4e\x50\x3d\x7c\x7a\xf6\xd4\xee\x2e\x85\x80\x7e\x88\x24\x4c\x75\x9e\xe4\xe1\x10\x3a\x1a\xc1\x29\x34\x3f\xd7\xf7\x1c\x42\xeb\x72\x0f\xbd\x10\x59\x37\xef\x43\xbe\x28\xf9\x9e\x97\x0f\xeb\xde\xbb\x87\xfa\xca\xb5\xa5\x77\xe9\xce\x5b\x93\x0d\xc9\x46\x2b\x81\xf0\x78\x10\x8a\x20\x35\xb7\x0c\x43\x57\xc2\x21\x76\x3b\x58\xa4\x71\x60\xfa\x54\x34\x18\xa2\xcf\x35\xbf\x3e\x1e\x9c\x1c\x57\x2a\xfd\xc9\xd0\x6d\xeb\x33\xf8\xdc\x4a\xc6\x11\x44\x12\x59\xd2\x2b\x74\x32\x37\xd7\x3f\x76\xa3\x21\xf8\x12\x0b\x8d\xac\x8d\x27\xa7\x2d\x9e\x05\xa5\xa1\x4f\x82\x8f\xef\x20\x10\xd4\x15\x6a\x0b\x35\x8f\x65\x06\xb1\x4b\xbe\x4d\x56\xe6\xa4\x88\xca\x6f\x8e\xf8\xd8\x21\xd5\x3c\x2c\xde\xc7\x26\x69\x24\x36\xb4\x87\x35\x77\x08\xec\x8a\x5e\xee\x05\xec\x8c\xd5\x49\x8b\xfb\xae\xc3\x42\x46\x79\xd8\x6b\x9a\xed\x52\xfa\x26\xcd\x53\x05\x95\x82\x30\xa9\xb4\x65\xfa\x12\xb0\x0d\x5b\x84\xec\x55\xa1\x78\x40\xda\xeb\x28\xd1\xdb\x2f\xd1\xed\xd5\xca\xa4\x42\x26\x4d\xb2\x8e\x43\x14\x90\xd8\xa8\xe8\x53\x44\x8d\xe0\x36\x5c\xd5\xe7\xae\x77\xb6\x7b\xa6\x03\x6a\xf3\xd8\x36\xe9\x70\x40\x15\x05\x7f\x86\x0f\x0e\xd1\x86\x5a\x0f\x02\xd5\xf3\xc3\x51\xd5\x2b\x26\x1f\x1e\xa8\x75\xad\xbe\x23\xa6\xed\x34\x51\x1b\xb9\x1b\xba\x22\x1e\x9c\xa8\x34\x30\xb5\xd8\x35\x6d\xa6\xeb\x6f\x4e\x4e\x51\x91\xa1\x07\xb4\x99\xa1\xcd\x43\xdd\x44\xd2\x0d\x24\xf4\x1f\xb3\x06\x76\x30\x7c\x7e\xd4\x22\x7a\x3c\xc0\xc7\xf5\xc6\x78\x08\x6c\x22\x19\x98\xce\xc2\x74\x7b\x21\x3e\x45\x5a\xa6\x14\x98\xbb\x78\xab\x9d\xe4\xd0\x35\x17\xc7\x03\x8c\x3b\x69\x4f\x8e\x2b\xcd\x02\x94\x8c\x15\x0b\xe9\x60\x2f\x65\xd1\x14\x02\xb1\x5a\x89\x9b\x92\xb8\xcd\xe2\x9a\xe2\xd9\x51\x71\xcc\x63\x0a\xdf\x15\xc2\x1f\xde\x5f\xcf\x71\x77\x7c\x99\x0b\x95\x2b\x63\x8d\x41\xb6\x5a\xd8\xc0\x22\x45\xa1\x27\xe4\xec\xdf\xa0\x6a\x3f\xf3\x3c\x5b\x62\xad\x44\xdc\xb3\x44\x40\x82\x15\x90\xd8\x6b\xd9\x6e\x8a\x54\x72\x10\x51\x9c\xb6\xbb\x69\xf2\x08\xb8\x42\xbb\xd2\x6c\x96\xef\x72\x7e\xf1\x74\xa2\xd9\x07\x7d\x69\x93\x24\x37\xd6\xa5\x99\xf8\x7e\xc2\x83\xa2\xc7\x1a\xdd\x23\xa4\x27\x9d\x50\x13\x23\x55\x33\xd7\x9f\xfe\xf9\xf6\x97\x7d\x76\x6e\xe0\x50\x00\x5c\x68\x8a\x2c\xa9\x9b\xfd\xb9\x41\xd1\x39\x34\x3c\xfb\xf4\xff\x1d\xe8\xec\xc2\x3e\x6b\xc5\x1b\xd1\xb9\x77\x99\xa2\x53\x86\x65\x4c\x2b\x36\x00\x95\xbf\x29\xbc\xed\xcc\xd2\xa0\xbc\x6b\x26\x05\xa3\x6b\xf5\x76\xee\xcf\xe6\xee\xdb\x36\x77\x87\x17\xcd\x87\x54\x2e\x5f\x7c\xf7\x6e\xee\x41\xd4\xc8\xca\xcd\xed\x1f\xbd\x4c\xd6\x6e\xf0\xbf\x7b\x9d\x2c\xff\xdc\xf0\x7f\x57\x28\xf7\x87\xd0\x9f\xb5\xf2\xfe\xb5\xf2\x1e\x08\x1f\x54\x2e\x77\x1b\xab\x56\x3c\x4b\x07\x34\x82\xa2\xe6\x88\xc7\x2e\x85\x08\xb7\x13\x95\xd4\xc4\xe2\x90\xde\x9e\x22\xa6\x69\xd4\xf4\x9b\xf1\x95\x14\x37\xe0\x2b\x30\x70\xa2\x25\xf2\x66\xcd\xbc\xd1\xce\x2f\x24\x49\x68\x1c\x0e\x80\xaf\x41\x0a\x23\xc5\xa4\x15\x17\xce\x30\xfa\xc1\xae\xeb\xda\xcb\xfb\x1f\x10\x9e\x78\x66\x74\x78\x38\xe3\x0e\xee\x3e\xf6\xbb\xe1\x51\x1d\xe1\xe6\xaf\x49\x8e\x8e\x26\xf6\x3a\xb3\xb8\xbe\x4c\x7d\x2e\x82\x8f\x4e\x22\x94\xe6\x82\x84\xce\x88\xd0\x27\xe4\x72\x31\x3a\x3f\x5b\x04\x4f\x9e\xf9\x17\x8b\xd0\x7f\x4a\x2e\xc6\x4f\xc6\xe1\xb3\x8b\xf1\x5f\xfd\x4b\x3c\xdb\xfd\x2d\x1a\xcd\x90\xcb\xe2\x85\x78\x29\x6e\x7f\x82\x24\x61\x5e\xd5\x8a\x48\xfa\x2e\x1b\x3b\xfa\x5c\xdc\x98\x9a\x0b\x53\xf4\x88\x45\xe6\x97\x36\x24\xd6\xcf\xef\x40\x2b\xa3\xc4\x6c\xe2\xd9\x9f\x9c\x4c\xbc\xec\xb7\x66\xff\x09\x00\x00\xff\xff\x7e\x3e\x92\x78\x7c\x26\x00\x00")

func static_index_html_bytes() ([]byte, error) {
	return bindata_read(
		_static_index_html,
		"static/index.html",
	)
}

func static_index_html() (*asset, error) {
	bytes, err := static_index_html_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "static/index.html", size: 9852, mode: os.FileMode(420), modTime: time.Unix(1426238000, 0)}
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
	"static/about.html": static_about_html,
	"static/index.html": static_index_html,
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
	for name := range node.Children {
		rv = append(rv, name)
	}
	return rv, nil
}

type _bintree_t struct {
	Func func() (*asset, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"static": &_bintree_t{nil, map[string]*_bintree_t{
		"about.html": &_bintree_t{static_about_html, map[string]*_bintree_t{
		}},
		"index.html": &_bintree_t{static_index_html, map[string]*_bintree_t{
		}},
	}},
}}

// Restore an asset under the given directory
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

// Restore assets under the given directory recursively
func RestoreAssets(dir, name string) error {
        children, err := AssetDir(name)
        if err != nil { // File
                return RestoreAsset(dir, name)
        } else { // Dir
                for _, child := range children {
                        err = RestoreAssets(dir, path.Join(name, child))
                        if err != nil {
                                return err
                        }
                }
        }
        return nil
}

func _filePath(dir, name string) string {
        cannonicalName := strings.Replace(name, "\\", "/", -1)
        return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

