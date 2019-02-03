# CopyMyScripts
自分用のスクリプトをコピーしてくるやつ

# コピー対象
例えばこんなリポジトリ

https://github.com/yumechi/LibModoki

# 使い方

これのリポジトリのローカルのパスを `templateDir` にセットする（home directory以下）

あとは下記のようなコマンドを打ち込む。

```
go run copy_template.go 
        -prefix CPP \
        -suffix .cpp \
        -input template \
        -output prog_a
```

すると `~/templates/comp_template/CPP/template.cpp` が現在のディレクトリに `prog_a.cpp` としてコピーされる。

# ラッパーも書いとく

```.zshrc
alias cpptouch='zsh ~/scripts/cpptouch.sh'
```

とか `.zshrc` にかいとく。

`cpptouch.sh` の中身はこんな感じ。

```cpptouch.sh
#!/bin/zsh -ue
# @(#) MyTouch sample
# @(#) Author: @yumechi

FROMFILE="template"
EXECUTOR=${HOME}"/scripts/copy_template.go"

PARAM=()
for opt in "$@"; do
    case "${opt}" in
        "-n" ) FROMFILE="$2"; shift 2 ;; # @(#) touch 
        '--' | '-' )
            shift
            PARAM+=( "$@" )
            break
            ;;
        * ) if [[ -n "$1" ]] && [[ ! "$1" =~ ^-+ ]]; then
                PARAM+=( "$1" ); shift
            fi
            ;;
    esac
done

# if OUTPUTFILE is None, Go Scripts create file by timedate.
TOFILE="${PARAM}"; PARAM=("${PARAM[@]:1}")

function my_copy() {
    local prefix=$1
    local suffix=$2
    local inputfile=$3
    local outputfile=$4
    if [ -n "$outputfile" ]; then
        go run ${EXECUTOR} \
        -prefix $prefix \
        -suffix $suffix \
        -input $inputfile \
        -output $outputfile
    else
        go run ${EXECUTOR} \
        -prefix $prefix \
        -suffix $suffix \
        -input $inputfile
    fi
}

my_copy CPP .cpp $FROMFILE $TOFILE
```

みたいなラッパーを書いとく。

`touchcpp prog_a` みたいに使うことができる。

# 機能

* template 以外もファイル名指定で取れるよ
* デフォルトは template ファイルがコピーされるよ
* 出力ファイル名は指定がなければ時間ベースで作成するよ
* コピー元がシンボリックリンクでも頑張って引っ張る（自分のローカルはシンボリックリンクになっているので）

くらいなもん。


