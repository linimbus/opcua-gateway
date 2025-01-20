package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"strings"
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

var _main_ico = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x1c\x0c\xe3\xf3\x00\x00\x01\x00\x01\x00\x40\x40\x00\x00\x00\x00\x08\x00\x06\x0c\x00\x00\x16\x00\x00\x00\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x40\x00\x00\x00\x40\x08\x06\x00\x00\x00\xaa\x69\x71\xde\x00\x00\x00\x01\x73\x52\x47\x42\x00\xae\xce\x1c\xe9\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8f\x0b\xfc\x61\x05\x00\x00\x00\x09\x70\x48\x59\x73\x00\x00\x16\x25\x00\x00\x16\x25\x01\x49\x52\x24\xf0\x00\x00\x0b\x9b\x49\x44\x41\x54\x78\x5e\xed\x5b\xf9\x6f\x14\xe7\x19\xde\x7f\x20\x52\xa5\xa8\x8a\xd2\x5f\x52\x29\xaa\x9a\x22\x4a\x53\x02\x26\x9c\xc6\x35\xf7\xe5\x10\x8e\x1a\x0c\x14\x48\x49\xc2\x15\x4a\x68\x48\x95\x28\x69\xb9\x44\x45\x50\x5b\x52\x35\x25\x49\x05\x4d\x1a\x82\xb8\x92\x18\x05\x52\xaa\x70\x63\x1b\x5f\xeb\x5d\xaf\xd7\xc7\xde\xb6\x67\xbe\x9d\xbd\x0f\xa0\x50\x25\xcd\x5b\x3d\xef\xec\xb7\xec\xce\xac\xc1\xa6\x4d\xb1\xd6\x59\xe9\x91\xbf\x9d\x79\x66\xf6\x7b\xde\xef\x9a\xf7\x99\xcf\x16\x8b\xc5\x62\x49\xa7\xd3\xe3\x35\x4d\x9b\xa2\x28\x4a\xf9\x50\x00\xb4\x42\x33\xb4\x5b\x54\x45\x39\xf6\xd5\x97\x5f\xd2\x3f\x6f\xdc\xa0\x1b\xd7\xaf\x0f\x09\x40\x2b\x34\xab\x8a\x72\xdc\x12\x8f\x46\x29\x28\x04\x09\x55\x1d\x52\x80\xe6\x58\x34\x4a\x96\xa1\x28\x5e\x02\xda\x2d\x85\x0e\x6a\xc1\x60\xf6\x3b\xca\x32\x48\xb9\xe7\xee\xc6\x0b\x8a\xbb\xf3\x06\x03\xf2\x02\x80\x8a\x05\xfc\x3e\xb2\xdb\x5a\xa8\x3b\xe0\xc7\x18\x21\x47\xab\x9d\x5c\x5d\x5d\x5c\xf1\xce\x8e\x0e\x6a\x73\xb4\x32\xcf\xef\xf3\x32\xaf\xa7\xbb\x9b\x79\xad\x76\x1b\xb9\x5d\x2e\xe6\x75\xb4\x3b\xc9\xd9\xe6\x60\x9e\xd7\xe3\xc9\xf2\x94\xde\x5e\xe6\x79\xdc\xee\x41\x13\x84\xbc\x00\xa0\xf2\x5e\x8f\x9b\x1a\xea\xeb\x58\x20\x84\x35\x36\x36\x90\xb3\xad\x8d\x42\x9a\x46\xad\xad\x76\xb2\x36\x37\x71\xe5\x3d\x6e\x17\x35\xd4\x5f\xa5\x80\xdf\xcf\xc2\x1a\x1b\xea\x59\x38\x78\x76\xbb\x8d\x5a\xac\xcd\x7c\x3f\x04\x0f\xbc\xee\x40\x80\x7a\x7b\x7a\x98\x87\x40\xe6\xf6\x8a\xbb\x41\xcd\xd4\x2d\x11\x8f\xe9\x88\xdd\x9e\xb7\x50\xc7\x48\x38\x44\xc9\x78\x8c\x11\x8b\x84\xf3\xaf\x8f\xc4\x48\x44\xe3\x3a\x82\x1a\x09\x55\xc9\x3b\x6f\x1a\x02\x40\x6e\xeb\xe4\x95\x55\x74\xed\x3e\xce\x19\xca\xc6\xef\x85\xca\xfd\x05\x37\x8c\xd7\x47\x17\x2e\xd5\xd0\x85\x4b\xb5\x74\xe9\x72\x1d\x75\x07\xba\xf9\x5e\x21\x2d\x48\x36\xbb\x83\xce\x5d\xb8\x42\xe7\x2f\xd6\x50\x53\x73\x8b\x1e\x5c\xfc\x8e\xa2\x90\x68\xaa\x25\x51\x7f\x91\xc4\xd5\x0b\x24\xdc\x9d\x99\x20\xdc\xbe\xb7\x69\x08\xf4\xf6\x74\x73\xeb\x2a\xbd\x3d\x7c\x0c\xdd\x15\xad\xac\x77\x7b\x1f\x77\x69\x1c\x47\x97\xd6\x79\xbd\xdc\x0a\xe0\x61\xd8\xe8\x3c\x2f\xf9\xbc\x1e\x2e\x17\xe6\x05\x06\x14\x08\xb4\x2c\x04\x8e\x2a\x9d\x4f\x25\x3f\x59\x40\x13\xa7\x57\x92\xb5\xc5\x4e\xb1\x48\x84\xae\xa5\x92\xb4\x6d\xf7\x3e\xfa\xe1\xd8\xd9\xf4\xa3\xf1\x73\x69\xdd\x96\xd7\xf5\x1e\xa2\x85\x48\x04\x7c\x24\x36\x2f\x20\xb1\x72\x22\x89\xaa\x31\x24\x4e\x1f\x25\x11\x4f\xe5\xdd\xdb\x34\x04\x50\xd9\x9a\x2b\x97\x58\x00\x2a\x7c\xb5\xae\x86\x1c\xad\xad\xdc\xb5\x5b\x6c\x56\xee\xce\xa8\xbc\xdb\xd5\x45\x35\x57\x2e\x73\x50\x20\xae\xb6\xe6\x0a\x8f\x7b\xf0\x30\x4c\xd0\xd5\xe5\xbc\x01\x1e\x82\x88\x21\x00\x5e\xbb\xd3\x39\xa0\x21\x20\x03\x30\xba\xec\x69\x1a\x33\x65\x21\x95\xce\x5c\x42\x2d\xb6\xd6\x6c\x00\xb6\xff\x76\x1f\x8d\x18\x37\x87\x1e\x9f\x30\x8f\xd6\xff\xf2\xd7\xf9\x01\x78\x71\x21\x89\x55\xa5\x24\x96\x3d\x49\xe2\xf4\xb1\x3b\x07\x40\xa8\x82\x5b\x1e\x2d\x08\x51\x38\x06\x81\x68\x45\x88\x46\xcb\x61\x92\xc4\x71\x88\x91\xf3\x84\x99\xe7\x67\xc1\xb8\x5f\x2e\x0f\x40\x19\x3c\x9c\x33\x0a\xed\x0b\x08\x00\xba\xf7\xb8\xa9\x8b\x69\xc2\x8c\x4a\x2a\x9f\xbb\x3c\x2f\x00\xbb\xde\xf8\x23\x3d\x51\xfa\x14\x07\x68\xd3\xcb\xdb\xf3\x03\xf0\x52\x25\x89\x35\xe5\x7a\x2f\xf8\xec\xf8\xdd\x02\xa0\x43\x8a\x92\x65\xf9\x3d\xb7\xdc\x5f\x5e\xa1\x6b\x72\xbf\xf7\x07\x08\x6a\x20\x10\xa0\xa6\x66\x1b\x35\x5b\x6d\x64\xb5\xda\x39\xb0\x38\xce\xbd\xac\xb3\x8b\x1a\x9b\x5a\x78\xfc\xb7\xb5\xb5\x67\x86\x57\x66\x0e\x68\xb3\x91\x68\x6d\x26\x61\x6f\x22\xe1\xf3\x92\xc8\x2c\xcf\x12\xa6\x21\x80\x55\x00\xdd\x5e\xb6\x1a\xba\x7c\x9b\x43\xef\xda\x98\xdd\x9b\x1a\x1b\xf8\x07\x30\x54\xae\xd6\xd5\x66\x57\x81\xfa\xab\x75\xd4\xee\xd4\x57\x0b\x5b\x8b\x95\x9a\x9b\x1a\xf9\x7e\x5d\x9d\x9d\xcc\x93\xab\x00\x78\x1d\xed\xed\x03\x1a\x02\x7a\xdd\xf0\xe4\x16\xd1\x11\x89\x64\xe7\x10\x55\x51\x29\x1c\xd2\x28\x9e\x39\x87\x15\x21\xef\xda\x70\x84\x44\x38\x4a\x22\x12\x25\x51\xe0\x37\x4d\x93\x20\x04\x61\xad\xef\xe9\x0e\x70\x00\x30\xae\xe5\xfa\x8e\x25\xad\xc3\xe9\xcc\xf0\x7c\x19\x9e\xfe\x1c\x00\x1e\x82\xa7\x8b\xee\xe0\x25\x51\x4e\x88\xe0\x41\x3c\x02\xe5\x74\x80\xa7\x4f\x90\xc6\xca\xdc\x0f\x98\x86\x80\xec\x56\xb2\x82\x79\x65\x11\xcc\xb6\x9c\xe4\xc9\xeb\x8c\xd7\xf4\x87\x37\x18\xf0\x4d\x2e\x80\x8c\x48\xb6\xca\x50\x02\x34\x73\x36\xa8\xaa\xea\xdb\xb7\x6e\xde\xa4\x74\x2a\x45\xa9\x64\x92\x12\xf1\x38\x26\x93\xaf\x8a\x11\xd0\x06\x8d\xe9\x64\x92\xa0\x59\x08\xf1\x0e\x9b\x22\xb1\x58\x6c\xac\xa6\x69\xe5\xc1\x60\xb0\xec\xcd\xb7\x0e\xec\xdc\xf0\xd2\x8e\x7f\xaf\xdd\xfc\x5a\x51\x01\x9a\xf6\xfd\xf9\xe0\x2e\x68\xd4\x14\xa5\x1c\x9a\x59\xbc\xf1\xf3\xc8\xf0\x49\xf3\xc7\x4e\xab\xe4\x87\x8a\x62\x02\x34\x3d\x32\xa2\xec\x69\xa3\x5e\xd3\xe7\xc1\xef\x8e\x9a\x57\x52\xbe\x88\x46\x4e\xaa\x28\x2a\x40\xd3\x83\x8f\x96\x54\x18\xf5\x5a\x3a\x3a\x3a\x1e\x48\x26\x93\xcf\x6b\x42\x6c\x11\x42\x6c\x3e\x79\xea\x1f\x1f\xec\x7d\xf3\x5d\xda\xf3\xfb\xfd\x45\x05\x68\xaa\xfe\xf4\xcc\x21\x68\x84\x56\x68\x8e\x44\x22\x0f\x60\x12\x3c\x4f\x44\xf4\xaf\x5b\xb7\x78\x62\xb8\x79\xe3\x06\x3f\x5f\x17\x23\xa0\x0d\x1a\xa1\x15\x1f\x55\x51\x2e\x58\xe2\xb1\x98\x69\x7d\x1c\x2a\xe0\x65\xd0\x78\x70\xa8\xa1\x60\x00\xf0\xa0\x90\x5b\x96\xdf\x73\xcb\x85\x78\x32\xc5\xbd\x3b\xcf\xfc\x9b\xf7\x0b\xa6\x6c\x10\xc9\x8b\xb5\xb9\x91\x93\x1d\x24\x39\xc8\xec\xa4\x87\x07\x23\xc3\x6e\xb3\xb1\x08\x18\x26\x30\x3e\x90\xe5\x21\xc9\x81\x07\xe8\xea\xea\x64\x1e\x12\x23\x98\x9f\x28\xc3\x01\x02\x4f\x9a\xa2\xcc\x73\xe9\x26\xab\xb1\x32\xf7\x03\xa6\x00\x40\x58\x53\x63\x7d\x36\x00\xa8\x3c\xd2\x57\xa4\xb9\x6d\x6d\x0e\x0e\x08\x02\x80\xcc\x0f\xa9\xb1\x0c\x00\xd2\x5f\x64\x81\xe0\xc1\x49\x86\x13\x8c\xfb\x21\x93\x04\x0f\x01\x40\x46\x08\x9e\x0c\x94\xb1\x32\x7d\x43\xd1\xbd\xbc\x78\x52\x47\x2c\xa9\x7b\x7e\x19\x7f\x01\xe9\x71\x3a\x95\xa0\x74\x32\x41\xf1\x58\x94\x53\x64\x79\x6d\x32\x11\xe7\xe3\x40\x38\x14\x32\x79\x11\xa6\x21\xa0\x3f\x27\x6b\xd9\xae\xca\x79\x42\xb0\xaf\x2c\x2f\x97\x97\x7f\x8d\x91\x27\xef\x7f\x4f\xd9\x20\xae\xf7\x74\x91\xf8\xbc\x9a\xc4\xd9\x93\x24\xce\x7d\xaa\xbb\x3d\xc1\x20\x8b\xaa\xbb\xda\x48\x27\x3e\x39\x4d\x1f\x55\x9f\x66\xe3\x14\x46\x29\x7e\x03\x0d\xf3\xd9\x99\xb3\xf4\xf1\xc9\xbf\xf3\x79\x87\xc3\xc9\xde\x41\xee\xbd\x4d\x7e\x00\x7c\x80\xce\x8e\x76\x36\x47\x71\x0c\x86\x86\xcf\xeb\xe5\x73\x1e\x8f\x9b\x3d\x01\x1c\x47\xcb\x63\x68\xa0\x55\x11\x55\xb4\x3e\x86\x0f\xf3\xdc\x6e\xf6\x0c\x51\x86\xbf\x00\x9e\x34\x45\x75\x9e\x6f\x60\x41\x88\x26\x48\xd4\x9e\xd3\x8d\xcd\xe5\xe3\x48\xac\x2e\x25\xe1\x68\x61\xb3\xe3\x7a\x3a\x45\xaf\xef\xfc\x1d\x7d\x7f\xd4\x34\xfa\x41\xc9\x0c\x7a\x6e\xd3\xab\x94\x8c\x45\x59\xa8\xdf\xe7\xa7\xe9\x4f\xad\xa4\x11\xe3\xe7\xd0\x63\xa3\xa7\xd3\xa1\xa3\x1f\xd3\xf5\x74\x32\xef\xde\xa6\x21\x00\xa7\xa7\xb6\xe6\x72\x8e\x29\x5a\x9b\x35\x45\x6d\xb6\x96\x3c\x53\x14\x06\xa7\x34\x45\xeb\x6a\x6b\xb2\xa6\x28\xc6\x39\x86\x11\xee\x87\x00\x82\x27\x4d\x51\xf0\x06\xec\x08\xc1\xd3\xaf\x3b\x4f\xe2\x67\x13\x48\xac\x9a\x4c\xe2\xb9\xa9\x24\xda\xf4\x00\x60\x7d\xcf\x75\x85\xd7\x67\x5c\x61\x19\x80\x39\x8b\xd6\xd0\xa8\xb2\xf9\x6c\x9a\x1e\x3e\x56\x4d\xd7\x52\x89\xbe\x03\x00\xa8\x4a\x6f\xd6\x0d\xc2\x77\x39\x79\xa1\x0c\x01\xba\xa1\xa9\xf2\x31\xe9\x06\xdd\x2b\xaf\xdf\x88\xc6\x48\x34\x5c\x22\xb1\x7e\x0e\x89\x8d\xf3\x48\xbc\xb8\x80\x84\xd3\xce\x01\xc0\xd8\xde\xf3\x87\xfd\x34\x69\xe6\x12\x2a\x9d\x5d\x45\x5b\x5f\xdb\x4d\x78\xe1\x2b\x03\xb0\x74\xf5\x2f\x68\x4a\xc5\x0a\x2a\x9d\xb5\x94\x4e\x54\x9f\xa2\x74\x32\x9e\x77\xef\x02\x01\x50\xfa\xac\x30\x84\xf5\x66\x84\xe9\x81\xba\x03\xaf\x47\x7f\xaf\xf0\x3f\x09\x00\x96\x57\xbc\xa7\xc0\x3c\xe0\x71\xe9\xc8\xdc\x0f\xbd\xd1\xef\xf7\x53\x57\x97\x8b\x81\x17\x28\x41\x71\xfb\x5a\xb7\xdb\x43\x2e\x97\x9b\xcf\xe1\xb7\x8d\x43\x6f\xc0\x43\xa0\xb1\xdf\x43\xa0\x21\x6f\x08\x48\x53\xf4\x9e\x86\x00\x80\x8a\xc3\xea\x96\xc8\xb1\xd5\x43\xc1\x20\x45\x42\x21\x06\x7e\x3f\xf7\x3a\xf4\x04\x79\x0e\xc6\xaa\xf1\xbe\xe6\x49\x30\x10\x60\x43\x13\x2d\xcd\x93\x56\x47\x07\x07\x43\x0f\x8e\x9b\x5c\x9d\x9d\xcc\x83\xf7\x0f\x21\x72\x12\xc4\xc4\xe9\xf7\x7a\x75\x9e\xcb\xc5\x4b\x9d\x9c\x04\xc1\x43\x90\x00\xe6\x0d\x74\x12\xfc\x1a\x61\x1a\x02\x72\xd9\x92\x15\xcc\x5d\xb6\xcc\xcb\x5b\x61\xb3\xb3\xbf\xbc\xc1\x00\xcb\x80\xbb\x62\x11\x01\xda\x2d\x91\x70\x98\x00\x7c\xc1\xf8\x91\xad\x57\xac\x90\x1a\xa1\x39\x1c\x0e\xb3\x29\xba\x23\x99\x48\x7c\x81\x13\xc1\x9c\x44\xa6\x68\x81\x49\x53\xd3\x28\x95\x48\x7c\x21\x14\x65\x17\xbb\x42\xa1\x50\xe8\xd1\x74\x34\x3a\x4c\x55\xd5\xc7\x5e\xdd\xb6\x77\xeb\x8a\xe7\xb7\x52\xd5\xcf\x37\x17\x15\xa0\xe9\x95\xdf\xbc\xf1\x32\x34\x42\x2b\x34\x1b\xdd\x31\xfe\x0c\x7b\x62\xea\xc2\x09\x33\x96\xd2\x98\xf2\x85\x45\x05\x68\x1a\x56\x32\x63\x91\x51\xaf\xe9\x33\xe4\x4c\xd1\x70\x38\xfc\x70\x2a\x95\xda\x1e\x0a\x06\xf7\x0a\x21\xf6\x9c\xbf\x58\x73\xea\x2f\xef\x1d\xa1\x77\x0e\x7e\x58\x54\x80\x26\x68\x83\x46\x68\x85\xe6\x48\x24\xf2\x1d\xec\x14\xb5\x4b\x53\x14\x90\xa6\x28\xb2\xa6\x62\x82\x34\x45\xa5\x4e\x7c\x84\xaa\xda\xf9\xdd\xa0\xd1\x24\x18\x0a\x60\x23\xe5\x1b\x53\xb4\xc0\xa3\x30\x80\x7d\x79\xc6\x63\x85\x70\x2f\xbc\xfe\x5e\xf3\xff\x82\x29\x1b\x44\xe2\xd3\xd8\x80\x0d\x90\xba\x27\x08\x0f\x0f\x66\x28\x7b\x7d\x0e\x07\x67\x7a\x78\xa0\x80\x27\x88\x9d\x60\x72\x8b\x0c\xb2\x3f\x24\x3a\xe0\x61\x43\x25\xbc\x43\xdc\x0f\x9e\x20\x78\x32\x1b\x04\x0f\x19\x22\xce\x19\x2b\x73\x3f\x60\x0a\x00\x6c\x2d\x88\x94\x5b\x65\x61\x6e\xca\x0a\xb7\xb7\x3b\xc9\x61\xb7\x73\x00\x10\xa8\x16\xab\x35\x6b\x8a\x42\xb0\x34\x3b\xb1\x57\x08\xc6\x28\xca\xd8\x0e\x03\x9e\xf4\x01\xc0\x93\x5b\x6e\x8c\x95\xe9\x0b\xa8\x07\x02\xcb\x13\x5a\xe6\x2d\x0f\x52\x5b\xee\x4d\x98\xbf\xb0\xff\x27\x99\xd6\x01\xf7\x28\xeb\x15\xa8\x6c\x98\xe8\x13\x61\x8a\x53\x62\xe3\x7c\x67\x1a\x02\x10\x07\xa3\x11\x7f\xf1\x5d\x3e\x3b\x73\x39\xf3\x2c\x5d\x88\x17\xd6\x90\x6f\x67\x78\x9a\x96\xe5\x69\xa6\xfb\xdd\xe6\xf5\x17\xc8\xe9\xdb\xdb\x3b\xe9\xc0\xfb\x47\xe8\xc0\xdf\x8e\xd2\x7b\x87\x8e\x93\xc7\xe3\xe5\xfa\x88\x50\x58\xf7\x0b\x8f\x1f\x20\x71\xe2\x20\x89\xf3\xa7\x48\x64\xea\xdc\xdb\xd3\x4b\x87\x8f\x56\xf3\x75\xef\xfe\xf5\x30\x35\x59\x6d\xa6\x4d\x54\x26\x3f\x00\x2d\x9f\xbb\x49\x0a\xad\x09\x1f\x00\x37\xe4\x4d\x52\x99\xcd\x4f\xfa\x26\x29\x47\xd6\xed\xc9\xdf\x24\xd5\xc9\xc3\x01\x3c\xe4\xfe\xe0\x65\x37\x49\x31\x6f\x60\x9b\xa4\x52\x89\x38\x7d\x7e\xee\x32\x0d\x1f\x33\x8b\x0d\xce\xd1\x93\xe7\x53\xb3\xd5\x4e\xf1\x48\x44\x6f\xf5\xb7\xb6\x91\x58\x38\x82\xc4\xe2\xc7\x49\xec\x5c\x47\x22\x16\xe7\xa0\xfb\xbc\x7e\x9a\x32\x6f\x05\x0d\x1f\x3b\x8b\xbe\x37\x72\x0a\x7d\x70\xe4\x23\xee\x0d\x7d\x06\x40\xef\xb2\x83\x6f\x9b\x1c\x36\x4a\x9e\xbd\x70\x85\x46\x4e\xac\x60\x83\x73\xfc\xb4\xc5\xd9\xad\xb2\x22\x91\x22\xf1\xf6\x4e\x12\x95\xa3\x48\x54\x95\x90\xd8\xfd\x02\x7b\x88\x08\x00\x3c\xc1\x59\x0b\x9e\xa1\x91\xa5\x15\x34\xfc\xc9\x59\x74\xf8\xd8\x27\x74\xfd\x6e\xa6\x28\x5e\x42\xc0\xef\x93\xdf\x11\x04\x39\x6e\xcc\xe5\xff\x8e\xd7\x5f\xc0\xe5\xad\xa9\x6d\xa0\xf9\x55\x6b\x69\xc1\xf2\xf5\xb4\x64\xd5\x26\x6a\x75\x38\x29\x1a\x0e\x93\x88\x27\x48\x7c\xf8\x27\x7d\x4f\xf0\x96\x45\x24\xf6\xef\xe0\x39\x01\x0d\x11\xf0\x07\x68\xcd\xc6\x57\xf8\xba\x39\x8b\xd7\xd0\xc9\x53\x67\xb8\x37\xe5\xde\xdb\x14\x00\xb4\xe6\x9d\xb7\xca\x62\x0b\x2c\x78\xd8\x02\xab\xaf\x14\x77\xe2\xa1\xd5\x25\x0f\x40\x59\x1a\xa6\x03\x85\xb4\xd5\x4c\xa6\x2a\xea\x80\x20\x33\xf2\x83\xab\xf4\x2a\xd9\x6b\x0a\x05\xde\x34\x04\x06\xe3\x66\x69\x00\x33\x3a\x4f\xc2\x19\x63\x23\xef\x3c\xe6\x13\xbc\x3d\xc2\x71\xc3\x56\x58\x70\xe5\x75\x85\xe6\x1d\xd3\x24\x88\x37\x42\x10\x27\x5b\x09\x01\x91\x26\xa6\xcf\xe7\xe5\x09\x11\xc7\xd1\xda\xe0\xc9\xc8\x62\x69\x93\xdb\xea\x11\x3c\xcc\x25\xb2\x37\x18\x79\x03\xdd\x2e\xff\x75\xc2\x34\x04\x50\x31\x44\x4d\x56\x30\xaf\x5c\x8c\x3b\x45\xe5\x7a\x3d\x14\x01\xed\xbc\x55\x16\x1b\x08\xa5\x39\xca\x66\xa1\xa6\xf1\xc3\x47\x51\x41\xd3\xf8\x21\x48\x6a\x84\x66\x68\x87\x29\xba\x31\x95\x4c\x46\x82\x42\xa4\x55\x45\x49\xf5\xf4\xf4\xdc\xf4\x79\x7d\xfc\x8a\xa9\x98\x00\x4d\xd0\x06\x8d\xd0\x0a\xcd\xaa\xaa\xbe\xc0\xae\x90\xdf\xef\xff\xd6\x35\x4d\x7b\x48\x08\xf1\xed\x67\xd6\x6d\x5d\x3b\x7f\xd9\x06\x9a\x57\xf9\x6c\x51\x01\x9a\x56\x6f\xf8\xd5\x3a\x68\xd4\x34\xed\x21\x68\x36\xba\x63\xfc\x19\x39\x7e\xee\xe2\x49\xb3\x97\xf1\xbf\xa7\x14\x13\xa0\xe9\xc7\x13\x2b\x7e\x6a\xd4\x6b\xfa\x94\xcd\x5c\xb2\x6a\xc6\xc2\x67\xa9\xbc\x62\x65\x51\x01\x9a\xca\x66\x57\xad\x36\xea\x35\x7d\x26\x4e\xaf\x7a\x78\x72\xc5\xf2\xe9\x93\x66\x56\x4e\x2d\x26\x40\x13\xb4\x19\xf5\xfe\x07\x89\x94\x98\x50\x5a\x4e\x50\x42\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\xe5\xd4\x8b\x5d\x1c\x0c\x00\x00")

