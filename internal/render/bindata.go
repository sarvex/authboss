package render

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

var _confirm_email_html_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x29\x2e\x29\xca\xcf\x4b\xb7\xab\xae\xd6\xf3\xc9\xcc\xcb\xae\xad\xb5\xd1\x87\x8a\x00\x02\x00\x00\xff\xff\xe1\x46\x1b\xff\x1a\x00\x00\x00")

func confirm_email_html_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_confirm_email_html_tpl,
		"confirm_email.html.tpl",
	)
}

func confirm_email_html_tpl() (*asset, error) {
	bytes, err := confirm_email_html_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "confirm_email.html.tpl", size: 26, mode: os.FileMode(438), modTime: time.Unix(1424060528, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _confirm_email_txt_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\xd6\xf3\xc9\xcc\xcb\xae\xad\x05\x04\x00\x00\xff\xff\x41\xf7\xa1\x3d\x09\x00\x00\x00")

func confirm_email_txt_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_confirm_email_txt_tpl,
		"confirm_email.txt.tpl",
	)
}

func confirm_email_txt_tpl() (*asset, error) {
	bytes, err := confirm_email_txt_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "confirm_email.txt.tpl", size: 9, mode: os.FileMode(438), modTime: time.Unix(1424060528, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _layout_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x95\xc1\x72\xd3\x30\x10\x86\xef\x79\x0a\x8d\x0e\xdc\x2c\xd1\x26\xc0\x4c\xeb\xe4\x06\x2f\xc0\x13\xc8\xd6\x3a\x56\x90\x25\xa3\x5d\xb7\xc9\x78\xfa\xee\xc8\x96\x1b\x5c\xd7\x40\x87\x03\x33\xe4\x60\x67\x37\xda\xff\xdf\xfd\x34\x52\xf2\x9a\x1a\x7b\xd8\xe4\x35\x28\x7d\xd8\xb0\xf8\xc9\xad\x71\xdf\x58\x1d\xa0\xda\x73\x29\x1d\x90\x76\x4a\x14\xde\x13\x52\x50\x6d\xa9\x9d\x28\x7d\x23\x2b\xef\x28\x53\x8f\x80\xbe\x01\xb9\x13\xef\xc5\x56\x96\x88\x2f\xd2\x22\x26\x38\x0b\x60\xf7\x1c\xe9\x62\x01\x6b\x00\xe2\x4c\xae\xd9\x34\xea\x3c\x28\xbf\xb2\xb9\x26\xe4\x56\x6c\xc5\xcd\xe8\x71\xcd\x89\xc6\xb8\x3f\x98\x60\x19\x4c\x4b\x0c\x43\xb9\xe7\x35\x51\x8b\x77\x52\xaa\x93\x3a\x8b\xa3\xf7\x47\x0b\xaa\x35\x38\xfa\x0c\x39\x69\x4d\x81\xf2\xf4\xbd\x83\x70\x91\xb7\xe2\x26\x8e\x94\x82\xd1\xe7\x84\xfc\x90\xcb\xa4\xb7\x22\xfe\xd6\x11\x6e\xe5\x69\x39\x41\x54\x66\x74\x69\x61\xcf\x09\xce\x24\x4f\xea\x41\x25\xe5\xb9\x61\x2e\xd3\x0e\xe5\x85\xd7\x97\xf8\xd2\xe6\x81\x95\x56\x21\xee\x79\x19\x99\x2b\xe3\x20\x64\x95\xed\x8c\xe6\xa9\xbb\xbe\x37\x15\x13\x5f\xe2\x92\xfa\x6b\x57\x96\x80\xf8\xf4\x94\xda\x9e\x95\x06\xff\x38\x2d\x5f\xfe\x52\x7a\x9b\x9d\x31\xf3\x55\x85\x40\xd9\x96\x0d\x71\xa3\xb3\x8f\xb3\xe5\xcb\x12\x65\x21\x10\x1b\x9f\x19\x26\xcb\x29\xd2\x06\x1b\x83\xa8\x0a\x0b\x9c\x8d\xdb\xb4\xe7\x8d\x0a\x47\xe3\x32\xf2\xed\x1d\xfb\xf4\xa1\x3d\xdf\x2f\x94\x47\xf5\xa2\x23\xf2\x6e\xa2\x93\x02\x7e\xed\xd0\x7a\x8c\x72\x5a\x91\x7a\x36\x98\x7a\x88\xdc\xb0\x55\xee\xf0\x8e\x4c\x03\x78\x1f\x21\x0e\x51\x2e\x93\xc0\x6b\x9b\xbe\x6f\x83\x71\xb4\x4a\xeb\xda\x8a\x8c\x93\xce\x50\xfd\x0c\x67\x5f\xfb\x1e\x9c\x9e\x0a\x67\xfc\x3f\x87\xe0\xc3\xbf\xa4\xaf\x95\x3b\x42\xf8\x1f\xe1\xcf\x51\xfd\x35\xfa\x25\xe3\xb5\xa9\x6f\x17\x53\x2f\xe9\x47\xda\x13\xfd\xdd\x33\xfd\xdd\x6f\xe8\xc7\x19\xc1\xb2\xf1\x99\x69\xa8\x54\x67\x69\x0d\xe9\xb2\x22\x1b\x0e\xb5\x71\xc7\xe1\xa4\xbf\x18\xf2\xd7\x15\xc3\xf9\x5f\x91\x4e\x0c\x08\x9a\xd6\x2a\x02\xc6\x55\x47\x75\xe1\x87\xbb\x51\x2c\x70\xae\x20\x7d\x13\xe5\xeb\x6b\xba\x81\x64\xfa\xe7\xd8\x6c\x7e\x04\x00\x00\xff\xff\x06\xe2\x8e\xc9\x43\x06\x00\x00")

func layout_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_layout_tpl,
		"layout.tpl",
	)
}

