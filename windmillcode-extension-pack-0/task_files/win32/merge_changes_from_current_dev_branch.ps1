Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation="",
    [string] $sourceBranch ="dev",
)




$utilsFile = $PSScriptRoot + '\utils.ps1'
. $utilsFile;


if ( $sourceBranch -eq "") {
    $sourceBranch = Read-Host -Prompt "the branch to merge changes from: "
    if ( $sourceBranch -eq "") {
        $sourceBranch =  "dev"
    }
}





cd $workspaceLocation

git checkout $sourceBranch ;
git pull origin $sourceBranch ;
git checkout -;
git merge $sourceBranch ;