func main_ico() ([]byte, error) {
	return bindata_read(
		_main_ico,
		"main.ico",
	)
}

var _start_ico = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x69\x06\x96\xf9\x00\x00\x01\x00\x01\x00\x40\x40\x00\x00\x00\x00\x08\x00\x53\x06\x00\x00\x16\x00\x00\x00\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x40\x00\x00\x00\x40\x08\x06\x00\x00\x00\xaa\x69\x71\xde\x00\x00\x00\x01\x73\x52\x47\x42\x00\xae\xce\x1c\xe9\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8f\x0b\xfc\x61\x05\x00\x00\x00\x09\x70\x48\x59\x73\x00\x00\x16\x25\x00\x00\x16\x25\x01\x49\x52\x24\xf0\x00\x00\x05\xe8\x49\x44\x41\x54\x78\x5e\xe5\x5a\x49\x88\x24\x55\x10\x7d\xf1\x2b\xab\x15\xa9\x16\x77\x1c\x15\x1b\x41\x54\x1a\x17\xc4\x15\x44\x05\x41\x19\x0f\x32\x1e\xec\x71\x19\xbd\x28\x78\x10\x66\xc0\x83\x2b\x48\x23\x73\x10\x15\xbd\xb8\xdc\x14\x11\x11\x07\x44\x46\x04\x0f\x0a\x82\x0a\x1e\x1c\x11\x54\x46\x0f\x0e\x08\x03\xa2\xa2\x4e\xa3\xb8\x8d\xcb\xb4\x44\x56\x44\xf9\xeb\x1b\xf9\x23\x6b\xcb\x2a\xaa\x1f\x24\xdd\xf9\xfa\x45\xc4\x7f\x99\xd9\xf1\x7f\x2e\x00\x80\xd5\xd5\xd5\xc0\x1b\xff\xbe\xb2\xb2\xd2\xda\x48\x9c\x2b\x98\x7b\xce\x15\xcc\x39\x57\x22\x27\x98\x77\xce\x15\x6c\x04\xce\x15\xcc\x33\xe7\x0a\x2c\x2e\x07\x2f\x76\xe6\x38\x57\x60\x9b\xdf\x04\x60\x07\x11\xbd\x08\xe0\x85\x10\xc2\x76\xe1\x14\x64\xc5\xce\x22\x57\x22\x27\x88\x39\xc1\x4d\x00\xbe\x07\xb0\xce\x1b\x11\x95\x3f\x01\xfc\x00\xe0\x5e\x00\x87\xa9\xd0\xcb\x37\x0b\x9c\x2b\x10\x8e\x78\x1f\xc0\x35\x6a\x3c\x31\x1f\x73\x9f\x01\xd8\x22\xfa\x12\x46\xbe\x99\xe2\x5c\xc1\xfa\xfa\x3a\x1f\x80\x02\xc0\x17\x8e\xf9\x78\xff\x75\x00\xe7\x94\x05\x6a\xd4\x98\x16\xe7\x0a\x94\x03\x70\x71\x85\xd1\x1c\x77\x10\xc0\xa3\x00\x8e\x91\x1c\xd9\x1a\x53\xe1\x5c\x81\x20\x84\x70\xa3\x61\xf4\x0f\x00\x7f\x56\x98\x8f\xf7\xf7\x03\xb8\x5d\x73\xf1\x15\x95\xd6\xb0\xea\x36\xc1\x95\xc8\x09\x74\x1f\xc0\x56\xc3\xe8\xb7\x45\x51\x5c\x08\xe0\xa5\x8c\xf9\x98\x7b\xaf\x28\x8a\x2b\xea\xd6\x6d\x82\x73\x05\xca\xe9\x15\x90\x98\x5a\x03\xb0\xc0\x7f\x6f\xb5\x5a\x9b\x01\xec\xc9\x98\x8f\xb7\xe7\x00\x2c\x71\x5c\x6f\x20\x15\x75\x9b\xe0\x7c\x41\x17\x5b\x0d\x53\x07\x3a\x9d\xce\x71\x91\x86\x63\x79\x4d\xf0\x5d\xc6\xbc\x72\x6b\x21\x84\xfb\x74\xda\x94\x46\x5b\xfe\x6b\xe4\xc6\x32\x6e\xce\x15\x28\x17\x1f\x80\xc8\xd4\xda\xe2\xe2\xe2\xb1\x89\x8e\xc1\x8b\xa2\xa7\x89\xe8\x50\x85\xf9\x78\xfb\x14\xc0\xf5\x1a\x28\x07\x22\x3b\x96\xb1\x72\xae\x40\x50\xd1\x04\x0f\xc4\x1d\x3e\x8e\x15\xf0\xcc\xf1\x76\xc6\x7c\xcc\xf5\x4d\x9b\x4d\x5d\x0d\x25\x72\x82\xc8\xd0\x8a\x61\x60\xad\xd3\xe9\x1c\x5f\x15\x1b\xe1\x66\x22\xda\x97\x31\xaf\xdb\x41\x22\x7a\x1c\xc0\xd1\x51\xec\x44\x97\xd5\xae\x40\xb9\x8a\x26\xd8\xbb\x02\x72\xb1\x82\x45\x00\x0f\x03\xf8\xa5\xc2\x7c\xcc\xf1\xb4\x79\x47\x14\xeb\x8e\x6f\x14\xce\x17\x74\xe1\x36\x41\x2b\x56\x38\x5d\x4a\x33\xce\x20\xa2\x5d\x19\xf3\x31\xf7\x3e\x80\xcb\xa3\xd8\x5c\x8d\xa1\x38\x57\xa0\x5c\xdd\x26\x68\xc5\xf6\x15\xfc\x0f\x3c\x6d\x7e\x94\x31\x1f\xef\x3f\x9f\x4e\x9b\xb9\x1a\x03\x71\xae\x40\x30\x68\x13\xcc\xe4\x8b\xaf\x06\xde\xdf\xc1\x0b\xaa\x8c\x79\xe5\x78\xcd\xd1\x77\xb7\x99\xa9\x51\x9b\xab\x95\x48\x65\xc6\xc0\xb2\x4d\xd0\xe1\xe2\x03\xb1\x89\x88\x9e\x01\x70\xc8\xa8\x91\x1e\x10\xbe\xdb\xec\x9b\x36\x33\x35\x5c\xce\x15\x28\x37\x4a\x13\xf4\x38\xfe\x29\xb8\x04\xc0\x5b\x19\xf3\x31\xb7\xbb\xdd\x6e\x9f\xab\x81\x5e\x8d\x1c\xe7\x0b\xba\x18\xa5\x09\x7a\x5c\xef\x2c\x32\x42\x08\xdb\x00\x7c\x99\x31\xaf\x1b\xdf\x6d\x3e\x96\xfe\x1b\x56\xd4\x30\x39\x57\xa0\xdc\xb8\x9a\xa0\xc7\x49\x2d\x86\x4e\x9b\xbf\x1a\x75\xd3\x03\xb2\x3f\x84\xd0\x9b\x36\xbd\x1a\x7d\x9c\x2b\x10\x8c\xb1\x09\xd6\xe6\x18\x0b\x0b\x0b\x67\x01\xd8\x95\x31\x1f\x6f\xaf\xc8\x83\x1b\x33\x9f\xc5\x95\xc8\x09\x74\x9f\x29\xa3\xe0\x28\x4d\xb0\x36\x27\xe8\x4d\x9b\x15\xe6\x95\x7b\x44\xe3\xab\xf2\xf5\x79\xf3\x04\xca\x4d\xb2\x09\xd6\xe1\x04\xfc\xfb\x76\x22\xea\xdd\x6d\x1a\x07\x84\xfb\xc2\xc9\x5e\xbe\x98\xf3\x05\x5d\x4c\xb2\x09\xd6\xe6\x04\x7c\xb7\xc9\xd3\xe6\x4f\x89\x79\xdd\x6e\x55\xa1\x9b\xcf\x13\x44\x67\xa0\x91\x26\x58\x83\x8b\xd7\x0f\xf7\x58\xb7\xdc\x21\x84\x07\x2b\x62\xff\xcf\xb9\x02\xc1\x34\x9a\xa0\xc5\x09\x4e\x00\xf0\x2c\x11\xf5\x66\x88\xe4\xe4\xdc\xa2\xf1\xb5\xf2\xe5\x04\xba\xcf\x94\x71\xa9\x35\xd2\x04\xa3\x31\x30\xf8\xe1\xea\xd7\xc6\x58\xd4\x3c\x1f\x94\x13\x55\x6c\xe5\xeb\xf3\xe6\x09\x94\x9b\x52\x13\xec\x5d\xee\xed\x76\xfb\x3c\x5d\x25\x66\xcc\xf3\xe5\x7f\xbf\x84\xd4\x7e\x8e\xe0\x0b\xba\x68\xbc\x09\x0a\x8e\x00\xb0\x53\xba\x7b\xd6\x3c\x11\xf1\xaa\xb0\x84\x95\xcf\xe2\x5c\x81\x72\x4d\x37\x41\xc1\x75\x00\x3e\x37\xea\xa6\xe6\xf7\x00\xb8\x4a\x83\xac\x7c\x95\x9c\x2b\x10\x34\xd1\x04\x35\x0f\x80\xd3\x00\xbc\x6c\x18\x4d\xcd\xff\xcc\x33\x81\xae\xfe\xd2\x7c\x56\x8d\x94\xab\x15\xa4\x32\x63\x10\xe3\x68\x82\xa4\x4f\x82\x19\x21\x84\xbb\x39\x6f\x0d\xf3\xaf\x02\x38\x5d\xe3\x06\xf9\x9f\xef\xf3\xe6\x09\x94\x9b\x50\x13\x8c\xe7\xf4\xcb\x88\xe8\x03\xa3\x46\x6a\x7e\x5f\xab\xd5\x5a\xd1\xa0\x1a\x35\x5c\xce\x17\x74\x31\xf6\x26\x28\xe0\x1e\xf2\x94\x61\x34\x35\xff\xb7\xbc\x68\x3d\x52\x03\xd3\x7c\x56\x0d\x8f\x73\x05\xd1\x60\xc7\xd2\x04\x7b\x85\xbb\xb8\x8d\x6f\x65\x6b\x98\x7f\x07\xc0\x05\x1a\xe4\xd5\x18\x88\x73\x05\x82\x31\x34\xc1\x78\x4e\x3f\x1b\xc0\x9b\x46\xbe\xd4\x3c\xdf\xf4\xdc\xa9\x71\x35\x6a\x0c\xcc\x95\xc8\x09\x74\x9f\x29\x63\xb0\xb5\x9b\xa0\xe0\x70\xa6\x01\xfc\xee\x99\x27\x22\x7e\x12\x7c\x52\x14\x3b\x91\x37\x45\xae\x40\xb9\x51\x9a\xa0\xe0\x5a\x79\xa0\x99\x9e\xe5\xd4\xfc\xc7\x00\xae\xd6\x20\x2b\xdf\xb8\x39\x5f\xd0\xc5\x40\x4d\x30\x8a\x3b\x15\x00\x7f\x4d\x96\xc6\xa6\xe6\xf9\x8d\xd1\x03\xfa\xba\x3d\xcd\x97\x1b\xdf\x28\x9c\x2b\x50\xae\x66\x13\x4c\xe7\x74\x7e\x55\xfe\x63\x0d\xf3\xbb\x01\x9c\xa9\x71\xc3\xce\xe9\x43\x71\xae\x40\x50\xa3\x09\xc6\x73\xfa\xa5\xf2\x5a\x2b\x35\x9a\x9a\xff\x2a\x84\xc0\x9f\xdd\x95\x48\xeb\x5a\x63\x19\x37\x57\x22\x27\xd0\x7d\xa6\x0c\x03\xbd\x26\x28\xe0\xb7\xba\x4f\x5a\x0f\x29\x92\xd8\x7f\x00\x3c\x01\xe0\x28\x0d\xb4\xea\x36\xc1\xb9\x02\xe5\x2a\x9a\x20\x2f\x59\xdb\xfc\x77\x7e\x05\xce\x67\xd4\x38\x48\xa9\xf9\x77\xe5\xbb\x81\x12\x5e\xdd\x26\x38\x5f\xd0\x45\xdf\x01\x10\x53\xdf\x14\x45\x71\x11\x80\xd7\x0c\xa3\xa9\x79\xfe\xba\xf4\xae\x28\x9f\x5b\xb7\x09\xce\x15\x28\x07\xe0\x06\xc3\x14\x7f\x26\xf7\x97\x67\x5e\xbe\x29\x3e\x45\xf2\x30\x26\x32\xa7\x0f\xc5\xb9\x02\x41\xbb\xdd\x3e\xdf\x33\x6a\x70\x9f\xc8\xfc\x5f\x22\x57\x63\x5a\x5c\x89\x9c\x40\xf7\xf9\xac\x11\xd1\x87\x15\x46\x53\xee\x37\x00\x0f\xc9\xca\xaf\x56\x8d\x69\x71\xae\x40\x38\x9d\xe2\xf8\xa3\x48\xbe\xec\x73\xe6\xdf\x00\xb0\x2c\x7a\x46\x73\x73\xfa\x90\x9c\x2b\x50\x8e\x51\x14\xc5\x95\x00\xf6\x1a\xe6\xf7\x86\x10\xfa\x5e\x48\x78\xf9\x66\x81\x73\x05\x16\xb7\xb4\xb4\xc4\x97\xf6\x16\x22\xda\x19\x42\x79\x73\xb3\x79\x79\x79\xb9\xb7\x84\xcd\xc5\xce\x1c\xe7\x0a\x2a\x38\x0b\x96\x6e\xd6\xb9\x12\x39\xc1\xbc\x73\xae\x60\x23\x70\xae\x60\x9e\x39\x57\x30\xf7\x9c\x2b\x98\x73\xae\x44\x4e\x30\xef\xdc\xbf\x62\xba\x18\x72\x44\x36\x65\x8a\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\xce\xc0\x5c\x37\x69\x06\x00\x00")

