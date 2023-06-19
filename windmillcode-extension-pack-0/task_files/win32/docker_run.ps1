Param (
   [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\',
  [string] $dockerContainerName =""
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation

try {
  if ( $dockerContainerName -eq "") {
      $dockerContainerName = Read-Host -Prompt "the name of the container"
      if ( $dockerContainerName -eq "") {
          throw "you must provide a container to ssh into"
      }
  }

  docker start $dockerContainerName
}
catch {
    Write-Host "An error occurred: $($_.Exception.Message)"
    exit 1
}
