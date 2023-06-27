Param (
    [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\',
    [string]  $envVarsLocation="",
    [string]  $pythonVersion=""
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
$currentScript = $PSScriptRoot + '\flask_backend_run.ps1'
. $utilsFile;

if ( $envVarsLocation -eq "") {
    $myPrompt = "where are the env vars located (choose empty string if the app does not have ENVIROMENT VARIABLES):"
    $myOptions = @(
        ".\ignore\Local\flask_backend_run.ps1",
        "None"
    )

    $programEnvVarsLocation = Show-Menu -Prompt $myPrompt -Options $myOptions
}

if ( $pythonVersion -eq "") {
    $defaultVar = "3.11.4"
    $myPythonVersion = Read-Host -Prompt "provide a python version for pyenv to use (default and recommended is $defaultVar)"
     if ( $myPythonVersion -eq "") {
        $myPythonVersion = $defaultVar
    }
}

while ($true){
    cd $workspaceLocation;
    if ( -not($programEnvVarsLocation -eq "None")) {
        Invoke-Expression $programEnvVarsLocation
    }
    cd $workspaceLocation
    if( -not($programPythonVersion -eq "") ){
        pyenv shell $myPythonVersion
    }
    Set-Location apps\backend\FlaskApp;
    python app.py;

}

