package service

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"
	repository "toz/pkg/url/repository"
)

type UrlService struct {
	rep  repository.Url
	l    *SingleList
	last string
}

func NewUrlService(rep repository.Url, l *SingleList, last string) *UrlService {
	return &UrlService{rep: rep, l: l, last: last}
}

var (
	checkUrl      = regexp.MustCompile(`^(?:http(s)?:\/\/)?[\w.-]+(?:\.[\w\.-]+)+[\w\-\._~:/?#[\]@!\$&'\(\)\*\+,;=.]+$`)
	checkShortUrl = regexp.MustCompile(`^[\w_]{10}$`)
)

func (s *UrlService) Create(full string) (string, error) {
	if !checkUrl.MatchString(full) {
		return "", fmt.Errorf("Invalid URL.")
	}
	short := s.l.GetUrl()
	s.last = short
	return s.rep.Create(short, transformURL(full))
}

func transformURL(s string) string {
	s = strings.ToLower(s)
	s = strings.Trim(s, " ")
	if len(s) >= 256 {
		s = s[:256]
	}
	return s
}

func normalizeURL(inputURL string) string {
	if !strings.HasPrefix(inputURL, "http://") && !strings.HasPrefix(inputURL, "https://") {
		if !strings.HasPrefix(inputURL, "www.") {
			inputURL = "www." + inputURL
		}
		inputURL = "http://" + inputURL
	}
	return inputURL
}

func pingURL(url string) bool {
	url = normalizeURL(url)
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		return false
	}
	defer response.Body.Close()
	return response.StatusCode == http.StatusOK
}

func (s *UrlService) GetFull(short string) (full string, err error) {
	if !checkShortUrl.MatchString(short) {
		return "1", fmt.Errorf("No short URL.")
	}
	if short < s.last {
		return "2", fmt.Errorf("Invalid short URL.")
	}
	return s.rep.GetFull(short)
}
