arg CADDY_VERSION 2
arg REPOSITORY_NAME
from caddy:${CADDY_VERSION}-builder as builder

add . /src/plugin

run xcaddy build \
    --with github.com/${REPOSITORY_NAME}=/src/plugin

from caddy:${CADDY_VERSION}
copy --from=builder /usr/bin/caddy /usr/bin/caddy
