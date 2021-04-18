package main

type Post struct {
    Id      string `json:"id"`
    User    string `json:"user"`
    Message string `json:"message"`
    Url     string `json:"url"` // post 媒体的存储路径
    Type    string `json:"type"`  //存图片还是视频, 前端用
}

// http handle function, why one is *, the other is not
// http request and writer defination