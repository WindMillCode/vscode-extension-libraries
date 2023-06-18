Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
    [Parameter]  $envVarsLocation ,
    [Parameter]  $pythonVersion
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
$currentScript = $PSScriptRoot + '\flask_backend_run.ps1'
. $utilsFile;

if ( $envVarsLocation -eq $null) {
    $myPrompt = "where are the env vars located (choose empty string if the app does not have ENVIROMENT VARIABLES):"
    $myOptions = @(
        ".\ignore\Windmillcode\flask_backend_run.ps1",
        "None"
    )

    $programEnvVarsLocation = Show-Menu -Prompt $myPrompt -Options $myOptions
}
if ( $pythonVersion -eq $null) {
    $programPythonVersion  = Read-Host -Prompt "provide a python version for pyenv to use"
}

while ($true){
    cd $workspaceLocation;
    if ( -not($programEnvVarsLocation -eq "None")) {
        Invoke-Expression $programEnvVarsLocation
    }
    cd $workspaceLocation
    if( -not($programPythonVersion -eq "") ){
        pyenv shell $programPythonVersion
    }
    Set-Location apps\backend\FlaskApp;
    python app.py;

}

