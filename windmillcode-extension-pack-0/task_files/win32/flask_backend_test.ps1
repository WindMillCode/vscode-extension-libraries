Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
    [string] $envVarsLocation ="",
    [string] $pythonVersion=""
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

if ( $envVarsLocation -eq "") {
    $myPrompt = "where are the env vars located (choose empty string if the app does not have ENVIROMENT VARIABLES):"
    $myOptions = @(
        ".\ignore\Windmillcode\flask_backend_test.ps1",
        "None"
    )

    $envVarsLocation = Show-Menu -Prompt $myPrompt -Options $myOptions

}
if ( $pythonVersion -eq "") {
    $pythonVersion  = Read-Host -Prompt "provide a python version for pyenv to use"
}

cd $workspaceLocation

if ( -not($envVarsLocation -eq "None")) {
    Invoke-Expression $envVarsLocation
}
if( -not($pythonVersion -eq "") ){
  pyenv shell $pythonVersion
}
Set-Location apps\backend\FlaskApp
python run_tests.py

$currentScript = $PSScriptRoot + '\flask_backend_run.ps1'
Invoke-Command  $currentScript $workspaceLocation $envVarsLocation $pythonVersion
