Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\'

)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation

try {
  cd .\\apps\\cloud\\FirebaseApp\\;
  npm run cleanup;
  npx firebase emulators:start --import='devData' --export-on-exit
}
catch {
    Write-Host "An error occurred: $($_.Exception.Message)"
    exit 1
}
