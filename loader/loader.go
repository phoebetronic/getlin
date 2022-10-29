package loader

import (
	"archive/zip"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Loader struct {
	git bool
	org string
	rep string
}

func New(con Config) *Loader {
	{
		con.Verify()
	}

	return &Loader{
		git: con.Git,
		org: con.Org,
		rep: con.Rep,
	}
}

func (l *Loader) Create(fil string) {
	var src string
	if l.git {
		src = fmt.Sprintf("https://raw.githubusercontent.com/%s/%s/main/%s.zip", l.org, l.rep, fil)
	} else {
		panic("github loader must be configured")
	}

	var dst string
	{
		dst = fmt.Sprintf("%s/data/tmp/%s.zip", l.rep, fil)
	}

	{
		l.download(src, dst)
	}

	{
		l.compress(fil)
	}
}

func (l *Loader) compress(fil string) {
	var err error

	var tmp string
	{
		tmp = fmt.Sprintf("./%s/data/tmp", l.rep)
	}

	var dst string
	{
		dst = fmt.Sprintf("./%s/data/", l.rep)
	}

	var src string
	{
		src = fmt.Sprintf("%s/%s.zip", tmp, fil)
	}

	var arc *zip.ReadCloser
	{
		arc, err = zip.OpenReader(src)
		if err != nil {
			panic(err)
		}
		defer arc.Close()
	}

	for _, f := range arc.File {
		if f.Name != fil {
			continue
		}

		var pat string
		{
			pat = filepath.Join(dst, f.Name)
		}

		{
			err = os.MkdirAll(dst, os.ModePerm)
			if err != nil {
				panic(err)
			}
		}

		var tar *os.File
		{
			tar, err = os.OpenFile(pat, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				panic(err)
			}
			defer tar.Close()
		}

		var ori io.ReadCloser
		{
			ori, err = f.Open()
			if err != nil {
				panic(err)
			}
			defer ori.Close()
		}

		{
			_, err = io.Copy(tar, ori)
			if err != nil {
				panic(err)
			}
		}
	}

	{
		err = os.RemoveAll(tmp)
		if err != nil {
			panic(err)
		}
	}
}

func (l *Loader) download(src string, dst string) {
	var err error

	{
		err = os.MkdirAll(filepath.Dir(dst), 0755)
		if err != nil {
			panic(err)
		}
	}

	var sta time.Time
	{
		sta = time.Now()
	}

	var fil *os.File
	{
		fil, err = os.Create(dst)
		if err != nil {
			panic(err)
		}
		defer fil.Close()
	}

	var res *http.Response
	{
		res, err = http.Head(src)
		if err != nil {
			panic(err)
		}
		defer res.Body.Close()
	}

	var siz int64
	{
		siz, err = strconv.ParseInt(res.Header.Get("Content-Length"), 10, 64)
		if err != nil {
			panic(err)
		}
	}

	var don chan struct{}
	{
		don = make(chan struct{})
	}

	{
		go prgrss(don, dst, siz)
	}

	var bod io.ReadCloser
	{
		res, err := http.Get(src)
		if err != nil {
			panic(err)
		}

		bod = res.Body

		defer bod.Close()
	}

	{
		_, err = io.Copy(fil, bod)
		if err != nil {
			panic(err)
		}
	}

	{
		don <- struct{}{}
	}

	fmt.Printf("Downloaded %s in %s\n", filepath.Base(dst), time.Since(sta).Round(100*time.Millisecond))
}