func layout_tpl() (*asset, error) {
	bytes, err := layout_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "layout.tpl", size: 1603, mode: os.FileMode(438), modTime: time.Unix(1423465727, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _layoutemail_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x49\xcd\xb5\xab\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\x4a\x2c\x2d\xc9\x48\xca\x2f\x2e\x56\x52\xd0\xab\xad\xb5\xd1\x07\xca\x02\x02\x00\x00\xff\xff\x3a\xdb\x96\xd1\x22\x00\x00\x00")

func layoutemail_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_layoutemail_tpl,
		"layoutEmail.tpl",
	)
}

func layoutemail_tpl() (*asset, error) {
	bytes, err := layoutemail_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "layoutEmail.tpl", size: 34, mode: os.FileMode(438), modTime: time.Unix(1423465727, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _login_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xcc\x94\xcd\x6e\xdc\x20\x10\xc7\xef\x95\xfa\x0e\x88\xbb\xe3\x17\xb0\x2d\xf5\xd0\x9e\xfa\x11\x65\x53\xa9\x57\x0c\xb3\x01\x2d\x30\x08\xc3\x26\x91\xb5\xef\x5e\x58\x7f\x14\x6f\x36\xdb\x6b\x2c\x21\x8d\xc6\xf3\xff\x7b\xe6\x07\xb8\xd9\xa3\x37\x84\xf1\xa0\xd0\xb6\xb4\xd6\xf8\xa4\x2c\x25\x06\x82\x44\xd1\xd2\xfb\x5f\xbb\x47\xda\x7d\xfe\x44\xd2\xd3\x08\x75\x24\x5c\xb3\x61\x68\x69\x16\x55\x4f\x1e\xa3\x1b\x47\xb5\x27\x77\x5f\xbd\x47\x7f\x3a\x11\xc9\x86\x0a\x72\x3c\x8e\x60\xc5\xe9\xb4\x68\x2f\xf5\xca\xba\x18\x26\x83\xb2\xe4\x5c\x36\x38\x66\xaf\xd4\x55\x4c\x08\xb4\xb4\x6b\xd4\xda\x04\x23\x7b\x56\xc5\x01\x7c\xca\xd6\x2a\xad\x2c\xbd\xb4\x3b\x5b\x90\xf0\xea\xa0\xa5\x01\x5e\x02\xdd\xcc\xc0\xd1\x06\x8f\x9a\x12\xcb\x4c\x2a\xc8\x66\x39\xa2\xc4\x69\xc6\x41\xa2\x16\xe0\x5b\xfa\x7b\x4d\x1f\x99\x8e\xa9\x6e\x1c\xef\x96\xdc\xc5\x90\x75\x9a\x72\x01\x56\xc6\x1f\x13\x9e\x46\x7e\xf8\x3f\xbc\x99\x9e\x4b\xc2\x67\xf4\xe2\x26\xc1\x7f\x45\x1b\x82\xf7\x4b\xfa\x1d\x56\x6f\x9a\x97\xa0\x5d\xd5\x4f\xfd\x25\xda\x33\xa3\x4d\x97\xa5\x7e\x22\xb9\x93\xf8\xfc\x00\x06\x4c\x0f\xa9\xf8\x2d\x79\x2e\x81\x1f\x7a\x7c\xd9\x74\xa1\x59\x0f\xfa\xd6\xa9\x59\x55\xf3\x88\xde\xac\xe7\x20\xf8\x08\xb4\x23\xcb\x37\xc9\x0f\x28\xc7\x2b\x9d\xb7\xcd\x9e\xf7\x77\x7e\xd1\xc7\x10\x70\x9d\xbb\x0f\x96\xa4\x55\x39\xaf\x0c\xf3\xaf\xe7\x78\xc2\x30\x77\x33\xc4\xde\xa8\x40\xbb\xef\xf9\xa6\x36\xf5\xa4\xbe\x42\x81\xe3\xb1\x80\xc0\x2e\xfd\xb5\xb2\x87\x77\xcd\x89\xf4\xb0\x4f\x7f\x03\x3f\xb9\xd0\x6e\xb6\x23\x5f\x38\xc7\x68\x43\x53\xb3\xab\xa3\x94\xd4\xa4\x12\x02\xec\xc2\x2c\x6d\xe1\x9f\xdd\xc3\xb7\x9f\xd3\x85\x29\xae\x51\xce\x3e\xe2\x01\x6c\x4e\xd7\xc9\xb5\xa9\xf3\xb1\xea\xfe\x06\x00\x00\xff\xff\xb5\x6b\x5e\x3d\x98\x04\x00\x00")

func login_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_login_tpl,
		"login.tpl",
	)
}

