

Param (
    [Parameter(Mandatory=$true)]
    [string] $workspaceLocation = "",
    [string] $databaseToBackup = ""
);

$ErrorActionPreference = "Stop";

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;


if ( $databaseToBackup -eq "") {
    $defaultVar = "mysql"
    $databaseToBackup = Read-Host -Prompt "Enter the database script location (default is mysql refer to the folder in apps\database for your project)"
    if ( $databaseToBackup -eq "") {
      $databaseToBackup =  $defaultVar
    }
    $databaseBackupLocation =  "apps\database\"+$databaseToBackup+"\schema_entries"
}

cd $workspaceLocation
cd $databaseBackupLocation
try {
  $MyEnvs = "dev" , "preview" , "prod"
  foreach ($MyEnv in $MyEnvs) {
    cd $workspaceLocation
    cd $databaseBackupLocation
    cd $MyEnv;
    $currentday = Get-Date -Format "M-dd-yy\_hh-mm-ss";
    Write-Host $currentday;
    New-Item -ItemType Directory -Path $currentday ;
    Move-Item -Path "*.sql" -Destination "$currentday";

  }

}
catch {
    Write-Host "An error occurred: $($_.Exception.Message)"
    exit 1
}
