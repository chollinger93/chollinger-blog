#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"
BASE_DIR=$DIR/../content #/posts/2022/09

for ix in $(find "$BASE_DIR" -name index.md -type f); do 
    dn=$(dirname $ix)
    for i in $(find "$dn" -type f | file --mime-type -f - | grep -F image/ |  rev | cut -d : -f 2- | rev); do
        img_gr=$(basename $i)
        if [[ -z $(grep $img_gr "$ix") ]]; then 
            echo "Unused: $ix,$img_gr"
            git rm $i
        fi
    done
done
