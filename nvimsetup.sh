#!/bin/bash

# Update system packages
sudo apt update

# Install essential dependencies
sudo apt install -y \
    curl \
    wget \
    git \
    unzip \
    tar \
    gzip \
    build-essential \
    software-properties-common \
    python3 \
    python3-pip \
    python3-venv \
    nodejs \
    npm \
    ripgrep \
    fd-find \
    fzf \
    tree-sitter-cli

# Install Neovim (latest stable version)
# Method 1: Using snap (recommended for latest version)
sudo snap install nvim --classic

# Alternative Method 2: Using AppImage (if snap is not available)
# curl -LO https://github.com/neovim/neovim/releases/latest/download/nvim.appimage
# chmod u+x nvim.appimage
# sudo mv nvim.appimage /usr/local/bin/nvim

# Alternative Method 3: Using apt (might be older version)
# sudo apt install -y neovim

# Install language servers and formatters
# Python
pip3 install --user pynvim black isort flake8 mypy

# Node.js packages for LSP
sudo npm install -g \
    pyright \
    typescript \
    typescript-language-server \
    vscode-langservers-extracted \
    yaml-language-server \
    bash-language-server \
    dockerfile-language-server-nodejs

# Install additional tools
# LaTeX support (if needed)
# sudo apt install -y texlive-full latexmk

# Install ripgrep alternative name
sudo ln -sf /usr/bin/fdfind /usr/local/bin/fd

echo "Prerequisites installed successfully!"
echo "Neovim version:"
nvim --version