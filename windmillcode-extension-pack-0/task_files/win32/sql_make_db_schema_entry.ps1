

Param (
    [Parameter(Mandatory=$true)]
    [string] $workspaceLocation = "",
    [string] $databaseBackupLocation = ""
);

$ErrorActionPreference = "Stop";

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;


if ( $databaseBackupLocation -eq "") {
    $defaultVar = "apps\database\mysql"
    $databaseBackupLocation = Read-Host -Prompt "Enter the database script location (leave blank for $defaultVar (dont provide an answer you will be able to configure this from settings.json soon)  )"
    if ( $databaseBackupLocation -eq "") {
        $databaseBackupLocation =  $defaultVar
    }
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
