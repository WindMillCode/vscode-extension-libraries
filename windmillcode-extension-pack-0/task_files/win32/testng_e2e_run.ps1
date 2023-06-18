Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation=""
    [string] $envVarsScript ="",
    [string] $suiteFile=""
    [string] $paramEnv=""
    [string] $E2EDir=""
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation


if ( $envVarsScript -eq "") {
    $envVarsScript  = Read-Host -Prompt "script where env vars are set for the app to run relative to workspace root"
    if ( $envVarsScript -eq "") {
      $envVarsScript = "apps\testing\testng"
    }
}


Invoke-Expression $envVarsScript
