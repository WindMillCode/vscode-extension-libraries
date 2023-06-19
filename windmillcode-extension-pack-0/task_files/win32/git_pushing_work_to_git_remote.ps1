Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\',
    [string] $repoLocation ="",
    [string] $commitType ="" ,
    [string] $commitMsg =""
)



$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;



if ( $repoLocation -eq "") {
    $myPrompt = "Choose an option:"
    $myOptions = @(
        ".",
        ".\apps\frontend\AngularApp",
        ".\apps\backend\RailsApp",
        ".\apps\backend\FlaskApp"
    )

    $repoLocation = Show-Menu -Prompt $myPrompt -Options $myOptions
}

if ( $commitType -eq "") {
    $myPrompt = "Choose an option:"
    $myOptions = @("UPDATE", "FIX", "PATCH", "BUG", "MERGE", "COMPLEX MERGE","CHECKPOINT")

    $commitType =  Show-Menu -Prompt $myPrompt -Options $myOptions
}

if ( $commitMsg -eq "") {
    $commitMsg = Read-Host -Prompt "Enter your commit msg: additional work"
    if ( $commitMsg -eq "") {
        $commitMsg =  "additional work"
    }
}


echo $repoLocation
echo $commitType
echo $commitMsg

cd $workspaceLocation
cd $repoLocation

git add .;
$commitCommand = "git commit -m '$commitType $commitMsg'; git branch --unset-upstream; git push origin HEAD;"
Invoke-Expression -Command $commitCommand

