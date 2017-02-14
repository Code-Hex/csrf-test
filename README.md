# 使い方
## インストール
    go get github.com/Code-Hex/csrf-test

## コンパイル
    cd cmd/csrf && go build
    cd cmd/sns && go build

## 実行
    ./cmd/csrf
    ./cmd/sns

## 確認
`http://localhost:8080` へアクセス。ボタンをポチる。  
ログインができたはずなので、ブラウザから Cookie を見てみると `token=...` といった jwt が作成されています。  
ここで CSRF リンクをクリック。  

CSRF を行う側のサーバーのログを見てみるとポストしたことを確認できるメッセージと Cookie に保存していた token の値が出力されているはずです。