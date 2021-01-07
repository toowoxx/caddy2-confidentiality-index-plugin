#!/bin/bash -e

set -euo pipefail

doit=false

while read file; do
	if [ "$file" -nt "pkged/pkged.go" ]; then
		doit=true
		break
	fi
done < <(find static -type f)

if $doit; then
	mkdir -p pkged
	cat <<EOF > pkged/ft.go
package pkged

var fileMap = map[string]string {
EOF
	for f in $(find static -type f); do
	  varname=$(basename "$f" | sed -re 's/[.]/_DOT_/g')
	  cat <<EOF > pkged/$(basename $f.go)
package pkged

const $varname string = "$(cat $f | base64 -w0)"

EOF
    cat <<EOF >> pkged/ft.go
    "$(echo "$f" | cut -c 1-)": $varname,
EOF
  done
  echo "}" >> pkged/ft.go
else
	echo "Static files haven't changed"
fi

