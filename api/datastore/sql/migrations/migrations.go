// Code generated by go-bindata.
// sources:
// 10_add_call_log_app_app_id.down.sql
// 10_add_call_log_app_app_id.up.sql
// 11_add_route_app_app_id.down.sql
// 11_add_route_app_app_id.up.sql
// 12_remove_route_app_name.down.sql
// 12_remove_route_app_name.up.sql
// 13_remove_call_app_name.down.sql
// 13_remove_call_app_name.up.sql
// 1_add_route_created_at.down.sql
// 1_add_route_created_at.up.sql
// 2_add_call_stats.down.sql
// 2_add_call_stats.up.sql
// 3_add_call_error.down.sql
// 3_add_call_error.up.sql
// 4_add_route_updated_at.down.sql
// 4_add_route_updated_at.up.sql
// 5_add_app_created_at.down.sql
// 5_add_app_created_at.up.sql
// 6_add_app_updated_at.down.sql
// 6_add_app_updated_at.up.sql
// 7_add_route_cpus.down.sql
// 7_add_route_cpus.up.sql
// 8_add_app_id.down.sql
// 8_add_app_id.up.sql
// 9_add_call_app_app_id.down.sql
// 9_add_call_app_app_id.up.sql
// DO NOT EDIT!

package migrations

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

var __10_add_call_log_app_app_idDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\xc9\x4f\x2f\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2c\x28\x88\xcf\x4c\xb1\xe6\x02\x04\x00\x00\xff\xff\xde\x3b\x4b\x8d\x25\x00\x00\x00")

func _10_add_call_log_app_app_idDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__10_add_call_log_app_app_idDownSql,
		"10_add_call_log_app_app_id.down.sql",
	)
}

func _10_add_call_log_app_app_idDownSql() (*asset, error) {
	bytes, err := _10_add_call_log_app_app_idDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10_add_call_log_app_app_id.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1515792190, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __10_add_call_log_app_app_idUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\xc8\xc9\x4f\x2f\x56\x70\x74\x71\x51\x48\x2c\x28\x88\xcf\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x32\x35\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\x5b\xa1\xcf\xde\x2a\x00\x00\x00")

func _10_add_call_log_app_app_idUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__10_add_call_log_app_app_idUpSql,
		"10_add_call_log_app_app_id.up.sql",
	)
}

func _10_add_call_log_app_app_idUpSql() (*asset, error) {
	bytes, err := _10_add_call_log_app_app_idUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "10_add_call_log_app_app_id.up.sql", size: 42, mode: os.FileMode(420), modTime: time.Unix(1515792190, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __11_add_route_app_app_idDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xca\x2f\x2d\x49\x2d\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2c\x28\x88\xcf\x4c\xb1\xe6\x02\x04\x00\x00\xff\xff\x8d\x6c\x53\x17\x27\x00\x00\x00")

func _11_add_route_app_app_idDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__11_add_route_app_app_idDownSql,
		"11_add_route_app_app_id.down.sql",
	)
}

func _11_add_route_app_app_idDownSql() (*asset, error) {
	bytes, err := _11_add_route_app_app_idDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "11_add_route_app_app_id.down.sql", size: 39, mode: os.FileMode(420), modTime: time.Unix(1515792190, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __11_add_route_app_app_idUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xca\x2f\x2d\x49\x2d\x56\x70\x74\x71\x51\x48\x2c\x28\x88\xcf\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x32\x35\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\x03\xb3\x3d\x6e\x2c\x00\x00\x00")

func _11_add_route_app_app_idUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__11_add_route_app_app_idUpSql,
		"11_add_route_app_app_id.up.sql",
	)
}

func _11_add_route_app_app_idUpSql() (*asset, error) {
	bytes, err := _11_add_route_app_app_idUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "11_add_route_app_app_id.up.sql", size: 44, mode: os.FileMode(420), modTime: time.Unix(1515792237, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __12_remove_route_app_nameDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xca\x2f\x2d\x49\x2d\x56\x70\x74\x71\x51\x48\x2c\x28\x88\xcf\x4b\xcc\x4d\x55\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x32\x35\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\xd7\x6d\x2f\x01\x2e\x00\x00\x00")

func _12_remove_route_app_nameDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__12_remove_route_app_nameDownSql,
		"12_remove_route_app_name.down.sql",
	)
}

