export CLIENT_ID=
export TRAKT_USERNAME=

function cleanup() {
    rm -f trakt-watchlist-converter
}
trap cleanup EXIT

# Compile Go
CGO_CFLAGS_ALLOW=-Xpreprocessor GO111MODULE=on GOGC=off go build -mod=vendor -v -o trakt-watchlist-converter .
./trakt-watchlist-converter
