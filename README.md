# golangでAPIのお勉強
---

## 概要
---
`waters`がDBの役割をしていて, httpメソッドで色々できる. (json)
初期で`waters`にはirohasuとevianのみが入っている.

ex:
```bash
curl -X POST -d '{"name":"CrystalGeyser","amount":500,"color":"blue","company": {"name":"CrystalGeyser","home":"US","establishment":1977}}' http://localhost:8080/

```
でPOSTできる.
実際に`http://localhost:8080/CrystalGeyser`にアクセス(or curl)すればJSONが返ってくる.


### TODO
---
* 一回一回初期化されちゃうのでRDB化したりしないとダメ
* RDBとの連携の勉強
* パッケージ化についての理解
* push時にlintなど


## 参考
---
日本語なさすぎたけどgolangめっちゃ書きやすくて好きになった.
https://qiita.com/roothybrid7/items/34578037d883c9a99ca8
https://pkg.go.dev/github.com/gorilla/mux
https://pkg.go.dev/net/http