func start_ico() ([]byte, error) {
	return bindata_read(
		_start_ico,
		"start.ico",
	)
}

var _status_ico = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x00\x86\x01\x79\xfe\x00\x00\x01\x00\x01\x00\x10\x10\x00\x00\x00\x00\x08\x00\x70\x01\x00\x00\x16\x00\x00\x00\x89\x50\x4e\x47\x0d\x0a\x1a\x0a\x00\x00\x00\x0d\x49\x48\x44\x52\x00\x00\x00\x10\x00\x00\x00\x10\x08\x06\x00\x00\x00\x1f\xf3\xff\x61\x00\x00\x00\x01\x73\x52\x47\x42\x00\xae\xce\x1c\xe9\x00\x00\x00\x04\x67\x41\x4d\x41\x00\x00\xb1\x8f\x0b\xfc\x61\x05\x00\x00\x00\x09\x70\x48\x59\x73\x00\x00\x12\x74\x00\x00\x12\x74\x01\xde\x66\x1f\x78\x00\x00\x01\x05\x49\x44\x41\x54\x38\x4f\x8d\x93\x4b\x0e\x44\x40\x18\x84\x3d\x62\x81\xad\xd8\x70\x08\x67\xe0\x60\x2c\x39\x06\xb7\xf2\xb8\x82\x10\x07\xd0\x93\xaa\x61\xa6\xfd\x4d\x66\x2a\xa9\x49\xa7\xeb\xfb\x2b\xd3\x9d\x66\xed\xfb\xae\x74\xff\x92\xe4\x3f\x05\xa7\x86\x61\x50\x6d\xdb\xaa\xba\xae\x55\x59\x96\x34\xd6\xd8\xeb\xfb\xfe\xa0\xbe\x45\x2c\x80\xd6\x75\x55\x96\x65\xfd\xe5\x65\x59\x38\x73\x29\xb0\x6d\xdb\x00\xc3\x30\xa4\xe5\x3e\x0c\xb1\x00\x0b\xfc\x6d\x09\x38\x8e\x43\x08\xc2\x5a\xe6\x98\x81\x58\x80\xf3\x49\x00\x3e\x75\x97\x61\x86\x19\x7e\x70\x49\x12\xf0\x3c\xcf\xd8\xd3\xdd\x34\xcd\xb7\x00\x37\x2d\x81\x79\x9e\x09\x40\xbe\xef\x1b\x39\x66\xa0\xc7\x82\x20\x08\x08\x40\x49\x92\x18\x79\x55\x55\xcc\x1e\x8f\x10\xc7\x31\x01\x28\x8a\x22\x23\xbf\x1c\xe1\xee\x12\x8b\xa2\x20\x00\xe5\x79\x6e\xe4\x5d\xd7\x31\x63\x01\x5e\x98\x04\xa4\x64\x3e\x4d\xd3\x7b\xff\x7c\x48\x12\x48\xd3\x54\x65\x59\x46\x63\xad\x67\xae\xeb\x72\xe6\xf2\x12\xf1\x3c\x75\xe8\xc9\x18\xde\xb6\x8d\x33\x9f\x82\xb3\x04\xba\xfb\x98\x70\x61\x38\xf3\x38\x8e\x07\x25\x3e\x26\xdd\xbf\x74\xe5\x77\xf5\x02\xda\x0b\x6c\xbc\x94\x95\x7c\xbb\x00\x00\x00\x00\x49\x45\x4e\x44\xae\x42\x60\x82\x01\x00\x00\xff\xff\x97\xe7\xb6\x58\x86\x01\x00\x00")

