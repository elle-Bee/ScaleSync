FROM golang:latest

RUN apt-get update && apt-get install -y \
    libgl1-mesa-glx \
    xorg-dev \
    libxcursor1 \
    libxi6 \
    libxinerama1 \
    libxrandr2 \
    xauth \
    git \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV DISPLAY=:0

CMD ["go", "run", "main.go"]