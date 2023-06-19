Param (
     [string] $workspaceLocation=$PSScriptRoot + '\..\..\..\'

)


cd $workspaceLocation
cd .\\apps\\zero\\frontend\\AngularApp\\;
yarn compodoc:build-and-serve
