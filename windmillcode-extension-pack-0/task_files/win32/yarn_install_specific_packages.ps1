Param (
    [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\',
    [string] $appLocation ="",
    [string] $reinstall ="false",
    [string] $packageList="",
    [string] $depType=""

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


if ( $packageList -eq "") {
    $defaultVar = ""
    $packageList = Read-Host -Prompt "Provide the names of the packages you would like to install seperated by spaces"
     if ( $packageList -eq "") {
        throw "You must provide packages for installation"
         $packageList = $defaultVar
    }
}

if ( $depType -eq "") {
    $myPrompt = "chose whether its a dev dependency (-D) or dependency (-s)"
    $myOptions = @(
        "-D",
        "-s"
    )
    $depType =  Show-Menu -Prompt $myPrompt -Options $myOptions
}

$myPrompt = "reinstall?"
$myOptions = @("true" , "false")
$reinstall =  Show-Menu -Prompt $myPrompt -Options $myOptions
if ( $reinstall -eq "") {
  $reinstall = "false"
}




cd  $appLocation
if ( $reinstall -eq "true"){
    yarn remove  $packageList
}
yarn add $depType $packageList
