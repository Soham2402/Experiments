#!/bin/bash

# Create Neovim config directory
mkdir -p ~/.config/nvim

# Backup existing configuration if it exists
if [ -d ~/.config/nvim ] && [ "$(ls -A ~/.config/nvim)" ]; then
    echo "Backing up existing Neovim configuration..."
    mv ~/.config/nvim ~/.config/nvim.backup.$(date +%Y%m%d_%H%M%S)
    mkdir -p ~/.config/nvim
fi

# Clone jdhao's Neovim configuration
echo "Cloning jdhao's Neovim configuration..."
cd ~/.config/nvim
git clone --depth=1 https://github.com/jdhao/nvim-config.git .

# Create necessary directories for data and cache
mkdir -p ~/.local/share/nvim
mkdir -p ~/.cache/nvim

# Make scripts executable if any
find ~/.config/nvim -name "*.sh" -exec chmod +x {} \;

echo "Configuration installed successfully!"
echo "Path: ~/.config/nvim"

# Show what was installed
echo "Contents of ~/.config/nvim:"
ls -la ~/.config/nvim/

echo ""
echo "Configuration setup complete!"
echo "When you first open Neovim, plugins will be installed automatically."
echo "This may take a few minutes on first launch."
