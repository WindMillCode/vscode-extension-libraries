Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\'
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
$tasksJsonFile =  $PSScriptRoot + '\..\tasks.json'
$windowsScriptsFolder = $PSScriptRoot
. $utilsFile;

cd $workspaceLocation

try {
    Write-Warning "this will delete your vscode/tasks.json in your workspace folder if you dont have a .vscode/tasks.json and you have not used this command before it is safe to pick TRUE otherwise consult with your manager before continuing"
    $myPrompt = "replace tasks.json?"
    $myOptions = @(
        "TRUE",
        "FALSE"
    )
    $proceed =  Show-Menu -Prompt $myPrompt -Options $myOptions
    if( $proceed -eq "FALSE"){
        exit 0;
    }
    cd .vscode
    Remove-Item -Path "tasks.json" -ErrorAction 'SilentlyContinue'
    cp $tasksJsonFile .

    cd ..\ignore
    Remove-Item -Path "Windmillcode" -Recurse -ErrorAction 'SilentlyContinue'
    mkdir Windmillcode
    cp  -r $windowsScriptsFolder Windmillcode
}
catch {
    Write-Host "An error occurred: $($_.Exception.Message)"
    exit 1
}
