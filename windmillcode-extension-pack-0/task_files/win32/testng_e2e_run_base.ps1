Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
    [string] $envVarsScript ="",
    [string]  $testNGFolder ="",
    [string] $suiteFile="",
    [string] $paramEnv=""
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation



if ( $envVarsScript -eq "") {
    $defaultVar = "ignore\Windmillcode\testng_e2e_shared.ps1"
    $envVarsScript  = Read-Host -Prompt "script where env vars are set for the app to run relative to workspace root (leave empty to default to ignore\Windmillcode\testng_e2e_run.ps1 )"
    if ( $envVarsScript -eq "") {
      $envVarsScript = $defaultVar
    }
}

if ( $testNGFolder -eq "") {
    $defaultVar = "apps\testing\testng"
    $myPrompt =  "testng app location(leave empty to default to " + $defaultVar
    $testNGFolder  = Read-Host -Prompt $myPrompt
    if ( $testNGFolder -eq "") {
      $testNGFolder = $defaultVar
    }
}

if ( $suiteFile -eq "") {
    $defaultVar = "src\test\resources\tests.xml"
    $myPrompt =  "xml suite file needed for testng(leave empty to default to " + $defaultVar
    $suiteFile  = Read-Host -Prompt $myPrompt
    if ( $suiteFile -eq "") {
      $suiteFile = $defaultVar
    }
}

if ( $paramEnv -eq "") {
    $defaultVar = "DEV"
    $myPrompt =  "the envionrment to test ( valid options are PROD,PREVIEW,DEV, leave empty to default to " + $defaultVar
    $paramEnv  = Read-Host -Prompt $myPrompt
    if ( $paramEnv -eq "") {
      $paramEnv = $defaultVar
    }
}




