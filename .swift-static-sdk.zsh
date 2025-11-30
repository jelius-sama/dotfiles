#!/usr/bin/env zsh

# Find the newest installed static Swift SDK inside ~/.swiftpm/swift-sdks
autoload -U colors && colors

SWIFT_SDK_DIR="$HOME/.swiftpm/swift-sdks"

if [[ ! -d "$SWIFT_SDK_DIR" ]]; then
    echo "$fg[red]✘ No static Swift SDK installed under $SWIFT_SDK_DIR$reset_color"
    return 1
fi

# Pick the latest artifactbundle
latest_bundle=("${SWIFT_SDK_DIR}"/*static-linux*.artifactbundle(N[1]))
if [[ ! -d "$latest_bundle" ]]; then
    echo "$fg[red]✘ No static Linux SDK artifactbundle found in $SWIFT_SDK_DIR$reset_color"
    return 1
fi

# There should be one folder inside the artifactbundle, grab it
bundle_root=("${latest_bundle}"/*(/N[1]))

# Detect musl sysroot automatically (e.g., musl-1.2.5.sdk/x86_64)
musl_sysroot=("${bundle_root}"/swift-linux-musl/*/x86_64(/N[1]))

if [[ ! -d "$musl_sysroot" ]]; then
    echo "$fg[red]✘ Failed to locate musl sysroot inside $bundle_root$reset_color"
    return 1
fi

export SWIFT_STATIC_SDK="$musl_sysroot"

# echo "$fg[green]✓ SWIFT_STATIC_SDK set to:$reset_color"
# echo "  $SWIFT_STATIC_SDK"
