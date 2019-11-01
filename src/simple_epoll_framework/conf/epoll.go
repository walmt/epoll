package conf

type EpollConf struct {
	Backlog string `xml:"backlog"`
	Ports   []int  `xml:"ports"`
}
