# yuque
yuque sdk for golang

## api
https://www.yuque.com/yuque/developer/api

https://www.yuque.com/api/space_posts?offset=0
url := "https://www.yuque.com/api/docs/" + data.Target.Slug + "?book_id=" + strconv.Itoa(data.Target.BookID)
url := "https://www.yuque.com/api/docs/" + strconv.Itoa(doc.Data.ID) + "/pager"
"https://www.yuque.com/api/docs/" + paper.Data.Next.Slug + "?book_id=" + strconv.Itoa(bookID)
https://www.yuque.com/api/docs/e3c1c012-ca5a-4de1-aa2c-f8387a172c55?book_id=215218
https://www.yuque.com/api/books/215218/docs?include_contributors=true&include_hits=true&limit=20&offset=0
https://www.yuque.com/api/books/215218/docs?include_contributors=true&include_hits=true&limit=20&offset=20
https://www.yuque.com/api/mine/books?limit=200&offset=0
https://www.yuque.com/api/mine/marks?limit=200&offset=0&type=all
