Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
    [string] $docLocation ="",
    [string] $targetName =""
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;

cd $workspaceLocation



if ( $docLocation -eq "") {
    $myPrompt = "Choose an option:"
    $myOptions = @(
      "docs\tutorials\",
      "docs\tasks_docs\",
      "docs\application_documentation\",
      "docs\issues\"
    )

    $docLocation =  Show-Menu -Prompt $myPrompt -Options $myOptions
}


if ( $targetName -eq "") {

  $targetName = Read-Host -Prompt "Enter your commit msg: additional work"
  if ($targetName -match "\s") {
      throw "The document name cannot contain any speaces PLEASE USE DASHES OR UNDERLINE FOR SPACES !!!!!!!!!     :)"
  }

}


$targetPath = $docLocation + $targetName
$templatePath = $docLocation + "\template"
cp -r $templatePath $targetPath
code $targetPath
