FROM linuxserver/chromium:version-5c5f851f
RUN curl -L https://go.dev/dl/go1.22.0.linux-arm64.tar.gz > go1.22.0.linux-arm64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -zxvf go1.22.0.linux-arm64.tar.gz
ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=/go
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOBIN
RUN go install github.com/cosmtrek/air@latest && go install honnef.co/go/tools/cmd/staticcheck@latest && go install github.com/kisielk/errcheck@latest
ENTRYPOINT []
CMD ["/bin/bash"]