func login_tpl() (*asset, error) {
	bytes, err := login_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "login.tpl", size: 1176, mode: os.FileMode(438), modTime: time.Unix(1424065052, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _recover_complete_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x54\xc1\x8e\xd3\x30\x10\xbd\xef\x57\x8c\xac\xbd\xa6\xb9\xa3\xa4\x17\x04\x37\x60\x45\xf7\xc0\xd5\x8d\xa7\x1b\x6b\x1d\xdb\x4c\x9c\xc2\xca\xf2\xbf\x33\x8e\x9b\x6e\x89\x5a\x81\x84\x38\x6c\x24\x2b\x9e\xb1\xe7\x79\xde\x7b\xd2\x34\x07\x47\x03\xc8\x2e\x68\x67\x5b\x51\x13\x76\xee\x88\x54\x77\x6e\xf0\x06\x03\x0a\x18\x30\xf4\x4e\xb5\xe2\xe1\xcb\xee\x51\x6c\xef\x80\xbf\x46\x5b\x3f\x05\x08\x2f\x1e\x5b\xd1\x6b\xa5\xd0\x0a\xb0\x72\xe0\x28\xb8\xe7\x1c\x1c\xa5\x99\x38\x8a\x71\xf3\x98\x13\x29\x09\xa8\x4b\x6d\x8c\xf7\x5e\x8e\xe3\x0f\x47\xea\x03\xd1\x08\xef\x5a\xd8\xf0\xe6\x93\xf4\x9b\x25\x9f\x52\x79\x45\xe9\x23\x74\x86\x93\xad\xc8\x4d\x56\x4f\xe4\x26\x1f\xa3\x3e\xc0\x6f\x10\x29\x41\x2f\xc7\x0a\x89\x1c\xc5\x88\x96\xeb\x4f\x7d\xae\x51\xe6\xb6\x0b\xcc\xc5\x8d\xf9\xd6\xe8\xa5\xbd\x72\xad\x92\x4a\x39\x2b\xb6\x8d\x3e\x77\x22\xe1\x20\x2b\xe3\xba\x67\xce\xd6\x9a\x57\x2e\x5d\xa1\x15\x7d\x2e\x7b\xef\x9c\x0d\xe4\x8c\x38\x89\x16\xf0\x67\x58\x24\x5b\xb8\x08\xf0\x46\x76\xd8\x3b\xa3\x90\x58\xef\x73\x9a\xf0\xfb\xa4\x09\xd5\x22\xe1\xfc\x44\xcd\xbc\x5e\xc3\x18\x49\xda\x27\x84\x7b\x56\x21\x4b\xba\x12\xe8\x36\xd7\x1e\x8d\xaf\xf6\x85\x4e\x8c\x9e\xb4\x0d\x33\x48\x4a\x6b\x62\x27\x69\xef\x2e\x5e\x5f\x0c\x65\x72\x07\x4d\xc3\xc3\x0d\x5f\x57\xc7\x7f\x61\xef\x15\xc0\xb7\xee\xf2\x8a\xd2\xca\xec\xf7\xe5\x14\xfe\xcd\xf4\xab\xb2\xfd\x27\xef\xe7\xed\x7e\x0a\xc1\x9d\xf1\xf6\xc1\x02\xaf\x8a\x91\x06\x49\x2f\xf3\xbe\xc0\x9f\xf4\x18\xa7\xfd\xa0\x83\xd8\xee\xe6\x7f\x53\x97\xf2\x3f\xce\x14\x9e\x22\xdf\x76\x5f\x3f\x7e\xe6\x20\x0f\x92\xd7\xd9\x92\xb3\x97\xf3\xa5\xa9\xb3\x0b\xdb\x5f\x01\x00\x00\xff\xff\x30\xbb\xc0\xa5\xd3\x04\x00\x00")

func recover_complete_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_recover_complete_tpl,
		"recover-complete.tpl",
	)
}

