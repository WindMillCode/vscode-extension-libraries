Param (
  [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
  [string] $openAIAPIKey ="",
  [string] $langCodes=""
)

$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;



try {
  if ( $openAIAPIKey -eq "") {
    $openAIAPIKey = Read-Host -Prompt "provide the open ai api key"
    if ( $openAIAPIKey -eq "") {
       throw "an open ai key is required to translate the app"
    }
  }

  if ( $langCodes -eq "") {
    $myprompt =" Provide a list of lang codes to run \n translation script. \n Provide them in comma separated format according to the options below. \n Example: 'zh, es, hi, bn' \n It's best to do 4 at a time. \n Options: zh, es, hi, uk, ar, bn, ms, fr, de, sw"


    $langCodes = Read-Host -Prompt $myprompt;
    if ( $langCodes -eq "") {
       throw "Lang codes are required"
    }
  }

  Set-Location Env:
  Set-Content -Path OPENAI_API_KEY_0 -Value $openAIAPIKey


  $depsFolder = ".\site-packages\windows";
  cd  $PSScriptRoot;
  cd i18n_script_via_ai;
  if ( -not (Test-Path -Path $depsFolder -PathType Container)) {

    pip install -r requirements.txt --upgrade --target .\site-packages\windows
  }

  $i18nLocation = $workspaceLocation + "\apps\frontend\AngularApp\src\assets\i18n"
  cd $PSScriptRoot;
  python i18n_script_via_ai\index.py --lang-codes $langCodes --location $i18nLocation --source-file en.json

}
catch {
    Write-Host "An error occurred: $($_.Exception.Message)"
    exit 1
}
