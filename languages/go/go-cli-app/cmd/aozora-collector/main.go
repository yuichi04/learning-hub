package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

type Entry struct {
	AuthorID string
	Author   string
	TitleID  string
	Title    string
	InfoURL  string
	ZipURL   string
}

func findEntries(siteURL string) ([]Entry, error) {
	// URL から HTTP GET リクエストを実行
	resp, err := http.Get(siteURL)
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
	if err != nil {
		log.Fatal(err)
	}

	// URL パターンをコンパイル
	pat := regexp.MustCompile(`.*/cards/([0-9]+)/card([0-9]+).html$`)

	// 解析したドキュメントを使用して、必要なデータを抽出
	doc.Find("ol li a").Each(func(n int, elem *goquery.Selection) {
		token := pat.FindStringSubmatch(elem.AttrOr("href", ""))
		if len(token) != 3 {
			return
		}
		pageURL := fmt.Sprintf("https://www.aozora.gr.jp/cards/%s/card%s.html",
			token[1], token[2])
		println(pageURL)
	})

	// 処理
	return nil, nil // ひとまず両方 nil を返しておく
}

func main() {
	listURL := "https://www.aozora.gr.jp/index_pages/person879.html"

	entries, err := findEntries(listURL)
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		fmt.Println(entry.Title, entry.ZipURL)
	}
}
