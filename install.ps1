# =============================================================================
# Devicebase CLI Installation Script (Windows PowerShell)
# Usage: irm https://git.uusense.cn/devicebase/devicebase-cli/releases/latest/download/install.ps1 | iex
# =============================================================================

$ErrorActionPreference = "Stop"

# Configuration
$Version = if ($env:VERSION) { $env:VERSION } else { "v2026.5.15" }
$BaseUrl = if ($env:BASE_URL) { $env:BASE_URL } else { "https://downloads.devicebase.cn/cli/releases" }
$InstallDir = if ($env:INSTALL_DIR) { $env:INSTALL_DIR } else { "$env:LOCALAPPDATA\Devicebase" }
$BinaryName = "devicebase"
$ChecksumUrl = if ($env:CHECKSUM_URL) { $env:CHECKSUM_URL } else { "" }

# Helper functions
function Write-Info {
    param([string]$Message)
    Write-Host "[INFO] " -ForegroundColor Green -NoNewline
    Write-Host $Message
}

function Write-Warn {
    param([string]$Message)
    Write-Host "[WARN] " -ForegroundColor Yellow -NoNewline
    Write-Host $Message
}

function Write-Err {
    param([string]$Message)
    Write-Host "[ERROR] " -ForegroundColor Red -NoNewline
    Write-Host $Message
}

# Detect architecture
function Get-Arch {
    $arch = $env:PROCESSOR_ARCHITECTURE
    if ($arch -eq "AMD64") { return "amd64" }
    if ($arch -eq "ARM64") { return "arm64" }
    if ($arch -eq "ARM32") { return "arm" }
    if ($arch -eq "x86")   { return "386" }

    # Fallback: check via .NET runtime
    try {
        $procArch = [System.Runtime.InteropServices.RuntimeInformation]::ProcessArchitecture
        if ($procArch -eq [System.Runtime.InteropServices.Architecture]::Arm64) { return "arm64" }
        if ($procArch -eq [System.Runtime.InteropServices.Architecture]::Arm)   { return "arm" }
        if ($procArch -eq [System.Runtime.InteropServices.Architecture]::X64)   { return "amd64" }
        if ($procArch -eq [System.Runtime.InteropServices.Architecture]::X86)   { return "386" }
    }
    catch {
        # .NET fallback not available, use pointer size
    }

    $ptrSize = [System.IntPtr]::Size
    if ($ptrSize -eq 8) { return "amd64" }
    return "386"
}

# Get latest version number
function Get-LatestVersion {
    param([string]$CurrentVersion)

    if ($CurrentVersion -ne "latest") {
        return $CurrentVersion
    }

    # GitHub API endpoint for latest release
    $releaseUrl = "https://api.github.com/repos/uusense/devicebase-cli/releases/latest"

    try {
        $response = Invoke-WebRequest -Uri $releaseUrl -UseBasicParsing -ErrorAction Stop
        $json = $response.Content | ConvertFrom-Json
        $tag = $json.tag_name
        if (-not [string]::IsNullOrEmpty($tag)) {
            return $tag
        }
    }
    catch {
        # API may fail due to rate limits; fall back to HTML scraping
        try {
            $htmlUrl = "$BaseUrl"
            $html = Invoke-WebRequest -Uri $htmlUrl -UseBasicParsing -ErrorAction Stop
            $match = [regex]::Match($html.Content, '(?<=<a href="/uusense/devicebase-cli/releases/tag/)[^"]+')
            if ($match.Success) {
                return $match.Value
            }
        }
        catch {
            # Fall through to error
        }
    }

    Write-Err "Failed to fetch latest version"
    exit 1
}

# Build download URL
function Get-DownloadUrl {
    param(
        [string]$Arch,
        [string]$Ver
    )
    return "$BaseUrl/download/$Ver/$BinaryName-windows-$Arch.exe"
}

