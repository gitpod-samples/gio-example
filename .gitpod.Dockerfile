# https://github.com/gitpod-io/workspace-images/tree/axonasif/add_docs/chunks/tool-vnc#workspace-full-vnc
FROM gitpod/workspace-full-vnc:latest

# https://gioui.org/doc/install/linux#dependencies
RUN sudo install-packages \
    gcc \
    pkg-config \
    libwayland-dev \
    libx11-dev \
    libx11-xcb-dev \
    libxkbcommon-x11-dev \
    libgles2-mesa-dev \
    libegl1-mesa-dev \
    libffi-dev \
    libxcursor-dev \
    libvulkan-dev

RUN go install gioui.org/cmd/gogio@latest