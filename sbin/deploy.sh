#!/bin/bash
set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

if [[ ! $UID -eq 0 ]]; then
	echo "Script must be called as root"
	exit 1
fi

cd "${DIR}/.."
git submodule update --init --recursive && git submodule update --remote
hugo --minify -d blog
mv blog /var/www/
chown -R www-data:www-data /var/www/blog
