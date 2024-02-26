# smil2png

## Usage

```
% ./smil2png.arm64.apple -h
Usage of ./smil2png.arm64.apple:
  -f int
    	frames (default 10)
  -i string
    	input (default "test.svg")
  -s float
    	seconds (default 1)
```

## Development

### Docker build

```
docker build -t svg2png:latest .
```

### Write code with hot reload

```
docker run --rm -it -v $(pwd):/go/src/app -w /go/src/app svg2png:latest air -c air.toml
```

### Build for Apple Silicon

```
docker run --rm -it -v $(pwd):/go/src/app -w /go/src/app -e GOOS=darwin -e GOARCH=arm64 golang:1.22.0-bookworm sh -c 'go mod 
tidy && go build -o smil2png.arm64.apple'
```