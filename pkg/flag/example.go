package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	h bool

	v, V bool
	t, T bool
	q    *bool

	s string
	p string
	c string
	g string
)

func init() {
	flag.BoolVar(&h, "h", false, "help info")

	flag.BoolVar(&v, "v", false, "show version")
	flag.BoolVar(&V, "V", false, "show version and configure options")

	flag.BoolVar(&t, "t", false, "test configure and exit")
	flag.BoolVar(&T, "T", false, "test configuration, dump it and exit")

	// 另一种绑定方式
	q = flag.Bool("q", false, "supress non-error messages during configuration testing")

	flag.StringVar(&s, "s", "", "send `signal` to a master process: stop, quit, reopen, reload")
	flag.StringVar(&p, "p", "/usr/local/nginx", "set `prefix` path")
	flag.StringVar(&c, "c", "conf/nginx.conf", "set configuration `file`")
	flag.StringVar(&g, "g", "conf/nginx.conf", "set global `dirextives` out of configuration file")

	// 改变默认的 usage, usage是一种函数类型，这里覆盖默认函数实现
	flag.Usage = usage

}

func usage() {
	fmt.Fprintf(os.Stderr, `nginx version: nginx/1.10.0 Usage: nginx [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives] Options:`)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
	}
}
