package ftp

import (
	"bytes"
	"github.com/jlaffaye/ftp"
	"go.uber.org/zap"
	"io"
	"mime/multipart"
	"net/url"
)

type FTP struct {
	client *ftp.ServerConn
	logger *zap.SugaredLogger
}

func NewFTP(rawUrl string, logger *zap.SugaredLogger) (*FTP, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return nil, err
	}
	c, err := ftp.Dial(uri.Host)
	if err != nil {
		return nil, err
	}
	password, isSet := uri.User.Password()
	if !isSet {
		return nil, err
	}
	err = c.Login(uri.User.Username(), password)
	if err != nil {
		return nil, err
	}
	err = c.ChangeDir(uri.Path)
	if err != nil {
		return nil, err
	}
	return &FTP{
		client: c,
		logger: logger,
	}, nil
}

func (f *FTP) Destroy() error {
	return f.client.Quit()
}

func (f *FTP) Upload(name string, file multipart.File) {
	defer f.closeFile(file)
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		f.logger.Errorf("copy file to buffer error: %v", err)
	}
	err := f.client.Stor(name, buffer)
	if err != nil {
		f.logger.Errorf("transfer file error: %v", err)
	}
}

func (f *FTP) closeFile(file multipart.File) {
	if err := file.Close(); err != nil {
		f.logger.Errorf("close file error: %v", err)
	}
}
