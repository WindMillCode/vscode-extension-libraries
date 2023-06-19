Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\'

)

cd $workspaceLocation
cd .\\apps\\frontend\\AngularApp\\; npx ng update
