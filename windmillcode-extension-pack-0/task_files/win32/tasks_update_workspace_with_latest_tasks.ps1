Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\'
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
$tasksJsonFile =  $PSScriptRoot + '\..\tasks.json'
$windowsScriptsFolder = $PSScriptRoot
. $utilsFile;

function Update-TasksJson {
    try {
        Write-Warning "This will delete your vscode/tasks.json in your workspace folder. If you don't have a .vscode/tasks.json or you have not used this command before, it is safe to choose TRUE. Otherwise, consult with your manager before continuing."

        $myPrompt = "Replace tasks.json?"
        $myOptions = @("TRUE", "FALSE")
        $proceed = Show-Menu -Prompt $myPrompt -Options $myOptions

        if ($proceed -eq "FALSE") {
            return
        }

        Set-Location .vscode
        Remove-Item -Path "tasks.json" -ErrorAction 'SilentlyContinue'
        Copy-Item $tasksJsonFile .

        Set-Location ..\ignore
        Remove-Item -Path "Windmillcode" -Recurse -ErrorAction 'SilentlyContinue'
        mkdir Windmillcode
        Copy-Item -Recurse $windowsScriptsFolder Windmillcode
    }
    catch {
        Write-Host "An error occurred: $($_.Exception.Message)"
    }
}

Update-TasksJson;

Set-Location $workspaceLocation
Set-Location ignore\Windmillcode
Write-Host "Updating golang..."
$version = "1.20.6"
$installDir = $workspaceLocation + "\ignore\Windmillcode"

try {
    $executable = $installDir + "Go\bin\go.exe"
    . $executable version
}
catch {
    $url = "https://dl.google.com/go/go$version.windows-amd64.zip"
    $oldGoDir = $installDir +"\go"
    $godir = $installDir+ "\Go"

    if (-not (Test-Path $installDir)) {
        New-Item -ItemType Directory -Force -Path $installDir
        $zipFilePath = Join-Path $installDir "go$version.zip"
        Invoke-WebRequest -Uri $url -OutFile $zipFilePath


        Expand-Archive -Path $zipFilePath -DestinationPath $installDir
        Remove-Item -Path $zipFilePath
        Move-Item -Path $oldGoDir -Destination $goDir
    }



    $executable = $godir + "\bin\go.exe"
    . $executable version

}
