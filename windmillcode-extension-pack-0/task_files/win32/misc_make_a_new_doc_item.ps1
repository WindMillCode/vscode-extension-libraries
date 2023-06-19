Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\',
    [string] $docLocation ="",
    [string] $targetName =""
)
$ErrorActionPreference = "Stop";

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation


try {
if ( $docLocation -eq "") {
    $myPrompt = "Choose an option:"
    $myOptions = @(
      "docs\tutorials\",
      "docs\tasks_docs\",
      "docs\application_documentation\",
      "issues\"
    )

    $docLocation =  Show-Menu -Prompt $myPrompt -Options $myOptions
}


if ( $targetName -eq "") {

  $targetName = Read-Host -Prompt "enter the name of the entity PLEASE USE DASHES OR UNDERLINE FOR SPACES"
  if ($targetName -match "\s") {
      throw "The document name cannot contain any speaces PLEASE USE DASHES OR UNDERLINE FOR SPACES !!!!!!!!!     :)"
  }

}


$targetPath = $docLocation + $targetName
$templatePath = $docLocation + "\template"
cp -r $templatePath $targetPath
code $targetPath
}
catch {
    Write-Host "An error occurred: $($_.Exception.Message)"
    exit 1
}
