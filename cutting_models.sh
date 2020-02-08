#!/bin/bash

#不要なコメントを削除
grep -l 'type' ./models/*.xo.go | xargs sed -i '' -e 's/\/\/ Package.*$//'
grep -l 'type' ./models/*.xo.go | xargs sed -i '' -e 's/\/\/ Code generated by xo.*$//'
grep -l 'type' ./models/*.xo.go | xargs sed -i '' -e 's/\/\/ xo fields.*$//'
grep -l 'type' ./models/*.xo.go | xargs sed -i '' -e 's/_exists.*$//'
grep -l 'type' ./models/*.xo.go | awk '{ sub("m/\/\/ Exists determines.*$",""); print $0; }'

#null系カラムの対処
grep -l 'type' ./models/*.xo.go | xargs sed -i '' -e 's/sql\.NullInt64/\*int/g'
grep -l 'type' ./models/*.xo.go | xargs sed -i '' -e 's/sql\.NullString/\*string/g'
grep -l 'type' ./models/*.xo.go | xargs sed -i '' -e 's/sql\.NullBool/\*bool/g'
grep -l 'type' ./models/*.xo.go | xargs sed -i '' -e 's/mysql\.NullTime/\*time.Time/g'

grep -l 'type' ./models/*.xo.go | while read file
do

  #構造体以降削除
  line=$(grep $file -e "Exists determines" -n |  sed -e "s/\(.*\):.*$/\1/g")
  if [ -n "$line" ]; then
    cmd="sed -i '' -e '$line,\$d' $file"
    eval ${cmd}
  fi

  #リネーム
  struct=$(grep $file -e "type \(.*\) struct.*$" | sed -e "s/type \(.*\) struct.*$/\1/g")
  mv $file ./models/${struct}.go
done

# 実際に使用するときには更にgofmtとgoimportをかけてます。
# make fmt
# make import