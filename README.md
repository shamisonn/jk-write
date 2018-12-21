# jk-write

これは静的WebページジェネレータJekyllにおいて
新たな投稿を行うための補助ツールです。
Jekyllにおける`_post`ディレクトリにファイル生成をしたり、
その中身を参照することができます。
なお`motemen/ghq`を強く参照して作られました。

# insall

```sh
$ go get shamisonn/jk-write
```

# usage

使い方を示します。

## prepare to use

jekyllにおいて`_post`の場所をフルパスで設定しておく必要があります。
```sh
$ export JK_WRITE_ROOT="/set/your/full/path/jekyll/_post"
```

## commands

### list
`_post`の中身のファイルを表示します。

```sh
$ jk-write list
2016-03-19-init.md
2016-12-31-diary.md
2017-01-05-movies.md
...
```

### new
`_post`に新たなファイルを生成します。必ず`title`を入力する必要があります。
```sh
$ jk-write new hoge # title を hoge としている。
Do you make new file?(2018-12-21-hoge.md)[y/n]: y # input y or n
new file: 2018-12-21-hoge.md
```

### root
`JK_WRITE_ROOT`の値を表示します。
```sh
$ jk-write root
/set/your/full/path/jekyll/_post
```

## for cli users

`fzf`(もしくは`peco`)と組み合わせて便利なコマンドを作ることができます。

`fish`での実際のコードを下記に示します。
```fish
alias jkn="jk-write new"
function jkw -d "write jekyll post"
	code (jk-write root)/(jk-write list -r | fzf)
end
```

## additional option
生成するファイルに予めテンプレートを書き込んでおきたい場合は
`$HOME/tmp.md`にそれを記述しておくとそれを検出し、
生成するさいにそのファイルのコピーを生成します。

```sh
$ cat ~/tmp.md
---
layout: default
title: taitoru
mode: post
---

<!more>
```


# for developers

you can use make
```sh
$ make build
$ make install
```

# if happen some problem

Write issue.
