# dotfiles

This document include:
* How my repo is structured
* How branches work in my dotfiles layout
* How the zinit + plugins submodules work
* How to clone, update, bootstrap a new machine
* How to manage updates in the future
* Clear diagrams and commands

---

# 📘 **Dotfiles Repository Guide**

*A complete guide for future-me so I don’t have to remember any paths.*

---

# 🗂 **Repository Structure**

My GitHub repo contains **multiple branches**, each representing a section of my system config:

| Branch     | Purpose                                       |
| ---------- | --------------------------------------------- |
| `nvim`     | Neovim configuration                          |
| `zsh`      | Zsh configuration                             |
| `tmux`     | Tmux configuration                            |
| `anifetch` | Anifetch config                               |
| `neofetch` | Neofetch config                               |
| `zinit`    | Zinit + plugin manager & plugins              |
| `main`     | Might be empty / not used yet                 |

This layout lets me clone only the config I care about on any system.

---

# 📍 **Zinit Directory Layout**

Inside the `zinit` branch, my actual runtime directory looks like:

```
~/.local/share/zinit/
│
├── zinit.git/                                  <-- submodule
│
└── plugins/
     ├── Aloxaf---fzf-tab/                      <-- submodule
     ├── romkatv---powerlevel10k/               <-- submodule
     ├── tolkonepiu---catppuccin-powerlevel10k-themes/  <-- submodule
     ├── zsh-users---zsh-autosuggestions/       <-- submodule
     ├── zsh-users---zsh-completions/           <-- submodule
     └── zsh-users---zsh-syntax-highlighting/   <-- submodule
```

All plugins are **proper Git submodules**, not copied folders.

---

# 🧩 **Submodules Overview**

Each plugin is stored as:

```
plugins/<github-user>---<repo-name>/
```

They are always added like:

```
git submodule add <url> plugins/<folder-name>
```

This keeps my dotfiles repo small, fast, and easy to update.

---

# 🌱 **Cloning the zinit branch on a new machine**

This is how to bootstrap a fresh install:

```sh
git clone --branch zinit --recursive https://github.com/jelius-sama/dotfiles.git ~/.local/share/zinit
```

The `--recursive` is essential — it pulls all plugin submodules.

If I ever forget:

```sh
git submodule update --init --recursive
```

---

# 🔄 **Updating plugins later**

To update **all plugins**:

```sh
git submodule update --remote --merge
```

To update only one plugin:

```sh
git submodule update --remote plugins/zsh-users---zsh-autosuggestions
```

Then commit the update:

```sh
git add .
git commit -m "Update zinit plugins"
git push
```

---

# 🛠 **Adding a new plugin in the future**

1. Clone it manually into the correct folder:

   ```sh
   cd ~/.local/share/zinit/plugins
   git -C SOMEplugin remote get-url origin  # get URL
   ```

2. Remove the embedded repo:

   ```sh
   git rm --cached -r plugins/SOMEplugin
   ```

3. Add it as a submodule from the correct location:

   ```sh
   cd ~/.local/share/zinit
   git submodule add <url> plugins/SOMEplugin
   ```

4. Commit:

   ```sh
   git add .
   git commit -m "Add new zinit plugin"
   git push
   ```

---

# 🔥 **Resetting / Fixing Broken Submodules**

If a plugin gets messed up:

```sh
git submodule deinit -f plugins/SOMEplugin
rm -rf .git/modules/plugins/SOMEplugin
rm -rf plugins/SOMEplugin
git submodule add <url> plugins/SOMEplugin
```

---

# 🌎 **Branch Workflow on GitHub**

My branches represent *different components of my system*:

* `zinit` branch → only zsh plugin manager & plugins
* `zsh` branch → zshrc, env vars, prompt
* `nvim`, `tmux`, etc → each component isolated
* Makes merging unnecessary and avoids huge conflicts

When working on e.g. tmux:

```sh
git switch tmux
```

When working on zsh:

```sh
git switch zsh
```

When working on zinit:

```sh
git switch zinit
```

---

# 📦 **Best Practices Summary**

✔ Keep each config type in its own branch
✔ Always commit after submodule updates
✔ Use `--recursive` when cloning
✔ Avoid mixing branches

---
