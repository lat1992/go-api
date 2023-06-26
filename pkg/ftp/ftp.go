package ftp

import (
	"bytes"
	"fmt"
	"github.com/jlaffaye/ftp"
	"io"
	"mime/multipart"
	"net/url"
)

type FTP struct {
	client *ftp.ServerConn
}

func NewFTP(rawUrl string) (*FTP, error) {
	uri, err := url.Parse(rawUrl)
	if err != nil {
		return nil, fmt.Errorf("NewFTP: %w", err)
	}
	c, err := ftp.Dial(uri.Host)
	if err != nil {
		return nil, fmt.Errorf("NewFTP: %w", err)
	}
	password, isSet := uri.User.Password()
	if !isSet {
		return nil, fmt.Errorf("NewFTP: %w", err)
	}
	if err = c.Login(uri.User.Username(), password); err != nil {
		return nil, fmt.Errorf("NewFTP: %w", err)
	}
	if err = c.ChangeDir(uri.Path); err != nil {
		return nil, fmt.Errorf("NewFTP: %w", err)
	}
	return &FTP{
		client: c,
	}, nil
}

func (f *FTP) Destroy() error {
	return f.client.Quit()
}

func (f *FTP) Upload(name string, file multipart.File) error {
	buffer := bytes.NewBuffer(nil)
	if _, err := io.Copy(buffer, file); err != nil {
		if err = f.closeFile(file); err != nil {
			return fmt.Errorf("close file error: %w", err)
		}
		return fmt.Errorf("copy file to buffer error: %w", err)
	}
	if err := f.client.Stor(name, buffer); err != nil {
		if err = f.closeFile(file); err != nil {
			return fmt.Errorf("close file error: %w", err)
		}
		return fmt.Errorf("transfer file error: %w", err)
	}
	if err := f.closeFile(file); err != nil {
		return fmt.Errorf("close file error: %w", err)
	}
	return nil
}

func (f *FTP) closeFile(file multipart.File) error {
	if err := file.Close(); err != nil {
		return fmt.Errorf("close file error: %w", err)
	}
	return nil
}
