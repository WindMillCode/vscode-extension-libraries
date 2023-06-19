Param (
  [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
  [string] $dockerContainerName ="",
  [string] $databaseSoftwareName="",
  [string] $mysqlUsername = "",
  [string] $mysqlPassword =""
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;
cd $workspaceLocation

try {

  if ( $dockerContainerName -eq "") {
    $dockerContainerName = Read-Host -Prompt "the name of the docker container"
    if ( $dockerContainerName -eq "") {
       throw "a docker container must be provided"
    }
  }

  if ( $databaseSoftwareName -eq "") {
    $myPrompt = "select from one of the follow databases"
    $myOptions = @(
      "mysql",
      "postgres",
      "mssql",
      "oraclesql"
    )
    $databaseSoftwareName =  Show-Menu -Prompt $myPrompt -Options $myOptions
  }

  if( $databaseSoftwareName -eq "mysql"){

    if ( $mysqlUsername -eq "") {
      $mysqlUsername = Read-Host -Prompt "enter the mysql username (default is myAdmin)"
       if ( $mysqlUsername -eq "") {
         $mysqlUsername =  "myAdmin"
      }
    }

    if ( $mysqlPassword -eq "") {
      $defaultVar = "my-secret-pw"
      $mysqlPassword = Read-Host -Prompt "enter the mysql password (default is $defaultVar) "
       if ( $mysqlPassword -eq "") {
         $mysqlPassword = $defaultVar
      }
    }


    echo $mysqlPassword
    docker exec --workdir /root $dockerContainerName mysqldump  -u $mysqlUsername --password=$mysqlPassword --single-transaction --no-data --no-create-db windmillcodesite_mysql_database_0 > backup.sql
  }
  else{
    Write-Host "$databaseSoftwareName is not supported as of right now"
  }
}
catch {
  Write-Host "An error occurred: $($_.Exception.Message)"
  exit 1
}
