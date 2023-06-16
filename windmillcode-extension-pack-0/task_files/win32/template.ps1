Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation=""

)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation
