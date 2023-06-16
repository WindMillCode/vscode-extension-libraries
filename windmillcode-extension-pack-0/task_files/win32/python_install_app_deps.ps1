Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
    [string] $appLocation =".\apps\backend\FlaskApp",
    [string] $reinstall ="false",
    [string] $pythonVersion =""

)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation

if ( $appLocation -eq "") {
    $myPrompt = "Choose an option:"
    $myOptions = @(
        ".\apps\backend\FlaskApp"
    )

    $appLocation =  Show-Menu -Prompt $myPrompt -Options $myOptions
    if ( $appLocation -eq "") {
      $appLocation = ".\apps\backend\FlaskApp"
    }
}

if ( $pythonVersion -eq "") {
    $pythonVersion  = Read-Host -Prompt "provide a python version for pyenv to use"
}



$myPrompt = "reinstall?"
$myOptions = @("true" , "false")
$reinstall =  Show-Menu -Prompt $myPrompt -Options $myOptions
if ( $reinstall -eq "") {
  $reinstall = "false"
}




cd  $appLocation
if ( $reinstall -eq "true"){
    rm -r .\\site-packages\\windows
}
if( -not($pythonVersion -eq "") ){
  pyenv shell $pythonVersion
}
pip install -r windows-requirements.txt --upgrade --target .\\site-packages\\windows
