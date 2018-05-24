// Code generated by go-bindata.
// sources:
// override/templates_test/singleton/psql_main.tpl
// DO NOT EDIT!

package driver

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

var _templates_testSingletonPsql_mainTpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x5b\x6f\xdb\x36\x14\x7e\xf7\xaf\x38\x13\x90\x42\x2a\x14\x66\xcf\x1e\xfc\xd0\xa4\x49\x56\xac\x6d\xdc\x38\x43\x1f\x86\x61\xa0\xcd\x23\x85\x88\x44\x32\x24\x15\xd7\x2b\xf2\xdf\x07\xde\x22\x29\x91\x9b\x14\x45\x81\x3d\xea\xf0\x3b\xf7\xab\xec\x4e\x21\xa8\xfa\x0a\x8d\x45\x0d\xc6\xea\x6e\x63\xe1\xeb\x0c\x80\xad\x4f\xa4\x10\xf0\xda\xdc\x36\xe4\xed\xf1\xcc\x53\x3e\xd2\x16\xc1\x81\xb8\xa8\x67\x00\xd7\xd2\x58\x80\x01\xa1\x33\xa8\x47\x04\x45\x8d\x19\x11\x8c\x69\x5a\xc9\x70\x80\x90\xda\xcb\xe0\xc2\x3a\x1d\xaa\x5e\x52\x63\xce\x78\xf3\x00\x99\x01\x58\x34\xf6\xed\xb1\x57\x1e\x89\xf7\xb3\x59\xd5\x89\x0d\x70\xc1\x6d\x5e\x44\x7b\x3f\x50\x2e\x60\x01\xaf\x92\x37\x5f\xef\x1d\xee\xe8\x08\x0c\xda\x4e\x01\xeb\x5a\x65\xc0\x5e\x23\x30\x6a\xe9\x9a\x1a\x04\xb3\xb9\xc6\x96\x02\x15\x0c\x78\xeb\x2c\x31\xc0\xad\x33\x45\x02\x05\x8b\x8e\x44\xf5\x0e\x34\x15\x4c\xb6\xcd\xce\xc9\xaa\x51\xa0\xa6\x16\x99\xb7\x6a\x20\x4a\x82\xbd\xa6\xd6\x53\x0d\x6c\xa8\x80\x35\x82\xee\x04\xd0\x9a\x72\x61\xac\x13\xdc\x19\x2e\x6a\x67\xc1\x58\x90\xb9\x6d\xd6\x92\x37\xa8\xe1\xe2\xf2\x03\x28\xba\xb9\xa1\x35\x92\xe0\x60\xae\xe0\x75\xf2\xa7\x08\x8e\xe4\x05\xa0\xd6\x52\x7b\xaf\xef\xa8\x76\x5f\x81\xe2\x03\x48\x62\x9a\x16\x70\xc7\x15\x6a\x72\x8e\x76\xe5\x83\x96\x67\xca\xe5\x92\xad\x05\x6d\x31\x2b\x3c\xd6\x67\x70\x1f\xd2\x3d\x46\x9c\x4f\xec\x3e\x9c\x7b\x8c\x38\x9f\xef\x7d\x38\xf7\x98\x70\x2e\xeb\x03\xdc\x3b\x61\x13\x48\xea\xa4\x34\xd5\xca\x3e\x79\xf1\xdd\xa3\x8f\x8e\xe0\x44\x23\xb5\x08\x34\xa6\x8b\xff\x8b\x0c\xd8\x1a\x9c\xb7\xc4\xcb\x1b\x94\xd1\xa2\x07\x91\x95\xa5\xeb\x06\xc3\x43\x9e\xc2\x57\xb8\x58\xf2\xca\xc7\x76\x01\x8a\xb4\xf4\x06\x97\xe7\xa9\x36\xf3\xe2\x37\xff\xf2\xcb\x02\x04\x6f\x7c\x22\x00\x34\xda\x4e\x0b\x47\x9f\x01\xdc\x3f\xe2\x67\x5a\xaa\x2b\xaf\xff\x25\xbc\x23\xd6\x8d\xf7\xeb\xe5\xcc\xae\x17\xba\x56\x9d\xb4\x0c\xe6\x0b\xc0\x2f\xb8\x21\x27\xb2\x6d\xa9\x60\x79\xa6\xea\x7f\xdc\x5b\x56\x42\x76\x78\x18\xaa\xff\x50\x8a\x66\x97\x95\xd0\x7b\xfe\xc0\x4f\x4e\xc5\x1d\x2c\x80\x2a\x85\x82\xe5\xd2\xb8\x6f\xae\xa5\xc8\x0b\x07\x57\xf5\xa9\xb8\xcb\x0b\x42\x88\x63\x09\x56\x4e\x2b\x35\xb7\x8d\x57\xd0\x27\x60\xc4\xf1\x72\x35\x33\x00\x5d\xc2\xd6\xa9\xe0\x92\x2c\xb9\xc2\x7c\x68\xee\xca\x32\xd9\xb9\xba\xda\x8e\xc4\xaf\x2c\xf3\x83\x41\xe0\xf6\xec\x0f\xdc\xbd\x45\x63\xb5\xdc\xa1\xce\x75\xfd\x65\x79\x7e\x76\x83\xbb\x12\xf4\x38\xe1\xbd\x44\xaa\xed\x33\x41\x97\xda\x90\xcf\x9a\xaa\x1c\xb5\x2e\x21\xab\x28\x6f\xdc\x74\x90\x60\x1c\x33\xc4\x90\xc3\x26\x84\xc3\x57\xeb\x28\xc3\x43\x43\x7f\x5c\x9b\xb9\x6d\x1e\xa9\x9a\xf2\xeb\x33\xe5\xd3\x8a\xaa\xd6\x92\xa5\xe6\xc2\x36\xc2\x69\x28\x5e\xaa\x7d\x4b\xb9\x85\x4a\xea\x3d\xee\xce\x00\xb6\xe4\xa4\x91\x06\xf3\xc2\xf5\xea\x9b\xca\x6d\x9b\x54\xa6\xdc\x00\x93\x02\x4b\xd8\x38\x84\x1f\xd1\x5b\xcd\x2d\x02\x0a\x06\xb2\xf2\x04\xc5\x15\xce\xa6\xc3\xf6\x13\x9d\x99\x8c\x66\x94\x20\x78\xf3\xb0\x87\xc6\x63\x5a\x77\xe2\xa4\x65\xb9\x71\x75\x57\x26\xfe\xb8\xba\x4a\xa0\xba\x36\x40\x08\x09\xdf\xc3\x61\xbe\x99\xe8\x9e\xc8\x1d\xd8\x52\xaf\x7d\x67\xcf\xf0\x0a\x1a\x14\xc1\x9e\xc2\x85\xe8\xd7\x18\xa0\xcd\xa0\x3b\x82\x39\x86\x7c\xc4\xed\x25\x52\x86\x3a\xe2\x93\xd7\x26\x34\xd7\x7c\x01\xaf\xd6\x3b\x8b\x86\x1c\x77\x55\xe5\x97\xac\x7f\x73\xc1\x9f\x7c\xdb\x0c\x1b\x33\x08\xe9\xa9\x21\x95\x81\xbd\x4f\xee\x7c\xe1\xdf\x2f\x3b\xf1\x6c\x5a\x53\xd6\x74\x27\x04\x17\xf5\x3c\x7b\x88\x77\x88\x58\xf1\x84\x23\x98\x40\xe2\x26\x29\x26\x01\xa8\xf5\x23\xc0\xd3\x19\xfb\x6c\x15\xc4\x24\xc0\x5f\x7f\x87\xd0\x7a\xeb\x23\x57\xa2\xf5\x0e\xad\x94\x53\x5f\xe5\xd9\xf2\xfc\xf7\x8b\xd5\xd5\xe2\xc0\xf8\x91\xe9\x76\x70\x51\x4e\xa1\x96\x17\x97\x57\x8b\x03\xe6\x51\x6e\x69\x4e\xa3\xfe\x5c\x9d\x5e\x26\x59\x6e\x4f\xef\x91\xf5\x66\xb5\x3a\x7b\xf7\xfe\x34\x21\xfb\x2b\xcc\xe3\xef\xf7\x78\xf8\x78\x29\x0e\x4a\xd9\xb6\xaa\x4c\xb9\xe4\xb2\xb3\xbc\x21\x57\xd8\x2a\x8f\xcb\xdc\xfa\x51\xf5\xc3\x39\x10\xb3\xfe\xbd\x03\x2f\x0c\x00\x90\xca\x72\x29\xa0\xe2\x0d\xf6\x2d\xea\xfc\x3b\x8b\xfe\x79\x53\xb2\x03\x33\x3f\x60\x73\x25\x8d\xad\x35\x9a\xf9\x20\xbc\x29\x80\x0f\x21\xea\x3b\x26\x1c\x34\xa3\x96\x79\x2a\x39\xc9\xf2\xd0\x38\xdc\x7b\x54\x23\x1c\xac\xf8\x96\x51\x07\x7b\xcd\x09\x1b\xf9\x7f\x68\x58\xbf\xc9\x7f\xba\x71\xc3\x6a\x84\x85\x2b\x2c\xe2\x2f\xb5\xa2\xef\x26\x47\x8b\xcb\x65\x4f\xa5\x8e\x6f\xa8\x41\x9d\x46\x09\x8a\xc4\x99\xed\x6b\x33\xa0\xd9\xfa\xc9\xd1\x32\x2d\x7c\x78\xdb\x3d\x27\xda\x61\xbd\xe0\xec\xf0\x90\x57\x87\xf8\x85\x1b\x6b\xa6\xf4\x1c\x1d\x81\x45\xaa\x99\xdc\x0a\xbf\x10\x3a\x8b\x06\x36\x0d\x52\xd1\x29\xb0\xd4\xdc\x18\xd8\x5e\xa3\xf0\xbb\x31\xfc\x73\x54\x5c\x70\x73\x9d\x26\xe1\x94\xa1\x49\xe0\x37\xfe\x20\xc6\x47\xab\xff\x05\x4c\xa1\x7d\xc9\xd9\x9a\x78\xc0\xa3\x7e\xe8\x06\xee\x23\x28\x0d\xb9\xc4\x56\xde\xb9\xeb\x7c\x30\x99\xf6\xa5\x5a\x0a\xe7\x60\x1e\x7f\x5c\xcb\xe0\x58\xf8\x47\xe4\x55\x6f\xe1\x94\xee\xf4\x58\x7a\xeb\xa3\x15\x8f\x03\xd4\x83\xe2\xf6\xba\x6d\xc8\x85\x42\x91\x67\x69\xba\x64\x25\x30\xcd\xef\x50\x93\xe5\xea\xd3\xfb\xe3\x8e\x37\xec\x53\x87\x7a\x17\x57\x4a\x68\x98\x54\xfc\x4f\xdb\xe9\x71\xb3\xc5\xff\x9c\xe2\x99\x59\x29\x78\x53\x4e\x44\x6f\xec\xd3\xfd\xec\xbf\x00\x00\x00\xff\xff\x03\x2b\x53\xb3\xf1\x0f\x00\x00")

func templates_testSingletonPsql_mainTplBytes() ([]byte, error) {
	return bindataRead(
		_templates_testSingletonPsql_mainTpl,
		"templates_test/singleton/psql_main.tpl",
	)
}

func templates_testSingletonPsql_mainTpl() (*asset, error) {
	bytes, err := templates_testSingletonPsql_mainTplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates_test/singleton/psql_main.tpl", size: 4081, mode: os.FileMode(420), modTime: time.Unix(1526660155, 0)}
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
	"templates_test/singleton/psql_main.tpl": templates_testSingletonPsql_mainTpl,
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
	"templates_test": &bintree{nil, map[string]*bintree{
		"singleton": &bintree{nil, map[string]*bintree{
			"psql_main.tpl": &bintree{templates_testSingletonPsql_mainTpl, map[string]*bintree{}},
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
