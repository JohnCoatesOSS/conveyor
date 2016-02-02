// Code generated by go-bindata.
// sources:
// db/migrations/1_initial_schema.sql
// db/migrations/2_add_sequences.sql
// DO NOT EDIT!

package conveyor

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

var _dbMigrations1_initial_schemaSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x52\x4d\x8f\xd3\x30\x10\xbd\xe7\x57\x3c\xed\x25\xa9\xa0\x20\x24\x6e\x15\x87\x42\xbd\x10\xa9\x4a\xa1\x4d\xb4\x7b\x8b\xdc\x64\x9a\x58\x24\x76\x70\xec\x96\xf0\xeb\x71\xbe\xca\xaa\x1c\x7a\xd8\xdb\xcc\x9b\xcf\xf7\x66\x96\x4b\xbc\xa9\x45\xa1\xb9\x21\x24\x8d\xf7\x65\xcf\xd6\x31\x03\x7b\x8e\x59\x74\x08\x77\x11\xc2\x47\x44\xbb\xd8\x01\xe1\x21\x3e\xa0\x6c\x8d\xd2\xb4\xba\x97\xf6\x60\xad\xc8\x97\xaa\x6d\x9b\x87\x95\x37\x27\xc7\xeb\xcf\x5b\x86\xa3\x15\x55\xde\x22\xf0\x00\x91\xa3\xcf\x1b\x0a\xa3\x64\xbb\xc5\x86\x3d\xae\x93\x6d\x3c\xa0\x69\x41\x92\xfa\xad\xd2\xf3\xc7\x60\x81\x46\x8b\x9a\xeb\x0e\x3f\xa9\x7b\xeb\x4a\x35\x35\xaa\x15\x6e\x97\x0e\x86\x7e\x9b\x1e\x3a\x6a\x2e\xb3\xf2\xea\xb6\x25\xff\x67\x9b\x9e\xdd\xec\x65\x9a\x9c\x9b\xa7\xdc\xc0\x88\x9a\x5c\xb0\x6e\x70\x11\xa6\x54\x76\x44\xf0\x47\x49\x42\x4e\x27\x6e\x2b\x83\x40\xaa\x8b\x5b\x80\xbf\x8c\xf9\xd6\x64\xfe\x62\x6a\xad\xef\x37\x1b\xc6\xaa\xba\xa9\xe8\x7e\xae\xb7\xb8\x95\xcc\x4d\x10\x27\x9e\x99\x57\xab\x36\x68\x9f\xfe\xd7\x40\xd3\x89\x34\xc9\x8c\xda\xe9\x3a\x81\xc8\x07\x72\xae\xb8\x18\x75\x1b\x96\x72\xaf\xf2\x44\x4e\x58\x65\xab\x1c\x24\x5b\xab\x5d\xb0\x74\x6c\x2e\x04\x25\xab\x0e\x25\x3f\x13\x3e\xa0\x21\x99\x0b\x59\xbc\x1f\x9a\x39\x63\xec\x8a\x93\xd2\xe0\xb2\x43\x21\xce\x24\xfb\xfb\xbc\x9b\x69\x26\x51\xf8\x23\x61\x08\xa3\x0d\x7b\x86\x95\xe2\x97\xa5\x74\xac\x71\x9f\x35\x3d\x4c\x72\x08\xa3\xaf\x38\x1a\x4d\x84\xc0\x15\x2f\xf0\xf4\x8d\xed\x99\xb3\x87\xe3\x7e\x82\x3f\x8f\xf3\xb1\xdb\xe3\x8a\x4e\xcb\xf8\x13\x81\xeb\xaf\x6f\xd4\x45\x7a\x9b\xfd\xee\xfb\xad\xc8\xab\x97\xe8\x38\x7c\xe5\xfd\x0d\x00\x00\xff\xff\x0c\xc4\x52\xb1\x26\x03\x00\x00")

func dbMigrations1_initial_schemaSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations1_initial_schemaSql,
		"db/migrations/1_initial_schema.sql",
	)
}

func dbMigrations1_initial_schemaSql() (*asset, error) {
	bytes, err := dbMigrations1_initial_schemaSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/1_initial_schema.sql", size: 806, mode: os.FileMode(420), modTime: time.Unix(1454046367, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _dbMigrations2_add_sequencesSql = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd2\xd5\x55\xd0\xce\xcd\x4c\x2f\x4a\x2c\x49\x55\x08\x2d\xe0\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2a\xcd\xcc\x49\x29\x56\x70\x74\x71\x51\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x4e\x2d\x54\x08\x76\x0d\xf2\x74\xf4\xb1\x46\x51\x99\x58\x54\x92\x99\x96\x98\x5c\x82\x53\x31\x97\x73\x90\xab\x63\x88\xab\x82\xa7\x9f\x8b\x6b\x84\x42\x66\x5e\x4a\x6a\x45\x3c\xc4\xf8\xf8\xfc\xbc\x78\x90\x4a\x7f\x3f\x98\x7d\xa1\xc1\x9e\x7e\xee\x0a\x49\x25\x45\xa9\xa9\x0a\x1a\x40\x29\x4d\x6b\x6c\xba\xe1\x56\x22\x19\x80\x70\x06\x16\x33\xb8\x90\x7d\xea\x92\x5f\x9e\x87\xcd\xaf\x2e\x41\xfe\x01\x48\xee\xc7\xe5\x4b\x0c\x65\x80\x00\x00\x00\xff\xff\x47\xe8\x1d\xc6\x46\x01\x00\x00")

func dbMigrations2_add_sequencesSqlBytes() ([]byte, error) {
	return bindataRead(
		_dbMigrations2_add_sequencesSql,
		"db/migrations/2_add_sequences.sql",
	)
}

func dbMigrations2_add_sequencesSql() (*asset, error) {
	bytes, err := dbMigrations2_add_sequencesSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "db/migrations/2_add_sequences.sql", size: 326, mode: os.FileMode(420), modTime: time.Unix(1454374721, 0)}
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
	"db/migrations/1_initial_schema.sql": dbMigrations1_initial_schemaSql,
	"db/migrations/2_add_sequences.sql": dbMigrations2_add_sequencesSql,
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
	"db": &bintree{nil, map[string]*bintree{
		"migrations": &bintree{nil, map[string]*bintree{
			"1_initial_schema.sql": &bintree{dbMigrations1_initial_schemaSql, map[string]*bintree{}},
			"2_add_sequences.sql": &bintree{dbMigrations2_add_sequencesSql, map[string]*bintree{}},
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

