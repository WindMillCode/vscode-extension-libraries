Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
    [string] $runWithCache ="false"

)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation


$myPrompt = "run with cache?"
$myOptions = @("true" , "false")
$runWithCache =  Show-Menu -Prompt $myPrompt -Options $myOptions
if ( $runWithCache -eq "") {
  $runWithCache = "false"
}

cd .\\apps\\frontend\\AngularApp\\;
if ( $runWithCache -eq "false") {
  Remove-Item -Path '.\\.angular' -Recurse -Force;
}
npx ng serve --ssl=true --ssl-key=$env:WML_CERT_KEY0 --ssl-cert=$env:WML_CERT0
