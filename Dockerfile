FROM golang:latest

# Install necessary packages for OpenGL and X11 support
RUN apt-get update && apt-get install -y \
    libgl1-mesa-dev \
    xorg-dev \
    libxcursor-dev \
    libxi-dev \
    libxinerama-dev \
    libxrandr-dev \
    xauth \
    x11-apps \
    git \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o ScaleSync


ENV DISPLAY=:0

CMD ["./ScaleSync"]
