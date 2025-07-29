package jm

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
)

type Service struct {
	Address string

	client *http.Client
}

func NewService(address string) *Service {
	return &Service{
		Address: address,
		client:  &http.Client{},
	}
}

// Ping 检查服务是否存活
func (s *Service) Ping() error {
	resp, err := s.client.Get(s.Address + "/ping")
	if err != nil {
		return fmt.Errorf("ping 请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("服务异常，状态码: %d", resp.StatusCode)
	}

	return nil
}

// Download 下载 PDF 并返回 base64 字符串
func (s *Service) Download(id int64) (string, error) {
	url := s.DownloadUrl(id)
	resp, err := s.client.Get(url)
	if err != nil {
		return "", fmt.Errorf("下载请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("下载失败，状态码: %d，返回内容: %s", resp.StatusCode, string(body))
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %w", err)
	}

	base64Str := base64.StdEncoding.EncodeToString(data)
	return base64Str, nil
}

func (s *Service) DownloadUrl(id int64) string {
	return fmt.Sprintf("%s/download?id=%d", s.Address, id)
}
