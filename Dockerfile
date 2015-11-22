FROM golang
RUN apt-get update && apt-get install -y --no-install-recommends \
  libtesseract-dev \
  tesseract-ocr-dev \
  libleptonica-dev \
  autoconf automake build-essential libass-dev libfreetype6-dev \
  libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev \
  libxcb-xfixes0-dev pkg-config texinfo zlib1g-dev yasm libx264-dev libmp3lame-dev \
  && rm -rf /var/lib/apt/lists/*
# Get tesseract data
RUN git clone https://github.com/tesseract-ocr/tessdata.git /usr/lib/tessdata
ENV TESSDATA_PREFIX=/usr/lib/tessdata
RUN mkdir -p /go/src/github.com/sklise/lawandorder

ADD install_ffmpeg.sh /root/
RUN /root/install_ffmpeg.sh

ADD ./*.go /go/src/github.com/sklise/lawandorder/
WORKDIR /go/src/github.com/sklise/lawandorder
RUN go get -u gopkg.in/godo.v1/cmd/godo
RUN go get
