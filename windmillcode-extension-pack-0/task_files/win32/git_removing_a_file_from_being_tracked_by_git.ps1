Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
    [string] $targetType ="file",
    [string] $targetName ="" ,
)




$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;





if ( $targetType -eq "") {
    $myPrompt = "Is it a file or directory:"
    $myOptions = @("file","directory")

    $targetType =  Show-Menu -Prompt $myPrompt -Options $myOptions
}
if ( $targetName -eq "") {
    $targetName = Read-Host -Prompt "name of the item: "

}
cd $workspaceLocation

$command0 =""
if ( $targetType -eq "file") {
  $command0 = "git rm --cached  --sparse  $targetName"
}
else {
    $command0 = "git rm -r --cached  --sparse  $targetName"
}

Invoke-Expression -Command $command0