func _12_remove_route_app_nameDownSql() (*asset, error) {
	bytes, err := _12_remove_route_app_nameDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "12_remove_route_app_name.down.sql", size: 46, mode: os.FileMode(420), modTime: time.Unix(1515792467, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __12_remove_route_app_nameUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xca\x2f\x2d\x49\x2d\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2c\x28\x88\xcf\x4b\xcc\x4d\xb5\xe6\x02\x04\x00\x00\xff\xff\x54\xab\x9f\x2f\x29\x00\x00\x00")

func _12_remove_route_app_nameUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__12_remove_route_app_nameUpSql,
		"12_remove_route_app_name.up.sql",
	)
}

func _12_remove_route_app_nameUpSql() (*asset, error) {
	bytes, err := _12_remove_route_app_nameUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "12_remove_route_app_name.up.sql", size: 41, mode: os.FileMode(420), modTime: time.Unix(1515792467, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __13_remove_call_app_nameDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4e\xcc\xc9\x29\x56\x70\x74\x71\x51\x48\x2c\x28\x88\xcf\x4b\xcc\x4d\x55\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x32\x35\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\x58\x78\xd5\x5e\x2d\x00\x00\x00")

func _13_remove_call_app_nameDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__13_remove_call_app_nameDownSql,
		"13_remove_call_app_name.down.sql",
	)
}

func _13_remove_call_app_nameDownSql() (*asset, error) {
	bytes, err := _13_remove_call_app_nameDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "13_remove_call_app_name.down.sql", size: 45, mode: os.FileMode(420), modTime: time.Unix(1515792618, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __13_remove_call_app_nameUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4e\xcc\xc9\x29\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2c\x28\x88\xcf\x4b\xcc\x4d\xb5\xe6\x02\x04\x00\x00\xff\xff\x92\x88\xe9\x6c\x28\x00\x00\x00")

func _13_remove_call_app_nameUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__13_remove_call_app_nameUpSql,
		"13_remove_call_app_name.up.sql",
	)
}

func _13_remove_call_app_nameUpSql() (*asset, error) {
	bytes, err := _13_remove_call_app_nameUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "13_remove_call_app_name.up.sql", size: 40, mode: os.FileMode(420), modTime: time.Unix(1515792618, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_add_route_created_atDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xca\x2f\x2d\x49\x2d\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2e\x4a\x4d\x2c\x49\x4d\x89\x4f\x2c\xb1\xe6\x02\x04\x00\x00\xff\xff\x47\xfd\x3b\xbe\x2b\x00\x00\x00")

func _1_add_route_created_atDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_add_route_created_atDownSql,
		"1_add_route_created_at.down.sql",
	)
}

func _1_add_route_created_atDownSql() (*asset, error) {
	bytes, err := _1_add_route_created_atDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_add_route_created_at.down.sql", size: 43, mode: os.FileMode(420), modTime: time.Unix(1511093370, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __1_add_route_created_atUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xca\x2f\x2d\x49\x2d\x56\x70\x74\x71\x51\x48\x2e\x4a\x4d\x2c\x49\x4d\x89\x4f\x2c\x51\x28\x49\xad\x28\xb1\xe6\x02\x04\x00\x00\xff\xff\x3b\x59\x9c\x54\x28\x00\x00\x00")

func _1_add_route_created_atUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__1_add_route_created_atUpSql,
		"1_add_route_created_at.up.sql",
	)
}

func _1_add_route_created_atUpSql() (*asset, error) {
	bytes, err := _1_add_route_created_atUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "1_add_route_created_at.up.sql", size: 40, mode: os.FileMode(420), modTime: time.Unix(1511093370, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_add_call_statsDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4e\xcc\xc9\x29\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x2e\x49\x2c\x29\xb6\xe6\x02\x04\x00\x00\xff\xff\xd3\x09\xeb\x22\x25\x00\x00\x00")

func _2_add_call_statsDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_add_call_statsDownSql,
		"2_add_call_stats.down.sql",
	)
}

