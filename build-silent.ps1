# Set webhook URL. Script looks like:
# $env:webhookUrl='<webhook URL>'
.\set-env.ps1

# Inject value of $env:webhookUrl into executable at build time
go build -ldflags "-X main.webhookUrl=$env:webhookUrl -H windowsgui" -o silent.exe .