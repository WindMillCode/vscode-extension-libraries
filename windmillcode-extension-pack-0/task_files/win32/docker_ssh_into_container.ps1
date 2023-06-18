Param (
  [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
  [string] $dockerContainerName ="",
  [string] $shell =""

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

  if ( $shell -eq "") {
      $myPrompt = "the command line shell (if uncertain select bash)"
      $myOptions = @(
        "sh",
        "bash",
        "dash",
        "zsh",
        "cmd",
        "fish",
        "ksh",
        "powershell"
      )

      $shell =  Show-Menu -Prompt $myPrompt -Options $myOptions 
  }

  docker exec -it $dockerContainerName $shell

}
catch {
    Write-Host "An error occurred: $($_.Exception.Message)"
    exit 1
}