func status_ico() ([]byte, error) {
	return bindata_read(
		_status_ico,
		"status.ico",
	)
}

var _stop_ico = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x62\x60\x60\x64\x60\x64\x70\x70\x60\x60\x60\x60\xe0\x60\x60\x61\x66\x60\x10\x63\x60\x60\xe8\x0c\xf0\x73\xe7\xe5\x92\xe2\x62\x60\x60\xe0\xf5\xf4\x70\x09\x62\x60\x60\x00\xa9\x70\xe0\x60\x63\x60\x60\x58\x95\x59\x78\x8f\x81\x81\x81\xb1\x38\xc8\xdd\x89\x61\xdd\x39\x99\x97\x0c\x0c\x0c\x2c\xe9\x8e\xbe\x8e\x0c\x0c\x1b\xfb\xb9\xff\x24\xb2\x32\x30\x30\x70\x16\x78\x44\x16\x33\x30\x88\xa9\x82\x30\xa3\x67\x90\xca\x07\x06\x06\xa6\x99\x9e\x2e\x8e\x21\x15\x71\x6f\x67\x05\xe6\xdd\x36\xe0\x69\xd3\xbc\xda\xd0\xa8\xc8\x39\x2d\xd5\xda\x61\xa3\xdc\xc2\x03\x7a\x46\xe6\x37\x37\xec\x7e\xb2\xf7\x40\xdf\xad\x25\xea\x19\x6c\x59\x29\x0a\xda\x93\x66\xf0\x4d\xd6\xd5\x50\x51\xb8\xf1\xef\x83\xb6\x50\xe6\xab\xfa\x4e\x46\x86\x84\xfd\xf7\xa3\x65\x2a\xec\x79\x0e\xdf\xd5\x35\x0c\xbc\x60\x23\xcc\xda\x6b\x2c\x6e\xe9\xf8\xda\x47\xe4\x0e\x93\x30\xab\xae\x31\x8a\x50\xcb\x9a\x6f\x16\x73\xfd\xd4\xee\x14\xeb\xf0\xcf\xef\x64\xb2\xff\xb9\xee\xb5\x1c\xe3\x83\xdd\x5b\xcb\xb8\x19\xfc\xaf\x45\xc9\x54\xd8\x9f\x39\xc3\x76\x6b\x76\x77\x8c\xdc\xff\x49\x7b\x0e\x56\x27\x9f\x31\x16\xbb\x29\x51\x7c\x68\xb5\xbd\x56\xd2\x0e\xf7\xed\x77\x3d\xe4\xa6\xef\xec\x64\xe2\x5a\x95\x97\x73\x50\x6a\xe6\xeb\x16\x83\xfe\x90\xbc\x48\x46\x7b\xe6\x03\x7a\x9b\x13\x42\xcf\x44\xfd\x7e\xfd\xea\x8e\x0a\xa3\xb1\x55\xfa\x45\x56\xed\x94\xdd\x21\x72\x86\x6a\xdc\xff\x8e\x05\xa7\x1c\x34\x67\xa8\x60\x6d\x60\x66\x90\xe9\x35\x74\xdf\x37\xbf\xc1\xb0\xbe\x5b\xe4\x88\xa8\xc2\x19\xe5\x59\xb3\x55\x3e\x9c\x3b\xb5\x93\xe1\x5e\xd3\xfa\xda\xfb\xf6\x0c\x07\x96\x1f\x4d\xfa\xd8\xba\x6e\xfe\xf4\xd3\x4f\x97\x99\x1c\xd6\x66\xb8\x23\x9d\xbc\x41\x60\x31\x53\x44\x76\xfd\x92\x4f\x3b\x12\x16\x35\x34\xf5\x3b\x04\xbc\x15\x67\x88\xd8\x90\xf6\xea\x46\x8e\x99\xf7\x2d\xe7\x1b\x33\x77\x5c\x5f\x5d\x7c\x9a\xeb\x87\x64\xe0\x92\xab\x7d\x5f\xc2\x26\x15\x7b\x1b\x76\x6d\x8d\x67\x38\xcf\x12\xcb\xbe\xe4\x59\x5f\xd1\x64\x96\x33\x42\x6a\x1d\x6f\xdb\xd7\x56\x84\x6e\xfd\x6f\x1e\xbd\x36\x6c\xc7\x36\xed\x03\x31\x6a\x15\x0b\x3e\xbf\x50\x32\xdb\x2e\x72\x58\x21\xbf\xa2\xb3\xe4\x59\x8a\x5a\xc5\x12\xdf\x94\xe7\xd3\x12\xb6\x2c\xef\x3f\x36\x63\x22\xfb\x83\x19\x97\x58\xce\xbb\x2b\x98\x6e\x57\x79\xcc\xd7\x51\xb2\x6c\x96\x02\x77\x52\xb1\xfb\x4e\x4f\xa6\xb4\xac\x26\x08\x9c\x18\x16\x2c\x1c\x02\x85\x67\xce\xea\x1e\xd5\xb8\xc2\x72\x15\x0a\x0f\x9f\x83\x49\xa4\xe4\x66\xad\x9b\x1f\x1d\x76\xe0\x7c\x87\xd2\xda\x65\x6a\x2b\xd4\x2a\x1e\xab\x55\x2c\xb0\xd7\x78\xde\xc5\x5f\x3c\x5b\xe1\x9f\x7d\xbd\x7f\x7a\x61\xaf\xae\x67\x7e\xea\xd9\xeb\x93\xd6\xbc\xff\xbd\x7e\x4b\xfb\x06\x63\x75\x63\xc7\x3f\x16\x8c\xe6\x0d\xf9\x8c\xec\x55\x4c\xdb\x54\xec\xa7\xa7\xbd\x63\xee\xdc\x60\xbc\xdc\xe4\xe0\x06\x85\xad\x1b\x94\xd4\x3a\xec\x67\x45\x6d\x30\xfb\x36\xd9\x5c\xf7\x99\x1a\xe7\x9e\xcd\x8f\xec\x26\x7b\xeb\x44\x88\xbe\xde\xf5\x2e\x5a\xab\x9f\x5b\x8d\xeb\x98\xdd\x07\x99\xe8\x0d\x46\xdf\xac\xa7\xbb\x9f\x88\x79\xc2\x62\xb7\xb8\x62\xcb\xb5\x77\x05\x0f\x36\x2a\x1c\xbe\xf9\xa8\x34\x6d\xdb\x53\xf3\xff\xf3\xde\xb7\xba\x3a\x9e\xc8\xea\xff\xae\x52\x5f\xb7\x71\xd2\x82\x54\x35\xae\x59\xe9\x8b\x16\x14\x3e\x53\x52\xcf\x98\xdb\xd9\x14\x9d\xb0\x08\x55\xe8\x2f\x23\xef\xfb\x3b\x7f\x0f\xfd\xfc\x93\x0f\xca\x05\x9e\xae\x7e\x2e\xeb\x9c\x12\x9a\x00\x01\x00\x00\xff\xff\x7e\x7c\xd1\xb2\x1a\x03\x00\x00")

func stop_ico() ([]byte, error) {
	return bindata_read(
		_stop_ico,
		"stop.ico",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
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
var _bindata = map[string]func() ([]byte, error){
	"main.ico": main_ico,
	"start.ico": start_ico,
	"status.ico": status_ico,
	"stop.ico": stop_ico,
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
	Func func() ([]byte, error)
	Children map[string]*_bintree_t
}
var _bintree = &_bintree_t{nil, map[string]*_bintree_t{
	"main.ico": &_bintree_t{main_ico, map[string]*_bintree_t{
	}},
	"start.ico": &_bintree_t{start_ico, map[string]*_bintree_t{
	}},
	"status.ico": &_bintree_t{status_ico, map[string]*_bintree_t{
	}},
	"stop.ico": &_bintree_t{stop_ico, map[string]*_bintree_t{
	}},
}}