func _2_add_call_statsDownSql() (*asset, error) {
	bytes, err := _2_add_call_statsDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_add_call_stats.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1511974963, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __2_add_call_statsUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4e\xcc\xc9\x29\x56\x70\x74\x71\x51\x28\x2e\x49\x2c\x29\x56\x28\x49\xad\x28\xb1\xe6\x02\x04\x00\x00\xff\xff\x29\xde\x11\xe8\x22\x00\x00\x00")

func _2_add_call_statsUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__2_add_call_statsUpSql,
		"2_add_call_stats.up.sql",
	)
}

func _2_add_call_statsUpSql() (*asset, error) {
	bytes, err := _2_add_call_statsUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "2_add_call_stats.up.sql", size: 34, mode: os.FileMode(420), modTime: time.Unix(1511974963, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_add_call_errorDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4e\xcc\xc9\x29\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2d\x2a\xca\x2f\xb2\xe6\x02\x04\x00\x00\xff\xff\xc1\x14\x26\x51\x25\x00\x00\x00")

func _3_add_call_errorDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__3_add_call_errorDownSql,
		"3_add_call_error.down.sql",
	)
}

func _3_add_call_errorDownSql() (*asset, error) {
	bytes, err := _3_add_call_errorDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "3_add_call_error.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1511983925, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __3_add_call_errorUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4e\xcc\xc9\x29\x56\x70\x74\x71\x51\x48\x2d\x2a\xca\x2f\x52\x28\x49\xad\x28\xb1\xe6\x02\x04\x00\x00\xff\xff\xaf\xba\x27\xcd\x22\x00\x00\x00")

func _3_add_call_errorUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__3_add_call_errorUpSql,
		"3_add_call_error.up.sql",
	)
}

func _3_add_call_errorUpSql() (*asset, error) {
	bytes, err := _3_add_call_errorUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "3_add_call_error.up.sql", size: 34, mode: os.FileMode(420), modTime: time.Unix(1511983925, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_add_route_updated_atDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xca\x2f\x2d\x49\x2d\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x2d\x48\x49\x2c\x49\x4d\x89\x4f\x2c\xb1\xe6\x02\x04\x00\x00\xff\xff\xa4\x67\xb0\xea\x2b\x00\x00\x00")

func _4_add_route_updated_atDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__4_add_route_updated_atDownSql,
		"4_add_route_updated_at.down.sql",
	)
}

func _4_add_route_updated_atDownSql() (*asset, error) {
	bytes, err := _4_add_route_updated_atDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4_add_route_updated_at.down.sql", size: 43, mode: os.FileMode(420), modTime: time.Unix(1514396019, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __4_add_route_updated_atUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xca\x2f\x2d\x49\x2d\x56\x70\x74\x71\x51\x28\x2d\x48\x49\x2c\x49\x4d\x89\x4f\x2c\x51\x28\x4b\x2c\x4a\xce\x48\x2c\xd2\x30\x32\x35\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\x54\xf7\xac\x11\x30\x00\x00\x00")

func _4_add_route_updated_atUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__4_add_route_updated_atUpSql,
		"4_add_route_updated_at.up.sql",
	)
}

func _4_add_route_updated_atUpSql() (*asset, error) {
	bytes, err := _4_add_route_updated_atUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "4_add_route_updated_at.up.sql", size: 48, mode: os.FileMode(420), modTime: time.Unix(1514396019, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_add_app_created_atDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2c\x28\x28\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2e\x4a\x4d\x2c\x49\x4d\x89\x4f\x2c\xb1\xe6\x02\x04\x00\x00\xff\xff\xd2\xde\x5c\x98\x29\x00\x00\x00")

func _5_add_app_created_atDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__5_add_app_created_atDownSql,
		"5_add_app_created_at.down.sql",
	)
}

