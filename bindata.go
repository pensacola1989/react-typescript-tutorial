// Code generated by go-bindata.
// sources:
// views/header.html
// views/index.html
// views/js.html
// DO NOT EDIT!

package main

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

var _viewsHeaderHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\x4e\x49\x4d\xcb\xcc\x4b\x55\x50\xca\x48\x4d\x4c\x49\x2d\x52\xaa\xad\xe5\xe2\xaa\xae\x4e\xcd\x4b\xa9\xad\x05\x04\x00\x00\xff\xff\xbc\xc5\xbc\xd3\x1c\x00\x00\x00")

func viewsHeaderHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsHeaderHtml,
		"views/header.html",
	)
}

func viewsHeaderHtml() (*asset, error) {
	bytes, err := viewsHeaderHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/header.html", size: 28, mode: os.FileMode(420), modTime: time.Unix(1472429417, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsIndexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x4c\x90\xbb\x4e\xeb\x40\x10\x86\xfb\xf3\x14\x73\xb6\x49\x83\xb1\x53\x40\xe1\xd8\x6e\x20\x35\x08\xa5\xa1\x1c\x76\x47\x78\xd0\xde\x64\x0f\x4e\xac\x28\xef\xce\x6e\x36\x48\x54\xfb\xef\xed\xff\x3e\x4d\xf7\xff\xf9\xe5\xe9\xf0\xfe\xba\x87\x51\x9c\x1d\xfe\x75\x65\x01\xe8\x46\x42\x93\x43\x8a\x8e\x04\xc1\xa3\xa3\x5e\x2d\x4c\xc7\x18\x26\x51\xa0\x83\x17\xf2\xd2\xab\x23\x1b\x19\x7b\x43\x0b\x6b\xaa\xae\x9b\x3b\x60\xcf\xc2\x68\xab\x59\xa3\xa5\x7e\x7b\xdf\xa8\xba\x74\x9d\xcf\x42\x2e\x5a\x14\x02\x95\x09\x34\xa9\xcb\xa5\x50\x84\xc5\xd2\xf0\x46\xa8\x05\x2a\x38\xac\x91\x66\x3d\x71\x14\xd8\x9f\x30\x7d\xa1\xae\x2e\x2f\xb2\x5c\x5d\xec\x72\xfc\x08\x66\xbd\x79\x1a\x5e\x80\x4d\xbf\xb9\xa9\x6d\x60\x96\x35\xe1\x95\xc3\xe9\x93\x7d\x8b\xdf\x12\x76\x70\x35\x6c\x1f\x9b\x26\x9e\x76\x50\x6e\x2a\x09\xb1\xdd\x3e\xa4\x13\x35\x74\x75\x6a\x29\x8c\xdf\xe6\xbf\xce\x5f\x73\xf6\x4d\xfc\x3c\xa6\x9f\x00\x00\x00\xff\xff\x6d\x31\x09\x83\x3d\x01\x00\x00")

func viewsIndexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsIndexHtml,
		"views/index.html",
	)
}

func viewsIndexHtml() (*asset, error) {
	bytes, err := viewsIndexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/index.html", size: 317, mode: os.FileMode(420), modTime: time.Unix(1472429453, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _viewsJsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x9c\xcf\xc1\x0a\xc2\x30\x0c\xc6\xf1\xbb\x4f\x51\x77\x4f\xa3\x78\x9d\x7b\x02\xc1\x67\xd8\xda\x08\x19\x6d\x57\x93\x0d\x91\xd2\x77\x57\x11\x2f\x22\x0c\xbc\x06\x7e\x7f\xf2\x95\xe2\xe9\xc2\x89\x4c\x33\x6a\x53\xeb\xc6\x98\x76\x0b\x60\x4e\x3c\x48\x2f\x4c\x6a\x00\xba\xd7\x51\x9d\x70\x9e\x8d\x8a\x3b\x36\x98\x97\x21\xb0\xc3\x51\x31\xf0\xa0\x38\x5e\x17\x92\x3b\x1c\xec\xde\xee\x6c\xe4\x64\x9f\xa9\xae\xc5\x37\x59\xd5\x42\xbd\x9b\xff\x64\xe0\xa7\xf8\x83\x7e\x56\x9c\x6f\x69\xe5\x7f\x37\xc5\xcc\x81\x3c\x6a\x1f\x73\xa0\xaf\x4e\x29\x94\x7c\xad\x8f\x00\x00\x00\xff\xff\xdb\xf4\xc0\xf4\x24\x01\x00\x00")

func viewsJsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_viewsJsHtml,
		"views/js.html",
	)
}

func viewsJsHtml() (*asset, error) {
	bytes, err := viewsJsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "views/js.html", size: 292, mode: os.FileMode(420), modTime: time.Unix(1472429724, 0)}
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
	"views/header.html": viewsHeaderHtml,
	"views/index.html":  viewsIndexHtml,
	"views/js.html":     viewsJsHtml,
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
	"views": &bintree{nil, map[string]*bintree{
		"header.html": &bintree{viewsHeaderHtml, map[string]*bintree{}},
		"index.html":  &bintree{viewsIndexHtml, map[string]*bintree{}},
		"js.html":     &bintree{viewsJsHtml, map[string]*bintree{}},
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
