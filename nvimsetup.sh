#!/bin/bash

set -e  # Exit immediately if a command fails
set -u  # Treat unset variables as error
set -o pipefail

echo "🔧 Starting Neovim configuration setup using jdhao/nvim-config..."

# ---- 1. Install system dependencies ----
echo "📦 Installing required system packages..."
sudo apt update
sudo apt install -y curl git unzip ripgrep python3-pip nodejs npm fd-find software-properties-common

# Ensure `fd` is accessible as expected by Telescope
mkdir -p ~/.local/bin
if ! command -v fd &> /dev/null; then
    echo "🔗 Linking fdfind to fd..."
    ln -sf "$(which fdfind)" ~/.local/bin/fd
    echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
    export PATH="$HOME/.local/bin:$PATH"
fi

# ---- 2. Install Neovim (latest stable) ----
if ! command -v nvim &> /dev/null || [[ "$(nvim --version | head -n1 | grep -o '[0-9]\.[0-9]\+' | head -n1)" < "0.9" ]]; then
    echo "📥 Installing latest Neovim..."
    sudo add-apt-repository -y ppa:neovim-ppa/unstable
    sudo apt update
    sudo apt install -y neovim
fi

# ---- 3. Install pynvim for Python and Node.js support ----
echo "🧠 Installing pynvim for Python and Node.js..."
pip3 install --user pynvim
npm install -g neovim

# ---- 4. Backup any existing nvim config ----
echo "📁 Backing up old Neovim config..."
CONFIG_DIR="$HOME/.config/nvim"
BACKUP_DIR="$HOME/.config/nvim_backup_$(date +%s)"
if [ -d "$CONFIG_DIR" ]; then
    mv "$CONFIG_DIR" "$BACKUP_DIR"
    echo "🗃️  Existing config moved to $BACKUP_DIR"
fi

# ---- 5. Clone jdhao/nvim-config ----
echo "🔻 Cloning jdhao/nvim-config..."
git clone https://github.com/jdhao/nvim-config.git "$CONFIG_DIR"

# ---- 6. Open Neovim to install plugins via packer ----
echo "🚀 Launching Neovim to install plugins..."
nvim --headless +PackerSync +qall

# ---- 7. Treesitter update (syntax highlighters) ----
echo "🌲 Updating Treesitter parsers..."
nvim --headless +TSUpdate +qall

# ---- 8. LSP installation via Mason ----
echo "🛠️ Setting up Mason LSP registry (manual steps recommended next)..."
echo "➡️ Open Neovim and run ':Mason' to install LSP servers for your languages."

# ---- 9. Final cleanup and summary ----
echo -e "\n✅ Neovim setup is complete with jdhao/nvim-config!"
echo "📂 Config installed at: $CONFIG_DIR"
echo "💡 Open Neovim and try commands like:"
echo "    - ;ff  → find files"
echo "    - ;fg  → live grep"
echo "    - gd   → go to definition"
echo "    - :Mason → install LSPs"
echo -e "\nHappy hacking! 🚀"