func _5_add_app_created_atDownSql() (*asset, error) {
	bytes, err := _5_add_app_created_atDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "5_add_app_created_at.down.sql", size: 41, mode: os.FileMode(420), modTime: time.Unix(1514396019, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __5_add_app_created_atUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2c\x28\x28\x56\x70\x74\x71\x51\x48\x2e\x4a\x4d\x2c\x49\x4d\x89\x4f\x2c\x51\x28\x4b\x2c\x4a\xce\x48\x2c\xd2\x30\x32\x35\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\x76\x6c\x0f\x45\x2e\x00\x00\x00")

func _5_add_app_created_atUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__5_add_app_created_atUpSql,
		"5_add_app_created_at.up.sql",
	)
}

func _5_add_app_created_atUpSql() (*asset, error) {
	bytes, err := _5_add_app_created_atUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "5_add_app_created_at.up.sql", size: 46, mode: os.FileMode(420), modTime: time.Unix(1514396019, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __6_add_app_updated_atDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2c\x28\x28\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x28\x2d\x48\x49\x2c\x49\x4d\x89\x4f\x2c\xb1\xe6\x02\x04\x00\x00\xff\xff\x31\x44\xd7\xcc\x29\x00\x00\x00")

func _6_add_app_updated_atDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__6_add_app_updated_atDownSql,
		"6_add_app_updated_at.down.sql",
	)
}

func _6_add_app_updated_atDownSql() (*asset, error) {
	bytes, err := _6_add_app_updated_atDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "6_add_app_updated_at.down.sql", size: 41, mode: os.FileMode(420), modTime: time.Unix(1514396019, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __6_add_app_updated_atUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2c\x28\x28\x56\x70\x74\x71\x51\x28\x2d\x48\x49\x2c\x49\x4d\x89\x4f\x2c\x51\x28\x4b\x2c\x4a\xce\x48\x2c\xd2\x30\x32\x35\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\x65\x01\x8b\x34\x2e\x00\x00\x00")

func _6_add_app_updated_atUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__6_add_app_updated_atUpSql,
		"6_add_app_updated_at.up.sql",
	)
}

func _6_add_app_updated_atUpSql() (*asset, error) {
	bytes, err := _6_add_app_updated_atUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "6_add_app_updated_at.up.sql", size: 46, mode: os.FileMode(420), modTime: time.Unix(1514396019, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __7_add_route_cpusDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xca\x2f\x2d\x49\x2d\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2e\x28\x2d\xb6\xe6\x02\x04\x00\x00\xff\xff\xec\x60\x24\xd0\x25\x00\x00\x00")

func _7_add_route_cpusDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__7_add_route_cpusDownSql,
		"7_add_route_cpus.down.sql",
	)
}

func _7_add_route_cpusDownSql() (*asset, error) {
	bytes, err := _7_add_route_cpusDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "7_add_route_cpus.down.sql", size: 37, mode: os.FileMode(420), modTime: time.Unix(1515792189, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __7_add_route_cpusUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x28\xca\x2f\x2d\x49\x2d\x56\x70\x74\x71\x51\x48\x2e\x28\x2d\x56\xc8\xcc\x2b\xb1\xe6\x02\x04\x00\x00\xff\xff\xf1\x18\xf8\xa9\x21\x00\x00\x00")

func _7_add_route_cpusUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__7_add_route_cpusUpSql,
		"7_add_route_cpus.up.sql",
	)
}

func _7_add_route_cpusUpSql() (*asset, error) {
	bytes, err := _7_add_route_cpusUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "7_add_route_cpus.up.sql", size: 33, mode: os.FileMode(420), modTime: time.Unix(1515792189, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __8_add_app_idDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2c\x28\x28\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\xc8\x4c\xb1\xe6\x02\x04\x00\x00\xff\xff\x6f\x16\xc9\xf6\x21\x00\x00\x00")

func _8_add_app_idDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__8_add_app_idDownSql,
		"8_add_app_id.down.sql",
	)
}

