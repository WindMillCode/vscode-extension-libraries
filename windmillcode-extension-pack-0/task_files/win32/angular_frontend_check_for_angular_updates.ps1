Param (
    [Parameter(Mandatory=$true)] [string] $workspaceLocation=""

)

cd $workspaceLocation
cd .\\apps\\frontend\\AngularApp\\; npx ng update
