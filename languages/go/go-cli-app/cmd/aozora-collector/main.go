package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/japanese"
)

type Entry struct {
	AuthorID string
	Author   string
	TitleID  string
	Title    string
	SiteURL  string
	ZipURL   string
}

func getResopnseBody(url string) (*goquery.Document, error) {
	// URL から HTTP GET リクエストを実行
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// レスポンスのステータスコードをチェック
	if resp.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	// レスポンスボディを goquery で解析
	doc, err := goquery.NewDocumentFromReader(resp.Body)

	return doc, err
}

func findEntries(siteURL string) ([]Entry, error) {
	doc, err := getResopnseBody(siteURL)
	if err != nil {
		log.Fatal(err)
	}

	// URL パターンをコンパイル
	pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html$`)

	entries := []Entry{}

	// 解析したドキュメントを使用して、必要なデータを抽出
	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		token := pat.FindStringSubmatch(elem.AttrOr("href", ""))
		if len(token) != 3 {
			return
		}
		title := elem.Text()
		pageURL := fmt.Sprintf("https://www.aozora.gr.jp/cards/%s/card%s.html",
			token[1], token[2])
		author, zipURL := findAuthorAndZIP(pageURL) // 作者とZIPファイルのURLを取得

		if zipURL != "" {
			entries = append(entries, Entry{
				AuthorID: token[1],
				Author:   author,
				TitleID:  token[2],
				Title:    title,
				SiteURL:  siteURL,
				ZipURL:   zipURL,
			})
		}
	})

	return entries, nil
}

func findAuthorAndZIP(siteURL string) (string, string) {
	doc, err := getResopnseBody(siteURL)
	if err != nil {
		return "", ""
	}

	author := doc.Find("table[summary=作家データ] tr:nth-child(1) td:nth-child(2)").Text()

	zipURL := ""
	doc.Find("table.download a").Each(func(n int, elem *goquery.Selection) {
		href := elem.AttrOr("href", "")
		if strings.HasSuffix(href, "zip") {
			zipURL = href
		}
	})

	if zipURL == "" {
		return author, ""
	}

	if strings.HasPrefix(zipURL, "http://") || strings.HasPrefix(zipURL, "https://") {
		return author, zipURL
	}

	u, err := url.Parse(siteURL)
	if err != nil {
		return author, ""
	}

	u.Path = path.Join(path.Dir(u.Path), zipURL)
	return author, u.String()
}

func extractText(zipURL string) (string, error) {
	resp, err := http.Get(zipURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	r, err := zip.NewReader(bytes.NewReader(b), int64(len(b)))
	for _, file := range r.File {
		if path.Ext(file.Name) == ".txt" {
			f, err := file.Open()

			if err != nil {
				return "", err
			}

			b, err := io.ReadAll(f)
			f.Close()

			if err != nil {
				return "", err
			}

			b, err = japanese.ShiftJIS.NewDecoder().Bytes(b)

			if err != nil {
				return "", err
			}

			return string(b), nil
		}
	}
	return "", errors.New("contents not found")
}

func main() {
	listURL := "https://www.aozora.gr.jp/index_pages/person879.html"

	entries, err := findEntries(listURL)

	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		content, err := extractText(entry.ZipURL)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(entry.SiteURL)
		fmt.Println(content)
	}
}

/*
Go でうまくプログラムを作る上で重要なポイント
- 最終的な処理は深く呼び出した関数の中では行わない
- エラーが起きた場合はその場で強制終了せず main に戻す
*/
