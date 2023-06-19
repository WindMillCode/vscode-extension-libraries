Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\',
    [string] $envVarsScript ="",
    [string] $pythonVersion=""
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

if ( $envVarsScript -eq "") {
    $myPrompt = "the file to run to set the env vars (choose empty string if the app does not have ENVIROMENT VARIABLES):"
    $myOptions = @(
        ".\ignore\Windmillcode\flask_backend_test.ps1",
        "None"
    )

    $envVarsScript = Show-Menu -Prompt $myPrompt -Options $myOptions

}
if ( $pythonVersion -eq "") {
    $pythonVersion  = Read-Host -Prompt "provide a python version for pyenv to use"
}

cd $workspaceLocation

if ( -not($envVarsScript -eq "None")) {
    Invoke-Expression $envVarsScript
}
if( -not($pythonVersion -eq "") ){
  pyenv shell $pythonVersion
}
Set-Location apps\backend\FlaskApp
python run_tests.py

$currentScript = $PSScriptRoot + '\flask_backend_run.ps1'
Invoke-Command  $currentScript $workspaceLocation $envVarsScript $pythonVersion
