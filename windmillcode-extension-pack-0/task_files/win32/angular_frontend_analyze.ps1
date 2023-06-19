Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\',
    [string] $envType =""
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation


if ( $envType -eq "") {
    $myPrompt = "Choose an option:"
    $myOptions = @("dev","preview","prod")

    $envType =  Show-Menu -Prompt $myPrompt -Options $myOptions
}
cd .\\apps\\frontend\\AngularApp\\;
yarn analyze:$envType
