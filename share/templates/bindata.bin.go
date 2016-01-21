// Code generated by go-bindata.
// sources:
// index.html
// DO NOT EDIT!

// +build bin

package templates

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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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

var _indexHtml = "\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x1d\x6b\x6f\xdb\x38\xf2\xfb\xfe\x0a\xae\xb0\x7b\xd8\x2e\x2a\x29\x71\xd2\xbd\x5e\xd7\xce\xa1\x4d\x7c\x3d\xe3\x92\xd4\x68\x92\xde\x2d\x16\x8b\x03\x2d\xd1\x36\x1b\x59\xd2\x92\x54\x1e\x67\xf8\xbf\xdf\x50\xef\xa7\x2d\xbf\x64\x37\x49\x81\xa4\x12\x35\x6f\x0e\x87\x33\x43\x39\x9e\x4e\x4d\x32\xa4\x36\x41\x4a\xf0\xff\x7f\x3d\x57\xd0\x09\x51\x66\xb3\x36\x77\xb1\x7d\x32\x9d\x6a\x67\x58\x60\x2d\x18\x86\x51\x3d\x1c\x26\xb6\x39\x9b\x15\xb0\xc7\x0e\x17\x36\x0e\xf0\x31\x1a\x33\x32\xec\x28\xba\x82\x04\x15\x16\xe9\x28\xd1\x53\x14\x51\x8d\x06\x66\x33\xe5\xa4\x38\xd6\xd6\x71\x25\x23\x0b\x2b\x72\xf4\x07\x17\xbd\xeb\xa0\x00\xd1\xc5\x0c\x4f\xf8\x6c\xf6\x5d\xdb\xa4\x77\xc8\xb0\x30\xe7\x1d\x65\xc4\xa8\xa9\x0e\x2c\xc7\xb8\x05\x69\x54\xe1\xb8\xca\xc9\x77\x08\xa5\x21\x0c\xc7\x52\x2d\x86\x2c\xcc\x46\x44\x3d\x44\x82\x3c\x08\x95\xd1\xd1\x58\x28\x27\x3e\xd8\x74\xea\x03\xce\x66\x1d\x65\x3a\xa5\x43\xf4\x83\xab\x9d\x63\x5b\xbb\x24\x23\x2c\xe8\x1d\x88\xc9\xc7\xce\xbd\x2a\x7f\x8d\xa9\x49\x40\x5e\x8b\xc3\xa0\xbc\x4e\x0f\x4a\x25\x7c\xd6\xc0\x7c\x7c\x18\xf1\x1e\x1f\x07\xfc\x9c\x3b\xc2\x86\x96\x73\x1f\x42\x00\x0c\x06\xbe\xff\x04\xf3\x5d\x03\xbf\x90\xe5\x6c\x16\xd9\xf1\xdc\xc1\x26\xc2\x77\x23\x64\x52\xee\x5a\xf8\x11\x39\x30\x3b\x8e\xcd\x13\x74\x39\x49\x69\xc9\x23\x49\x54\xf9\x0b\xc1\xa5\x20\x00\x62\x10\xd5\x65\x04\x01\x67\x2c\x54\x8b\x0c\x41\x65\x4d\xd3\x50\x38\xc7\x11\x17\x39\x0d\x81\xe0\xfa\xf8\xf0\xa4\xad\x83\x51\x7c\x1b\xc6\x17\x95\xc6\x3c\xdc\x9a\x09\x4b\x67\x38\x51\xdf\xb3\xa2\xa7\x13\x62\x7b\x28\xd6\x5e\x5e\xc4\x50\x00\x67\xd1\xe4\x26\x4b\x95\xda\xae\x27\xd4\x11\x73\x3c\x57\x49\xc3\x54\x42\xa9\x16\x1e\x10\x0b\xf1\x09\xb6\xac\x60\x52\x6d\xe7\x9e\x61\xc0\x36\x09\x4c\x51\x6c\xad\xc5\x84\x06\x9e\x10\x8e\x0d\xa6\xc3\xd1\xd3\x60\x24\xa4\xcd\x89\xe1\xd8\x26\x66\x8f\x72\xf6\xf1\xc0\x22\xa6\x5c\x3b\xbe\x49\xcd\x60\xcd\xac\xc0\x6b\x3a\xbd\xa7\x62\x8c\xce\x09\xe7\x67\xb1\xc7\x99\x28\xcb\x39\xad\x15\x2c\x33\x94\xf9\x07\xd2\x3a\xf6\xa9\x45\x8d\xdb\xce\x54\x8c\x29\xd7\xc6\xd8\x36\x2d\xe2\x8f\xcc\x82\x48\x00\x2b\x5c\xba\x34\xf8\x71\xe2\x10\x30\x76\x1a\x5c\xe6\xe8\x21\x24\x23\xc2\x07\x6c\x8e\x88\x8f\xa0\x5d\x03\xf7\x4c\x4c\xc8\x2a\xb8\xb2\xce\x17\x0e\x23\x7b\xa5\x73\xa0\x28\x4a\xb4\xaf\xad\x73\x6e\xa0\xad\xa7\xfd\xbb\x19\x67\x67\xce\x3d\x6f\xc4\xd7\x6d\xed\xfd\x80\x3b\x96\x27\xc8\x26\x9c\xfe\x32\x09\xb3\x39\x07\xe0\x9e\x61\x00\xc0\x13\x76\xfe\xbd\xd4\x7d\x2b\x8b\xa0\xad\x7b\x56\xb4\x95\x49\xa0\xd2\xcd\x89\xfc\x89\x72\x0e\x86\x0e\x82\xcd\xa8\xce\x1e\x84\x60\x27\x17\xd4\xc0\x60\x3c\xc1\xa8\x4b\x92\x3d\xa9\x1c\x5c\x8c\x09\x06\xa7\x0e\x36\xec\xf0\x31\x79\x80\x1b\x13\xc1\x5e\x1a\xcc\x83\x0a\xfb\x68\x9f\x30\xea\x98\xe1\xce\xbc\x00\x3a\x93\xc3\xa0\x2f\xd8\xf2\xc8\x7c\x44\xe0\x1f\x3e\x4f\xdb\xce\xb7\x46\x90\x5a\x41\x82\x60\x42\x26\x10\x9b\x7e\x3a\x65\xd8\x1e\x11\xf4\x83\x85\x93\xfc\x2b\x04\xd2\xce\x29\x97\xbe\x9a\xb3\x6d\x7a\x9b\x5e\xa8\xed\x77\x65\x6e\x9c\x68\x85\xee\xa9\x29\xc6\xea\x11\x23\x13\x3f\x1e\x58\x58\x0b\xec\x33\x9b\x4d\xc2\x99\x5d\xc1\x50\x01\x21\xdf\x5c\x71\x9e\x5b\x41\x00\x72\x9c\x9c\xcd\xb2\xae\x99\xbe\x8b\x04\x4a\x7e\x27\xc9\x53\x06\xb7\x90\xe3\x4e\x40\xbd\xdd\x25\xb9\x17\x64\xd2\x78\x96\x2b\x79\x26\x69\x2e\xdc\x39\x41\xe8\xdf\x5e\x92\x1b\xf0\xd8\x4a\x8a\xbb\xae\x01\x5f\x72\xdc\xe2\xbe\x0f\x36\xdd\x68\x92\x2b\xe9\xed\x3e\xe3\x6b\x38\xcb\xdd\x0f\xa5\x5f\xd2\xdc\x5a\xee\xbe\x9d\x3c\x57\x12\xde\x9f\x64\xaf\xe1\x44\x77\xbf\x94\xdf\x75\xa6\x9b\xf1\xb1\x7d\x48\x75\xa3\x2d\x79\xf9\x0c\xee\xda\x11\xd8\x5a\x05\xf1\x86\x93\x95\x72\xeb\x7f\x30\xb2\x20\xb5\xae\xe4\xf7\xe3\x9a\x29\xf9\xc4\xb7\x52\x31\x23\x87\xf1\x24\x5b\x0c\x80\x36\x91\x91\x43\x38\x02\x62\xda\xbf\xa8\x6d\x2e\x4a\x8e\x2b\xb3\x6b\x49\xc0\x9f\xa2\xb5\x28\xc8\xb9\x5a\x8b\x80\x9c\xb4\x1c\x81\x94\x55\xe6\x12\x41\x83\x91\xea\x71\xe2\x1a\x42\x41\x26\x18\x38\xbc\xe9\x24\xa2\xf5\x0d\xb0\x74\x5a\x56\x7f\x60\xc1\x64\x37\x53\x4e\x98\xc3\x1d\x56\x13\x67\xc3\xc6\x8b\x09\x60\x99\xd4\x12\x67\x94\xdf\x22\x8f\x63\x58\x22\x5b\xad\x27\x12\x3e\x5b\xa9\x29\xd6\x34\xe3\x4b\x49\x51\xcc\xb1\xce\x86\x1b\xad\x28\x80\xdc\xee\x73\xeb\x86\x0b\x8a\xbd\xd0\xf9\xa5\x9e\xa8\xe3\xeb\xdb\x29\x27\x80\xee\xfe\x24\xd4\x0d\x57\x13\x7b\xa5\xfb\xae\x8b\x89\xb4\x83\x35\x50\x4b\x4c\xa7\xfa\xcf\xc8\x75\x6e\x09\x72\x3d\x46\x87\x8f\x06\xe7\xef\x32\x28\x26\x73\x5c\xb9\x3a\x33\x9b\x2d\x66\x04\xf2\x37\x3d\xf4\x7d\xf4\xb3\x1e\xbe\xec\xb0\x28\x8f\x8c\x16\x6b\xe8\x04\x5f\x2c\x6a\xdf\xc6\x4e\x70\x8b\x0e\x91\x72\x46\xee\xa8\x41\xb2\x93\x5d\x6f\xa2\xcf\x81\x56\xd9\x64\xcb\xf1\xb2\x09\x3f\x09\x9f\x85\x1e\x9e\x4b\x5d\xa4\x97\x48\x2d\x43\xcc\x24\xb7\x8c\x9d\x61\x99\xfc\x79\x81\xde\x2d\xa4\x5c\x38\x9e\x2d\x20\xca\x3c\x49\xc5\x33\xf9\x6d\x85\x0d\x7e\x41\x8a\x5f\xd9\x3c\x5f\x0b\xbc\x41\x8a\xac\xcc\x9e\xaf\x01\x8e\x90\xf2\xfe\x0e\xd3\x67\xec\x02\xc7\xbe\x0b\xfc\xf8\xed\x1b\x60\x61\xf3\x05\xd2\xaa\xdb\x1b\x59\xe4\x15\xfb\x2f\xe6\x30\x29\xaf\x63\xb0\xf5\x3b\x30\xf9\xa2\x17\xd2\x39\x73\xa8\xc1\x7e\x73\x89\x27\x0b\x8f\x2b\xab\x88\xa0\x90\x0a\x65\xab\x53\x49\x77\x47\x46\x0c\x3f\xe6\x34\x09\xca\xe4\x54\x07\x06\xf8\x95\xf5\x80\xc2\x32\xbd\x67\x3b\x26\x81\xcc\xc9\x8f\xa5\x11\x7c\x30\x98\x20\x34\x23\x65\x49\x9f\x29\x2b\xa4\x27\xc3\x5d\x2c\xa3\x97\x01\x6f\x46\x44\x3f\xde\xcc\x95\x71\xc8\x08\x49\x64\x1c\x66\x3a\x5f\x4b\x36\xc0\xaa\xe4\x9c\xd3\x13\x0b\xac\x18\xb4\xc4\xea\xd8\xbb\xa4\x57\x56\x30\xf9\x8f\x19\x93\x67\x31\xf6\xa0\xc9\x66\xb8\xde\x0e\xbb\x6c\xa7\xfd\x9b\xc6\xdb\x6c\x92\x67\xd2\x67\x83\xbb\xed\x36\xd8\x80\xc1\x56\x3a\x6b\xeb\x9a\xee\xa5\xb5\x56\x6c\x37\x80\x4d\x37\xda\x5b\x93\xf4\x76\xdf\x68\x6a\xb8\xb9\xb6\x1f\x4a\xbf\x74\xd7\x6a\xb9\xfb\x76\xda\x6b\x92\xf0\xfe\xf4\x98\x1a\xee\xaf\xed\x97\xf2\xbb\x6e\xb0\x65\x7c\x6c\x1f\x4e\xeb\x4f\x61\xaa\x56\x3c\x02\x67\x0b\x8e\x45\xcb\x11\xaf\x1e\xf9\x4a\x78\xff\xc6\x54\xac\x84\xd8\x03\xaf\x59\xf7\xb4\x1e\xd2\xc2\x62\xa9\x08\x83\x49\x92\x08\x37\x9b\xaa\x12\x93\x7e\x9d\x64\xa1\x5d\xae\x7f\xd8\x1d\x79\x6d\x2e\xbf\x97\xd4\xe5\x3c\x26\x87\xde\x99\x91\xc8\x68\x55\xe9\xdf\x06\xb8\x83\x33\x64\x99\x47\x03\x0d\xf0\x96\x0e\x95\x65\x1e\x8f\x6c\x88\xbb\x4a\x6d\x58\xb5\x9c\xcc\x93\x42\x7a\x67\x56\x8a\x78\x64\x2f\xde\x3a\xa0\xbb\x7c\xeb\xa0\xd7\xfc\x5b\x07\xbd\xf4\x5b\x07\x3d\x5b\x00\x3c\x54\x34\x7c\xbb\x45\x51\xc2\x67\x2b\xb5\xd1\x9a\x66\x7c\x29\x8d\x8a\xb9\x62\x6f\xb3\x6f\x1d\xf4\xf6\xe1\x04\xbe\xe1\xc2\x68\x2f\x74\x7e\xa9\x8b\xea\xf8\xfa\x76\xca\xa2\xde\x3e\x9d\xbc\x37\x5c\x15\xed\x95\xee\xbb\x2e\x8a\x7a\xcd\xbe\x75\xb0\xb0\x26\x8a\xb7\xe3\x95\xca\x8d\x7e\x59\x03\xfc\x8c\x39\x2e\x7f\xdd\x65\xcc\x61\x1c\xf5\x6c\xfd\x93\x27\x90\x4b\x58\xb8\xe8\x94\x4a\xe2\xad\x74\x62\x99\x59\xfc\xe8\xdc\x01\x6f\xe9\x7d\x42\x2e\x2f\x63\xd8\xc7\xc6\x2d\x11\x9b\xe2\x15\x53\xab\x62\xf7\x81\x6e\x8c\x17\xf0\x68\xd3\x93\x41\x5b\xa7\x27\x39\x5e\xcb\xd7\x6d\x74\x88\x4d\x93\xf1\x62\xed\x46\x53\xc7\x7c\x21\xd0\x56\x0e\xf9\xe8\x50\xdb\xc8\xd9\x5c\x44\xac\xd7\xaf\x49\xaa\xda\xc0\x7f\xb1\x07\xdc\xfd\x35\x33\x7b\xfe\x89\x1e\x32\x7d\x1f\x25\x81\x8f\x4e\x1c\xd3\xb3\x1c\x74\xfc\x51\x49\x16\x52\x72\xa8\x55\xc4\x83\xc9\x4f\xe1\x84\xd2\xfa\x5e\xdf\xb3\xab\x0b\xd8\x38\x08\x44\xc0\xd2\x7f\x14\x58\x83\xb9\x20\xa0\xaf\x49\xa0\x4c\x60\x09\x59\x2e\x31\x3c\x89\x45\x7e\x5d\xa2\x30\x89\x96\x71\x11\x3f\x58\xe1\x29\x95\xf5\x6a\xfc\x52\x01\x02\x02\x69\x09\xc2\xff\xd2\xd3\x96\x37\x80\x9c\xfa\x9f\x22\x2b\x10\x4b\xe0\xd0\xf0\x48\x39\x50\x5e\xa1\x9f\x1c\x56\xf2\x34\x32\xd4\xab\xaa\x67\x80\x9a\x7b\x16\x29\x17\x92\x2d\x3e\x8a\xf0\x66\x33\xe9\x2c\xf1\xec\x15\x23\xa1\xb4\x5d\x2a\x4c\x44\xc6\x4f\x89\x5e\xd3\x67\x0a\xea\xac\xe4\x38\x0b\xa8\x64\x45\xcf\x85\xb8\x82\xec\xd5\xde\x13\x87\xff\x2a\xfc\x05\xde\xd3\x4d\xfc\x66\x2e\x81\x12\xef\xd9\x60\xcc\x58\x14\x0e\xdc\x78\xe7\x29\x7a\x77\xb8\x8f\x2c\x58\x1f\x11\x85\xd2\x05\x12\x92\x58\xb0\x42\x0a\xbb\x60\x95\xc5\x16\x49\xd4\x4f\xc9\x32\x9f\x44\x99\xd5\x1b\x30\xf7\x87\xdf\xae\xbb\x57\xa5\xc6\xfe\xf0\x28\xc8\x22\x53\x07\xd8\xa5\x86\xf6\xd1\xeb\x99\xf9\x43\xef\xfa\x6a\x8e\x8d\x65\x72\x50\x29\x87\x8f\x3b\xc7\xba\x12\xb9\xbe\x69\x9b\x69\x8a\xb9\x7c\x87\x4d\xb1\x3e\x6f\xbc\x29\x06\x2c\x93\xa6\x58\x9f\x39\xb2\x50\xd9\x76\x4f\x2c\x66\xb3\x95\x96\xd8\x9a\x46\x7c\x69\x89\x15\xdb\x04\x7d\xbe\xd1\x96\x18\x90\xdb\x7d\x7b\xa8\xe1\x96\xd8\x5e\xe8\xfc\xd2\x12\xab\xe3\xeb\xdb\x69\x89\x01\xdd\xfd\x69\x0b\x35\xdc\x12\xdb\x2b\xdd\x77\xdd\x12\x4b\x3b\xd8\x13\xfd\x20\xce\xa2\x57\xd2\xfb\xdc\xff\x4c\x4e\xbf\x77\xf6\xed\xbf\x91\xbe\xaa\xfe\x2d\xa4\xdc\x3c\x55\xfd\x2b\x95\x3e\x02\xa5\xaf\xba\x9f\x9f\xa6\xd6\x75\x66\xfd\x18\xbc\xfe\x19\xab\xff\x06\x29\x97\xbd\xe7\xab\xfe\x2f\x48\xf9\xd2\xfb\x7c\xfd\x7c\x0d\xf0\x57\xa4\x7c\xee\x5e\x3d\x65\xfd\x0d\x22\x4f\x98\x2a\x0d\xf0\x16\x29\xd7\xbd\x8b\xee\xd3\xb4\x40\xa5\xd6\x7f\x43\xca\xe9\xa7\x8b\x8b\xf7\x97\xdf\xd4\x7e\x57\x71\xf0\xe4\x32\xc7\x28\x39\x76\x72\x79\xaa\x63\x24\x41\x36\x75\xe8\x94\x1c\x12\xb9\x5c\x83\x94\x69\xe5\xbf\xb0\x03\xe8\x37\xf5\xd1\x13\x1c\x4e\xd8\x3a\x3c\xfb\x8c\x3a\x8c\x8a\xc7\x75\x68\x5c\x52\x63\xc9\x83\xb6\x2c\xfe\x15\xfd\xdf\x5a\xf8\x9f\x09\x87\x2c\xdd\x16\xcb\xd1\x88\x62\x41\x48\xe4\x9a\x2e\x3e\x2d\x4c\x34\xce\x1c\x2d\x6e\xaa\xdd\xd9\xfe\xfe\xec\xd3\xe9\xf5\x6f\xfd\x2e\x1a\x8b\x09\x14\x2c\xed\xe0\x3f\xd9\x3b\x84\x62\x21\xac\x3e\x26\x44\x40\xcd\x3c\xc6\x8c\x13\xd1\x51\x3c\x31\x54\xdf\x46\x85\x89\xdf\x27\x94\xeb\x2b\xf7\x7d\x00\x08\xae\x40\xd9\xb6\x1e\x00\xa4\xe8\x8c\x85\x70\x55\xf2\xa7\x47\xef\x3a\xca\x7f\xd4\x9b\xf7\xea\xa9\x33\x71\xb1\xa0\x50\x71\x2b\xa0\xb0\x2d\xd1\x3a\x4a\xaf\xdb\x21\x50\x8b\x29\x69\x4c\x49\xb9\xa3\xdc\x51\x72\xef\x3a\x4c\xa4\x80\xfd\x3f\x31\xdb\x31\xfd\x0f\xf4\xab\xfe\xcd\x6b\x44\x6d\x2a\x28\xb6\x54\x0e\xf5\x11\xe9\x1c\x6a\x07\xb0\xb8\xbe\x57\xd5\xdf\x61\xd5\xf6\xba\x7f\x84\x64\xfd\xa8\xc4\x88\xe5\x77\x2d\x99\x30\x3c\x81\x28\x50\x55\xa2\x6f\x43\x18\xe2\x3b\x79\xaf\xc1\x2f\x89\xfe\x3b\x98\x8c\x0e\xff\x50\xd5\x02\x7a\x80\xc5\xc1\xa9\x60\xf2\x8e\x5a\x0f\x47\xad\x98\x06\xd8\xe6\x93\x6f\x8b\x2f\x84\x71\xea\xd8\xb3\x59\x44\x96\x47\x17\xea\x51\x4b\x73\xed\x91\x82\xc4\xa3\x0b\x1a\x5e\xe2\x4b\xa5\xc0\x01\xbb\xae\x45\x54\xe1\x78\xc6\x58\xcd\x70\x3b\x6c\x1d\x3c\xc0\xcf\x52\xfc\x00\xde\x67\xb8\x04\x97\xe3\xe3\x07\xf8\x59\x8e\xcb\xf1\xf1\xb2\x5c\xde\xb4\x1e\xe0\x67\x39\x2e\x6f\x5a\xcb\x72\x79\x0b\x16\x7b\xbb\xa4\xc5\xde\x56\x58\x8c\x8b\x47\x8b\xf0\x31\x91\xd5\x71\x30\x7d\x72\xa5\xeb\x50\x54\xcf\xa3\x0f\x8f\x75\x6a\x9b\xe4\x41\x93\x80\xc1\xe2\x8c\x16\x5c\x7b\xe0\x98\x8f\xc5\xba\x5f\x38\xae\x3a\xc0\xac\xb4\xc8\x0f\x9f\x85\x8d\xf5\xa4\x1f\x91\xed\x4d\x67\xfa\x6d\xb9\x16\x9d\x5c\xf0\x2d\xe4\x2b\x03\xc0\x98\x8d\xa8\xad\x0e\x1c\x21\x9c\xc9\x3b\x74\xa0\x24\xa7\x08\xd1\x16\xe5\xbf\xdd\x1c\x9c\x90\x74\x0a\x5f\x28\x12\x81\x7b\x2e\x80\x11\x2e\xb7\x2a\x41\x26\xae\x05\x37\xc5\x6f\x1f\x41\x5a\x1c\xd5\x50\xdb\x8f\x96\x7e\x38\xa1\xb2\xa7\x31\xf0\x1f\xf9\x63\x60\x9f\x56\xb6\xf5\x92\xeb\x31\xc2\x6d\x26\x84\xfa\xb1\xd3\x73\x33\xd1\xb5\x4c\xea\xf0\x4b\x54\x52\x68\x28\x94\x3b\x77\x78\x51\xaa\x46\x84\xad\xa5\x42\x73\xa6\x13\x1a\xb7\x81\x72\xdd\xa2\xb2\xe9\x0b\xb7\x98\xda\xf3\x07\xd9\x52\x66\xc6\xfc\xa9\x0a\x9c\x4e\x46\x59\xfe\x4e\xd7\xef\xef\xef\x35\xb0\x34\x83\x1f\xcd\x70\x26\x7a\x10\x96\x75\x70\x5d\x82\x39\xe1\xba\x54\x86\x8b\xbf\x1b\x13\xb7\x53\x74\x53\x25\xdf\x28\x8d\x5b\xa4\x63\xc7\xf2\x0f\x91\xa2\x28\x8f\xeb\xa9\x1d\x6d\x3d\x10\x84\x6d\x82\x7e\xb2\x88\x1d\xe6\x47\xb0\xdf\x41\xfc\xa5\xf6\xe8\x95\x6c\x80\xcd\xeb\x79\xa5\xce\xd6\xf2\x50\xe1\x49\x90\x72\x52\x69\xe8\xe8\xac\x28\x3e\x3d\x4c\x04\x5e\xae\xc3\x36\x07\xa5\xd6\x8b\x67\x2d\xe5\xa4\x1b\xa9\x8c\x84\x83\x4a\x0e\x81\x07\xd8\xdf\xff\xa6\xd3\x12\x2b\x2d\xf9\xe1\x75\xd8\xfb\xba\xb6\xe9\x3a\x54\x4e\x55\x26\x7f\x08\x53\x55\x49\x38\x49\x56\x53\x6c\x16\xaa\x5b\x43\x51\x48\x5f\x24\xc5\xa5\xde\x8d\x3a\x48\xd0\x22\xc9\xe7\xe6\x3e\xf9\xc5\x55\x74\xb9\x54\x6e\x24\x15\x29\x8b\x04\x16\x2e\x46\x2e\x54\xb2\xe6\x25\x5c\xb0\xe0\x13\x2e\x55\x34\xe5\xdf\x85\xaf\x13\x0e\x7d\xb8\x9a\x34\xcd\x61\x3d\x39\x25\x5c\x5d\x9a\xf2\xb3\xf0\x75\xe4\xf4\xe1\x6a\xd2\xa4\x35\xe5\xa4\x4b\xc8\xe9\xf2\x7a\x34\x25\x5c\x9e\x26\x37\x60\x0d\x8b\xf4\x2e\xfd\x15\xdf\xe1\x60\x54\x39\xb9\xc3\x0c\x49\xf7\x47\x1d\x14\xa6\xb3\xb3\xd9\xaf\xe0\x72\xfe\xe3\x3a\x14\xf2\x19\x32\xe2\xcc\x28\xdf\xfd\xbf\x72\x3d\x28\x1f\xaf\xdf\x7f\xfc\xd8\x3d\x1b\x50\x18\x9c\x50\x5b\x1f\x78\xb2\xd8\xd5\xe0\x52\xfb\xca\xa3\xa3\x69\xa0\x12\x3d\xf0\x07\xc3\xd7\xac\x52\x82\xb5\xf5\x20\x63\x80\x2d\x52\x26\xef\xff\x0f\x00\x00\xff\xff\xfc\xf7\xae\x97\x27\x6c\x00\x00"

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 27687, mode: os.FileMode(420), modTime: time.Unix(1400000000, 0)}
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
	"index.html": indexHtml,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
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

