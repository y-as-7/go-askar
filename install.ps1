# askar Framework Tool
# This script handles both CLI installation and project setup for Windows PowerShell

$ErrorActionPreference = "Stop"

# ANSI Colors (PowerShell supports these in modern terminals)
$E = [char]27
$RED = "$E[0;31m"
$GREEN = "$E[0;32m"
$BLUE = "$E[0;34m"
$YELLOW = "$E[1;33m"
$NC = "$E[0m"

# Check if we are inside an existing askar project
$isProject = Test-Path "go.mod"
if ($isProject) {
    $content = Get-Content "go.mod" -Raw
    if ($content -notmatch "github.com/y-as-7/go-askar") {
        $isProject = $false
    }
}

if (-not $isProject) {
    # --- CLI INSTALLER MODE ---
    Write-Host "🚀 $($BLUE)askar CLI Installer (Windows)$($NC)"
    Write-Host "==============================="
    
    # 1. Install Binary
    Write-Host "📥 $($GREEN)Installing go-askar CLI v1.3.3...$($NC)"
    go install github.com/y-as-7/go-askar/cmd/go-askar@v1.3.3
    if ($LASTEXITCODE -ne 0) {
        Write-Host "$($RED)❌ Failed to install go-askar. Make sure Go is installed and in your PATH.$($NC)"
        exit 1
    }
    
    # 2. Configure PowerShell Profile (The Magic Trick)
    $magicFunc = @"

# askar CLI shortcut
function go-askar-alias {
    param([Parameter(ValueFromRemainingArguments = `$true)] `$args)
    if (`$args[0] -eq "askar") {
        `$rest = `$args[1..(`$args.Count-1)]
        go-askar `$rest
    } else {
        & go.exe `$args
    }
}
Set-Alias -Name go -Value go-askar-alias -Option AllScope -Force -ErrorAction SilentlyContinue
"@

    Write-Host "📝 $($GREEN)Configuring PowerShell profile...$($NC)"
    if (-not (Test-Path $PROFILE)) {
        New-Item -Path $PROFILE -Type File -Force | Out-Null
    }
    
    $profileContent = Get-Content $PROFILE -Raw
    if ($profileContent -notmatch "askar CLI shortcut") {
        Add-Content -Path $PROFILE -Value $magicFunc
        Write-Host "✅ $($GREEN)Added 'go askar' magic trick to your profile!$($NC)"
    }

    Write-Host ""
    Write-Host "✅ $($GREEN)CLI installed successfully!$($NC)"
    Write-Host ""
    Write-Host "$($YELLOW)IMPORTANT:$($NC) Please restart your PowerShell session or run:"
    Write-Host "   $($BLUE). `$PROFILE$($NC)"
    Write-Host ""
    Write-Host "Then you can start your project with:"
    Write-Host "   $($BLUE)go askar create/project my-app$($NC)"
    Write-Host ""
    exit 0
}

# --- PROJECT SETUP MODE (Post-Clone) ---
Write-Host "🚀 $($BLUE)askar Project Setup$($NC)"
Write-Host "==============================="
Write-Host ""

# Get project name
$projectName = $args[0]
if (-not $projectName) {
    $projectName = Read-Host "Enter your project name (e.g., my-shop)"
}

if (-not $projectName) {
    Write-Host "$($RED)❌ Project name cannot be empty$($NC)"
    exit 1
}

Write-Host "`n📦 $($GREEN)Setting up project: $projectName$($NC)`n"

# Update go.mod
Write-Host "📝 Updating go.mod..."
(Get-Content "go.mod") -replace "module github.com/y-as-7/go-askar", "module $projectName" | Set-Content "go.mod"

# Update all import paths
Write-Host "📝 Updating import paths..."
Get-ChildItem -Filter "*.go" -Recurse | ForEach-Object {
    (Get-Content $_.FullName) -replace "github.com/y-as-7/go-askar/", "$projectName/" | Set-Content $_.FullName
}

# Create .env
if (-not (Test-Path ".env")) {
    Write-Host "📝 Creating .env file..."
    Copy-Item ".env.example" ".env"
    $jwtSecret = [Guid]::NewGuid().ToString("N") + [Guid]::NewGuid().ToString("N")
    (Get-Content ".env") -replace "CHANGE_THIS_TO_RANDOM_SECRET", $jwtSecret | Set-Content ".env"
    Write-Host "✅ Generated random JWT_SECRET"
}

# Install dependencies
Write-Host "📥 Installing dependencies..."
go mod tidy

# Remove installers and .git
if (Test-Path ".git") {
    Write-Host "🧹 Cleaning up..."
    Remove-Item ".git" -Recurse -Force
}
Remove-Item "install.sh" -Force -ErrorAction SilentlyContinue
Remove-Item $MyInvocation.MyCommand.Path -Force -ErrorAction SilentlyContinue # Remove self

Write-Host "`n✅ $($GREEN)Project setup complete!$($NC)`n"
Write-Host "📋 $($BLUE)Next steps:$($NC)"
Write-Host "  1. Run: $($YELLOW)go run main.go$($NC) (or use make if installed)"
Write-Host "  2. API at: $($YELLOW)http://localhost:8080$($NC)`n"
Write-Host "🎉 $($GREEN)Happy coding with askar!$($NC)"
