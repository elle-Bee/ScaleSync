FROM golang:1.23

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y \
    libx11-dev \
    libxtst-dev \
    libxrender-dev \
    libfontconfig1 \
    dbus-x11 \
    libxext6 \
    libxi6 \
    libglib2.0-0 \
    libsm6 \
    mesa-common-dev \
    libgl1-mesa-dev \
    libglu1-mesa-dev \
    freeglut3-dev && \
    rm -rf /var/lib/apt/lists/*

ENV DISPLAY=host.docker.internal:0
CMD ["go", "run", "main.go"]