# Verify checksum
function Confirm-Checksum {
    param(
        [string]$FilePath,
        [string]$ExpectedChecksum
    )

    if ([string]::IsNullOrEmpty($ExpectedChecksum)) {
        Write-Warn "No checksum provided, skipping verification"
        return
    }

    Write-Info "Verifying checksum ..."

    # Detect algorithm from checksum length (match bash script logic)
    $algo = switch ($ExpectedChecksum.Length) {
        64  { "SHA256" }
        128 { "SHA512" }
        default { "SHA256" }
    }

    $hash = (Get-FileHash -Path $FilePath -Algorithm $algo).Hash

    if ($hash -ne $ExpectedChecksum.ToUpper()) {
        Write-Err "Checksum mismatch! Expected: $ExpectedChecksum, Got: $hash"
        exit 1
    }

    Write-Info "Checksum verified ($algo)."
}

# Ensure install directory exists and add to PATH
function Install-Binary {
    param(
        [string]$SourcePath,
        [string]$DestDir
    )

    # Create install directory if needed
    if (-not (Test-Path -Path $DestDir)) {
        Write-Info "Creating install directory: $DestDir"
        New-Item -ItemType Directory -Path $DestDir -Force | Out-Null
    }

    $dest = Join-Path $DestDir "$BinaryName.exe"

    Write-Info "Installing to $dest ..."

    # Copy binary
    Copy-Item -Path $SourcePath -Destination $dest -Force

    # Add to user PATH if not already present
    $userPath = [Environment]::GetEnvironmentVariable("Path", "User")
    if ([string]::IsNullOrEmpty($userPath)) {
        Write-Info "Adding $DestDir to user PATH ..."
        [Environment]::SetEnvironmentVariable("Path", $DestDir, "User")
        $env:Path = "$env:Path;$DestDir"
    }
    elseif ($userPath -notlike "*$DestDir*") {
        Write-Info "Adding $DestDir to user PATH ..."
        $newPath = if ($userPath.EndsWith(";")) { "$userPath$DestDir" } else { "$userPath;$DestDir" }
        [Environment]::SetEnvironmentVariable("Path", $newPath, "User")
        $env:Path = "$env:Path;$DestDir"
    }

    Write-Info "Installed successfully: $dest"
}

# Main
function Main {
    Write-Info "Devicebase CLI Installer"
    Write-Info "Version: $Version"
    Write-Host ""

    # Detect platform
    $arch = Get-Arch
    Write-Info "Detected: windows/$arch"

    # Resolve version
    $Version = Get-LatestVersion -CurrentVersion $Version
    Write-Info "Installing version: $Version"
    Write-Host ""

    # Build download URL
    $url = Get-DownloadUrl -Arch $arch -Ver $Version
    $tmpDir = Join-Path $env:TEMP "devicebase-install-$(Get-Random)"
    $tmpFile = Join-Path $tmpDir "$BinaryName.exe"

    try {
        # Create temp directory
        New-Item -ItemType Directory -Path $tmpDir -Force | Out-Null

        # Download
        Write-Info "Downloading $url ..."
        Invoke-WebRequest -Uri $url -OutFile $tmpFile -UseBasicParsing
        if (-not (Test-Path $tmpFile) -or (Get-Item $tmpFile).Length -eq 0) {
            Write-Err "Download failed: file is empty or missing"
            exit 1
        }

        # Verify checksum if provided
        if (-not [string]::IsNullOrEmpty($ChecksumUrl)) {
            $checksumFile = Join-Path $tmpDir "checksum.txt"
            Invoke-WebRequest -Uri $ChecksumUrl -OutFile $checksumFile -UseBasicParsing

            $pattern = "windows-$arch"
            $checksumLine = Get-Content $checksumFile | Where-Object { $_ -match $pattern } | Select-Object -First 1
            $expectedHash = if ($checksumLine) { ($checksumLine -split '\s+')[0] } else { "" }

            Confirm-Checksum -FilePath $tmpFile -ExpectedChecksum $expectedHash
        }

        # Install
        Install-Binary -SourcePath $tmpFile -DestDir $InstallDir

        Write-Host ""
        Write-Info "Devicebase CLI installed successfully!"
        Write-Info "Run 'devicebase --help' to get started."
        Write-Info "Don't forget to set DEVICEBASE_API_KEY environment variable."
        Write-Host ""
        Write-Warn "You may need to restart your terminal for PATH changes to take effect."
    }
    finally {
        # Cleanup temp directory
        if (Test-Path $tmpDir) {
            Remove-Item -Path $tmpDir -Recurse -Force -ErrorAction SilentlyContinue
        }
    }
}

Main
