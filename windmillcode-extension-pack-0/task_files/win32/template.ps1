Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\'

)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation

try {

}
catch {
    Write-Host "An error occurred: $($_.Exception.Message)"
    exit 1
}
