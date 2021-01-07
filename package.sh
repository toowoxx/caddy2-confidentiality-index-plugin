#!/bin/bash -e

doit=false

while read file; do
	if [ "$file" -nt "pkged/pkged.go" ]; then
		doit=true
		break
	fi
done < <(find static -type f)

if $doit; then
	go get github.com/markbates/pkger/cmd/pkger
	mkdir -p pkged
	$(go env GOPATH)/bin/pkger -include /static -o pkged/
else
	echo "Static files haven't changed"
fi

