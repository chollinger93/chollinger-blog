#!/bin/bash
set -euo pipefail

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)"

if [[ ! $UID -eq 0 ]]; then
	echo "Script must be called as root"
	exit 1
fi

WWW_ROOT=/var/www/
BACKUP_DIR=$HOME/backups/blog/"$(date --iso)"

mkdir -p "$BACKUP_DIR"

cd "${DIR}/.."
git submodule update --init --recursive && git submodule update --remote
hugo --minify -d blog
mv "$WWW_ROOT/blog" "$BACKUP_DIR"
rm -rf "$WWW_ROOT/blog"
mv blog "$WWW_ROOT"
chown -R www-data:www-data "$WWW_ROOT/blog"
