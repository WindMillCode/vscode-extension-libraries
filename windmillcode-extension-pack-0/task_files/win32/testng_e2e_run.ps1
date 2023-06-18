Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation=""

)

try {
    $utilsFile = Join-Path $PSScriptRoot 'testng_e2e_run_base.ps1'
    . $utilsFile -workspaceLocation $workspaceLocation


    cd $workspaceLocation


    if ( -not($envVarsScript -eq "")) {
      Invoke-Expression $envVarsScript
    }

    cd $workspaceLocation
    cd  $testNGFolder
    mvn clean test -DsuiteFile="$suiteFile" -DparamEnv="$defaultVar"

}
catch {
    Write-Host "An error occurred: $($_.Exception.Message)"
    exit 1
}
