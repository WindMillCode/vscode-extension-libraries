Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
    [string] $appLocation ="",
    [string] $reinstall ="false"

)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation

if ( $appLocation -eq "") {
    $myPrompt = "Choose an option:"
    $myOptions = @(
        ".\apps\frontend\AngularApp",
        ".\apps\cloud\FirebaseApp"
    )

    $appLocation =  Show-Menu -Prompt $myPrompt -Options $myOptions
    if ( $appLocation -eq "") {
      $appLocation = ".\apps\frontend\AngularApp"
    }
}



$myPrompt = "reinstall?"
$myOptions = @("true" , "false")
$reinstall =  Show-Menu -Prompt $myPrompt -Options $myOptions
if ( $reinstall -eq "") {
  $reinstall = "false"
}




cd  $appLocation
if ( $reinstall -eq "true"){
    rm yarn.lock;rm -r node_modules
}
yarn install