func _8_add_app_idDownSql() (*asset, error) {
	bytes, err := _8_add_app_idDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "8_add_app_id.down.sql", size: 33, mode: os.FileMode(420), modTime: time.Unix(1515792190, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __8_add_app_idUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x2c\x28\x28\x56\x70\x74\x71\x51\xc8\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x32\x35\xd3\xb4\xe6\x02\x04\x00\x00\xff\xff\x7d\x6d\x8e\x1a\x26\x00\x00\x00")

func _8_add_app_idUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__8_add_app_idUpSql,
		"8_add_app_id.up.sql",
	)
}

func _8_add_app_idUpSql() (*asset, error) {
	bytes, err := _8_add_app_idUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "8_add_app_id.up.sql", size: 38, mode: os.FileMode(420), modTime: time.Unix(1515792190, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __9_add_call_app_app_idDownSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4e\xcc\xc9\x29\x56\x70\x09\xf2\x0f\x50\x70\xf6\xf7\x09\xf5\xf5\x53\x48\x2c\x28\x88\xcf\x4c\xb1\xe6\x02\x04\x00\x00\xff\xff\x5e\x6a\x7c\x07\x26\x00\x00\x00")

func _9_add_call_app_app_idDownSqlBytes() ([]byte, error) {
	return bindataRead(
		__9_add_call_app_app_idDownSql,
		"9_add_call_app_app_id.down.sql",
	)
}

func _9_add_call_app_app_idDownSql() (*asset, error) {
	bytes, err := _9_add_call_app_app_idDownSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "9_add_call_app_app_id.down.sql", size: 38, mode: os.FileMode(420), modTime: time.Unix(1515792190, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var __9_add_call_app_app_idUpSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x72\xf4\x09\x71\x0d\x52\x08\x71\x74\xf2\x71\x55\x48\x4e\xcc\xc9\x29\x56\x70\x74\x71\x51\x48\x2c\x28\x88\xcf\x4c\x51\x08\x73\x0c\x72\xf6\x70\x0c\xd2\x30\x32\x35\xd3\xb4\xb6\xe6\x02\x04\x00\x00\xff\xff\x11\xea\x70\x92\x2c\x00\x00\x00")

func _9_add_call_app_app_idUpSqlBytes() ([]byte, error) {
	return bindataRead(
		__9_add_call_app_app_idUpSql,
		"9_add_call_app_app_id.up.sql",
	)
}

func _9_add_call_app_app_idUpSql() (*asset, error) {
	bytes, err := _9_add_call_app_app_idUpSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "9_add_call_app_app_id.up.sql", size: 44, mode: os.FileMode(420), modTime: time.Unix(1515792190, 0)}
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
	"10_add_call_log_app_app_id.down.sql": _10_add_call_log_app_app_idDownSql,
	"10_add_call_log_app_app_id.up.sql": _10_add_call_log_app_app_idUpSql,
	"11_add_route_app_app_id.down.sql": _11_add_route_app_app_idDownSql,
	"11_add_route_app_app_id.up.sql": _11_add_route_app_app_idUpSql,
	"12_remove_route_app_name.down.sql": _12_remove_route_app_nameDownSql,
	"12_remove_route_app_name.up.sql": _12_remove_route_app_nameUpSql,
	"13_remove_call_app_name.down.sql": _13_remove_call_app_nameDownSql,
	"13_remove_call_app_name.up.sql": _13_remove_call_app_nameUpSql,
	"1_add_route_created_at.down.sql": _1_add_route_created_atDownSql,
	"1_add_route_created_at.up.sql": _1_add_route_created_atUpSql,
	"2_add_call_stats.down.sql": _2_add_call_statsDownSql,
	"2_add_call_stats.up.sql": _2_add_call_statsUpSql,
	"3_add_call_error.down.sql": _3_add_call_errorDownSql,
	"3_add_call_error.up.sql": _3_add_call_errorUpSql,
	"4_add_route_updated_at.down.sql": _4_add_route_updated_atDownSql,
	"4_add_route_updated_at.up.sql": _4_add_route_updated_atUpSql,
	"5_add_app_created_at.down.sql": _5_add_app_created_atDownSql,
	"5_add_app_created_at.up.sql": _5_add_app_created_atUpSql,
	"6_add_app_updated_at.down.sql": _6_add_app_updated_atDownSql,
	"6_add_app_updated_at.up.sql": _6_add_app_updated_atUpSql,
	"7_add_route_cpus.down.sql": _7_add_route_cpusDownSql,
	"7_add_route_cpus.up.sql": _7_add_route_cpusUpSql,
	"8_add_app_id.down.sql": _8_add_app_idDownSql,
	"8_add_app_id.up.sql": _8_add_app_idUpSql,
	"9_add_call_app_app_id.down.sql": _9_add_call_app_app_idDownSql,
	"9_add_call_app_app_id.up.sql": _9_add_call_app_app_idUpSql,
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
	"10_add_call_log_app_app_id.down.sql": &bintree{_10_add_call_log_app_app_idDownSql, map[string]*bintree{}},
	"10_add_call_log_app_app_id.up.sql": &bintree{_10_add_call_log_app_app_idUpSql, map[string]*bintree{}},
	"11_add_route_app_app_id.down.sql": &bintree{_11_add_route_app_app_idDownSql, map[string]*bintree{}},
	"11_add_route_app_app_id.up.sql": &bintree{_11_add_route_app_app_idUpSql, map[string]*bintree{}},
	"12_remove_route_app_name.down.sql": &bintree{_12_remove_route_app_nameDownSql, map[string]*bintree{}},
	"12_remove_route_app_name.up.sql": &bintree{_12_remove_route_app_nameUpSql, map[string]*bintree{}},
	"13_remove_call_app_name.down.sql": &bintree{_13_remove_call_app_nameDownSql, map[string]*bintree{}},
	"13_remove_call_app_name.up.sql": &bintree{_13_remove_call_app_nameUpSql, map[string]*bintree{}},
	"1_add_route_created_at.down.sql": &bintree{_1_add_route_created_atDownSql, map[string]*bintree{}},
	"1_add_route_created_at.up.sql": &bintree{_1_add_route_created_atUpSql, map[string]*bintree{}},
	"2_add_call_stats.down.sql": &bintree{_2_add_call_statsDownSql, map[string]*bintree{}},
	"2_add_call_stats.up.sql": &bintree{_2_add_call_statsUpSql, map[string]*bintree{}},
	"3_add_call_error.down.sql": &bintree{_3_add_call_errorDownSql, map[string]*bintree{}},
	"3_add_call_error.up.sql": &bintree{_3_add_call_errorUpSql, map[string]*bintree{}},
	"4_add_route_updated_at.down.sql": &bintree{_4_add_route_updated_atDownSql, map[string]*bintree{}},
	"4_add_route_updated_at.up.sql": &bintree{_4_add_route_updated_atUpSql, map[string]*bintree{}},
	"5_add_app_created_at.down.sql": &bintree{_5_add_app_created_atDownSql, map[string]*bintree{}},
	"5_add_app_created_at.up.sql": &bintree{_5_add_app_created_atUpSql, map[string]*bintree{}},
	"6_add_app_updated_at.down.sql": &bintree{_6_add_app_updated_atDownSql, map[string]*bintree{}},
	"6_add_app_updated_at.up.sql": &bintree{_6_add_app_updated_atUpSql, map[string]*bintree{}},
	"7_add_route_cpus.down.sql": &bintree{_7_add_route_cpusDownSql, map[string]*bintree{}},
	"7_add_route_cpus.up.sql": &bintree{_7_add_route_cpusUpSql, map[string]*bintree{}},
	"8_add_app_id.down.sql": &bintree{_8_add_app_idDownSql, map[string]*bintree{}},
	"8_add_app_id.up.sql": &bintree{_8_add_app_idUpSql, map[string]*bintree{}},
	"9_add_call_app_app_id.down.sql": &bintree{_9_add_call_app_app_idDownSql, map[string]*bintree{}},
	"9_add_call_app_app_id.up.sql": &bintree{_9_add_call_app_app_idUpSql, map[string]*bintree{}},
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

