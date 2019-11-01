package conf

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

var Gconf GlobalConf

type GlobalConf struct {
	Epoll EpollConf `xml:"epoll"`
}

func (g *GlobalConf) FlushConf(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("os.Open %v failed, err:%v", path, err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("ioutil.ReadAll file failed, path:%v, err:%v", path, err)
	}

	conf := GlobalConf{}
	if err = xml.Unmarshal(data, &conf); err != nil {
		return fmt.Errorf("xml unmarshal failed, err:%v", err)
	}

	*g = conf
	return nil
}
