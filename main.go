package main

import (
    "fmt"
    "math/rand"
    "os"
    "path/filepath"
    "strings"
    "time"

    "github.com/joho/godotenv"
)

const (
    grubConfigPath = "/etc/default/grub"
    logoFilename   = "logo.png"
    // Images must already be converted to the correct format/dimensions.
    // Use the .grub.png suffix convention to identify prepared images.
    imageSuffix = ".grub.png"
)

func main() {
    if len(os.Args) < 2 {
        fmt.Fprintf(os.Stderr, "Usage: %s <images-dir>\n", os.Args[0])
        fmt.Fprintf(os.Stderr, "  images-dir: directory containing prepared *%s files\n", imageSuffix)
        os.Exit(1)
    }

    imagesDir := os.Args[1]

    env, err := godotenv.Read(grubConfigPath)
    if err != nil {
        fatalf("Failed to read %s: %v", grubConfigPath, err)
    }

    themePath, ok := env["GRUB_THEME"]
    if !ok || strings.TrimSpace(themePath) == "" {
        fatalf("GRUB_THEME is not set in %s", grubConfigPath)
    }

    // e.g. /usr/share/grub/themes/catppuccin-mocha-grub-theme/theme.txt
    //   → /usr/share/grub/themes/catppuccin-mocha-grub-theme/
    themeDir := filepath.Dir(themePath)
    logoTarget := filepath.Join(themeDir, logoFilename)

    fmt.Printf("Active theme dir : %s\n", themeDir)
    fmt.Printf("Logo target      : %s\n", logoTarget)

    // Verify the theme directory actually exists
    if _, err := os.Stat(themeDir); os.IsNotExist(err) {
        fatalf("Theme directory does not exist: %s", themeDir)
    }

    entries, err := os.ReadDir(imagesDir)
    if err != nil {
        fatalf("Failed to read images directory %s: %v", imagesDir, err)
    }

    var images []string
    for _, e := range entries {
        if !e.IsDir() && strings.HasSuffix(e.Name(), imageSuffix) {
            images = append(images, filepath.Join(imagesDir, e.Name()))
        }
    }

    if len(images) == 0 {
        fatalf("No *%s files found in %s", imageSuffix, imagesDir)
    }

    fmt.Printf("Found %d image(s)\n", len(images))

    rng := rand.New(rand.NewSource(time.Now().UnixNano()))
    chosen := images[rng.Intn(len(images))]
    fmt.Printf("Selected         : %s\n", chosen)

    backupPath := logoTarget + ".bak"
    if _, err := os.Stat(backupPath); os.IsNotExist(err) {
        if err := copyFile(logoTarget, backupPath); err != nil {
            fatalf("Failed to back up original logo: %v", err)
        }
        fmt.Printf("Backed up original logo → %s\n", backupPath)
    }

    if err := copyFile(chosen, logoTarget); err != nil {
        fatalf("Failed to copy logo: %v", err)
    }

    fmt.Printf("Done — GRUB logo updated to %s\n", filepath.Base(chosen))
}

// copyFile copies src to dst, preserving permissions of dst if it exists.
func copyFile(src, dst string) error {
    data, err := os.ReadFile(src)
    if err != nil {
        return fmt.Errorf("read %s: %w", src, err)
    }

    // Use 0644 as default; if dst already exists keep its mode
    mode := os.FileMode(0644)
    if info, err := os.Stat(dst); err == nil {
        mode = info.Mode()
    }

    if err := os.WriteFile(dst, data, mode); err != nil {
        return fmt.Errorf("write %s: %w", dst, err)
    }
    return nil
}

func fatalf(format string, args ...any) {
    fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", args...)
    os.Exit(1)
}