func recover_complete_tpl() (*asset, error) {
	bytes, err := recover_complete_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "recover-complete.tpl", size: 1235, mode: os.FileMode(438), modTime: time.Unix(1424063388, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _recover_html_email = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb2\x29\x2e\x29\xca\xcf\x4b\xb7\xab\xae\xd6\xf3\xc9\xcc\xcb\xae\xad\xb5\xd1\x87\x8a\x00\x02\x00\x00\xff\xff\xe1\x46\x1b\xff\x1a\x00\x00\x00")

func recover_html_email_bytes() ([]byte, error) {
	return bindata_read(
		_recover_html_email,
		"recover-html.email",
	)
}

func recover_html_email() (*asset, error) {
	bytes, err := recover_html_email_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "recover-html.email", size: 26, mode: os.FileMode(438), modTime: time.Unix(1422773459, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _recover_text_email = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xaa\xae\xd6\xf3\xc9\xcc\xcb\xae\xad\x05\x04\x00\x00\xff\xff\x41\xf7\xa1\x3d\x09\x00\x00\x00")

func recover_text_email_bytes() ([]byte, error) {
	return bindata_read(
		_recover_text_email,
		"recover-text.email",
	)
}

func recover_text_email() (*asset, error) {
	bytes, err := recover_text_email_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "recover-text.email", size: 9, mode: os.FileMode(438), modTime: time.Unix(1422773459, 0)}
	a := &asset{bytes: bytes, info:  info}
	return a, nil
}

var _recover_tpl = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xd4\x94\x4f\x8f\x9b\x30\x10\xc5\xef\xfb\x29\x46\xd6\x5e\x09\xf7\x0a\xb8\x44\xed\xad\x7f\xb4\xbb\x91\x7a\x9d\xc0\x10\xac\x18\x1b\x0d\x26\xea\xca\xe2\xbb\xd7\xc6\x24\x9b\x58\x89\xd4\x1e\x7a\x28\x12\x92\x19\xc6\xbf\xe1\xbd\x67\x51\xb4\x86\x7b\xc0\xda\x4a\xa3\x4b\x91\x33\xd5\xe6\x44\x2c\xa0\x27\xdb\x99\xa6\x14\x3f\xbe\xbf\xbe\x89\xea\x09\xfc\xe5\xdc\xf3\x34\x12\x6b\xec\xe9\x33\xf3\x08\x9f\x4a\xd8\xf8\xc5\x57\x1c\x36\xe7\xfa\x3c\x2f\x9d\x45\x23\x4f\x50\x2b\x1c\xc7\x52\x04\x7e\x76\x60\x33\x0d\xce\xc9\x16\x6e\x10\xf3\x0c\x1d\x8e\x19\x31\x1b\x76\x8e\x74\x33\xcf\xeb\xac\x94\x22\xf5\x30\xd9\x88\xb9\xea\x58\xba\xc6\x01\xf5\x9d\xb6\x0c\x9b\xc6\x68\x51\x15\xf2\xf2\x25\x08\x2d\x66\x61\xbe\xaf\xe6\xd2\xdf\x61\x6b\x42\x5b\x08\x37\xdf\x5e\x1b\x6d\xd9\x28\x01\xf6\x7d\xa0\x52\x58\xfa\x65\x05\x04\x05\xa5\x38\x6b\x11\x30\x28\xac\xa9\x33\xaa\x21\x2e\xc5\xee\x52\x3e\xa1\x9a\x7c\x9f\x73\x9b\xdd\xc5\x21\x01\xf9\x95\xc6\xdc\x8b\xfc\x78\x74\x8e\x51\x1f\x08\x9e\xbd\x25\xc1\xdf\xc4\xad\xc7\xc2\x3b\x52\x43\xb6\x57\xa6\x3e\x8a\xca\xb9\x81\xa5\xb6\x0b\x64\x9e\x53\x95\xab\xcf\x4f\x57\xd3\xcf\xe9\x7a\xa5\xad\xe4\x7e\xf7\x20\xe4\xe4\xf5\x1f\x64\x7d\x07\xf8\xbf\x47\x9e\x48\x4a\x92\xdf\xc6\xb7\x70\xef\x04\x6c\x53\xfb\xfe\xea\x20\xdc\xb5\xf2\x1f\x9d\x87\xb0\x8c\xcf\xfb\xc9\x5a\x73\x81\xee\xad\x06\x7f\x67\x1e\xd7\x23\xbf\x2f\xeb\x38\x63\x35\x6a\x9c\xf6\xbd\xb4\xa2\x7a\x89\x7f\x91\x22\x8f\xfb\x23\xb1\xc0\x94\xa3\xa4\x3e\x3e\x84\x40\xc7\xd4\xfa\x3f\x92\x32\x07\xe9\x33\xdd\xa2\xae\x49\x15\x39\xae\xb0\x18\x5a\xdc\xd1\xc9\xa6\x21\x7d\x4e\xc8\x5b\xfd\xf3\xf5\xe5\xcb\xb7\xd5\xe3\x8f\x00\x42\xf5\xcd\x1c\x49\xaf\xd6\x17\x79\x08\xbb\xfa\x1d\x00\x00\xff\xff\xfd\xfd\xd2\x06\x02\x05\x00\x00")

func recover_tpl_bytes() ([]byte, error) {
	return bindata_read(
		_recover_tpl,
		"recover.tpl",
	)
}

func recover_tpl() (*asset, error) {
	bytes, err := recover_tpl_bytes()
	if err != nil {
		return nil, err
	}

	info := bindata_file_info{name: "recover.tpl", size: 1282, mode: os.FileMode(438), modTime: time.Unix(1424063407, 0)}
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
	"confirm_email.html.tpl": confirm_email_html_tpl,
	"confirm_email.txt.tpl": confirm_email_txt_tpl,
	"layout.tpl": layout_tpl,
	"layoutEmail.tpl": layoutemail_tpl,
	"login.tpl": login_tpl,
	"recover-complete.tpl": recover_complete_tpl,
	"recover-html.email": recover_html_email,
	"recover-text.email": recover_text_email,
	"recover.tpl": recover_tpl,
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
	"confirm_email.html.tpl": &_bintree_t{confirm_email_html_tpl, map[string]*_bintree_t{
	}},
	"confirm_email.txt.tpl": &_bintree_t{confirm_email_txt_tpl, map[string]*_bintree_t{
	}},
	"layout.tpl": &_bintree_t{layout_tpl, map[string]*_bintree_t{
	}},
	"layoutEmail.tpl": &_bintree_t{layoutemail_tpl, map[string]*_bintree_t{
	}},
	"login.tpl": &_bintree_t{login_tpl, map[string]*_bintree_t{
	}},
	"recover-complete.tpl": &_bintree_t{recover_complete_tpl, map[string]*_bintree_t{
	}},
	"recover-html.email": &_bintree_t{recover_html_email, map[string]*_bintree_t{
	}},
	"recover-text.email": &_bintree_t{recover_text_email, map[string]*_bintree_t{
	}},
	"recover.tpl": &_bintree_t{recover_tpl, map[string]*_bintree_t{
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
