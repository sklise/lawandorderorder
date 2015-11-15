FROM golang
RUN apt-get update && apt-get install -y --no-install-recommends \
  libtesseract-dev \
  tesseract-ocr-dev \
  libleptonica-dev \
  && rm -rf /var/lib/apt/lists/*
# Get tesseract data
RUN git clone https://github.com/tesseract-ocr/tessdata.git /usr/lib/tessdata
ENV TESSDATA_PREFIX=/usr/lib/tessdata
