FROM golang:1.23

WORKDIR /app
COPY . .

RUN apt-get update && apt-get install -y \
    libx11-dev \
    libxtst-dev \
    libxrender-dev \
    libfontconfig1 \
    libxcursor-dev \
    libxrandr-dev \
    libxinerama-dev \
    libxi-dev \
    libglfw3-dev \
    dbus-x11 \
    libxext6 \
    libxi6 \
    libglib2.0-0 \
    libsm6 \
    mesa-common-dev \
    libgl1-mesa-dev \
    libglu1-mesa-dev \
    freeglut3-dev \
    libxxf86vm-dev \
    # Additional Libraries/Packages
    libcurl4-openssl-dev \         # For curl integration (common in many projects)
    libssl-dev \                   # For OpenSSL libraries
    build-essential \               # For building C++ applications
    python3-dev \                  # For Python3 development headers
    pkg-config \                   # For managing library paths
    clang \                        # For C/C++ language support
    cmake \                        # For building from source
    git \                          # Git for version control
    vim \                          # Vim text editor (optional)
    htop \                         # Monitoring tools for container
    unzip \                        # Unzip tools for archives
    wget \                         # For downloading files
    && rm -rf /var/lib/apt/lists/*  # Clean up to reduce image size

ENV DISPLAY=host.docker.internal:0
CMD ["go", "run", "main.go"]
